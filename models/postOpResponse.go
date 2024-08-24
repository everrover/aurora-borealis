package models

type PostOperationResponse struct {
	Response      string `json:"response"` // just a message
	ID            string `json:"id"`
	ResponsePost  Post   `json:"post"`
	ResponsePosts []Post `json:"posts"`
}

// Constructor for PostOperationRequest
func NewPostOperationResponse(media string, post Post, posts []Post, id string) *PostOperationResponse {
	return &PostOperationResponse{
		Response:      media,
		ID:            id,
		ResponsePost:  post,
		ResponsePosts: posts,
	}
}

type PostAnalysisResponse struct {
	ShortSummary string   `json:"short_summary"`
	Title        string   `json:"title"`
	Sentiment    string   `json:"sentiment"`
	Category     string   `json:"category"`
	Hashtags     []string `json:"hashtags"`
	WhatWeSay    string   `json:"what-we-say"`
	IdealFor     []string `json:"ideal_for"`
	Slug         string   `json:"slug"`
}

type PostAnalysisResponseActual struct {
	Message  string               `json:"message"`
	Response PostAnalysisResponse `json:"response"`
}

// Constructor for PostAnalysisResponse not needed - unmarshalled from JSON
