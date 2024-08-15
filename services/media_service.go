package services

import (
	"aurora-borealis/models"
	utils "aurora-borealis/utils"
	"context"
	"time"
)

// SaveToElasticsearch indexes the post in Elasticsearch
func RunMediaAnalysis(post models.Post) error {
	doc := map[string]interface{}{
		//"title":   post.Title,
		"content": post.Content,
		"date":    time.Now().Format(time.RFC3339),
	}

	_, err := utils.ESClient.Index().
		Index(utils.ESIndexName).
		BodyJson(doc).
		Do(context.Background())
	return err
}
