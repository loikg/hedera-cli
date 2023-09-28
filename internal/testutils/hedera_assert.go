package testutils

import (
	"testing"

	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/stretchr/testify/assert"
)

func AssertValidAccountID(t *testing.T, accountID interface{}) {
	t.Helper()
	assert.IsType(t, "string", accountID)
	_, err := hedera.AccountIDFromString(accountID.(string))
	if err != nil {
		t.Errorf("expected %s to be a valid account id: %v", accountID, err)
	}
}

func AssertValidKeyPair(t *testing.T, privateKey interface{}, publicKey interface{}) {
	t.Helper()
	assert.IsType(t, "string", privateKey)
	assert.IsType(t, "string", publicKey)
	key, err := hedera.PrivateKeyFromStringEd25519(privateKey.(string))
	if err != nil {
		t.Fatalf("failed parse private key %s: %v", privateKey, err)
	}
	assert.Equal(t, key.PublicKey().StringRaw(), publicKey)
}
