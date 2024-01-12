package cidaas

import (
	"context"

	"terraform-provider-cidaas/helper/cidaas"
	"terraform-provider-cidaas/helper/util"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRegistrationField() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRegistrationFieldCreate,
		ReadContext:   resourceRegistrationFieldRead,
		UpdateContext: resourceRegistrationFieldUpdate,
		DeleteContext: resourceRegistrationFieldDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"required": {
				Type:     schema.TypeBool,
				Required: true,
			},

			"internal": {
				Type:     schema.TypeBool,
				Required: true,
			},

			"claimable": {
				Type:     schema.TypeBool,
				Required: true,
			},

			"scopes": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},

			"is_group": {
				Type:     schema.TypeBool,
				Required: true,
			},

			"is_list": {
				Type:     schema.TypeBool,
				Computed: true,
			},

			"parent_group_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			"field_type": {
				Type:     schema.TypeString,
				Required: true,
			},

			"data_type": {
				Type:     schema.TypeString,
				Required: true,
			},

			"field_key": {
				Type:     schema.TypeString,
				Required: true,
			},

			"read_only": {
				Type:     schema.TypeBool,
				Required: true,
			},

			"order": {
				Type:     schema.TypeInt,
				Required: true,
			},

			"locale_text_locale": {
				Type:     schema.TypeString,
				Required: true,
			},

			"locale_text_name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"locale_text_language": {
				Type:     schema.TypeString,
				Required: true,
			},
			"locale_text_min_length": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"locale_text_max_length": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"min_length_error_msg": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"max_length_error_msg": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"required_msg": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"registration_field_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"base_data_type": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"app_attributes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"datatype": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

type registrationFeildCreationResponse struct {
	Success bool
	Status  int
}

func resourceRegistrationFieldCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	var registrationFieldConfig cidaas.RegistrationFieldConfig
	cidaas_client := m.(cidaas.CidaasClient)

	registrationFieldConfig.ParentGroupId = d.Get("parent_group_id").(string)
	registrationFieldConfig.Scopes = util.InterfaceArray2StringArray(d.Get("scopes").([]interface{}))
	registrationFieldConfig.DataType = d.Get("data_type").(string)
	registrationFieldConfig.FieldKey = d.Get("field_key").(string)
	registrationFieldConfig.Required = d.Get("required").(bool)
	registrationFieldConfig.IsGroup = d.Get("is_group").(bool)
	registrationFieldConfig.Enabled = d.Get("enabled").(bool)
	registrationFieldConfig.ReadOnly = d.Get("read_only").(bool)
	registrationFieldConfig.Internal = d.Get("internal").(bool)
	registrationFieldConfig.Claimable = d.Get("claimable").(bool)
	registrationFieldConfig.Order = d.Get("order").(int)
	registrationFieldConfig.FieldType = d.Get("field_type").(string)
	registrationFieldConfig.BaseDataType = "string"
	registrationFieldConfig.LocaleText.Locale = d.Get("locale_text_locale").(string)
	registrationFieldConfig.LocaleText.Name = d.Get("locale_text_name").(string)
	registrationFieldConfig.LocaleText.Language = d.Get("locale_text_language").(string)
	registrationFieldConfig.LocaleText.MinLengthErrorMsg = d.Get("min_length_error_msg").(string)
	registrationFieldConfig.LocaleText.MaxLengthErrorMsg = d.Get("max_length_error_msg").(string)
	registrationFieldConfig.LocaleText.RequiredMsg = d.Get("required_msg").(string)
	registrationFieldConfig.FieldDefinition.Locale = d.Get("locale_text_locale").(string)
	registrationFieldConfig.FieldDefinition.Name = d.Get("locale_text_name").(string)
	registrationFieldConfig.FieldDefinition.Language = d.Get("locale_text_language").(string)
	registrationFieldConfig.FieldDefinition.MinLength = d.Get("locale_text_min_length").(int)
	registrationFieldConfig.FieldDefinition.MaxLength = d.Get("locale_text_max_length").(int)

	if registrationFieldConfig.FieldDefinition.MinLength > 0 {
		if registrationFieldConfig.FieldDefinition.MaxLength <= 0 {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "locale_text_max_length can not be empty or less than equal to 0",
			})
			return diags
		}

		if registrationFieldConfig.FieldDefinition.MinLength > registrationFieldConfig.FieldDefinition.MaxLength {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "locale_text_min_length can not be greater than locale_text_max_length",
			})
			return diags
		}
	}

	if registrationFieldConfig.FieldDefinition.MinLength > 0 && registrationFieldConfig.LocaleText.MinLengthErrorMsg == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "min_length_error_msg can not be empty when locale_text_min_length is greater than 0",
		})
		return diags
	}

	if registrationFieldConfig.FieldDefinition.MaxLength > 0 && registrationFieldConfig.LocaleText.MaxLengthErrorMsg == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "max_length_error_msg can not be empty when locale_text_max_length is greater than 0",
		})
		return diags
	}

	if registrationFieldConfig.Required {
		if registrationFieldConfig.LocaleText.RequiredMsg == "" {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "required_msg can not be empty when required is true",
			})
			return diags
		}
	}

	response, err := cidaas_client.CreateRegistrationField(registrationFieldConfig)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "failed to create registration field",
			Detail:   err.Error(),
		})
		return diags
	}
	if err := d.Set("registration_field_id", response.Data.Id); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "error while settiing registration_field_id",
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId(d.Get("field_key").(string))
	resourceRegistrationFieldRead(ctx, d, m)
	return diags
}

func resourceRegistrationFieldRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	cidaas_client := m.(cidaas.CidaasClient)
	registration_field_key := d.Id()
	response, err := cidaas_client.GetRegistrationField(registration_field_key)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "failed to read registration field",
			Detail:   err.Error(),
		})
		return diags
	}
	d.Set("enabled", response.Data.Enabled)
	d.Set("field_key", response.Data.FieldKey)
	d.Set("parent_group_id", response.Data.ParentGroupId)
	d.Set("is_group", response.Data.IsGroup)
	d.Set("data_type", response.Data.DataType)
	d.Set("required", response.Data.Required)
	d.Set("read_only", response.Data.ReadOnly)
	d.Set("internal", response.Data.Internal)
	d.Set("scopes", response.Data.Scopes)
	d.Set("claimable", response.Data.Claimable)
	d.Set("order", response.Data.Order)
	d.Set("field_type", response.Data.FieldType)
	d.Set("registration_field_id", response.Data.Id)
	d.Set("base_data_type", response.Data.BaseDataType)
	if len(response.Data.LocaleText) > 0 {
		d.Set("locale_text_locale", response.Data.LocaleText[0]["locale"])
		d.Set("locale_text_name", response.Data.LocaleText[0]["name"])
		d.Set("locale_text_language", response.Data.LocaleText[0]["language"])
		d.Set("min_length_error_msg", response.Data.LocaleText[0]["minLength"])
		d.Set("max_length_error_msg", response.Data.LocaleText[0]["maxLength"])
		d.Set("required_msg", response.Data.LocaleText[0]["required"])
	}
	d.Set("locale_text_locale", response.Data.FieldDefinition.Locale)
	d.Set("locale_text_name", response.Data.FieldDefinition.Name)
	d.Set("locale_text_language", response.Data.FieldDefinition.Language)
	d.Set("locale_text_min_length", response.Data.FieldDefinition.MinLength)
	d.Set("locale_text_max_length", response.Data.FieldDefinition.MaxLength)

	return diags
}

func resourceRegistrationFieldUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	cidaas_client := m.(cidaas.CidaasClient)
	var rfg cidaas.RegistrationFieldConfig

	rfg.ParentGroupId = d.Get("parent_group_id").(string)
	rfg.Scopes = util.InterfaceArray2StringArray(d.Get("scopes").([]interface{}))
	rfg.DataType = d.Get("data_type").(string)
	rfg.FieldKey = d.Get("field_key").(string)
	rfg.Required = d.Get("required").(bool)
	rfg.IsGroup = d.Get("is_group").(bool)
	rfg.Enabled = d.Get("enabled").(bool)
	rfg.ReadOnly = d.Get("read_only").(bool)
	rfg.Internal = d.Get("internal").(bool)
	rfg.Claimable = d.Get("claimable").(bool)
	rfg.Order = d.Get("order").(int)
	rfg.FieldType = d.Get("field_type").(string)
	rfg.BaseDataType = d.Get("base_data_type").(string)
	rfg.Id = d.Get("registration_field_id").(string)
	rfg.LocaleText.Locale = d.Get("locale_text_locale").(string)
	rfg.LocaleText.Name = d.Get("locale_text_name").(string)
	rfg.LocaleText.Language = d.Get("locale_text_language").(string)
	rfg.LocaleText.MinLengthErrorMsg = d.Get("min_length_error_msg").(string)
	rfg.LocaleText.MaxLengthErrorMsg = d.Get("max_length_error_msg").(string)
	rfg.LocaleText.RequiredMsg = d.Get("required_msg").(string)
	rfg.FieldDefinition.Locale = d.Get("locale_text_locale").(string)
	rfg.FieldDefinition.Name = d.Get("locale_text_name").(string)
	rfg.FieldDefinition.Language = d.Get("locale_text_language").(string)
	rfg.FieldDefinition.MinLength = d.Get("locale_text_min_length").(int)
	rfg.FieldDefinition.MaxLength = d.Get("locale_text_max_length").(int)

	if rfg.FieldDefinition.MinLength > 0 {
		if rfg.FieldDefinition.MaxLength <= 0 {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "locale_text_max_length can not be empty or less than equal to 0",
			})
			return diags
		}

		if rfg.FieldDefinition.MinLength > rfg.FieldDefinition.MaxLength {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "locale_text_min_length can not be greater than locale_text_max_length",
			})
			return diags
		}
	}

	if rfg.FieldDefinition.MinLength > 0 && rfg.LocaleText.MinLengthErrorMsg == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "min_length_error_msg can not be empty when locale_text_min_length is greater than 0",
		})
		return diags
	}

	if rfg.FieldDefinition.MaxLength > 0 && rfg.LocaleText.MaxLengthErrorMsg == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "max_length_error_msg can not be empty when locale_text_max_length is greater than 0",
		})
		return diags
	}

	if rfg.Required {
		if rfg.LocaleText.RequiredMsg == "" {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "required_msg can not be empty when required is true",
			})
			return diags
		}
	}

	_, err := cidaas_client.UpdateRegistrationField(rfg)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "failed to update registration field",
			Detail:   err.Error(),
		})
	}
	d.SetId(d.Get("field_key").(string))
	return diags
}

func resourceRegistrationFieldDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	cidaas_client := m.(cidaas.CidaasClient)
	registration_field_key := d.Get("field_key").(string)
	_, err := cidaas_client.DeleteRegistrationField(registration_field_key)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "failed to delete registration field",
			Detail:   err.Error(),
		})
	}
	return diags
}
