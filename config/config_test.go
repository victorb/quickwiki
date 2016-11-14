package config_test

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/victorbjelkholm/quickwiki/config"
)

func TestConstructSitepaths(t *testing.T) {
	workingDirectory := "/pwd"
	configPath := "/pwd/config.toml"
	pluginsDirectory := "/pwd/plugins"
	inputPagesDirectory := "/pwd/pages"
	inputMediaDirectory := "/pwd/media"
	inputThemesDirectory := "/pwd/themes"
	outputPublicDirectory := "/pwd/public"
	outputPublicMediaDirectory := "/pwd/public/media"

	fmt.Println("Using " + workingDirectory + " as cwd")
	sitePaths := config.CreatePaths(workingDirectory)

	if sitePaths.WorkingDirectory != workingDirectory {
		t.Error("SitePaths.WorkingDirectory was incorrect")
	}
	if sitePaths.Config != configPath {
		t.Error("SitePaths.Config was incorrect")
	}
	if sitePaths.PluginsDirectory != pluginsDirectory {
		t.Error("SitePaths.PluginsDirectory was incorrect")
	}
	if sitePaths.Input.Pages != inputPagesDirectory {
		t.Error("SitePaths.Input.Pages was incorrect")
	}
	if sitePaths.Input.Media != inputMediaDirectory {
		t.Error("SitePaths.Input.Media was incorrect")
	}
	if sitePaths.Input.Themes != inputThemesDirectory {
		t.Error("SitePaths.Input.Themes was incorrect")
	}
	if sitePaths.Output.Public != outputPublicDirectory {
		t.Error("SitePaths.Output.Public was incorrect")
	}
	if sitePaths.Output.Media != outputPublicMediaDirectory {
		t.Error("SitePaths.Output.Media was incorrect")
	}
	spew.Dump(sitePaths)
}
