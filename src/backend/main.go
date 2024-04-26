package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
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

/* checkListOfPrefixes: Checks url against slice of prefixes to ensure the linke is to an article.*/
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

func getNeighbours(url string) []string {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	var links []string
	doc.Find("#mw-content-text").Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists && strings.HasPrefix(href, "/wiki/") && !checkListOfPrefixes(href) {
			title := toTitleCase(strings.TrimPrefix(href, "/wiki/"))
			links = append(links, title)
		}
	})
	return links
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
	var tempQueue []Node
	var uniqueNodes []NodeInfo
	var arrIndexNodes []Node
	// var goalIndexNodes []Node
	var solution [][]string
	// var isGetSolution bool
	maxUniqueIndex := 0
	tempQueue = append(tempQueue, Node{ParentIndex: -1, CurrIndex: 0})
	uniqueNodes = append(uniqueNodes, NodeInfo{Title: startArticle, Depth: 0})
	for len(tempQueue) > 0 {
		currNode := tempQueue[0]
		tempQueue = tempQueue[1:]

		if uniqueNodes[currNode.CurrIndex].Title == goalArticle {
			// Backward traversal to get the path solution
			var path []string
			for uniqueNodes[currNode.CurrIndex].Depth != 0 {
				path = append(path, uniqueNodes[currNode.CurrIndex].Title)
				currNode = arrIndexNodes[currNode.ParentIndex]
			}
			path = append(path, startArticle)
			solution = append(solution, path)
		}

		neighbours := getNeighbours("https://en.wikipedia.org/wiki/" + uniqueNodes[currNode.CurrIndex].Title)
		for _, neighbour := range neighbours {
			if !containsNode(uniqueNodes, neighbour) {
				maxUniqueIndex += 1
				tempQueue = append(tempQueue, Node{ParentIndex: currNode.CurrIndex, CurrIndex: maxUniqueIndex})
				uniqueNodes = append(uniqueNodes, NodeInfo{Title: neighbour, Depth: uniqueNodes[currNode.CurrIndex].Depth + 1})
			}
		}
	}
	return solution
}

func depthLimitedDFS(startNode string, goalNode string, depth int) [][]string {
	var solution [][]string
	var visited []string
	var path []string
	var isGetSolution bool
	var recursiveDFS func(string, string, int)
	recursiveDFS = func(currNode string, goalNode string, depth int) {
		if depth == 0 {
			return
		}
		if currNode == goalNode {
			isGetSolution = true
			path = append(path, currNode)
			solution = append(solution, path)
			return
		}
		visited = append(visited, currNode)
		neighbours := getNeighbours("https://en.wikipedia.org/wiki/" + currNode)
		for _, neighbour := range neighbours {
			if !isGetSolution && !containsNode(visited, neighbour) {
				path = append(path, currNode)
				recursiveDFS(neighbour, goalNode, depth-1)
				path = path[:len(path)-1]
			}
		}
	}
	recursiveDFS(startNode, goalNode, depth)
	return solution
}

func shortestPathIDS(startNode string, goalNode string) [][]string {
	var depth int
	for {
		path := depthLimitedDFS(startNode, goalNode, depth)
		if path != nil {
			return path
		}
		depth++
	}
}

func main() {
	wikipediaURL := "https://en.wikipedia.org/wiki/Indonesia"

	var test = getNeighbours(wikipediaURL)
	fmt.Println(test)

}
