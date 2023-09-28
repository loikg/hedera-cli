package internal

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

type HederaNetwork string

const (
	HederaNetworkLocal   HederaNetwork = "local"
	HederaNetworkTestNet HederaNetwork = "testnet"
	HederaNetworkMainnet HederaNetwork = "mainnet"
)

type BuildHederaClientOptions struct {
	Network     HederaNetwork
	OperatorID  string
	OperatorKey string
}

func BuildHederaClient(opts BuildHederaClientOptions) (*hedera.Client, error) {
	parsedOperatorID, err := hedera.AccountIDFromString(opts.OperatorID)
	if err != nil {
		return nil, fmt.Errorf("failed to parse operator account id %s: %v", opts.OperatorID, err)
	}
	parsedOperatorKey, err := hedera.PrivateKeyFromStringEd25519(opts.OperatorKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse operator private key: %v", err)
	}

	var client *hedera.Client

	switch {
	case opts.Network == HederaNetworkLocal:
		client = buildLocalHederaClient()
	case opts.Network == HederaNetworkTestNet:
		client = buildTestnetClient()
	case opts.Network == HederaNetworkMainnet:
		client = buildMainnetClient()
	default:
		panic(fmt.Errorf("unknown network: %s", opts.Network))
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

// func parseNetworkString(network string) (HederaNetwork, error) {
// 	switch {
// 	case network == "local":
// 		return HederaNetworkLocal, nil
// 	case network == "testnet":
// 		return HederaNetworkTestNet, nil
// 	case network == "mainnet":
// 		return HederaNetworkMainnet, nil
// 	default:
// 		return "", fmt.Errorf("unsupported network string: %s", network)
// 	}
// }
