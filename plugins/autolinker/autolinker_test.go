package autolinker_test

import (
	"fmt"
	"testing"

	"github.com/victorbjelkholm/quickwiki/config"
	"github.com/victorbjelkholm/quickwiki/plugins/autolinker"
)

func TestShouldParseAnyFile(t *testing.T) {
	parser := autolinker.Load()
	files := []string{"hello.md", "test.markdown", "another.html"}
	for _, filename := range files {
		if !parser.ShouldParseFile(filename) {
			t.Error("AutoLinker should parse any files! File " + filename + " was not accepted")
		}
	}
}

type testCase struct {
	Input          string
	ExpectedOutput string
	Name           string
}

func TestConvertPageNameInContentToLink(t *testing.T) {
	globalConfig := config.Config{}
	globalConfig.PageNames = []string{"content"}
	parser := autolinker.Load()

	testCases := []testCase{
		testCase{
			Name:           "Simple Link",
			Input:          "Testing with some content that is here",
			ExpectedOutput: "Testing with some [content](../content) that is here",
		},
		testCase{
			Name:           "No links",
			Input:          "Testing with stuff that should have no links in them",
			ExpectedOutput: "Testing with stuff that should have no links in them",
		},
		testCase{
			Name:           "No links in links!",
			Input:          "Testing with some [content](../content) that is here",
			ExpectedOutput: "Testing with some [content](../content) that is here",
		},
		testCase{
			Name:           "Can link in pages in the end of input",
			Input:          "Ends with content",
			ExpectedOutput: "Ends with [content](../content)",
		},
		testCase{
			Name:           "Does not care about casing",
			Input:          "Ends with Content",
			ExpectedOutput: "Ends with [Content](../content)",
		},
		testCase{
			Name:           "Works if text is bold",
			Input:          "Somehow, **content** in middle",
			ExpectedOutput: "Somehow, **[content](../content)** in middle",
		},
		testCase{
			Name:           "Links with commas around",
			Input:          "Together with content, themes and plugins",
			ExpectedOutput: "Together with [content](../content), themes and plugins",
		},
		testCase{
			Name:           "Many links in one text",
			Input:          "Many content inside this content",
			ExpectedOutput: "Many [content](../content) inside this [content](../content)",
		},
	}

	for _, aTestCase := range testCases {
		actualOutput := string(parser.TransformContent(globalConfig, []byte(aTestCase.Input)))

		if actualOutput != aTestCase.ExpectedOutput {
			fmt.Println("### " + aTestCase.Name + " ###")
			fmt.Println("## output:")
			fmt.Println(actualOutput)
			fmt.Println("## expected output:")
			fmt.Println(aTestCase.ExpectedOutput)
			t.Error(aTestCase.Name + " = Did not transform the link correctly")
			fmt.Println("")
		}
	}
}
