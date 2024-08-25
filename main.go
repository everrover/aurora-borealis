package main

import (
	"aurora-borealis/handlers"
	"aurora-borealis/services"
	"aurora-borealis/utils"
	"fmt"
	gmux "github.com/gorilla/mux"
	"net/http"
)

func main() {
	// Load configuration and initialize services
	utils.Initialize()
	services.InitElasticsearch()

	// Set up routes
	//mux := http.NewServeMux()
	router := gmux.NewRouter()
	router.Methods("POST").Path("/post").HandlerFunc(handlers.CreatePostHandler)
	router.Methods("GET").Path("/post").HandlerFunc(handlers.CreatePostHandler)
	router.Methods("GET").Path("/posts").HandlerFunc(handlers.CreatePostHandler)
	router.Methods("PUT").Path("/post").HandlerFunc(handlers.CreatePostHandler)
	router.Methods("DELETE").Path("/post").HandlerFunc(handlers.CreatePostHandler)

	router.Methods("POST").Path("/media").HandlerFunc(handlers.CreatePostHandler)
	router.Methods("DELETE").Path("/media").HandlerFunc(handlers.CreatePostHandler)

	router.Methods("POST").Path("/habit").HandlerFunc(handlers.CreatePostHandler)
	router.Methods("GET").Path("/habit").HandlerFunc(handlers.CreatePostHandler)
	router.Methods("GET").Path("/habits").HandlerFunc(handlers.CreatePostHandler)
	router.Methods("PUT").Path("/habit").HandlerFunc(handlers.CreatePostHandler)
	router.Methods("DELETE").Path("/habit").HandlerFunc(handlers.CreatePostHandler)

	router.Methods("POST").Path("/habit/mark").HandlerFunc(handlers.CreatePostHandler)
	router.Methods("GET").Path("/search").HandlerFunc(handlers.CreatePostHandler)

	// Start the server
	fmt.Println("Starting server on :8080...")
	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		return
	}
}
