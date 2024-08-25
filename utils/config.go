package utils

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/olivere/elastic/v7"
)

const (
	S3BucketName                = "your-s3-bucket-name"
	ESIndexName                 = "posts"
	LOCAL_DIR                   = "./posts/" // TB imported at run configs
	ES_POST_INDEX_MAPPING       = "./utils/es-mapping.json"
	ES_POST_MEDIA_INDEX_MAPPING = "./utils/es-media-mapping.json"
	ES_POST_INDEX_NAME          = "post_idx_v0001"
	ES_POST_MEDIA_INDEX_NAME    = "post_media_idx_v0001"
	ES_POST_INDEX_ALIAS         = "post_index"
	ES_POST_MEDIA_INDEX_ALIAS   = "post_media_index"
	ElasticSearchURL            = "http://localhost:9200"
)

var (
	S3Client *s3.S3
	ESClient *elastic.Client
)

// Initialize initializes the AWS session, S3 client, and Elasticsearch client
func Initialize() {
	// Initialize AWS session
	//sess := session.Must(session.NewSession(&aws.Config{
	//	Region: aws.String("us-east-1"),
	//}))
	//S3Client = s3.New(sess)
	//
	//// Initialize Elasticsearch client
	//var err error
	//ESClient, err = elastic.NewClient(elastic.SetURL(ElasticSearchURL), elastic.SetSniff(false))
	//if err != nil {
	//	log.Fatalf("Error initializing Elasticsearch client: %v", err)
	//}
}
