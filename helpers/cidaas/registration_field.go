package cidaas

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/util"
)

type RegistrationFieldResponse struct {
	Success bool                    `json:"success"`
	Status  int                     `json:"status"`
	Data    RegistrationFieldConfig `json:"data"`
}

type AllRegFieldResponse struct {
	Success bool                      `json:"success"`
	Status  int                       `json:"status"`
	Data    []RegistrationFieldConfig `json:"data"`
}

type RegistrationFieldConfig struct {
	Internal                                 bool             `json:"internal"`
	ReadOnly                                 bool             `json:"readOnly"`
	Claimable                                bool             `json:"claimable"`
	Required                                 bool             `json:"required"`
	Unique                                   bool             `json:"unique"`
	IsSearchable                             bool             `json:"isSearchable"`
	OverwriteWithNullValueFromSocialProvider bool             `json:"overwriteWithNullValueFromSocialProvider"`
	ConsentRefs                              []string         `json:"consent_refs,omitempty"`
	Scopes                                   []string         `json:"scopes,omitempty"`
	Enabled                                  bool             `json:"enabled"`
	LocaleTexts                              []*LocaleText    `json:"localeTexts,omitempty"`
	IsGroup                                  bool             `json:"is_group"`
	IsList                                   bool             `json:"is_list"`
	ParentGroupID                            string           `json:"parent_group_id,omitempty"`
	FieldType                                string           `json:"fieldType,omitempty"`
	ID                                       string           `json:"_id,omitempty"`
	FieldKey                                 string           `json:"fieldKey,omitempty"`
	DataType                                 string           `json:"dataType,omitempty"`
	Order                                    int64            `json:"order,omitempty"`
	BaseDataType                             string           `json:"baseDataType,omitempty"`
	FieldDefinition                          *FieldDefinition `json:"fieldDefinition,omitempty"`
	ClassName                                string           `json:"className,omitempty"`
}

type FieldDefinition struct {
	MinLength       *int64     `json:"minLength,omitempty"`
	MaxLength       *int64     `json:"maxLength,omitempty"`
	MinDate         *time.Time `json:"minDate,omitempty"`
	MaxDate         *time.Time `json:"maxDate,omitempty"`
	InitialDate     *time.Time `json:"initialDate,omitempty"`
	InitialDateView string     `json:"initialDateView,omitempty"`
	AttributesKeys  []string   `json:"attributesKeys,omitempty"`
}

type LocaleText struct {
	MinLengthErrorMsg string       `json:"minLength,omitempty"`
	MaxLengthErrorMsg string       `json:"maxLength,omitempty"`
	RequiredMsg       string       `json:"required,omitempty"`
	Language          string       `json:"language,omitempty"`
	Locale            string       `json:"locale,omitempty"`
	Name              string       `json:"name,omitempty"`
	Attributes        []*Attribute `json:"attributes,omitempty"`
	ConsentLabel      *Consent     `json:"consentLabel,omitempty"`
}

type Attribute struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type Consent struct {
	Label     string `json:"label,omitempty"`
	LabelText string `json:"label_text,omitempty"`
}

var _ RegFieldService = &RegField{}

type RegField struct {
	ClientConfig
}

type RegFieldService interface {
	Upsert(rfc RegistrationFieldConfig) (*RegistrationFieldResponse, error)
	Get(fieldKey string) (*RegistrationFieldResponse, error)
	Delete(fieldKey string) error
	GetAll() ([]RegistrationFieldConfig, error)
}

func NewRegField(clientConfig ClientConfig) RegFieldService {
	return &RegField{clientConfig}
}

func (r *RegField) Upsert(rfc RegistrationFieldConfig) (*RegistrationFieldResponse, error) {
	var response RegistrationFieldResponse
	url := fmt.Sprintf("%s/%s", r.BaseURL, "fieldsetup-srv/fields")
	httpClient := util.NewHTTPClient(url, http.MethodPost, r.AccessToken)

	res, err := httpClient.MakeRequest(rfc)
	if err = util.HandleResponseError(res, err); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err = util.ProcessResponse(res, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (r *RegField) Get(fieldKey string) (*RegistrationFieldResponse, error) {
	var response RegistrationFieldResponse
	url := fmt.Sprintf("%s/%s/%s", r.BaseURL, "fieldsetup-srv/fields", fieldKey)
	httpClient := util.NewHTTPClient(url, http.MethodGet, r.AccessToken)

	res, err := httpClient.MakeRequest(nil)
	if err = util.HandleResponseError(res, err); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err = util.ProcessResponse(res, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (r *RegField) Delete(fieldKey string) error {
	url := fmt.Sprintf("%s/%s/%s", r.BaseURL, "fieldsetup-srv/fields", fieldKey)
	httpClient := util.NewHTTPClient(url, http.MethodDelete, r.AccessToken)

	res, err := httpClient.MakeRequest(nil)
	if err = util.HandleResponseError(res, err); err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}

func (r *RegField) GetAll() ([]RegistrationFieldConfig, error) {
	var response AllRegFieldResponse
	url := fmt.Sprintf("%s/%s", r.BaseURL, "registration-setup-srv/fields/list")
	httpClient := util.NewHTTPClient(url, http.MethodGet, r.AccessToken)

	res, err := httpClient.MakeRequest(nil)
	if err = util.HandleResponseError(res, err); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err = util.ProcessResponse(res, &response); err != nil {
		return nil, err
	}
	return response.Data, nil
}
