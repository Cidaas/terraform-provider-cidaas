package cidaas

import (
	"context"
	"encoding/json"
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

		Schema: map[string]*schema.Schema{
			"client_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"accent_color": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"primary_color": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"media_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"client_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"custom_provider_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"content_align": {
				Type:     schema.TypeString,
				Computed: true,
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
			"application_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"required_fields": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"application_meta_data": {
				Type:     schema.TypeMap,
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
				Type:     schema.TypeSet,
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

			"logo_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"policy_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tos_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"default_max_age": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"template_group_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"token_lifetime_in_seconds": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"id_token_lifetime_in_seconds": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"initiate_login_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"contacts": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"client_secret_expires_at": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"client_id_issued_at": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"registration_client_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"registration_access_token": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"jwks_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"jwks": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sector_identifier_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subject_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id_token_signed_response_alg": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id_token_encrypted_response_alg": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id_token_encrypted_response_enc": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"userinfo_signed_response_alg": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"userinfo_encrypted_response_alg": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"userinfo_encrypted_response_enc": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"request_object_signing_alg": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"request_object_encryption_alg": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"request_object_encryption_enc": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"request_uris": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"token_endpoint_auth_method": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"token_endpoint_auth_signing_alg": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"require_auth_time": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"default_acr_values": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"default_scopes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"pending_scopes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"app_owner": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"jwe_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"user_consent": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"deleted": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"allowed_fields": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"consent_page_group": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"always_ask_mfa": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"allowed_mfa": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"password_policy_ref": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"captcha_ref": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"captcha_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"consent_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"email_verification_required": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"mobile_number_verification_required": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"accept_roles_in_the_registration": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"allowed_roles": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"default_roles": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"enable_classical_provider": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"sub": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"role": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"is_remember_me_selected": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"mfa_configuration": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"suggest_mfa": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"custom_providers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"provider_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"social_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"social_providers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"provider_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"logo_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"saml_providers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"provider_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"logo_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"ad_providers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"provider_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"logo_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"allowed_groups": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"roles": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"default_roles": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"operations_allowed_groups": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"roles": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"default_roles": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},

			"app_key": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"push_config": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
	// appConfig.EnableDeduplication = d.Get("enable_deduplication").(bool)
	// appConfig.AllowDisposableEmail = d.Get("allow_disposable_email").(bool)
	// appConfig.ValidatePhoneNumber = d.Get("validate_phone_number").(bool)
	// appConfig.FdsEnabled = d.Get("fds_enabled").(bool)
	// appConfig.LoginProviders = interfaceArray2StringArray(d.Get("login_providers").([]interface{}))
	// appConfig.AdditionalAccessTokenPayload = d.Get("additional_access_token_payload").(string)
	// appConfig.RequiredFields = interfaceArray2StringArray(d.Get("required_fields").([]interface{}))
	// appConfig.ApplicationMetaData = d.Get("template_group_id").(string)
	// appConfig.IsHybridApp = d.Get("is_hybrid_app").(bool)
	// appConfig.AllowedWebOrigins = interfaceArray2StringArray(d.Get("allowed_web_origins").([]interface{}))
	// appConfig.AllowedOrigins = interfaceArray2StringArray(d.Get("allowed_origins").([]interface{}))
	// appConfig.MobileSettings = d.Get("mobile_settings").(string)

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

	var linkCpPayload cidaas_sdk.LinkCustomProviderStruct

	linkCpPayload.ClientId = appcreationresponse.Data.ClientId
	linkCpPayload.Test = bool(false)
	linkCpPayload.Type = "CUSTOM_OPENID_CONNECT"
	linkCpPayload.DisplayName = d.Get("custom_provider_name").(string)

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

	if !appupdateresponse.Success {
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

	var linkCpPayload cidaas_sdk.LinkCustomProviderStruct

	t, _ := strconv.ParseBool("true")
	linkCpPayload.ClientId = client_id
	linkCpPayload.Test = t
	linkCpPayload.Type = "CUSTOM_OPENID_CONNECT"
	linkCpPayload.DisplayName = d.Get("custom_provider_name").(string)

	json_payload, _ := json.Marshal(linkCpPayload)

	payload_string := string(json_payload)

	response := cidaas_sdk.LinkCustomProvider(cidaas_client, linkCpPayload)

	if !response.Success {
		str := fmt.Sprintf("Unable to unlink custom provider to the client %+v", payload_string)
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  str,
			Detail:   response.Error.Error,
		})
		return diags
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
