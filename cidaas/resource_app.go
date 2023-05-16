package cidaas

import (
	"context"
	"fmt"
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
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"client_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"accent_color": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"primary_color": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"media_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"content_align": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"allow_login_with": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"redirect_uris": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"allowed_logout_urls": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"enable_deduplication": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"auto_login_after_register": {
				Type:     schema.TypeBool,
				Required: true,
			},

			"enable_passwordless_auth": {
				Type:     schema.TypeBool,
				Required: true,
			},

			"register_with_login_information": {
				Type:     schema.TypeBool,
				Required: true,
			},

			"allow_disposable_email": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"validate_phone_number": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"fds_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"hosted_page_group": {
				Type:     schema.TypeString,
				Required: true,
			},

			"client_name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"client_display_name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"company_name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"company_address": {
				Type:     schema.TypeString,
				Required: true,
			},

			"company_website": {
				Type:     schema.TypeString,
				Required: true,
			},

			"allowed_scopes": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"response_types": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"grant_types": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"login_providers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"additional_access_token_payload": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"required_fields": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"application_meta_data": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"is_hybrid_app": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"allowed_web_origins": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"allowed_origins": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"mobile_settings": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"team_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"bundle_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"package_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"key_hash": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"refresh_token_lifetime_in_seconds": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"template_group_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"client_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"custom_provider_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"app_retrieval_status": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"app_retrieval_success": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"app_creation_status": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"app_creation_success": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"app_creation_error": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"app_creation_error_code": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"app_creation_error_type": {
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

func arrayOfInterface(interfaceArray []interface{}) (stringArray []map[string]string) {

	for _, txt := range interfaceArray {
		elements := map[string]string{
			"scope_name": txt.(string),
		}
		stringArray = append(stringArray, elements)
	}

	return stringArray
}

func resourceAppCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	cidaas_client := m.(cidaas_sdk.CidaasClient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	appConfig := preparePayload(d)
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

	error_string := ""
	error_code := int(appcreationresponse.Errors.Code)
	error_type := appcreationresponse.Errors.Type

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

	if !appcreationresponse.Success {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create app",
			Detail:   appcreationresponse.Errors.Error,
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

	d.SetId(appcreationresponse.Data.ClientId)
	custom_provider_name := d.Get("custom_provider_name").(string)

	if custom_provider_name != "" {
		var linkCpPayload cidaas_sdk.LinkCustomProviderStruct

		linkCpPayload.ClientId = appcreationresponse.Data.ClientId
		linkCpPayload.Test = bool(false)
		linkCpPayload.Type = "CUSTOM_OPENID_CONNECT"
		linkCpPayload.DisplayName = custom_provider_name

		response := cidaas_sdk.LinkCustomProvider(cidaas_client, linkCpPayload)

		if !response.Success {
			str := fmt.Sprintf("Unable to link custom provide to the client %+v", linkCpPayload.ClientId)
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  str,
				Detail:   appcreationresponse.Errors.Error,
			})
			return diags
		}
	}

	// resourceAppRead(ctx, d, m)

	return diags
}

func resourceAppRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	cidaas_client := m.(cidaas_sdk.CidaasClient)

	client_id := d.Id()
	appreadresponse := cidaas_sdk.GetApp(cidaas_client, client_id)

	if err := d.Set("client_id", appreadresponse.Data.ClientId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("client_name", appreadresponse.Data.ClientName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("client_type", appreadresponse.Data.ClientType); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("allow_login_with", appreadresponse.Data.AllowLoginWith); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("auto_login_after_register", appreadresponse.Data.AutoLoginAfterRegister); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("enable_passwordless_auth", appreadresponse.Data.EnablePasswordlessAuth); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("register_with_login_information", appreadresponse.Data.RegisterWithLoginInformation); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("hosted_page_group", appreadresponse.Data.HostedPageGroup); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("client_display_name", appreadresponse.Data.ClientDisplayName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("company_name", appreadresponse.Data.CompanyName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("company_address", appreadresponse.Data.CompanyAddress); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("company_website", appreadresponse.Data.CompanyWebsite); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("allowed_scopes", appreadresponse.Data.AllowedScopes); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("response_types", appreadresponse.Data.ResponseTypes); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("grant_types", appreadresponse.Data.GrantTypes); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("template_group_id", appreadresponse.Data.TemplateGroupId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("redirect_uris", appreadresponse.Data.RedirectURIS); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("allowed_logout_urls", appreadresponse.Data.AllowedLogoutUrls); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("fds_enabled", appreadresponse.Data.FdsEnabled); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("login_providers", appreadresponse.Data.LoginProviders); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceAppUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	cidaas_client := m.(cidaas_sdk.CidaasClient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	appConfig := preparePayload(d)

	response := cidaas_sdk.UpdateApp(cidaas_client, appConfig)

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "App Update Success",
		Detail:   strconv.FormatBool(response.Success),
	})

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "App Update Status",
		Detail:   strconv.Itoa(response.Status),
	})

	if !response.Success {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "App Update Failed",
			Detail:   "App Update Failed",
		})
	}

	if response.Success == false {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create app",
			Detail:   response.Errors.Error,
		})
		return diags
	}

	resourceAppRead(ctx, d, m)

	return diags
}

func resourceAppDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	cidaas_client := m.(cidaas_sdk.CidaasClient)

	client_id := d.Id()
	custom_provider_name := d.Get("custom_provider_name").(string)

	if custom_provider_name != "" {
		var linkCpPayload cidaas_sdk.LinkCustomProviderStruct

		t, _ := strconv.ParseBool("true")
		linkCpPayload.ClientId = client_id
		linkCpPayload.Test = t
		linkCpPayload.Type = "CUSTOM_OPENID_CONNECT"
		linkCpPayload.DisplayName = custom_provider_name

		response := cidaas_sdk.LinkCustomProvider(cidaas_client, linkCpPayload)

		if !response.Success {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to unlink custom provider to the client",
				Detail:   response.Error.Error,
			})
			return diags
		}
	}

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

	if !appdeleteresponse.Success {
		str := fmt.Sprintf("App deletion failed %+v", client_id)
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  str,
			Detail:   str,
		})
	}

	return diags
}

func preparePayload(d *schema.ResourceData) cidaas_sdk.AppConfig {
	var appConfig cidaas_sdk.AppConfig

	appConfig.ClientType = d.Get("client_type").(string)
	appConfig.AccentColor = d.Get("accent_color").(string)
	appConfig.PrimaryColor = d.Get("primary_color").(string)
	appConfig.MediaType = d.Get("media_type").(string)
	appConfig.ContentAlign = d.Get("content_align").(string)
	appConfig.AllowLoginWith = interfaceArray2StringArray(d.Get("allow_login_with").([]interface{}))
	appConfig.RedirectURIS = interfaceArray2StringArray(d.Get("redirect_uris").([]interface{}))
	appConfig.AllowedLogoutUrls = interfaceArray2StringArray(d.Get("allowed_logout_urls").([]interface{}))
	appConfig.EnableDeduplication = d.Get("enable_deduplication").(bool)
	appConfig.AutoLoginAfterRegister = d.Get("auto_login_after_register").(bool)
	appConfig.EnablePasswordlessAuth = d.Get("enable_passwordless_auth").(bool)
	appConfig.RegisterWithLoginInformation = d.Get("register_with_login_information").(bool)
	// appConfig.AllowDisposableEmail = d.Get("allow_disposable_email").(bool)
	// appConfig.ValidatePhoneNumber = d.Get("validate_phone_number").(bool)
	appConfig.FdsEnabled = d.Get("fds_enabled").(bool)
	appConfig.HostedPageGroup = d.Get("hosted_page_group").(string)
	appConfig.ClientName = d.Get("client_name").(string)
	appConfig.ClientDisplayName = d.Get("client_display_name").(string)
	appConfig.CompanyName = d.Get("company_name").(string)
	appConfig.CompanyAddress = d.Get("company_address").(string)
	appConfig.CompanyWebsite = d.Get("company_website").(string)
	appConfig.AllowedScopes = interfaceArray2StringArray(d.Get("allowed_scopes").([]interface{}))
	appConfig.ResponseTypes = interfaceArray2StringArray(d.Get("response_types").([]interface{}))
	appConfig.GrantTypes = interfaceArray2StringArray(d.Get("grant_types").([]interface{}))
	appConfig.LoginProviders = interfaceArray2StringArray(d.Get("login_providers").([]interface{}))
	// appConfig.AdditionalAccessTokenPayload = d.Get("additional_access_token_payload").(string)
	// appConfig.RequiredFields = interfaceArray2StringArray(d.Get("required_fields").([]interface{}))
	// appConfig.ApplicationMetaData = interfaceArray2StringArray(d.Get("application_meta_data").([]interface{}))
	// appConfig.IsHybridApp = d.Get("is_hybrid_app").(bool)
	// appConfig.AllowedWebOrigins = interfaceArray2StringArray(d.Get("allowed_web_origins").([]interface{}))
	// appConfig.AllowedOrigins = interfaceArray2StringArray(d.Get("allowed_origins").([]interface{}))
	// appConfig.MobileSettings = d.Get("mobile_settings").(string)
	// appConfig.RefreshTokenLifetimeInSeconds = d.Get("refresh_token_lifetime_in_seconds").(int)
	appConfig.TemplateGroupId = d.Get("template_group_id").(string)

	return appConfig
}
