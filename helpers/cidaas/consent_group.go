package cidaas

import (
	"encoding/json"
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
	HTTPClient util.HTTPClientInterface
}
type ConsentGroupService interface {
	Upsert(cg ConsentGroupConfig) (*ConsentGroupResponse, error)
	Get(consentGroupID string) (*ConsentGroupResponse, error)
	Delete(consentGroupID string) error
}

func NewConsentGroup(httpClient util.HTTPClientInterface) ConsentGroupService {
	return &ConsentGroup{HTTPClient: httpClient}
}

func (c *ConsentGroup) Upsert(cg ConsentGroupConfig) (*ConsentGroupResponse, error) {
	c.HTTPClient.SetURL(fmt.Sprintf("%s/%s", c.HTTPClient.GetHost(), "consent-management-srv/v2/groups"))
	c.HTTPClient.SetMethod(http.MethodPost)
	res, err := c.HTTPClient.MakeRequest(cg)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var response ConsentGroupResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %w", err)
	}
	return &response, nil
}

func (c *ConsentGroup) Get(consentGroupID string) (*ConsentGroupResponse, error) {
	c.HTTPClient.SetURL(fmt.Sprintf("%s/%s/%s", c.HTTPClient.GetHost(), "consent-management-srv/v2/groups", consentGroupID))
	c.HTTPClient.SetMethod(http.MethodGet)
	res, err := c.HTTPClient.MakeRequest(nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var response ConsentGroupResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %w, %s", err, consentGroupID)
	}
	return &response, nil
}

func (c *ConsentGroup) Delete(consentGroupID string) error {
	c.HTTPClient.SetURL(fmt.Sprintf("%s/%s/%s", c.HTTPClient.GetHost(), "consent-management-srv/v2/groups", consentGroupID))
	c.HTTPClient.SetMethod(http.MethodDelete)
	res, err := c.HTTPClient.MakeRequest(nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}
