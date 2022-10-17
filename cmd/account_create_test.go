package cmd_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/loikg/hedera-cli/cmd"
	"github.com/loikg/hedera-cli/internal"
	"github.com/loikg/hedera-cli/pkg/hederatest"
)

func TestAccountCreateCommand(t *testing.T) {
	actual := new(bytes.Buffer)
	cmd.RootCmd.SetOut(actual)
	cmd.RootCmd.SetErr(actual)

	cmd.RootCmd.SetArgs([]string{"account", "--network", "local", "create", "--balance", "10.5"})
	cmd.RootCmd.Execute()

	var data internal.M
	err := json.Unmarshal(actual.Bytes(), &data)
	if err != nil {
		t.Fatalf("failed to unmarshal command output: %v", err)
	}

	hederatest.AssertValidAccountID(t, data["accountId"])
	hederatest.AssertValidKeyPair(t, data["privateKey"], data["publicKey"])
}
