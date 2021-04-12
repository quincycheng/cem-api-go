package cem_test

import (
	"testing"

	"github.com/quincycheng/cem-api-go/pkg/cem"
	"github.com/quincycheng/cem-api-go/pkg/test"
)

func TestGetEntityRemediationsSuccess(t *testing.T) {
	token := login(t)
	_, err := cem.GetEntityRemediations(token, &cem.EntityQuery{
		Platform:  "azure",
		AccountID: "43ff4d9a-e162-4c20-9c0d-5f249199cf48",
		EntityID:  "71b781f5-bd5d-4e7b-af5e-ed9884b24b15",
	})
	test.AssertNoError(t, err)
}

func TestGetEntityRemediationsInvalidAccountID(t *testing.T) {
	token := login(t)
	_, err := cem.GetEntityRemediations(token, &cem.EntityQuery{
		Platform:  "invalidPlatform",
		AccountID: "invalidAccountID",
		EntityID:  "invalidEntityID",
	})

	test.AssertError(t, err, "invalid query was provided")
	test.AssertStringContains(t, err.Error(), "400 Bad Request")
}

func TestGetEntityRemediationsInvalidToken(t *testing.T) {
	_, err := cem.GetEntityRemediations("invalidToken", &cem.EntityQuery{})
	test.AssertError(t, err, "Should not be able to list entities with invalid token")
}
