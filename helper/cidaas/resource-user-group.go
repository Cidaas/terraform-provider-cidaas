package cidaas

import (
	"encoding/json"
	"fmt"
	"terraform-provider-cidaas/helper/util"
)

type UserGroup struct {
	GroupType                   string                 `json:"groupType,omitempty"`
	GroupId                     string                 `json:"groupId,omitempty"`
	GroupName                   string                 `json:"groupName,omitempty"`
	LogoUrl                     string                 `json:"logoUrl,omitempty"`
	Description                 string                 `json:"description,omitempty"`
	MakeFirstUserAdmin          bool                   `json:"make_first_user_admin,omitempty"`
	CustomFields                map[string]interface{} `json:"customFields,omitempty"`
	MemberProfileVisibility     string                 `json:"memberProfileVisibility,omitempty"`
	NoneMemberProfileVisibility string                 `json:"noneMemberProfileVisibility,omitempty"`
	GroupOwner                  string                 `json:"groupOwner,omitempty"`
	ParentId                    string                 `json:"parentId,omitempty"`
}

type UserGroupResponse struct {
	Success bool      `json:"success,omitempty"`
	Status  int       `json:"status,omitempty"`
	Data    UserGroup `json:"data,omitempty"`
}

type DeleteUserGroupResponse struct {
	Success bool `json:"success,omitempty"`
	Status  int  `json:"status,omitempty"`
	Data    struct {
		Deleted bool `json:"deleted,omitempty"`
	} `json:"data,omitempty"`
}

func (c *CidaasClient) CreateUserGroup(userGroup UserGroup) (response *UserGroupResponse, err error) {
	h := util.HttpClient{
		Token: c.TokenData.AccessToken,
	}
	url := c.BaseUrl + "/groups-srv/usergroups"
	res, err := h.Post(url, userGroup)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %v", err)
	}
	return response, nil
}

func (c *CidaasClient) GetUserGroup(group_id string) (response *UserGroupResponse, err error) {
	url := c.BaseUrl + "/groups-srv/usergroups/" + group_id
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

func (c *CidaasClient) UpdateUserGroup(userGroup UserGroup) (response *UserGroupResponse, err error) {
	h := util.HttpClient{
		Token: c.TokenData.AccessToken,
	}
	url := c.BaseUrl + "/groups-srv/usergroups"
	res, err := h.Put(url, userGroup)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %v", err)
	}
	return response, nil
}

func (c *CidaasClient) DeleteUserGroup(group_id string) (response *DeleteUserGroupResponse, err error) {
	url := c.BaseUrl + "/groups-srv/usergroups/" + group_id
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
