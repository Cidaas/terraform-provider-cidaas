package cidaas

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/util"
)

type ConsentModel struct {
	ID             string `json:"_id,omitempty"`
	ConsentGroupID string `json:"consent_group_id,omitempty"`
	ConsentName    string `json:"consent_name,omitempty"`
	Enabled        bool   `json:"enabled"`
	CreatedTime    string `json:"createdTime,omitempty"`
	UpdatedTime    string `json:"updatedTime,omitempty"`
}

type ConsentResponse struct {
	Success bool         `json:"success,omitempty"`
	Status  int          `json:"status,omitempty"`
	Data    ConsentModel `json:"data,omitempty"`
}

type ConsentInstanceResponse struct {
	Success bool           `json:"success,omitempty"`
	Status  int            `json:"status,omitempty"`
	Data    []ConsentModel `json:"data,omitempty"`
}

type ConsentClient struct {
	HTTPClient util.HTTPClientInterface
}
type ConsentService interface {
	Upsert(consent ConsentModel) (*ConsentResponse, error)
	GetConsentInstances(consentGroupID string) (*ConsentInstanceResponse, error)
	Delete(consentID string) error
}

func NewConsent(httpClient util.HTTPClientInterface) ConsentService {
	return &ConsentClient{HTTPClient: httpClient}
}

func (c *ConsentClient) Upsert(consentConfig ConsentModel) (*ConsentResponse, error) {
	c.HTTPClient.SetURL(fmt.Sprintf("%s/%s", c.HTTPClient.GetHost(), "consent-management-srv/v2/consent/instance"))
	c.HTTPClient.SetMethod(http.MethodPost)
	res, err := c.HTTPClient.MakeRequest(consentConfig)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var response ConsentResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %w", err)
	}
	return &response, nil
}

func (c *ConsentClient) GetConsentInstances(consentGroupID string) (*ConsentInstanceResponse, error) {
	c.HTTPClient.SetURL(fmt.Sprintf("%s/%s/%s", c.HTTPClient.GetHost(), "consent-management-srv/v2/consent/instance", consentGroupID))
	c.HTTPClient.SetMethod(http.MethodGet)
	res, err := c.HTTPClient.MakeRequest(nil)
	if res.StatusCode == http.StatusNoContent {
		return &ConsentInstanceResponse{
			Success: false,
			Status:  http.StatusNoContent,
		}, nil
	}
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var response ConsentInstanceResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %w, %s", err, consentGroupID)
	}
	return &response, nil
}

func (c *ConsentClient) Delete(consentID string) error {
	c.HTTPClient.SetURL(fmt.Sprintf("%s/%s/%s", c.HTTPClient.GetHost(), "consent-management-srv/v2/consent/instance", consentID))
	c.HTTPClient.SetMethod(http.MethodDelete)
	res, err := c.HTTPClient.MakeRequest(nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}
