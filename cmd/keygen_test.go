package cmd_test

import (
	"encoding/json"
	"testing"

	"github.com/loikg/hedera-cli/internal"
	"github.com/loikg/hedera-cli/internal/hederatest"
	"github.com/loikg/hedera-cli/internal/testutils"
	"github.com/stretchr/testify/require"
)

func TestKeygen(t *testing.T) {
	actual, err := testutils.RunCLI(t, "keygen")
	require.NoError(t, err)

	var data internal.M
	err = json.Unmarshal(actual, &data)
	require.NoError(t, err)

	hederatest.AssertValidKeyPair(t, data["privateKey"], data["publicKey"])
}
