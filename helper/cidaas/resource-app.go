package cidaas

import (
	"encoding/json"
	"fmt"
	"terraform-provider-cidaas/helper/util"
)

type AppResponse struct {
	Success bool      `json:"success,omitempty"`
	Status  int       `json:"status,omitempty"`
	Data    AppConfig `json:"data,omitempty"`
}

type AppConfig struct {
	ClientId                      string      `json:"client_id,omitempty"`
	ClientName                    string      `json:"client_name,omitempty"`
	ClientDisplayName             string      `json:"client_display_name,omitempty"`
	TemplateGroupId               string      `json:"template_group_id,omitempty"`
	HostedPageGroup               string      `json:"hosted_page_group,omitempty"`
	ClientType                    string      `json:"client_type,omitempty"`
	AllowLoginWith                []string    `json:"allow_login_with,omitempty"`
	AutoLoginAfterRegister        bool        `json:"auto_login_after_register,omitempty"`
	EnablePasswordlessAuth        bool        `json:"enable_passwordless_auth,omitempty"`
	RegisterWithLoginInformation  bool        `json:"register_with_login_information,omitempty"`
	EnableDeduplication           bool        `json:"enable_deduplication,omitempty"`
	AllowDisposableEmail          bool        `json:"allow_disposable_email,omitempty"`
	ValidatePhoneNumber           bool        `json:"validate_phone_number,omitempty"`
	FdsEnabled                    bool        `json:"fds_enabled,omitempty"`
	CompanyName                   string      `json:"company_name,omitempty"`
	CompanyAddress                string      `json:"company_address,omitempty"`
	CompanyWebsite                string      `json:"company_website,omitempty"`
	AllowedScopes                 []string    `json:"allowed_scopes,omitempty"`
	ResponseTypes                 []string    `json:"response_types,omitempty"`
	LoginProviders                []string    `json:"login_providers,omitempty"`
	AdditionalAccessTokenPayload  []string    `json:"additional_access_token_payload,omitempty"`
	GrantTypes                    []string    `json:"grant_types,omitempty"`
	RequiredFields                []string    `json:"required_fields,omitempty"`
	IsHybridApp                   bool        `json:"is_hybrid_app,omitempty"`
	RedirectURIS                  []string    `json:"redirect_uris,omitempty"`
	AllowedLogoutUrls             []string    `json:"allowed_logout_urls,omitempty"`
	AllowedWebOrigins             []string    `json:"allowed_web_origins,omitempty"`
	AllowedOrigins                []string    `json:"allowed_origins,omitempty"`
	MobileSettings                interface{} `json:"mobile_settings,omitempty"`
	AccentColor                   string      `json:"accent_color,omitempty"`
	PrimaryColor                  string      `json:"primary_color,omitempty"`
	MediaType                     string      `json:"media_type,omitempty"`
	ContentAlign                  string      `json:"contentAlign,omitempty"`
	ApplicationType               string      `json:"application_type,omitempty"`
	ApplicationMetaData           interface{} `json:"application_meta_data,omitempty"`
	RefreshTokenLifetimeInSeconds int         `json:"refresh_token_lifetime_in_seconds,omitempty"`
}

func (c *CidaasClient) CreateApp(app AppConfig) (response *AppResponse, err error) {
	h := util.HttpClient{
		Token: c.TokenData.AccessToken,
	}
	res, err := h.Post(c.AppUrl, app)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %v", err)
	}
	return response, nil
}

func (c *CidaasClient) UpdateApp(app AppConfig) (response *AppResponse, err error) {
	h := util.HttpClient{
		Token: c.TokenData.AccessToken,
	}
	res, err := h.Put(c.AppUrl, app)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %v", err)
	}
	return response, nil
}

func (c *CidaasClient) DeleteApp(app AppConfig) (response *AppResponse, err error) {
	h := util.HttpClient{
		Token: c.TokenData.AccessToken,
	}
	url := c.BaseUrl + "/apps-srv/clients/" + app.ClientId
	res, err := h.Delete(url)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %v", err)
	}
	return response, nil
}

func (c *CidaasClient) GetApp(app AppConfig) (response *AppResponse, err error) {
	h := util.HttpClient{
		Token: c.TokenData.AccessToken,
	}
	url := c.BaseUrl + "/apps-srv/clients/" + app.ClientId
	res, err := h.Get(url)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %v", err)
	}
	return response, nil
}
