package finality

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v2/testing/spectest/shared/deneb/finality"
)

func TestMinimal_Deneb_Finality(t *testing.T) {
	finality.RunFinalityTest(t, "minimal")
}
