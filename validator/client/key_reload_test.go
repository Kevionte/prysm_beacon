package client

import (
	"context"
	"testing"

	validator2 "github.com/Kevionte/prysm_beacon/v2/consensus-types/validator"
	"github.com/Kevionte/prysm_beacon/v2/validator/client/iface"

	fieldparams "github.com/Kevionte/prysm_beacon/v2/config/fieldparams"
	ethpb "github.com/Kevionte/prysm_beacon/v2/proto/prysm/v1alpha1"
	"github.com/Kevionte/prysm_beacon/v2/testing/assert"
	"github.com/Kevionte/prysm_beacon/v2/testing/require"
	validatormock "github.com/Kevionte/prysm_beacon/v2/testing/validator-mock"
	"github.com/Kevionte/prysm_beacon/v2/validator/client/testutil"
	"github.com/pkg/errors"
	logTest "github.com/sirupsen/logrus/hooks/test"
	"go.uber.org/mock/gomock"
)

func TestValidator_HandleKeyReload(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("active", func(t *testing.T) {
		hook := logTest.NewGlobal()

		inactive := randKeypair(t)
		active := randKeypair(t)

		client := validatormock.NewMockValidatorClient(ctrl)
		beaconClient := validatormock.NewMockBeaconChainClient(ctrl)
		prysmBeaconClient := validatormock.NewMockPrysmBeaconChainClient(ctrl)
		v := validator{
			validatorClient:   client,
			keyManager:        newMockKeymanager(t, inactive),
			genesisTime:       1,
			beaconClient:      beaconClient,
			prysmBeaconClient: prysmBeaconClient,
		}

		resp := testutil.GenerateMultipleValidatorStatusResponse([][]byte{inactive.pub[:], active.pub[:]})
		resp.Statuses[0].Status = ethpb.ValidatorStatus_UNKNOWN_STATUS
		resp.Statuses[1].Status = ethpb.ValidatorStatus_ACTIVE
		client.EXPECT().MultipleValidatorStatus(
			gomock.Any(),
			&ethpb.MultipleValidatorStatusRequest{
				PublicKeys: [][]byte{inactive.pub[:], active.pub[:]},
			},
		).Return(resp, nil)
		prysmBeaconClient.EXPECT().GetValidatorCount(
			gomock.Any(),
			"head",
			[]validator2.Status{validator2.Active},
		).Return([]iface.ValidatorCount{}, nil)

		anyActive, err := v.HandleKeyReload(context.Background(), [][fieldparams.BLSPubkeyLength]byte{inactive.pub, active.pub})
		require.NoError(t, err)
		assert.Equal(t, true, anyActive)
		assert.LogsContain(t, hook, "Waiting for deposit to be observed by beacon node")
		assert.LogsContain(t, hook, "Validator activated")
	})

	t.Run("no active", func(t *testing.T) {
		hook := logTest.NewGlobal()

		client := validatormock.NewMockValidatorClient(ctrl)
		beaconClient := validatormock.NewMockBeaconChainClient(ctrl)
		prysmBeaconClient := validatormock.NewMockPrysmBeaconChainClient(ctrl)
		kp := randKeypair(t)
		v := validator{
			validatorClient:   client,
			keyManager:        newMockKeymanager(t, kp),
			genesisTime:       1,
			beaconClient:      beaconClient,
			prysmBeaconClient: prysmBeaconClient,
		}

		resp := testutil.GenerateMultipleValidatorStatusResponse([][]byte{kp.pub[:]})
		resp.Statuses[0].Status = ethpb.ValidatorStatus_UNKNOWN_STATUS
		client.EXPECT().MultipleValidatorStatus(
			gomock.Any(),
			&ethpb.MultipleValidatorStatusRequest{
				PublicKeys: [][]byte{kp.pub[:]},
			},
		).Return(resp, nil)
		prysmBeaconClient.EXPECT().GetValidatorCount(
			gomock.Any(),
			"head",
			[]validator2.Status{validator2.Active},
		).Return([]iface.ValidatorCount{}, nil)

		anyActive, err := v.HandleKeyReload(context.Background(), [][fieldparams.BLSPubkeyLength]byte{kp.pub})
		require.NoError(t, err)
		assert.Equal(t, false, anyActive)
		assert.LogsContain(t, hook, "Waiting for deposit to be observed by beacon node")
		assert.LogsDoNotContain(t, hook, "Validator activated")
	})

	t.Run("error when getting status", func(t *testing.T) {
		kp := randKeypair(t)
		client := validatormock.NewMockValidatorClient(ctrl)
		v := validator{
			validatorClient: client,
			keyManager:      newMockKeymanager(t, kp),
			genesisTime:     1,
		}

		client.EXPECT().MultipleValidatorStatus(
			gomock.Any(),
			&ethpb.MultipleValidatorStatusRequest{
				PublicKeys: [][]byte{kp.pub[:]},
			},
		).Return(nil, errors.New("error"))

		_, err := v.HandleKeyReload(context.Background(), [][fieldparams.BLSPubkeyLength]byte{kp.pub})
		assert.ErrorContains(t, "error", err)
	})
}
