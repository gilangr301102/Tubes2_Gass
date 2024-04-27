package utils

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"
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

func convertStrToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return num
}
