package cidaas

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ScopeGroupConfig struct {
	// ID          string `json:"_id,omitempty"`
	GroupName   string `json:"group_name,omitempty"`
	Description string `json:"description,omitempty"`
}

type ScopeGroupResponse struct {
	Success bool             `json:"success,omitempty"`
	Status  int              `json:"status,omitempty"`
	Data    ScopeGroupConfig `json:"data,omitempty"`
}

func (c *Client) UpsertScopeGroup(scopeGroup ScopeGroupConfig) (response *ScopeGroupResponse, err error) {
	c.HTTPClient.URL = fmt.Sprintf("%s/%s", c.Config.BaseURL, "scopes-srv/group")
	c.HTTPClient.HTTPMethod = http.MethodPost
	res, err := c.HTTPClient.MakeRequest(scopeGroup)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %w", err)
	}
	return response, nil
}

func (c *Client) GetScopeGroup(scopeGroupName string) (response *ScopeGroupResponse, err error) {
	c.HTTPClient.URL = fmt.Sprintf("%s/%s?group_name=%s", c.Config.BaseURL, "scopes-srv/group", scopeGroupName)
	c.HTTPClient.HTTPMethod = http.MethodGet
	res, err := c.HTTPClient.MakeRequest(nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %w", err)
	}
	return response, nil
}

func (c *Client) DeleteScopeGroup(scopeGroup string) (response *DeleteScopeResponse, err error) {
	c.HTTPClient.URL = fmt.Sprintf("%s/%s/%s", c.Config.BaseURL, "scopes-srv/group", scopeGroup)
	c.HTTPClient.HTTPMethod = http.MethodDelete
	res, err := c.HTTPClient.MakeRequest(nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %w", err)
	}
	return response, nil
}
