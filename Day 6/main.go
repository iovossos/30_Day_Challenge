package main

import (
	"flag"
	"fmt"
	"markdown_converter/converter"
	"os"
)

func main() {
	// Command-line arguments
	inputMarkdown := flag.String("input", "", "Path to the input Markdown file")
	outputHTML := flag.String("output", "output.html", "Path to the output HTML file")
	theme := flag.String("theme", "themes/theme1.html", "Path to the theme HTML file")

	flag.Parse()

	// Check if the input Markdown file is provided
	if *inputMarkdown == "" {
		fmt.Println("Error: Input Markdown file is required.")
		flag.Usage()
		os.Exit(1)
	}

	// Convert Markdown to HTML with the selected theme
	err := converter.ConvertMarkdownToHTML(*inputMarkdown, *outputHTML, *theme)
	if err != nil {
		fmt.Printf("Error converting Markdown to HTML: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully converted %s to %s using theme %s\n", *inputMarkdown, *outputHTML, *theme)
}
