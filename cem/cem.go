package cem

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	httpJson "github.com/infamousjoeg/cybr-cli/pkg/cybr/helpers/httpjson"
)

func Login(org string, apikey string) (string, error) {

	if org == "" {
		return "", errors.New("empty organization")
	}
	if apikey == "" {
		return "", errors.New("empty API Key")
	}
	url := "https://api.cem.cyberark.com/apis/login"
	var jsonStr = []byte(`{ "organization" : ` + org + `, "accessKey": ` + apikey + `}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	s := strings.Split(string(body), `"`)
	return s[3], nil
}

func GetAccounts(token string) (string, error) {
	url := "https://api.cem.cyberark.com/customer/platforms/accounts"

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body), nil
}

type EntityQuery struct {
	Platform  string `query_key:"platform"`
	AccountId string `query_key:"account_id"`
	EntityId  string `query_key:"entity_id"`
}

func GetEntityRemediations(token string, query *EntityQuery) (string, error) {
	baseUrl := "https://api.cem.cyberark.com/recommendations/remediations"
	url := fmt.Sprintf("%s%s", baseUrl, httpJson.GetURLQuery(query))

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body), nil
}

func GetEntityRecommendations(token string, query *EntityQuery) (string, error) {
	baseUrl := "https://api.cem.cyberark.com/recommendations/api/metadata"
	url := fmt.Sprintf("%s%s", baseUrl, httpJson.GetURLQuery(query))

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body), nil
}

func GetEntityDetail(token string, query *EntityQuery) (string, error) {
	baseUrl := "https://api.cem.cyberark.com/cloudEntities/api/get-entity-details"
	url := fmt.Sprintf("%s%s", baseUrl, httpJson.GetURLQuery(query))

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body), nil
}

type GetEntitiesQuery struct {
	Platform    string `query_key:"platform"`
	AccountId   string `query_key:"account_id"`
	FullAdmin   string `query_key:"full_admin"`
	ShadowAdmin string `query_key:"shadow_admin"`
	NextToken   string `query_key:"next_token"`
}

func GetEntities(token string, query *GetEntitiesQuery) (string, error) {
	baseUrl := "https://api.cem.cyberark.com/cloudEntities/api/search"
	url := fmt.Sprintf("%s%s", baseUrl, httpJson.GetURLQuery(query))

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body), nil
}
