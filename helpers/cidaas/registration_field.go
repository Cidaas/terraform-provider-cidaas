package cidaas

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	ParentGroupID   string                   `json:"parent_group_id,omitempty"`
	FieldType       string                   `json:"fieldType,omitempty"`
	ID              string                   `json:"_id,omitempty"`
	FieldKey        string                   `json:"fieldKey,omitempty"`
	DataType        string                   `json:"dataType,omitempty"`
	Order           int                      `json:"order,omitempty"`
	BaseDataType    string                   `json:"baseDataType,omitempty"`
	FieldDefinition FieldDefinition          `json:"fieldDefinition,omitempty"`
}

type RegistrationFieldConfig struct {
	Internal        bool            `json:"internal"`
	ReadOnly        bool            `json:"readOnly"`
	Claimable       bool            `json:"claimable"`
	Required        bool            `json:"required"`
	Scopes          []string        `json:"scopes,omitempty"`
	Enabled         bool            `json:"enabled"`
	LocaleText      LocaleText      `json:"localeText,omitempty"`
	IsGroup         bool            `json:"is_group"`
	IsList          bool            `json:"is_list"`
	ParentGroupID   string          `json:"parent_group_id,omitempty"`
	FieldType       string          `json:"fieldType,omitempty"`
	ID              string          `json:"_id,omitempty"`
	FieldKey        string          `json:"fieldKey,omitempty"`
	DataType        string          `json:"dataType,omitempty"`
	Order           int             `json:"order,omitempty"`
	BaseDataType    string          `json:"baseDataType,omitempty"`
	FieldDefinition FieldDefinition `json:"fieldDefinition,omitempty"`
	ClassName       string          `json:"className,omitempty"`
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

func (c *Client) UpsertRegistrationField(rfg RegistrationFieldConfig) (response *RegistrationFieldResponse, err error) {
	c.HTTPClient.URL = fmt.Sprintf("%s/%s", c.Config.BaseURL, "fieldsetup-srv/field")
	c.HTTPClient.HTTPMethod = http.MethodPost
	res, err := c.HTTPClient.MakeRequest(rfg)
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

func (c *Client) GetRegistrationField(key string) (response *GetRegistrationFieldResponse, err error) {
	c.HTTPClient.URL = fmt.Sprintf("%s/%s/%s", c.Config.BaseURL, "registration-setup-srv/fields/flat/field", key)
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

func (c *Client) DeleteRegistrationField(key string) (response *RegistrationFieldResponse, err error) {
	c.HTTPClient.URL = fmt.Sprintf("%s/%s/%s", c.Config.BaseURL, "registration-setup-srv/fields", key)
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
