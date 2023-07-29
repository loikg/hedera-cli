package cmd

import (
	"runtime/debug"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display version",
	Run: func(cmd *cobra.Command, args []string) {
		if info, ok := debug.ReadBuildInfo(); ok {
			for _, setting := range info.Settings {
				if setting.Key == "vcs.revision" {
					cmd.Println(setting.Value)
					return
				}
			}
		}
		cmd.Println("no version")
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
