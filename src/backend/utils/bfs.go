package utils

import (
	"example/user/hello/models"
	"log"
	"runtime"
	"sync"
)

var numNodesPerLevel int = runtime.NumCPU() * 10

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

func GetShortestPathBFS(startUrl string, endUrl string) (*[]string, *map[string]string) {
	// Initialize the map to store the parent of each article
	articleToParent := make(map[string]string)
	emptyPath := make([]string, 0)

	// start url and end url should be reachable
	startArticle, err := GetArticleNameFromURLString(startUrl)

	// if the start url is not reachable, return empty path
	if err != nil || !IsReachable(startUrl) {
		log.Printf("Invalid StartURL: %s\n", startUrl)
		return &emptyPath, &articleToParent
	}
	//	start url and end url should be reachable
	endArticle, err := GetArticleNameFromURLString(endUrl)

	// if the end url is not reachable, return empty path
	if err != nil || !IsReachable(endUrl) {
		log.Printf("Invalid EndURL: %s\n", endUrl)
		return &emptyPath, &articleToParent
	}
	// Initialize the parent of the start article as "root"
	articleToParent[startArticle] = "root"

	// If the start and end articles are the same, return the path with the start article
	if startUrl == endUrl {
		emptyPath = append(emptyPath, startUrl)
		return &emptyPath, &articleToParent
	}

	// Initialize the output channel and wait group
	outputCh := make(chan models.ArticleInfo1)
	// Wait group to wait for all the goroutines to finish
	var wg sync.WaitGroup
	wg.Add(1)
	// Start the first goroutine to scrap the start article
	go scrappedArticleAndSync(startArticle, outputCh, &wg)
	go closeChannelOnWg(outputCh, &wg)

	level := 0
	found := false

	for {
		// Collect outputs from scrapping at the current level
		// Skip articles that are already visited
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
		// If the solution is found, break the loop
		if found {
			log.Printf("Successfully Found the Solution!")
			break
		}

		// If no more articles to scrap, break the loop
		log.Printf(
			"Collected %d outputs from scrapping data the child\n",
			len(scrappedDatas))
		level++

		// Start the next level
		inputCh := make(chan string)
		nextOutputCh := make(chan models.ArticleInfo1, 1000)
		var nextWg sync.WaitGroup
		nextWg.Add(numNodesPerLevel)
		for i := 0; i < numNodesPerLevel; i++ {
			go scrappedArticlesAndSync(inputCh, nextOutputCh, &nextWg)
		}

		// close the output channel when all the goroutines are done
		go closeChannelOnWg(nextOutputCh, &nextWg)
		log.Printf("Level %d: Started %d Scrapeds\n", level, numNodesPerLevel)

		// feed articles into the input channel
		go feedArticlesIntoChannel(scrappedDatas, inputCh)
		outputCh = nextOutputCh
	}

	return getPath(&articleToParent, endArticle), &articleToParent
}
