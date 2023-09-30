package cmd_test

import (
	"encoding/json"
	"github.com/hashgraph/hedera-sdk-go/v2"
	"testing"

	"github.com/loikg/hedera-cli/internal"
	"github.com/loikg/hedera-cli/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccountCreateCommand(t *testing.T) {
	t.Parallel()
	t.Run("with balance argument", func(t *testing.T) {
		t.Parallel()
		testClient := testutils.NewHederaTestClient(t)
		actual := testutils.RunCLI(t, "account", "create", "10.5")

		var data internal.M
		err := json.Unmarshal(actual, &data)
		require.NoError(t, err)

		testutils.AssertValidAccountID(t, data["accountId"])
		testutils.AssertValidKeyPair(t, data["privateKey"], data["publicKey"])
		_, err = testClient.GetAccount(data["accountId"].(string))
		assert.NoError(t, err)
	})

	t.Run("without balance argument", func(t *testing.T) {
		t.Parallel()
		testClient := testutils.NewHederaTestClient(t)
		actual := testutils.RunCLI(t, "account", "create")

		var data internal.M
		err := json.Unmarshal(actual, &data)
		require.NoError(t, err)

		testutils.AssertValidAccountID(t, data["accountId"])
		testutils.AssertValidKeyPair(t, data["privateKey"], data["publicKey"])
		_, err = testClient.GetAccount(data["accountId"].(string))
		assert.NoError(t, err)
	})
}

func TestAccountShowCommand(t *testing.T) {
	t.Parallel()
	accountID := "0.0.1026"
	expectedOutput := testutils.Testdata(t, "account_show.golden")

	actual := testutils.RunCLI(t, "--network", "local", "account", "show", accountID)

	assert.JSONEq(t, string(expectedOutput), string(actual))
}

func TestAccountDeleteCommand(t *testing.T) {
	t.Parallel()
	testClient := testutils.NewHederaTestClient(t)
	accountID, privateKey := testClient.MustCreateAccount(0)

	actual := testutils.RunCLI(t, "--network", "local", "account", "delete", accountID.String(), privateKey.String())

	assert.Equal(t, "Status: SUCCESS\n", string(actual))
	var hederaError hedera.ErrHederaPreCheckStatus
	_, err := testClient.GetAccount(accountID.String())
	assert.ErrorAsf(t, err, &hederaError, "expected an ErrHederaPreCheckStatus but got %v instead", err)
	assert.Equal(t, hedera.StatusAccountDeleted, hederaError.Status)
}
