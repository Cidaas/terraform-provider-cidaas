package cidaas_sdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type baseCustomTemplateResponse struct {
	Success bool                 `json:"success"`
	Status  int                  `json:"status"`
	Data    CustomTemplateConfig `json:"data"`
	Error   Error                `json:"error"`
}

type CustomTemplateConfig struct {
	GroupId       string `json:"group_id"`
	ClassName     string `json:"className"`
	Content       string `json:"content"`
	Subject       string `json:"subject"`
	Language      string `json:"language"`
	Locale        string `json:"locale"`
	UsageType     string `json:"usageType"`
	TemplateType  string `json:"templateType"`
	TemplateKey   string `json:"templateKey"`
	LastSeededBy  string `json:"last_seeded_by"`
	TemplateOwner string `json:"templateOwner"`
	Id            string `json:"Id"`
	CreatedTime   string `json:"createdTime"`
	UpdatedTime   string `json:"updatedTime"`
	Deleted       bool   `json:"deleted"`
}

func CreateCustomTemplate(cidaas_client CidaasClient, custom_template_config CustomTemplateConfig) (base_response baseCustomTemplateResponse) {

	url := cidaas_client.BaseUrl + "/templates-srv/template/custom"
	method := "POST"

	json_payload, err := json.Marshal(custom_template_config)
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

func GetCustomTemplate(cidaas_client CidaasClient, template_id string) (base_response baseCustomTemplateResponse) {

	url := cidaas_client.BaseUrl + "/templates-srv/template/custom/" + template_id
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

func DeleteCustomTemplate(cidaas_client CidaasClient, template_key string, template_type string) (base_response baseCustomTemplateResponse) {

	url := cidaas_client.BaseUrl + "/templates-srv/template/custom/" + template_key + "/" + template_type
	method := "DELETE"

	fmt.Println(url)

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
