package utils

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

func wikiraceBFS(c *gin.Context) {
	sourceTitle := c.PostForm("sourceTitle")
	goalTitle := c.PostForm("goalTitle")
	isFindAll := c.PostForm("isFindAll")

	sourceUrl := URL_SCRAPPING_WIKIPEDIA + handleUnderScore(sourceTitle)
	goalUrl := URL_SCRAPPING_WIKIPEDIA + handleUnderScore(goalTitle)
	if len(sourceUrl) != 0 && len(goalUrl) != 0 {
		if isFindAll == "0" {
			startTime := time.Now()
			path, numOfArticlesChecked, articlesLoop := GetShortestSinglePathBFS(sourceUrl, goalUrl)
			elapsedTime := time.Since(startTime).Seconds()
			c.JSON(http.StatusOK, gin.H{
				"path":                        path,
				"message":                     "Success",
				"num_of_article_checked":      numOfArticlesChecked,
				"num_of_node_article_visited": articlesLoop,
				"elapsed_time":                elapsedTime,
			})
			log.Printf("Path: %v\n", path)
			fmt.Println(path)
		} else {
			startTime := time.Now()
			path, numberPath, numOfArticlesChecked, articlesLoop := GetShortestMultiPathBFS(sourceUrl, goalUrl)
			elapsedTime := time.Since(startTime).Seconds()
			c.JSON(http.StatusOK, gin.H{
				"path":                        path,
				"message":                     "Success",
				"num_of_article_checked":      numOfArticlesChecked,
				"num_of_node_article_visited": articlesLoop,
				"number_of_path":              numberPath,
				"elapsed_time":                elapsedTime,
			})
			log.Printf("Path: %v\n", path)
			fmt.Println(path)
		}
	} else {
		c.JSON(http.StatusBadRequest, nil)
	}
}

func wikiraceIDS(c *gin.Context) {
	sourceTitle := c.PostForm("sourceTitle")
	goalTitle := c.PostForm("goalTitle")
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

	startNode := Node{Title: sourceTitle, URL: URL_SCRAPPING_WIKIPEDIA + handleUnderScore(sourceTitle)}
	targetNode := Node{Title: goalTitle, URL: URL_SCRAPPING_WIKIPEDIA + handleUnderScore(goalTitle)}

	fmt.Printf("Start URL: %s, End URL: %s, Max Depth: %d\n", startNode.URL, targetNode.URL, maxDepthNum)

	if len(sourceTitle) != 0 && len(goalTitle) != 0 {
		startTime := time.Now()
		numOfArticlesChecked, articlesLoop, numberPath, paths := getShortestPathIDS(startNode, targetNode, maxDepthNum, isFindAllReq)
		elapsedTime := time.Since(startTime).Seconds()
		c.JSON(http.StatusOK, gin.H{
			"path":                        paths,
			"message":                     "Success",
			"num_of_article_checked":      numOfArticlesChecked,
			"num_of_node_article_visited": articlesLoop,
			"number_of_path":              numberPath,
			"elapsed_time":                elapsedTime,
		})
	} else {
		c.JSON(http.StatusBadRequest, nil)
	}
}

func ServeRoutes() *gin.Engine {

	log.Printf("Initializing to Listening and Serving Server...")

	router := gin.Default()

	// Use the Cors middleware
	router.Use(Cors())

	router.POST("/wikiraceBFS", wikiraceBFS)
	router.POST("/wikiraceIDS", wikiraceIDS)

	return router
}
