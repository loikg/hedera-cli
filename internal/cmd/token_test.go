package cmd_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/hashgraph/hedera-sdk-go/v2"

	"github.com/loikg/hedera-cli/internal"
	"github.com/loikg/hedera-cli/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTokenCreate(t *testing.T) {
	t.Parallel()
	testClient := testutils.NewHederaTestClient(t)

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
	actual := testutils.RunCLI(t, args...)

	var data internal.M
	err := json.Unmarshal(actual, &data)
	require.NoError(t, err)

	assert.Equal(t, "MTK", data["symbol"])
	assert.Equal(t, "myToken", data["name"])
	assert.Equal(t, "TOKEN_TYPE_FUNGIBLE_COMMON", data["tokenType"])
	assert.Equal(t, float64(0), data["totalSupply"])
	_, err = testClient.GetToken(data["tokenId"].(string))
	assert.NoError(t, err)
}

func TestTokenShow(t *testing.T) {
	t.Parallel()
	testClient := testutils.NewHederaTestClient(t)
	treasuryID, treasuryKey := testClient.MustCreateAccount(0)
	supplyKey := testClient.MustGenerateKey()
	tokenID := testClient.MustCreateToken(&testutils.CreateTokenOptions{
		Name:              "my token",
		Symbol:            "MT",
		Type:              hedera.TokenTypeFungibleCommon,
		Decimals:          0,
		InitialSupply:     0,
		TreasuryAccountID: *treasuryID,
		TreasuryKey:       treasuryKey,
		SupplyType:        hedera.TokenSupplyTypeInfinite,
		SupplyKey:         supplyKey,
	})
	expectedTpl := string(testutils.Testdata(t, "token_show.json"))
	expected := fmt.Sprintf(expectedTpl, tokenID.String(), treasuryID.String())

	actual := testutils.RunCLI(t, "--network", "local", "token", "show", tokenID.String())

	assert.Equal(t, expected, string(actual))
	_, err := testClient.GetToken(tokenID.String())
	assert.NoError(t, err)
}
