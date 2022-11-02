/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/loikg/hedera-cli/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	// flagOperatorID  string
	// flagOperatorKey string
	flagNetwork string
	flagVerbose bool
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "hedera-cli",
	Short: "hedera-cli make it easy to interact with the hedera blockchain",
	Long: `hedera-cli make it easy to interact with the hedera blockchain form the command line.
It can connect to a local hedera node, testnet and mainnet.
Operator and network can be configured in the config file located at $HOME/.hedera-cli.yaml by default.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	RootCmd.SetOut(os.Stdout)
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.hedera-cli.yaml)")

	RootCmd.PersistentFlags().BoolVar(&flagVerbose, "verbose", false, "enable debug mesage useful for debugging")
	viper.BindPFlag(internal.ConfigKeyVerbose, RootCmd.PersistentFlags().Lookup("verbose"))

	RootCmd.PersistentFlags().StringVar(&flagNetwork, "network", "", "Network to connect to either local,testnet or mainnet")
	RootCmd.MarkPersistentFlagRequired("network")
	viper.BindPFlag(internal.ConfigKeyNetwork, RootCmd.PersistentFlags().Lookup("network"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		wd, err := os.Getwd()
		cobra.CheckErr(err)

		// Search config in home directory with name ".hedera-cli" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath(wd)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".hedera-cli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		if viper.GetBool(internal.ConfigKeyVerbose) {
			fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		}
	}
}
