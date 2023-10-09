package testutils

import (
	"os"
	"testing"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

// HederaTestClient is a simple wrapper around a *hedera.Client use to make assertions
// about the state of the chain in tests.
type HederaTestClient struct {
	client *hedera.Client
	t      *testing.T
}

// NewHederaTestClient create a new HederaTestClient. Once client should be instanciated for each tests.
func NewHederaTestClient(t *testing.T) *HederaTestClient {
	t.Helper()

	// TODO: find a better way to get the operator id / key or document it
	parsedOperatorID, err := hedera.AccountIDFromString(os.Getenv("OPERATOR_ID"))
	if err != nil {
		t.Fatalf("failed to parse operator account id: %v", err)
	}
	parsedOperatorKey, err := hedera.PrivateKeyFromStringEd25519(os.Getenv("OPERATOR_KEY"))
	if err != nil {
		t.Fatalf("failed to parse operator private key: %v", err)
	}

	client := hedera.ClientForNetwork(map[string]hedera.AccountID{"127.0.0.1:50211": {Account: 3}})
	client.SetMirrorNetwork([]string{"127.0.0.1:5600"})
	client.SetOperator(parsedOperatorID, parsedOperatorKey)

	t.Cleanup(func() {
		if err := client.Close(); err != nil {
			t.Logf("failed to close hedera client: %v", err)
		}
	})

	return &HederaTestClient{
		client: client,
	}
}

// GetAccount query an hedera account by account id.
func (c HederaTestClient) GetAccount(accountID string) (hedera.AccountInfo, error) {
	id, err := hedera.AccountIDFromString(accountID)
	if err != nil {
		c.t.Fatalf("failed to parse account id: %v", err)
	}
	return hedera.NewAccountInfoQuery().SetAccountID(id).Execute(c.client)
}

// MustCreateAccount attempt to create an hedera account, it fails the current tests if any error occur.
func (c HederaTestClient) MustCreateAccount(balance float64) (*hedera.AccountID, hedera.PrivateKey) {
	privateKey, err := hedera.PrivateKeyGenerateEd25519()
	if err != nil {
		c.t.Fatalf("failed to create private key: %v", err)
	}
	publicKey := privateKey.PublicKey()

	createAccountTx, err := hedera.NewAccountCreateTransaction().
		SetKey(publicKey).
		SetInitialBalance(hedera.NewHbar(balance)).
		Execute(c.client)
	if err != nil {
		c.t.Fatalf("failed to execute create account transaction: %v", err)
	}

	receipt, err := createAccountTx.GetReceipt(c.client)
	if err != nil {
		c.t.Fatalf("failed to get receipt of create account transaction: %v", err)
	}

	return receipt.AccountID, privateKey
}

// GetToken query a token by token id.
func (c HederaTestClient) GetToken(tokenIDStr string) (hedera.TokenInfo, error) {
	tokenID, err := hedera.TokenIDFromString(tokenIDStr)
	if err != nil {
		c.t.Fatalf("failed to parse token id: %v", err)
	}

	return hedera.NewTokenInfoQuery().SetTokenID(tokenID).Execute(c.client)
}

// CreateTokenOptions are options pass to HederaTestClient.MustCreateToken to create a new token.
type CreateTokenOptions struct {
	Name              string
	Symbol            string
	Type              hedera.TokenType
	Decimals          uint
	InitialSupply     uint64
	TreasuryAccountID hedera.AccountID
	TreasuryKey       hedera.PrivateKey
	SupplyType        hedera.TokenSupplyType
	SupplyKey         hedera.PrivateKey
}

// MustCreateToken creates a new token with the given params. It fails the current test if any error occur.
func (c HederaTestClient) MustCreateToken(opts *CreateTokenOptions) *hedera.TokenID {
	tokenCreateTx, err := hedera.NewTokenCreateTransaction().
		SetTokenName(opts.Name).
		SetTokenSymbol(opts.Symbol).
		SetTokenType(opts.Type).
		SetDecimals(opts.Decimals).
		SetInitialSupply(opts.InitialSupply).
		SetTreasuryAccountID(opts.TreasuryAccountID).
		SetSupplyType(opts.SupplyType).
		SetSupplyKey(opts.SupplyKey).
		FreezeWith(c.client)
	if err != nil {
		c.t.Fatalf("failed to build the token create transaction: %v", err)
	}

	tokenCreateSign := tokenCreateTx.Sign(opts.TreasuryKey)
	tokenCreateSubmit, err := tokenCreateSign.Execute(c.client)
	if err != nil {
		c.t.Fatalf("failed to execute token create transaction: %v", err)
	}

	tokenCreateRx, err := tokenCreateSubmit.GetReceipt(c.client)
	if err != nil {
		c.t.Fatalf("failed to get receipt for token create transaction: %v", err)
	}

	if tokenCreateRx.Status != hedera.StatusSuccess {
		c.t.Fatalf("create token transaction status is unsuccessful, actual status is: %v", tokenCreateRx.Status)
	}

	return tokenCreateRx.TokenID
}

// MustGenerateKey generate a new key pair. It fails the current test if any error occur.
func (c HederaTestClient) MustGenerateKey() hedera.PrivateKey {
	key, err := hedera.PrivateKeyGenerateEd25519()
	if err != nil {
		c.t.Fatalf("failed to generate ed25519 key: %v", err)
	}

	return key
}
