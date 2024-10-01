package download

import (
	"downloader/progress" // Import the progress package
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// DownloadFile fetches the file from the given URL and saves it to the specified filename
func DownloadFile(url string, filename string) error {
	// Create the output file
	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()

	// Create an HTTP client and request
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	// Set a User-Agent header to simulate a browser request
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")

	// Send the HTTP GET request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check if the server response is OK
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Get the total file size from the "Content-Length" header
	totalBytes := resp.ContentLength

	// Create a ProgressReader to monitor the download progress
	progressReader := &progress.ProgressReader{
		Reader:     resp.Body,
		TotalBytes: totalBytes,
		StartTime:  time.Now(), // Initialize start time
	}

	// Copy the file from the response to the local file while tracking progress
	_, err = io.Copy(out, progressReader)
	if err != nil {
		return err
	}

	return nil
}
