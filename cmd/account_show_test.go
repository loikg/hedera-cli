package cmd_test

import (
	"testing"

	"github.com/loikg/hedera-cli/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccountShowCommand(t *testing.T) {
	accountID := "0.0.1026"
	expectedOutput := testutils.Testdata(t, "account_show.golden")

	actual, err := testutils.RunCLI(t, "--network", "local", "account", "show", "--account-id", accountID)
	require.NoError(t, err)

	assert.JSONEq(t, string(expectedOutput), string(actual))
}
