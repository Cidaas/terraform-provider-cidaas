package cidaas

import (
	"encoding/json"
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
	HTTPClient util.HTTPClientInterface
}
type ScopeGroupService interface {
	Upsert(sg ScopeGroupConfig) (*ScopeGroupResponse, error)
	Get(scopeGroupName string) (*ScopeGroupResponse, error)
	Delete(scopeGroupName string) error
}

func NewScopeGroup(httpClient util.HTTPClientInterface) ScopeGroupService {
	return &ScopeGroup{HTTPClient: httpClient}
}

func (c *ScopeGroup) Upsert(sg ScopeGroupConfig) (*ScopeGroupResponse, error) {
	c.HTTPClient.SetURL(fmt.Sprintf("%s/%s", c.HTTPClient.GetHost(), "scopes-srv/group"))
	c.HTTPClient.SetMethod(http.MethodPost)
	res, err := c.HTTPClient.MakeRequest(sg)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var response ScopeGroupResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %w", err)
	}
	return &response, nil
}

func (c *ScopeGroup) Get(scopeGroupName string) (*ScopeGroupResponse, error) {
	c.HTTPClient.SetURL(fmt.Sprintf("%s/%s?group_name=%s", c.HTTPClient.GetHost(), "scopes-srv/group", scopeGroupName))
	c.HTTPClient.SetMethod(http.MethodGet)
	res, err := c.HTTPClient.MakeRequest(nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var response ScopeGroupResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %w", err)
	}
	return &response, nil
}

func (c *ScopeGroup) Delete(scopeGroupName string) error {
	c.HTTPClient.SetURL(fmt.Sprintf("%s/%s/%s", c.HTTPClient.GetHost(), "scopes-srv/group", scopeGroupName))
	c.HTTPClient.SetMethod(http.MethodDelete)
	res, err := c.HTTPClient.MakeRequest(nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}
