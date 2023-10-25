package cidaas

import (
	"context"
	"fmt"
	"strconv"

	"terraform-provider-cidaas/helper/cidaas"
	"terraform-provider-cidaas/helper/util"

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
			"default_max_age": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"token_lifetime_in_seconds": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"id_token_lifetime_in_seconds": {
				Type:     schema.TypeInt,
				Optional: true,
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
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"custom_provider_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceAppCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	cidaas_client := m.(cidaas.CidaasClient)
	appConfig := preparePayload(d)
	response, err := cidaas_client.CreateApp(appConfig)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("failed to create client %+v", appConfig.ClientName),
			Detail:   err.Error(),
		})
		return diags
	}
	if err := d.Set("client_id", response.Data.ClientId); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("error while setting client_id to cidaas_app %+v", appConfig.ClientName),
			Detail:   err.Error(),
		})
		return diags
	}
	custom_provider_name := d.Get("custom_provider_name").(string)
	if custom_provider_name != "" {
		var enableCpPayload cidaas.CustomProviderConfigPayload

		enableCpPayload.ClientId = response.Data.ClientId
		enableCpPayload.Test = bool(false)
		enableCpPayload.Type = "CUSTOM_OPENID_CONNECT"
		enableCpPayload.DisplayName = custom_provider_name

		_, err := cidaas_client.ConfigureCustomProvider(enableCpPayload)
		if err != nil {
			errStr := fmt.Sprintf("custom provider configuration failed for the client %+v. app created successfully", enableCpPayload.ClientId)
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  errStr,
				Detail:   err.Error(),
			})
			return diags
		}
	}
	d.SetId(response.Data.ClientId)
	return diags
}

func resourceAppRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client_id := d.Id()
	cidaas_client := m.(cidaas.CidaasClient)

	var appConfig cidaas.AppConfig
	appConfig.ClientId = client_id
	response, err := cidaas_client.GetApp(appConfig)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("failed to get app for the client_id %+v", client_id),
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("client_id", response.Data.ClientId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("client_name", response.Data.ClientName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("client_type", response.Data.ClientType); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("allow_login_with", response.Data.AllowLoginWith); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("auto_login_after_register", response.Data.AutoLoginAfterRegister); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("enable_passwordless_auth", response.Data.EnablePasswordlessAuth); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("register_with_login_information", response.Data.RegisterWithLoginInformation); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("hosted_page_group", response.Data.HostedPageGroup); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("client_display_name", response.Data.ClientDisplayName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("company_name", response.Data.CompanyName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("company_address", response.Data.CompanyAddress); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("company_website", response.Data.CompanyWebsite); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("allowed_scopes", response.Data.AllowedScopes); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("response_types", response.Data.ResponseTypes); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("grant_types", response.Data.GrantTypes); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("template_group_id", response.Data.TemplateGroupId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("redirect_uris", response.Data.RedirectURIS); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("allowed_logout_urls", response.Data.AllowedLogoutUrls); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("fds_enabled", response.Data.FdsEnabled); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("login_providers", response.Data.LoginProviders); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("default_max_age", response.Data.DefaultMaxAge); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("token_lifetime_in_seconds", response.Data.TokenLifetimeInSeconds); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("id_token_lifetime_in_seconds", response.Data.IdTokenLifetimeInSeconds); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("refresh_token_lifetime_in_seconds", response.Data.RefreshTokenLifetimeInSeconds); err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceAppUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	cidaas_client := m.(cidaas.CidaasClient)
	appConfig := preparePayload(d)
	_, err := cidaas_client.UpdateApp(appConfig)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("failed to update client %+v", appConfig.ClientId),
			Detail:   err.Error(),
		})
		return diags
	}
	resourceAppRead(ctx, d, m)
	return diags
}

func resourceAppDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	cidaas_client := m.(cidaas.CidaasClient)
	client_id := d.Id()
	custom_provider_name := d.Get("custom_provider_name").(string)

	if custom_provider_name != "" {
		var disableCpPayload cidaas.CustomProviderConfigPayload
		t, _ := strconv.ParseBool("true")
		disableCpPayload.ClientId = client_id
		disableCpPayload.Test = t
		disableCpPayload.Type = "CUSTOM_OPENID_CONNECT"
		disableCpPayload.DisplayName = custom_provider_name

		_, err := cidaas_client.ConfigureCustomProvider(disableCpPayload)
		if err != nil {
			errStr := fmt.Sprintf("custom provider configuration failed for the client %+v", disableCpPayload.ClientId)
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  errStr,
				Detail:   err.Error(),
			})
			return diags
		}
	}
	var appConfig cidaas.AppConfig
	appConfig.ClientId = client_id
	resp, err := cidaas_client.DeleteApp(appConfig)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("failed to delete client %+v with status %+v", client_id, resp.Status),
			Detail:   err.Error(),
		})
		return diags
	}
	return diags
}

func preparePayload(d *schema.ResourceData) cidaas.AppConfig {
	var appConfig cidaas.AppConfig

	appConfig.ClientType = d.Get("client_type").(string)
	appConfig.AccentColor = d.Get("accent_color").(string)
	appConfig.PrimaryColor = d.Get("primary_color").(string)
	appConfig.MediaType = d.Get("media_type").(string)
	appConfig.ContentAlign = d.Get("content_align").(string)
	appConfig.AllowLoginWith = util.InterfaceArray2StringArray(d.Get("allow_login_with").([]interface{}))
	appConfig.RedirectURIS = util.InterfaceArray2StringArray(d.Get("redirect_uris").([]interface{}))
	appConfig.AllowedLogoutUrls = util.InterfaceArray2StringArray(d.Get("allowed_logout_urls").([]interface{}))
	appConfig.EnableDeduplication = d.Get("enable_deduplication").(bool)
	appConfig.AutoLoginAfterRegister = d.Get("auto_login_after_register").(bool)
	appConfig.EnablePasswordlessAuth = d.Get("enable_passwordless_auth").(bool)
	appConfig.RegisterWithLoginInformation = d.Get("register_with_login_information").(bool)
	appConfig.FdsEnabled = d.Get("fds_enabled").(bool)
	appConfig.HostedPageGroup = d.Get("hosted_page_group").(string)
	appConfig.ClientName = d.Get("client_name").(string)
	appConfig.ClientDisplayName = d.Get("client_display_name").(string)
	appConfig.CompanyName = d.Get("company_name").(string)
	appConfig.ClientId = d.Get("client_id").(string)
	appConfig.CompanyAddress = d.Get("company_address").(string)
	appConfig.CompanyWebsite = d.Get("company_website").(string)
	appConfig.AllowedScopes = util.InterfaceArray2StringArray(d.Get("allowed_scopes").([]interface{}))
	appConfig.ResponseTypes = util.InterfaceArray2StringArray(d.Get("response_types").([]interface{}))
	appConfig.GrantTypes = util.InterfaceArray2StringArray(d.Get("grant_types").([]interface{}))
	appConfig.LoginProviders = util.InterfaceArray2StringArray(d.Get("login_providers").([]interface{}))
	appConfig.TemplateGroupId = d.Get("template_group_id").(string)
	appConfig.DefaultMaxAge = d.Get("default_max_age").(int)
	appConfig.TokenLifetimeInSeconds = d.Get("token_lifetime_in_seconds").(int)
	appConfig.IdTokenLifetimeInSeconds = d.Get("id_token_lifetime_in_seconds").(int)
	appConfig.RefreshTokenLifetimeInSeconds = d.Get("refresh_token_lifetime_in_seconds").(int)

	return appConfig
}
