package services

import (
	"aurora-borealis/models"
	utils "aurora-borealis/utils"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"io/ioutil"
	"log"
	"os"
)

// SaveToElasticsearch indexes the post in Elasticsearch
func SaveToElasticsearch(post models.Post) (string, error) {
	// Marshal the post into JSON
	postJSON, err := json.Marshal(post)
	if err != nil {
		return "", err
	}

	// Create the request
	req := esapi.IndexRequest{
		Index:      utils.ES_POST_INDEX_NAME,
		DocumentID: post.Slug,
		Body:       bytes.NewReader(postJSON),
		Refresh:    "true",
	}

	// Perform the request with the client
	res, err := req.Do(context.Background(), esClient)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.IsError() {
		return "", fmt.Errorf("Error indexing document: %s", res.String())
	}

	return post.Slug, nil
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
	createInitialIndex()
}

func GetElasticsearchClient() *elasticsearch.Client { return esClient }

func createNewIndexIfNotExists(indexName string, indexMapping string) {

	exists, err := esClient.Indices.Exists([]string{indexName})
	if err != nil || exists.StatusCode == 200 {
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
	// Check if the alias already exists
	alias, err := esClient.Indices.GetAlias(
		esClient.Indices.GetAlias.WithContext(context.Background()),
		esClient.Indices.GetAlias.WithName(aliasName),
	)
	if err != nil || alias.IsError() {
		log.Println("Some bullshit happened")
	}

	var putAliasReq esapi.IndicesPutAliasRequest
	if alias.StatusCode != 200 {
		log.Println("Alias does not exist")
		// Create an alias for the index

		// Add the alias to the index
		putAliasReq = esapi.IndicesPutAliasRequest{
			Body: bytes.NewReader([]byte(fmt.Sprintf(`{
				"actions": [
					{ "add": { "alias": "%s", "index": "%s" } }
				]
			}`, aliasName, indexName))),
		}
	} else {
		var aliasResponse map[string]interface{}
		if err := json.NewDecoder(alias.Body).Decode(&aliasResponse); err != nil {
			log.Fatalf("Error parsing response body: %s", err)
		}

		var currentIndex string
		for index := range aliasResponse {
			currentIndex = index
			break
		}

		if currentIndex == indexName {
			log.Println("Alias already mapped to the index")
			return
		}

		putAliasReq = esapi.IndicesPutAliasRequest{
			Body: bytes.NewReader([]byte(fmt.Sprintf(`{
				"actions": [
					{ "remove": { "alias": "%s", "index": "%s" } },
					{ "add": { "alias": "%s", "index": "%s" } }
				]
			}`, aliasName, currentIndex, aliasName, indexName))),
		}

		fmt.Println("Req is: ", putAliasReq.Body)
	}
	aliasRes, err := putAliasReq.Do(context.Background(), esClient)
	if err != nil || aliasRes.IsError() {
		log.Fatalf("Error creating alias/Error response from Elasticsearch: %s / %s", err, aliasRes.String())
	}
	defer aliasRes.Body.Close()
}

func createInitialIndex() {
	createNewIndexIfNotExists(utils.ES_POST_INDEX_NAME, utils.ES_POST_INDEX_MAPPING)
	//mapAliasToIndex(utils.ES_POST_INDEX_ALIAS, utils.ES_POST_INDEX_NAME)

	createNewIndexIfNotExists(utils.ES_POST_MEDIA_INDEX_NAME, utils.ES_POST_MEDIA_INDEX_MAPPING)
	//mapAliasToIndex(utils.ES_POST_MEDIA_INDEX_ALIAS, utils.ES_POST_MEDIA_INDEX_NAME)

	fmt.Println("Index and alias created successfully")
}

func CreateNewIndexWithMapping(indexName string, mappingFile string, aliasName string) {
	if indexName == "" || mappingFile == "" {
		log.Fatalf("Index name and mapping are required")
	}

	createNewIndexIfNotExists(indexName, mappingFile)
	mapAliasToIndex(aliasName, indexName)

	fmt.Println("Index and alias created and remapped successfully")
}
