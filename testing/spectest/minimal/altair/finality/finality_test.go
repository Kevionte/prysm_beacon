package finality

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v2/testing/spectest/shared/altair/finality"
)

func TestMinimal_Altair_Finality(t *testing.T) {
	finality.RunFinalityTest(t, "minimal")
}
