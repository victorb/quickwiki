// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/victorbjelkholm/quickwiki/assets"
	"github.com/victorbjelkholm/quickwiki/config"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Creates a new QuickWiki site",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("You need to provide the name of your new quickwiki atleast")
			os.Exit(4)
		}
		directoryForWiki := path.Join(cmd.Flag("path").Value.String(), args[0])
		somethingAlreadyExists, err := Exists(directoryForWiki)
		PanicIfErr(err, "Couldnt check if it exists")
		if somethingAlreadyExists {
			fmt.Println("Seems like something is already at " + directoryForWiki)
			fmt.Println("Cannot create a new wiki there...")
			os.Exit(5)
		}
		paths := config.CreatePaths(directoryForWiki)
		CreateDirectoryIfNotExists(paths.WorkingDirectory)
		CreateDirectoryIfNotExists(paths.Input.Pages)
		CreateDirectoryIfNotExists(paths.Input.Media)
		CreateDirectoryIfNotExists(paths.Input.Themes)
		CreateDirectoryIfNotExists(path.Join(paths.Input.Themes, "simpleblue"))
		// What we need to need to have as assets and ready to copy
		loadAndWriteAsset("default-config", paths.Config)
		loadAndWriteAsset("default-page", path.Join(paths.Input.Pages, "home.md"))
		loadAndWriteAsset("simpleblue/readme.md", path.Join(paths.Input.Themes, "simpleblue", "readme.md"))
		loadAndWriteAsset("simpleblue/style.css", path.Join(paths.Input.Themes, "simpleblue", "style.css"))
		loadAndWriteAsset("simpleblue/template.html", path.Join(paths.Input.Themes, "simpleblue", "template.html"))
		fmt.Println("Created new QuickWiki over at " + directoryForWiki)
	},
}

func loadAndWriteAsset(name, path string) {
	file, err := assets.Asset(name)
	PanicIfErr(err, "Could not load "+name)
	err = ioutil.WriteFile(path, file, 0644)
	PanicIfErr(err, "Could not write "+name+" to "+path)
}

func init() {
	RootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	newCmd.Flags().StringP("path", "p", wd, "Path to where to create site")

}
