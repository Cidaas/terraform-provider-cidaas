package cidaas

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/util"
)

type SocialProviderModel struct {
	ID                    string                `json:"id,omitempty"`
	ClientID              string                `json:"client_id,omitempty"`
	ClientSecret          string                `json:"client_secret,omitempty"`
	Name                  string                `json:"name,omitempty"`
	ProviderName          string                `json:"provider_name,omitempty"`
	Claims                *ClaimsModel          `json:"claims,omitempty"`
	EnabledForAdminPortal bool                  `json:"enabled_for_admin_portal"`
	Enabled               bool                  `json:"enabled"`
	Scopes                []string              `json:"scopes"`
	UserInfoFields        []UserInfoFieldsModel `json:"userinfo_fields"`
}

type ClaimsModel struct {
	RequiredClaims RequiredClaimsModel `json:"required_claims,omitempty"`
	OptionalClaims OptionalClaimsModel `json:"optional_claims,omitempty"`
}

type RequiredClaimsModel struct {
	UserInfo []string `json:"user_info,omitempty"`
	IDToken  []string `json:"id_token,omitempty"`
}

type OptionalClaimsModel struct {
	UserInfo []string `json:"user_info,omitempty"`
	IDToken  []string `json:"id_token,omitempty"`
}

type UserInfoFieldsModel struct {
	InnerKey      string `json:"inner_key,omitempty"`
	ExternalKey   string `json:"external_key,omitempty"`
	IsCustomField bool   `json:"is_custom_field,omitempty"`
	IsSystemField bool   `json:"is_system_field,omitempty"`
}

type SocialProviderResponse struct {
	Success bool `json:"success,omitempty"`
	Status  int  `json:"status,omitempty"`
	Data    SocialProviderModel
}

type SocialProvider struct {
	HTTPClient util.HTTPClientInterface
}

type SocialProviderService interface {
	Upsert(cp *SocialProviderModel) (*SocialProviderResponse, error)
	Get(providerName, providerID string) (*SocialProviderResponse, error)
	Delete(providerName, providerID string) error
}

func NewSocialProvider(httpClient util.HTTPClientInterface) SocialProviderService {
	return &SocialProvider{HTTPClient: httpClient}
}

func (c *SocialProvider) Upsert(cp *SocialProviderModel) (*SocialProviderResponse, error) {
	c.HTTPClient.SetURL(fmt.Sprintf("%s/%s", c.HTTPClient.GetHost(), "providers-srv/multi/providers"))
	c.HTTPClient.SetMethod(http.MethodPost)
	res, err := c.HTTPClient.MakeRequest(cp)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var response SocialProviderResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %w", err)
	}
	return &response, nil
}

func (c *SocialProvider) Get(providerName, providerID string) (*SocialProviderResponse, error) {
	c.HTTPClient.SetURL(fmt.Sprintf("%s/%s?provider_name=%s&provider_id=%s", c.HTTPClient.GetHost(), "providers-srv/multi/providers", providerName, providerID))
	c.HTTPClient.SetMethod(http.MethodGet)
	res, err := c.HTTPClient.MakeRequest(nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var response SocialProviderResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %w", err)
	}
	return &response, nil
}

func (c *SocialProvider) Delete(providerName, providerID string) error {
	c.HTTPClient.SetURL(fmt.Sprintf("%s/%s/%s/%s", c.HTTPClient.GetHost(), "providers-srv/multi/providers", providerName, providerID))
	c.HTTPClient.SetMethod(http.MethodDelete)
	res, err := c.HTTPClient.MakeRequest(nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}
