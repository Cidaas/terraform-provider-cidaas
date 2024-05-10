package cidaas

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func (c *Client) CreateUserGroup(ug UserGroup) (response *UserGroupResponse, err error) {
	c.HTTPClient.URL = fmt.Sprintf("%s/%s", c.Config.BaseURL, "groups-srv/usergroups")
	c.HTTPClient.HTTPMethod = http.MethodPost
	res, err := c.HTTPClient.MakeRequest(ug)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %v", err)
	}
	return response, nil
}

func (c *Client) GetUserGroup(group_id string) (response *UserGroupResponse, err error) {
	c.HTTPClient.URL = fmt.Sprintf("%s/%s/%s", c.Config.BaseURL, "groups-srv/usergroups", group_id)
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

func (c *Client) UpdateUserGroup(ug UserGroup) (response *UserGroupResponse, err error) {
	c.HTTPClient.URL = fmt.Sprintf("%s/%s", c.Config.BaseURL, "groups-srv/usergroups")
	c.HTTPClient.HTTPMethod = http.MethodPut
	res, err := c.HTTPClient.MakeRequest(ug)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %v", err)
	}
	return response, nil
}

func (c *Client) DeleteUserGroup(group_id string) (response *DeleteUserGroupResponse, err error) {
	c.HTTPClient.URL = fmt.Sprintf("%s/%s/%s", c.Config.BaseURL, "groups-srv/usergroups", group_id)
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
