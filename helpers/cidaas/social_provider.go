package cidaas

import (
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
	ClientConfig
}

type SocialProviderService interface {
	Upsert(cp *SocialProviderModel) (*SocialProviderResponse, error)
	Get(providerName, providerID string) (*SocialProviderResponse, error)
	Delete(providerName, providerID string) error
}

func NewSocialProvider(clientConfig ClientConfig) SocialProviderService {
	return &SocialProvider{clientConfig}
}

func (s *SocialProvider) Upsert(sp *SocialProviderModel) (*SocialProviderResponse, error) {
	var response SocialProviderResponse
	url := fmt.Sprintf("%s/%s", s.BaseURL, "providers-srv/multi/providers")
	httpClient := util.NewHTTPClient(url, http.MethodPost, s.AccessToken)

	res, err := httpClient.MakeRequest(sp)
	if err = util.HandleResponseError(res, err); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err = util.ProcessResponse(res, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (s *SocialProvider) Get(providerName, providerID string) (*SocialProviderResponse, error) {
	var response SocialProviderResponse
	url := fmt.Sprintf("%s/%s?provider_name=%s&provider_id=%s", s.BaseURL, "providers-srv/multi/providers", providerName, providerID)
	httpClient := util.NewHTTPClient(url, http.MethodGet, s.AccessToken)

	res, err := httpClient.MakeRequest(nil)
	if err = util.HandleResponseError(res, err); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err = util.ProcessResponse(res, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (s *SocialProvider) Delete(providerName, providerID string) error {
	url := fmt.Sprintf("%s/%s/%s/%s", s.BaseURL, "providers-srv/multi/providers", providerName, providerID)
	httpClient := util.NewHTTPClient(url, http.MethodDelete, s.AccessToken)

	res, err := httpClient.MakeRequest(nil)
	if err = util.HandleResponseError(res, err); err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}
