//go:build go1.18
// +build go1.18

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/blobs"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

const accountName = "your-storage-account-name"
const accountKey = "your-storage-account-key"
const containerName = "your-container-name"
const blobName = "example-blob.txt"

func main() {
	// Create a pipeline using your storage account credentials
	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		log.Fatal("Error creating shared key credential:", err)
	}
	pipeline := azblob.NewPipeline(credential, azblob.PipelineOptions{})

	// Create a blob service URL
	serviceURL := azblob.NewServiceURL(
		azblob.NewAccountURL(accountName, azblob.NewPipeline(credential, azblob.PipelineOptions{})),
		nil)

	// Create a container URL
	containerURL := serviceURL.NewContainerURL(containerName)

	// Create the container (ignore error if it already exists)
	_, _ = containerURL.Create(context.Background(), azblob.Metadata{}, azblob.PublicAccessNone)

	// Upload a blob to the container
	content := []byte("Hello, Azure Blob Storage!")
	blobURL := containerURL.NewBlockBlobURL(blobName)
	_, err = blobs.UploadBufferToBlockBlob(context.Background(), content, blobURL, azblob.UploadToBlockBlobOptions{})
	if err != nil {
		log.Fatal("Error uploading blob:", err)
	}

	// List blobs in the container
	fmt.Println("Blobs in the container:")
	for marker := (azblob.Marker{}); marker.NotDone(); {
		listBlob, err := containerURL.ListBlobsFlatSegment(context.Background(), marker, azblob.ListBlobsSegmentOptions{})
		if err != nil {
			log.Fatal("Error listing blobs:", err)
		}
		for _, blobItem := range listBlob.Segment.BlobItems {
			fmt.Println(" -", blobItem.Name)
		}
		marker = listBlob.NextMarker
	}
}
