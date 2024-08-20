package cidaas

import (
	"fmt"
	"net/http"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/util"
)

type ConsentGroupConfig struct {
	ID          string `json:"_id,omitempty"`
	GroupName   string `json:"group_name,omitempty"`
	Description string `json:"description"` // description needs to set to empty string, so omitempty is removed here
	CreatedTime string `json:"createdTime,omitempty"`
	UpdatedTime string `json:"updatedTime,omitempty"`
}

type ConsentGroupResponse struct {
	Success bool               `json:"success,omitempty"`
	Status  int                `json:"status,omitempty"`
	Data    ConsentGroupConfig `json:"data,omitempty"`
}

type ConsentGroup struct {
	ClientConfig
}
type ConsentGroupService interface {
	Upsert(cg ConsentGroupConfig) (*ConsentGroupResponse, error)
	Get(consentGroupID string) (*ConsentGroupResponse, error)
	Delete(consentGroupID string) error
}

func NewConsentGroup(clientConfig ClientConfig) ConsentGroupService {
	return &ConsentGroup{clientConfig}
}

func (c *ConsentGroup) Upsert(cg ConsentGroupConfig) (*ConsentGroupResponse, error) {
	var response ConsentGroupResponse
	url := fmt.Sprintf("%s/%s", c.BaseURL, "consent-management-srv/v2/groups")
	httpClient := util.NewHTTPClient(url, http.MethodPost, c.AccessToken)

	res, err := httpClient.MakeRequest(cg)
	if err = util.HandleResponseError(res, err); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err = util.ProcessResponse(res, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *ConsentGroup) Get(consentGroupID string) (*ConsentGroupResponse, error) {
	var response ConsentGroupResponse
	url := fmt.Sprintf("%s/%s/%s", c.BaseURL, "consent-management-srv/v2/groups", consentGroupID)
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

func (c *ConsentGroup) Delete(consentGroupID string) error {
	url := fmt.Sprintf("%s/%s/%s", c.BaseURL, "consent-management-srv/v2/groups", consentGroupID)
	httpClient := util.NewHTTPClient(url, http.MethodDelete, c.AccessToken)

	res, err := httpClient.MakeRequest(nil)
	if err = util.HandleResponseError(res, err); err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}
