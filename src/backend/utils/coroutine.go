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

func threadScrappingProcessing(resultQueue chan<- []Node, errQueue chan<- error, urlQueue <-chan string, wg *sync.WaitGroup) {
	for url := range urlQueue {
		ch := make(chan []Node)
		errCh := make(chan error)

		go getChilds(url, ch, errCh)

		select {
		case result := <-ch:
			resultQueue <- result
		case err := <-errCh:
			errQueue <- err
		}
	}

	wg.Done()
}

func multithreadScrappingProcessing(resultQueue chan<- []Node, errQueue chan<- error, urlQueue <-chan string, wg *sync.WaitGroup) {
	for i := 0; i < NumOfNodeWORKERS; i++ {
		wg.Add(1)
		go threadScrappingProcessing(resultQueue, errQueue, urlQueue, wg)
	}
}
