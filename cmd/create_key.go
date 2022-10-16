/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/loikg/hedera-cli/internal"
	"github.com/spf13/cobra"
)

// createKeyCmd represents the createKey command
var createKeyCmd = &cobra.Command{
	Use:   "create-key",
	Short: "Create a private key.",
	Long:  `Create a private key. Useful to create supply keys used when creating tokens.`,
	Run:   runCreateKey,
}

func init() {
	RootCmd.AddCommand(createKeyCmd)
}

func runCreateKey(cmd *cobra.Command, args []string) {
	privateKey, err := hedera.GeneratePrivateKey()
	cobra.CheckErr(err)

	cmd.Println(internal.M{
		"privateKey": privateKey.StringRaw(),
		"publicKey":  privateKey.PublicKey().StringRaw(),
	})
}
