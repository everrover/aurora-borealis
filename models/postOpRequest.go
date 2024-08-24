package models

type PostOperationRequest struct {
	Media    []string `json:"media"`   // ip - user - video, audio, image, and others path - Update/Create
	Content  string   `json:"content"` // ip - user - Update/Create
	ID       string   `json:"id"`      // ip - user - Update/Delete/Get
	User     string   `json:"user"`    // ip - user - Update/Create/Delete/Get/GetAll
	Page     int      `json:"page"`    // ip - user - Get/GetAll
	Weather  string   `json:"weather"`
	Location string   `json:"location"`
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

type PostRequestForAnalysis struct {
	Content  string   `json:"post_contents"`
	Media    []string `json:"media_description"`
	Author   string   `json:"author"`
	Location string   `json:"location"`
	Weather  string   `json:"weather"`
	Tags     []string `json:"hashtags"`
	PostedAt string   `json:"date"`
}

func NewPostRequestForAnalysis(
	content string, media []string, author string,
	location string, weather string, tags []string,
	postedAt string) *PostRequestForAnalysis {
	return &PostRequestForAnalysis{
		Content:  content,
		Media:    media,
		Author:   author,
		Location: location,
		Weather:  weather,
		Tags:     tags,
		PostedAt: postedAt,
	}
}
