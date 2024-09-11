package sanity

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v2/testing/spectest/shared/capella/sanity"
)

func TestMinimal_Capella_Sanity_Blocks(t *testing.T) {
	sanity.RunBlockProcessingTest(t, "minimal", "sanity/blocks/pyspec_tests")
}
