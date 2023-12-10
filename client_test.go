package microsoftgraph

import (
	"os"
	"testing"
)

func TestNewClientByDeviceCodeCredential(t *testing.T) {
	_, err := NewClientByDeviceCodeCredential(
		os.Getenv("TENANT_ID"),
		os.Getenv("CLIENT_ID"),
		[]string{"Files.Read"},
	)
	if err != nil {
		t.Error(err)
	}
}

func TestNewClientBySecretCredential(t *testing.T) {
	_, err := NewClientBySecretCredential(
		os.Getenv("TENANT_ID"),
		os.Getenv("CLIENT_ID"),
		os.Getenv("CLIENT_SECRET"),
	)
	if err != nil {
		t.Error(err)
	}
}
