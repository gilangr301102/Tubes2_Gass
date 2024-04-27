package routes

import (
	"backend/wikirace/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func wikiraceBFS(c *gin.Context) {
	sourceUrl := c.PostForm("sourceUrl")
	goalUrl := c.PostForm("goalUrl")

	var path *[]string = nil
	if len(sourceUrl) != 0 && len(goalUrl) != 0 {
		path, _ = utils.GetShortestPathBFS(sourceUrl, goalUrl)
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
	// maxDepth := c.PostForm("maxDepth")
	// maxDepthNum, err := strconv.Atoi(maxDepth)
	// if err != nil {
	// 	maxDepthNum = 0
	// }

	var path *[]string = nil
	var message string = ""
	if len(sourceUrl) != 0 && len(goalUrl) != 0 {
		path, _, message = utils.GetShortestPathIDDFS(sourceUrl, goalUrl, 1)
		c.JSON(http.StatusOK, gin.H{
			"path":    path,
			"message": message,
		})
	} else {
		c.JSON(http.StatusBadRequest, nil)
	}

	log.Printf(
		"Request server with Start URL: %s, End URL: %s,  Max Depth: %d, Path: %s\n",
		sourceUrl, goalUrl, 1, path)
}

func ServeRoutes() *gin.Engine {

	router := gin.Default()

	log.Printf("Listening and Serving Server...")

	// router.POST("/wikiraceBFS", wikiraceBFS)
	router.POST("/wikiraceIDS", wikiraceIDS)

	return router
}
