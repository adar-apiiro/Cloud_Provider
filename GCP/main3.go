package main

import (
	"context"
	"fmt"
	"time"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"github.com/googleapis/gax-go/v2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

func main() {
	// Set your Google Cloud project ID
	projectID := "your-project-id"

	// Authentication using Application Default Credentials (ADC)
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx, option.WithEndpoint("your-endpoint"), option.WithProjectID(projectID))
	if err != nil {
		fmt.Printf("Error creating Secret Manager client: %v\n", err)
		return
	}
	defer client.Close()

	// Timeout for the RPC call
	timeout := 10 * time.Second
	ctxWithTimeout, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	// Create a secret (replace "your-secret-id" and "your-secret-data" accordingly)
	secretID := "your-secret-id"
	secretData := []byte("your-secret-data")

	req := &secretmanagerpb.CreateSecretRequest{
		Parent:   fmt.Sprintf("projects/%s", projectID),
		SecretId: secretID,
		Secret:   &secretmanagerpb.Secret{Replication: &secretmanagerpb.Replication{Replication: &secretmanagerpb.Replication_Automatic_{}}},
	}
	secret, err := client.CreateSecret(ctxWithTimeout, req)
	if err != nil {
		fmt.Printf("Error creating secret: %v\n", err)
		return
	}

	// Add a version to the secret
	// Replace "your-version-id" and "your-version-data" accordingly
	versionID := "your-version-id"
	versionData := []byte("your-version-data")

	req = &secretmanagerpb.AddSecretVersionRequest{
		Parent:  secret.GetName(),
		Payload: &secretmanagerpb.SecretPayload{Data: versionData},
	}
	_, err = client.AddSecretVersion(ctxWithTimeout, req)
	if err != nil {
		fmt.Printf("Error adding secret version: %v\n", err)
		return
	}

	// Fetch the secret version (replace "latest" with your version ID)
	versionName := fmt.Sprintf("%s/versions/latest", secret.GetName())
	version, err := client.AccessSecretVersion(ctxWithTimeout, &secretmanagerpb.AccessSecretVersionRequest{Name: versionName})
	if err != nil {
		fmt.Printf("Error accessing secret version: %v\n", err)
		return
	}

	fmt.Printf("Secret Data: %s\n", string(version.Payload.Data))
}
