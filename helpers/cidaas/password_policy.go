package cidaas

import (
	"encoding/json"
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
	HTTPClient util.HTTPClientInterface
}
type PasswordPolicyService interface {
	Get() (*PasswordPolicyResponse, error)
	Update(cp PasswordPolicyModel) error
}

func NewPasswordPolicy(httpClient util.HTTPClientInterface) PasswordPolicyService {
	return &PasswordPolicy{HTTPClient: httpClient}
}

func (c *PasswordPolicy) Get() (*PasswordPolicyResponse, error) {
	c.HTTPClient.SetURL(fmt.Sprintf("%s/%s", c.HTTPClient.GetHost(), "password-policy-srv/policy"))
	c.HTTPClient.SetMethod(http.MethodGet)
	res, err := c.HTTPClient.MakeRequest(nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var response PasswordPolicyResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %w", err)
	}
	return &response, nil
}

func (c *PasswordPolicy) Update(payload PasswordPolicyModel) error {
	c.HTTPClient.SetURL(fmt.Sprintf("%s/%s", c.HTTPClient.GetHost(), "password-policy-srv/policy"))
	c.HTTPClient.SetMethod(http.MethodPut)
	res, err := c.HTTPClient.MakeRequest(payload)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}
