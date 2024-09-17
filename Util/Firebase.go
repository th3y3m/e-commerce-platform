package Util

import (
	"context"
	"fmt"
	"io"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func UploadFileToFireBase(bucketName, objectName, filePath string) (string, error) {
	ctx := context.Background()

	client, err := storage.NewClient(ctx, option.WithCredentialsFile("path/to/firebase_credentials.json"))
	if err != nil {
		return "", fmt.Errorf("failed to create client: %v", err)
	}
	defer client.Close()

	// Open the file to upload
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Get the bucket handle
	bucket := client.Bucket(bucketName)

	// Create a new writer to upload the file to the specified path (objectName) in the bucket
	wc := bucket.Object(objectName).NewWriter(ctx)

	// Write the file to Google Cloud Storage
	if _, err = io.Copy(wc, file); err != nil {
		return "", fmt.Errorf("failed to write file: %v", err)
	}

	// Close the writer to complete the upload
	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("failed to close writer: %v", err)
	}

	// Construct the public URL of the uploaded file
	publicURL := fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucketName, objectName)

	fmt.Printf("File %s uploaded to bucket %s as %s\n", filePath, bucketName, objectName)
	fmt.Printf("Public URL: %s\n", publicURL)

	return publicURL, nil
}
