package cidaas

import (
	"encoding/json"
	"fmt"
	"strings"
	"terraform-provider-cidaas/helper/util"
)

type Template struct {
	ID            string `json:"_id,omitempty"`
	Locale        string `json:"locale,omitempty"`
	TemplateKey   string `json:"templateKey,omitempty"`
	TemplateType  string `json:"templateType,omitempty"`
	Content       string `json:"content,omitempty"`
	Subject       string `json:"subject,omitempty"`
	TemplateOwner string `json:"templateOwner,omitempty"`
	UsageType     string `json:"usageType,omitempty"`
	Language      string `json:"language,omitempty"`
	GroupId       string `json:"group_id,omitempty"`
}

type TemplateResponse struct {
	Success bool     `json:"success,omitempty"`
	Status  int      `json:"status,omitempty"`
	Data    Template `json:"data,omitempty"`
}

type DeleteTemplateResponse struct {
	Success bool `json:"success,omitempty"`
	Status  int  `json:"status,omitempty"`
	Data    bool `json:"data,omitempty"`
}

func (c *CidaasClient) UpsertTemplate(sc Template) (response *TemplateResponse, err error) {
	url := c.BaseUrl + "/templates-srv/template/custom"
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

func (c *CidaasClient) GetTemplate(template Template) (response *TemplateResponse, err error) {
	url := c.BaseUrl + "/templates-srv/template/custom/find"
	h := util.HttpClient{
		Token: c.TokenData.AccessToken,
	}
	res, err := h.Post(url, template)
	if res.StatusCode == 204 {
		return nil, fmt.Errorf("204 no content")
	}
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %v", err)
	}
	return response, nil
}

func (c *CidaasClient) DeleteTemplate(template Template) (response *DeleteScopeResponse, err error) {
	url := c.BaseUrl + "/templates-srv/template/custom/" + strings.ToUpper(template.TemplateKey) + "/" + strings.ToUpper(template.TemplateType)
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

func PrepareTemplate(id string) (template Template) {
	switch split_id := strings.Split(id, "_"); split_id[len(split_id)-1] {
	case "EMAIL":
		template.TemplateKey = strings.TrimSuffix(id, "_EMAIL")
		template.TemplateType = "EMAIL"
	case "SMS":
		template.TemplateKey = strings.TrimSuffix(id, "_SMS")
		template.TemplateType = "SMS"
	case "IVR":
		template.TemplateKey = strings.TrimSuffix(id, "_IVR")
		template.TemplateType = "IVR"
	case "PUSH":
		template.TemplateKey = strings.TrimSuffix(id, "_PUSH")
		template.TemplateType = "PUSH"
	}
	return template
}
