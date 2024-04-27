package main

import (
	"backend/wikirace/utils"
	"log"
	"time"
)

func main() {
	startTime := time.Now()
	utils.ServeRoutes().Run(":8080")
	elapsed := time.Since(startTime)
	log.Printf("Time Execution: %s", elapsed)
}
