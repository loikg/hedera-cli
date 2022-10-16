package cmd_test

import (
	"bytes"
	"testing"

	"github.com/loikg/hedera-cli/cmd"
	"github.com/stretchr/testify/assert"
)

func TestAccountShowCommand(t *testing.T) {
	accountID := "0.0.1026"
	actual := new(bytes.Buffer)
	cmd.RootCmd.SetOut(actual)
	cmd.RootCmd.SetErr(actual)

	cmd.RootCmd.SetArgs([]string{"account", "show", "--account-id", accountID})
	cmd.RootCmd.Execute()

	expectedOutput := `{
	"accountId": "0.0.1026",
	"accountMemo": "",
	"isDeleted": false,
	"ownedNfts": 0,
	"tinyBarBalance": 10000000000000
}`
	assert.JSONEq(t, expectedOutput, actual.String())
}
