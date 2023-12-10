package microsoftgraph

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/pkg/errors"
	"log"
)

// Client is a struct that holds a pointer to a GraphServiceClient.
type Client struct {
	client *msgraphsdk.GraphServiceClient
}

// NewClientByDeviceCodeCredential creates a new Client using a Device Code Credential.
// tenantID: The ID of the tenant.
// clientID: The ID of the client.
// scopes: The scopes for the request.
// Returns a pointer to a Client and an error, if any occurred.
func NewClientByDeviceCodeCredential(tenantID, clientID string, scopes []string) (*Client, error) {
	cred, err := azidentity.NewDeviceCodeCredential(&azidentity.DeviceCodeCredentialOptions{
		TenantID: tenantID,
		ClientID: clientID,
		UserPrompt: func(ctx context.Context, message azidentity.DeviceCodeMessage) error {
			log.Println(message.Message)
			return nil
		},
	})

	if err != nil {
		return nil, errors.New("new device code credential fail: " + err.Error())
	}

	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)

	if err != nil {
		return nil, errors.New("error creating client: " + err.Error())
	}
	return &Client{client: client}, nil
}

// NewClientBySecretCredential creates a new Client using a Secret Credential.
// tenantID: The ID of the tenant.
// clientID: The ID of the client.
// clientSecret: The secret of the client.
// Returns a pointer to a Client and an error, if any occurred.
func NewClientBySecretCredential(tenantID, clientID, clientSecret string) (*Client, error) {
	cred, err := azidentity.NewClientSecretCredential(
		tenantID,
		clientID,
		clientSecret,
		nil,
	)

	if err != nil {
		log.Fatal("new device code credential fail: ", err)
	}

	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(
		cred, []string{"https://graph.microsoft.com/.default"},
	)
	if err != nil {
		log.Fatal("error creating client: ", err)
	}
	return &Client{client: client}, nil
}
