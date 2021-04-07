package cem_test

import (
	"testing"

	"github.com/quincycheng/cem-api-go/pkg/cem"
	"github.com/quincycheng/cem-api-go/pkg/test"
)

func TestGetEntityDetailSuccess(t *testing.T) {
	token := login(t)
	_, err := cem.GetEntityDetail(token, &cem.EntityQuery{
		Platform:  "azure",
		AccountID: "43ff4d9a-e162-4c20-9c0d-5f249199cf48",
		EntityID:  "597382af-b4b3-467e-a970-c3df0ae813da",
	})
	test.AssertNoError(t, err)
}

func TestGetEntityDetailInvalidAccountID(t *testing.T) {
	token := login(t)
	_, err := cem.GetEntityDetail(token, &cem.EntityQuery{
		Platform:  "invalidPlatform",
		AccountID: "invalidAccountID",
		EntityID:  "invalidEntityID",
	})

	test.AssertError(t, err, "invalid query was provided")
	test.AssertStringContains(t, err.Error(), "400 Bad Request")
}

func TestGetEntityDetailInvalidToken(t *testing.T) {
	_, err := cem.GetEntityDetail("invalidToken", &cem.EntityQuery{})
	test.AssertError(t, err, "Should not be able to list entities with invalid token")
}
