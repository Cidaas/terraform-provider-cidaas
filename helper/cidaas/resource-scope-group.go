package cidaas

import (
	"encoding/json"
	"fmt"
	"terraform-provider-cidaas/helper/util"
)

type ScopeGroupConfig struct {
	//ID          string `json:"_id,omitempty"`
	GroupName   string `json:"group_name,omitempty"`
	Description string `json:"description,omitempty"`
}

type ScopeGroupResponse struct {
	Success bool             `json:"success,omitempty"`
	Status  int              `json:"status,omitempty"`
	Data    ScopeGroupConfig `json:"data,omitempty"`
}

func (c *CidaasClient) UpsertScopeGroup(scopeGroup ScopeGroupConfig) (response *ScopeGroupResponse, err error) {
	h := util.HttpClient{
		Token: c.TokenData.AccessToken,
	}
	url := c.BaseUrl + "/scopes-srv/group"
	res, err := h.Post(url, scopeGroup)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %v", err)
	}
	return response, nil
}

func (c *CidaasClient) GetScopeGroup(scopeGroupName string) (response *ScopeGroupResponse, err error) {
	url := c.BaseUrl + "/scopes-srv/group?group_name=" + scopeGroupName
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

func (c *CidaasClient) DeleteScopeGroup(scopeGroup string) (response *DeleteScopeResponse, err error) {
	url := c.BaseUrl + "/scopes-srv/group/" + scopeGroup
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
