package fork_transition

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v2/testing/spectest/shared/deneb/fork"
)

func TestMainnet_Deneb_Transition(t *testing.T) {
	fork.RunForkTransitionTest(t, "mainnet")
}
