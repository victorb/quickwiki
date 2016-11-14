package config

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/naoina/toml"
)

// WikiConfig is the wiki configuration
type WikiConfig struct {
	Title       string
	Description string
	Theme       struct {
		Name string
	}
	Options struct {
		OutputDirectory string
		PagesDirectory  string
	}
	Plugins struct {
		Activated []string
	}
}

// Config is the global, internal configuration that gets created in the
// beginning (and passed to plugins etc) and used to contain some state
type Config struct {
	WikiConfig
	PageNames []string
}

// SitePaths contains all the paths for a wiki
type SitePaths struct {
	WorkingDirectory string
	Config           string
	PluginsDirectory string
	Input            struct {
		Pages  string
		Media  string
		Themes string
	}
	Output struct {
		Public string
		Media  string
	}
}

// LoadConfig loads a .toml and unmarshals it into a WikiConfig
func LoadConfig(path string) (WikiConfig, error) {
	f, err := os.Open(path)
	if err != nil {
		return WikiConfig{}, err
	}
	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return WikiConfig{}, err
	}
	var config WikiConfig
	if err := toml.Unmarshal(buf, &config); err != nil {
		return WikiConfig{}, err
	}
	return config, nil
}

// CreatePaths creates a SitePaths based on working directory
func CreatePaths(root string) SitePaths {
	paths := SitePaths{}
	paths.WorkingDirectory = root
	paths.Config = path.Join(root, "config.toml")
	paths.PluginsDirectory = path.Join(root, "plugins")
	paths.Input.Pages = path.Join(root, "pages")
	paths.Input.Media = path.Join(root, "media")
	paths.Input.Themes = path.Join(root, "themes")
	paths.Output.Public = path.Join(root, "public")
	paths.Output.Media = path.Join(root, "public", "media")
	return paths
}
