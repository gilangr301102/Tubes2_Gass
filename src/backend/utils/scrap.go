package utils

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

func getChilds2(article string, ch chan ArticleInfo) {
	if Articles.Contains(article) {
		value, _ := Articles.Get(article)
		childArticles := value.(*[]string)
		for _, childArticle := range *childArticles {
			ch <- ArticleInfo{childArticle, article}
		}
		fmt.Println("Already scrapped1: ", ch)
		return
	}

	resp, err := http.Get(URL_SCRAPPING_WIKIPEDIA + article)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	childArticles := make([]string, 0)
	visited := make(map[string]bool)
	visited[article] = true
	z := html.NewTokenizer(resp.Body)

	for {
		tokenType := z.Next()
		switch tokenType {
		case html.ErrorToken:
			return
		case html.StartTagToken:
			token := z.Token()
			if token.Data == "a" {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						nextArticle, err := GetArticleNameFromURLString(attr.Val)
						if err == nil && !visited[nextArticle] {
							ch <- ArticleInfo{Article: nextArticle, ParentArticle: article}
							visited[nextArticle] = true
							childArticles = append(childArticles, nextArticle)
						}
					}
				}
			}
		}

		if tokenType == html.ErrorToken {
			break
		}
	}

	Articles.Add(article, &childArticles)
	fmt.Println("Already scrapped2: ", Articles)
}

// function to fetch links from a given url
func getChilds(urlString string, ch chan<- []Node, errCh chan<- error) {
	// reserve a spot in the semaphore to limit concurrent HTTP requests
	sem <- struct{}{}
	defer func() { <-sem }() // release the semaphore when it's done

	// cek if the result url has already in cache
	if nodes, ok := urlCache[urlString]; ok {
		ch <- nodes
		return
	}

	// create a new http get request
	req, _ := http.NewRequest("GET", urlString, nil)
	req.Header.Set("Connection", "keep-alive")

	// perform the http request
	res, err := client.Do(req)
	if err != nil {
		errCh <- err
		return
	}
	defer res.Body.Close() // ensure the response body is closed after processing

	// check if the response status code indicates success
	if res.StatusCode != 200 {
		errCh <- fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
		return
	}

	// parse the html content to extract links
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		errCh <- err
		return
	}

	// store extracted nodes
	var nodes []Node
	doc.Find("#mw-content-text a").Each(func(_ int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists && strings.HasPrefix(href, "/wiki/") && !strings.Contains(href, ":") {
			// Extract the title from the URL
			title := strings.TrimPrefix(href, "/wiki/")
			title, err = url.QueryUnescape(title) // Decode URL-encoded characters
			if err != nil {
				errCh <- err
				return
			}
			title = strings.ReplaceAll(title, "_", " ") // Replace underscores with spaces
			fullURL := WIKIPEDIA_URL_EN + href
			nodes = append(nodes, Node{Title: title, URL: fullURL, Path: []string{title}}) // create a new node with the link's details
		}
	})

	// store the result to cache
	urlCache[urlString] = nodes
	ch <- nodes
}
