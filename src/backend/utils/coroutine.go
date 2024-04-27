package utils

import (
	"sync"
)

func scrappedArticleAndSync(article string, outputCh chan ArticleInfo, wg *sync.WaitGroup) {
	getChilds2(article, outputCh)
	(*wg).Done()
}

func scrappedArticlesAndSync(inputCh chan string, outputCh chan ArticleInfo, wg *sync.WaitGroup) {
	for article := range inputCh {
		getChilds2(article, outputCh)
	}
	(*wg).Done()
}

func closeChannelOnWg(ch chan ArticleInfo, wg *sync.WaitGroup) {
	(*wg).Wait()
	close(ch)
}

func feedArticlesIntoChannel(articles []string, ch chan string) {
	for _, article := range articles {
		ch <- article
	}
	close(ch)
}
