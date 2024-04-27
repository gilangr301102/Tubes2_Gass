package utils

import (
	"backend/wikirace/models"
	"sync"
)

func scrappedArticleAndSync(article string, outputCh chan models.ArticleInfo1, wg *sync.WaitGroup) {
	getChilds(article, outputCh)
	(*wg).Done()
}

func scrappedArticlesAndSync(inputCh chan string, outputCh chan models.ArticleInfo1, wg *sync.WaitGroup) {
	for article := range inputCh {
		getChilds(article, outputCh)
	}
	(*wg).Done()
}

func closeChannelOnWg(ch chan models.ArticleInfo1, wg *sync.WaitGroup) {
	(*wg).Wait()
	close(ch)
}

func feedArticlesIntoChannel(articles []string, ch chan string) {
	for _, article := range articles {
		ch <- article
	}
	close(ch)
}
