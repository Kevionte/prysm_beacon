package util

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v2/beacon-chain/core/signing"
	"github.com/Kevionte/prysm_beacon/v2/beacon-chain/core/time"
	"github.com/Kevionte/prysm_beacon/v2/config/params"
	"github.com/Kevionte/prysm_beacon/v2/crypto/hash"
	"github.com/Kevionte/prysm_beacon/v2/encoding/ssz"
	"github.com/Kevionte/prysm_beacon/v2/testing/require"
)

func TestGenerateBLSToExecutionChange(t *testing.T) {
	st, keys := DeterministicGenesisStateCapella(t, 64)
	change, err := GenerateBLSToExecutionChange(st, keys[0], 0)
	require.NoError(t, err)

	message := change.Message
	val, err := st.ValidatorAtIndex(message.ValidatorIndex)
	require.NoError(t, err)

	cred := val.WithdrawalCredentials
	require.DeepEqual(t, cred[0], params.BeaconConfig().BLSWithdrawalPrefixByte)

	fromPubkey := message.FromBlsPubkey
	hashFn := ssz.NewHasherFunc(hash.CustomSHA256Hasher())
	digest := hashFn.Hash(fromPubkey)
	require.DeepEqual(t, digest[1:], digest[1:])

	domain, err := signing.Domain(st.Fork(), time.CurrentEpoch(st), params.BeaconConfig().DomainBLSToExecutionChange, st.GenesisValidatorsRoot())
	require.NoError(t, err)

	require.NoError(t, signing.VerifySigningRoot(message, fromPubkey, change.Signature, domain))
}
