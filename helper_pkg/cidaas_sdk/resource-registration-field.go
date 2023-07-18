package cidaas_sdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type baseResponseRegistrationField struct {
	Success bool                    `json:"success"`
	Status  int                     `json:"status"`
	Data    RegistrationFieldConfig `json:"data"`
	Error   Error                   `json:"error"`
}

type baseGetResponseRegistrationField struct {
	Success bool                       `json:"success"`
	Status  int                        `json:"status"`
	Data    GetRegistrationFieldConfig `json:"data"`
	Error   Error                      `json:"error"`
}
type GetRegistrationFieldConfig struct {
	Internal      bool                     `json:"internal,omitempty"`
	ReadOnly      bool                     `json:"read_only,omitempty"`
	Claimable     bool                     `json:"claimable,omitempty"`
	Required      bool                     `json:"required,omitempty"`
	Scopes        []string                 `json:"scopes,omitempty"`
	Enabled       bool                     `json:"enabled,omitempty"`
	LocaleText    []map[string]interface{} `json:"localeText,omitempty"`
	IsGroup       bool                     `json:"is_group,omitempty"`
	IsList        bool                     `json:"is_list,omitempty"`
	ParentGroupId string                   `json:"parent_group_id,omitempty"`
	FieldType     string                   `json:"fieldType,omitempty"`
	Id            string                   `json:"_id,omitempty"`
	FieldKey      string                   `json:"fieldKey,omitempty"`
	DataType      string                   `json:"dataType,omitempty"`
	Order         int                      `json:"order,omitempty"`
	BaseDataType  string                   `json:"baseDataType,omitempty"`
}

type RegistrationFieldConfig struct {
	Internal      bool                   `json:"internal,omitempty"`
	ReadOnly      bool                   `json:"read_only,omitempty"`
	Claimable     bool                   `json:"claimable,omitempty"`
	Required      bool                   `json:"required,omitempty"`
	Scopes        []string               `json:"scopes,omitempty"`
	Enabled       bool                   `json:"enabled,omitempty"`
	LocaleText    map[string]interface{} `json:"localeText"`
	IsGroup       bool                   `json:"is_group,omitempty"`
	IsList        bool                   `json:"is_list,omitempty"`
	ParentGroupId string                 `json:"parent_group_id,omitempty"`
	FieldType     string                 `json:"fieldType,omitempty"`
	Id            string                 `json:"_id,omitempty"`
	FieldKey      string                 `json:"fieldKey,omitempty"`
	DataType      string                 `json:"dataType,omitempty"`
	Order         int                    `json:"order,omitempty"`
	BaseDataType  string                 `json:"baseDataType,omitempty"`
}

func CreateRegistrationField(cidaas_client CidaasClient, registration_field_config RegistrationFieldConfig) (base_response baseResponseRegistrationField) {

	url := cidaas_client.BaseUrl + "/fieldsetup-srv/fields"
	method := "POST"

	json_payload, err := json.Marshal(registration_field_config)
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

func UpdateRegistrationField(cidaas_client CidaasClient, registration_field_config RegistrationFieldConfig) (base_response baseResponseRegistrationField) {

	url := cidaas_client.BaseUrl + "/fieldsetup-srv/fields"
	method := "POST"

	json_payload, err := json.Marshal(registration_field_config)
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

func GetRegistrationField(cidaas_client CidaasClient, registration_field_key string) (base_response baseGetResponseRegistrationField) {

	url := cidaas_client.BaseUrl + "/registration-setup-srv/fields/flat/field/" + registration_field_key
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

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

func DeleteRegistrationField(cidaas_client CidaasClient, registration_field_key string) (base_response baseResponseRegistrationField) {

	url := cidaas_client.BaseUrl + "/registration-setup-srv/fields/" + registration_field_key
	method := "DELETE"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

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
