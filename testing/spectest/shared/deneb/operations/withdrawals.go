package operations

import (
	"context"
	"math/big"
	"path"
	"testing"

	"github.com/Kevionte/prysm_beacon/v2/beacon-chain/core/blocks"
	"github.com/Kevionte/prysm_beacon/v2/beacon-chain/state"
	consensusblocks "github.com/Kevionte/prysm_beacon/v2/consensus-types/blocks"
	"github.com/Kevionte/prysm_beacon/v2/consensus-types/interfaces"
	enginev1 "github.com/Kevionte/prysm_beacon/v2/proto/engine/v1"
	ethpb "github.com/Kevionte/prysm_beacon/v2/proto/prysm/v1alpha1"
	"github.com/Kevionte/prysm_beacon/v2/testing/require"
	"github.com/Kevionte/prysm_beacon/v2/testing/spectest/utils"
	"github.com/Kevionte/prysm_beacon/v2/testing/util"
	"github.com/golang/snappy"
)

func RunWithdrawalsTest(t *testing.T, config string) {
	require.NoError(t, utils.SetConfig(t, config))
	testFolders, testsFolderPath := utils.TestFolders(t, config, "deneb", "operations/withdrawals/pyspec_tests")
	for _, folder := range testFolders {
		t.Run(folder.Name(), func(t *testing.T) {
			folderPath := path.Join(testsFolderPath, folder.Name())
			payloadFile, err := util.BazelFileBytes(folderPath, "execution_payload.ssz_snappy")
			require.NoError(t, err)
			payloadSSZ, err := snappy.Decode(nil /* dst */, payloadFile)
			require.NoError(t, err, "Failed to decompress")
			payload := &enginev1.ExecutionPayloadDeneb{}
			require.NoError(t, payload.UnmarshalSSZ(payloadSSZ), "Failed to unmarshal")

			body := &ethpb.BeaconBlockBodyDeneb{ExecutionPayload: payload}
			RunBlockOperationTest(t, folderPath, body, func(_ context.Context, s state.BeaconState, b interfaces.SignedBeaconBlock) (state.BeaconState, error) {
				payload, err := b.Block().Body().Execution()
				if err != nil {
					return nil, err
				}
				withdrawals, err := payload.Withdrawals()
				if err != nil {
					return nil, err
				}
				p, err := consensusblocks.WrappedExecutionPayloadDeneb(&enginev1.ExecutionPayloadDeneb{Withdrawals: withdrawals}, big.NewInt(0))
				require.NoError(t, err)
				return blocks.ProcessWithdrawals(s, p)
			})
		})
	}
}
