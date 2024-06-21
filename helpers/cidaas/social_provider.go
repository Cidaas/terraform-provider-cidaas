package cidaas

import (
	"encoding/json"
	"fmt"
	"github.com/Cidaas/terraform-provider-cidaas/helpers/util"
	"net/http"
)

type SocialProviderModel struct {
	ID                    string      `json:"id,omitempty"`
	ClientID              string      `json:"client_id,omitempty"`
	ClientSecret          string      `json:"client_secret,omitempty"`
	Name                  string      `json:"name,omitempty"`
	ProviderName          string      `json:"provider_name,omitempty"`
	Claims                ClaimsModel `json:"claims,omitempty"`
	EnabledForAdminPortal bool        `json:"enabled_for_admin_portal"`
	Enabled               bool        `json:"enabled"`
	SPScopes              []string    `json:"scopes"`
	SPUserInfoFields      []string    `json:"userinfo_fields"`
}

type ClaimsModel struct {
	RequiredClaims RequiredClaimsModel `json:"required_claims,omitempty"`
	OptionalClaims OptionalClaimsModel `json:"optional_claims,omitempty"`
}

type RequiredClaimsModel struct {
	UserInfo []string `json:"user_info,omitempty"`
	IdToken  []string `json:"id_token,omitempty"`
}

type OptionalClaimsModel struct {
	UserInfo []string `json:"user_info,omitempty"`
	IdToken  []string `json:"id_token,omitempty"`
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
	UpsertSocialProvider(cp *SocialProviderModel) (*SocialProviderResponse, error)

	GetSocialProvider(providerName, providerId string) (*SocialProviderResponse, error)
	DeleteSocialProvider(providerName, providerId string) error
	//ConfigureCustomProvider(cp CustomProviderConfigPayload) (*CustomProviderConfigureResponse, error)
	//UpdateCustomProvider(cp *CustomProviderModel) error
	//https://kube-nightlybuild-dev.cidaas.de/providers-srv/multi/providers/google/bc34e3db-aeaf-4a41-9b75-c12b825ea987
}

func NewSocialProvider(httpClient util.HTTPClientInterface) SocialProviderService {
	return &SocialProvider{HTTPClient: httpClient}
}

func (c *SocialProvider) UpsertSocialProvider(cp *SocialProviderModel) (*SocialProviderResponse, error) {
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

func (c *SocialProvider) GetSocialProvider(providerName, providerId string) (*SocialProviderResponse, error) {
	c.HTTPClient.SetURL(fmt.Sprintf("%s/%s%s&provider_id=%s", c.HTTPClient.GetHost(), "providers-srv/multi/providers?provider_name=", providerName, providerId))
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

func (c *SocialProvider) DeleteSocialProvider(providerName, providerId string) error {
	c.HTTPClient.SetURL(fmt.Sprintf("%s/%s/%s/%s", c.HTTPClient.GetHost(), "providers-srv/multi/providers", providerName, providerId))
	c.HTTPClient.SetMethod(http.MethodDelete)
	res, err := c.HTTPClient.MakeRequest(nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}
