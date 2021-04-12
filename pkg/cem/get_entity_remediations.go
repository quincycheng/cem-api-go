package cem

import (
	"encoding/json"

	"github.com/quincycheng/cem-api-go/pkg/utils"
)

type EntityQuery struct {
	Platform  string `query_key:"platform"`
	AccountID string `query_key:"account_id"`
	EntityID  string `query_key:"entity_id"`
}

func GetEntityRemediations(token string, query *EntityQuery) (*map[string]interface{}, error) {
	req := utils.Request{
		BaseURL:  BaseURL,
		Method:   "GET",
		Endpoint: "/recommendations/remediations",
		Token:    token,
		Query:    query,
	}
	obj := &map[string]interface{}{}

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
