package utils

import (
	"fmt"
	"log"
	"strings"
	"sync"
)

func isSearchCondition(flag bool, depth int, foundCount int) bool {
	if !flag {
		return foundCount == 0
	}

	return depth <= foundCount
}

// IDS Algorithm
func getShortestPathIDS(start Node, target Node, maxDepth int, findAll bool) (int, int, int, [][]string) {
	var numOfArticlesChecked int
	var articlesLoop int
	var foundCount int = maxDepth + 1

	if !findAll {
		foundCount = 0
	}

	urlQueue := make(chan string, 500)
	resultQueue := make(chan []Node, 500)
	errQueue := make(chan error, 500)
	results := make([][]string, 0)
	pathSet := make(map[string]bool)

	// create multithread to get the scrapping data
	var wg sync.WaitGroup
	multithreadScrappingProcessing(resultQueue, errQueue, urlQueue, &wg)

	// iterate through increasing depth
	for depth := 0; depth <= maxDepth && isSearchCondition(findAll, depth, foundCount); depth++ {
		visited := make(map[string]bool)
		var stack []Node

		// set the initial path for the start node
		start.Path = []string{start.Title}
		stack = append(stack, start)

		// iterate through the stack
		for len(stack) > 0 && (foundCount == 0 || findAll) {
			// pop the last element from the stack
			current := stack[len(stack)-1]
			pathKey := strings.Join(current.Path, "^")
			stack = stack[:len(stack)-1]

			// check if the current node is the target
			if current.Title == target.Title && !pathSet[pathKey] {
				results = append(results, current.Path)
				pathSet[pathKey] = true
				if findAll && depth < foundCount {
					foundCount = depth
				} else if !findAll {
					foundCount = 1
					break
				}
			}
			if len(current.Path) > depth || visited[current.URL] {
				continue
			}

			// mark the current node as visited
			visited[current.URL] = true
			urlQueue <- current.URL

			// get the neighbors of the current node
			select {
			case neighbors, ok := <-resultQueue:
				if !ok {
					// resultQueue has been closed
					break
				}
				for _, neighbor := range neighbors {
					if !visited[neighbor.URL] {
						neighbor.Path = append([]string(nil), current.Path...)
						neighbor.Path = append(neighbor.Path, neighbor.Title)
						stack = append(stack, neighbor)
						articlesLoop++
					}
				}
				numOfArticlesChecked++
			case err := <-errQueue:
				log.Println(err)
				// Optionally handle error and continue the loop
			}
		}
	}

	close(urlQueue)
	wg.Wait()
	close(resultQueue)
	close(errQueue)

	numberofPath := len(results)

	if numberofPath != 0 {
		fmt.Printf("Successfully get the Result: %v\n", results)
	} else {
		fmt.Println("No path found.")
	}

	return numOfArticlesChecked, articlesLoop, numberofPath, results
}
