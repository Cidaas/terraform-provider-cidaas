package cidaas

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/util"
)

type UserGroupData struct {
	ID                          string            `json:"_id,omitempty"`
	GroupType                   string            `json:"groupType,omitempty"`
	GroupID                     string            `json:"groupId,omitempty"`
	GroupName                   string            `json:"groupName,omitempty"`
	ParentID                    string            `json:"parentId,omitempty"`
	LogoURL                     string            `json:"logoUrl,omitempty"`
	Description                 string            `json:"description,omitempty"`
	MakeFirstUserAdmin          bool              `json:"make_first_user_admin,omitempty"`
	MemberProfileVisibility     string            `json:"memberProfileVisibility,omitempty"`
	NoneMemberProfileVisibility string            `json:"noneMemberProfileVisibility,omitempty"`
	GroupOwner                  string            `json:"groupOwner,omitempty"`
	CustomFields                map[string]string `json:"customFields,omitempty"`
	CreatedTime                 string            `json:"createdTime,omitempty"`
	UpdatedTime                 string            `json:"updatedTime,omitempty"`
}

type UserGroupResponse struct {
	Success bool          `json:"success,omitempty"`
	Status  int           `json:"status,omitempty"`
	Data    UserGroupData `json:"data,omitempty"`
}

type SubGroupResponse struct {
	Success bool `json:"success,omitempty"`
	Status  int  `json:"status,omitempty"`
	Data    struct {
		Groups []UserGroupData `json:"groups"`
	} `json:"data,omitempty"`
}

type UserGroup struct {
	HTTPClient util.HTTPClientInterface
}
type UserGroupService interface {
	Create(ug UserGroupData) (*UserGroupResponse, error)
	Get(groupID string) (*UserGroupResponse, error)
	Update(ug UserGroupData) error
	Delete(groupID string) error
	GetSubGroups(parentID string) ([]UserGroupData, error)
}

func NewUserGroup(httpClient util.HTTPClientInterface) UserGroupService {
	return &UserGroup{HTTPClient: httpClient}
}

func (c *UserGroup) Create(ug UserGroupData) (*UserGroupResponse, error) {
	c.HTTPClient.SetURL(fmt.Sprintf("%s/%s", c.HTTPClient.GetHost(), "groups-srv/usergroups"))
	c.HTTPClient.SetMethod(http.MethodPost)
	res, err := c.HTTPClient.MakeRequest(ug)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var response UserGroupResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %w", err)
	}
	return &response, nil
}

func (c *UserGroup) Get(groupID string) (*UserGroupResponse, error) {
	c.HTTPClient.SetURL(fmt.Sprintf("%s/%s/%s", c.HTTPClient.GetHost(), "groups-srv/usergroups", groupID))
	c.HTTPClient.SetMethod(http.MethodGet)
	res, err := c.HTTPClient.MakeRequest(nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var response UserGroupResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %w", err)
	}
	return &response, nil
}

func (c *UserGroup) Update(ug UserGroupData) error {
	c.HTTPClient.SetURL(fmt.Sprintf("%s/%s", c.HTTPClient.GetHost(), "groups-srv/usergroups"))
	c.HTTPClient.SetMethod(http.MethodPut)
	res, err := c.HTTPClient.MakeRequest(ug)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}

func (c *UserGroup) Delete(groupID string) error {
	c.HTTPClient.SetURL(fmt.Sprintf("%s/%s/%s", c.HTTPClient.GetHost(), "groups-srv/usergroups", groupID))
	c.HTTPClient.SetMethod(http.MethodDelete)
	res, err := c.HTTPClient.MakeRequest(nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}

func (c *UserGroup) GetSubGroups(parentID string) ([]UserGroupData, error) {
	c.HTTPClient.SetURL(fmt.Sprintf("%s/%s", c.HTTPClient.GetHost(), "groups-srv/graph/usergroups"))
	c.HTTPClient.SetMethod(http.MethodPost)
	payload := map[string]string{"parentId": parentID}
	res, err := c.HTTPClient.MakeRequest(payload)
	if res.StatusCode == http.StatusNoContent {
		return []UserGroupData{}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to fetch sub groups: %w", err)
	}
	defer res.Body.Close()
	var response SubGroupResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode subgroup response: %w", err)
	}
	return response.Data.Groups, nil
}
