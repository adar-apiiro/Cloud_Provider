package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"cloud.google.com/go/billing/apiv1"
	"cloud.google.com/go/cloudbuild/apiv1/v2"
	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	billingpb "google.golang.org/genproto/googleapis/cloud/billing/v1"
	cloudbuildpb "google.golang.org/genproto/googleapis/devtools/cloudbuild/v2"
)

const (
	projectID        = "your-project-id"
	credentialsFile = "path/to/your/credentials-file.json"
)

func main() {
	// Google Cloud Storage Example
	listBucketsExample()

	// Cloud Build Example
	createBuildExample()
}

func listBucketsExample() {
	ctx := context.Background()

	// Set up a Google Cloud Storage client
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialsFile))
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
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		bucketNames = append(bucketNames, bucketAttrs.Name)
	}

	return bucketNames, nil
}

func createBuildExample() {
	ctx := context.Background()

	// Set up a Cloud Build client
	buildClient, err := cloudbuild.NewClient(ctx, option.WithCredentialsFile(credentialsFile))
	if err != nil {
		log.Fatalf("Error creating Cloud Build client: %v", err)
	}
	defer buildClient.Close()

	// Create a build
	buildID, err := createBuild(buildClient)
	if err != nil {
		log.Fatalf("Error creating build: %v", err)
	}

	fmt.Printf("Build created with ID: %s\n", buildID)
}

func createBuild(client *cloudbuild.Client) (string, error) {
	// Define a build request
	buildRequest := &cloudbuildpb.Build{
		Source: &cloudbuildpb.Source{
			StorageSource: &cloudbuildpb.StorageSource{
				Bucket: "your-source-bucket",
				Object: "your-source-object",
			},
		},
		Steps: []*cloudbuildpb.BuildStep{
			{
				Name: "gcr.io/cloud-builders/go",
				Args: []string{"go", "build", "-o", "app"},
			},
		},
		Timeout: &durationpb.Duration{Seconds: 600}, // Set a timeout of 10 minutes
	}

	// Create a build in Cloud Build
	resp, err := client.CreateBuild(context.Background(), &cloudbuildpb.CreateBuildRequest{ProjectId: projectID, Build: buildRequest})
	if err != nil {
		return "", err
	}

	return resp.Id, nil
}
