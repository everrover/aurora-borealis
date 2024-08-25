package services

import (
	"aurora-borealis/models"
	utils "aurora-borealis/utils"
	"bytes"
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// SaveToElasticsearch indexes the post in Elasticsearch
func SaveToElasticsearch(post models.Post) error {
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

var esClient *elasticsearch.Client

func InitElasticsearch() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	esClient = client
	createInitialIndex("")
}

func GetElasticsearchClient() *elasticsearch.Client { return esClient }

func createNewIndexIfNotExists(indexName string, indexMapping string) {

	exists, err := esClient.Indices.Exists([]string{indexName})
	if err != nil || exists.StatusCode != 200 {
		log.Println("Index already exists")
		return
	}

	mappingFile, err := os.Open(indexMapping)
	if err != nil {
		log.Fatalf("Error opening mapping file: %s", err)
	}
	defer mappingFile.Close()

	mappingBytes, err := ioutil.ReadAll(mappingFile)
	if err != nil {
		log.Fatalf("Error reading mapping file: %s", err)
	}

	// Create the index with the specified mapping
	req := esapi.IndicesCreateRequest{
		Index: indexName,
		Body:  bytes.NewReader(mappingBytes),
	}
	res, err := req.Do(context.Background(), esClient)
	if err != nil || res.IsError() {
		log.Fatalf("Error creating index/Error response from Elasticsearch: %s / %s", err, res.String())
	}
	defer res.Body.Close()
}

func mapAliasToIndex(aliasName, indexName string) {
	// Create an alias for the index
	aliasReq := esapi.IndicesPutAliasRequest{
		Body: bytes.NewReader([]byte(fmt.Sprintf(`{
			"actions": [
				{ "remove": { "alias": "%s", "index": "*" } },
				{ "add": { "alias": "%s", "index": "%s" } }
			]
		}`, aliasName, aliasName, indexName))),
	}
	aliasRes, err := aliasReq.Do(context.Background(), esClient)
	if err != nil || aliasRes.IsError() {
		log.Fatalf("Error creating alias/Error response from Elasticsearch: %s / %s", err, aliasRes.String())
	}
	defer aliasRes.Body.Close()
}

func createInitialIndex(indexName string) {
	if indexName == "" {
		indexName = "posts_idx_v0001"
	}

	createNewIndexIfNotExists(indexName, "utils/es-mapping.json")
	mapAliasToIndex("post_index", indexName)

	fmt.Println("Index and alias created successfully")
}

func CreateNewIndexWithMapping(indexName string, mappingFile string) {
	if indexName == "" || mappingFile == "" {
		log.Fatalf("Index name and mapping are required")
	}

	createNewIndexIfNotExists(indexName, mappingFile)
	mapAliasToIndex("post_index", indexName)

	fmt.Println("Index and alias created and remapped successfully")
}
