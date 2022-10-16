/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/loikg/hedera-cli/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// accountBalanceCmd represents the balance command
var accountBalanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: runAccountBalance,
}

var (
	flagAccountID string
)

func init() {
	accountCmd.AddCommand(accountBalanceCmd)

	accountBalanceCmd.Flags().StringVarP(&flagAccountID, "accountId", "i", "", "Account id of which the balance should be displayed")
}

func runAccountBalance(cmd *cobra.Command, args []string) {
	client, err := internal.BuildHederaClientFromConfig()
	cobra.CheckErr(err)

	of := viper.GetString(internal.ConfigKeyOperatorAccountID)
	if flagAccountID != "" {
		of = flagAccountID
	}

	accountID, err := hedera.AccountIDFromString(of)
	cobra.CheckErr(err)

	balanceCheckTx, err := hedera.NewAccountBalanceQuery().SetAccountID(accountID).Execute(client)
	cobra.CheckErr(err)

	cmd.Println(internal.M{
		"tokens":  balanceCheckTx.Tokens,
		"tinybar": balanceCheckTx.Hbars.AsTinybar(),
	})
}
