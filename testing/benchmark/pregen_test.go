package benchmark

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v2/testing/require"
)

func TestPreGenFullBlock(t *testing.T) {
	_, err := PreGenFullBlock()
	require.NoError(t, err)
}

func TestPreGenState1Epoch(t *testing.T) {
	_, err := PreGenState1Epoch()
	require.NoError(t, err)
}

func TestPreGenstateFullEpochs(t *testing.T) {
	_, err := PreGenstateFullEpochs()
	require.NoError(t, err)
}
