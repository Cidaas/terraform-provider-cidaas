package cidaas

import (
	"fmt"
	"net/http"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/util"
)

type PasswordPolicyModel struct {
	ID             string  `json:"_id"`
	PolicyName     string  `json:"policy_name"`
	PasswordPolicy *Policy `json:"passwordPolicy"`
}

type Policy struct {
	BlockCompromised  bool              `json:"blockCompromised"`
	DenyUsageCount    int64             `json:"denyUsageCount"`
	StrengthRegexes   []string          `json:"strengthRegexes"`
	ChangeEnforcement ChangeEnforcement `json:"changeEnforcement"`
}

type ChangeEnforcement struct {
	ExpirationInDays       int64 `json:"expirationInDays"`
	NotifyUserBeforeInDays int64 `json:"notifyUserBeforeInDays"`
}

type PasswordPolicyResponse struct {
	Success bool                `json:"success"`
	Status  int                 `json:"status"`
	Data    PasswordPolicyModel `json:"data,omitempty"`
}

type PasswordPolicyUpdateResponse struct {
	Success bool `json:"success"`
	Status  int  `json:"status"`
	Data    bool `json:"data"`
}

type PasswordPolicy struct {
	ClientConfig
}
type PasswordPolicyService interface {
	Get(id string) (*PasswordPolicyResponse, error)
	Create(cp PasswordPolicyModel) (*PasswordPolicyResponse, error)
	Update(cp PasswordPolicyModel) (*PasswordPolicyUpdateResponse, error)
	Delete(id string) error
}

func NewPasswordPolicy(clientConfig ClientConfig) PasswordPolicyService {
	return &PasswordPolicy{clientConfig}
}

func (p *PasswordPolicy) Get(id string) (*PasswordPolicyResponse, error) {
	var response PasswordPolicyResponse
	url := fmt.Sprintf("%s/%s/%s", p.BaseURL, "verification-actions-srv/policies", id)
	httpClient := util.NewHTTPClient(url, http.MethodGet, p.AccessToken)

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

func (p *PasswordPolicy) Create(payload PasswordPolicyModel) (*PasswordPolicyResponse, error) {
	var response PasswordPolicyResponse
	url := fmt.Sprintf("%s/%s", p.BaseURL, "verification-actions-srv/policies")
	httpClient := util.NewHTTPClient(url, http.MethodPost, p.AccessToken)

	res, err := httpClient.MakeRequest(payload)
	if err = util.HandleResponseError(res, err); err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err = util.ProcessResponse(res, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (p *PasswordPolicy) Update(payload PasswordPolicyModel) (*PasswordPolicyUpdateResponse, error) {
	var response PasswordPolicyUpdateResponse
	url := fmt.Sprintf("%s/%s", p.BaseURL, "verification-actions-srv/policies")
	httpClient := util.NewHTTPClient(url, http.MethodPut, p.AccessToken)

	res, err := httpClient.MakeRequest(payload)
	if err = util.HandleResponseError(res, err); err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err = util.ProcessResponse(res, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (p *PasswordPolicy) Delete(id string) error {
	url := fmt.Sprintf("%s/%s/%s", p.BaseURL, "verification-actions-srv/policies", id)
	httpClient := util.NewHTTPClient(url, http.MethodDelete, p.AccessToken)

	res, err := httpClient.MakeRequest(nil)
	if err = util.HandleResponseError(res, err); err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}
