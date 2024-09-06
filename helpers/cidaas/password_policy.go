package cidaas

import (
	"fmt"
	"net/http"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/util"
)

type PasswordPolicyModel struct {
	MaximumLength     int64 `json:"maximumLength"`
	MinimumLength     int64 `json:"minimumLength"`
	NoOfSpecialChars  int64 `json:"noOfSpecialChars"`
	NoOfDigits        int64 `json:"noOfDigits"`
	LowerAndUppercase bool  `json:"lowerAndUpperCase"`
	ReuseLimit        int64 `json:"reuseLimit"`
	ExpirationInDays  int64 `json:"expirationInDays"`
	NoOfDaysToRemind  int64 `json:"noOfDaysToRemind"`
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
	Get() (*PasswordPolicyResponse, error)
	Update(cp PasswordPolicyModel) error
}

func NewPasswordPolicy(clientConfig ClientConfig) PasswordPolicyService {
	return &PasswordPolicy{clientConfig}
}

func (p *PasswordPolicy) Get() (*PasswordPolicyResponse, error) {
	var response PasswordPolicyResponse
	url := fmt.Sprintf("%s/%s", p.BaseURL, "password-policy-srv/policy")
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

func (p *PasswordPolicy) Update(payload PasswordPolicyModel) error {
	url := fmt.Sprintf("%s/%s", p.BaseURL, "password-policy-srv/policy")
	httpClient := util.NewHTTPClient(url, http.MethodPut, p.AccessToken)

	res, err := httpClient.MakeRequest(payload)
	if err = util.HandleResponseError(res, err); err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}
