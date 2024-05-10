package cidaas

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
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

func (c *Client) UpsertTemplate(template Template) (response *TemplateResponse, err error) {
	c.HTTPClient.URL = fmt.Sprintf("%s/%s", c.Config.BaseURL, "templates-srv/template/custom")
	c.HTTPClient.HTTPMethod = http.MethodPost
	res, err := c.HTTPClient.MakeRequest(template)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	// handle empty response body
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read template response body")
	}
	bodyString := string(bodyBytes)
	if bodyString == "" {
		return nil, fmt.Errorf("response code %+v with empty response body", res.StatusCode)
	}
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body %s, status code %+v, error %s", bodyString, res.StatusCode, err.Error())
	}
	return response, nil
}

func (c *Client) GetTemplate(template Template) (response *TemplateResponse, err error) {
	c.HTTPClient.URL = fmt.Sprintf("%s/%s", c.Config.BaseURL, "templates-srv/template/custom/find")
	c.HTTPClient.HTTPMethod = http.MethodPost
	res, err := c.HTTPClient.MakeRequest(template)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode == 204 {
		return nil, fmt.Errorf("template not found for the templateKey %s", template.TemplateKey)
	}
	// handle empty response body
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read template response body")
	}
	bodyString := string(bodyBytes)
	if bodyString == "" {
		return nil, fmt.Errorf("response code %+v with empty response body", res.StatusCode)
	}
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body %s, status code %+v, error %s", bodyString, res.StatusCode, err.Error())
	}
	return response, nil
}

func (c *Client) DeleteTemplate(template Template) error {
	c.HTTPClient.URL = fmt.Sprintf("%s/%s/%s/%s", c.Config.BaseURL, "templates-srv/template/custom", strings.ToUpper(template.TemplateKey), strings.ToUpper(template.TemplateType))
	c.HTTPClient.HTTPMethod = http.MethodDelete
	_, err := c.HTTPClient.MakeRequest(nil)
	if err != nil {
		return err
	}
	return nil
}

func PrepareTemplate(id string) (template Template) {
	split_id := strings.Split(id, "_")

	template.TemplateKey = strings.TrimSuffix(id, fmt.Sprintf("_%s_%s", split_id[len(split_id)-2], split_id[len(split_id)-1]))
	template.TemplateType = split_id[len(split_id)-2]
	template.Locale = split_id[len(split_id)-1]

	return template
}
