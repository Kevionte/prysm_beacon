package operations

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v2/testing/spectest/shared/altair/operations"
)

func TestMinimal_Altair_Operations_AttesterSlashing(t *testing.T) {
	operations.RunAttesterSlashingTest(t, "minimal")
}
