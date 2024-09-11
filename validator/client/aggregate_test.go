package client

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/Kevionte/prysm_beacon/v2/validator/client/iface"

	fieldparams "github.com/Kevionte/prysm_beacon/v2/config/fieldparams"
	"github.com/Kevionte/prysm_beacon/v2/config/params"
	"github.com/Kevionte/prysm_beacon/v2/consensus-types/primitives"
	"github.com/Kevionte/prysm_beacon/v2/crypto/bls"
	ethpb "github.com/Kevionte/prysm_beacon/v2/proto/prysm/v1alpha1"
	"github.com/Kevionte/prysm_beacon/v2/testing/assert"
	"github.com/Kevionte/prysm_beacon/v2/testing/require"
	"github.com/Kevionte/prysm_beacon/v2/testing/util"
	"github.com/Kevionte/prysm_beacon/v2/time/slots"
	"github.com/prysmaticlabs/go-bitfield"
	logTest "github.com/sirupsen/logrus/hooks/test"
	"go.uber.org/mock/gomock"
)

func TestSubmitAggregateAndProof_GetDutiesRequestFailure(t *testing.T) {
	for _, isSlashingProtectionMinimal := range [...]bool{false, true} {
		t.Run(fmt.Sprintf("SlashingProtectionMinimal:%v", isSlashingProtectionMinimal), func(t *testing.T) {
			hook := logTest.NewGlobal()
			validator, _, validatorKey, finish := setup(t, isSlashingProtectionMinimal)
			validator.duties = &ethpb.DutiesResponse{CurrentEpochDuties: []*ethpb.DutiesResponse_Duty{}}
			defer finish()

			var pubKey [fieldparams.BLSPubkeyLength]byte
			copy(pubKey[:], validatorKey.PublicKey().Marshal())
			validator.SubmitAggregateAndProof(context.Background(), 0, pubKey)

			require.LogsContain(t, hook, "Could not fetch validator assignment")
		})
	}
}

func TestSubmitAggregateAndProof_SignFails(t *testing.T) {
	for _, isSlashingProtectionMinimal := range [...]bool{false, true} {
		t.Run(fmt.Sprintf("SlashingProtectionMinimal:%v", isSlashingProtectionMinimal), func(t *testing.T) {
			validator, m, validatorKey, finish := setup(t, isSlashingProtectionMinimal)
			defer finish()
			var pubKey [fieldparams.BLSPubkeyLength]byte
			copy(pubKey[:], validatorKey.PublicKey().Marshal())
			validator.duties = &ethpb.DutiesResponse{
				CurrentEpochDuties: []*ethpb.DutiesResponse_Duty{
					{
						PublicKey: validatorKey.PublicKey().Marshal(),
					},
				},
			}

			m.validatorClient.EXPECT().DomainData(
				gomock.Any(), // ctx
				gomock.Any(), // epoch
			).Return(&ethpb.DomainResponse{SignatureDomain: make([]byte, 32)}, nil /*err*/)

			m.validatorClient.EXPECT().SubmitAggregateSelectionProof(
				gomock.Any(), // ctx
				gomock.AssignableToTypeOf(&ethpb.AggregateSelectionRequest{}),
				gomock.Any(),
				gomock.Any(),
			).Return(&ethpb.AggregateSelectionResponse{
				AggregateAndProof: &ethpb.AggregateAttestationAndProof{
					AggregatorIndex: 0,
					Aggregate: util.HydrateAttestation(&ethpb.Attestation{
						AggregationBits: make([]byte, 1),
					}),
					SelectionProof: make([]byte, 96),
				},
			}, nil)

			m.validatorClient.EXPECT().DomainData(
				gomock.Any(), // ctx
				gomock.Any(), // epoch
			).Return(&ethpb.DomainResponse{SignatureDomain: nil}, errors.New("bad domain root"))

			validator.SubmitAggregateAndProof(context.Background(), 0, pubKey)
		})
	}
}

func TestSubmitAggregateAndProof_Ok(t *testing.T) {
	for _, isSlashingProtectionMinimal := range [...]bool{false, true} {
		t.Run(fmt.Sprintf("SlashingProtectionMinimal:%v", isSlashingProtectionMinimal), func(t *testing.T) {
			validator, m, validatorKey, finish := setup(t, isSlashingProtectionMinimal)
			defer finish()
			var pubKey [fieldparams.BLSPubkeyLength]byte
			copy(pubKey[:], validatorKey.PublicKey().Marshal())
			validator.duties = &ethpb.DutiesResponse{
				CurrentEpochDuties: []*ethpb.DutiesResponse_Duty{
					{
						PublicKey: validatorKey.PublicKey().Marshal(),
					},
				},
			}

			m.validatorClient.EXPECT().DomainData(
				gomock.Any(), // ctx
				gomock.Any(), // epoch
			).Return(&ethpb.DomainResponse{SignatureDomain: make([]byte, 32)}, nil /*err*/)

			m.validatorClient.EXPECT().SubmitAggregateSelectionProof(
				gomock.Any(), // ctx
				gomock.AssignableToTypeOf(&ethpb.AggregateSelectionRequest{}),
				gomock.Any(),
				gomock.Any(),
			).Return(&ethpb.AggregateSelectionResponse{
				AggregateAndProof: &ethpb.AggregateAttestationAndProof{
					AggregatorIndex: 0,
					Aggregate: util.HydrateAttestation(&ethpb.Attestation{
						AggregationBits: make([]byte, 1),
					}),
					SelectionProof: make([]byte, 96),
				},
			}, nil)

			m.validatorClient.EXPECT().DomainData(
				gomock.Any(), // ctx
				gomock.Any(), // epoch
			).Return(&ethpb.DomainResponse{SignatureDomain: make([]byte, 32)}, nil /*err*/)

			m.validatorClient.EXPECT().SubmitSignedAggregateSelectionProof(
				gomock.Any(), // ctx
				gomock.AssignableToTypeOf(&ethpb.SignedAggregateSubmitRequest{}),
			).Return(&ethpb.SignedAggregateSubmitResponse{AttestationDataRoot: make([]byte, 32)}, nil)

			validator.SubmitAggregateAndProof(context.Background(), 0, pubKey)
		})
	}
}

func TestSubmitAggregateAndProof_Distributed(t *testing.T) {
	validatorIdx := primitives.ValidatorIndex(123)
	slot := primitives.Slot(456)
	ctx := context.Background()
	for _, isSlashingProtectionMinimal := range [...]bool{false, true} {
		t.Run(fmt.Sprintf("SlashingProtectionMinimal:%v", isSlashingProtectionMinimal), func(t *testing.T) {
			validator, m, validatorKey, finish := setup(t, isSlashingProtectionMinimal)
			defer finish()

			var pubKey [fieldparams.BLSPubkeyLength]byte
			copy(pubKey[:], validatorKey.PublicKey().Marshal())
			validator.duties = &ethpb.DutiesResponse{
				CurrentEpochDuties: []*ethpb.DutiesResponse_Duty{
					{
						PublicKey:      validatorKey.PublicKey().Marshal(),
						ValidatorIndex: validatorIdx,
						AttesterSlot:   slot,
					},
				},
			}

			validator.distributed = true
			validator.attSelections = make(map[attSelectionKey]iface.BeaconCommitteeSelection)
			validator.attSelections[attSelectionKey{
				slot:  slot,
				index: 123,
			}] = iface.BeaconCommitteeSelection{
				SelectionProof: make([]byte, 96),
				Slot:           slot,
				ValidatorIndex: validatorIdx,
			}

			m.validatorClient.EXPECT().SubmitAggregateSelectionProof(
				gomock.Any(), // ctx
				gomock.AssignableToTypeOf(&ethpb.AggregateSelectionRequest{}),
				gomock.Any(),
				gomock.Any(),
			).Return(&ethpb.AggregateSelectionResponse{
				AggregateAndProof: &ethpb.AggregateAttestationAndProof{
					AggregatorIndex: 0,
					Aggregate: util.HydrateAttestation(&ethpb.Attestation{
						AggregationBits: make([]byte, 1),
					}),
					SelectionProof: make([]byte, 96),
				},
			}, nil)

			m.validatorClient.EXPECT().DomainData(
				gomock.Any(), // ctx
				gomock.Any(), // epoch
			).Return(&ethpb.DomainResponse{SignatureDomain: make([]byte, 32)}, nil /*err*/)

			m.validatorClient.EXPECT().SubmitSignedAggregateSelectionProof(
				gomock.Any(), // ctx
				gomock.AssignableToTypeOf(&ethpb.SignedAggregateSubmitRequest{}),
			).Return(&ethpb.SignedAggregateSubmitResponse{AttestationDataRoot: make([]byte, 32)}, nil)

			validator.SubmitAggregateAndProof(ctx, slot, pubKey)
		})
	}
}

func TestWaitForSlotTwoThird_WaitCorrectly(t *testing.T) {
	for _, isSlashingProtectionMinimal := range [...]bool{false, true} {
		t.Run(fmt.Sprintf("SlashingProtectionMinimal:%v", isSlashingProtectionMinimal), func(t *testing.T) {
			validator, _, _, finish := setup(t, isSlashingProtectionMinimal)
			defer finish()
			currentTime := time.Now()
			numOfSlots := primitives.Slot(4)
			validator.genesisTime = uint64(currentTime.Unix()) - uint64(numOfSlots.Mul(params.BeaconConfig().SecondsPerSlot))
			oneThird := slots.DivideSlotBy(3 /* one third of slot duration */)
			timeToSleep := oneThird + oneThird

			twoThirdTime := currentTime.Add(timeToSleep)
			validator.waitToSlotTwoThirds(context.Background(), numOfSlots)
			currentTime = time.Now()
			assert.Equal(t, twoThirdTime.Unix(), currentTime.Unix())
		})
	}
}

func TestWaitForSlotTwoThird_DoneContext_ReturnsImmediately(t *testing.T) {
	for _, isSlashingProtectionMinimal := range [...]bool{false, true} {
		t.Run(fmt.Sprintf("SlashingProtectionMinimal:%v", isSlashingProtectionMinimal), func(t *testing.T) {
			validator, _, _, finish := setup(t, isSlashingProtectionMinimal)
			defer finish()
			currentTime := time.Now()
			numOfSlots := primitives.Slot(4)
			validator.genesisTime = uint64(currentTime.Unix()) - uint64(numOfSlots.Mul(params.BeaconConfig().SecondsPerSlot))

			expectedTime := time.Now()
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			validator.waitToSlotTwoThirds(ctx, numOfSlots)
			currentTime = time.Now()
			assert.Equal(t, expectedTime.Unix(), currentTime.Unix())
		})
	}
}

func TestAggregateAndProofSignature_CanSignValidSignature(t *testing.T) {
	for _, isSlashingProtectionMinimal := range [...]bool{false, true} {
		t.Run(fmt.Sprintf("SlashingProtectionMinimal:%v", isSlashingProtectionMinimal), func(t *testing.T) {
			validator, m, validatorKey, finish := setup(t, isSlashingProtectionMinimal)
			defer finish()

			var pubKey [fieldparams.BLSPubkeyLength]byte
			copy(pubKey[:], validatorKey.PublicKey().Marshal())
			m.validatorClient.EXPECT().DomainData(
				gomock.Any(), // ctx
				&ethpb.DomainRequest{Epoch: 0, Domain: params.BeaconConfig().DomainAggregateAndProof[:]},
			).Return(&ethpb.DomainResponse{SignatureDomain: make([]byte, 32)}, nil /*err*/)

			agg := &ethpb.AggregateAttestationAndProof{
				AggregatorIndex: 0,
				Aggregate: util.HydrateAttestation(&ethpb.Attestation{
					AggregationBits: bitfield.NewBitlist(1),
				}),
				SelectionProof: make([]byte, 96),
			}
			sig, err := validator.aggregateAndProofSig(context.Background(), pubKey, agg, 0 /* slot */)
			require.NoError(t, err)
			_, err = bls.SignatureFromBytes(sig)
			require.NoError(t, err)
		})
	}
}
