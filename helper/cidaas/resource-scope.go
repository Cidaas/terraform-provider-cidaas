package cidaas

import (
	"encoding/json"
	"fmt"
	"strings"
	"terraform-provider-cidaas/helper/util"
)

type Scope struct {
	ID                    string                  `json:"_id,omitempty"`
	LocaleWiseDescription []ScopeLocalDescription `json:"localeWiseDescription,omitempty"`
	SecurityLevel         string                  `json:"securityLevel,omitempty"`
	ScopeKey              string                  `json:"scopeKey,omitempty"`
	RequiredUserConsent   bool                    `json:"requiredUserConsent,omitempty"`
	GroupName             []string                `json:"group_name,omitempty"`
	ScopeOwner            string                  `json:"scopeOwner,omitempty"`
}

type ScopeLocalDescription struct {
	Locale      string `json:"locale,omitempty"`
	Language    string `json:"language,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

type ScopeResponse struct {
	Success bool   `json:"success,omitempty"`
	Status  int    `json:"status,omitempty"`
	Data    Scope  `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

type DeleteScopeResponse struct {
	Success bool   `json:"success,omitempty"`
	Status  int    `json:"status,omitempty"`
	Data    bool   `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

func (c *CidaasClient) CreateOrUpdateScope(sc Scope) (response *ScopeResponse, err error) {
	url := c.BaseUrl + "/scopes-srv/scope"
	h := util.HttpClient{
		Token: c.TokenData.AccessToken,
	}
	res, err := h.Post(url, sc)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %v", err)
	}
	return response, nil
}

func (c *CidaasClient) GetScope(scope_key string) (response *ScopeResponse, err error) {
	url := c.BaseUrl + "/scopes-srv/scope?scopekey=" + strings.ToLower(scope_key)
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

func (c *CidaasClient) DeleteScope(scope_key string) (response *DeleteScopeResponse, err error) {
	url := c.BaseUrl + "/scopes-srv/scope/" + strings.ToLower(scope_key)
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
