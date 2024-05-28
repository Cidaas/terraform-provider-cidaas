package cidaas

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/util"
)

type ScopeModel struct {
	ID                    string                  `json:"_id,omitempty"`
	LocaleWiseDescription []ScopeLocalDescription `json:"localeWiseDescription,omitempty"`
	SecurityLevel         string                  `json:"securityLevel,omitempty"`
	ScopeKey              string                  `json:"scopeKey,omitempty"`
	RequiredUserConsent   bool                    `json:"requiredUserConsent"`
	GroupName             []string                `json:"group_name,omitempty"`
	ScopeOwner            string                  `json:"scopeOwner,omitempty"`
}

type ScopeLocalDescription struct {
	Locale      string `json:"locale,omitempty"`
	Language    string `json:"language,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

type ScopeResponse struct {
	Success bool       `json:"success,omitempty"`
	Status  int        `json:"status,omitempty"`
	Data    ScopeModel `json:"data,omitempty"`
	Error   string     `json:"error,omitempty"`
}

type ScopeImpl struct {
	HTTPClient util.HTTPClientInterface
}
type ScopeService interface {
	Upsert(sc ScopeModel) (*ScopeResponse, error)
	Get(scopeKey string) (*ScopeResponse, error)
	Delete(scopeKey string) error
}

func NewScope(httpClient util.HTTPClientInterface) ScopeService {
	return &ScopeImpl{HTTPClient: httpClient}
}

func (c *ScopeImpl) Upsert(sc ScopeModel) (*ScopeResponse, error) {
	c.HTTPClient.SetURL(fmt.Sprintf("%s/%s", c.HTTPClient.GetHost(), "scopes-srv/scope"))
	c.HTTPClient.SetMethod(http.MethodPost)
	res, err := c.HTTPClient.MakeRequest(sc)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var response ScopeResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %w", err)
	}
	return &response, nil
}

func (c *ScopeImpl) Get(scopeKey string) (*ScopeResponse, error) {
	c.HTTPClient.SetURL(fmt.Sprintf("%s/%s?scopekey=%s", c.HTTPClient.GetHost(), "scopes-srv/scope", strings.ToLower(scopeKey)))
	c.HTTPClient.SetMethod(http.MethodGet)
	res, err := c.HTTPClient.MakeRequest(nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var response ScopeResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %w", err)
	}
	return &response, nil
}

func (c *ScopeImpl) Delete(scopeKey string) error {
	c.HTTPClient.SetURL(fmt.Sprintf("%s/%s/%s", c.HTTPClient.GetHost(), "scopes-srv/scope", strings.ToLower(scopeKey)))
	c.HTTPClient.SetMethod(http.MethodDelete)
	res, err := c.HTTPClient.MakeRequest(nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}
