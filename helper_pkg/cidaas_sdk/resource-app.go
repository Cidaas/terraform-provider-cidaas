package cidaas_sdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Error struct {
	Code   int    `json:"code,omitempty"`
	Type   string `json:"type,omitempty"`
	Status int    `json:"status,omitempty"`
	Error  string `json:"error,omitempty"`
}

type DeleteError struct {
	Code   int    `json:"code,omitempty"`
	Type   string `json:"type,omitempty"`
	Status int    `json:"status,omitempty"`
	Error  string `json:"error,omitempty"`
}
type baseResponse struct {
	Success bool      `json:"success,omitempty"`
	Status  int       `json:"status,omitempty"`
	Data    AppConfig `json:"data,omitempty"`
	Errors  Error     `json:"error,omitempty"`
	// Error   string    `json:"error,omitempty"`
	// RefNum  string    `json:"renum,omitempty"`
}

type AppConfig struct {
	ClientId                     string      `json:"client_id,omitempty"`
	ClientName                   string      `json:"client_name,omitempty"`
	ClientDisplayName            string      `json:"client_display_name,omitempty"`
	TemplateGroupId              string      `json:"template_group_id,omitempty"`
	HostedPageGroup              string      `json:"hosted_page_group,omitempty"`
	ClientType                   string      `json:"client_type,omitempty"`
	AllowLoginWith               []string    `json:"allow_login_with,omitempty"`
	AutoLoginAfterRegister       bool        `json:"auto_login_after_register,omitempty"`
	EnablePasswordlessAuth       bool        `json:"enable_passwordless_auth,omitempty"`
	RegisterWithLoginInformation bool        `json:"register_with_login_information,omitempty"`
	EnableDeduplication          bool        `json:"enable_deduplication,omitempty"`
	AllowDisposableEmail         bool        `json:"allow_disposable_email,omitempty"`
	ValidatePhoneNumber          bool        `json:"validate_phone_number,omitempty"`
	FdsEnabled                   bool        `json:"fds_enabled,omitempty"`
	CompanyName                  string      `json:"company_name,omitempty"`
	CompanyAddress               string      `json:"company_address,omitempty"`
	CompanyWebsite               string      `json:"company_website,omitempty"`
	AllowedScopes                []string    `json:"allowed_scopes,omitempty"`
	ResponseTypes                []string    `json:"response_types,omitempty"`
	LoginProviders               []string    `json:"login_providers,omitempty"`
	AdditionalAccessTokenPayload string      `json:"additional_access_token_payload,omitempty"`
	GrantTypes                   []string    `json:"grant_types,omitempty"`
	RequiredFields               []string    `json:"required_fields,omitempty"`
	IsHybridApp                  bool        `json:"is_hybrid_app,omitempty"`
	RedirectURIS                 []string    `json:"redirect_uris,omitempty"`
	AllowedLogoutUrls            []string    `json:"allowed_logout_urls,omitempty"`
	AllowedWebOrigins            []string    `json:"allowed_web_origins,omitempty"`
	AllowedOrigins               []string    `json:"allowed_origins,omitempty"`
	MobileSettings               interface{} `json:"mobile_settings,omitempty"`

	AccentColor                   string      `json:"accent_color,omitempty"`
	PrimaryColor                  string      `json:"primary_color,omitempty"`
	MediaType                     string      `json:"media_type,omitempty"`
	ContentAlign                  string      `json:"contentAlign,omitempty"`
	ApplicationType               string      `json:"application_type,omitempty"`
	ApplicationMetaData           interface{} `json:"application_meta_data,omitempty"`
	RefreshTokenLifetimeInSeconds int         `json:"refresh_token_lifetime_in_seconds,omitempty"`
}

func CreateApp(cidaas_client CidaasClient, app_config AppConfig) (base_response baseResponse) {

	url := cidaas_client.AppUrl
	method := "POST"

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

func operate_app(cidaas_client CidaasClient, client_id string, operation_type string) (base_response baseResponse) {
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

func preparePayload(d *schema.ResourceData) AppConfig {
	var appConfig AppConfig
	appConfig.ClientType = d.Get("client_type").(string)
	appConfig.AccentColor = d.Get("accent_color").(string)
	appConfig.PrimaryColor = d.Get("primary_color").(string)
	appConfig.MediaType = d.Get("media_type").(string)
	appConfig.ContentAlign = d.Get("content_align").(string)
	appConfig.AllowLoginWith = interfaceArray2StringArray(d.Get("allow_login_with").([]interface{}))
	appConfig.RedirectURIS = interfaceArray2StringArray(d.Get("redirect_uris").([]interface{}))
	appConfig.AllowedLogoutUrls = interfaceArray2StringArray(d.Get("allowed_logout_urls").([]interface{}))
	appConfig.EnableDeduplication = d.Get("enable_deduplication").(bool)
	appConfig.AutoLoginAfterRegister = d.Get("auto_login_after_register").(bool)
	appConfig.EnablePasswordlessAuth = d.Get("enable_passwordless_auth").(bool)
	appConfig.RegisterWithLoginInformation = d.Get("register_with_login_information").(bool)
	appConfig.AllowDisposableEmail = d.Get("allow_disposable_email").(bool)
	appConfig.ValidatePhoneNumber = d.Get("validate_phone_number").(bool)
	appConfig.FdsEnabled = d.Get("fds_enabled").(bool)
	appConfig.HostedPageGroup = d.Get("hosted_page_group").(string)
	appConfig.ClientName = d.Get("client_name").(string)
	appConfig.ClientDisplayName = d.Get("client_display_name").(string)
	appConfig.CompanyName = d.Get("company_name").(string)
	appConfig.CompanyAddress = d.Get("company_address").(string)
	appConfig.CompanyWebsite = d.Get("company_website").(string)
	appConfig.AllowedScopes = interfaceArray2StringArray(d.Get("allowed_scopes").([]interface{}))
	appConfig.ResponseTypes = interfaceArray2StringArray(d.Get("response_types").([]interface{}))
	appConfig.GrantTypes = interfaceArray2StringArray(d.Get("grant_types").([]interface{}))
	appConfig.LoginProviders = interfaceArray2StringArray(d.Get("login_providers").([]interface{}))
	appConfig.AdditionalAccessTokenPayload = d.Get("additional_access_token_payload").(string)
	appConfig.RequiredFields = interfaceArray2StringArray(d.Get("required_fields").([]interface{}))
	appConfig.ApplicationMetaData = interfaceArray2StringArray(d.Get("application_meta_data").([]interface{}))
	appConfig.IsHybridApp = d.Get("is_hybrid_app").(bool)
	appConfig.AllowedWebOrigins = interfaceArray2StringArray(d.Get("allowed_web_origins").([]interface{}))
	appConfig.AllowedOrigins = interfaceArray2StringArray(d.Get("allowed_origins").([]interface{}))
	appConfig.MobileSettings = d.Get("mobile_settings").(string)
	appConfig.RefreshTokenLifetimeInSeconds = d.Get("refresh_token_lifetime_in_seconds").(int)
	appConfig.TemplateGroupId = d.Get("template_group_id").(string)

	return appConfig
}

func interfaceArray2StringArray(interfaceArray []interface{}) (stringArray []string) {

	stringArray = make([]string, 0)
	for _, txt := range interfaceArray {
		stringArray = append(stringArray, txt.(string))
	}

	return stringArray
}
