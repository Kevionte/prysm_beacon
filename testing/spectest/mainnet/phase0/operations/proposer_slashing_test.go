package operations

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v2/testing/spectest/shared/phase0/operations"
)

func TestMainnet_Phase0_Operations_ProposerSlashing(t *testing.T) {
	operations.RunProposerSlashingTest(t, "mainnet")
}
