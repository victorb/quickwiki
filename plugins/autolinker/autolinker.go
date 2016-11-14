// Package autolinker is a plugin that would automatically
// create links from words it can find in a page.
// Supposed to run after any other processing and ends with outputting HTML
// BUG
// does not handle text in code blocks very well
package autolinker

import (
	// "github.com/gijsbers/go-pcre"

	"strings"

	"github.com/victorbjelkholm/quickwiki/config"
	"github.com/victorbjelkholm/quickwiki/plugin"
)

// Plugin ...
type Plugin struct {
	plugin.Details
	plugin.Interface
}

// Load ...
func Load() *Plugin {
	plugin := Plugin{}
	plugin.Details.Name = "AutoLinker"
	plugin.Details.Description = "Automatically links to pages in text, if the page exists already"
	plugin.Details.Type = "parser"
	plugin.Details.Creator = "Victor Bjelkholm"
	return &plugin
}

var allowedPadding = []string{
	"*",
	"**",
}

var allowedEndings = []string{
	",",
	".",
}

// TransformContent gives you a chance to modify the content that gets written
func (p Plugin) TransformContent(config config.Config, content []byte) []byte {
	contentStr := string(content)
	for _, pageName := range config.PageNames {
		for _, word := range strings.Fields(contentStr) {
			pageLink := "[" + word + "](../" + pageName + ")"
			if strings.ToLower(pageName) == strings.ToLower(word) {
				contentStr = strings.Replace(contentStr, word, pageLink, -1)
				break
			}
			for _, padding := range allowedPadding {
				if padding+strings.ToLower(pageName)+padding == strings.ToLower(word) {
					paddingWord := strings.TrimPrefix(word, padding)
					paddingWord = strings.TrimSuffix(paddingWord, padding)
					paddingLink := string(padding + "[" + paddingWord + "](../" + pageName + ")" + padding)
					contentStr = strings.Replace(contentStr, word, paddingLink, -1)
					break
				}
			}
			for _, ending := range allowedEndings {
				if strings.ToLower(pageName)+ending == strings.ToLower(word) {
					endingWord := strings.TrimPrefix(word, ending)
					endingWord = strings.TrimSuffix(endingWord, ending)
					endingLink := string("[" + endingWord + "](../" + pageName + ")" + ending)
					contentStr = strings.Replace(contentStr, word, endingLink, -1)
					break
				}
			}
		}
	}
	return []byte(contentStr)
}

// ShouldParseFile decides if this file should be parsed with this parser
func (p Plugin) ShouldParseFile(filename string) bool {
	return true
}
