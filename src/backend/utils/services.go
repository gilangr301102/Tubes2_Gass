package utils

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func getNeighbors(searchSource string, searchGoal string) [][]string {
	url := URL_SCRAPPING_WIKIPEDIA + searchSource

	// Define a user-agent header to simulate a browser request
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36",
	}

	// Create an HTTP client with custom headers
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Send an HTTP GET request to the URL with the headers
	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	data := [][]string{}

	// Check if the request was successful (status code 200)
	if response.StatusCode == 200 {
		// Parse the HTML content of the page using goquery
		doc, err := goquery.NewDocumentFromReader(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		// Iterate through the rows of the table
		doc.Find("tr").Each(func(rowIdx int, row *goquery.Selection) {
			// Skip the header row
			if rowIdx == 0 {
				return
			}

			// Extract data from each column and append it to the data slice
			rowData := []string{}
			row.Find("th,td").Each(func(colIdx int, col *goquery.Selection) {
				rowData = append(rowData, col.Text())
			})
			data = append(data, rowData)
		})
	} else {
		fmt.Println("Failed to retrieve the web page. Status code:", response.StatusCode)
	}

	return data
}
