package cmd

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/victorbjelkholm/quickwiki/assets"
	"github.com/victorbjelkholm/quickwiki/config"
	"github.com/victorbjelkholm/quickwiki/plugins"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Builds your QuickWiki into a static website",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Try to figure out where to go
		// Defaults to current directory or can be specified as first argument
		// `quickwiki ./pathToWiki`
		cwd := GetCwd()
		wikiPaths := config.CreatePaths(cwd)

		// Make sure that the configuration file actually exists here
		haveConfig, _ := Exists(wikiPaths.Config)
		if !haveConfig {
			log.Fatal("Wiki Config does not exists at " + wikiPaths.Config)
		}

		// Load configuration
		userConfig, err := config.LoadConfig(wikiPaths.Config)
		PanicIfErr(err, "Could not load config.toml from "+wikiPaths.Config)

		globalConfig := config.Config{}
		globalConfig.WikiConfig = userConfig

		// ####
		// #### Procedure for reading all files from directory
		// Find all the pages
		// stay in the same directory for linking
		// Covered by new pages
		files, err := ioutil.ReadDir(wikiPaths.Input.Pages)

		// directories := filter(files, func() { return files.isDir() })
		PanicIfErr(err, "Could not find pages directory")
		// If you only want to load one file, useful for debugging
		// file, _ := os.Stat("./pages/home.md")
		// files := make([]os.FileInfo, 0)
		// files = append(files, file)

		// Covered by new pages
		allPages := globalConfig.PageNames
		for _, f := range files {
			pageName := strings.TrimSuffix(f.Name(), filepath.Ext(f.Name()))
			allPages = append(allPages, pageName)
		}
		sort.Strings(allPages)
		globalConfig.PageNames = allPages
		// Covered by new pages

		// Debug output of the paths and configuration we'll use
		dumpIfDebugEnabled("wikiPaths", wikiPaths)
		dumpIfDebugEnabled("globalConfig", globalConfig)

		// Make sure output directories exists
		CreateDirectoryIfNotExists(wikiPaths.Output.Public)
		CreateDirectoryIfNotExists(wikiPaths.Output.Media)

		// This part copies media files, should really go into it's own
		// media plugin
		// TODO only copy media files that are actually being used
		mediaFiles, err := ioutil.ReadDir(wikiPaths.Input.Media)
		PanicIfErr(err, "Could not find media directory")
		printIfDebugEnabled("<media-copy>", color.FgGreen)
		for _, mf := range mediaFiles {
			src := path.Join(wikiPaths.Input.Media, mf.Name())
			dst := path.Join(wikiPaths.Output.Media, mf.Name())
			err := copy(src, dst)
			PanicIfErr(err, "Could not copy file "+src+" to "+dst)
		}
		printIfDebugEnabled("</media-copy>", color.FgGreen)

		printIfDebugEnabled("<wiki-resources>", color.FgGreen)

		// Load and output the default "redirect to home" HTML page
		// TODO should be able to change where it redirects
		redirectToHomeContents, err := assets.Asset("redirect-to-home")
		PanicIfErr(err, "Could not load default redirect-to-home")
		publicIndexPath := path.Join(wikiPaths.Output.Public, "index.html")
		err = ioutil.WriteFile(publicIndexPath, redirectToHomeContents, 0644)
		PanicIfErr(err, "Could not write default redirect-to-home")
		printIfDebugEnabled("Created "+publicIndexPath, color.FgMagenta)

		// Make sure that theme actually exists in wiki
		themePath := path.Join(wikiPaths.Input.Themes, globalConfig.WikiConfig.Theme.Name)
		themeExists, err := Exists(themePath)
		PanicIfErr(err, "Could not check if "+themePath+" exists")
		if !themeExists {
			color.Set(color.FgRed)
			panic("Could not find '" + globalConfig.WikiConfig.Theme.Name + "' in " + wikiPaths.Input.Themes)
		}

		// Copy the style.css from theme to output directory
		themeStyleCSSPath := path.Join(themePath, "style.css")
		cssOutputPath := path.Join(wikiPaths.Output.Public, "style.css")
		copy(themeStyleCSSPath, cssOutputPath)

		// Read the main template that we'll use to parse all the pages
		themeTemplatePath := path.Join(themePath, "template.html")
		themeTemplate, err := ioutil.ReadFile(themeTemplatePath)
		// themeTemplate := string(themeTemplateBytes)
		globalTemplate := template.Must(template.New("themeTemplate").Parse(string(themeTemplate)))

		printIfDebugEnabled("</wiki-resources>", color.FgGreen)

		// Load all plugins
		// for _, plugin := range globalConfig.WikiConfig.Plugins.Activated {
		// 	fmt.Println(plugin)
		// }
		activatedPlugins := plugins.LoadPlugins()
		dumpIfDebugEnabled("Loaded Plugins", activatedPlugins)

		// Load, parse and output all the pages
		for _, f := range files {
			printIfDebugEnabled("Processing "+f.Name(), color.FgGreen)
			// Helper variables for pages
			// TODO split into it's own pages?
			pageFilename := f.Name()
			pageName := strings.TrimSuffix(pageFilename, filepath.Ext(pageFilename))
			pagePath := path.Join(wikiPaths.Input.Pages, pageFilename)

			if f.IsDir() {
				printIfDebugEnabled("Is directory, nothing to do", color.FgBlue)
				continue
			}

			debugMsg := "Loading " + pagePath
			printIfDebugEnabled(debugMsg, color.FgBlue)

			// Actually read the data
			pageContents, err := ioutil.ReadFile(pagePath)

			PanicIfErr(err, "Could not read page "+f.Name())

			pageContentsString := string(pageContents)

			// Create a directory for every page, later we write to :pageName/index.html
			// to have nice paths in the browser
			CreateDirectoryIfNotExists(path.Join(wikiPaths.Output.Public, pageName))

			// Part of markdown plugin
			// parsers receives input as []byte and needs to have output as []byte
			// parsers also need to define what files to process, markdown one uses .md
			outputAfterPlugins := []byte(pageContents)
			for _, plugin := range activatedPlugins {
				if plugin.ShouldParseFile(pageFilename) {
					outputAfterPlugins = plugin.TransformContent(globalConfig, outputAfterPlugins)
				}
			}
			pageContentsString = string(outputAfterPlugins)

			// Take theme template and replace contents with actual page contents
			// outputToWrite := strings.Replace(string(themeTemplate), "{{contents}}", pageContentsString, 1)

			renderedFilePath := path.Join(wikiPaths.Output.Public, pageName, "index.html")

			// PanicIfErr(err, "Error in parsing template")
			file, err := os.Create(renderedFilePath)
			PanicIfErr(err, "Could not write built file of "+pageName)
			// Have to insert it HERE!
			data := struct {
				Page struct {
					Title    string
					Contents template.HTML
				}
			}{}
			data.Page.Title = pageName
			data.Page.Contents = template.HTML(pageContentsString)
			err = globalTemplate.Execute(file, data)
			// err = ioutil.WriteFile(renderedFilePath, []byte(outputToWrite), 0644)
			PanicIfErr(err, "Could not write built file of "+pageName)
			// fmt.Println(string(output))
		}
		fmt.Println("Done!")
	},
}

func init() {
	RootCmd.AddCommand(buildCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func removeDuplicates(xs *[]string) {
	found := make(map[string]bool)
	j := 0
	for i, x := range *xs {
		if !found[x] {
			found[x] = true
			(*xs)[j] = (*xs)[i]
			j++
		}
	}
	*xs = (*xs)[:j]
}

// PanicIfErr ...
func PanicIfErr(err error, msg string) {
	if err != nil {
		color.Set(color.FgRed)
		defer color.Unset()
		panic(err)
	}
}

// Exists checks if a path exists or not
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

// CreateDirectoryIfNotExists does what the name says
func CreateDirectoryIfNotExists(path string) {
	dirExists, err := Exists(path)
	PanicIfErr(err, "Could not check if directory exists "+path)
	if !dirExists {
		err := os.Mkdir(path, 0700)
		PanicIfErr(err, "Could not create directory "+path)
	}
}
func copy(src, dst string) error {
	debugMsg := "[copy] " + src + " => " + dst
	printIfDebugEnabled(debugMsg, color.FgCyan)
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	cerr := out.Close()
	if err != nil {
		return err
	}
	return cerr
}

// GetCwd figured out the working directory for the wiki
func GetCwd() string {
	if len(os.Args) < 3 {
		pwd, err := os.Getwd()
		PanicIfErr(err, "Could not get current working directory")
		return pwd
	}
	pwd := os.Args[2]
	if path.IsAbs(pwd) {
		return pwd
	}
	wd, err := os.Getwd()
	PanicIfErr(err, "Could not get current working directory")
	finalPath := path.Join(wd, pwd)
	return finalPath
}

func dumpIfDebugEnabled(name string, toDump interface{}) {
	debugEnvVar := os.Getenv("QUICKWIKI_DEBUG")
	if debugEnvVar != "" {
		color.Red("## Dump of: " + name + " ##")
		color.Set(color.FgBlue)
		spew.Dump(toDump)
		color.Unset()
	}
}
func printIfDebugEnabled(str string, clr color.Attribute) {
	debugEnvVar := os.Getenv("QUICKWIKI_DEBUG")
	if debugEnvVar != "" {
		color.Set(clr)
		fmt.Println(str)
		color.Unset()
	}
}
