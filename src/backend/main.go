package main

import (
	"example/user/hello/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/urfave/negroni"
)

func main() {
	port := "8080"

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// Setup for Router, Endpoints, Handlers and middleware
	router := routes.GetAllRoutes()
	middleWare := negroni.Classic()
	middleWare.UseHandler(router)

	// Serves API - Creates a new thread and if fails it will log the error.
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":"+port, middleWare))
}
