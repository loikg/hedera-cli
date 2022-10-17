/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// accountCmd represents the account command
var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Create, update, delete accounts",
	Long:  `Manage account on the hedera blockchain`,
}

func init() {
	RootCmd.AddCommand(accountCmd)
}
