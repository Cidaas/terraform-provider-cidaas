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

type RegistrationFieldConfig struct {
	Internal        bool                   `json:"internal"`
	ReadOnly        bool                   `json:"read_only"`
	Claimable       bool                   `json:"claimable"`
	Required        bool                   `json:"required"`
	Scopes          []string               `json:"scopes"`
	Enabled         bool                   `json:"enabled"`
	LocaleText      map[string]interface{} `json:"localeText"`
	IsGroup         bool                   `json:"is_group"`
	IsList          bool                   `json:"is_list"`
	ParentGroupId   string                 `json:"parent_group_id"`
	FieldType       string                 `json:"fieldType"`
	ConsentRefs     []string               `json:"consent"`
	ClassName       string                 `json:"className"`
	Id              string                 `json:"_id"`
	FieldKey        string                 `json:"fieldKey"`
	DataType        string                 `json:"dataType"`
	Order           int                    `json:"order"`
	FieldDefinition map[string]interface{} `json:"fieldDefinition"`
	CreatedTime     string                 `json:"createdTime"`
	UpdatedTime     string                 `json:"updatedTime"`
	Deleted         bool                   `json:"deleted"`
}

func CreateRegistrationField(cidaas_client CidaasClient, registration_field_config RegistrationFieldConfig) (base_response baseResponseRegistrationField) {

	url := cidaas_client.BaseUrl + "/registration-setup-srv/fields"
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
	// fmt.Println(string(body))

	json.Unmarshal([]byte(body), &base_response)

	return base_response

}

func UpdateRegistrationField(cidaas_client CidaasClient, registration_field_config RegistrationFieldConfig) (base_response baseResponseRegistrationField) {

	url := cidaas_client.BaseUrl + "/registration-setup-srv/fields"
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
	// fmt.Println(string(body))

	json.Unmarshal([]byte(body), &base_response)

	return base_response

}

func GetRegistrationField(cidaas_client CidaasClient, registration_field_key string) (base_response baseResponseRegistrationField) {

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
	// fmt.Println(string(body))

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
	// fmt.Println(string(body))

	json.Unmarshal([]byte(body), &base_response)

	return base_response

}
