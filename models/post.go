package models

type CreateRequestPost struct {
	Media   string `json:"media"`   // ip - user - video, audio, image, and others path
	Content string `json:"content"` // ip - user
}

type UpdateRequestPost struct {
	Media   string `json:"media"`   // ip - user - video, audio, image, and others path
	Content string `json:"content"` // ip - user
	ID      string `json:"id"`      // ip - user
}

type DeleteRequestPost struct {
	ID string `json:"id"` // ip - user
}

// Post represents a markdown post
type Post struct {
	Media      string            `json:"media"`     // ip - user - video, audio, image, and others path
	Content    string            `json:"content"`   // ip - user
	SelfLike   bool              `json:"self_like"` // ip - user
	ID         string            `json:"id"`        // auto generated
	PostedAt   string            `json:"posted_at"` // ip - system
	Author     string            `json:"author"`    // ip - system
	Slug       string            `json:"slug"`      // ip - system
	Metadata   string            `json:"metadata"`  // auto-gen via AI analysis
	Comments   []Comment         `json:"comments"`
	Likes      int               `json:"likes"`
	WhatAISays map[string]string `json:"what_ai_says"`
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
