package microsoftgraph

import (
	"context"
	"os"
	"testing"
)

func TestGetDrive(t *testing.T) {
	// t.SkipNow()
	client, err := NewClientByDeviceCodeCredential(
		os.Getenv("TENANT_ID"),
		os.Getenv("CLIENT_ID"),
		[]string{"Files.Read"},
	)
	if err != nil {
		t.Error(err)
	}
	_, err = client.ListDrives(context.Background(), nil)
	if err != nil {
		t.Error(err)
	}
}
