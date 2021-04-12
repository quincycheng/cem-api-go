package cem_test

import (
	"testing"

	"github.com/quincycheng/cem-api-go/pkg/cem"
	"github.com/quincycheng/cem-api-go/pkg/test"
)

func TestGetEntitiesSuccess(t *testing.T) {
	token := login(t)
	resp, err := cem.GetEntities(token, &cem.GetEntitiesQuery{})
	test.AssertNoError(t, err)
	test.AssertIntGreaterThan(t, resp.Total, 0)
}

func TestGetEntitiesInvalidToken(t *testing.T) {
	_, err := cem.GetEntities("invalidToken", &cem.GetEntitiesQuery{})
	test.AssertError(t, err, "Should fail with invalid token")
}
