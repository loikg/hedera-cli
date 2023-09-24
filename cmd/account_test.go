package cmd_test

import (
	"encoding/json"
	"testing"

	"github.com/loikg/hedera-cli/internal"
	"github.com/loikg/hedera-cli/internal/hederatest"
	"github.com/loikg/hedera-cli/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccountBalance(t *testing.T) {
	expected := testutils.Testdata(t, "account_balance.golden")

	actual := testutils.RunCLI(t, "account", "balance", "--accountId", "0.0.1023")

	assert.JSONEq(t, string(expected), string(actual))
}

func TestAccountCreateCommand(t *testing.T) {
	actual := testutils.RunCLI(t, "account", "create", "--balance", "10.5")

	var data internal.M
	err := json.Unmarshal(actual, &data)
	require.NoError(t, err)

	hederatest.AssertValidAccountID(t, data["accountId"])
	hederatest.AssertValidKeyPair(t, data["privateKey"], data["publicKey"])
}

func TestAccountShowCommand(t *testing.T) {
	accountID := "0.0.1026"
	expectedOutput := testutils.Testdata(t, "account_show.golden")

	actual := testutils.RunCLI(t, "--network", "local", "account", "show", "--account-id", accountID)

	assert.JSONEq(t, string(expectedOutput), string(actual))
}
