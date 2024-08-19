package cidaas

import (
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
	ClientConfig
}
type GroupTypeService interface {
	Create(gt GroupTypeData) (*GroupTypeResponse, error)
	Get(groupType string) (*GroupTypeResponse, error)
	Update(gt GroupTypeData) error
	Delete(groupType string) error
}

func NewGroupType(clientConfig ClientConfig) GroupTypeService {
	return &GroupType{clientConfig}
}

func (c *GroupType) Create(gt GroupTypeData) (*GroupTypeResponse, error) {
	var response GroupTypeResponse
	url := fmt.Sprintf("%s/%s", c.BaseURL, "groups-srv/grouptypes")
	httpClient := util.NewHTTPClient(url, http.MethodPost, c.AccessToken)

	res, err := httpClient.MakeRequest(gt)
	if err = util.HandleResponseError(res, err); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err = util.ProcessResponse(res, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *GroupType) Get(groupType string) (*GroupTypeResponse, error) {
	var response GroupTypeResponse
	url := fmt.Sprintf("%s/%s?groupType=%s", c.BaseURL, "groups-srv/grouptypes", groupType)
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

func (c *GroupType) Update(gt GroupTypeData) error {
	url := fmt.Sprintf("%s/%s", c.BaseURL, "groups-srv/grouptypes")
	httpClient := util.NewHTTPClient(url, http.MethodPut, c.AccessToken)

	res, err := httpClient.MakeRequest(gt)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}

func (c *GroupType) Delete(groupType string) error {
	url := fmt.Sprintf("%s/%s/%s", c.BaseURL, "groups-srv/grouptypes", groupType)
	httpClient := util.NewHTTPClient(url, http.MethodDelete, c.AccessToken)

	res, err := httpClient.MakeRequest(nil)
	if err = util.HandleResponseError(res, err); err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}
