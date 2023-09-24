/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"os"

	"github.com/loikg/hedera-cli/cmd"
)

func main() {
	if err := cmd.App.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
