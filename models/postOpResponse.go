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
