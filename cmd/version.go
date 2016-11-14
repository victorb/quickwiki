package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version string
)

// versionCmd represents the new command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Shows current version of QuickWiki",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("QuickWiki version: " + version)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
