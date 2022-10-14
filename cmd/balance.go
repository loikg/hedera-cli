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

// balanceCmd represents the balance command
var balanceCmd = &cobra.Command{
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
	accountCmd.AddCommand(balanceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// balanceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// By default set --authorId value to the operator accountId from the config file
	// If --authorId id provided on the command line then operator accountId will be ignore.
	balanceCmd.Flags().StringVarP(&flagAccountID, "accountId", "i", "", "Account id of which the balance should be displayed")
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

	internal.PrettyPrintJSON(internal.M{
		"tokens":  balanceCheckTx.Tokens,
		"tinybar": balanceCheckTx.Hbars.AsTinybar(),
	})
}
