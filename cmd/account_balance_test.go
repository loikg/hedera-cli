package cmd_test

import (
	"testing"

	"github.com/loikg/hedera-cli/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccountBalance(t *testing.T) {
	expected := testutils.Testdata(t, "account_balance.golden")

	actual, err := testutils.RunCLI(t, "account", "balance", "--accountId", "0.0.1023")
	require.NoError(t, err)

	assert.JSONEq(t, string(expected), string(actual))
}
