package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"
	"web_scraper/exporter"
	"web_scraper/scraper"
)

// Result struct holds the URL and the data scraped from that URL.
type Result struct {
	URL  string
	Data []string
}

func main() {
	// Define CLI flags: input URLs (comma-separated) and output CSV file.
	urlsArg := flag.String("urls", "", "Comma-separated list of URLs to scrape")
	outputFile := flag.String("output", "output.csv", "Output CSV file")
	flag.Parse()

	// Split the comma-separated URLs into a slice.
	urls := strings.Split(*urlsArg, ",")

	// Channel to collect results from concurrent scraping Goroutines.
	results := make(chan Result)

	// WaitGroup to wait for all Goroutines to finish before proceeding.
	var wg sync.WaitGroup

	// Start timer to measure the total scraping time.
	startTime := time.Now()

	// Loop over the list of URLs, spawning a new Goroutine for each URL to scrape.
	for _, url := range urls {
		wg.Add(1) // Increment the WaitGroup counter for each Goroutine.
		go func(url string) {
			defer wg.Done() // Decrement the WaitGroup counter when this Goroutine is done.

			// Scrape the page and capture the data.
			data, err := scraper.ScrapePage(url)
			if err != nil {
				// If there's an error scraping the page, log it and return.
				log.Printf("Error scraping %s: %v", url, err)
				return
			}

			// Send the scraped data to the results channel.
			results <- Result{URL: url, Data: data}
		}(url) // Pass the URL to the Goroutine.
	}

	// Close the results channel once all Goroutines have finished.
	go func() {
		wg.Wait()
		close(results)
	}()

	// Map to store the final scraped data (URL as key, data as value).
	scrapedData := make(map[string][]string)

	// Receive the results from the results channel and store them in the scrapedData map.
	for result := range results {
		scrapedData[result.URL] = result.Data
	}

	// Export the scraped data to a CSV file using the ExportToCSV function.
	err := exporter.ExportToCSV(*outputFile, scrapedData)
	if err != nil {
		// Log any error that occurs during the CSV export process.
		log.Fatalf("Error exporting to CSV: %v", err)
	}

	// Log the total scraping time and CSV export completion.
	fmt.Printf("Scraping completed in %v. Data saved to %s\n", time.Since(startTime), *outputFile)
}
