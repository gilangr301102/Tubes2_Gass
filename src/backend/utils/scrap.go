package utils

import (
	"backend/wikirace/models"
	"net/http"

	"golang.org/x/net/html"
)

func getChilds(article string, ch chan models.ArticleInfo1) {
	if Articles.Contains(article) {
		value, _ := Articles.Get(article)
		childArticles := value.(*[]string)
		for _, childArticle := range *childArticles {
			ch <- models.ArticleInfo1{childArticle, article}
		}
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
							ch <- models.ArticleInfo1{Article: nextArticle, ParentArticle: article}
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
}
