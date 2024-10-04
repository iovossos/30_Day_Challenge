package converter

import (
	"fmt"
	"html/template"
	"os"

	"github.com/gomarkdown/markdown"
)

// ConvertMarkdownToHTML converts a Markdown file to HTML and applies a theme
func ConvertMarkdownToHTML(markdownFile, htmlFile, themeFile string) error {
	// Read the Markdown file using os.ReadFile (replacing ioutil.ReadFile)
	mdContent, err := os.ReadFile(markdownFile)
	if err != nil {
		return fmt.Errorf("error reading markdown file: %v", err)
	}

	// Convert Markdown to HTML
	htmlContent := markdown.ToHTML(mdContent, nil, nil)

	// Load the selected theme (HTML template)
	themeTemplate, err := template.ParseFiles(themeFile)
	if err != nil {
		return fmt.Errorf("error loading theme: %v", err)
	}

	// Create the output HTML file
	outputFile, err := os.Create(htmlFile)
	if err != nil {
		return fmt.Errorf("error creating output file: %v", err)
	}
	defer outputFile.Close()

	// Apply the theme to the HTML content and write to the file
	err = themeTemplate.Execute(outputFile, struct {
		Content template.HTML
	}{
		Content: template.HTML(htmlContent),
	})
	if err != nil {
		return fmt.Errorf("error applying theme: %v", err)
	}

	return nil
}
