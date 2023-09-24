package cmd

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/loikg/hedera-cli/internal"
	"github.com/urfave/cli/v2"
)

var accountCmd = &cli.Command{
	Name:    "account",
	Usage:   "Manage hedera accounts",
	Aliases: []string{"a"},
	Subcommands: []*cli.Command{
		{
			Flags: []cli.Flag{
				&cli.Float64Flag{
					Name:  "balance",
					Value: 0,
					Usage: "Initial balance to transfer to the newly created account",
				},
			},
			Name:   "create",
			Usage:  "Create hedera accounts",
			Action: createAccountAction,
		},
		{
			Name:  "show",
			Usage: "Show hedera accounts informations",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "account-id",
					Usage: "account id to query",
				},
			},
			Action: showAccountAction,
		},
		{
			Name:  "balance",
			Usage: "Show balances of given account",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "accountId",
					Usage: "account id to query",
				},
			},
			Action: accountBalanceAction,
		},
	},
}

func createAccountAction(ctx *cli.Context) error {
	client, err := internal.BuildHederaClient(internal.BuildHederaClientOptions{
		Network:     internal.HederaNetwork(ctx.String("network")),
		OperatorID:  ctx.String("operator-id"),
		OperatorKey: ctx.String("operator-key"),
	})
	if err != nil {
		return err
	}

	newAccountPrivateKey, err := hedera.PrivateKeyGenerateEd25519()
	if err != nil {
		return err
	}

	newAccountPublicKey := newAccountPrivateKey.PublicKey()

	newAccount, err := hedera.NewAccountCreateTransaction().
		SetKey(newAccountPublicKey).
		SetInitialBalance(hedera.NewHbar(ctx.Float64("balance"))).
		Execute(client)
	if err != nil {
		return err
	}

	receipt, err := newAccount.GetReceipt(client)
	if err != nil {
		return err
	}

	toPrint := internal.M{"accountId": receipt.AccountID.String(), "privateKey": newAccountPrivateKey.StringRaw(), "publicKey": newAccountPublicKey.StringRaw()}

	fmt.Println(toPrint)

	return nil
}

func showAccountAction(ctx *cli.Context) error {
	client, err := internal.BuildHederaClient(internal.BuildHederaClientOptions{
		Network:     internal.HederaNetwork(ctx.String("network")),
		OperatorID:  ctx.String("operator-id"),
		OperatorKey: ctx.String("operator-key"),
	})
	if err != nil {
		return err
	}

	accountID, err := hedera.AccountIDFromString(ctx.String("account-id"))
	if err != nil {
		return err
	}

	accountInfo, err := hedera.NewAccountInfoQuery().SetAccountID(accountID).Execute(client)
	if err != nil {
		return err
	}

	fmt.Println(internal.M{
		"accountId":      accountInfo.AccountID.String(),
		"accountMemo":    accountInfo.AccountMemo,
		"tinyBarBalance": accountInfo.Balance.AsTinybar(),
		"isDeleted":      accountInfo.IsDeleted,
		"ownedNfts":      accountInfo.OwnedNfts,
	})

	return nil
}

func accountBalanceAction(ctx *cli.Context) error {
	client, err := internal.BuildHederaClient(internal.BuildHederaClientOptions{
		Network:     internal.HederaNetwork(ctx.String("network")),
		OperatorID:  ctx.String("operator-id"),
		OperatorKey: ctx.String("operator-key"),
	})
	if err != nil {
		return err
	}

	accountID, err := hedera.AccountIDFromString(ctx.String("accountId"))
	if err != nil {
		return err
	}

	balanceCheckTx, err := hedera.NewAccountBalanceQuery().SetAccountID(accountID).Execute(client)
	if err != nil {
		return err
	}

	fmt.Println(internal.M{
		"tokens":  balanceCheckTx.Tokens,
		"tinybar": balanceCheckTx.Hbars.AsTinybar(),
	})

	return nil
}
