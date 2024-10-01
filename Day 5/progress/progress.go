package progress

import (
	"fmt"
	"io"
	"time"
)

// ProgressReader wraps an io.Reader to track and display the progress of the file download
type ProgressReader struct {
	Reader     io.Reader
	TotalBytes int64
	Downloaded int64
	StartTime  time.Time // Track when the download started
}

// Read method tracks the number of bytes downloaded and displays the progress bar
func (pr *ProgressReader) Read(p []byte) (int, error) {
	n, err := pr.Reader.Read(p)
	pr.Downloaded += int64(n)
	pr.displayProgress()
	return n, err
}

// displayProgress prints the download progress as a bar and the size in megabytes (on the same line)
func (pr *ProgressReader) displayProgress() {
	// Convert bytes to megabytes (MB)
	downloadedMB := float64(pr.Downloaded) / (1024 * 1024)

	// If TotalBytes is not provided, just show the downloaded amount
	if pr.TotalBytes <= 0 {
		fmt.Printf("\rDownloaded: %.2f MB", downloadedMB)
		return
	}

	// Calculate the percentage of download completed
	percent := float64(pr.Downloaded) / float64(pr.TotalBytes) * 100
	if percent > 100 {
		percent = 100 // Cap the percentage at 100%
	}

	// Create the progress bar with 50 segments
	barLength := 50
	progress := int(percent / 2) // 50 segments, so we divide by 2
	bar := "[" + fillBar(progress, barLength) + "]"

	totalMB := float64(pr.TotalBytes) / (1024 * 1024)

	// Calculate the elapsed time and estimate remaining time
	elapsed := time.Since(pr.StartTime).Seconds()
	var remainingTime float64
	mbPerSecond := 0.0
	if elapsed > 0 {
		mbPerSecond = downloadedMB / elapsed
		if mbPerSecond > 0 {
			remainingTime = (totalMB - downloadedMB) / mbPerSecond
		} else {
			remainingTime = 0
		}
	}

	// Use carriage return '\r' to overwrite the current line in the terminal
	fmt.Printf("\r%s %.2f%% %.2f/%.2f MB, %.2f MB/s, ETA: %.0f seconds", bar, percent, downloadedMB, totalMB, mbPerSecond, remainingTime)
}

// fillBar fills the progress bar with "=" for completed parts and " " for incomplete parts
func fillBar(progress, total int) string {
	bar := ""
	for i := 0; i < total; i++ {
		if i < progress {
			bar += "="
		} else {
			bar += " "
		}
	}
	return bar
}
