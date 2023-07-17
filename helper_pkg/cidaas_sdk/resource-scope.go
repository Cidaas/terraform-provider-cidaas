package cidaas_sdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Scope struct {
	ID                    string                  `json:"_id,omitempty"`
	LocaleWiseDescription []ScopeLocalDescription `json:"localeWiseDescription,omitempty"`
	SecurityLevel         string                  `json:"securityLevel,omitempty"`
	ScopeKey              string                  `json:"scopeKey,omitempty"`
	RequiredUserConsent   bool                    `json:"requiredUserConsent,omitempty"`
	GroupName             []string                `json:"group_name,omitempty"`
	ScopeOwner            string                  `json:"scopeOwner,omitempty"`
}

type ScopeLocalDescription struct {
	Locale      string `json:"locale,omitempty"`
	Language    string `json:"language,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

type ScopeResponse struct {
	Success bool  `json:"success,omitempty"`
	Status  int   `json:"status,omitempty"`
	Data    Scope `json:"data,omitempty"`
	// Errors  Error `json:"error,omitempty"`
	Error string `json:"error,omitempty"`
	// RefNum  string    `json:"renum,omitempty"`
}

func CreateOrUpdateScope(cidaas_client CidaasClient, sc Scope) (base_response ScopeResponse) {

	url := cidaas_client.BaseUrl + "/scopes-srv/scope"
	method := "POST"

	json_payload, err := json.Marshal(sc)
	if err != nil {
		fmt.Println(err)
		return
	}
	payload_string := string(json_payload)
	payload := strings.NewReader(payload_string)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Authorization", "Bearer "+cidaas_client.TokenData.AccessToken)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json, text/plain, */*")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	json.Unmarshal([]byte(body), &base_response)

	return base_response
}

func GetScope(cidaas_client CidaasClient, scope_key string) (base_response ScopeResponse) {
	url := cidaas_client.BaseUrl + "/scopes-srv/scope?scopekey=" + strings.ToLower(scope_key)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Authorization", "Bearer "+cidaas_client.TokenData.AccessToken)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	json.Unmarshal([]byte(body), &base_response)

	return base_response
}

func DeleteScope(cidaas_client CidaasClient, scope_key string) (base_response ScopeResponse) {
	url := cidaas_client.BaseUrl + "/scopes-srv/scope/" + strings.ToLower(scope_key)
	method := "DELETE"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Authorization", "Bearer "+cidaas_client.TokenData.AccessToken)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	json.Unmarshal([]byte(body), &base_response)

	return base_response
}
