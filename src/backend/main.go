package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var prefixes = [...]string{"/wiki/Main_Page", "/wiki/Special", "/wiki/File", "/wiki/Help", "/wiki/Wikipedia:"}

/* checkListOfPrefixes: Checks url against slice of prefixes to ensure the linke is to an article.*/
func checkListOfPrefixes(href string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(href, prefix) {
			return true
		}
	}
	return false
}

func toTitleCase(s string) string {
	return strings.ReplaceAll(s, "_", " ")
}

func getNeighbours(url string) []string {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	var links []string
	doc.Find("#mw-content-text").Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists && strings.HasPrefix(href, "/wiki/") && !checkListOfPrefixes(href) {
			title := toTitleCase(strings.TrimPrefix(href, "/wiki/"))
			links = append(links, title)
		}
	})
	return links
}

func main() {
	wikipediaURL := "https://en.wikipedia.org/wiki/Indonesia"

	var test = getNeighbours(wikipediaURL)
	fmt.Println(test)

}
