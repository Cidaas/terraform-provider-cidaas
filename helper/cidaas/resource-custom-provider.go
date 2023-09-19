package cidaas

import (
	"encoding/json"
	"fmt"
	"strings"
	"terraform-provider-cidaas/helper/util"
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
type CustomProviderResponse struct {
	Success bool           `json:"success,omitempty"`
	Status  int            `json:"status,omitempty"`
	Data    CustomProvider `json:"data,omitempty"`
}

type CustomProviderConfigPayload struct {
	ClientId    string `json:"client_id,omitempty"`
	Test        bool   `json:"deleted"`
	Type        string `json:"type,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
}

type CustomProviderConfigureResponse struct {
	Success bool `json:"success,omitempty"`
	Status  int  `json:"status,omitempty"`
	Data    struct {
		Updated bool `json:"updated,omitempty"`
	} `json:"data,omitempty"`
}

func (c *CidaasClient) CreateCustomProvider(cp *CustomProvider) (response *CustomProviderResponse, err error) {
	h := util.HttpClient{
		Token: c.TokenData.AccessToken,
	}
	res, err := h.Post(c.ProvideUrl, cp)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %v", err)
	}
	return response, nil
}

func (c *CidaasClient) UpdateCustomProvider(body *CustomProvider) (response *CustomProviderResponse, err error) {
	h := util.HttpClient{
		Token: c.TokenData.AccessToken,
	}
	res, err := h.Put(c.ProvideUrl, body)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %v", err)
	}
	return response, nil
}

func (c *CidaasClient) GetCustomProvider(provider_name string) (response *CustomProviderResponse, err error) {
	url := c.BaseUrl + "/providers-srv/custom/" + provider_name
	h := util.HttpClient{
		Token: c.TokenData.AccessToken,
	}
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

func (c *CidaasClient) DeleteCustomProvider(provider string) (response *CustomProviderResponse, err error) {
	url := c.BaseUrl + "/providers-srv/custom/" + strings.ToLower(provider)
	h := util.HttpClient{
		Token: c.TokenData.AccessToken,
	}
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

func (c *CidaasClient) ConfigureCustomProvider(cp CustomProviderConfigPayload) (response *CustomProviderConfigureResponse, err error) {
	url := c.BaseUrl + "/apps-srv/loginproviders/update/" + cp.DisplayName
	h := util.HttpClient{
		Token: c.TokenData.AccessToken,
	}
	res, err := h.Put(url, cp)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %v", err)
	}
	return response, nil
}
