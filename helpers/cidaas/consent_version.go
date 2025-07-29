package cidaas

import (
	"context"
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

func NewConsentVersion(clientConfig ClientConfig) *ConsentVersion {
	return &ConsentVersion{clientConfig}
}

func (c *ConsentVersion) Upsert(ctx context.Context, consentVersionConfig ConsentVersionModel) (*ConsentVersionResponse, error) {
	var response ConsentVersionResponse
	url := fmt.Sprintf("%s/%s", c.BaseURL, "consent-management-srv/v2/consent/versions")
	client, err := util.NewHTTPClient(url, http.MethodPost, c.AccessToken)
	if err != nil {
		return nil, err
	}
	res, err := client.MakeRequest(ctx, consentVersionConfig)
	if err = util.HandleResponseError(res, err); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err = util.ProcessResponse(res, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *ConsentVersion) Get(ctx context.Context, consentID string) (*ConsentVersionReadResponse, error) {
	var response ConsentVersionReadResponse
	url := fmt.Sprintf("%s/%s/%s", c.BaseURL, "consent-management-srv/v2/consent/versions/list", consentID)
	client, err := util.NewHTTPClient(url, http.MethodGet, c.AccessToken)
	if err != nil {
		return nil, err
	}
	res, err := client.MakeRequest(ctx, nil)
	if err = util.HandleResponseError(res, err); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err = util.ProcessResponse(res, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *ConsentVersion) UpsertLocal(ctx context.Context, consentLocal ConsentLocalModel) (*ConsentLocalResponse, error) {
	var response ConsentLocalResponse
	url := fmt.Sprintf("%s/%s", c.BaseURL, "consent-management-srv/v2/consent/locale")
	client, err := util.NewHTTPClient(url, http.MethodPost, c.AccessToken)
	if err != nil {
		return nil, err
	}
	res, err := client.MakeRequest(ctx, consentLocal)
	if err = util.HandleResponseError(res, err); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err = util.ProcessResponse(res, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *ConsentVersion) GetLocal(ctx context.Context, consentVersionID string, locale string) (*ConsentLocalResponse, error) {
	var response ConsentLocalResponse
	url := fmt.Sprintf("%s/%s/%s?locale=%s", c.BaseURL, "consent-management-srv/v2/consent/locale", consentVersionID, locale)
	client, err := util.NewHTTPClient(url, http.MethodGet, c.AccessToken)
	if err != nil {
		return nil, err
	}
	res, err := client.MakeRequest(ctx, nil)
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
