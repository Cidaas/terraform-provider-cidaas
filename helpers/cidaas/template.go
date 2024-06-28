package cidaas

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/util"
)

type TemplateModel struct {
	ID               string `json:"_id,omitempty"`
	Locale           string `json:"locale,omitempty"`
	TemplateKey      string `json:"templateKey,omitempty"`
	TemplateType     string `json:"templateType,omitempty"`
	Content          string `json:"content,omitempty"`
	Subject          string `json:"subject,omitempty"`
	TemplateOwner    string `json:"templateOwner,omitempty"`
	UsageType        string `json:"usageType,omitempty"`
	ProcessingType   string `json:"processingType,omitempty"`
	VerificationType string `json:"verificationType,omitempty"`
	Language         string `json:"language,omitempty"`
	GroupID          string `json:"group_id,omitempty"`
}

type TemplateResponse struct {
	Success bool          `json:"success,omitempty"`
	Status  int           `json:"status,omitempty"`
	Data    TemplateModel `json:"data,omitempty"`
}

type MasterListResponse struct {
	Success bool         `json:"success,omitempty"`
	Status  int          `json:"status,omitempty"`
	Data    []MasterList `json:"data,omitempty"`
}
type MasterList struct {
	TemplateKey   string         `json:"templateKey,omitempty"`
	Requirement   string         `json:"requirement,omitempty"`
	Enabled       bool           `json:"enabled,omitempty"`
	TemplateTypes []TemplateType `json:"templateTypes,omitempty"`
}

type TemplateType struct {
	TemplateType    string           `json:"templateType,omitempty"`
	ProcessingTypes []ProcessingType `json:"processingTypes,omitempty"`
}

type ProcessingType struct {
	ProcessingType    string             `json:"processingType,omitempty"`
	VerificationTypes []VerificationType `json:"verificationTypes,omitempty"`
}

type VerificationType struct {
	VerificationType string   `json:"verificationType,omitempty"`
	UsageTypes       []string `json:"usageTypes,omitempty"`
}

var _ TemplateService = &Template{}

type Template struct {
	HTTPClient util.HTTPClientInterface
}

type TemplateService interface {
	Upsert(template TemplateModel, isSystemTemplate bool) (response *TemplateResponse, err error)
	Get(template TemplateModel, isSystemTemplate bool) (response *TemplateResponse, err error)
	Delete(templateKey string, templateType string) error
	GetMasterList(groupID string) (*MasterListResponse, error)
}

func NewTemplate(httpClient util.HTTPClientInterface) TemplateService {
	return &Template{HTTPClient: httpClient}
}

func (t *Template) Upsert(template TemplateModel, isSystemTemplate bool) (response *TemplateResponse, err error) {
	url := fmt.Sprintf("%s/%s", t.HTTPClient.GetHost(), "templates-srv/template")
	if !isSystemTemplate {
		url += "/custom"
	}
	t.HTTPClient.SetURL(url)
	t.HTTPClient.SetMethod(http.MethodPost)
	res, err := t.HTTPClient.MakeRequest(template)
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

func (t *Template) Get(template TemplateModel, isSystemTemplate bool) (response *TemplateResponse, err error) {
	url := fmt.Sprintf("%s/%s", t.HTTPClient.GetHost(), "templates-srv/template")
	if isSystemTemplate {
		url += "/find"
	} else {
		url += "/custom/find"
	}
	t.HTTPClient.SetURL(url)
	t.HTTPClient.SetMethod(http.MethodPost)
	res, err := t.HTTPClient.MakeRequest(template)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode == 204 {
		return nil, fmt.Errorf("template not found for the  template_key %s with template type %s and locale %s", template.TemplateKey, template.TemplateType, template.Locale)
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

func (t *Template) Delete(templateKey string, templateType string) error {
	t.HTTPClient.SetURL(fmt.Sprintf("%s/%s/%s/%s", t.HTTPClient.GetHost(), "templates-srv/template/custom", strings.ToUpper(templateKey), strings.ToUpper(templateType)))
	t.HTTPClient.SetMethod(http.MethodDelete)
	res, err := t.HTTPClient.MakeRequest(nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}

func (t *Template) GetMasterList(groupID string) (*MasterListResponse, error) {
	url := fmt.Sprintf("%s/%s/%s", t.HTTPClient.GetHost(), "templates-srv/master/settings", groupID)
	t.HTTPClient.SetURL(url)
	t.HTTPClient.SetMethod(http.MethodGet)
	res, err := t.HTTPClient.MakeRequest(nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var response MasterListResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %w", err)
	}
	return &response, nil
}
