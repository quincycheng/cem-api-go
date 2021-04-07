package cem

import (
	"encoding/json"
	"time"

	"github.com/quincycheng/cem-api-go/pkg/utils"
)

type GetEntityRecommendationsResponse struct {
	Platform        string            `json:"platform"`
	AccountID       string            `json:"account_id"`
	EntityID        string            `json:"entity_id"`
	Recommendations []Recommendations `json:"recommendations"`
}
type Recommendations struct {
	CreatedDate           time.Time `json:"created_date"`
	ActiveRecommendations []string  `json:"active_recommendations"`
	Status                string    `json:"status"`
	ExecTime              string    `json:"exec_time"`
}

func GetEntityRecommendations(token string, query *EntityQuery) (*GetEntityRecommendationsResponse, error) {
	req := utils.Request{
		BaseURL:  BaseURL,
		Method:   "GET",
		Endpoint: "/recommendations/api/metadata",
		Token:    token,
		Query:    query,
	}
	obj := &GetEntityRecommendationsResponse{}

	response, err := utils.SendHTTPRequest(req)
	if err != nil {
		return obj, err
	}

	err = json.Unmarshal(response, obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}
