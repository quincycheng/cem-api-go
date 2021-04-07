package cem

import (
	"encoding/json"

	"github.com/quincycheng/cem-api-go/pkg/utils"
)

type GetEntitiesQuery struct {
	Platform    string `query_key:"platform"`
	AccountId   string `query_key:"account_id"`
	FullAdmin   string `query_key:"full_admin"`
	ShadowAdmin string `query_key:"shadow_admin"`
	NextToken   string `query_key:"next_token"`
}

type GetEntitiesResponse struct {
	Hits      []Hits        `json:"hits"`
	Total     int           `json:"total"`
	NextToken []interface{} `json:"next_token"`
}

type Hits struct {
	EntityID        string   `json:"entityId"`
	EntityName      string   `json:"entityName"`
	EntityType      string   `json:"entityType"`
	AccountID       string   `json:"accountId"`
	AccountName     string   `json:"accountName"`
	PlatformName    string   `json:"platformName"`
	IsShadowAdmin   bool     `json:"isShadowAdmin"`
	IsFullAdmin     bool     `json:"isFullAdmin"`
	Recommendations []string `json:"recommendations"`
	RiskTotalScore  int      `json:"riskTotalScore"`
	Status          string   `json:"status"`
}

func GetEntities(token string, query *GetEntitiesQuery) (*GetEntitiesResponse, error) {
	req := utils.Request{
		BaseURL:  BaseURL,
		Method:   "GET",
		Endpoint: "/cloudEntities/api/search",
		Token:    token,
		Query:    query,
	}
	obj := &GetEntitiesResponse{}

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
