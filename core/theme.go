package core

import (
	"errors"
	"path"
	"strings"

	"github.com/spf13/afero"
)

// Theme is
type Theme struct {
	Name     string
	Style    []byte
	Template []byte
}

// A theme needs to have at least two files
// template.html and style.css

// LoadTheme loads a theme from a directory
func LoadTheme(fs afero.Fs, directory string) (Theme, error) {
	theme := Theme{}
	files, err := afero.ReadDir(fs, directory)
	if err != nil {
		return theme, err
	}
	foundTemplate := false
	foundStylesheet := false
	for _, file := range files {
		if file.Name() == "template.html" {
			foundTemplate = true
			template, err := afero.ReadFile(fs, path.Join(directory, file.Name()))
			if err != nil {
				panic(err)
			}
			theme.Template = template
			continue
		}
		if file.Name() == "style.css" {
			foundStylesheet = true
			stylesheet, err := afero.ReadFile(fs, path.Join(directory, file.Name()))
			if err != nil {
				panic(err)
			}
			theme.Style = stylesheet
			continue
		}
	}
	if !foundTemplate && !foundStylesheet {
		return theme, errors.New("Missing both template.html and style.css")
	}
	if !foundTemplate {
		return theme, errors.New("Missing template.html")
	}
	if !foundStylesheet {
		return theme, errors.New("Missing style.css")
	}
	theme.Name = strings.Title(path.Base(directory))
	return theme, nil
}
