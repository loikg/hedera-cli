package cmd

import (
	"fmt"

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

func keygenAction(_ *cli.Context) error {
	privateKey, err := hedera.GeneratePrivateKey()
	if err != nil {
		return err
	}

	fmt.Println(internal.M{
		"privateKey": privateKey.StringRaw(),
		"publicKey":  privateKey.PublicKey().StringRaw(),
	})

	return nil
}
