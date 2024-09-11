package random

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v2/testing/spectest/shared/deneb/sanity"
)

func TestMainnet_Deneb_Random(t *testing.T) {
	sanity.RunBlockProcessingTest(t, "mainnet", "random/random/pyspec_tests")
}
