package main

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	lru "github.com/hashicorp/golang-lru"
)

var prefixes = [...]string{"/wiki/Main_Page", "/wiki/Special", "/wiki/File", "/wiki/Help", "/wiki/Wikipedia:"}

type Node struct {
	ParentIndex int
	CurrIndex   int
}

type NodeInfo struct {
	Title string
	Depth int
}

var cache *lru.Cache
var wg sync.WaitGroup

func init() {
	// Initialize LRU cache with a capacity of 1000 entries
	var err error
	cache, err = lru.New(1000)
	if err != nil {
		log.Fatal(err)
	}
}

func checkListOfPrefixes(href string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(href, prefix) {
			return true
		}
	}
	return false
}

func toTitleCase(s string) string {
	return strings.ReplaceAll(s, "_", " ")
}

func getNeighbours(url string, ch chan<- []string) {
	defer wg.Done()

	if val, ok := cache.Get(url); ok {
		if neighbors, ok := val.([]string); ok {
			ch <- neighbors
			return
		}
	}

	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	var links []string
	doc.Find("#mw-content-text").Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists && strings.HasPrefix(href, "/wiki/") && !checkListOfPrefixes(href) {
			title := toTitleCase(strings.TrimPrefix(href, "/wiki/"))
			links = append(links, title)
		}
	})

	cache.Add(url, links)
	ch <- links
}

func containsNode(nodes []NodeInfo, node string) bool {
	for _, n := range nodes {
		if n.Title == node {
			return true
		}
	}
	return false
}

func shortestPathBFS(startArticle string, goalArticle string) [][]string {
	var tempQ1 []Node
	var tempQ2 []Node
	var uniqueNodes []NodeInfo
	var arrIndexNodes []Node
	var solution [][]string
	isGetSolution := false
	maxUniqueIndex := 0
	tempQ1 = append(tempQ1, Node{ParentIndex: -1, CurrIndex: maxUniqueIndex})
	uniqueNodes = append(uniqueNodes, NodeInfo{Title: startArticle, Depth: 0})
	arrIndexNodes = append(arrIndexNodes, Node{ParentIndex: -1, CurrIndex: maxUniqueIndex})
	for len(tempQ1) > 0 {
		currNode := tempQ1[0]
		if uniqueNodes[currNode.CurrIndex].Title == goalArticle {
			var path []string
			for uniqueNodes[currNode.CurrIndex].Depth != 0 {
				path = append(path, uniqueNodes[currNode.CurrIndex].Title)
				currNode = arrIndexNodes[currNode.ParentIndex]
			}
			path = append(path, startArticle)
			solution = append(solution, path)
			isGetSolution = true
			return solution
		}
		tempQ2 = append(tempQ2, tempQ1[0])
		tempQ1 = tempQ1[1:]

		if len(tempQ1) == 0 {
			if !isGetSolution {
				count := 0
				for len(tempQ2) > 0 {
					currNode = tempQ2[0]
					tempQ2 = tempQ2[1:]
					ch := make(chan []string)
					wg.Add(1)
					go getNeighbours("https://en.wikipedia.org/wiki/"+uniqueNodes[currNode.CurrIndex].Title, ch)
					neighbours := <-ch
					wg.Wait()
					close(ch)
					for _, neighbour := range neighbours {
						isExist := false
						countUniqueNode := 0
						for _, n := range uniqueNodes {
							if n.Title == neighbour {
								isExist = true
								if n.Depth < uniqueNodes[currNode.CurrIndex].Depth+1 {
									break
								}
								if n.Depth == uniqueNodes[currNode.CurrIndex].Depth+1 {
									tempQ1 = append(tempQ1, Node{ParentIndex: currNode.CurrIndex, CurrIndex: countUniqueNode})
									arrIndexNodes = append(arrIndexNodes, Node{ParentIndex: currNode.CurrIndex, CurrIndex: countUniqueNode})
									break
								}
							}
							countUniqueNode++
						}
						if !isExist {
							maxUniqueIndex += 1
							tempQ1 = append(tempQ1, Node{ParentIndex: currNode.CurrIndex, CurrIndex: maxUniqueIndex})
							uniqueNodes = append(uniqueNodes, NodeInfo{Title: neighbour, Depth: uniqueNodes[currNode.CurrIndex].Depth + 1})
							arrIndexNodes = append(arrIndexNodes, Node{ParentIndex: currNode.CurrIndex, CurrIndex: maxUniqueIndex})
						}
					}
					count++
				}
			}
		}
	}
	return solution
}

func main() {
	startTime := time.Now()
	var test = shortestPathBFS("Sweden", "Dracula")
	fmt.Println(test)
	elapsed := time.Since(startTime)
	fmt.Println("Time elapsed: ", elapsed)
}
