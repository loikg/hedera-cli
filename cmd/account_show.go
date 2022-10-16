/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/loikg/hedera-cli/internal"
	"github.com/spf13/cobra"
)

// accountShowCmd represents the accountShow command
var accountShowCmd = &cobra.Command{
	Use:   "show",
	Short: "show the current state of the account",
	Long: `Show the current state of the account.
This query does not include the list of records associated with the account.
Anyone on the network can request an account info for a given account.
Queries do not change the state of the account or require network consensus.
The information is returned from a single node processing the query.`,
	Run: runAccountShow,
}

var (
	fAccountShowAccountID string
)

func init() {
	accountCmd.AddCommand(accountShowCmd)

	accountShowCmd.Flags().StringVarP(&fAccountShowAccountID, "account-id", "i", "", "sets the account ID for which the records should be retrieved")
	accountShowCmd.MarkFlagRequired("account-id")
}

func runAccountShow(cmd *cobra.Command, args []string) {
	accountID, err := hedera.AccountIDFromString(fAccountShowAccountID)
	cobra.CheckErr(err)

	client, err := internal.BuildHederaClientFromConfig()
	cobra.CheckErr(err)

	accountInfo, err := hedera.NewAccountInfoQuery().SetAccountID(accountID).Execute(client)
	cobra.CheckErr(err)

	internal.PrettyPrintJSON(internal.M{
		"accountId":      accountInfo.AccountID.String(),
		"accountMemo":    accountInfo.AccountMemo,
		"tinyBarBalance": accountInfo.Balance.AsTinybar(),
		"isDeleted":      accountInfo.IsDeleted,
		"ownedNfts":      accountInfo.OwnedNfts,
	})
}
