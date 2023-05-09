package cidaas_sdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Scopes struct {
	DisplayLabel string              `json:"display_label,omitempty"`
	Scopes       []map[string]string `json:"scopes,omitempty"`
}

type CustomProvider struct {
	ID                    string   `json:"_id,omitempty"`
	ClientId              string   `json:"client_id,omitempty"`
	ClientSecret          string   `json:"client_secret,omitempty"`
	DisplayName           string   `json:"display_name,omitempty"`
	StandardType          string   `json:"standard_type,omitempty"`
	AuthorizationEndpoint string   `json:"authorization_endpoint,omitempty"`
	TokenEndpoint         string   `json:"token_endpoint,omitempty"`
	ProviderName          string   `json:"provider_name,omitempty"`
	LogoUrl               string   `json:"logo_url,omitempty"`
	UserinfoEndpoint      string   `json:"userinfo_endpoint,omitempty"`
	UserinfoFields        UserInfo `json:"userinfo_fields,omitempty"`
	Scopes                Scopes   `json:"scopes,omitempty"`
}

type UserInfo struct {
	Name string `json:"name,omitempty"`
}

type CpBaseResponse struct {
	Success bool           `json:"success,omitempty"`
	Status  int            `json:"status,omitempty"`
	Data    CustomProvider `json:"data,omitempty"`
	Errors  Error          `json:"error,omitempty"`
	// Error   string    `json:"error,omitempty"`
	// RefNum  string    `json:"renum,omitempty"`
}

type LinkCustomProviderStruct struct {
	ClientId    string `json:"client_id,omitempty"`
	Test        bool   `json:"deleted"`
	Type        string `json:"type,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
}

type LinkCpResponse struct {
	Success bool `json:"success,omitempty"`
	Status  int  `json:"status,omitempty"`
	Data    struct {
		Updated bool `json:"updated,omitempty"`
	} `json:"data,omitempty"`
	Error Error `json:"error,omitempty"`
}

func CreateCustomProvider(cidaas_client CidaasClient, cp CustomProvider) (base_response CpBaseResponse) {

	url := cidaas_client.ProvideUrl
	method := "POST"

	json_payload, err := json.Marshal(cp)
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
	// fmt.Println(string(body))

	json.Unmarshal([]byte(body), &base_response)

	return base_response
}

func UpdateCustomProvider(cidaas_client CidaasClient, body CustomProvider) (base_response CpBaseResponse) {

	url := cidaas_client.ProvideUrl
	method := "PUT"

	json_payload, err := json.Marshal(body)
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

	resp, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	json.Unmarshal([]byte(resp), &base_response)

	return base_response
}

func GetCustomProvider(cidaas_client CidaasClient, provider_name string) (base_response CpBaseResponse) {
	url := "https://kube-nightlybuild-dev.cidaas.de/providers-srv/custom/" + provider_name
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

func DeleteCustomProvider(cidaas_client CidaasClient, provider string) (base_response CpBaseResponse) {
	url := cidaas_client.BaseUrl + "/providers-srv/custom/" + provider
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

func LinkCustomProvider(cidaas_client CidaasClient, cp LinkCustomProviderStruct) (base_response LinkCpResponse) {

	url := cidaas_client.BaseUrl + "/apps-srv/loginproviders/update/" + cp.DisplayName
	method := "PUT"

	json_payload, err := json.Marshal(cp)
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
