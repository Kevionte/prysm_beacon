package state_native

import (
	"github.com/Kevionte/prysm_beacon/v2/beacon-chain/state/state-native/types"
	ethpb "github.com/Kevionte/prysm_beacon/v2/proto/prysm/v1alpha1"
	"github.com/Kevionte/prysm_beacon/v2/runtime/version"
)

// SetCurrentSyncCommittee for the beacon state.
func (b *BeaconState) SetCurrentSyncCommittee(val *ethpb.SyncCommittee) error {
	b.lock.Lock()
	defer b.lock.Unlock()

	if b.version == version.Phase0 {
		return errNotSupported("SetCurrentSyncCommittee", b.version)
	}

	b.currentSyncCommittee = val
	b.markFieldAsDirty(types.CurrentSyncCommittee)
	return nil
}

// SetNextSyncCommittee for the beacon state.
func (b *BeaconState) SetNextSyncCommittee(val *ethpb.SyncCommittee) error {
	b.lock.Lock()
	defer b.lock.Unlock()

	if b.version == version.Phase0 {
		return errNotSupported("SetNextSyncCommittee", b.version)
	}

	b.nextSyncCommittee = val
	b.markFieldAsDirty(types.NextSyncCommittee)
	return nil
}
