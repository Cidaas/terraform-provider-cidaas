package cidaas

import (
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
	ClientConfig
}

type RoleService interface {
	UpsertRole(role RoleModel) (*RoleResponse, error)
	GetRole(role string) (*RoleResponse, error)
	DeleteRole(role string) error
}

func NewRole(clientConfig ClientConfig) RoleService {
	return &Role{clientConfig}
}

func (r *Role) UpsertRole(role RoleModel) (*RoleResponse, error) {
	var response RoleResponse
	url := fmt.Sprintf("%s/%s", r.BaseURL, "roles-srv/role")
	httpClient := util.NewHTTPClient(url, http.MethodPost, r.AccessToken)

	res, err := httpClient.MakeRequest(role)
	if err = util.HandleResponseError(res, err); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err = util.ProcessResponse(res, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (r *Role) GetRole(role string) (*RoleResponse, error) {
	var response RoleResponse
	url := fmt.Sprintf("%s/%s?role=%s", r.BaseURL, "roles-srv/role", role)

	httpClient := util.NewHTTPClient(url, http.MethodGet, r.AccessToken)

	res, err := httpClient.MakeRequest(role)
	if err = util.HandleResponseError(res, err); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err = util.ProcessResponse(res, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (r *Role) DeleteRole(role string) error {
	url := fmt.Sprintf("%s/%s?role=%s", r.BaseURL, "roles-srv/role", role)

	httpClient := util.NewHTTPClient(url, http.MethodDelete, r.AccessToken)

	res, err := httpClient.MakeRequest(nil)
	if err = util.HandleResponseError(res, err); err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}
