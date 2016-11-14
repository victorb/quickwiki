package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd is top level
var RootCmd = &cobra.Command{
	Use:   "quickwiki",
	Short: "QuickWiki is a fast, easy and extensible way of running your own, personal wiki",
	Long: `
	QuickWiki is a simple wiki,
driven by a simple binary that you run yourself
and generates a static website of it.`,
}
