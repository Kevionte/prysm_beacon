//go:build minimal

package field_params_test

import (
	"testing"

	fieldparams "github.com/Kevionte/prysm_beacon/v2/config/fieldparams"
	"github.com/Kevionte/prysm_beacon/v2/config/params"
	"github.com/Kevionte/prysm_beacon/v2/testing/require"
)

func TestFieldParametersValues(t *testing.T) {
	params.SetupTestConfigCleanup(t)
	min := params.MinimalSpecConfig().Copy()
	params.OverrideBeaconConfig(min)
	require.Equal(t, "minimal", fieldparams.Preset)
	testFieldParametersMatchConfig(t)
}
