package cidaas

import (
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
	ClientConfig
}
type ConsentService interface {
	Upsert(consent ConsentModel) (*ConsentResponse, error)
	GetConsentInstances(consentGroupID string) (*ConsentInstanceResponse, error)
	Delete(consentID string) error
	GetAll() ([]ConsentModel, error)
}

func NewConsent(clientConfig ClientConfig) ConsentService {
	return &ConsentClient{clientConfig}
}

func (c *ConsentClient) Upsert(consentConfig ConsentModel) (*ConsentResponse, error) {
	var response ConsentResponse
	url := fmt.Sprintf("%s/%s", c.BaseURL, "consent-management-srv/v2/consent/instance")
	httpClient := util.NewHTTPClient(url, http.MethodPost, c.AccessToken)

	res, err := httpClient.MakeRequest(consentConfig)
	if err = util.HandleResponseError(res, err); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err = util.ProcessResponse(res, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *ConsentClient) GetConsentInstances(consentGroupID string) (*ConsentInstanceResponse, error) {
	var response ConsentInstanceResponse
	url := fmt.Sprintf("%s/%s/%s", c.BaseURL, "consent-management-srv/v2/consent/instance", consentGroupID)
	httpClient := util.NewHTTPClient(url, http.MethodGet, c.AccessToken)

	res, err := httpClient.MakeRequest(nil)
	if res.StatusCode == http.StatusNoContent {
		return &ConsentInstanceResponse{
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

func (c *ConsentClient) Delete(consentID string) error {
	url := fmt.Sprintf("%s/%s/%s", c.BaseURL, "consent-management-srv/v2/consent/instance", consentID)
	httpClient := util.NewHTTPClient(url, http.MethodDelete, c.AccessToken)

	res, err := httpClient.MakeRequest(nil)
	if err = util.HandleResponseError(res, err); err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}

func (c *ConsentClient) GetAll() ([]ConsentModel, error) {
	var response ConsentInstanceResponse
	url := fmt.Sprintf("%s/%s", c.BaseURL, "consent-management-srv/v2/consent/instance/all/list")
	httpClient := util.NewHTTPClient(url, http.MethodGet, c.AccessToken)

	res, err := httpClient.MakeRequest(nil)
	if err = util.HandleResponseError(res, err); err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err = util.ProcessResponse(res, &response); err != nil {
		return nil, err
	}
	return response.Data, nil
}
