package handlers

import (
	"aurora-borealis/models"
	"aurora-borealis/services"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
	//"aurora-borealis/services"
)

func extractTags(content string) ([]string, error) {
	// Extract tags from the contents with # prefix and separated by `-` and `_`. #tag-alpha_beta
	words := strings.Split(content, " ")
	tags := []string{}
	var err error = nil
	for _, word := range words {
		if strings.HasPrefix(word, "#") {
			tag := strings.Replace(word, "_", "-", -1)
			// remove all extra symbols ❗️ TODO
			tags = append(tags, tag)
		}
	}
	return tags, err
}

// CreatePostHandler handles the creation of a markdown post
func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the incoming JSON request
	var postReq models.PostOperationRequest
	err := json.NewDecoder(r.Body).Decode(&postReq)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var postContents string = postReq.Content
	var postMediaLinks []string = postReq.Media

	if len(postContents) == 0 && len(postMediaLinks) == 0 {
		http.Error(w, "Empty post", http.StatusBadRequest)
		return
	}

	tags, err := extractTags(postContents)
	if err != nil {
		http.Error(w, "Error extracting tags", http.StatusInternalServerError)
	}

	postedAt := time.Now().Format(time.RFC3339)
	/*
		generate slug
		generate short summary
		generate title
		generate sentiment
		generate category
		generate AI thoughts
		generate AI tags
		Run postReq analysis*/
	analysisMeta, err := services.RunNLPAnalysis(postContents, tags, postMediaLinks)
	if err != nil {
		http.Error(w, "Error running analysis", http.StatusInternalServerError)
	}

	// Run media analysis
	// analysisMeta, err := services.RunNLPAnalysis(postReq.Content)

	fmt.Println(analysisMeta)

	//var post models.Post = models.NewPost(postMediaLink, postContents, "user")

	//// Save to local file system
	//err = services.SaveToLocal(postReq.Title, postReq.Content)
	//if err != nil {
	//	http.Error(w, "Error saving to local file", http.StatusInternalServerError)
	//	return
	//}
	//
	//// Save to S3
	//err = services.SaveToS3(postReq.Title, postReq.Content)
	//if err != nil {
	//	http.Error(w, "Error saving to S3", http.StatusInternalServerError)
	//	return
	//}
	//
	//// Create Elasticsearch document
	//err = services.SaveToElasticsearch(postReq)
	//if err != nil {
	//	http.Error(w, "Error saving to Elasticsearch", http.StatusInternalServerError)
	//	return
	//}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Post created successfully"))
}

//func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
//	// Parse the incoming JSON request
//	var post models.Post
//	err := json.NewDecoder(r.Body).Decode(&post)
//	if err != nil {
//		http.Error(w, "Invalid request payload", http.StatusBadRequest)
//		return
//	}
//
//	//// Delete from local file system
//	//err = services.DeleteFromLocal(post.Title)
//	//if err != nil {
//	//	http.Error(w, "Error deleting from local file", http.StatusInternalServerError)
//	//	return
//	//}
//	//
//	//// Delete from S3
//	//err = services.DeleteFromS3(post.Title)
//	//if err != nil {
//	//	http.Error(w, "Error deleting from S3", http.StatusInternalServerError)
//	//	return
//	//}
//
//	// Delete from Elasticsearch
//	//err = services.DeleteFromElasxticsearch(post.Title)
//	if err != nil {
//		http.Error(w, "Error deleting from Elasticsearch", http.StatusInternalServerError)
//		return
//	}
//
//	w.WriteHeader(http.StatusNoContent)
//}
//
//func GetPostHandler(w http.ResponseWriter, r *http.Request) {
//	// Parse the incoming JSON request
//	var post models.Post
//	err := json.NewDecoder(r.Body).Decode(&post)
//	if err != nil {
//		http.Error(w, "Invalid request payload", http.StatusBadRequest)
//		return
//	}
//
//	// Get from local file system
//	//err = services.GetFromLocal(post.Title)
//	//if err != nil {
//	//	http.Error(w, "Error getting from local file", http.StatusInternalServerError)
//	//	return
//	//}
//
//	// Get from S3
//	//err = services.GetFromS3(post.Title)
//	//if err != nil {
//	//	http.Error(w, "Error getting from S3", http.StatusInternalServerError)
//	//	return
//	//}
//
//	// Get from Elasticsearch
//	//err = services.GetFromElasxticsearch(post.Title)
//	if err != nil {
//		http.Error(w, "Error getting from Elasticsearch", http.StatusInternalServerError)
//		return
//	}
//
//	w.WriteHeader(http.StatusOK)
//}
