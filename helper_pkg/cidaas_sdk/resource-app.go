package cidaas_sdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Error struct {
	Code   int    `json:"code"`
	Type   string `json:"type"`
	Status int    `json:"status"`
	Error  string `json:"error"`
}

type baseResponse struct {
	Success bool      `json:"success"`
	Status  int       `json:"status"`
	Data    AppConfig `json:"data"`
	Error   Error     `json:"error"`
}

type AppConfig struct {
	ClientId                     string      `json:"client_id"`
	ClientName                   string      `json:"client_name"`
	ClientDisplayName            string      `json:"client_display_name"`
	TemplateGroupId              string      `json:"template_group_id"`
	HostedPageGroup              string      `json:"hosted_page_group"`
	ClientType                   string      `json:"client_type"`
	AllowLoginWith               []string    `json:"allow_login_with"`
	AutoLoginAfterRegister       bool        `json:"auto_login_after_register"`
	EnablePasswordlessAuth       bool        `json:"enable_passwordless_auth"`
	RegisterWithLoginInformation bool        `json:"register_with_login_information"`
	EnableDeduplication          bool        `json:"enable_deduplication"`
	AllowDisposableEmail         bool        `json:"allow_disposable_email"`
	ValidatePhoneNumber          bool        `json:"validate_phone_number"`
	FdsEnabled                   bool        `json:"fds_enabled"`
	CompanyName                  string      `json:"company_name"`
	CompanyAddress               string      `json:"company_address"`
	CompanyWebsite               string      `json:"company_website"`
	AllowedScopes                []string    `json:"allowed_scopes"`
	ResponseTypes                []string    `json:"response_types"`
	LoginProviders               []string    `json:"login_providers"`
	AdditionalAccessTokenPayload string      `json:"additional_access_token_payload"`
	GrantTypes                   []string    `json:"grant_types"`
	RequiredFields               string      `json:"required_fields"`
	ApplicationMetadata          string      `json:"application_metadata"`
	IsHybridApp                  bool        `json:"is_hybrid_app"`
	RedirectURIS                 []string    `json:"redirect_uris"`
	AllowedLogoutUrls            []string    `json:"allowed_logout_urls"`
	AllowedWebOrigins            []string    `json:"allowed_web_origins"`
	AllowedOrigins               []string    `json:"allowed_origins"`
	MobileSettings               interface{} `json:"mobile_settings"`
}

func CreateApp(cidaas_client CidaasClient, app_config AppConfig) (base_response baseResponse) {

	url := cidaas_client.AppUrl
	method := "POST"

	json_payload, err := json.Marshal(app_config)
	if err != nil {
		// fmt.Println(err)
		log.Printf("Error: %s", err.Error())

		log.Printf("HERE 1")
		return
	}
	
	payload_string := string(json_payload)
	payload := strings.NewReader(payload_string)


	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Printf("Error: %s", err.Error())

		log.Printf("HERE 2")

		return
	}

	req.Header.Add("Authorization", "Bearer "+cidaas_client.TokenData.AccessToken)
	// req.Header.Add("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6ImJjODQ3M2Q0LWViNWItNDUyNC1hODU5LTA4ZjdlODUzZTZhOCJ9.eyJhbXIiOlsiMTAzIiwiMTAxIl0sInVhX2hhc2giOiI1MjdmZjZiNDg5YjI2N2UyNjJiZDhhZmZjMTgyZTRjOSIsInNpZCI6Ijc1ODU4MmZlLTc3ZmQtNGY5Yi1iYWQ2LTlmZmViZTQyYjZmMCIsInN1YiI6IjVkM2ExZjY2LTQ1MDQtNDEzOC1hZjg4LWQzODQwZDNjYWEzZCIsImlzdWIiOiI2NGYxMmMwYS1kYzBiLTRlNzktYWFiNy04ZTAyNGRjMjJhZmIiLCJhdWQiOiI3ZjVjOGFmZC1lNjEzLTQzOTMtODM0Yi1kMTc1YjM0NWYzNTEiLCJpYXQiOjE2NTk5NDIxODMsImF1dGhfdGltZSI6MTY1OTk0MjE3OCwiaXNzIjoiaHR0cHM6Ly9pZHAtZGV2LnN0YWNraXQuY2xvdWQiLCJqdGkiOiIwYmMyODQxNS0yNzEzLTQ5OGUtOTBiNC02YTYyODMzMGU4MGUiLCJzY29wZXMiOlsib3BlbmlkIiwicHJvZmlsZSIsImNpZGFhczphcHBzX2RlbGV0ZSIsImNpZGFhczphcHBzX3dyaXRlIiwiY2lkYWFzOmFwcHNfcmVhZCIsImNpZGFhczpob3N0ZWRfcGFnZXNfcmVhZCIsImNpZGFhczpob3N0ZWRfcGFnZXNfZGVsZXRlIiwiY2lkYWFzOmhvc3RlZF9wYWdlc193cml0ZSIsImNpZGFhczphZG1pbl9kZWxldGUiLCJjaWRhYXM6YWRtaW5fd3JpdGUiLCJjaWRhYXM6YWRtaW5fcmVhZCIsImNpZGFhczp1c2VyYWN0aXZpdHlfcmVhZCIsIm9mZmxpbmVfYWNjZXNzIl0sInJvbGVzIjpbIlVTRVIiXSwiZ3JvdXBzIjpbeyJncm91cElkIjoiQ0lEQUFTX0FETUlOUyIsInJvbGVzIjpbIkFQUF9SRUFEIiwiQVBQX0NSRUFURSIsIkFQUF9ERUxFVEUiXX1dLCJleHAiOjE2NjAwMjg1ODN9.Vq2X5qGwX85WT2_ouOIA0rF_KwnKAhH42j1ObUXcAsXdslXalfV_yImtwIKWu_nxY2OggzHywo8eSxTCJ5oFS8IclY-cyT8XGB4BG5QGjvThcRwUad_-oVJS_jvGsl6iJMYhqkDSfAuTu-WkZaY18shMNZd4-b9e0Y5JjkXkb1yacVEiOlC2h02VLC88aufUAK65P-LWHDHnEjcrzWKSDouuJKRhcYI3JVjZtS9eFG0gmw3aFhFxxZXIWf_QBcocN9nu4GDcB4aU6NtII2IFZbF-iSm7hLy_xwWJ6-Qlc5s2f5zydppgfB0on927DR-MtRs4sy77v1QTaNCwVnapnJG9rhjPofQwzKMyLaKLEPp4x0lulMXB4T9ttqwTz5fHPgUXHoRWROytAdLe63FtSmRcyqlg2b60Xp7nJPxDrel-1-6Q2vGTH7in5ovEGxXOElcYiWU0I-Go8cF3IJeL-fUwOE6KIqeLBBxGivTAyZjPKNgZNDdj4zcB5vhxtyeU8jYLff0zniIN6luuq8ln8RU5qeHnusuadn2iXVh5TBsC-C4AgZHPgx_1TBnMuaDQ636XYjQ1RDqoQFQEXA-3s_w5sqJAM-NuY1oljY4s0QwKumlmhHHoBqP3ITJPGVrqjvOe3CG7xThRTN1AL3tSwdMA7iiLXCMMc_I_GJsfy-8")


	// log.Println("CIDAAS CLIENT TOKEN  ", cidaas_client.TokenData.AccessToken)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json, text/plain, */*")

	res, err := client.Do(req)
	if err != nil {
		// fmt.Println(err)
		log.Printf("Error: %s", err.Error())
		
		log.Printf("HERE 3")
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// fmt.Println(err)
		log.Printf("Error: %s", err.Error())
	
		log.Printf("HERE 4")

		return
	}
	log.Println(string(body))


	if err := json.Unmarshal([]byte(body), &base_response); err != nil {
		log.Printf("JSON UNMARSHAL")
		log.Printf("Error: %s", err.Error())
		return
	 }

	
	

	return base_response
}

func UpdateApp(cidaas_client CidaasClient, app_config AppConfig) (base_response baseResponse) {

	url := cidaas_client.AppUrl
	method := "PUT"

	json_payload, err := json.Marshal(app_config)
	if err != nil {
		fmt.Println(err)
		return
	}
	payload_string := string(json_payload)
	payload := strings.NewReader(payload_string)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		// fmt.Println(err)
		log.Printf("Error: %s", err.Error())

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

func operate_app(cidaas_client CidaasClient, client_id string, operation_type string) (base_response baseResponse) {
	//url := "https://terraform-cidaas-test-free.cidaas.de/apps-srv/clients/" + client_id
	//url := "https://terraform-cidaas-test-free.cidaas.de/apps-srv/clients/" + client_id
	url := cidaas_client.BaseUrl + "/apps-srv/clients/" + client_id

	method := "GET"

	if operation_type == "delete_app" {
		method = "DELETE"
	} else if operation_type == "get_app" {
		method = "GET"
	} else {
		method = "GET"
	}

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
	// fmt.Println(string(body))

	json.Unmarshal([]byte(body), &base_response)

	return base_response
}

func DeleteApp(cidaas_client CidaasClient, client_id string) (base_response baseResponse) {
	base_response = operate_app(cidaas_client, client_id, "delete_app")

	return base_response
}

func GetApp(cidaas_client CidaasClient, client_id string) (base_response baseResponse) {
	base_response = operate_app(cidaas_client, client_id, "get_app")

	return base_response
}
