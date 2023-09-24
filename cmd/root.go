package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var App = &cli.App{
	Name:  "hedera-cli",
	Usage: "hedera-cli make it easy to interact with the hedera blockchain",
	Description: `hedera-cli make it easy to interact with the hedera blockchain form the command line.
It can connect to a local hedera node, testnet and mainnet.
Operator and network can be configured in the config file located at $HOME/.hedera-cli.yaml by default.`,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "network",
			Value:    "local",
			Usage:    "Hedera network to connect to, possible values are local,testnet,mainnet",
			EnvVars:  []string{"NETWORK"},
			Required: true,
			Action: func(ctx *cli.Context, value string) error {
				if value != "local" && value != "testnet" && value != "mainnet" {
					return fmt.Errorf("invalid network flag value: %s", value)
				}
				return nil
			},
		},
		&cli.StringFlag{
			Name:     "operator-id",
			Usage:    "Hedera operator account id",
			EnvVars:  []string{"OPERATOR_ID"},
			Required: true,
		},
		&cli.StringFlag{
			Name:     "operator-key",
			Usage:    "Hedera operator account private key",
			EnvVars:  []string{"OPERATOR_KEY"},
			Required: true,
		},
	},
	Commands: []*cli.Command{
		accountCmd,
		keygenCmd,
		tokenCmd,
	},
}
