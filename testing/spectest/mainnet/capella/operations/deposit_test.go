package operations

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v2/testing/spectest/shared/capella/operations"
)

func TestMainnet_Capella_Operations_Deposit(t *testing.T) {
	operations.RunDepositTest(t, "mainnet")
}
