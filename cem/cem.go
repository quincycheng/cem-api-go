package cem

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
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
	if err != nil {
		panic(err)
	}
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
	if err != nil {
		panic(err)
	}
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
	url := fmt.Sprintf("%s%s", baseUrl, GetURLQuery(query))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
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
	url := fmt.Sprintf("%s%s", baseUrl, GetURLQuery(query))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
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
	url := fmt.Sprintf("%s%s", baseUrl, GetURLQuery(query))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
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
	url := fmt.Sprintf("%s%s", baseUrl, GetURLQuery(query))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
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

// GetURLQuery converts a struct into a url query. The struct must have a tagname of 'query_key'
func GetURLQuery(queryParams interface{}) string {
	val := reflect.ValueOf(queryParams).Elem()
	req, _ := http.NewRequest("GET", "http://localhost", nil)
	query := req.URL.Query()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		tag := val.Type().Field(i).Tag

		key := tag.Get("query_key")
		value := valueField.Interface()

		// If value is nil or equal to 0 do not include the query param
		if value == nil || fmt.Sprintf("%v", value) == "0" || key == "" || fmt.Sprintf("%v", value) == "" {
			continue
		}

		query.Add(key, fmt.Sprintf("%v", value))
	}

	// possible for query to be completely empty
	if query.Encode() == "" {
		return ""
	}

	// append '?' to prefix
	return "?" + query.Encode()
}
