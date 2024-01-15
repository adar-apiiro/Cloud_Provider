package firstversion

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	// Create an AWS session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"), // Specify your AWS region
	})
	if err != nil {
		log.Fatal("Error creating session:", err)
	}

	// Create an S3 client
	s3Client := s3.New(sess)

	// Example: List buckets
	result, err := s3Client.ListBuckets(nil)
	if err != nil {
		log.Fatal("Error listing buckets:", err)
	}

	// Display bucket names
	fmt.Println("Buckets:")
	for _, bucket := range result.Buckets {
		fmt.Println(*bucket.Name)
	}
}
