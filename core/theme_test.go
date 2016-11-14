package core

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
	"github.com/victorbjelkholm/quickwiki/core"
)

func TestErrorsLoadingTheme(t *testing.T) {
	require := require.New(t)
	appFS := afero.NewMemMapFs()
	appFS.MkdirAll("/themes/simpleblue", 0755)

	_, err := core.LoadTheme(appFS, "/themes/simpleblue")
	require.NotNil(err)
	require.Equal("Missing both template.html and style.css", err.Error())

	afero.WriteFile(appFS, "/themes/simpleblue/template.html", []byte("file One"), 0644)

	_, err = core.LoadTheme(appFS, "/themes/simpleblue")
	require.NotNil(err)
	require.Equal("Missing style.css", err.Error())

	appFS.Remove("/themes/simpleblue/template.html")
	afero.WriteFile(appFS, "/themes/simpleblue/style.css", []byte("file One"), 0644)

	_, err = core.LoadTheme(appFS, "/themes/simpleblue")
	require.NotNil(err)
	require.Equal("Missing template.html", err.Error())
}

func TestLoadingThemeFromPath(t *testing.T) {
	require := require.New(t)
	appFS := afero.NewMemMapFs()
	appFS.MkdirAll("/themes/simpleblue", 0755)
	cssFile := []byte("My CSS File")
	templateFile := []byte("My Template")
	afero.WriteFile(appFS, "/themes/simpleblue/style.css", cssFile, 0644)
	afero.WriteFile(appFS, "/themes/simpleblue/template.html", templateFile, 0644)

	theme, _ := core.LoadTheme(appFS, "/themes/simpleblue")
	require.Equal("Simpleblue", theme.Name)
	require.Equal(cssFile, theme.Style)
	require.Equal(templateFile, theme.Template)
}
