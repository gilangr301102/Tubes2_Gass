package utils

import (
	"log"
	"runtime"
	"sync"
)

var numNodesPerLevelBFS int = runtime.NumCPU() * 10

func getSinglePathBFS(articleToParent *map[string]string, endArticle string) *[]string {
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
		path[i] = reversedPath[pathLen-i-1]
	}

	// Return a pointer to the final path slice
	return &path
}

func GetShortestSinglePathBFS(startUrl string, endUrl string) (*[]string, int, int) {
	// Initialize the map to store the parent of each article
	articleToParent := make(map[string]string)
	emptyPath := make([]string, 0)

	// start url and end url should be reachable
	startArticle, err := GetArticleNameFromURLString(startUrl)

	// if the start url is not reachable, return empty path
	if err != nil || !IsReachable(startUrl) {
		log.Printf("Invalid StartURL: %s\n", startUrl)
		return &emptyPath, 0, 0
	}
	//	start url and end url should be reachable
	endArticle, err := GetArticleNameFromURLString(endUrl)

	// if the end url is not reachable, return empty path
	if err != nil || !IsReachable(endUrl) {
		log.Printf("Invalid EndURL: %s\n", endUrl)
		return &emptyPath, 0, 0
	}
	// Initialize the parent of the start article as "root"
	articleToParent[startArticle] = "root"

	// If the start and end articles are the same, return the path with the start article
	if startUrl == endUrl {
		emptyPath = append(emptyPath, startUrl)
		return &emptyPath, 0, 1
	}

	// Initialize the output channel and wait group
	outputCh := make(chan ArticleInfo1)
	// Wait group to wait for all the goroutines to finish
	var wg sync.WaitGroup
	wg.Add(1)

	numOfArticlesChecked := 0
	articlesLoop := 1
	// Start the first goroutine to scrap the start article
	go scrappedArticleAndSync(startArticle, outputCh, &wg)
	go closeChannelOnWg(outputCh, &wg)

	level := 0
	found := false

	for {
		// Collect outputs from scrapping at the current level
		// Skip articles that are already visited
		scrappedDatas := make([]string, 0)
		tempCountLeaves := 0
		for articleWithParent := range outputCh {
			nextArticle := articleWithParent.Article
			if articleToParent[nextArticle] == "" {
				articleToParent[nextArticle] = articleWithParent.ParentArticle
				articlesLoop++
				tempCountLeaves++
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
			numOfArticlesChecked = articlesLoop - tempCountLeaves
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
		nextOutputCh := make(chan ArticleInfo1, 1000)
		var nextWg sync.WaitGroup
		nextWg.Add(numNodesPerLevelBFS)
		for i := 0; i < numNodesPerLevelBFS; i++ {
			go scrappedArticlesAndSync(inputCh, nextOutputCh, &nextWg)
		}

		// close the output channel when all the goroutines are done
		go closeChannelOnWg(nextOutputCh, &nextWg)
		log.Printf("Level %d: Started %d Scrapeds\n", level, numNodesPerLevelBFS)

		// feed articles into the input channel
		go feedArticlesIntoChannel(scrappedDatas, inputCh)
		outputCh = nextOutputCh
	}

	path := getSinglePathBFS(&articleToParent, endArticle)
	return path, numOfArticlesChecked, articlesLoop
}

func getMultiPathBFS(articleToParent *map[string][]ArticleNode, endArticle string) (*[][]string, int) {
	// Initialize an empty slice to store the reversed path
	reversedPath := make([][]string, 0)

	// Start from the end article and trace back to the root
	for _, articleParent := range (*articleToParent)[endArticle] {
		tempReversedPath := make([]string, 0)
		tempReversedPath = append(tempReversedPath, endArticle)
		currentArticle := articleParent.Name
		for currentArticle != "root" {
			// Add the current article to the reversed path
			tempReversedPath = append(tempReversedPath, currentArticle)
			// Move to the parent article
			currentArticle = (*articleToParent)[currentArticle][0].Name
		}
		reversedPath = append(reversedPath, tempReversedPath)
	}

	// Calculate the length of the reversed path
	pathLen := len(reversedPath)

	// Initialize a new slice to store the final path
	path := make([][]string, pathLen)

	// Get All the path
	for i := 0; i < pathLen; i++ {
		tempPath := make([]string, 0)
		subPathLen := len(reversedPath[i])
		// Reverse the order of articles to get the correct path
		for j := 0; j < subPathLen; j++ {
			tempPath = append(tempPath, reversedPath[i][subPathLen-j-1])
		}
		path[i] = tempPath
	}

	// Return a pointer to the final path slice
	return &path, pathLen
}

func GetShortestMultiPathBFS(startUrl string, endUrl string) (*[][]string, int, int, int) {
	// Initialize the map to store the parent of each article
	articleToParent := make(map[string][]ArticleNode)
	emptyPath := make([][]string, 0)

	// start url and end url should be reachable
	startArticle, err := GetArticleNameFromURLString(startUrl)

	// if the start url is not reachable, return empty path
	if err != nil || !IsReachable(startUrl) {
		log.Printf("Invalid StartURL: %s\n", startUrl)
		return &emptyPath, 0, 0, 0
	}
	//	start url and end url should be reachable
	endArticle, err := GetArticleNameFromURLString(endUrl)

	// if the end url is not reachable, return empty path
	if err != nil || !IsReachable(endUrl) {
		log.Printf("Invalid EndURL: %s\n", endUrl)
		return &emptyPath, 0, 0, 0
	}
	// Initialize the parent of the start article as "root"
	articleToParent[startArticle] = append(articleToParent[startArticle], ArticleNode{"root", -1})

	// If the start and end articles are the same, return the path with the start article
	if startUrl == endUrl {
		tempEmptyPath := make([]string, 0)
		tempEmptyPath = append(tempEmptyPath, startUrl)
		emptyPath = append(emptyPath, tempEmptyPath)
		return &emptyPath, 1, 0, 1
	}

	// Initialize the output channel and wait group
	outputCh := make(chan ArticleInfo1)
	// Wait group to wait for all the goroutines to finish
	var wg sync.WaitGroup
	wg.Add(1)
	// Start the first goroutine to scrap the start article
	numOfArticlesChecked := 0
	articlesLoop := 1
	go scrappedArticleAndSync(startArticle, outputCh, &wg)
	go closeChannelOnWg(outputCh, &wg)

	level := 1
	found := false
	for {
		// Collect outputs from scrapping at the current level
		// Skip articles that are already visited
		scrappedDatas := make([]string, 0)
		tempCountLeaves := 0
		for articleWithParent := range outputCh {
			flag := false
			nextArticle := articleWithParent.Article
			tempCurrArticleToParent := articleToParent[nextArticle]
			if len(tempCurrArticleToParent) > 0 {
				for _, articleParent := range tempCurrArticleToParent {
					if level-1 > articleParent.Level ||
						articleParent.Name == articleWithParent.ParentArticle {
						flag = true
						break
					}
				}
			}
			if !flag {
				articlesLoop++
				tempCountLeaves++
				articleToParent[nextArticle] = append(articleToParent[nextArticle],
					ArticleNode{articleWithParent.ParentArticle, level - 1})
				scrappedDatas = append(scrappedDatas, nextArticle)
				if nextArticle == endArticle {
					found = true
				}
			}
		}
		log.Printf("Level %d\n", level)
		// If the solution is found, break the loop
		if found {
			log.Printf("Successfully Found the Solution!")
			numOfArticlesChecked = articlesLoop - tempCountLeaves
			break
		}

		// If no more articles to scrap, break the loop
		log.Printf(
			"Collected %d outputs from scrapping data the child\n",
			len(scrappedDatas))
		level++

		// Start the next level
		inputCh := make(chan string)
		nextOutputCh := make(chan ArticleInfo1, 1000)
		var nextWg sync.WaitGroup
		nextWg.Add(numNodesPerLevelBFS)
		for i := 0; i < numNodesPerLevelBFS; i++ {
			go scrappedArticlesAndSync(inputCh, nextOutputCh, &nextWg)
		}

		// close the output channel when all the goroutines are done
		go closeChannelOnWg(nextOutputCh, &nextWg)
		log.Printf("Level %d: Started %d Scrapeds\n", level, numNodesPerLevelBFS)

		// feed articles into the input channel
		go feedArticlesIntoChannel(scrappedDatas, inputCh)
		outputCh = nextOutputCh
	}

	path, numberPath := getMultiPathBFS(&articleToParent, endArticle)
	return path, numberPath, numOfArticlesChecked, articlesLoop
}
