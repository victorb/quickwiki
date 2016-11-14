package plugin

import "github.com/victorbjelkholm/quickwiki/config"

// Details contains the details about the plugin
type Details struct {
	Name        string
	Description string
	Type        string
	Creator     string
}

// Interface is the interface of a plugin
type Interface interface {
	Load() Details
	ShouldParseFile(filename string) bool
	TransformContent(config.Config, []byte) []byte
}

// Parser is a plugin type that allows you to change how files looks before/after
// they have been written to disk. Default one is the Markdown Parser
type Parser struct {
	Name string
}
