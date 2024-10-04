package scraper

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// ScrapePage fetches the HTML content of a URL and extracts data (e.g., <h1> and <h2> tags).
func ScrapePage(url string) ([]string, error) {
	// Send an HTTP GET request to the URL.
	resp, err := http.Get(url)
	if err != nil {
		// If there's an error during the HTTP request, log it and return the error.
		return nil, fmt.Errorf("error fetching URL %s: %v", url, err)
	}
	defer resp.Body.Close()

	// Check if the server returned a status other than 200 (OK).
	if resp.StatusCode != 200 {
		// If not, log and return the error, including the HTTP status code.
		return nil, fmt.Errorf("error: non-200 status code %d received from %s", resp.StatusCode, url)
	}

	// Use goquery to parse the HTML content of the page.
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		// If there's an error while parsing the HTML, log and return it.
		return nil, fmt.Errorf("error parsing HTML for URL %s: %v", url, err)
	}

	// Initialize a slice to store the extracted headings (or any other data).
	var data []string

	// Find all <h1> and <h2> tags in the HTML document and extract the text content.
	doc.Find("h1, h2").Each(func(i int, s *goquery.Selection) {
		heading := s.Text()          // Extract the text content of each <h1> and <h2> tag.
		data = append(data, heading) // Append the heading to the data slice.
	})

	// Return the extracted data (slice of headings) and no error.
	return data, nil
}
