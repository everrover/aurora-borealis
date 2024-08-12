package services

import (
	utils "aurora-borealis/utils"
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

// SaveToS3 uploads the markdown content to S3
func SaveToS3(title, content string) error {
	_, err := utils.S3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(utils.S3BucketName),
		Key:    aws.String(title + ".md"),
		Body:   bytes.NewReader([]byte(content)),
		ACL:    aws.String("private"),
	})
	return err
}
