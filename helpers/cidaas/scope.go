package cidaas

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
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

func (c *Client) UpsertScope(sc Scope) (response *ScopeResponse, err error) {
	c.HTTPClient.URL = fmt.Sprintf("%s/%s", c.Config.BaseURL, "scopes-srv/scope")
	c.HTTPClient.HTTPMethod = http.MethodPost
	res, err := c.HTTPClient.MakeRequest(sc)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %v", err)
	}
	return response, nil
}

func (c *Client) GetScope(scope_key string) (response *ScopeResponse, err error) {
	c.HTTPClient.URL = fmt.Sprintf("%s/%s?scopekey=%s", c.Config.BaseURL, "scopes-srv/scope", strings.ToLower(scope_key))
	c.HTTPClient.HTTPMethod = http.MethodGet
	res, err := c.HTTPClient.MakeRequest(nil)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %v", err)
	}
	return response, nil
}

func (c *Client) DeleteScope(scope_key string) (response *DeleteScopeResponse, err error) {
	c.HTTPClient.URL = fmt.Sprintf("%s/%s/%s", c.Config.BaseURL, "scopes-srv/scope", strings.ToLower(scope_key))
	c.HTTPClient.HTTPMethod = http.MethodDelete
	res, err := c.HTTPClient.MakeRequest(nil)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %v", err)
	}
	return response, nil
}
