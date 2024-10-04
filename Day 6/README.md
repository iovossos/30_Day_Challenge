
# Markdown to HTML Converter with Theming

This project is a simple Markdown to HTML converter written in Go. It allows users to convert Markdown files to HTML and apply different themes to the resulting HTML output. The project demonstrates how Go can handle file input, HTML template processing, and theming.

## How It Works

The program works as follows:

1. **Convert Markdown to HTML**: The program reads the Markdown file, converts it to HTML using the `gomarkdown/markdown` library, and then applies a user-selected theme.
2. **Apply HTML Themes**: The program uses Go’s `html/template` package to apply different themes to the generated HTML content.
3. **CLI Input**: The user provides the input Markdown file, output HTML file, and theme file as command-line arguments.
4. **Generate HTML with Themes**: The resulting HTML file includes the user-selected theme to format the output.

## Project Structure

```
Day 6/
├── go.mod               # Go module file (already initialized for you)
├── main.go              # Main file containing CLI logic and program flow
├── converter/           # Folder containing conversion logic
│   └── converter.go     # File with Markdown to HTML conversion logic
├── themes/              # Folder containing different HTML themes
│   ├── theme1.html      # Light theme template
│   ├── theme2.html      # Dark theme template
│   └── theme3.html      # Minimalist theme template
```

## How to Run the Program

### 1. Clone the Repository

Clone this repository and navigate to the `Day 6` directory:

```bash
git clone https://github.com/iovossos/30_Day_Challenge.git
cd 30_Day_Challenge/Day\ 6
```

### 2. Run the Program

The project is already initialized with `go.mod`, so there's no need to initialize it again. Just run the following command to convert a Markdown file to HTML:

```bash
go run main.go -input=<markdown_file> -output=<html_file> -theme=<theme_file>
```

For example, to convert `myfile.md` using the dark theme:

```bash
go run main.go -input=myfile.md -output=output.html -theme=themes/theme2.html
```

### Example Output

After running the above command, the output will be an HTML file (`output.html`) that is styled with the selected theme.

```bash
go run main.go -input=myfile.md -output=output.html -theme=themes/theme2.html
```

Output HTML:

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Markdown to HTML - Dark Theme</title>
    <style>
        body {
            background-color: #2b2b2b;
            color: #f0f0f0;
            font-family: Arial, sans-serif;
        }
    </style>
</head>
<body>
    <div class="content">
        <h1>My Markdown Heading</h1>
        <p>This is some content from my Markdown file.</p>
    </div>
</body>
</html>
```

## File Breakdown

### `main.go`

- **`main()`**: The entry point of the program. It parses the command-line arguments for the input Markdown file, output HTML file, and theme file. It then calls the `converter.ConvertMarkdownToHTML()` function.

### `converter.go` (in the `converter` package)

- **`ConvertMarkdownToHTML(markdownFile, htmlFile, themeFile string)`**: This function handles reading the Markdown file, converting it to HTML, applying the selected theme, and writing the final output to an HTML file.

### Themes

The `themes` folder contains several HTML templates for styling the generated HTML content. You can create your own themes by modifying the styles in these templates.

## Features

- **Markdown to HTML Conversion**: Converts any valid Markdown file into HTML.
- **Theming**: Supports applying different themes (HTML templates with CSS) to the generated HTML file.
- **Easy to Use**: Takes in the Markdown file, output file, and theme file as command-line arguments for easy conversion.

