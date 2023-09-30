package testutils

import (
	"github.com/hashgraph/hedera-sdk-go/v2"
	"os"
	"testing"
)

type HederaTestClient struct {
	client *hedera.Client
	t      *testing.T
}

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

func (c HederaTestClient) GetAccount(accountID string) (hedera.AccountInfo, error) {
	id, err := hedera.AccountIDFromString(accountID)
	if err != nil {
		c.t.Fatalf("failed to parse account id: %v", err)
	}
	return hedera.NewAccountInfoQuery().SetAccountID(id).Execute(c.client)
}

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
