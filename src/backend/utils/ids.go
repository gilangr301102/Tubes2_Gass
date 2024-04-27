package utils

import (
	"backend/wikirace/models"
	"log"
	"runtime"
	"sync"
)

var numNodesPerLevel int = runtime.NumCPU() * 10

// getPathIDDFS retrieves the path using IDDFS.
func getPathIDDFS(articleToParent *map[string]string, endArticle string) *[]string {
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

// IDDFS performs Iterative Deepening Depth-First Search.
func IDDFS(startArticle string, endArticle string, maxDepth int, articleToParent *map[string]string, outputCh chan models.ArticleInfo1, wg *sync.WaitGroup) bool {
	if startArticle == endArticle {
		return true
	}
	if maxDepth <= 0 {
		return false
	}
	for articleWithParent := range outputCh {
		nextArticle := articleWithParent.Article
		if (*articleToParent)[nextArticle] == "" {
			(*articleToParent)[nextArticle] = articleWithParent.ParentArticle
			if IDDFS(nextArticle, endArticle, maxDepth-1, articleToParent, outputCh, wg) {
				return true
			}
		}
	}
	return false
}

// GetShortestPathIDDFS finds the shortest path using IDDFS.
func GetShortestPathIDDFS(startUrl string, endUrl string, maxDepth int) (*[]string, *map[string]string, string) {
	articleToParent := make(map[string]string)
	emptyPath := make([]string, 0)
	maxDepthReachedMsg := ""

	startArticle, err := GetArticleNameFromURLString(startUrl)
	if err != nil || !IsReachable(startUrl) {
		log.Printf("Invalid StartURL: %s\n", startUrl)
		return &emptyPath, &articleToParent, "Invalid StartURL"
	}
	endArticle, err := GetArticleNameFromURLString(endUrl)
	if err != nil || !IsReachable(endUrl) {
		log.Printf("Invalid EndURL: %s\n", endUrl)
		return &emptyPath, &articleToParent, "Invalid EndURL"
	}
	articleToParent[startArticle] = "root"
	if startUrl == endUrl {
		emptyPath = append(emptyPath, startUrl)
		return &emptyPath, &articleToParent, "Start and End URLs are the same"
	}

	outputCh := make(chan models.ArticleInfo1)
	var wg sync.WaitGroup
	wg.Add(1)
	go scrappedArticleAndSync(startArticle, outputCh, &wg)
	go closeChannelOnWg(outputCh, &wg)

	level := 0
	found := false

	for {
		scrappedDatas := make([]string, 0)
		for articleWithParent := range outputCh {
			nextArticle := articleWithParent.Article
			if articleToParent[nextArticle] == "" {
				articleToParent[nextArticle] = articleWithParent.ParentArticle
				if nextArticle == endArticle {
					found = true
					break
				}
				scrappedDatas = append(scrappedDatas, nextArticle)
			}
		}
		log.Printf("Level %d\n", level)
		if found {
			log.Printf("Successfully Found the Solution!")
			break
		}
		log.Printf(
			"Collected %d outputs from scrapping data the child\n",
			len(scrappedDatas))
		level++

		inputCh := make(chan string)
		nextOutputCh := make(chan models.ArticleInfo1, 1000)
		var nextWg sync.WaitGroup
		nextWg.Add(numNodesPerLevel)
		for i := 0; i < numNodesPerLevel; i++ {
			go scrappedArticlesAndSync(inputCh, nextOutputCh, &nextWg)
		}

		go closeChannelOnWg(nextOutputCh, &nextWg)
		log.Printf("Level %d: Started %d Scrapeds\n", level, numNodesPerLevel)

		go feedArticlesIntoChannel(scrappedDatas, inputCh)
		outputCh = nextOutputCh

		if level >= maxDepth {
			maxDepthReachedMsg = "Maximum Depth Reached!"
			log.Printf(maxDepthReachedMsg)
			break
		}

		// Perform IDDFS at the current level
		if IDDFS(startArticle, endArticle, level, &articleToParent, outputCh, &nextWg) {
			found = true
			break
		}
	}

	if maxDepthReachedMsg != "" {
		return &emptyPath, &articleToParent, maxDepthReachedMsg
	}

	return getPathIDDFS(&articleToParent, endArticle), &articleToParent, ""
}
