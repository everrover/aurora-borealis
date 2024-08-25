package handlers

import (
	"aurora-borealis/models"
	"aurora-borealis/services"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
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
			tag = strings.Replace(tag, ".", "", -1)
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
	var author string = postReq.User
	var location string = postReq.Location
	var weather string = postReq.Weather
	postedAt := time.Now().Format(time.RFC3339)

	if len(postMediaLinks) > 10 {
		http.Error(w, "Media links should be less than 10 and content should not be empty", http.StatusBadRequest)
		return
	}

	if len(postContents) == 0 && len(postMediaLinks) == 0 {
		http.Error(w, "Empty post", http.StatusBadRequest)
		return
	}

	hashtags, err := extractTags(postContents)
	if err != nil {
		http.Error(w, "Error extracting tags", http.StatusInternalServerError)
	}

	var postForAnalysis *models.PostRequestForAnalysis = models.NewPostRequestForAnalysis(
		postContents, postMediaLinks, author,
		location, weather, hashtags,
		string(postedAt))
	// call api to get media analysis
	//analysisMeta, err := services.RunMediaAnalysis(postForAnalysis)

	analysisMeta, err := services.RunNLPAnalysis(*postForAnalysis)
	if analysisMeta != nil && (len(analysisMeta.Hashtags) > 0) {
		hashtags = append(hashtags, analysisMeta.Hashtags...)
	}
	if err != nil {
		http.Error(w, "Error running analysis", http.StatusInternalServerError)
	}

	fmt.Println(analysisMeta)

	// Create post entity
	var post *models.Post = models.NewPost(postReq.ID, postMediaLinks, postContents, author, hashtags, postedAt)

	// Add post to ES

	// Create post file and upload it
	services.SaveFileContents(post.Slug, *post, 1)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Post created successfully"))
}
