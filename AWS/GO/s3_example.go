package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	// Create a new AWS session using the default credentials and region.
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"), // Replace with your desired AWS region
	}))

	// Create an S3 service client.
	svc := s3.New(sess)

	// Specify the bucket name you want to list objects from.
	bucketName := "your-bucket-name"

	// Call the ListObjectsV2 operation to list objects in the bucket.
	result, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to list objects, %v\n", err)
		os.Exit(1)
	}

	// Print the object keys.
	fmt.Println("Objects in the bucket:")
	for _, obj := range result.Contents {
		fmt.Println(*obj.Key)
	}
}
