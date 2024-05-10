package cidaas

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HostedPagePayload struct {
	ID            string       `json:"_id,omitempty"`
	GroupOwner    string       `json:"groupOwner,omitempty"`
	DefaultLocale string       `json:"default_locale,omitempty"`
	HostedPages   []HostedPage `json:"hosted_pages,omitempty"`
}

type HostedPage struct {
	HostedPageID string `json:"hosted_page_id,omitempty"`
	Content      string `json:"content,omitempty"`
	Locale       string `json:"locale,omitempty"`
	URL          string `json:"url,omitempty"`
}

type HostedPageResponse struct {
	Success bool              `json:"success,omitempty"`
	Status  int               `json:"status,omitempty"`
	Data    HostedPagePayload `json:"data,omitempty"`
	Error   string            `json:"error,omitempty"`
}

func (c *Client) UpsertHostedPage(sc HostedPagePayload) (response *HostedPageResponse, err error) {
	c.HTTPClient.URL = fmt.Sprintf("%s/%s", c.Config.BaseURL, "hostedpages-srv/hpgroup")
	c.HTTPClient.HTTPMethod = http.MethodPost
	res, err := c.HTTPClient.MakeRequest(sc)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %v", err)
	}
	return response, nil
}

func (c *Client) GetHostedPage(hpGroupName string) (response *HostedPageResponse, err error) {
	c.HTTPClient.URL = fmt.Sprintf("%s/%s/%s", c.Config.BaseURL, "hostedpages-srv/hpgroup", hpGroupName)
	c.HTTPClient.HTTPMethod = http.MethodGet
	res, err := c.HTTPClient.MakeRequest(nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %v", err)
	}
	return response, nil
}

func (c *Client) DeleteHostedPage(hpGroupName string) (response *HostedPageResponse, err error) {
	c.HTTPClient.URL = fmt.Sprintf("%s/%s/%s", c.Config.BaseURL, "hostedpages-srv/hpgroup", hpGroupName)
	c.HTTPClient.HTTPMethod = http.MethodDelete
	res, err := c.HTTPClient.MakeRequest(nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %v", err)
	}
	return response, nil
}
