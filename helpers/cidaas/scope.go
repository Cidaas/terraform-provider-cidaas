package cidaas

import (
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
	ClientConfig
}
type ScopeService interface {
	Upsert(sc ScopeModel) (*ScopeResponse, error)
	Get(scopeKey string) (*ScopeResponse, error)
	Delete(scopeKey string) error
}

func NewScope(clientConfig ClientConfig) ScopeService {
	return &ScopeImpl{clientConfig}
}

func (c *ScopeImpl) Upsert(sc ScopeModel) (*ScopeResponse, error) {
	var response ScopeResponse
	url := fmt.Sprintf("%s/%s", c.BaseURL, "scopes-srv/scope")
	httpClient := util.NewHTTPClient(url, http.MethodPost, c.AccessToken)

	res, err := httpClient.MakeRequest(sc)
	if err = util.HandleResponseError(res, err); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err = util.ProcessResponse(res, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *ScopeImpl) Get(scopeKey string) (*ScopeResponse, error) {
	var response ScopeResponse
	url := fmt.Sprintf("%s/%s?scopekey=%s", c.BaseURL, "scopes-srv/scope", strings.ToLower(scopeKey))
	httpClient := util.NewHTTPClient(url, http.MethodGet, c.AccessToken)

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

func (c *ScopeImpl) Delete(scopeKey string) error {
	url := fmt.Sprintf("%s/%s/%s", c.BaseURL, "scopes-srv/scope", strings.ToLower(scopeKey))
	httpClient := util.NewHTTPClient(url, http.MethodDelete, c.AccessToken)

	res, err := httpClient.MakeRequest(nil)
	if err = util.HandleResponseError(res, err); err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}
