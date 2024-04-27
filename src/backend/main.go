package main

import (
	"example/user/hello/routes"
	"log"
	"time"
)

func main() {
	startTime := time.Now()
	routes.ServeRoutes().Run(":8080")
	elapsed := time.Since(startTime)
	log.Printf("Time Execution: %s", elapsed)
}
