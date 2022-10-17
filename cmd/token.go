/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// tokenCmd represents the token command
var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Create, update, delete fungible and non fungible tokens",
	Long:  `Manage token fungible (ft) and non-fungible (nft) tokens.`,
}

func init() {
	RootCmd.AddCommand(tokenCmd)
}
