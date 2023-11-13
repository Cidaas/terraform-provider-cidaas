package cidaas

import (
	"encoding/json"
	"fmt"
	"strings"
	"terraform-provider-cidaas/helper/util"
)

type HostedPagePayload struct {
	ID            string       `json:"_id,omitempty"`
	GroupOwner    string       `json:"groupOwner,omitempty"`
	DefaultLocale string       `json:"default_locale,omitempty"`
	HostedPages   []HostedPage `json:"hosted_pages,omitempty"`
}

type HostedPage struct {
	HostedPageId string `json:"hosted_page_id,omitempty"`
	Content      string `json:"content,omitempty"`
	Locale       string `json:"locale,omitempty"`
	Url          string `json:"url,omitempty"`
}

type HostedPageResponse struct {
	Success bool              `json:"success,omitempty"`
	Status  int               `json:"status,omitempty"`
	Data    HostedPagePayload `json:"data,omitempty"`
	Error   string            `json:"error,omitempty"`
}

func (c *CidaasClient) CreateOrUpdateHostedPage(sc HostedPagePayload) (response *HostedPageResponse, err error) {
	url := c.BaseUrl + "/hostedpages-srv/hpgroup"
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

func (c *CidaasClient) GetHostedPage(hp_group_name string) (response *HostedPageResponse, err error) {
	url := c.BaseUrl + "/hostedpages-srv/hpgroup/" + strings.ToLower(hp_group_name)
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

func (c *CidaasClient) DeleteHostedPage(hp_group_name string) (response *HostedPageResponse, err error) {
	url := c.BaseUrl + "/hostedpages-srv/hpgroup/" + strings.ToLower(hp_group_name)
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
