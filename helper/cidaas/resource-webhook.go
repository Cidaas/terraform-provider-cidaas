package cidaas

import (
	"encoding/json"
	"fmt"
	"terraform-provider-cidaas/helper/util"
)

type WebhookRequestPayload struct {
	AuthType      string                 `json:"auth_type,omitempty"`
	Url           string                 `json:"url,omitempty"`
	Events        []string               `json:"events,omitempty"`
	ApiKeyDetails map[string]interface{} `json:"apikeyDetails,omitempty"`
	ID            string                 `json:"_id,omitempty"`
}

type ApiKeyDetails struct {
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
	ID            string            `json:"_id,omitempty"`
	AuthType      string            `json:"auth_type,omitempty"`
	Url           string            `json:"url,omitempty"`
	Events        []string          `json:"events,omitempty"`
	ApiKeyDetails map[string]string `json:"apikeyDetails,omitempty"`
	Disable       bool              `json:"disable,omitempty"`
}

func (c *CidaasClient) CreateOrUpdateWebhook(wb WebhookRequestPayload) (response *WebhookResponse, err error) {
	url := c.BaseUrl + "/webhook-srv/webhook"
	h := util.HttpClient{
		Token: c.TokenData.AccessToken,
	}
	res, err := h.Post(url, wb)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %v", err)
	}
	return response, nil
}

func (c *CidaasClient) GetWebhook(id string) (response *WebhookResponse, err error) {
	url := c.BaseUrl + "/webhook-srv/webhook?id=" + id
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

func (c *CidaasClient) DeleteWebhook(wb_id string) (response *WebhookResponse, err error) {
	url := c.BaseUrl + "/webhook-srv/webhook/" + wb_id
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
