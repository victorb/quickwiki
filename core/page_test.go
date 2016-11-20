package core

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
)

func initializeFS() afero.Fs {
	appFS := afero.NewMemMapFs()
	appFS.MkdirAll("/pages", 0755)
	afero.WriteFile(appFS, "/pages/hello-world.md", []byte("file One"), 0644)
	afero.WriteFile(appFS, "/pages/second-file.md", []byte("file Two"), 0644)
	afero.WriteFile(appFS, "/pages/template.md", []byte("{{ .Name }} is here"), 0644)
	return appFS
}

func TestGetPageNamesFromDirectory(t *testing.T) {
	require := require.New(t)
	fs := initializeFS()

	pages := ReadPagesFromDirectory(fs, "/pages")

	require.Equal(pages[0].Name, "Hello World")
	require.Equal(pages[0].Filename, "hello-world.md")
	require.Equal(pages[0].Slug, "hello-world")
	require.Equal(pages[0].Path, "/pages/hello-world.md")
	require.Equal(pages[0].Contents, []byte("file One"))
}

func TestTurnFilenamesIntoPageNames(t *testing.T) {
	cases := []struct {
		input  string
		output string
	}{
		{"hello-world.md", "Hello World"},
		{"second-file.md", "Second File"},
	}

	for _, tc := range cases {
		actual := FilenameIntoTitle(tc.input)
		require.Equal(t, tc.output, actual)
	}
}

func TestParsePage(t *testing.T) {
	fs := initializeFS()
	pages := ReadPagesFromDirectory(fs, "/pages")
	parsedPage := ParsePage(pages[2], pages)
	require.Equal(t, "Template", pages[2].Name)
	require.Equal(t, []byte("Template is here"), parsedPage)
}
