/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/loikg/hedera-cli/internal"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create hedera accounts",
	Long:  `Create hedera accounts with the configured operator account.`,
	Run:   accountCreate,
}

var (
	balanceFlag float64
)

func init() {
	accountCmd.AddCommand(createCmd)

	createCmd.Flags().Float64VarP(&balanceFlag, "balance", "b", 0, "Initial balance to transfer to the newly created account")
}

func accountCreate(cmd *cobra.Command, args []string) {
	client, err := internal.BuildHederaClientFromConfig()
	cobra.CheckErr(err)

	newAccountPrivateKey, err := hedera.PrivateKeyGenerateEd25519()
	cobra.CheckErr(err)

	newAccountPublicKey := newAccountPrivateKey.PublicKey()

	newAccount, err := hedera.NewAccountCreateTransaction().
		SetKey(newAccountPublicKey).
		SetInitialBalance(hedera.NewHbar(balanceFlag)).
		Execute(client)
	cobra.CheckErr(err)

	receipt, err := newAccount.GetReceipt(client)
	cobra.CheckErr(err)

	toPrint := internal.M{"accountId": receipt.AccountID.String(), "privateKey": newAccountPrivateKey.StringRaw(), "publicKey": newAccountPublicKey.StringRaw()}
	internal.PrettyPrintJSON(toPrint)
}
