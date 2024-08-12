package handlers

import (
	"encoding/json"
	"net/http"

	"aurora-borealis/models"
	//"aurora-borealis/services"
)

// CreatePostHandler handles the creation of a markdown post
func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the incoming JSON request
	var post models.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	//// Save to local file system
	//err = services.SaveToLocal(post.Title, post.Content)
	//if err != nil {
	//	http.Error(w, "Error saving to local file", http.StatusInternalServerError)
	//	return
	//}
	//
	//// Save to S3
	//err = services.SaveToS3(post.Title, post.Content)
	//if err != nil {
	//	http.Error(w, "Error saving to S3", http.StatusInternalServerError)
	//	return
	//}
	//
	//// Create Elasticsearch document
	//err = services.SaveToElasticsearch(post)
	//if err != nil {
	//	http.Error(w, "Error saving to Elasticsearch", http.StatusInternalServerError)
	//	return
	//}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Post created successfully"))
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the incoming JSON request
	var post models.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	//// Delete from local file system
	//err = services.DeleteFromLocal(post.Title)
	//if err != nil {
	//	http.Error(w, "Error deleting from local file", http.StatusInternalServerError)
	//	return
	//}
	//
	//// Delete from S3
	//err = services.DeleteFromS3(post.Title)
	//if err != nil {
	//	http.Error(w, "Error deleting from S3", http.StatusInternalServerError)
	//	return
	//}

	// Delete from Elasticsearch
	//err = services.DeleteFromElasxticsearch(post.Title)
	if err != nil {
		http.Error(w, "Error deleting from Elasticsearch", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetPostHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the incoming JSON request
	var post models.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Get from local file system
	//err = services.GetFromLocal(post.Title)
	//if err != nil {
	//	http.Error(w, "Error getting from local file", http.StatusInternalServerError)
	//	return
	//}

	// Get from S3
	//err = services.GetFromS3(post.Title)
	//if err != nil {
	//	http.Error(w, "Error getting from S3", http.StatusInternalServerError)
	//	return
	//}

	// Get from Elasticsearch
	//err = services.GetFromElasxticsearch(post.Title)
	if err != nil {
		http.Error(w, "Error getting from Elasticsearch", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
