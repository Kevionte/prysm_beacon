package fork

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v2/testing/spectest/shared/deneb/fork"
)

func TestMinimal_UpgradeToDeneb(t *testing.T) {
	fork.RunUpgradeToDeneb(t, "minimal")
}
