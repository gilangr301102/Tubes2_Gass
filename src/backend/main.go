package main

import (
	"example/user/hello/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// port := os.Getenv("PORT")
	// fmt.Println("PORT: ")
	// fmt.Println(port)
	// if port == "" {
	// 	log.Fatal("$PORT must be set")
	// }

	// Create a new router
	router := mux.NewRouter()

	// Define the route with a searchName parameter
	router.HandleFunc("/api/scrap/{searchName}", routes.GetScrappingRoutes).Methods("GET")

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
