/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"github.com/loikg/hedera-cli/internal/cmd"
	"os"
)

func main() {
	if err := cmd.App.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
