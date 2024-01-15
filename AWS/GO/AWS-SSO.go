package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssooidc"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Error loading AWS config:", err)
		return
	}

	client := ssooidc.NewFromConfig(cfg)

	// Example: Initiate the authentication flow
	startURL := "https://your-sso-instance.awsapps.com/start"
	clientID := "your-client-id"
	clientSecret := "your-client-secret"

	resp, err := client.StartDeviceAuthorization(context.TODO(), &ssooidc.StartDeviceAuthorizationInput{
		ClientId:     &clientID,
		ClientSecret: &clientSecret,
		StartUrl:     &startURL,
	})
	if err != nil {
		fmt.Println("Error starting device authorization:", err)
		return
	}

	fmt.Println("Device code:", *resp.DeviceCode)
	fmt.Println("User code:", *resp.UserCode)
	fmt.Println("Verification URI:", *resp.VerificationUri)
}
