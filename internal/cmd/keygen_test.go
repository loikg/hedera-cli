package cmd_test

import (
	"encoding/json"
	"testing"

	"github.com/loikg/hedera-cli/internal"
	"github.com/loikg/hedera-cli/internal/testutils"
	"github.com/stretchr/testify/require"
)

func TestKeygen(t *testing.T) {
	t.Parallel()

	actual := testutils.RunCLI(t, "keygen")

	var data internal.M
	err := json.Unmarshal(actual, &data)
	require.NoError(t, err)

	testutils.AssertValidKeyPair(t, data["privateKey"], data["publicKey"])
}
