package cidaas

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/util"
)

type RoleModel struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Role        string `json:"role,omitempty"`
}

type RoleResponse struct {
	Success bool      `json:"success,omitempty"`
	Status  int       `json:"status,omitempty"`
	Data    RoleModel `json:"data,omitempty"`
}

type Role struct {
	HTTPClient util.HTTPClientInterface
}

type RoleService interface {
	UpsertRole(role RoleModel) (*RoleResponse, error)
	GetRole(role string) (*RoleResponse, error)
	DeleteRole(role string) error
}

func NewRole(httpClient util.HTTPClientInterface) RoleService {
	return &Role{HTTPClient: httpClient}
}

func (c *Role) UpsertRole(role RoleModel) (*RoleResponse, error) {
	c.HTTPClient.SetURL(fmt.Sprintf("%s/%s", c.HTTPClient.GetHost(), "roles-srv/role"))
	c.HTTPClient.SetMethod(http.MethodPost)
	res, err := c.HTTPClient.MakeRequest(role)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var response RoleResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response body, %v", err)
	}
	return &response, nil
}

func (c *Role) GetRole(role string) (*RoleResponse, error) {
	c.HTTPClient.SetURL(fmt.Sprintf("%s/%s?role=%s", c.HTTPClient.GetHost(), "roles-srv/role", role))
	c.HTTPClient.SetMethod(http.MethodGet)
	res, err := c.HTTPClient.MakeRequest(nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var response RoleResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response body, %v", err)
	}
	return &response, nil
}

func (c *Role) DeleteRole(role string) error {
	c.HTTPClient.SetURL(fmt.Sprintf("%s/%s?role=%s", c.HTTPClient.GetHost(), "roles-srv/role", role))
	c.HTTPClient.SetMethod(http.MethodDelete)
	res, err := c.HTTPClient.MakeRequest(nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}
