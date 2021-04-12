package cem_test

import (
	"testing"

	"github.com/quincycheng/cem-api-go/pkg/cem"
	"github.com/quincycheng/cem-api-go/pkg/test"
)

func TestGetAccountsSuccess(t *testing.T) {
	token := login(t)
	resp, err := cem.GetAccounts(token)
	test.AssertNoError(t, err)

	validPlatforms := []string{
		"aws",
		"azure",
		"gcp",
	}

	for _, data := range resp.Data {
		test.AssertStringInList(t, data.Platform, validPlatforms)
	}
}

func TestGetAccountsInvalidToken(t *testing.T) {
	_, err := cem.GetAccounts("invalidToken")
	test.AssertError(t, err, "Expecting error when getting accounts with invalid token")
}
