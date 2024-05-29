package cidaas

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/util"
)

type GroupTypeData struct {
	ID           string   `json:"_id,omitempty"`
	RoleMode     string   `json:"roleMode,omitempty"`
	GroupType    string   `json:"groupType,omitempty"`
	Description  string   `json:"description,omitempty"`
	AllowedRoles []string `json:"allowedRoles,omitempty"`
	CreatedTime  string   `json:"createdTime,omitempty"`
	UpdatedTime  string   `json:"updatedTime,omitempty"`
}

type GroupTypeResponse struct {
	Success bool          `json:"success,omitempty"`
	Status  int           `json:"status,omitempty"`
	Data    GroupTypeData `json:"data,omitempty"`
}

type GroupType struct {
	HTTPClient util.HTTPClientInterface
}
type GroupTypeService interface {
	Create(gt GroupTypeData) (*GroupTypeResponse, error)
	Get(groupType string) (*GroupTypeResponse, error)
	Update(gt GroupTypeData) error
	Delete(groupType string) error
}

func NewGroupType(httpClient util.HTTPClientInterface) GroupTypeService {
	return &GroupType{HTTPClient: httpClient}
}

func (c *GroupType) Create(gt GroupTypeData) (*GroupTypeResponse, error) {
	c.HTTPClient.SetURL(fmt.Sprintf("%s/%s", c.HTTPClient.GetHost(), "groups-srv/grouptypes"))
	c.HTTPClient.SetMethod(http.MethodPost)
	res, err := c.HTTPClient.MakeRequest(gt)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var response GroupTypeResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %w", err)
	}
	return &response, nil
}

func (c *GroupType) Get(groupType string) (*GroupTypeResponse, error) {
	c.HTTPClient.SetURL(fmt.Sprintf("%s/%s?groupType=%s", c.HTTPClient.GetHost(), "groups-srv/grouptypes", groupType))
	c.HTTPClient.SetMethod(http.MethodGet)
	res, err := c.HTTPClient.MakeRequest(nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var response GroupTypeResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %w", err)
	}
	return &response, nil
}

func (c *GroupType) Update(gt GroupTypeData) error {
	c.HTTPClient.SetURL(fmt.Sprintf("%s/%s", c.HTTPClient.GetHost(), "groups-srv/grouptypes"))
	c.HTTPClient.SetMethod(http.MethodPut)
	res, err := c.HTTPClient.MakeRequest(gt)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}

func (c *GroupType) Delete(groupType string) error {
	c.HTTPClient.SetURL(fmt.Sprintf("%s/%s/%s", c.HTTPClient.GetHost(), "groups-srv/grouptypes", groupType))
	c.HTTPClient.SetMethod(http.MethodDelete)
	res, err := c.HTTPClient.MakeRequest(nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}
