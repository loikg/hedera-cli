/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/loikg/hedera-cli/internal"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var tokenCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new token on the chain.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: runTokenCreate,
}

var (
	flagTokenCreateTokenType      string
	flagTokenCreateTokenName      string
	flagTokenCreateTokenSymbol    string
	flagTreasuryId                string
	flagTreasuryKey               string
	flagSupplyKey                 string
	flagTokenCreateInitialBalance uint64
	flagTokenCreateDecimals       uint
	flagTokenCreateIsFiniteSupply bool
)

func init() {
	tokenCmd.AddCommand(tokenCreateCmd)

	tokenCreateCmd.Flags().StringVarP(&flagTokenCreateTokenName, "name", "n", "", "Sets the publicly visible name of the token, specified as a string of only ASCII characters.")
	tokenCreateCmd.MarkFlagRequired("name")

	tokenCreateCmd.Flags().StringVarP(&flagTokenCreateTokenSymbol, "symbol", "s", "", "Sets the publicly visible token symbol. It is UTF-8 capitalized alphabetical string identifying the token")
	tokenCreateCmd.MarkFlagRequired("symbol")

	tokenCreateCmd.Flags().StringVarP(&flagTokenCreateTokenType, "type", "t", "", "Specifies the token type.")
	tokenCreateCmd.MarkFlagRequired("type")

	tokenCreateCmd.Flags().StringVarP(&flagTreasuryId, "treasury-id", "i", "", "Sets the account which will act as a treasury for the token. This account will receive the specified initial supply")
	tokenCreateCmd.MarkFlagRequired("treasury-id")
	tokenCreateCmd.Flags().StringVarP(&flagTreasuryKey, "treasury-key", "k", "", "The private key of the treasury account to use to create this token")
	tokenCreateCmd.MarkFlagRequired("treasury-key")

	tokenCreateCmd.Flags().StringVarP(&flagSupplyKey, "supply-key", "u", "", "The suply key to use to create this token")

	tokenCreateCmd.Flags().Uint64VarP(&flagTokenCreateInitialBalance, "balance", "b", 0, "Specifies the initial supply of tokens to be put in circulation. The initial supply is sent to the Treasury Account. The supply is in the lowest denomination possible.")
	tokenCreateCmd.Flags().UintVarP(&flagTokenCreateDecimals, "decimals", "d", 0, "Sets the number of decimal places a token is divisible by. This field can never be changed!")

	tokenCreateCmd.Flags().BoolVarP(&flagTokenCreateIsFiniteSupply, "supply-type", "p", false, "Specifies the token supply type. Default infinit.")
}

func runTokenCreate(cmd *cobra.Command, args []string) {
	client, err := internal.BuildHederaClientFromConfig()
	cobra.CheckErr(err)

	var (
		supplyKey   hedera.Key
		treasuryId  hedera.AccountID
		treasuryKey hedera.PrivateKey
		tokenType   hedera.TokenType
		supplyType  = hedera.TokenSupplyTypeInfinite
	)

	treasuryId, err = hedera.AccountIDFromString(flagTreasuryId)
	cobra.CheckErr(err)
	treasuryKey, err = hedera.PrivateKeyFromStringEd25519(flagTreasuryKey)
	cobra.CheckErr(err)

	if flagTokenCreateIsFiniteSupply {
		supplyType = hedera.TokenSupplyTypeFinite
	}

	switch {
	case flagTokenCreateTokenType == "nft":
		tokenType = hedera.TokenTypeNonFungibleUnique
	case flagTokenCreateTokenType == "ft":
		tokenType = hedera.TokenTypeFungibleCommon
	default:
		fmt.Printf("Invalid value for --type flags. Allowed values are nft and ft")
	}

	if flagSupplyKey != "" {
		supplyKey, err = hedera.PrivateKeyFromStringEd25519(flagSupplyKey)
		cobra.CheckErr(err)
	}

	tokenCreateTx, err := hedera.NewTokenCreateTransaction().
		SetTokenName(flagTokenCreateTokenName).
		SetTokenSymbol(flagTokenCreateTokenSymbol).
		SetTokenType(tokenType).
		SetDecimals(flagTokenCreateDecimals).
		SetInitialSupply(flagTokenCreateInitialBalance).
		SetTreasuryAccountID(treasuryId).
		SetSupplyType(supplyType).
		SetSupplyKey(supplyKey).
		FreezeWith(client)
	cobra.CheckErr(err)

	tokenCreateSign := tokenCreateTx.Sign(treasuryKey)
	tokenCreateSubmit, err := tokenCreateSign.Execute(client)
	cobra.CheckErr(err)

	tokenCreateRx, err := tokenCreateSubmit.GetReceipt(client)
	cobra.CheckErr(err)

	internal.PrettyPrintJSON(internal.M{
		"name":        flagTokenCreateTokenName,
		"symbol":      flagTokenCreateTokenSymbol,
		"tokenType":   tokenType,
		"tokenId":     tokenCreateRx.TokenID.String(),
		"totalSupply": tokenCreateRx.TotalSupply,
	})
}
