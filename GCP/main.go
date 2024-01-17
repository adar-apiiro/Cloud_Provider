package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

const (
	projectID      = "your-project-id"
	credentialsFile = "path/to/your/credentials-file.json"
)

func main() {
	// Set up a Google Cloud Storage client
	client, err := storage.NewClient(context.Background(), option.WithCredentialsFile(credentialsFile))
	if err != nil {
		log.Fatalf("Error creating Google Cloud Storage client: %v", err)
	}
	defer client.Close()

	// List buckets in the project
	buckets, err := listBuckets(client)
	if err != nil {
		log.Fatalf("Error listing buckets: %v", err)
	}

	fmt.Println("Google Cloud Storage Buckets:")
	for _, bucket := range buckets {
		fmt.Println(bucket)
	}
}

func listBuckets(client *storage.Client) ([]string, error) {
	var bucketNames []string

	// List all buckets in the project
	it := client.Buckets(context.Background(), projectID)
	for {
		bucketAttrs, err := it.Next()
		if err == storage.IterationDone {
			break
		}
		if err != nil {
			return nil, err
		}
		bucketNames = append(bucketNames, bucketAttrs.Name)
	}

	return bucketNames, nil
}
