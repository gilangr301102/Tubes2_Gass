package routes

import (
	"example/user/hello/controllers"
	"log"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func GetScrappingRoute(router *mux.Router) *mux.Router {

	router.Handle(
		"/api/scrap/{searchName}",
		negroni.New(
			negroni.HandlerFunc(controllers.GetScrappingData),
		)).Methods("GET")

	return router
}

func GetAllRoutes() *mux.Router {
	router := mux.NewRouter()

	//Set Router Routes
	router = GetScrappingRoute(router)

	log.Printf("Listening All Routes...")

	return router
}
