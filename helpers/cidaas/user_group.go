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
	ClientConfig
}
type UserGroupService interface {
	Create(ug UserGroupData) (*UserGroupResponse, error)
	Get(groupID string) (*UserGroupResponse, error)
	Update(ug UserGroupData) error
	Delete(groupID string) error
	GetSubGroups(parentID string) ([]UserGroupData, error)
}

func NewUserGroup(clientConfig ClientConfig) UserGroupService {
	return &UserGroup{clientConfig}
}

func (c *UserGroup) Create(ug UserGroupData) (*UserGroupResponse, error) {
	var response UserGroupResponse
	url := fmt.Sprintf("%s/%s", c.BaseURL, "groups-srv/usergroups")
	httpClient := util.NewHTTPClient(url, http.MethodPost, c.AccessToken)

	res, err := httpClient.MakeRequest(ug)
	if err = util.HandleResponseError(res, err); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err = util.ProcessResponse(res, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *UserGroup) Get(groupID string) (*UserGroupResponse, error) {
	var response UserGroupResponse
	url := fmt.Sprintf("%s/%s/%s", c.BaseURL, "groups-srv/usergroups", groupID)
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

func (c *UserGroup) Update(ug UserGroupData) error {
	url := fmt.Sprintf("%s/%s", c.BaseURL, "groups-srv/usergroups")
	httpClient := util.NewHTTPClient(url, http.MethodPut, c.AccessToken)

	res, err := httpClient.MakeRequest(ug)
	if err = util.HandleResponseError(res, err); err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}

func (c *UserGroup) Delete(groupID string) error {
	url := fmt.Sprintf("%s/%s/%s", c.BaseURL, "groups-srv/usergroups", groupID)
	httpClient := util.NewHTTPClient(url, http.MethodDelete, c.AccessToken)

	res, err := httpClient.MakeRequest(nil)
	if err = util.HandleResponseError(res, err); err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}

func (c *UserGroup) GetSubGroups(parentID string) ([]UserGroupData, error) {
	var response SubGroupResponse
	url := fmt.Sprintf("%s/%s", c.BaseURL, "groups-srv/graph/usergroups")
	payload := map[string]string{"parentId": parentID}
	httpClient := util.NewHTTPClient(url, http.MethodPost, c.AccessToken)

	res, err := httpClient.MakeRequest(payload)
	if res.StatusCode == http.StatusNoContent {
		return []UserGroupData{}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to fetch sub groups: %w", err)
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode subgroup response: %w", err)
	}
	return response.Data.Groups, nil
}
