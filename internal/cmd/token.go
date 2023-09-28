package cmd

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/loikg/hedera-cli/internal"
	"github.com/urfave/cli/v2"
)

var tokenCmd = &cli.Command{
	Name:    "token",
	Usage:   "Create, update, delete fungible and non fungible tokens",
	Aliases: []string{"t"},
	Subcommands: []*cli.Command{
		{
			Name:  "create",
			Usage: "Create a new token on the chain.",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "name",
					Aliases:  []string{"n"},
					Usage:    "Sets the publicly visible name of the token, specified as a string of only ASCII characters.",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "symbol",
					Aliases:  []string{"s"},
					Usage:    "Sets the publicly visible token symbol. It is UTF-8 capitalized alphabetical string identifying the token",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "type",
					Aliases:  []string{"t"},
					Usage:    "Specifies the token type. The value can be ft or nft",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "treasury-id",
					Aliases:  []string{"i"},
					Usage:    "Sets the account which will act as a treasury for the token. This account will receive the specified initial supply",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "treasury-key",
					Aliases:  []string{"k"},
					Usage:    "The private key of the treasury account to use to create this token",
					Required: true,
				},
				&cli.StringFlag{
					Name:    "supply-key",
					Aliases: []string{"u"},
					Usage:   "The suply key to use to create this token",
				},
				&cli.Uint64Flag{
					Name:     "balance",
					Aliases:  []string{"b"},
					Usage:    "Specifies the initial supply of tokens to be put in circulation. The initial supply is sent to the Treasury Account. The supply is in the lowest denomination possible.",
					Required: true,
				},
				&cli.UintFlag{
					Name:    "decimals",
					Aliases: []string{"d"},
					Usage:   "Sets the number of decimal places a token is divisible by. This field can never be changed!",
					Value:   0,
				},
				&cli.StringFlag{
					Name:     "supply-type",
					Aliases:  []string{"p"},
					Usage:    "Set the supply to be infinite for the token",
					Required: true,
				},
			},
			Action: tokenCreateAction,
		},
	},
}

func tokenCreateAction(ctx *cli.Context) error {
	client, err := internal.BuildHederaClient(internal.BuildHederaClientOptions{
		Network:     internal.HederaNetwork(ctx.String("network")),
		OperatorID:  ctx.String("operator-id"),
		OperatorKey: ctx.String("operator-key"),
	})
	if err != nil {
		return err
	}

	var (
		supplyKey   hedera.Key
		treasuryID  hedera.AccountID
		treasuryKey hedera.PrivateKey
		tokenType   hedera.TokenType
		supplyType  hedera.TokenSupplyType
	)

	treasuryID, err = hedera.AccountIDFromString(ctx.String("treasury-id"))
	if err != nil {
		return fmt.Errorf("invalid treasury id")
	}
	treasuryKey, err = hedera.PrivateKeyFromStringEd25519(ctx.String("treasury-key"))
	if err != nil {
		return fmt.Errorf("invalid treasury private key")
	}

	tokenSupplyTypeValue := ctx.String("supply-type")
	switch {
	case tokenSupplyTypeValue == "finite":
		supplyType = hedera.TokenSupplyTypeFinite
	case tokenSupplyTypeValue == "infinite":
		supplyType = hedera.TokenSupplyTypeInfinite
	default:
		return fmt.Errorf("invalid value %s for flag --supply-type", tokenSupplyTypeValue)
	}

	tokenTypeValue := ctx.String("type")
	switch {
	case tokenTypeValue == "nft":
		tokenType = hedera.TokenTypeNonFungibleUnique
	case tokenTypeValue == "ft":
		tokenType = hedera.TokenTypeFungibleCommon
	default:
		return fmt.Errorf("invalid value for --type flags. Allowed values are nft and ft")
	}

	if supplyKeyValue := ctx.String("supply-key"); supplyKeyValue != "" {
		supplyKey, err = hedera.PrivateKeyFromStringEd25519(ctx.String("supply-key"))
		if err != nil {
			return err
		}
	}

	tokenName := ctx.String("name")
	tokenSymbol := ctx.String("symbol")

	tokenCreateTx, err := hedera.NewTokenCreateTransaction().
		SetTokenName(tokenName).
		SetTokenSymbol(tokenSymbol).
		SetTokenType(tokenType).
		SetDecimals(ctx.Uint("decimals")).
		SetInitialSupply(ctx.Uint64("balance")).
		SetTreasuryAccountID(treasuryID).
		SetSupplyType(supplyType).
		SetSupplyKey(supplyKey).
		FreezeWith(client)
	if err != nil {
		return err
	}

	tokenCreateSign := tokenCreateTx.Sign(treasuryKey)
	tokenCreateSubmit, err := tokenCreateSign.Execute(client)
	if err != nil {
		return err
	}

	tokenCreateRx, err := tokenCreateSubmit.GetReceipt(client)
	if err != nil {
		return err
	}

	fmt.Println(internal.M{
		"name":        tokenName,
		"symbol":      tokenSymbol,
		"tokenType":   tokenType.String(),
		"tokenId":     tokenCreateRx.TokenID.String(),
		"totalSupply": tokenCreateRx.TotalSupply,
	})

	return nil
}
