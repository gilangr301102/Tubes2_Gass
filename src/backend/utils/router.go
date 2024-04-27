package utils

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func wikiraceBFS(c *gin.Context) {
	sourceUrl := c.PostForm("sourceUrl")
	goalUrl := c.PostForm("goalUrl")

	var path *[]string = nil
	if len(sourceUrl) != 0 && len(goalUrl) != 0 {
		path, _ = GetShortestPathBFS(sourceUrl, goalUrl)
		c.JSON(http.StatusOK, gin.H{
			"path": path,
		})
	} else {
		c.JSON(http.StatusBadRequest, nil)
	}

	log.Printf(
		"Request server with Start URL: %s, End URL: %s, Path: %s\n",
		sourceUrl, goalUrl, path)
}

func wikiraceIDS(c *gin.Context) {
	sourceUrl := c.PostForm("sourceUrl")
	goalUrl := c.PostForm("goalUrl")
	maxDepth := c.PostForm("maxDepth")
	isFindAll := c.PostForm("isFindAll")
	isFindAllReq := false
	if isFindAll == "1" {
		isFindAllReq = true
	}
	maxDepthNum, err := strconv.Atoi(maxDepth)
	if err != nil {
		maxDepthNum = 0
	}

	startNode := Node{Title: sourceUrl, URL: URL_SCRAPPING_WIKIPEDIA + handleUnderScore(sourceUrl)}
	targetNode := Node{Title: goalUrl, URL: URL_SCRAPPING_WIKIPEDIA + handleUnderScore(goalUrl)}

	startTime := time.Now()
	numOfArticlesChecked, articlesLoop, numberPath, paths := getShortestPathIDS(startNode, targetNode, maxDepthNum, isFindAllReq)
	elapsedTime := time.Since(startTime).Seconds()

	var path *[]string = nil
	var message string = ""
	if len(sourceUrl) != 0 && len(goalUrl) != 0 {
		c.JSON(http.StatusOK, gin.H{
			"path":                        paths,
			"message":                     message,
			"num_of_article_checked":      numOfArticlesChecked,
			"num_of_node_article_visited": articlesLoop,
			"number_of_path":              numberPath,
			"elapsed_time":                elapsedTime,
		})
	} else {
		c.JSON(http.StatusBadRequest, nil)
	}

	log.Printf(
		"Request server with Start URL: %s, End URL: %s,  Max Depth: %d, Path: %s\n",
		sourceUrl, goalUrl, maxDepthNum, path)
}

func ServeRoutes() *gin.Engine {

	log.Printf("Initializing to Listening and Serving Server...")

	router := gin.Default()

	router.POST("/wikiraceBFS", wikiraceBFS)
	router.POST("/wikiraceIDS", wikiraceIDS)

	return router
}
