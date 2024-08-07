package cidaas

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/util"
)

type Scopes struct {
	DisplayLabel string       `json:"display_label,omitempty"`
	Scopes       []ScopeChild `json:"scopes,omitempty"`
}

type ScopeChild struct {
	ScopeName   string `json:"scope_name,omitempty"`
	Required    bool   `json:"required,omitempty"`
	Recommended bool   `json:"recommended,omitempty"`
}
type CustomProviderModel struct {
	ID                    string            `json:"_id,omitempty"`
	ClientID              string            `json:"client_id,omitempty"`
	ClientSecret          string            `json:"client_secret,omitempty"`
	DisplayName           string            `json:"display_name,omitempty"`
	StandardType          string            `json:"standard_type,omitempty"`
	AuthorizationEndpoint string            `json:"authorization_endpoint,omitempty"`
	TokenEndpoint         string            `json:"token_endpoint,omitempty"`
	ProviderName          string            `json:"provider_name,omitempty"`
	LogoURL               string            `json:"logo_url,omitempty"`
	UserinfoEndpoint      string            `json:"userinfo_endpoint,omitempty"`
	UserinfoFields        map[string]string `json:"userinfo_fields,omitempty"`
	Scopes                Scopes            `json:"scopes,omitempty"`
	Domains               []string          `json:"domains,omitempty"`
}

type CustomProviderResponse struct {
	Success bool                `json:"success,omitempty"`
	Status  int                 `json:"status,omitempty"`
	Data    CustomProviderModel `json:"data,omitempty"`
}

type CustomProviderConfigPayload struct {
	ClientID    string `json:"client_id,omitempty"`
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

type CustomProvider struct {
	HTTPClient util.HTTPClientInterface
}
type CustomProvideService interface {
	CreateCustomProvider(cp *CustomProviderModel) (*CustomProviderResponse, error)
	UpdateCustomProvider(cp *CustomProviderModel) error
	GetCustomProvider(providerName string) (*CustomProviderResponse, error)
	DeleteCustomProvider(providerName string) error
	ConfigureCustomProvider(cp CustomProviderConfigPayload) (*CustomProviderConfigureResponse, error)
}

func NewCustomProvider(httpClient util.HTTPClientInterface) CustomProvideService {
	return &CustomProvider{HTTPClient: httpClient}
}

func (c *CustomProvider) CreateCustomProvider(cp *CustomProviderModel) (*CustomProviderResponse, error) {
	c.HTTPClient.SetURL(fmt.Sprintf("%s/%s", c.HTTPClient.GetHost(), "providers-srv/custom"))
	c.HTTPClient.SetMethod(http.MethodPost)
	res, err := c.HTTPClient.MakeRequest(cp)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var response CustomProviderResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %w", err)
	}
	return &response, nil
}

func (c *CustomProvider) UpdateCustomProvider(cp *CustomProviderModel) error {
	c.HTTPClient.SetURL(fmt.Sprintf("%s/%s", c.HTTPClient.GetHost(), "providers-srv/custom"))
	c.HTTPClient.SetMethod(http.MethodPut)
	res, err := c.HTTPClient.MakeRequest(cp)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}

func (c *CustomProvider) GetCustomProvider(providerName string) (*CustomProviderResponse, error) {
	c.HTTPClient.SetURL(fmt.Sprintf("%s/%s/%s", c.HTTPClient.GetHost(), "providers-srv/custom", providerName))
	c.HTTPClient.SetMethod(http.MethodGet)
	res, err := c.HTTPClient.MakeRequest(nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var response CustomProviderResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %w", err)
	}
	return &response, nil
}

func (c *CustomProvider) DeleteCustomProvider(providerName string) error {
	c.HTTPClient.SetURL(fmt.Sprintf("%s/%s/%s", c.HTTPClient.GetHost(), "providers-srv/custom", strings.ToLower(providerName)))
	c.HTTPClient.SetMethod(http.MethodDelete)
	res, err := c.HTTPClient.MakeRequest(nil)
	if err != nil {
		return nil
	}
	defer res.Body.Close()
	return nil
}

func (c *CustomProvider) ConfigureCustomProvider(cp CustomProviderConfigPayload) (*CustomProviderConfigureResponse, error) {
	c.HTTPClient.SetURL(fmt.Sprintf("%s/%s/%s", c.HTTPClient.GetHost(), "apps-srv/loginproviders/update", cp.DisplayName))
	c.HTTPClient.SetMethod(http.MethodPut)
	res, err := c.HTTPClient.MakeRequest(cp)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var response CustomProviderConfigureResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %w", err)
	}
	return &response, nil
}
