package cidaas

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UserGroupCategory struct {
	RoleMode     string   `json:"roleMode,omitempty"`
	GroupType    string   `json:"groupType,omitempty"`
	Description  string   `json:"description,omitempty"`
	AllowedRoles []string `json:"allowedRoles,omitempty"`
	ObjectOwner  string   `json:"objectOwner,omitempty"`
}

type UserGroupCategoryResponse struct {
	Success bool              `json:"success,omitempty"`
	Status  int               `json:"status,omitempty"`
	Data    UserGroupCategory `json:"data,omitempty"`
}

func (c *Client) CreateUserGroupCategory(ugc UserGroupCategory) (response *UserGroupCategoryResponse, err error) {
	c.HTTPClient.URL = fmt.Sprintf("%s/%s", c.Config.BaseURL, "groups-srv/grouptypes")
	c.HTTPClient.HTTPMethod = http.MethodPost
	res, err := c.HTTPClient.MakeRequest(ugc)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %v", err)
	}
	return response, nil
}

func (c *Client) GetUserGroupCategory(group_type string) (response *UserGroupCategoryResponse, err error) {
	c.HTTPClient.URL = fmt.Sprintf("%s/%s?groupType=%s", c.Config.BaseURL, "groups-srv/grouptypes", group_type)
	c.HTTPClient.HTTPMethod = http.MethodGet
	res, err := c.HTTPClient.MakeRequest(nil)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %v", err)
	}
	return response, nil
}

func (c *Client) UpdateUserGroupCategory(ugc UserGroupCategory) (response *UserGroupCategoryResponse, err error) {
	c.HTTPClient.URL = fmt.Sprintf("%s/%s", c.Config.BaseURL, "groups-srv/grouptypes")
	c.HTTPClient.HTTPMethod = http.MethodPut
	res, err := c.HTTPClient.MakeRequest(ugc)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %v", err)
	}
	return response, nil
}
func (c *Client) DeleteUserGroupCategory(group_type string) (response *UserGroupCategoryResponse, err error) {
	c.HTTPClient.URL = fmt.Sprintf("%s/%s/%s", c.Config.BaseURL, "groups-srv/grouptypes", group_type)
	c.HTTPClient.HTTPMethod = http.MethodDelete
	res, err := c.HTTPClient.MakeRequest(nil)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %v", err)
	}
	return response, nil
}
