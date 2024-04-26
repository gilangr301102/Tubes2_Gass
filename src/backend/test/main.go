// // package main

// // import (
// // 	"example/user/hello/routes"
// // 	"fmt"
// // 	"log"
// // 	"net/http"

// // 	"github.com/urfave/negroni"
// // )

// // func main() {
// // 	// port := "8080"

// // 	// if port == "" {
// // 	// 	log.Fatal("$PORT must be set")
// // 	// }

// // 	// // Setup for Router, Endpoints, Handlers and middleware
// // 	// router := routes.GetAllRoutes()
// // 	// middleWare := negroni.Classic()
// // 	// middleWare.UseHandler(router)

// // 	// // Serves API - Creates a new thread and if fails it will log the error.
// // 	// fmt.Println("Server is running on http://localhost:8080")
// // 	// log.Fatal(http.ListenAndServe(":"+port, middleWare))

// // }

// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/PuerkitoBio/goquery"
// )

// func main() {
// 	// Define the URL of the Wikipedia page
// 	url := "https://en.wikipedia.org/wiki/Pepin_of_Landen"

// 	// Define a user-agent header to simulate a browser request
// 	headers := map[string]string{
// 		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36",
// 	}

// 	// Create an HTTP client with custom headers
// 	client := &http.Client{}
// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for key, value := range headers {
// 		req.Header.Set(key, value)
// 	}

// 	// Send an HTTP GET request to the URL with the headers
// 	response, err := client.Do(req)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer response.Body.Close()

// 	// Check if the request was successful (status code 200)
// 	if response.StatusCode == 200 {
// 		// Parse the HTML content of the page using goquery
// 		doc, err := goquery.NewDocumentFromReader(response.Body)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		// Find the table with the specified class name
// 		// table := doc.Find("table.wikitable.sortable")

// 		// Initialize empty slice to store the table data
// 		// data := [][]string{}

// 		// // Iterate through the rows of the table
// 		// doc.Find("tr").Each(func(rowIdx int, row *goquery.Selection) {
// 		// 	// Skip the header row
// 		// 	if rowIdx == 0 {
// 		// 		return
// 		// 	}

// 		// 	// Extract data from each column and append it to the data slice
// 		// 	rowData := []string{}
// 		// 	row.Find("th,td").Each(func(colIdx int, col *goquery.Selection) {
// 		// 		rowData = append(rowData, col.Text())
// 		// 	})
// 		// 	data = append(data, rowData)
// 		// })

// 		// fmt.Println(data[1])

// 		data := [][]string{}

// 		// doc.Find("a href").Each(func(i int, s *goquery.Selection) {
// 		// 	// fmt.Println(s.Text())
// 		// 	// // For each item found, get the title
// 		// 	// title := s.Find("a").Text()
// 		// 	// fmt.Printf("Review %d: %s\n", i, title)
// 		// 	data = append(data, s.Text())
// 		// })
// 		// doc.Find(".left-content article .post-title").Each(func(i int, s *goquery.Selection) {
// 		// 	// For each item found, get the title
// 		// 	title := s.Find("a").Text()
// 		// 	// fmt.Printf("Review %d: %s\n", i, title)
// 		// 	data = append(data, title)
// 		// })
// 		doc.Find("tr").Each(func(rowIdx int, row *goquery.Selection) {
// 			// Skip the header row
// 			if rowIdx == 0 {
// 				return
// 			}

// 			// Extract data from each column and append it to the data slice
// 			rowData := []string{}
// 			row.Find("th,td").Each(func(colIdx int, col *goquery.Selection) {
// 				rowData = append(rowData, col.Text())
// 			})
// 			data = append(data, rowData)
// 		})

// 		fmt.Println(data[0])

// 		// // Print the scraped data for all presidents
// 		// for _, presidentData := range data {
// 		// 	fmt.Println("President Data:")
// 		// 	fmt.Println("Number:", presidentData[0])
// 		// 	fmt.Println("Name:", presidentData[2])
// 		// 	fmt.Println("Term:", presidentData[3])
// 		// 	fmt.Println("Party:", presidentData[5])
// 		// 	fmt.Println("Election:", presidentData[6])
// 		// 	fmt.Println("Vice President:", presidentData[7])
// 		// 	fmt.Println()
// 		// }
// 	} else {
// 		fmt.Println("Failed to retrieve the web page. Status code:", response.StatusCode)
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"

// 	"github.com/PuerkitoBio/goquery"
// )

// type WikipediaLink struct {
// 	Title string
// 	URL   string
// }

// func main() {

// 	// Specify the Wikipedia URL to scrape
// 	wikipediaURL := "https://en.wikipedia.org/wiki/Indonesian_language"

// 	// Scrape Wikipedia article
// 	doc, err := goquery.NewDocument(wikipediaURL)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Error scraping Wikipedia article: %v\n", err)
// 		return
// 	}

// 	// Extract hyperlinks
// 	var links []WikipediaLink
// 	title := doc.Find("title").Text()
// 	cleanedTitle := strings.TrimSuffix(title, " - Wikipedia")

// 	found := false
// 	doc.Find("#mw-content-text").Find("a").Each(func(i int, s *goquery.Selection) {
// 		if found {
// 			return
// 		}
// 		href, exists := s.Attr("href")
// 		if exists && strings.HasPrefix(href, "/wiki/") {
// 			links = append(links, WikipediaLink{
// 				Title: cleanedTitle,
// 				URL:   "https://en.wikipedia.org" + href,
// 			})
// 			found = true
// 		}
// 	})
// 	fmt.Println(links)
// }

package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type node struct {
	url    string
	path   int
	parent *node
}

// Slice of prefixes of wikipedia pages to ignore in order to only get pages of actual articles
var prefixes = [...]string{"/wiki/Main_Page", "/wiki/Special", "/wiki/File", "/wiki/Help", "/wiki/Wikipedia:"}

/* checkListOfPrefixes: Checks url against slice of prefixes to ensure the linke is to an article.*/
func checkListOfPrefixes(href string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(href, prefix) {
			return true
		}
	}
	return false
}

/*
processElement: Checks if url is valid for the race and formats it to the full url

	to prepare it for a get request.
*/
func processElement(index int, element *goquery.Selection) string {
	// See if the href attribute exists on the element
	href, exists := element.Attr("href")
	if exists && strings.HasPrefix(href, "/wiki/") && !checkListOfPrefixes(href) {
		return "https://en.m.wikipedia.org" + string(href)
	}
	return ""
}

/*
getNewLinks: Ansynchronous function that performs a get request to a url to

	ensure it is valid, creates a slice of all the valid urls from
	the page for the race, returns the slice in the provided channel.
*/
func getNewLinks(a node, out chan<- []node, path_num int) {
	//slic to add new links to
	q := make([]node, 0)

	//Get request for the url
	resp, err := http.Get(a.url)
	if err != nil {
		log.Print("Error getting page", err)
		out <- q
		return
	}

	document, queryError := goquery.NewDocumentFromReader(resp.Body)
	if queryError != nil {
		log.Fatal("Error loading HTTP response body. ", queryError)
	}

	document.Find("#mw-content-text").Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists && strings.HasPrefix(href, "/wiki/") {
			q = append(q, node{url: "https://en.wikipedia.org" + href, path: path_num, parent: &a})
		}
	})

	//send slice to the channel
	out <- q
}

func toTitleCase(s string) string {
	// Replace underscores with spaces
	s = strings.ReplaceAll(s, "_", " ")

	// Convert to title case
	words := strings.Fields(s)
	for i, word := range words {
		words[i] = strings.Title(word)
	}
	return strings.Join(words, " ")
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
			links = append(links, toTitleCase(strings.TrimPrefix(href, "/wiki/")))
		}
	})
	return links
}

// func sorthestPathBFS(start string, end string) string {

// 	// // Create a queue for BFS
// 	// queue := []string{start}
// 	// // Create a map to store visited nodes
// 	// visited := make(map[string]bool)
// 	// visited[start] = true
// 	// // Iterate through the queue
// 	// for len(queue) > 0 {
// 	// 	// Dequeue the first element
// 	// 	node := queue[0]
// 	// 	queue = queue[1:]
// 	// 	// Get the neighbours of the current node
// 	// 	neighbours := getNeighbours(node)
// 	// 	// Iterate through the neighbours
// 	// 	for _, neighbour := range neighbours {
// 	// 		// Check if the neighbour is the end node
// 	// 		if neighbour == end {
// 	// 			return node + " -> " + neighbour
// 	// 		}
// 	// 		// Check if the neighbour has been visited
// 	// 		if !visited[neighbour] {
// 	// 			// Mark the neighbour as visited
// 	// 			visited[neighbour] = true
// 	// 			// Enqueue the neighbour
// 	// 			queue = append(queue, neighbour)
// 	// 		}
// 	// 	}
// 	// }
// 	// return "No path found"

// }

/*
getPath: concatenate the path from node1 to the starting node with the path of

	node2 to the ending node.
*/
func getPath(node1 *node, node2 *node) string {
	path := ""
	for node1 != nil {
		path = node1.url + " -> " + path
		node1 = node1.parent
	}
	path = strings.TrimRight(path, " -> ")
	node2 = node2.parent
	for node2 != nil {
		path = path + " -> " + node2.url
		node2 = node2.parent
	}
	return path
}

func VisitNode(m map[string]node, q []node, threadCount int, out chan<- []node, path int, paths []string) (map[string]node, []node, int, []string) {
	if len(q) > 0 {
		var a node
		var val node
		var found bool

		//pop a url from the queue
		a, q = q[0], q[1:]
		//if the url is already in the map and it was not found in this half
		//of the search, then path has been found from start to finish
		if val, found = m[a.url]; found && val.path != path {
			paths = append(paths, getPath(&a, &val))
		}
		if !found {
			//add url to map
			m[a.url] = a
			threadCount += 1
			//asynch thread to get new links on page
			go getNewLinks(a, out, path)
		}
	}
	return m, q, threadCount, paths
}

/*
FindShortestWikiPath: Returns a string indicating the shortest path between

	two wikipeida articles. Utilizes a BFS that starts from
	both the start and finish articles.
*/
func FindShortestWikiPath(article1 string, article2 string) (string, string) {

	//convert the proviced article names to wikipedia url's
	start := node{
		url:    "https://en.m.wikipedia.org/wiki/" + strings.Replace(article1, " ", "_", -1),
		path:   1,
		parent: nil,
	}
	end := node{
		url:    "https://en.m.wikipedia.org/wiki/" + strings.Replace(article2, " ", "_", -1),
		path:   2,
		parent: nil,
	}

	//map from a url string to node representing a webpage that has been visited
	m := make(map[string]node)

	//preliminary queues for the next round of BFS
	next_q1 := make([]node, 1)
	next_q1[0] = start
	next_q2 := make([]node, 1)
	next_q2[0] = end

	//current queue of url's to be visited in the BFS
	q1 := make([]node, 0)
	q2 := make([]node, 0)

	//for each round of bfs
	for len(next_q1) > 0 || len(next_q2) > 0 {
		q1 = next_q1
		q2 = next_q2

		next_q1 = make([]node, 0)
		// next_q2 = make([]node, 0)

		//contains all valid paths from start to finish found in a round
		paths := make([]string, 0)

		//chnnels for each queue to return new links found on each page
		out1 := make(chan []node)
		// out2 := make(chan []node)

		//number of threads for each channel
		out1Count := 0
		// out2Count := 0

		//for each element in the queues in this round of BFS
		for len(q1) > 0 || len(q2) > 0 {

			m, q1, out1Count, paths = VisitNode(m, q1, out1Count, out1, 1, paths)

			// m, q2, out2Count, paths = VisitNode(m, q2, out2Count, out2, 2, paths)

		}

		//If a valid path from starting node to finishing node is found
		if len(paths) != 0 {
			//find the shortest path in the list of paths
			shortestLength := len(paths[0])
			index := 0
			for i, path := range paths {
				if length := strings.Count(path, "->"); length < shortestLength {
					shortestLength = length
					index = i
				}
			}
			return paths[index], ""
		}

		//For each thread in ech channel, get the result from the thread
		//and add it to next queue
		for i := 0; i < out1Count; i++ {
			next_q1 = append(next_q1, <-out1...)
		}
		// for i := 0; i < out2Count; i++ {
		// 	next_q2 = append(next_q2, <-out2...)
		// }

	}
	return "No Path Found.", ""

}

func FindArticleAdress(article string) (bool, string) {
	var wikiURL = "https://en.wikipedia.org/wiki/"
	wikiURL = wikiURL + strings.Replace(article, " ", "_", -1)
	resp, err := http.Get(wikiURL)
	status := resp.StatusCode
	if err != nil || status == 404 {
		return false, "The article " + article + " does not exist."
	}
	return true, ""

}

type WikipediaLink struct {
	Title string
	URL   string
}

func main() {
	// if len(os.Args) < 3 {
	// 	fmt.Println("Please provide two wikipedia article names as arguments.")
	// 	return
	// }
	// var article1 = os.Args[1]
	// var article2 = os.Args[2]

	// exists1, err := FindArticleAdress("indonesia")
	// if !exists1 {
	// 	fmt.Println(err)
	// }
	// exists2, err := FindArticleAdress("malaysia")
	// if !exists2 {
	// 	fmt.Println(err)
	// }

	// path, _ := FindShortestWikiPath("malaysia", "indonesia")
	// fmt.Println("Path: ", path)
	// func main() {

	// Specify the Wikipedia URL to scrape
	wikipediaURL := "https://en.wikipedia.org/wiki/Indonesia"

	var test = getNeighbours(wikipediaURL)
	fmt.Println(test[0])

	// // Scrape Wikipedia article
	// doc, err := goquery.NewDocument(wikipediaURL)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error scraping Wikipedia article: %v\n", err)
	// 	return
	// }

	// // Extract hyperlinks
	// var links []WikipediaLink
	// title := doc.Find("title").Text()
	// cleanedTitle := strings.TrimSuffix(title, " - Wikipedia")

	// found := false
	// doc.Find("#mw-content-text").Find("a").Each(func(i int, s *goquery.Selection) {
	// 	if found {
	// 		return
	// 	}
	// 	href, exists := s.Attr("href")
	// 	if exists && strings.HasPrefix(href, "/wiki/") {
	// 		links = append(links, WikipediaLink{
	// 			Title: cleanedTitle,
	// 			URL:   "https://en.wikipedia.org" + href,
	// 		})
	// 		found = true
	// 	}
	// })
	// fmt.Println(links)

}
