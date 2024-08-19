package cidaas

import (
	"fmt"
	"net/http"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/util"
)

type ScopeGroupConfig struct {
	ID          string `json:"_id,omitempty"`
	GroupName   string `json:"group_name,omitempty"`
	Description string `json:"description,omitempty"`
	CreatedTime string `json:"createdTime,omitempty"`
	UpdatedTime string `json:"updatedTime,omitempty"`
}

type ScopeGroupResponse struct {
	Success bool             `json:"success,omitempty"`
	Status  int              `json:"status,omitempty"`
	Data    ScopeGroupConfig `json:"data,omitempty"`
}

type ScopeGroup struct {
	ClientConfig
}
type ScopeGroupService interface {
	Upsert(sg ScopeGroupConfig) (*ScopeGroupResponse, error)
	Get(scopeGroupName string) (*ScopeGroupResponse, error)
	Delete(scopeGroupName string) error
}

func NewScopeGroup(clientConfig ClientConfig) ScopeGroupService {
	return &ScopeGroup{clientConfig}
}

func (c *ScopeGroup) Upsert(sg ScopeGroupConfig) (*ScopeGroupResponse, error) {
	var response ScopeGroupResponse
	url := fmt.Sprintf("%s/%s", c.BaseURL, "scopes-srv/group")
	httpClient := util.NewHTTPClient(url, http.MethodPost, c.AccessToken)

	res, err := httpClient.MakeRequest(sg)
	if err = util.HandleResponseError(res, err); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err = util.ProcessResponse(res, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *ScopeGroup) Get(scopeGroupName string) (*ScopeGroupResponse, error) {
	var response ScopeGroupResponse
	url := fmt.Sprintf("%s/%s?group_name=%s", c.BaseURL, "scopes-srv/group", scopeGroupName)
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

func (c *ScopeGroup) Delete(scopeGroupName string) error {
	url := fmt.Sprintf("%s/%s/%s", c.BaseURL, "scopes-srv/group", scopeGroupName)
	httpClient := util.NewHTTPClient(url, http.MethodDelete, c.AccessToken)

	res, err := httpClient.MakeRequest(nil)
	if err = util.HandleResponseError(res, err); err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}
