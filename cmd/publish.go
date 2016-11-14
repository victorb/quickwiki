package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"

	"github.com/victorbjelkholm/quickwiki/config"
)

var quietFlag bool

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publishes your wiki to somewhere",
	Long: `Publishes a built copy of your quickwiki somewhere,
	current only IPFS is supported`,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := exec.LookPath("ipfs")
		if err != nil {
			log.Fatal("`ipfs` seems to not be installed")
		}

		cwd := args[0]
		wikiPaths := config.CreatePaths(cwd)

		haveConfig, _ := Exists(wikiPaths.Config)
		if !haveConfig {
			log.Fatal("Wiki Config does not exists at " + wikiPaths.Config)
		}

		ipfsAddCmd := exec.Command("ipfs", "add", "-r", "-q", wikiPaths.Output.Public)
		out, err := ipfsAddCmd.Output()
		if err != nil {
			panic(err)
		}
		lines := bytes.Split(out, []byte("\n"))
		directoryHash := string(lines[len(lines)-2])
		if quietFlag {
			fmt.Println(directoryHash)
		} else {
			fmt.Println("Wiki published at: " + directoryHash)
		}
	},
}

func init() {
	publishCmd.Flags().BoolVarP(&quietFlag, "quiet", "q", false, "Silents as much output as possible")
	RootCmd.AddCommand(publishCmd)
}
