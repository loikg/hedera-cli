package testutils

import (
	"testing"

	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/stretchr/testify/assert"
)

// AssertValidAccountID tries to parse the given accountID as an hedera.AccountID
func AssertValidAccountID(t *testing.T, accountID interface{}) {
	t.Helper()

	assert.IsType(t, "string", accountID)
	_, err := hedera.AccountIDFromString(accountID.(string))
	if err != nil {
		t.Fatalf("expected %s to be a valid account id: %v", accountID, err)
	}
}

// AssertValidKeyPair tries to parse the given key pair as an hedera.PrivateKey
func AssertValidKeyPair(t *testing.T, privateKey interface{}, publicKey interface{}) {
	t.Helper()

	assert.IsType(t, "string", privateKey)
	assert.IsType(t, "string", publicKey)
	_, err := hedera.PrivateKeyFromString(privateKey.(string))
	if err != nil {
		t.Fatalf("failed parse private key %s: %v", privateKey, err)
	}
}
