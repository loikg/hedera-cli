/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"os"

	"github.com/loikg/hedera-cli/internal/cmd"
)

var version string

func main() {
	cmd.App.Version = version
	if err := cmd.App.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
