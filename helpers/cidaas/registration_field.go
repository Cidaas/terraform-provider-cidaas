package cidaas

import (
	"encoding/json"
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

type RegistrationFieldConfig struct {
	Internal                                 bool             `json:"internal"`
	ReadOnly                                 bool             `json:"readOnly"`
	Claimable                                bool             `json:"claimable"`
	Required                                 bool             `json:"required"`
	Unique                                   bool             `json:"unique"`
	IsSearchable                             bool             `json:"isSearchable"`
	OverwriteWithNullValueFromSocialProvider bool             `json:"overwriteWithNullValueFromSocialProvider"`
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
	HTTPClient util.HTTPClientInterface
}

type RegFieldService interface {
	Upsert(rfc RegistrationFieldConfig) (*RegistrationFieldResponse, error)
	Get(fieldKey string) (*RegistrationFieldResponse, error)
	Delete(fieldKey string) error
}

func NewRegField(httpClient util.HTTPClientInterface) RegFieldService {
	return &RegField{HTTPClient: httpClient}
}

func (rf *RegField) Upsert(rfc RegistrationFieldConfig) (*RegistrationFieldResponse, error) {
	rf.HTTPClient.SetURL(fmt.Sprintf("%s/%s", rf.HTTPClient.GetHost(), "fieldsetup-srv/fields"))
	rf.HTTPClient.SetMethod(http.MethodPost)
	res, err := rf.HTTPClient.MakeRequest(rfc)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var response RegistrationFieldResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %w", err)
	}
	return &response, nil
}

func (r *RegField) Get(fieldKey string) (*RegistrationFieldResponse, error) {
	r.HTTPClient.SetURL(fmt.Sprintf("%s/%s/%s", r.HTTPClient.GetHost(), "fieldsetup-srv/fields", fieldKey))
	r.HTTPClient.SetMethod(http.MethodGet)
	res, err := r.HTTPClient.MakeRequest(nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var response RegistrationFieldResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %w", err)
	}
	return &response, nil
}

func (r *RegField) Delete(fieldKey string) error {
	r.HTTPClient.SetURL(fmt.Sprintf("%s/%s/%s", r.HTTPClient.GetHost(), "fieldsetup-srv/fields", fieldKey))
	r.HTTPClient.SetMethod(http.MethodDelete)
	res, err := r.HTTPClient.MakeRequest(nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}
