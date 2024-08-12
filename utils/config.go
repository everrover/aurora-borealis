package utils

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/olivere/elastic/v7"
)

const (
	S3BucketName     = "your-s3-bucket-name"
	ESIndexName      = "posts"
	LocalDir         = "./posts" // TB imported at run configs
	ElasticSearchURL = "http://localhost:9200"
)

var (
	S3Client *s3.S3
	ESClient *elastic.Client
)

// Initialize initializes the AWS session, S3 client, and Elasticsearch client
func Initialize() {
	// Initialize AWS session
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))
	S3Client = s3.New(sess)

	// Initialize Elasticsearch client
	var err error
	ESClient, err = elastic.NewClient(elastic.SetURL(ElasticSearchURL), elastic.SetSniff(false))
	if err != nil {
		log.Fatalf("Error initializing Elasticsearch client: %v", err)
	}
}
