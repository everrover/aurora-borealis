package main

import (
	"aurora-borealis/handlers"
	"aurora-borealis/utils"
	"fmt"
	gmux "github.com/gorilla/mux"
	"net/http"
)

func main() {
	// Load configuration and initialize services
	utils.Initialize()

	// Set up routes
	//mux := http.NewServeMux()
	router := gmux.NewRouter()
	router.Methods("POST").Path("/create-post").HandlerFunc(handlers.CreatePostHandler)
	//mux.HandleFunc("POST /create-post", handlers.CreatePostHandler)
	//http.HandleFunc("/delete-post", handlers.CreatePostHandler)
	//http.HandleFunc("/update-post", handlers.CreatePostHandler)
	//http.HandleFunc("/get-post", handlers.CreatePostHandler)

	// Start the server
	fmt.Println("Starting server on :8080...")
	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		return
	}
}
