package internal

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/spf13/viper"
)

const (
	HederaNetworkLocal = iota
	HederaNetworkTestNet
	HederaNetworkMainnet
)

func BuildHederaClientFromConfig() (*hedera.Client, error) {
	operatorID := viper.GetString(ConfigKeyOperatorAccountID)
	operatorKey := viper.GetString(ConfigKeyOperatorPrivateKey)

	return buildHederaClient(HederaNetworkLocal, operatorID, operatorKey)
}

func buildHederaClient(network int, operatorID, operatorKey string) (*hedera.Client, error) {
	switch network {
	case HederaNetworkLocal:
		return buildLocalHederaClient(operatorID, operatorKey)
	default:
		return nil, fmt.Errorf("unknow hedera network: %d", network)
	}
}

func buildLocalHederaClient(operatorID, operatorKey string) (*hedera.Client, error) {
	parsedOperatorID, err := hedera.AccountIDFromString(operatorID)
	if err != nil {
		return nil, fmt.Errorf("failed to parse operator account id %s: %v", operatorID, err)
	}
	parsedOperatorKey, err := hedera.PrivateKeyFromStringEd25519(operatorKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse operator private key: %v", err)
	}

	client := hedera.ClientForNetwork(map[string]hedera.AccountID{"127.0.0.1:50211": {Account: 3}})
	client.SetMirrorNetwork([]string{"127.0.0.1:5600"})
	client.SetOperator(parsedOperatorID, parsedOperatorKey)

	return client, nil
}
