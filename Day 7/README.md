
# Web Scraper with Concurrency and CSV Export

This project is a web scraper written in Go that allows users to scrape multiple websites concurrently and export the scraped data to a CSV file. The program demonstrates how to use Goroutines for concurrent scraping, HTML parsing for extracting data, and CSV export for saving the results.

## How It Works

The program works as follows:

1. **Scrape Multiple Websites Concurrently**: 
   - The program uses Goroutines to scrape multiple URLs simultaneously. Each URL is processed in its own Goroutine, which fetches the HTML content of the page and extracts specific data (like headings).
   - A `sync.WaitGroup` is used to ensure that the main program waits for all Goroutines to finish before proceeding to export the data.

2. **Extract Data from Web Pages**: 
   - The scraper uses the `goquery` library to parse HTML and extract data from specific elements (like `<h1>`, `<h2>`, etc.).
   - The extracted data from each URL is stored in a map, with the URL as the key and the list of extracted data as the value.

3. **Export Data to CSV**: 
   - Once all the scraping is complete, the program exports the collected data to a CSV file using Go's built-in `encoding/csv` package.
   - Each row in the CSV contains the URL and the corresponding extracted data (e.g., headings).

4. **Command-Line Interface**: 
   - The user provides input URLs (as a comma-separated list) and an output CSV filename via command-line flags.
   - The program processes these inputs and writes the scraped data to the specified output file.

## Project Structure

```
Day 7/
├── go.mod                  # Go module file (already initialized for you)
├── main.go                 # Main file containing CLI and concurrency logic
├── scraper/                # Folder containing the web scraping logic
│   └── scraper.go          # Web scraper that fetches and parses HTML
├── exporter/               # Folder containing the CSV export logic
│   └── exporter.go         # Logic for exporting scraped data to CSV
```

## How to Run the Program

### 1. Clone the Repository

First, clone the repository and navigate to the `Day 7` directory:

```bash
git clone https://github.com/iovossos/30_Day_Challenge.git
cd 30_Day_Challenge/Day\ 7
```

### 2. Run the Program

The project is already initialized with `go.mod`, so there's no need to initialize it again. To run the scraper, provide a list of URLs and specify an output CSV file:

```bash
go run main.go -urls="<comma-separated URLs>" -output=<output CSV filename>
```

For example, to scrape two websites (`https://golang.org/doc/` and `https://example.com`), and save the results in `scraped_data.csv`, use the following command:

```bash
go run main.go -urls="https://golang.org/doc/,https://example.com" -output="scraped_data.csv"
```

### Example Output

After running the command, the program will scrape the given websites concurrently and export the data to the specified CSV file. The program will output something like this:

```
Scraping completed in 2.1s. Data saved to scraped_data.csv
```

The `scraped_data.csv` file will contain the URLs and the corresponding data extracted from the web pages, like this:

```
URL,Data
https://golang.org/doc/,The Go Programming Language
https://golang.org/doc/,Getting Started
https://example.com,Example Domain
```

### File Breakdown

#### `main.go`

- **`main()`**: The entry point of the program. It processes command-line arguments to get the list of URLs and the output CSV file. The program uses Goroutines to scrape multiple URLs concurrently, collects the scraped data, and exports it to a CSV file using the `exporter.ExportToCSV` function.

#### `scraper.go` (in the `scraper` package)

- **`ScrapePage(url string) ([]string, error)`**: 
   - This function takes a URL as input, sends an HTTP GET request to fetch the HTML content, and uses the `goquery` library to extract specific elements (like headings).
   - The function returns a slice of strings representing the extracted data (e.g., headings) and any error encountered during the process.
   
- **Concurrency**: Each URL is scraped in a separate Goroutine for concurrent execution, allowing the program to handle multiple URLs simultaneously.

#### `exporter.go` (in the `exporter` package)

- **`ExportToCSV(filename string, data map[string][]string) error`**: 
   - This function takes the scraped data and writes it to a CSV file. The map contains URLs as keys and the corresponding scraped data as values.
   - Each row in the CSV file contains the URL and an associated piece of data (e.g., a heading).
   
- **CSV Export**: The program uses the built-in `encoding/csv` package to create and write to a CSV file.

## Features

- **Concurrency**: The program leverages Go's Goroutines to scrape multiple websites concurrently, improving efficiency and speed.
- **Web Scraping**: Uses the `goquery` package to parse HTML content and extract data from specific elements (e.g., headings, titles).
- **CSV Export**: Scraped data is exported to a CSV file, making it easy to analyze or process further.
- **CLI Interface**: Simple to use, with command-line flags for input URLs and output file name.


## Optional Enhancements

Here are a few additional features that could be added to improve the project:

1. **Error Handling**: Implement retries for failed HTTP requests and handle non-200 HTTP status codes more gracefully.
2. **Scraping More Data**: Extend the scraper to extract other elements like paragraphs, images, or links.
3. **Pagination**: Add support for scraping paginated content by detecting and following "Next" links.
4. **Rate Limiting**: Add delays between requests to avoid overloading servers or triggering rate limits.
