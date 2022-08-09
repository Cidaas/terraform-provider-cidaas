package cidaas

import (
	"context"
	"strconv"
	"time"

	"terraform-provider-cidaas/helper_pkg/cidaas_sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceApp() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAppCreate,
		ReadContext:   resourceAppRead,
		UpdateContext: resourceAppUpdate,
		DeleteContext: resourceAppDelete,

		Schema: map[string]*schema.Schema{

			"client_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"client_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"redirect_uris": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"allowed_logout_urls": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"auth_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"app_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"allow_login_with": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"auto_login_after_register": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},

			"enable_passwordless_auth": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},

			"register_with_login_information": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},

			"hosted_page_group": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"client_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"client_display_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"company_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"company_address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"company_website": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"allowed_scopes": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"response_types": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"grant_types": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"template_group_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"app_attributes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"datatype": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"app_retrieval_status": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"app_retrieval_success": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"app_creation_status": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"app_creation_success": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"app_creation_error": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"app_creation_error_code": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"app_creation_error_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

type appCreationResponse struct {
	Success bool
	Status  int
}

func interfaceArray2StringArray(interfaceArray []interface{}) (stringArray []string) {

	stringArray = make([]string, 0)
	for _, txt := range interfaceArray {
		stringArray = append(stringArray, txt.(string))
	}

	return stringArray
}

func resourceAppCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	cidaas_client := m.(cidaas_sdk.CidaasClient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	var appConfig cidaas_sdk.AppConfig
	appConfig.ClientType = d.Get("client_type").(string)
	appConfig.AllowLoginWith = interfaceArray2StringArray(d.Get("allow_login_with").([]interface{}))
	appConfig.AutoLoginAfterRegister = d.Get("auto_login_after_register").(bool)
	appConfig.EnablePasswordlessAuth = d.Get("enable_passwordless_auth").(bool)
	appConfig.RegisterWithLoginInformation = d.Get("register_with_login_information").(bool)
	appConfig.HostedPageGroup = d.Get("hosted_page_group").(string)
	appConfig.ClientName = d.Get("client_name").(string)
	appConfig.ClientDisplayName = d.Get("client_display_name").(string)
	appConfig.CompanyName = d.Get("company_name").(string)
	appConfig.CompanyAddress = d.Get("company_address").(string)
	appConfig.CompanyWebsite = d.Get("company_website").(string)
	appConfig.AllowedScopes = interfaceArray2StringArray(d.Get("allowed_scopes").([]interface{}))
	appConfig.ResponseTypes = interfaceArray2StringArray(d.Get("response_types").([]interface{}))
	appConfig.GrantTypes = interfaceArray2StringArray(d.Get("grant_types").([]interface{}))
	appConfig.AllowedLogoutUrls = interfaceArray2StringArray(d.Get("allowed_logout_urls").([]interface{}))
	appConfig.RedirectURIS = interfaceArray2StringArray(d.Get("redirect_uris").([]interface{}))
	appConfig.TemplateGroupId = d.Get("template_group_id").(string)
	appConfig.CustomFields = map[string]string{"foo": "bar"}
	appConfig.AdditionalAccessTokenPayload = []string{"foo"}

	appcreationresponse := cidaas_sdk.CreateApp(cidaas_client, appConfig)

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	if err := d.Set("app_creation_success", appcreationresponse.Success); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error Occured while setting app_creation_success to resourceData",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("app_creation_status", int(appcreationresponse.Status)); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error Occured while setting app_creation_status to resourceData",
			Detail:   err.Error(),
		})
		return diags
	}

	error_string := appcreationresponse.Error.Error
	error_code := int(appcreationresponse.Error.Code)
	error_type := appcreationresponse.Error.Type

	if err := d.Set("app_creation_error", error_string); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error Occured while setting app_creation_error to resourceData",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("app_creation_error_code", error_code); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error Occured while setting app_creation_error to resourceData",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("app_creation_error_type", error_type); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error Occured while setting app_creation_error to resourceData",
			Detail:   err.Error(),
		})
		return diags
	}

	if appcreationresponse.Success == false {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create app",
			Detail:   error_string,
		})
		return diags
	}

	if err := d.Set("client_id", appcreationresponse.Data.ClientId); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error Occured while setting app_creation_error to resourceData",
			Detail:   err.Error(),
		})
		return diags
	}

	resourceAppRead(ctx, d, m)

	return diags
}

func resourceAppRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	cidaas_client := m.(cidaas_sdk.CidaasClient)

	client_id := d.Get("client_id").(string)

	appreadresponse := cidaas_sdk.GetApp(cidaas_client, client_id)

	if err := d.Set("app_retrieval_success", appreadresponse.Success); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error Occured while setting app_retrieval_success to resourceData",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("app_retrieval_status", int(appreadresponse.Status)); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error Occured while setting app_retrieval_status to resourceData",
			Detail:   err.Error(),
		})
		return diags
	}

	return diags
}

func resourceAppUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	cidaas_client := m.(cidaas_sdk.CidaasClient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	var appConfig cidaas_sdk.AppConfig

	appConfig.ClientType = d.Get("client_type").(string)
	appConfig.AllowLoginWith = interfaceArray2StringArray(d.Get("allow_login_with").([]interface{}))
	appConfig.AutoLoginAfterRegister = d.Get("auto_login_after_register").(bool)
	appConfig.EnablePasswordlessAuth = d.Get("enable_passwordless_auth").(bool)
	appConfig.RegisterWithLoginInformation = d.Get("register_with_login_information").(bool)
	appConfig.HostedPageGroup = d.Get("hosted_page_group").(string)
	appConfig.ClientName = d.Get("client_name").(string)
	appConfig.ClientDisplayName = d.Get("client_display_name").(string)
	appConfig.CompanyName = d.Get("company_name").(string)
	appConfig.CompanyAddress = d.Get("company_address").(string)
	appConfig.CompanyWebsite = d.Get("company_website").(string)
	appConfig.AllowedScopes = interfaceArray2StringArray(d.Get("allowed_scopes").([]interface{}))
	appConfig.ResponseTypes = interfaceArray2StringArray(d.Get("response_types").([]interface{}))
	appConfig.GrantTypes = interfaceArray2StringArray(d.Get("grant_types").([]interface{}))
	appConfig.AllowedLogoutUrls = interfaceArray2StringArray(d.Get("allowed_logout_urls").([]interface{}))
	appConfig.RedirectURIS = interfaceArray2StringArray(d.Get("redirect_uris").([]interface{}))
	appConfig.TemplateGroupId = d.Get("template_group_id").(string)
	appConfig.ClientId = d.Get("client_id").(string)

	appupdateresponse := cidaas_sdk.UpdateApp(cidaas_client, appConfig)

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "App Update Success",
		Detail:   strconv.FormatBool(appupdateresponse.Success),
	})

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "App Update Status",
		Detail:   strconv.Itoa(appupdateresponse.Status),
	})

	if appupdateresponse.Success == false {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "App Update Failed",
			Detail:   "App Update Failed",
		})
	}

	// if err := d.Set("app_creation_success", appcreationresponse.Success); err != nil {
	// 	diags = append(diags, diag.Diagnostic{
	// 		Severity: diag.Error,
	// 		Summary:  "Error Occured while setting app_creation_success to resourceData",
	// 		Detail:   err.Error(),
	// 	})
	// 	return diags
	// }

	// if err := d.Set("app_creation_status", int(appcreationresponse.Status)); err != nil {
	// 	diags = append(diags, diag.Diagnostic{
	// 		Severity: diag.Error,
	// 		Summary:  "Error Occured while setting app_creation_status to resourceData",
	// 		Detail:   err.Error(),
	// 	})
	// 	return diags
	// }

	// error_string := appupdateresponse.Error.Error
	// error_code := int(appupdateresponse.Error.Code)
	// error_type := appupdateresponse.Error.Type

	// if err := d.Set("app_creation_error", error_string); err != nil {
	// 	diags = append(diags, diag.Diagnostic{
	// 		Severity: diag.Error,
	// 		Summary:  "Error Occured while setting app_creation_error to resourceData",
	// 		Detail:   err.Error(),
	// 	})
	// 	return diags
	// }

	// if err := d.Set("app_creation_error_code", error_code); err != nil {
	// 	diags = append(diags, diag.Diagnostic{
	// 		Severity: diag.Error,
	// 		Summary:  "Error Occured while setting app_creation_error to resourceData",
	// 		Detail:   err.Error(),
	// 	})
	// 	return diags
	// }

	// if err := d.Set("app_creation_error_type", error_type); err != nil {
	// 	diags = append(diags, diag.Diagnostic{
	// 		Severity: diag.Error,
	// 		Summary:  "Error Occured while setting app_creation_error to resourceData",
	// 		Detail:   err.Error(),
	// 	})
	// 	return diags
	// }

	// if appcreationresponse.Success == false {
	// 	diags = append(diags, diag.Diagnostic{
	// 		Severity: diag.Error,
	// 		Summary:  "Unable to create app",
	// 		Detail:   error_string,
	// 	})
	// 	return diags
	// }

	// if err := d.Set("client_id", appcreationresponse.Data.ClientId); err != nil {
	// 	diags = append(diags, diag.Diagnostic{
	// 		Severity: diag.Error,
	// 		Summary:  "Error Occured while setting app_creation_error to resourceData",
	// 		Detail:   err.Error(),
	// 	})
	// 	return diags
	// }

	resourceAppRead(ctx, d, m)

	return diags
}

func resourceAppDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	cidaas_client := m.(cidaas_sdk.CidaasClient)

	client_id := d.Get("client_id").(string)

	appdeleteresponse := cidaas_sdk.DeleteApp(cidaas_client, client_id)

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "App Deletion Success",
		Detail:   strconv.FormatBool(appdeleteresponse.Success),
	})

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "App Deletion Status",
		Detail:   strconv.Itoa(appdeleteresponse.Status),
	})

	if appdeleteresponse.Success == false {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "App Deletion Failed",
			Detail:   "App Deletion Failed",
		})
	}

	return diags
}
