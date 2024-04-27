package utils

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
)

// GetArticleNameFromURLString extracts the article name from the input URL string.
func GetArticleNameFromURLString(inputUrl string) (string, error) {
	parsedUrl, err := url.Parse(inputUrl)
	if err != nil {
		return "", err
	}

	return GetArticleNameFromParsedUrl(parsedUrl)
}

// GetArticleNameFromParsedUrl retrieves the article name from the parsed URL.
func GetArticleNameFromParsedUrl(inputUrl *url.URL) (string, error) {
	if inputUrl != nil &&
		(inputUrl.Host == "" || inputUrl.Host == "en.wikipedia.org") &&
		strings.HasPrefix(inputUrl.Path, "/wiki/") &&
		!noToVisit[inputUrl.Path] {
		return inputUrl.Path[6:], nil
	}
	return "", errors.New("invalid URL format for Wikipedia article")
}

// PanicIfError panics if the input error is not nil.
func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

// IsReachable checks if the input URL is reachable (returns a 200 OK status).
func IsReachable(inputUrl string) bool {
	resp, err := http.Get(inputUrl)
	if err != nil {
		return false
	}
	resp.Body.Close()
	return resp.StatusCode == 200
}

func handleUnderScore(name string) string {
	return strings.ReplaceAll(name, " ", "_")
}

func getPath(articleToParent *map[string]string, endArticle string) *[]string {
	// Initialize an empty slice to store the reversed path
	reversedPath := make([]string, 0)

	// Start from the end article and trace back to the root
	currentArticle := endArticle
	for currentArticle != "root" {
		// Add the current article to the reversed path
		reversedPath = append(reversedPath, currentArticle)
		// Move to the parent article
		currentArticle = (*articleToParent)[currentArticle]
	}

	// Calculate the length of the reversed path
	pathLen := len(reversedPath)

	// Initialize a new slice to store the final path
	path := make([]string, pathLen)

	// Reverse the order of articles to get the correct path
	for i := 0; i < pathLen; i++ {
		// Get the URL string from the reversed path and store it in the correct order
		path[i] = URL_SCRAPPING_WIKIPEDIA + reversedPath[pathLen-i-1]
	}

	// Return a pointer to the final path slice
	return &path
}
