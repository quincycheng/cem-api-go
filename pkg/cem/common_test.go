package cem_test

import (
	"os"
	"testing"

	"github.com/quincycheng/cem-api-go/pkg/cem"
)

var (
	Organization = os.Getenv("CEM_ORG")
	ApiKey       = os.Getenv("CEM_API_KEY")
)

func login(t *testing.T) string {
	token, err := cem.Login(Organization, ApiKey)
	if err != nil {
		t.Fatalf("Failed to login. %s", err)
	}

	return token
}
