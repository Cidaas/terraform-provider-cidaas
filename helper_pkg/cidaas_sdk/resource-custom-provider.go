package cidaas_sdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Scopes struct {
	DisplayLabel string        `json:"display_label,omitempty"`
	Scopes       []ScopesChild `json:"scopes,omitempty"`
}

type ScopesChild struct {
	ScopeName  string `json:"scope_name,omitempty"`
	Required   bool   `json:"required,omitempty"`
	Recommened bool   `json:"recommened,omitempty"`
}
type CustomProvider struct {
	ID                    string                 `json:"_id,omitempty"`
	ClientId              string                 `json:"client_id,omitempty"`
	ClientSecret          string                 `json:"client_secret,omitempty"`
	DisplayName           string                 `json:"display_name,omitempty"`
	StandardType          string                 `json:"standard_type,omitempty"`
	AuthorizationEndpoint string                 `json:"authorization_endpoint,omitempty"`
	TokenEndpoint         string                 `json:"token_endpoint,omitempty"`
	ProviderName          string                 `json:"provider_name,omitempty"`
	LogoUrl               string                 `json:"logo_url,omitempty"`
	UserinfoEndpoint      string                 `json:"userinfo_endpoint,omitempty"`
	UserinfoFields        map[string]interface{} `json:"userinfo_fields,omitempty"`
	Scopes                Scopes                 `json:"scopes,omitempty"`
}

type UserInfo struct {
	Name              string        `json:"name,omitempty"`
	FamilyName        string        `json:"family_name,omitempty"`
	GivenName         string        `json:"given_name,omitempty"`
	MiddleName        string        `json:"middle_name,omitempty"`
	Nickname          string        `json:"nickname,omitempty"`
	PreferredUsername string        `json:"preferred_username,omitempty"`
	Profile           string        `json:"profile,omitempty"`
	Picture           string        `json:"picture,omitempty"`
	Website           string        `json:"website,omitempty"`
	Gender            string        `json:"gender,omitempty"`
	Birthdate         string        `json:"birthdate,omitempty"`
	Zoneinfo          string        `json:"zoneinfo,omitempty"`
	Locale            string        `json:"locale,omitempty"`
	Updated_at        string        `json:"updated_at,omitempty"`
	Email             string        `json:"email,omitempty"`
	EmailVerified     string        `json:"email_verified,omitempty"`
	PhoneNumber       string        `json:"phone_number,omitempty"`
	MobileNumber      string        `json:"mobile_number,omitempty"`
	Address           string        `json:"address,omitempty"`
	CustomFields      []interface{} `json:"custom_fields,omitempty"`
}

type CustomFields struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
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
	url := cidaas_client.BaseUrl + "/providers-srv/custom/" + provider_name
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
	url := cidaas_client.BaseUrl + "/providers-srv/custom/" + strings.ToLower(provider)
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
