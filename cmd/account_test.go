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

func TestAccountCreateCommand(t *testing.T) {
	t.Parallel()
	t.Run("with balance argument", func(t *testing.T) {
		t.Parallel()
		actual := testutils.RunCLI(t, "account", "create", "10.5")

		var data internal.M
		err := json.Unmarshal(actual, &data)
		require.NoError(t, err)

		hederatest.AssertValidAccountID(t, data["accountId"])
		hederatest.AssertValidKeyPair(t, data["privateKey"], data["publicKey"])
	})

	t.Run("without balance argument", func(t *testing.T) {
		t.Parallel()
		actual := testutils.RunCLI(t, "account", "create")

		var data internal.M
		err := json.Unmarshal(actual, &data)
		require.NoError(t, err)

		hederatest.AssertValidAccountID(t, data["accountId"])
		hederatest.AssertValidKeyPair(t, data["privateKey"], data["publicKey"])
	})
}

func TestAccountShowCommand(t *testing.T) {
	t.Parallel()
	accountID := "0.0.1026"
	expectedOutput := testutils.Testdata(t, "account_show.golden")

	actual := testutils.RunCLI(t, "--network", "local", "account", "show", accountID)

	assert.JSONEq(t, string(expectedOutput), string(actual))
}
