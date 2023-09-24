package cmd_test

import (
	"encoding/json"
	"errors"
	"os/exec"
	"testing"

	"github.com/loikg/hedera-cli/internal"
	"github.com/loikg/hedera-cli/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTokenCreate(t *testing.T) {
	args := []string{
		"token",
		"create",
		"--balance",
		"100",
		"--decimals",
		"2",
		"--name",
		"myToken",
		"--symbol",
		"MTK",
		"--type",
		"ft",
		"--supply-type",
		"infinite",
		"--treasury-id",
		"0.0.1022",
		"--treasury-key",
		"851a12c2561f12014d51e30bfa6342d34275c77866118f18a29659cdc12a485b",
	}
	actual, err := testutils.RunCLI(t, args...)
	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) {
		t.Log(string(exitErr.Stderr))
	}
	require.NoError(t, err)

	var data internal.M
	err = json.Unmarshal(actual, &data)
	require.NoError(t, err)

	assert.Equal(t, "MTK", data["symbol"])
	assert.Equal(t, "myToken", data["name"])
	assert.Equal(t, "TOKEN_TYPE_FUNGIBLE_COMMON", data["tokenType"])
	assert.Equal(t, float64(0), data["totalSupply"])
}
