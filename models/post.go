package models

// Post represents a markdown post
type Post struct {
	Media      string            `json:"media"`     // ip - user - video, audio, image, and others path
	Content    string            `json:"content"`   // ip - user
	SelfLike   bool              `json:"self_like"` // ip - user
	Tags       []string          `json:"tags"`      // prefixed with #
	ID         string            `json:"id"`        // auto generated
	PostedAt   string            `json:"posted_at"` // ip - system
	Author     string            `json:"author"`    // ip - system
	Slug       string            `json:"slug"`      // ip - system
	Metadata   string            `json:"metadata"`  // auto-gen via AI analysis
	Comments   []Comment         `json:"comments"`  // CommentIDs
	Likes      int               `json:"likes"`     // Count of likes
	WhatAISays map[string]string `json:"what_ai_says"`
}

func NewPost(slug string, media string, content string, author string, hashtags []string, postedAt string) *Post {
	return &Post{
		Media:      media,
		Content:    content,
		SelfLike:   false,
		Tags:       hashtags,
		ID:         "",
		PostedAt:   postedAt,
		Author:     author,
		Slug:       slug,
		Metadata:   "",
		Comments:   []Comment{},
		Likes:      0,
		WhatAISays: map[string]string{},
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
