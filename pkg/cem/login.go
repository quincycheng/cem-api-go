package cem

import (
	"encoding/json"
	"errors"

	"github.com/quincycheng/cem-api-go/pkg/utils"
)

type LoginRequest struct {
	Organization string `json:"organization"`
	AccessKey    string `json:"accessKey"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func Login(org string, apikey string) (string, error) {
	if org == "" {
		return "", errors.New("empty organization")
	}
	if apikey == "" {
		return "", errors.New("empty API Key")
	}

	req := utils.Request{
		BaseURL:  BaseURL,
		Method:   "POST",
		Endpoint: "/apis/login",
		Data: LoginRequest{
			Organization: org,
			AccessKey:    apikey,
		},
	}

	response, err := utils.SendHTTPRequest(req)
	if err != nil {
		return "", err
	}

	loginResponse := &LoginResponse{}
	err = json.Unmarshal(response, loginResponse)
	if err != nil {
		return "", err
	}

	return loginResponse.Token, nil
}
