package utils

import (
	"example/user/hello/models"
	"log"
	"sync"
)

// var numNodesPerLevelIDS int = runtime.NumCPU() * 10

func getPathIDS(articleToParent *map[string]string, endArticle string) *[]string {
	reversedPath := make([]string, 0)
	currentArticle := endArticle
	for currentArticle != "root" {
		reversedPath = append(reversedPath, currentArticle)
		currentArticle = (*articleToParent)[currentArticle]
	}
	pathLen := len(reversedPath)
	path := make([]string, pathLen)
	for i := 0; i < pathLen; i++ {
		path[i] = URL_SCRAPPING_WIKIPEDIA + reversedPath[pathLen-i-1]
	}
	return &path
}

func performIDS(startArticle string, endArticle string, depthLimit int, articleToParent *map[string]string, outputCh chan models.ArticleInfo1, wg *sync.WaitGroup) {
	defer wg.Done()
	if startArticle == endArticle {
		log.Printf("Successfully Found the Solution!")
		return
	}
	if depthLimit <= 0 {
		return
	}
	scrappedDatas := make([]string, 0)
	go scrappedArticleAndSync(startArticle, outputCh, nil)
	for articleWithParent := range outputCh {
		nextArticle := articleWithParent.Article
		if nextArticle != startArticle {
			scrappedDatas = append(scrappedDatas, nextArticle)
		}
	}
	log.Printf("Depth %d: Collected %d outputs from scrapping data the child\n", depthLimit, len(scrappedDatas))
	for _, article := range scrappedDatas {
		(*articleToParent)[article] = startArticle
		wg.Add(1)
		go performIDS(article, endArticle, depthLimit-1, articleToParent, outputCh, wg)
	}
}

func GetShortestPathIDS(startUrl string, endUrl string, maxDepth int) (*[]string, *map[string]string) {
	articleToParent := make(map[string]string)
	emptyPath := make([]string, 0)

	startArticle, err := GetArticleNameFromURLString(startUrl)
	if err != nil || !IsReachable(startUrl) {
		log.Printf("Invalid StartURL: %s\n", startUrl)
		return &emptyPath, &articleToParent
	}

	endArticle, err := GetArticleNameFromURLString(endUrl)
	if err != nil || !IsReachable(endUrl) {
		log.Printf("Invalid EndURL: %s\n", endUrl)
		return &emptyPath, &articleToParent
	}

	articleToParent[startArticle] = "root"

	if startUrl == endUrl {
		emptyPath = append(emptyPath, startUrl)
		return &emptyPath, &articleToParent
	}

	outputCh := make(chan models.ArticleInfo1)
	var wg sync.WaitGroup
	wg.Add(1)
	go performIDS(startArticle, endArticle, maxDepth, &articleToParent, outputCh, &wg)
	wg.Wait()

	if articleToParent[endArticle] == "" {
		log.Printf("Solution not found within depth limit.")
		return &emptyPath, &articleToParent
	}

	return getPathIDS(&articleToParent, endArticle), &articleToParent
}
