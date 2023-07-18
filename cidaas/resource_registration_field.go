package cidaas

import (
	"context"
	"strconv"

	"terraform-provider-cidaas/helper_pkg/cidaas_sdk"

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
			"registration_field_retrieval_status": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"registration_field_retrieval_success": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"registration_field_creation_status": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"registration_field_creation_success": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"registration_field_creation_error": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"registration_field_creation_error_code": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"registration_field_creation_error_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

type registrationFeildCreationResponse struct {
	Success bool
	Status  int
}

func resourceRegistrationFieldCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	cidaas_client := m.(cidaas_sdk.CidaasClient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	var registrationFieldConfig cidaas_sdk.RegistrationFieldConfig

	registrationFieldConfig.ParentGroupId = d.Get("parent_group_id").(string)
	registrationFieldConfig.Scopes = interfaceArray2StringArray(d.Get("scopes").([]interface{}))
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
	registrationFieldConfig.LocaleText = make(map[string]interface{})
	registrationFieldConfig.BaseDataType = "string"
	registrationFieldConfig.LocaleText["locale"] = d.Get("locale_text_locale").(string)
	registrationFieldConfig.LocaleText["name"] = d.Get("locale_text_name").(string)
	registrationFieldConfig.LocaleText["language"] = d.Get("locale_text_language").(string)

	registrationFieldcreationresponse := cidaas_sdk.CreateRegistrationField(cidaas_client, registrationFieldConfig)

	if err := d.Set("registration_field_creation_success", registrationFieldcreationresponse.Success); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error Occured while setting registration_field_creation_success to resourceData",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("registration_field_creation_status", int(registrationFieldcreationresponse.Status)); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error Occured while setting registration_field_creation_status to resourceData",
			Detail:   err.Error(),
		})
		return diags
	}

	error_string := registrationFieldcreationresponse.Error.Error
	error_code := int(registrationFieldcreationresponse.Error.Code)
	error_type := registrationFieldcreationresponse.Error.Type

	if err := d.Set("registration_field_creation_error", error_string); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error Occured while setting registration_field_creation_error to resourceData",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("registration_field_creation_error_code", error_code); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error Occured while setting registration_field_creation_error to resourceData",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("registration_field_creation_error_type", error_type); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error Occured while setting registration_field_creation_error to resourceData",
			Detail:   err.Error(),
		})
		return diags
	}

	if registrationFieldcreationresponse.Success == false {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create registration",
			Detail:   error_string,
		})
		return diags
	}

	if err := d.Set("registration_field_id", registrationFieldcreationresponse.Data.Id); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error Occured while setting app_creation_error to resourceData",
			Detail:   err.Error(),
		})
		return diags
	}

	d.SetId(d.Get("field_key").(string))
	resourceRegistrationFieldRead(ctx, d, m)

	return diags
}

func resourceRegistrationFieldRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	cidaas_client := m.(cidaas_sdk.CidaasClient)

	registration_field_key := d.Id()

	registrationFieldreadresponse := cidaas_sdk.GetRegistrationField(cidaas_client, registration_field_key)

	d.Set("enabled", registrationFieldreadresponse.Data.Enabled)
	d.Set("field_key", registrationFieldreadresponse.Data.FieldKey)
	d.Set("parent_group_id", registrationFieldreadresponse.Data.ParentGroupId)
	d.Set("is_group", registrationFieldreadresponse.Data.IsGroup)
	d.Set("data_type", registrationFieldreadresponse.Data.DataType)
	d.Set("required", registrationFieldreadresponse.Data.Required)
	d.Set("read_only", registrationFieldreadresponse.Data.ReadOnly)
	d.Set("internal", registrationFieldreadresponse.Data.Internal)
	d.Set("scopes", registrationFieldreadresponse.Data.Scopes)
	d.Set("claimable", registrationFieldreadresponse.Data.Claimable)
	d.Set("order", registrationFieldreadresponse.Data.Order)
	d.Set("field_type", registrationFieldreadresponse.Data.FieldType)
	d.Set("locale_text_locale", registrationFieldreadresponse.Data.LocaleText[0]["locale"])
	d.Set("locale_text_name", registrationFieldreadresponse.Data.LocaleText[0]["name"])
	d.Set("locale_text_language", registrationFieldreadresponse.Data.LocaleText[0]["language"])
	d.Set("registration_field_id", registrationFieldreadresponse.Data.Id)
	d.Set("base_data_type", registrationFieldreadresponse.Data.BaseDataType)

	if err := d.Set("registration_field_retrieval_success", registrationFieldreadresponse.Success); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error Occured while setting registration_field_retrieval_success to resourceData",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("registration_field_retrieval_status", int(registrationFieldreadresponse.Status)); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error Occured while setting registrationField_retrieval_status to resourceData",
			Detail:   err.Error(),
		})
		return diags
	}
	return diags
}

func resourceRegistrationFieldUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	cidaas_client := m.(cidaas_sdk.CidaasClient)

	// // Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	var registrationFieldConfig cidaas_sdk.RegistrationFieldConfig

	registrationFieldConfig.ParentGroupId = d.Get("parent_group_id").(string)
	registrationFieldConfig.Scopes = interfaceArray2StringArray(d.Get("scopes").([]interface{}))
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
	registrationFieldConfig.LocaleText = make(map[string]interface{})
	registrationFieldConfig.BaseDataType = d.Get("base_data_type").(string)
	registrationFieldConfig.Id = d.Get("registration_field_id").(string)
	registrationFieldConfig.LocaleText["locale"] = d.Get("locale_text_locale").(string)
	registrationFieldConfig.LocaleText["name"] = d.Get("locale_text_name").(string)
	registrationFieldConfig.LocaleText["language"] = d.Get("locale_text_language").(string)

	registrationFieldupdateresponse := cidaas_sdk.UpdateRegistrationField(cidaas_client, registrationFieldConfig)

	d.SetId(d.Get("field_key").(string))

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Registration Field Update Success",
		Detail:   strconv.FormatBool(registrationFieldupdateresponse.Success),
	})

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Registration field Update Status",
		Detail:   strconv.Itoa(registrationFieldupdateresponse.Status),
	})

	if registrationFieldupdateresponse.Success == false {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Registration Field Update Failed",
			Detail:   "Registration Field Update Failed",
		})
	}
	return diags
}

func resourceRegistrationFieldDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	cidaas_client := m.(cidaas_sdk.CidaasClient)

	registration_field_key := d.Get("field_key").(string)

	registrationFielddeleteresponse := cidaas_sdk.DeleteRegistrationField(cidaas_client, registration_field_key)

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Registration Field Deletion Success",
		Detail:   strconv.FormatBool(registrationFielddeleteresponse.Success),
	})

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Registration Field Deletion Status",
		Detail:   strconv.Itoa(registrationFielddeleteresponse.Status),
	})

	if registrationFielddeleteresponse.Success == false {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Registration Field Deletion Failed",
			Detail:   "Registration Field Deletion Failed",
		})
	}

	return diags
}
