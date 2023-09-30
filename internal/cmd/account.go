package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/loikg/hedera-cli/internal"
	"github.com/urfave/cli/v2"
	"golang.org/x/sync/errgroup"
)

var errMissingAccountID = errors.New("you must provide a valid account id as an argument")

var accountCmd = &cli.Command{
	Name:    "account",
	Usage:   "Manage hedera accounts",
	Aliases: []string{"a"},
	Subcommands: []*cli.Command{
		{
			Name:      "create",
			Usage:     "Create hedera accounts",
			ArgsUsage: "[<initial_balance>]",
			Action:    createAccountAction,
		},
		{
			Name:      "show",
			Usage:     "Show hedera accounts informations",
			ArgsUsage: "<account_id>",
			Action:    showAccountAction,
		},
		{
			Name:      "delete",
			Usage:     "Delete an hedera account",
			Aliases:   []string{"d"},
			ArgsUsage: "<account_id> <account_private_key>",
			Action:    deleteAccountAction,
		},
	},
}

func createAccountAction(ctx *cli.Context) error {
	balance := float64(0)
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

	if ctx.Args().Present() {
		balanceArg := ctx.Args().First()
		parsedValue, err := strconv.ParseFloat(balanceArg, 64)
		if err != nil {
			return fmt.Errorf("invalid balance argument %s", balanceArg)
		}
		balance = parsedValue
	}

	newAccount, err := hedera.NewAccountCreateTransaction().
		SetKey(newAccountPublicKey).
		SetInitialBalance(hedera.NewHbar(balance)).
		Execute(client)
	if err != nil {
		return err
	}

	receipt, err := newAccount.GetReceipt(client)
	if err != nil {
		return err
	}

	toPrint := internal.M{"accountId": receipt.AccountID.String(), "privateKey": newAccountPrivateKey.String(), "publicKey": newAccountPublicKey.String()}

	_, err = fmt.Fprintf(ctx.App.Writer, "%v", toPrint)
	if err != nil {
		return err
	}

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

	if !ctx.Args().Present() {
		return errMissingAccountID
	}

	accountID, err := hedera.AccountIDFromString(ctx.Args().First())
	if err != nil {
		return err
	}

	data, err := getHederaAccount(client, accountID)
	if err != nil {
		return nil
	}

	fmt.Println(internal.M{
		"accountId":      data.info.AccountID.String(),
		"accountMemo":    data.info.AccountMemo,
		"tinyBarBalance": data.info.Balance.AsTinybar(),
		"isDeleted":      data.info.IsDeleted,
		"ownedNfts":      data.info.OwnedNfts,
		"tokens":         data.balance.Tokens,
	})

	return nil
}

type hederaAccountInfo struct {
	info    hedera.AccountInfo
	balance hedera.AccountBalance
}

func getHederaAccount(c *hedera.Client, accountID hedera.AccountID) (*hederaAccountInfo, error) {
	var (
		accountInfo    hedera.AccountInfo
		accountBalance hedera.AccountBalance
		g              errgroup.Group
	)

	g.Go(func() error {
		data, err := hedera.NewAccountInfoQuery().SetAccountID(accountID).Execute(c)
		if err != nil {
			return err
		}
		accountInfo = data
		return nil
	})

	g.Go(func() error {
		data, err := hedera.NewAccountBalanceQuery().SetAccountID(accountID).Execute(c)
		if err != nil {
			return err
		}
		accountBalance = data
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return &hederaAccountInfo{
		info:    accountInfo,
		balance: accountBalance,
	}, nil
}

func deleteAccountAction(ctx *cli.Context) error {
	client, err := internal.BuildHederaClient(internal.BuildHederaClientOptions{
		Network:     internal.HederaNetwork(ctx.String("network")),
		OperatorID:  ctx.String("operator-id"),
		OperatorKey: ctx.String("operator-key"),
	})
	if err != nil {
		return err
	}

	if ctx.NArg() != 2 {
		cli.ShowSubcommandHelpAndExit(ctx, 1)
	}

	accountID, err := hedera.AccountIDFromString(ctx.Args().First())
	if err != nil {
		return err
	}

	accountKey, err := hedera.PrivateKeyFromString(ctx.Args().Get(1))
	if err != nil {
		return err
	}

	operatorID, err := operatorIDFromCtx(ctx)
	if err != nil {
		return fmt.Errorf("Invalid operator-id: %v", err)
	}
	operatorKey, err := operatorKeyFromCtx(ctx)
	if err != nil {
		return fmt.Errorf("Invalid operator-key: %v", err)
	}

	tx, err := hedera.NewAccountDeleteTransaction().
		SetAccountID(accountID).
		SetTransferAccountID(operatorID).
		FreezeWith(client)
	if err != nil {
		return err
	}

	txResponse, err := tx.Sign(accountKey).Sign(operatorKey).Execute(client)
	if err != nil {
		return err
	}

	receipt, err := txResponse.GetReceipt(client)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(ctx.App.Writer, "Status: %s\n", receipt.Status)
	if err != nil {
		return err
	}

	return nil
}

func operatorIDFromCtx(ctx *cli.Context) (hedera.AccountID, error) {
	return hedera.AccountIDFromString(ctx.String("operator-id"))
}

func operatorKeyFromCtx(ctx *cli.Context) (hedera.PrivateKey, error) {
	return hedera.PrivateKeyFromString(ctx.String("operator-key"))
}
