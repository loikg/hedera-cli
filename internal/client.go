package internal

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type HederaNetwork string

const (
	HederaNetworkLocal   HederaNetwork = "local"
	HederaNetworkTestNet HederaNetwork = "testnet"
	HederaNetworkMainnet HederaNetwork = "mainnet"
)

func BuildHederaClientFromConfig() (*hedera.Client, error) {
	network, err := parseNetworkString(viper.GetString(ConfigKeyNetwork))
	cobra.CheckErr(err)
	operatorID, operatorKey := resolveOperatorForNetork(network)

	if viper.GetBool(ConfigKeyVerbose) {
		fmt.Printf("Using network: %s, operatorId: %s, operatorKey: %s\n", network, operatorID, operatorKey)
	}

	parsedOperatorID, err := hedera.AccountIDFromString(operatorID)
	if err != nil {
		return nil, fmt.Errorf("failed to parse operator account id %s: %v", operatorID, err)
	}
	parsedOperatorKey, err := hedera.PrivateKeyFromStringEd25519(operatorKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse operator private key: %v", err)
	}

	var client *hedera.Client

	switch {
	case network == HederaNetworkLocal:
		client = buildLocalHederaClient()
	case network == HederaNetworkTestNet:
		client = buildTestnetClient()
	case network == HederaNetworkMainnet:
		client = buildMainnetClient()
	default:
		panic(fmt.Errorf("unknown network: %s", network))
	}

	client.SetOperator(parsedOperatorID, parsedOperatorKey)

	return client, nil
}

func buildTestnetClient() *hedera.Client {
	return hedera.ClientForTestnet()
}

func buildLocalHederaClient() *hedera.Client {
	client := hedera.ClientForNetwork(map[string]hedera.AccountID{"127.0.0.1:50211": {Account: 3}})
	client.SetMirrorNetwork([]string{"127.0.0.1:5600"})
	return client
}

func buildMainnetClient() *hedera.Client {
	return hedera.ClientForMainnet()
}

func parseNetworkString(network string) (HederaNetwork, error) {
	switch {
	case network == "local":
		return HederaNetworkLocal, nil
	case network == "testnet":
		return HederaNetworkTestNet, nil
	case network == "mainnet":
		return HederaNetworkMainnet, nil
	default:
		return "", fmt.Errorf("unsupported network string: %s", network)
	}
}
