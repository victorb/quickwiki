package plugins

import (
	"github.com/victorbjelkholm/quickwiki/plugin"
	"github.com/victorbjelkholm/quickwiki/plugins/autolinker"
	"github.com/victorbjelkholm/quickwiki/plugins/markdown"
)

// LoadPlugins loads all the activated plugins
func LoadPlugins() []plugin.Interface {
	installedPlugins := make([]plugin.Interface, 0)
	installedPlugins = append(installedPlugins, autolinker.Load())
	installedPlugins = append(installedPlugins, markdown.Load())
	return installedPlugins
}
