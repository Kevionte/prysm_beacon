package finality

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v2/testing/spectest/shared/bellatrix/finality"
)

func TestMinimal_Bellatrix_Finality(t *testing.T) {
	finality.RunFinalityTest(t, "minimal")
}
