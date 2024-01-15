package cidaas

import (
	"encoding/json"
	"fmt"
	"terraform-provider-cidaas/helper/util"
)

type RegistrationFieldResponse struct {
	Success bool                    `json:"success"`
	Status  int                     `json:"status"`
	Data    RegistrationFieldConfig `json:"data"`
}

type GetRegistrationFieldResponse struct {
	Success bool                       `json:"success"`
	Status  int                        `json:"status"`
	Data    GetRegistrationFieldConfig `json:"data"`
}

type GetRegistrationFieldConfig struct {
	Internal        bool                     `json:"internal,omitempty"`
	ReadOnly        bool                     `json:"readOnly,omitempty"`
	Claimable       bool                     `json:"claimable,omitempty"`
	Required        bool                     `json:"required,omitempty"`
	Scopes          []string                 `json:"scopes,omitempty"`
	Enabled         bool                     `json:"enabled,omitempty"`
	LocaleText      []map[string]interface{} `json:"localeText,omitempty"`
	IsGroup         bool                     `json:"is_group,omitempty"`
	IsList          bool                     `json:"is_list,omitempty"`
	ParentGroupId   string                   `json:"parent_group_id,omitempty"`
	FieldType       string                   `json:"fieldType,omitempty"`
	Id              string                   `json:"_id,omitempty"`
	FieldKey        string                   `json:"fieldKey,omitempty"`
	DataType        string                   `json:"dataType,omitempty"`
	Order           int                      `json:"order,omitempty"`
	BaseDataType    string                   `json:"baseDataType,omitempty"`
	FieldDefinition FieldDefinition          `json:"fieldDefinition,omitempty"`
}

type RegistrationFieldConfig struct {
	Internal        bool            `json:"internal,omitempty"`
	ReadOnly        bool            `json:"readOnly,omitempty"`
	Claimable       bool            `json:"claimable,omitempty"`
	Required        bool            `json:"required,omitempty"`
	Scopes          []string        `json:"scopes,omitempty"`
	Enabled         bool            `json:"enabled,omitempty"`
	LocaleText      LocaleText      `json:"localeText,omitempty"`
	IsGroup         bool            `json:"is_group,omitempty"`
	IsList          bool            `json:"is_list,omitempty"`
	ParentGroupId   string          `json:"parent_group_id,omitempty"`
	FieldType       string          `json:"fieldType,omitempty"`
	Id              string          `json:"_id,omitempty"`
	FieldKey        string          `json:"fieldKey,omitempty"`
	DataType        string          `json:"dataType,omitempty"`
	Order           int             `json:"order,omitempty"`
	BaseDataType    string          `json:"baseDataType,omitempty"`
	FieldDefinition FieldDefinition `json:"fieldDefinition,omitempty"`
}

type FieldDefinition struct {
	MinLength int    `json:"minLength,omitempty"`
	MaxLength int    `json:"maxLength,omitempty"`
	Language  string `json:"language,omitempty"`
	Locale    string `json:"locale,omitempty"`
	Name      string `json:"name,omitempty"`
}

type LocaleText struct {
	MinLengthErrorMsg string `json:"minLength,omitempty"`
	MaxLengthErrorMsg string `json:"maxLength,omitempty"`
	RequiredMsg       string `json:"required,omitempty"`
	Language          string `json:"language,omitempty"`
	Locale            string `json:"locale,omitempty"`
	Name              string `json:"name,omitempty"`
}

func (c *CidaasClient) CreateRegistrationField(rfg RegistrationFieldConfig) (response *RegistrationFieldResponse, err error) {
	url := c.BaseUrl + "/fieldsetup-srv/fields"
	h := util.HttpClient{
		Token: c.TokenData.AccessToken,
	}
	res, err := h.Post(url, rfg)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %v", err)
	}
	return response, nil
}

func (c *CidaasClient) UpdateRegistrationField(rfg RegistrationFieldConfig) (response *RegistrationFieldResponse, err error) {
	url := c.BaseUrl + "/fieldsetup-srv/fields"
	h := util.HttpClient{
		Token: c.TokenData.AccessToken,
	}
	res, err := h.Post(url, rfg)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %v", err)
	}
	return response, nil
}

func (c *CidaasClient) GetRegistrationField(key string) (response *GetRegistrationFieldResponse, err error) {
	url := c.BaseUrl + "/registration-setup-srv/fields/flat/field/" + key
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

func (c *CidaasClient) DeleteRegistrationField(key string) (response *RegistrationFieldResponse, err error) {
	url := c.BaseUrl + "/registration-setup-srv/fields/" + key
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

func ValidateRequest(registrationFieldConfig RegistrationFieldConfig) (bool, string) {
	if registrationFieldConfig.FieldDefinition.MinLength > 0 {
		if registrationFieldConfig.FieldDefinition.MaxLength <= 0 {
			return false, "locale_text_max_length can not be empty or less than equal to 0"
		}

		if registrationFieldConfig.FieldDefinition.MinLength > registrationFieldConfig.FieldDefinition.MaxLength {
			return false, "locale_text_min_length can not be greater than locale_text_max_length"
		}
	}

	if registrationFieldConfig.FieldDefinition.MinLength > 0 && registrationFieldConfig.LocaleText.MinLengthErrorMsg == "" {
		return false, "min_length_error_msg can not be empty when locale_text_min_length is greater than 0"
	}

	if registrationFieldConfig.FieldDefinition.MaxLength > 0 && registrationFieldConfig.LocaleText.MaxLengthErrorMsg == "" {
		return false, "max_length_error_msg can not be empty when locale_text_max_length is greater than 0"
	}

	if registrationFieldConfig.Required {
		if registrationFieldConfig.LocaleText.RequiredMsg == "" {
			return false, "required_msg can not be empty when required is true"
		}
	}
	return true, ""
}
