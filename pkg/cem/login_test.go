package cem_test

import (
	"testing"

	"github.com/quincycheng/cem-api-go/pkg/cem"
	"github.com/quincycheng/cem-api-go/pkg/test"
)

func TestLogin(t *testing.T) {
	token, err := cem.Login(Organization, ApiKey)
	test.AssertNoError(t, err)
	test.AssertStringNotEmpty(t, token)
}

func TestLoginEmptyOrg(t *testing.T) {
	_, err := cem.Login("", "mock")
	test.AssertError(t, err, "Organization is empty")
}

func TestLoginEmptyApiKey(t *testing.T) {
	_, err := cem.Login("mock", "")
	test.AssertError(t, err, "Api Key is empty")
}

func TestLoginInvalidCredentials(t *testing.T) {
	_, err := cem.Login("mock", "mock")
	test.AssertError(t, err, "Credentials are invalid")
}
