package models

type PostOperationRequest struct {
	Media   []string `json:"media"`   // ip - user - video, audio, image, and others path - Update/Create
	Content string   `json:"content"` // ip - user - Update/Create
	ID      string   `json:"id"`      // ip - user - Update/Delete/Get
	User    string   `json:"user"`    // ip - user - Update/Create/Delete/Get/GetAll
	Page    int      `json:"page"`    // ip - user - Get/GetAll
}

// Constructor for PostOperationRequest
func NewPostOperationRequest(media []string, content string, id string, user string, page int) *PostOperationRequest {
	return &PostOperationRequest{
		Media:   media,
		Content: content,
		ID:      id,
		User:    user,
		Page:    page,
	}
}
