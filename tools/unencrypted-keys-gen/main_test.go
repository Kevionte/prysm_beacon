package main

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/Kevionte/prysm_beacon/v2/testing/assert"
	"github.com/Kevionte/prysm_beacon/v2/testing/require"
	"github.com/Kevionte/prysm_beacon/v2/tools/unencrypted-keys-gen/keygen"
)

func TestSavesUnencryptedKeys(t *testing.T) {
	keys := 2
	numKeys = &keys
	ctnr := generateUnencryptedKeys(0 /* start index */)
	buf := new(bytes.Buffer)
	require.NoError(t, keygen.SaveUnencryptedKeysToFile(buf, ctnr))
	enc := buf.Bytes()
	dec := &keygen.UnencryptedKeysContainer{}
	require.NoError(t, json.Unmarshal(enc, dec))
	assert.DeepEqual(t, ctnr, dec)
}
