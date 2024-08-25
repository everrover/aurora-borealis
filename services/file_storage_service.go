package services

import (
	"aurora-borealis/models"
	"aurora-borealis/utils"
	//"context"
	//"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	//"cloud.google.com/go/storage"
)

// Save file to storage mediums based on provided bitmask
// 1 - Local
// 2 - AWS S3
// 4 - GCP Cloud Storage
func SaveFileContents(slug string, post models.Post, storageMediums int) []error {
	var errs []error
	content := CreatePostFileContents(post)
	// Save to local file system
	if storageMediums&1 != 0 {
		err := saveToLocal(slug, content)
		if err != nil {
			errs = append(errs, err)
		}
	}

	// Save to S3
	//if storageMediums&2 != 0 {
	//	err := SaveToS3(title, content)
	//	if err != nil {
	//		return err
	//	}
	//}

	// Save to GCP Cloud Storage
	//if storageMediums&4 != 0 {
	//	err := SaveToGCP(title, content)
	//	if err != nil {
	//		return
	//}

	return nil
}

// SaveToLocal saves the markdown content to the local file system
func saveToLocal(slug, content string) error {
	// Ensure the local directory exists
	if _, err := os.Stat(utils.LOCAL_DIR); os.IsNotExist(err) {
		err := os.MkdirAll(utils.LOCAL_DIR, os.ModePerm)
		if err != nil {
			return err
		}
	}
	// Check if the file already exists
	_, err := os.Stat(filepath.Join(utils.LOCAL_DIR, slug+".md"))
	if err != nil {
		slug = slug + "-(1)"
		_, err = os.Stat(filepath.Join(utils.LOCAL_DIR, slug+".md"))
		if err == nil {
			return err
		}
	}

	// Create the markdown file
	filename := filepath.Join(utils.LOCAL_DIR, slug+".md")
	err = ioutil.WriteFile(filename, []byte(content), 0644)
	return err
}

func saveToGCP(slug, content string) error {
	//ctx := context.Background()
	//client, err := storage.NewClient(ctx, option.WithCredentialsFile("path/to/your/service-account-file.json"))
	//if err != nil {
	//	return fmt.Errorf("failed to create client: %v", err)
	//}
	//defer client.Close()
	//
	//bucketName := "your-bucket-name"
	//bucket := client.Bucket(bucketName)
	//object := bucket.Object(slug + ".md")
	//writer := object.NewWriter(ctx)
	//defer writer.Close()
	//
	//_, err = writer.Write([]byte(content))
	//if err != nil {
	//	return fmt.Errorf("failed to write to bucket: %v", err)
	//}

	return nil
}
