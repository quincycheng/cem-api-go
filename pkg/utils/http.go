package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Request struct {
	BaseURL  string
	Method   string
	Endpoint string
	Token    string
	Query    interface{}
	Data     interface{}
}

func GetHTTPRequest(req Request) (*http.Request, error) {
	var body []byte
	var err error

	url := fmt.Sprintf("%s%s%s", req.BaseURL, req.Endpoint, GetURLQuery(req.Query))

	// Add body if provided
	if req.Data != nil {
		body, err = json.Marshal(req.Data)
	}
	if err != nil {
		return nil, err
	}

	// Create http.Request
	httpRequest, err := http.NewRequest(req.Method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// Add token to header if provided
	if req.Token != "" {
		httpRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", req.Token))
	}

	// Always set Content-Type
	httpRequest.Header.Set("Content-Type", "application/json")

	return httpRequest, nil
}

func GetDefaultHTTPClient() *http.Client {
	return &http.Client{}
}

func SendHTTPRequest(req Request) ([]byte, error) {
	client := GetDefaultHTTPClient()
	httpReq, err := GetHTTPRequest(req)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return body, fmt.Errorf("Received invalid status code '%s'", resp.Status)
	}

	return body, nil
}
