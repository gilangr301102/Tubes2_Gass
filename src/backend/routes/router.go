package routes

import (
	"example/user/hello/utils"
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

func ServeRoutes() *gin.Engine {

	router := gin.Default()

	log.Printf("Listening and Serving Server...")

	router.POST("/wikirace", wikiraceBFS)

	return router
}
