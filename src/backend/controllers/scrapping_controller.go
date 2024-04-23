package controllers

import (
	"encoding/json"
	"example/user/hello/models"
	"example/user/hello/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/mux"
)

func GetScrappingData(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var err error

	params := mux.Vars(r)
	searchName := params["searchName"]

	// Define the URL of the Wikipedia page
	url := utils.URL_SCRAPPING_WIKIPEDIA + searchName

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

	// data := utils.getNeighbors(searchName)

	// Status 200 OK
	if len(data) > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(data); err != nil {
			panic(err)
		}
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)

	if err := json.NewEncoder(w).Encode(models.JsonError{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}
