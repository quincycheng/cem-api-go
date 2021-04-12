package cem

import (
	"encoding/json"

	"github.com/quincycheng/cem-api-go/pkg/utils"
)

type GetEntityDetailResponse struct {
	EntityID      string `json:"entity_id"`
	EntityName    string `json:"entity_name"`
	EntityType    string `json:"entity_type"`
	AccountID     string `json:"account_id"`
	Platform      string `json:"platform"`
	ShadowAdmin   bool   `json:"shadow_admin"`
	Admin         bool   `json:"admin"`
	ExposureLevel int    `json:"exposure_level"`
	AccountName   string `json:"account_name"`
}

func GetEntityDetail(token string, query *EntityQuery) (*GetEntityDetailResponse, error) {
	req := utils.Request{
		BaseURL:  BaseURL,
		Method:   "GET",
		Endpoint: "/cloudEntities/api/get-entity-details",
		Token:    token,
		Query:    query,
	}
	obj := &GetEntityDetailResponse{}

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
