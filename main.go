package main

import (
	"fmt"
	"net/http"

	"github.com/yourusername/markdown-webapp/handlers"
	"github.com/yourusername/markdown-webapp/utils"
)

func main() {
	// Load configuration and initialize services
	utils.Initialize()

	// Set up routes
	http.HandleFunc("/create-post", handlers.CreatePostHandler)

	// Start the server
	fmt.Println("Starting server on :8080...")
	http.ListenAndServe(":8080", nil)
}
