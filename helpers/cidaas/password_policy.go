package cidaas

import (
	"fmt"
	"net/http"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/util"
)

type PasswordPolicyModel struct {
	ID                string `json:"id"`
	PolicyName        string `json:"policy_name"`
	MaximumLength     int64  `json:"maximumLength"`
	MinimumLength     int64  `json:"minimumLength"`
	NoOfSpecialChars  int64  `json:"noOfSpecialChars"`
	NoOfDigits        int64  `json:"noOfDigits"`
	LowerAndUppercase bool   `json:"lowerAndUpperCase"`
}

type PasswordPolicyResponse struct {
	Success bool                `json:"success"`
	Status  int                 `json:"status"`
	Data    PasswordPolicyModel `json:"data,omitempty"`
}

type PasswordPolicy struct {
	ClientConfig
}
type PasswordPolicyService interface {
	Get(id string) (*PasswordPolicyResponse, error)
	Upsert(cp PasswordPolicyModel) (*PasswordPolicyResponse, error)
	Delete(id string) error
}

func NewPasswordPolicy(clientConfig ClientConfig) PasswordPolicyService {
	return &PasswordPolicy{clientConfig}
}

func (p *PasswordPolicy) Get(id string) (*PasswordPolicyResponse, error) {
	var response PasswordPolicyResponse
	url := fmt.Sprintf("%s/%s?id=%s", p.BaseURL, "password-policy-srv/policy", id)
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

func (p *PasswordPolicy) Upsert(payload PasswordPolicyModel) (*PasswordPolicyResponse, error) {
	var response PasswordPolicyResponse
	url := fmt.Sprintf("%s/%s", p.BaseURL, "password-policy-srv/policy")
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

func (p *PasswordPolicy) Delete(id string) error {
	url := fmt.Sprintf("%s/%s/%s", p.BaseURL, "password-policy-srv/policy", id)
	httpClient := util.NewHTTPClient(url, http.MethodDelete, p.AccessToken)

	res, err := httpClient.MakeRequest(nil)
	if err = util.HandleResponseError(res, err); err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}
