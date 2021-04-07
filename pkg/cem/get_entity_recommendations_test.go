package cem_test

import (
	"testing"

	"github.com/quincycheng/cem-api-go/pkg/cem"
	"github.com/quincycheng/cem-api-go/pkg/test"
)

func TestGetEntityRecomendationsSuccess(t *testing.T) {
	token := login(t)
	resp, err := cem.GetEntityRecommendations(token, &cem.EntityQuery{
		Platform:  "azure",
		AccountID: "43ff4d9a-e162-4c20-9c0d-5f249199cf48",
		EntityID:  "71b781f5-bd5d-4e7b-af5e-ed9884b24b15",
	})
	test.AssertNoError(t, err)
	test.AssertIntGreaterThan(t, len(resp.Recommendations), 0)
}

func TestGetEntityRecomendationsInvalidAccountID(t *testing.T) {
	token := login(t)
	_, err := cem.GetEntityRecommendations(token, &cem.EntityQuery{
		Platform:  "invalidPlatform",
		AccountID: "invalidAccountID",
		EntityID:  "invalidEntityID",
	})

	test.AssertError(t, err, "invalid query was provided")
	test.AssertStringContains(t, err.Error(), "400 Bad Request")
}

func TestGetEntityRecomendationsInvalidToken(t *testing.T) {
	_, err := cem.GetEntityRecommendations("invalidToken", &cem.EntityQuery{})
	test.AssertError(t, err, "Should not be able to list entities with invalid token")
}
