package main

import (
	"aurora-borealis/handlers"
	"aurora-borealis/utils"
	"fmt"
	"net/http"
)

func main() {
	// Load configuration and initialize services
	utils.Initialize()

	// Set up routes
	http.HandleFunc("/create-post", handlers.CreatePostHandler)
	//http.HandleFunc("/delete-post", handlers.CreatePostHandler)
	//http.HandleFunc("/update-post", handlers.CreatePostHandler)
	http.HandleFunc("/get-post", handlers.CreatePostHandler)

	// Start the server
	fmt.Println("Starting server on :8080...")
	http.ListenAndServe(":8080", nil)
}
