package cmd_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/loikg/hedera-cli/cmd"
	"github.com/loikg/hedera-cli/internal"
	"github.com/loikg/hedera-cli/internal/hederatest"
)

func TestAccountCreateCommand(t *testing.T) {
	actual := new(bytes.Buffer)
	cmd.RootCmd.SetOut(actual)
	cmd.RootCmd.SetErr(actual)

	cmd.RootCmd.SetArgs([]string{"account", "--network", "local", "create", "--balance", "10.5"})
	err := cmd.RootCmd.Execute()
	require.NoError(t, err)

	var data internal.M
	err = json.Unmarshal(actual.Bytes(), &data)
	require.NoError(t, err)

	hederatest.AssertValidAccountID(t, data["accountId"])
	hederatest.AssertValidKeyPair(t, data["privateKey"], data["publicKey"])
}
