// Package markdown is a example parser plugin
// it implements two different functions
//
// TransformContent does the actual transformation
// ShouldParseFile decides which files the transformation should apply to
package markdown

import (
	"path/filepath"

	"github.com/russross/blackfriday"
	// "github.com/shurcooL/github_flavored_markdown"
	"github.com/victorbjelkholm/quickwiki/config"
	"github.com/victorbjelkholm/quickwiki/plugin"
)

// Plugin ...
type Plugin struct {
	plugin.Details
	plugin.Interface
}

// Load Loads this plugin with details filled out
func Load() *Plugin {
	plugin := Plugin{}
	plugin.Details.Name = "Markdown"
	plugin.Details.Description = "Parses .md and .markdown files with black friday markdown parser"
	plugin.Details.Type = "parser"
	plugin.Details.Creator = "Victor Bjelkholm"
	return &plugin
}

// TransformContent gives you a chance to modify the content that gets written
func (p Plugin) TransformContent(config config.Config, content []byte) []byte {
	// return github_flavored_markdown.Markdown(content)
	return blackfriday.MarkdownCommon(content)
}

// TODO More advanced version to be implemented in the future
// htmlFlags := 0
// htmlFlags |= blackfriday.HTML_USE_SMARTYPANTS
// htmlFlags |= blackfriday.HTML_SMARTYPANTS_LATEX_DASHES
// htmlFlags |= blackfriday.HTML_TOC
// extensions := 0
// extensions |= blackfriday.EXTENSION_NO_INTRA_EMPHASIS
// extensions |= blackfriday.EXTENSION_TABLES
// extensions |= blackfriday.EXTENSION_FENCED_CODE
// extensions |= blackfriday.EXTENSION_AUTOLINK
// extensions |= blackfriday.EXTENSION_STRIKETHROUGH
// extensions |= blackfriday.EXTENSION_SPACE_HEADERS
// renderer := blackfriday.HtmlRenderer(htmlFlags, "QuickWiki", "")
// output := blackfriday.Markdown([]byte(dataAsString), renderer, extensions)

// ShouldParseFile decides if this file should be parsed with this parser
func (p Plugin) ShouldParseFile(filename string) bool {
	extension := filepath.Ext(filename)
	if extension == ".md" || extension == ".markdown" {
		return true
	}
	return false
}
