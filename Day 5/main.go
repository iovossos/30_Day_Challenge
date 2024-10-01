package main

import (
	"downloader/download" // Import the download package
	"fmt"
	"os"
)

func main() {
	// Check if enough arguments are provided
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <url> <filename>")
		return
	}

	// Get URL and filename from command-line arguments
	url := os.Args[1]
	filename := os.Args[2]

	fmt.Println("Starting the download...")

	// Download the file using the download package
	err := download.DownloadFile(url, filename)
	if err != nil {
		fmt.Printf("\nError downloading file: %v\n", err)
		return
	}

	// Print completion message
	fmt.Println("\nDownload completed successfully!")
}
