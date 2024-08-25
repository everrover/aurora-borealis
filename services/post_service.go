package services

import (
	"aurora-borealis/models"
	"fmt"
	//"aurora-borealis/utils"
)

func CreatePostFileContents(post models.Post) string {
	mediaString := ""
	for _, media := range post.Media {
		mediaString += media + ","
	}
	tagsString := ""
	for _, tag := range post.Tags {
		tagsString += tag + ","
	}

	content := fmt.Sprintf(`
---
media: %s
self_like: %t
tags: %s
id: %s
posted_at: %s
author: %s
slug: %s
metadata: %s	
likes: %d
what_ai_says: %s

short_summary: %s
title: %s
sentiment: %s
category: %s
ideal_for: %s
---
%s
`, mediaString, post.SelfLike, tagsString, post.ID, post.PostedAt,
		post.Author, post.Slug, post.Metadata, post.Likes, post.WhatAISays,
		post.ShortSummary, post.Title, post.Sentiment, post.Category, post.IdealFor, post.Content)
	return content
}
