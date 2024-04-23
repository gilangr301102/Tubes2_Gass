package routes

import (
	"example/user/hello/controllers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// func GetScrappingRoutes(router *mux.Router) *mux.Router {
// 	router.Handle(
// 		"/scrap/{srapName}",
// 		negroni.New(
// 			negroni.HandlerFunc(controllers.GetScrappingData2),
// 		)).Methods("GET")

// 	return router
// }

func GetScrappingRoutes(w http.ResponseWriter, r *http.Request) {
	// Get the search name parameter from the URL
	params := mux.Vars(r)
	searchName := params["searchName"]

	// Call the GetScrappingData function from the controller
	data := controllers.GetScrappingData(searchName)

	// Send the data as JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%v", data)
}
