package cmd_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/loikg/hedera-cli/internal"
	"github.com/loikg/hedera-cli/internal/hederatest"
	"github.com/loikg/hedera-cli/internal/testutils"
)

func TestAccountCreateCommand(t *testing.T) {
	actual, err := testutils.RunCLI(t, "account", "--network", "local", "create", "--balance", "10.5")
	require.NoError(t, err)

	var data internal.M
	err = json.Unmarshal(actual, &data)
	require.NoError(t, err)

	hederatest.AssertValidAccountID(t, data["accountId"])
	hederatest.AssertValidKeyPair(t, data["privateKey"], data["publicKey"])
}
