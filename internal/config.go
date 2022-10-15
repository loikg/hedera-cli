package internal

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	ConfigKeyOperatorAccountID  = "operatorId"
	ConfigKeyOperatorPrivateKey = "operatorKey"
	ConfigKeyNetwork            = "network"
	ConfigKeyNetworks           = "networks"

	FlagDefaultNetwork = "local"
)

func resolveOperatorForNetork(network HederaNetwork) (string, string) {
	accountID := viper.GetString(fmt.Sprintf("%s.%s.%s", ConfigKeyNetworks, network, ConfigKeyOperatorAccountID))
	privateKey := viper.GetString(fmt.Sprintf("%s.%s.%s", ConfigKeyNetworks, network, ConfigKeyOperatorPrivateKey))

	return accountID, privateKey
}
