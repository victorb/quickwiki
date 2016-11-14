package markdown_test

import (
	"fmt"
	"testing"

	"github.com/victorbjelkholm/quickwiki/config"
	"github.com/victorbjelkholm/quickwiki/plugins/markdown"
)

func TestShouldParseFile(t *testing.T) {
	parser := markdown.Load()
	shouldParseThese := [2]string{"home.md", "helloworld.markdown"}
	for _, filename := range shouldParseThese {
		if !parser.ShouldParseFile(filename) {
			t.Error("ShouldParse file should have returned true for " + filename + " but didnt")
		}
	}
	shouldNotParseThese := [5]string{"home.mdd", "home.mmd", "home.html", "home.css", "home.js"}
	for _, filename := range shouldNotParseThese {
		if parser.ShouldParseFile(filename) {
			t.Error("ShouldParse file should have returned false for " + filename + " but didnt")
		}
	}
}

func TestTransformContent(t *testing.T) {
	config := config.Config{}
	parser := markdown.Load()
	testContent := []byte("## Testing some markdown")
	expectedContent := []byte("<h2>Testing some markdown</h2>\n")
	actualOutput := parser.TransformContent(config, testContent)

	expected := string(expectedContent)
	actual := string(actualOutput)
	if expected != actual {
		fmt.Println("expected")
		fmt.Println(expected)
		fmt.Println("actual")
		fmt.Println(actual)
		t.Error("Did not transform markdown correctly!")
	}
}
