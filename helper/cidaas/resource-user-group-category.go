package cidaas

import (
	"encoding/json"
	"fmt"
	"terraform-provider-cidaas/helper/util"
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

func (c *CidaasClient) CreateUserGroupCategory(userGroupCategory UserGroupCategory) (response *UserGroupCategoryResponse, err error) {
	h := util.HttpClient{
		Token: c.TokenData.AccessToken,
	}
	url := c.BaseUrl + "/groups-srv/grouptypes"
	res, err := h.Post(url, userGroupCategory)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %v", err)
	}
	return response, nil
}

func (c *CidaasClient) GetUserGroupCategory(group_type string) (response *UserGroupCategoryResponse, err error) {
	url := c.BaseUrl + "/groups-srv/grouptypes?groupType=" + group_type
	h := util.HttpClient{
		Token: c.TokenData.AccessToken,
	}
	res, err := h.Get(url)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %v", err)
	}
	return response, nil
}

func (c *CidaasClient) UpdateUserGroupCategory(userGroupCategory UserGroupCategory) (response *UserGroupCategoryResponse, err error) {
	h := util.HttpClient{
		Token: c.TokenData.AccessToken,
	}
	url := c.BaseUrl + "/groups-srv/grouptypes"
	res, err := h.Put(url, userGroupCategory)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %v", err)
	}
	return response, nil
}
func (c *CidaasClient) DeleteUserGroupCategory(group_type string) (response *UserGroupCategoryResponse, err error) {
	url := c.BaseUrl + "/groups-srv/grouptypes/" + group_type
	h := util.HttpClient{
		Token: c.TokenData.AccessToken,
	}
	res, err := h.Delete(url)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %v", err)
	}
	return response, nil
}
