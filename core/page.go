package core

import (
	"bytes"
	"html/template"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
)

// Page is a page in QuickWiki
type Page struct {
	Name     string
	Filename string
	Slug     string
	Path     string
	Contents []byte
}

// ReadPagesFromDirectory does
func ReadPagesFromDirectory(fs afero.Fs, directory string) []Page {
	files, err := afero.ReadDir(fs, directory)
	pages := []Page{}
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		pageName := strings.TrimSuffix(f.Name(), filepath.Ext(f.Name()))
		fullPath := path.Join(directory, f.Name())
		pageContents, err := afero.ReadFile(fs, fullPath)
		if err != nil {
			panic(err)
		}
		pages = append(pages, Page{
			Name:     FilenameIntoTitle(f.Name()),
			Filename: f.Name(),
			Slug:     pageName,
			Path:     fullPath,
			Contents: pageContents,
		})
	}
	return pages
}

// FilenameIntoTitle turns a filename into a title
func FilenameIntoTitle(filename string) string {
	filename = strings.TrimSuffix(filename, filepath.Ext(filename))
	filename = strings.Title(filename)
	filename = strings.Replace(filename, "-", " ", -1)
	return filename
}

// ParsePage parses a page
func ParsePage(page Page, pages []Page) []byte {
	t := template.Must(template.New("page").Parse(string(page.Contents)))
	var pageContents bytes.Buffer
	err := t.Execute(&pageContents, page)
	if err != nil {
		panic(err)
	}
	return pageContents.Bytes()
}
