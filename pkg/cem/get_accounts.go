package cem

import (
	"encoding/json"

	"github.com/quincycheng/cem-api-go/pkg/utils"
)

type GetAccountsResponse struct {
	Data []Data `json:"data"`
}
type Accounts struct {
	WorkspaceID     string `json:"workspace_id"`
	WorkspaceStatus string `json:"workspace_status"`
	WorkspaceName   string `json:"workspace_name,omitempty"`
}
type Data struct {
	Platform string     `json:"platform"`
	Accounts []Accounts `json:"accounts"`
}

func GetAccounts(token string) (*GetAccountsResponse, error) {
	req := utils.Request{
		BaseURL:  BaseURL,
		Method:   "GET",
		Endpoint: "/customer/platforms/accounts",
		Token:    token,
	}

	response, err := utils.SendHTTPRequest(req)
	if err != nil {
		return &GetAccountsResponse{}, err
	}

	obj := &GetAccountsResponse{}
	err = json.Unmarshal(response, obj)
	if err != nil {
		return &GetAccountsResponse{}, err
	}

	return obj, nil
}
