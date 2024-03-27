package main

import (
	"fmt"
	"log"
	"os"
	"test-html/router"
)

func main() {
	fmt.Println("Started Running")
	r := router.Router()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8089" // Default to 8080 if PORT is not set
	}
	log.Fatal(r.Run("0.0.0.0:" + port))
}
