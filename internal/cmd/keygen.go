package cmd

import (
	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/loikg/hedera-cli/internal"
	"github.com/urfave/cli/v2"
)

var keygenCmd = &cli.Command{
	Name:    "keygen",
	Usage:   "Create a private key.",
	Aliases: []string{"kg"},
	Action:  keygenAction,
}

func keygenAction(ctx *cli.Context) error {
	privateKey, err := hedera.GeneratePrivateKey()
	if err != nil {
		return err
	}

	return internal.ConsolePrint(ctx.App.Writer, internal.M{
		"privateKey": privateKey.String(),
		"publicKey":  privateKey.PublicKey().String(),
	})
}
