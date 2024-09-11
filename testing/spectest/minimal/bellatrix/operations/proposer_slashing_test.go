package operations

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v2/testing/spectest/shared/bellatrix/operations"
)

func TestMinimal_Bellatrix_Operations_ProposerSlashing(t *testing.T) {
	operations.RunProposerSlashingTest(t, "minimal")
}
