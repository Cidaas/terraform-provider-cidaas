package cidaas

import (
	"encoding/json"
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
	HTTPClient util.HTTPClientInterface
}

type ConsentVersionService interface {
	Upsert(consent ConsentVersionModel) (*ConsentVersionResponse, error)
	Get(consentID string) (*ConsentVersionReadResponse, error)
	UpsertLocal(consentLocal ConsentLocalModel) (*ConsentLocalResponse, error)
	GetLocal(consentVersionID string, locale string) (*ConsentLocalResponse, error)
}

func NewConsentVersion(httpClient util.HTTPClientInterface) ConsentVersionService {
	return &ConsentVersion{HTTPClient: httpClient}
}

func (c *ConsentVersion) Upsert(consentVersionConfig ConsentVersionModel) (*ConsentVersionResponse, error) {
	c.HTTPClient.SetURL(fmt.Sprintf("%s/%s", c.HTTPClient.GetHost(), "consent-management-srv/v2/consent/versions"))
	c.HTTPClient.SetMethod(http.MethodPost)
	res, err := c.HTTPClient.MakeRequest(consentVersionConfig)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var response ConsentVersionResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %w", err)
	}
	return &response, nil
}

func (c *ConsentVersion) Get(consentID string) (*ConsentVersionReadResponse, error) {
	c.HTTPClient.SetURL(fmt.Sprintf("%s/%s/%s", c.HTTPClient.GetHost(), "consent-management-srv/v2/consent/versions/list", consentID))
	c.HTTPClient.SetMethod(http.MethodGet)
	res, err := c.HTTPClient.MakeRequest(nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var response ConsentVersionReadResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %w, %s", err, consentID)
	}
	return &response, nil
}

func (c *ConsentVersion) UpsertLocal(consentLocal ConsentLocalModel) (*ConsentLocalResponse, error) {
	c.HTTPClient.SetURL(fmt.Sprintf("%s/%s", c.HTTPClient.GetHost(), "consent-management-srv/v2/consent/locale"))
	c.HTTPClient.SetMethod(http.MethodPost)
	res, err := c.HTTPClient.MakeRequest(consentLocal)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var response ConsentLocalResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %w", err)
	}
	return &response, nil
}

func (c *ConsentVersion) GetLocal(consentVersionID string, locale string) (*ConsentLocalResponse, error) {
	c.HTTPClient.SetURL(fmt.Sprintf("%s/%s/%s?locale=%s", c.HTTPClient.GetHost(), "consent-management-srv/v2/consent/locale", consentVersionID, locale))
	c.HTTPClient.SetMethod(http.MethodGet)
	res, err := c.HTTPClient.MakeRequest(nil)
	if res.StatusCode == http.StatusNoContent {
		return &ConsentLocalResponse{
			Success: false,
			Status:  http.StatusNoContent,
		}, nil
	}
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var response ConsentLocalResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %w, %s", err, consentVersionID)
	}
	return &response, nil
}
