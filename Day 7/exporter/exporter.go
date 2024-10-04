package exporter

import (
	"encoding/csv"
	"fmt"
	"os"
)

// ExportToCSV writes scraped data to a CSV file.
func ExportToCSV(filename string, data map[string][]string) error {
	// Create (or overwrite) a new CSV file with the given filename.
	file, err := os.Create(filename)
	if err != nil {
		// If there's an error creating the file, log and return it.
		return fmt.Errorf("error creating CSV file: %v", err)
	}
	defer file.Close() // Ensure the file is closed after writing.

	// Create a new CSV writer to write to the file.
	writer := csv.NewWriter(file)
	defer writer.Flush() // Ensure all data is written (flushed) at the end.

	// Write the header row to the CSV file.
	writer.Write([]string{"URL", "Data"})

	// Iterate over the scraped data map, where the key is the URL and the value is the list of data.
	for url, items := range data {
		// For each item (e.g., each heading), write a row to the CSV file.
		for _, item := range items {
			writer.Write([]string{url, item}) // Write the URL and the corresponding item (e.g., heading).
		}
	}

	// Return no error, indicating success.
	return nil
}
