package cidaas_sdk

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"strings"
// )

// type baseTemplateResponse struct {
// 	Success bool           `json:"success"`
// 	Status  int            `json:"status"`
// 	Data    TemplateConfig `json:"data"`
// 	Error   Error          `json:"error"`
// }

// type TemplateConfig struct {
// 	GroupId       string `json:"group_id"`
// 	ClassName     string `json:"className"`
// 	Content       string `json:"content"`
// 	Subject       string `json:"subject"`
// 	Language      string `json:"language"`
// 	Locale        string `json:"locale"`
// 	UsageType     string `json:"usageType"`
// 	TemplateType  string `json:"templateType"`
// 	TemplateKey   string `json:"templateKey"`
// 	LastSeededBy  string `json:"last_seeded_by"`
// 	TemplateOwner string `json:"templateOwner"`
// 	Id            string `json:"Id"`
// 	CreatedTime   string `json:"createdTime"`
// 	UpdatedTime   string `json:"updatedTime"`
// 	Deleted       bool   `json:"deleted"`
// }

// func GetTemplate(cidaas_client CidaasClient, template_id string) (base_response baseTemplateResponse) {

// 	url := cidaas_client.BaseUrl + "/templates-srv/template/" + template_id
// 	method := "GET"

// 	client := &http.Client{}
// 	req, err := http.NewRequest(method, url, nil)

// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	req.Header.Add("Authorization", "Bearer "+cidaas_client.TokenData.AccessToken)
// 	req.Header.Add("Content-Type", "application/json")
// 	req.Header.Add("Accept", "application/json, text/plain, */*")

// 	res, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer res.Body.Close()

// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	// fmt.Println(string(body))

// 	json.Unmarshal([]byte(body), &base_response)

// 	return base_response

// }

// func DeleteCustomTemplate(cidaas_client CidaasClient, template_key string, template_type string) (base_response baseCustomTemplateResponse) {

// 	url := cidaas_client.BaseUrl + "/templates-srv/template/custom/" + template_key + "/" + template_type
// 	method := "DELETE"

// 	fmt.Println(url)

// 	client := &http.Client{}
// 	req, err := http.NewRequest(method, url, nil)

// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	req.Header.Add("Authorization", "Bearer "+cidaas_client.TokenData.AccessToken)
// 	req.Header.Add("Content-Type", "application/json")
// 	req.Header.Add("Accept", "application/json, text/plain, */*")

// 	res, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer res.Body.Close()

// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	// fmt.Println(string(body))

// 	json.Unmarshal([]byte(body), &base_response)

// 	return base_response

// }
