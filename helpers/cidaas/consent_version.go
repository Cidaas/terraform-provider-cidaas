package cidaas

import (
	"fmt"
	"net/http"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/util"
)

type ConsentVersionResponse struct {
	Success bool                `json:"success,omitempty"`
	Status  int                 `json:"status,omitempty"`
	Data    ConsentVersionModel `json:"data,omitempty"`
}

type ConsentVersionReadResponse struct {
	Success bool                  `json:"success,omitempty"`
	Status  int                   `json:"status,omitempty"`
	Data    []ConsentVersionModel `json:"data,omitempty"`
}

type ConsentVersionModel struct {
	ID             string        `json:"_id,omitempty"`
	Version        float64       `json:"version,omitempty"`
	ConsentID      string        `json:"consent_id,omitempty"`
	ConsentType    string        `json:"consentType,omitempty"`
	Scopes         []string      `json:"scopes,omitempty"`
	RequiredFields []string      `json:"required_fields,omitempty"`
	ConsentLocale  ConsentLocale `json:"consent_locale,omitempty"`
	CreatedAt      string        `json:"created_at,omitempty"`
	UpdatedAt      string        `json:"updated_at,omitempty"`
}

type ConsentLocalResponse struct {
	Success bool              `json:"success,omitempty"`
	Status  int               `json:"status,omitempty"`
	Data    ConsentLocalModel `json:"data,omitempty"`
}

type ConsentLocalModel struct {
	ConsentVersionID string   `json:"consent_version_id,omitempty"`
	ConsentID        string   `json:"consent_id,omitempty"`
	Content          string   `json:"content,omitempty"`
	Locale           string   `json:"locale,omitempty"`
	URL              string   `json:"url,omitempty"`
	Scopes           []string `json:"scopes,omitempty"`
	RequiredFields   []string `json:"required_fields,omitempty"`
}
type ConsentLocale struct {
	Locale  string `json:"locale,omitempty"`
	Content string `json:"content,omitempty"`
	URL     string `json:"url,omitempty"`
}

type ConsentVersion struct {
	ClientConfig
}

type ConsentVersionService interface {
	Upsert(consent ConsentVersionModel) (*ConsentVersionResponse, error)
	Get(consentID string) (*ConsentVersionReadResponse, error)
	UpsertLocal(consentLocal ConsentLocalModel) (*ConsentLocalResponse, error)
	GetLocal(consentVersionID string, locale string) (*ConsentLocalResponse, error)
}

func NewConsentVersion(clientConfig ClientConfig) ConsentVersionService {
	return &ConsentVersion{clientConfig}
}

func (c *ConsentVersion) Upsert(consentVersionConfig ConsentVersionModel) (*ConsentVersionResponse, error) {
	var response ConsentVersionResponse
	url := fmt.Sprintf("%s/%s", c.BaseURL, "consent-management-srv/v2/consent/versions")
	httpClient := util.NewHTTPClient(url, http.MethodPost, c.AccessToken)

	res, err := httpClient.MakeRequest(consentVersionConfig)
	if err = util.HandleResponseError(res, err); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err = util.ProcessResponse(res, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *ConsentVersion) Get(consentID string) (*ConsentVersionReadResponse, error) {
	var response ConsentVersionReadResponse
	url := fmt.Sprintf("%s/%s/%s", c.BaseURL, "consent-management-srv/v2/consent/versions/list", consentID)
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

func (c *ConsentVersion) UpsertLocal(consentLocal ConsentLocalModel) (*ConsentLocalResponse, error) {
	var response ConsentLocalResponse
	url := fmt.Sprintf("%s/%s", c.BaseURL, "consent-management-srv/v2/consent/locale")
	httpClient := util.NewHTTPClient(url, http.MethodPost, c.AccessToken)

	res, err := httpClient.MakeRequest(consentLocal)
	if err = util.HandleResponseError(res, err); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err = util.ProcessResponse(res, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *ConsentVersion) GetLocal(consentVersionID string, locale string) (*ConsentLocalResponse, error) {
	var response ConsentLocalResponse
	url := fmt.Sprintf("%s/%s/%s?locale=%s", c.BaseURL, "consent-management-srv/v2/consent/locale", consentVersionID, locale)
	httpClient := util.NewHTTPClient(url, http.MethodGet, c.AccessToken)

	res, err := httpClient.MakeRequest(nil)
	if res.StatusCode == http.StatusNoContent {
		return &ConsentLocalResponse{
			Success: false,
			Status:  http.StatusNoContent,
		}, nil
	}
	if err = util.HandleResponseError(res, err); err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err = util.ProcessResponse(res, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
