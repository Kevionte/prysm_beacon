package client

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Kevionte/prysm_beacon/v2/beacon-chain/core/signing"
	fieldparams "github.com/Kevionte/prysm_beacon/v2/config/fieldparams"
	"github.com/Kevionte/prysm_beacon/v2/config/params"
	"github.com/Kevionte/prysm_beacon/v2/consensus-types/primitives"
	"github.com/Kevionte/prysm_beacon/v2/crypto/bls"
	"github.com/Kevionte/prysm_beacon/v2/monitoring/tracing"
	"github.com/Kevionte/prysm_beacon/v2/network/httputil"
	ethpb "github.com/Kevionte/prysm_beacon/v2/proto/prysm/v1alpha1"
	validatorpb "github.com/Kevionte/prysm_beacon/v2/proto/prysm/v1alpha1/validator-client"
	prysmTime "github.com/Kevionte/prysm_beacon/v2/time"
	"github.com/Kevionte/prysm_beacon/v2/time/slots"
	"github.com/pkg/errors"
	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// SubmitAggregateAndProof submits the validator's signed slot signature to the beacon node
// via gRPC. Beacon node will verify the slot signature and determine if the validator is also
// an aggregator. If yes, then beacon node will broadcast aggregated signature and
// proof on the validator's behalf.
func (v *validator) SubmitAggregateAndProof(ctx context.Context, slot primitives.Slot, pubKey [fieldparams.BLSPubkeyLength]byte) {
	ctx, span := trace.StartSpan(ctx, "validator.SubmitAggregateAndProof")
	defer span.End()

	span.AddAttributes(trace.StringAttribute("validator", fmt.Sprintf("%#x", pubKey)))
	fmtKey := fmt.Sprintf("%#x", pubKey[:])

	duty, err := v.duty(pubKey)
	if err != nil {
		log.WithError(err).Error("Could not fetch validator assignment")
		if v.emitAccountMetrics {
			ValidatorAggFailVec.WithLabelValues(fmtKey).Inc()
		}
		return
	}

	// Avoid sending beacon node duplicated aggregation requests.
	k := validatorSubscribeKey(slot, duty.CommitteeIndex)
	v.aggregatedSlotCommitteeIDCacheLock.Lock()
	if v.aggregatedSlotCommitteeIDCache.Contains(k) {
		v.aggregatedSlotCommitteeIDCacheLock.Unlock()
		return
	}
	v.aggregatedSlotCommitteeIDCache.Add(k, true)
	v.aggregatedSlotCommitteeIDCacheLock.Unlock()

	var slotSig []byte
	if v.distributed {
		slotSig, err = v.getAttSelection(attSelectionKey{slot: slot, index: duty.ValidatorIndex})
		if err != nil {
			log.WithError(err).Error("Could not find aggregated selection proof")
			if v.emitAccountMetrics {
				ValidatorAggFailVec.WithLabelValues(fmtKey).Inc()
			}
			return
		}
	} else {
		slotSig, err = v.signSlotWithSelectionProof(ctx, pubKey, slot)
		if err != nil {
			log.WithError(err).Error("Could not sign slot")
			if v.emitAccountMetrics {
				ValidatorAggFailVec.WithLabelValues(fmtKey).Inc()
			}
			return
		}
	}

	// As specified in spec, an aggregator should wait until two thirds of the way through slot
	// to broadcast the best aggregate to the global aggregate channel.
	// https://github.com/ethereum/consensus-specs/blob/v0.9.3/specs/validator/0_beacon-chain-validator.md#broadcast-aggregate
	v.waitToSlotTwoThirds(ctx, slot)

	res, err := v.validatorClient.SubmitAggregateSelectionProof(ctx, &ethpb.AggregateSelectionRequest{
		Slot:           slot,
		CommitteeIndex: duty.CommitteeIndex,
		PublicKey:      pubKey[:],
		SlotSignature:  slotSig,
	}, duty.ValidatorIndex, uint64(len(duty.Committee)))
	if err != nil {
		// handle grpc not found
		s, ok := status.FromError(err)
		grpcNotFound := ok && s.Code() == codes.NotFound
		// handle http not found
		jsonErr := &httputil.DefaultJsonError{}
		httpNotFound := errors.As(err, &jsonErr) && jsonErr.Code == http.StatusNotFound

		if grpcNotFound || httpNotFound {
			log.WithField("slot", slot).WithError(err).Warn("No attestations to aggregate")
		} else {
			log.WithField("slot", slot).WithError(err).Error("Could not submit aggregate selection proof to beacon node")
			if v.emitAccountMetrics {
				ValidatorAggFailVec.WithLabelValues(fmtKey).Inc()
			}
		}

		return
	}

	sig, err := v.aggregateAndProofSig(ctx, pubKey, res.AggregateAndProof, slot)
	if err != nil {
		log.WithError(err).Error("Could not sign aggregate and proof")
		return
	}
	_, err = v.validatorClient.SubmitSignedAggregateSelectionProof(ctx, &ethpb.SignedAggregateSubmitRequest{
		SignedAggregateAndProof: &ethpb.SignedAggregateAttestationAndProof{
			Message:   res.AggregateAndProof,
			Signature: sig,
		},
	})
	if err != nil {
		log.WithError(err).Error("Could not submit signed aggregate and proof to beacon node")
		if v.emitAccountMetrics {
			ValidatorAggFailVec.WithLabelValues(fmtKey).Inc()
		}
		return
	}

	if err := v.saveSubmittedAtt(res.AggregateAndProof.Aggregate.Data, pubKey[:], true); err != nil {
		log.WithError(err).Error("Could not add aggregator indices to logs")
		if v.emitAccountMetrics {
			ValidatorAggFailVec.WithLabelValues(fmtKey).Inc()
		}
		return
	}
	if v.emitAccountMetrics {
		ValidatorAggSuccessVec.WithLabelValues(fmtKey).Inc()
	}
}

// Signs input slot with domain selection proof. This is used to create the signature for aggregator selection.
func (v *validator) signSlotWithSelectionProof(ctx context.Context, pubKey [fieldparams.BLSPubkeyLength]byte, slot primitives.Slot) (signature []byte, err error) {
	domain, err := v.domainData(ctx, slots.ToEpoch(slot), params.BeaconConfig().DomainSelectionProof[:])
	if err != nil {
		return nil, err
	}

	var sig bls.Signature
	sszUint := primitives.SSZUint64(slot)
	root, err := signing.ComputeSigningRoot(&sszUint, domain.SignatureDomain)
	if err != nil {
		return nil, err
	}
	sig, err = v.keyManager.Sign(ctx, &validatorpb.SignRequest{
		PublicKey:       pubKey[:],
		SigningRoot:     root[:],
		SignatureDomain: domain.SignatureDomain,
		Object:          &validatorpb.SignRequest_Slot{Slot: slot},
		SigningSlot:     slot,
	})
	if err != nil {
		return nil, err
	}

	return sig.Marshal(), nil
}

// waitToSlotTwoThirds waits until two third through the current slot period
// such that any attestations from this slot have time to reach the beacon node
// before creating the aggregated attestation.
func (v *validator) waitToSlotTwoThirds(ctx context.Context, slot primitives.Slot) {
	ctx, span := trace.StartSpan(ctx, "validator.waitToSlotTwoThirds")
	defer span.End()

	oneThird := slots.DivideSlotBy(3 /* one third of slot duration */)
	twoThird := oneThird + oneThird
	delay := twoThird

	startTime := slots.StartTime(v.genesisTime, slot)
	finalTime := startTime.Add(delay)
	wait := prysmTime.Until(finalTime)
	if wait <= 0 {
		return
	}
	t := time.NewTimer(wait)
	defer t.Stop()
	select {
	case <-ctx.Done():
		tracing.AnnotateError(span, ctx.Err())
		return
	case <-t.C:
		return
	}
}

// This returns the signature of validator signing over aggregate and
// proof object.
func (v *validator) aggregateAndProofSig(ctx context.Context, pubKey [fieldparams.BLSPubkeyLength]byte, agg *ethpb.AggregateAttestationAndProof, slot primitives.Slot) ([]byte, error) {
	d, err := v.domainData(ctx, slots.ToEpoch(agg.Aggregate.Data.Slot), params.BeaconConfig().DomainAggregateAndProof[:])
	if err != nil {
		return nil, err
	}
	var sig bls.Signature
	root, err := signing.ComputeSigningRoot(agg, d.SignatureDomain)
	if err != nil {
		return nil, err
	}
	sig, err = v.keyManager.Sign(ctx, &validatorpb.SignRequest{
		PublicKey:       pubKey[:],
		SigningRoot:     root[:],
		SignatureDomain: d.SignatureDomain,
		Object:          &validatorpb.SignRequest_AggregateAttestationAndProof{AggregateAttestationAndProof: agg},
		SigningSlot:     slot,
	})
	if err != nil {
		return nil, err
	}

	return sig.Marshal(), nil
}
