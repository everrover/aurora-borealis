package main

//
//import (
//	"bytes"
//	"context"
//	"encoding/json"
//	"fmt"
//	"io/ioutil"
//	"net/http"
//	"os"
//	"time"
//
//	"github.com/aws/aws-sdk-go/aws"
//	"github.com/aws/aws-sdk-go/aws/session"
//	"github.com/aws/aws-sdk-go/service/s3"
//	"github.com/olivere/elastic/v7"
//)
//
//const (
//	s3BucketName = "your-s3-bucket-name"
//	esIndexName  = "posts"
//	localDir     = "./posts"
//)
//
//var (
//	s3Client         *s3.S3
//	esClient         *elastic.Client
//	awsSession       *session.Session
//	elasticsearchURL = "http://localhost:9200"
//)
//
//type Post struct {
//	Title   string `json:"title"`
//	Content string `json:"content"`
//}
//
//func init() {
//	// Initialize AWS session
//	awsSession = session.Must(session.NewSession(&aws.Config{
//		Region: aws.String("us-east-1"),
//	}))
//	s3Client = s3.New(awsSession)
//
//	// Initialize Elasticsearch client
//	var err error
//	esClient, err = elastic.NewClient(elastic.SetURL(elasticsearchURL), elastic.SetSniff(false))
//	if err != nil {
//		panic(err)
//	}
//}
//
//func main() {
//	http.HandleFunc("/create-post", createPostHandler)
//	fmt.Println("Starting server on :8080...")
//	http.ListenAndServe(":8080", nil)
//}
//
//func createPostHandler(w http.ResponseWriter, r *http.Request) {
//	// Parse the incoming JSON request
//	var post Post
//	err := json.NewDecoder(r.Body).Decode(&post)
//	if err != nil {
//		http.Error(w, "Invalid request payload", http.StatusBadRequest)
//		return
//	}
//
//	// Create the markdown content
//	markdownContent := fmt.Sprintf("# %s\n\n%s", post.Title, post.Content)
//
//	// Save to local file system
//	err = saveToLocal(post.Title, markdownContent)
//	if err != nil {
//		http.Error(w, "Error saving to local file", http.StatusInternalServerError)
//		return
//	}
//
//	// Save to S3
//	err = saveToS3(post.Title, markdownContent)
//	if err != nil {
//		http.Error(w, "Error saving to S3", http.StatusInternalServerError)
//		return
//	}
//
//	// Create Elasticsearch document
//	err = saveToElasticsearch(post)
//	if err != nil {
//		http.Error(w, "Error saving to Elasticsearch", http.StatusInternalServerError)
//		return
//	}
//
//	w.WriteHeader(http.StatusCreated)
//	w.Write([]byte("Post created successfully"))
//}
//
//func saveToLocal(title, content string) error {
//	// Ensure the local directory exists
//	if _, err := os.Stat(localDir); os.IsNotExist(err) {
//		err := os.MkdirAll(localDir, os.ModePerm)
//		if err != nil {
//			return err
//		}
//	}
//
//	// Create the markdown file
//	filename := fmt.Sprintf("%s/%s.md", localDir, title)
//	err := ioutil.WriteFile(filename, []byte(content), 0644)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func saveToS3(title, content string) error {
//	// Upload the markdown content to S3
//	_, err := s3Client.PutObject(&s3.PutObjectInput{
//		Bucket: aws.String(s3BucketName),
//		Key:    aws.String(fmt.Sprintf("%s.md", title)),
//		Body:   bytes.NewReader([]byte(content)),
//		ACL:    aws.String("private"),
//	})
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func saveToElasticsearch(post Post) error {
//	// Prepare the document for Elasticsearch
//	doc := map[string]interface{}{
//		"title":   post.Title,
//		"content": post.Content,
//		"date":    time.Now().Format(time.RFC3339),
//	}
//
//	// Index the document
//	_, err := esClient.Index().
//		Index(esIndexName).
//		BodyJson(doc).
//		Do(context.Background())
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
