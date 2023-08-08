package microsoftgraph

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/pkg/errors"
	"log"
)

type Client struct {
	client *msgraphsdk.GraphServiceClient
}

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
