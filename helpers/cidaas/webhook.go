package cidaas

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	AllowedAuthType          = []string{"APIKEY", "TOTP", "CIDAAS_OAUTH2"}
	AllowedKeyPlacementValue = []string{"query", "header"}
)

type WebhookRequestPayload struct {
	AuthType          string        `json:"auth_type,omitempty"`
	URL               string        `json:"url,omitempty"`
	Events            []string      `json:"events,omitempty"`
	ID                string        `json:"_id,omitempty"`
	APIKeyDetails     APIKeyDetails `json:"apikeyDetails,omitempty"`
	TotpDetails       TotpDetails   `json:"totpDetails,omitempty"`
	CidaasAuthDetails AuthDetails   `json:"cidaasAuthDetails,omitempty"`
}

type APIKeyDetails struct {
	ApikeyPlaceholder string `json:"apikey_placeholder,omitempty"`
	ApikeyPlacement   string `json:"apikey_placement,omitempty"`
	Apikey            string `json:"apikey,omitempty"`
}

type WebhookResponse struct {
	Success bool         `json:"success,omitempty"`
	Status  int          `json:"status,omitempty"`
	Data    ResponseData `json:"data,omitempty"`
	Error   string       `json:"error,omitempty"`
}

type ResponseData struct {
	ID                string        `json:"_id,omitempty"`
	AuthType          string        `json:"auth_type,omitempty"`
	URL               string        `json:"url,omitempty"`
	Events            []string      `json:"events,omitempty"`
	APIKeyDetails     APIKeyDetails `json:"apikeyDetails,omitempty"`
	Disable           bool          `json:"disable,omitempty"`
	TotpDetails       TotpDetails   `json:"totpDetails,omitempty"`
	CidaasAuthDetails AuthDetails   `json:"cidaasAuthDetails,omitempty"`
}

type TotpDetails struct {
	TotpPlaceholder string `json:"totp_placeholder,omitempty"`
	TotpPlacement   string `json:"totp_placement,omitempty"`
	TotpKey         string `json:"totpkey,omitempty"`
}
type AuthDetails struct {
	ClientID string `json:"client_id,omitempty"`
}

func (c *Client) CreateOrUpdateWebhook(wb *WebhookRequestPayload) (response *WebhookResponse, err error) {
	c.HTTPClient.URL = fmt.Sprintf("%s/%s", c.Config.BaseURL, "webhook-srv/webhook")
	c.HTTPClient.HTTPMethod = http.MethodPost
	res, err := c.HTTPClient.MakeRequest(wb)
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

func (c *Client) GetWebhook(id string) (response *WebhookResponse, err error) {
	c.HTTPClient.URL = fmt.Sprintf("%s/%s?id=%s", c.Config.BaseURL, "webhook-srv/webhook", id)
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

func (c *Client) DeleteWebhook(id string) (response *WebhookResponse, err error) {
	c.HTTPClient.URL = fmt.Sprintf("%s/%s/%s", c.Config.BaseURL, "webhook-srv/webhook", id)
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
