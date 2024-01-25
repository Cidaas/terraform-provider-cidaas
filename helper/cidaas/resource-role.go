package cidaas

import (
	"encoding/json"
	"fmt"
	"terraform-provider-cidaas/helper/util"
)

type RoleConfig struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Role        string `json:"role,omitempty"`
}

type RoleResponse struct {
	Success bool       `json:"success,omitempty"`
	Status  int        `json:"status,omitempty"`
	Data    RoleConfig `json:"data,omitempty"`
}

type DeleteRoleResponse struct {
	Success bool   `json:"success,omitempty"`
	Status  int    `json:"status,omitempty"`
	Data    bool   `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

func (c *CidaasClient) UpsertRole(roleConfig RoleConfig) (response *RoleResponse, err error) {
	h := util.HttpClient{
		Token: c.TokenData.AccessToken,
	}
	url := c.BaseUrl + "/roles-srv/role"
	res, err := h.Post(url, roleConfig)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %v", err)
	}
	return response, nil
}

func (c *CidaasClient) GetRole(role string) (response *RoleResponse, err error) {
	url := c.BaseUrl + "/roles-srv/role?role=" + role
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

func (c *CidaasClient) DeleteRole(role string) (response *DeleteRoleResponse, err error) {
	url := c.BaseUrl + "/roles-srv/role?role=" + role
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
