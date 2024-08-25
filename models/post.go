package models

// Post represents a markdown post
type Post struct {
	Media       []string  `json:"media"`         // ip - user - video, audio, image, and others path
	Content     string    `json:"content"`       // ip - user
	SelfLike    bool      `json:"self_like"`     // ip - user
	Tags        []string  `json:"tags"`          // prefixed with #
	AutoGenTags []string  `json:"auto_gen_tags"` // prefixed with #
	ID          string    `json:"id"`            // auto generated
	PostedAt    string    `json:"posted_at"`     // ip - system
	Author      string    `json:"author"`        // ip - system
	Slug        string    `json:"slug"`          // ip - system
	Metadata    string    `json:"metadata"`      // auto-gen via AI analysis
	Comments    []Comment `json:"comments"`      // CommentIDs
	Likes       int       `json:"likes"`         // Count of likes
	WhatAISays  string    `json:"what_ai_says"`

	ShortSummary string   `json:"short_summary"`
	Title        string   `json:"title"`
	Sentiment    string   `json:"sentiment"`
	Category     string   `json:"category"`
	IdealFor     []string `json:"ideal_for"`
}

// NewPost creates a new post - with all fields
func NewPost(media []string, content string, selfLike bool, tags []string,
	autoGenTags []string, id string, postedAt string, author string, slug string,
	metadata string, likes int, whatAISays string, shortSummary string,
	title string, sentiment string, category string, idealFor []string) *Post {
	return &Post{
		Media:        media,
		Content:      content,
		SelfLike:     selfLike,
		Tags:         tags,
		AutoGenTags:  autoGenTags,
		ID:           id,
		PostedAt:     postedAt,
		Author:       author,
		Slug:         slug,
		Metadata:     metadata,
		Comments:     []Comment{},
		Likes:        likes,
		WhatAISays:   whatAISays,
		ShortSummary: shortSummary,
		Title:        title,
		Sentiment:    sentiment,
		Category:     category,
		IdealFor:     idealFor,
	}
}

// Comment represents a comment on a post
type Comment struct {
	Content  string `json:"content"`
	PostedAt string `json:"posted_at"`
	Author   string `json:"author"`
	Likes    int    `json:"likes"`
	SelfLike bool   `json:"self_like"`
}

type Like struct {
	PostID string `json:"post_id"`
	UserID string `json:"user_id"`
}
