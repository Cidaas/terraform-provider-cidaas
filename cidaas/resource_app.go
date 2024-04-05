package cidaas

import (
	"context"
	"fmt"

	"terraform-provider-cidaas/helper/cidaas"
	"terraform-provider-cidaas/helper/util"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"VIDEO", "IMAGE"}, false),
			},
			"content_align": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"CENTER", "LEFT", "RIGHT"}, false),
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
			"is_hybrid_app": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
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
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"team_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"bundle_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"package_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"key_hash": {
							Type:     schema.TypeString,
							Optional: true,
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
				Optional: true,
			},
			"client_secret": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"policy_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tos_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"imprint_uri": {
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
			"token_endpoint_auth_method": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"token_endpoint_auth_signing_alg": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"default_acr_values": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"editable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"web_message_uris": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"social_providers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"provider_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"social_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"custom_providers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"logo_url": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"provider_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"saml_providers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"logo_url": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"provider_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"ad_providers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"logo_url": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"provider_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"app_owner": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"jwe_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"user_consent": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"allowed_groups": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"group_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"roles": {
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
					},
				},
			},
			"operations_allowed_groups": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"group_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"roles": {
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
					},
				},
			},
			"deleted": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
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
			"app_key": {
				Type:      schema.TypeList,
				Sensitive: true,
				Computed:  true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"key_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"public_key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"public_key_jwk": {
							Type:     schema.TypeMap,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"created_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"always_ask_mfa": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"smart_mfa": {
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
			"communication_medium_verification": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"email_verification_required": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"mobile_number_verification_required": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
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
			"is_remember_me_selected": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"bot_provider": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"allow_guest_login_groups": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"group_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"roles": {
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
					},
				},
			},
			"is_login_success_page_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_register_success_page_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"group_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"admin_client": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_group_login_selection_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"group_selection": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"always_show_group_selection": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"selectable_groups": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"selectable_group_types": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"group_types": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"backchannel_logout_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"post_logout_redirect_uris": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"logo_align": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mfa": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"setting": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"OFF", "ALWAYS", "SMART", "TIME_BASED", "SMART_PLUS_TIME_BASED"}, true),
						},
						"time_interval_in_seconds": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"allowed_methods": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"push_config": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tenant_key": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"vendor": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"key": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"secret": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"owner": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"webfinger": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"application_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 50),
			},
			"logo_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"initiate_login_uri": {
				Type:     schema.TypeString,
				Optional: true,
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
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 150),
			},
			"id_token_signed_response_alg": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 150),
			},
			"id_token_encrypted_response_alg": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 150),
			},
			"id_token_encrypted_response_enc": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 150),
			},
			"userinfo_signed_response_alg": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 150),
			},
			"userinfo_encrypted_response_alg": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 150),
			},
			"userinfo_encrypted_response_enc": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 150),
			},
			"request_object_signing_alg": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 150),
			},
			"request_object_encryption_alg": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 150),
			},
			"request_object_encryption_enc": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 150),
			},
			"request_uris": {
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
			"consent_page_group": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"password_policy_ref": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"blocking_mechanism_ref": {
				Type:     schema.TypeString,
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
			"login_spi": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oauth_client_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"spi_url": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"background_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"video_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"bot_captcha_ref": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"updated_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"application_meta_data": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
	if err := d.Set("client_secret", response.Data.ClientSecret); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("error while setting client_secret to cidaas_app %+v", appConfig.ClientName),
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId(response.Data.ClientId)
	resourceAppRead(ctx, d, m)
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

	if err := d.Set("client_type", response.Data.ClientType); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("accent_color", response.Data.AccentColor); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("primary_color", response.Data.PrimaryColor); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("media_type", response.Data.MediaType); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("content_align", response.Data.ContentAlign); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("allow_login_with", response.Data.AllowLoginWith); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("redirect_uris", response.Data.RedirectURIS); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("allowed_logout_urls", response.Data.AllowedLogoutUrls); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("enable_deduplication", response.Data.EnableDeduplication); err != nil {
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
	if err := d.Set("allow_disposable_email", response.Data.AllowDisposableEmail); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("validate_phone_number", response.Data.ValidatePhoneNumber); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("fds_enabled", response.Data.FdsEnabled); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("hosted_page_group", response.Data.HostedPageGroup); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("client_name", response.Data.ClientName); err != nil {
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
	if err := d.Set("login_providers", response.Data.LoginProviders); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("additional_access_token_payload", response.Data.AdditionalAccessTokenPayload); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("required_fields", response.Data.RequiredFields); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("is_hybrid_app", response.Data.IsHybridApp); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("allowed_web_origins", response.Data.AllowedWebOrigins); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("allowed_origins", response.Data.AllowedOrigins); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("mobile_settings", cidaas.FlattenMobileSettings(response.Data.MobileSettings)); err != nil {
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
	if err := d.Set("template_group_id", response.Data.TemplateGroupId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("client_id", response.Data.ClientId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("client_secret", response.Data.ClientSecret); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("policy_uri", response.Data.PolicyUri); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("tos_uri", response.Data.TosUri); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("imprint_uri", response.Data.ImprintUri); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("contacts", response.Data.Contacts); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("token_endpoint_auth_method", response.Data.TokenEndpointAuthMethod); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("token_endpoint_auth_signing_alg", response.Data.TokenEndpointAuthSigningAlg); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("default_acr_values", response.Data.DefaultAcrValues); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("editable", response.Data.Editable); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("web_message_uris", response.Data.WebMessageUris); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("social_providers", cidaas.FlattenSocialProvider(&response.Data.SocialProviders)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("custom_providers", cidaas.FlattenProviders(&response.Data.CustomProviders)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("saml_providers", cidaas.FlattenProviders(&response.Data.SamlProviders)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("ad_providers", cidaas.FlattenProviders(&response.Data.AdProviders)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("app_owner", response.Data.AppOwner); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("jwe_enabled", response.Data.JweEnabled); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("user_consent", response.Data.UserConsent); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("allowed_groups", cidaas.FlattenAllowedGroups(&response.Data.AllowedGroups)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("operations_allowed_groups", cidaas.FlattenAllowedGroups(&response.Data.OperationsAllowedGroups)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("deleted", response.Data.Deleted); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("enabled", response.Data.Enabled); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("allowed_fields", response.Data.AllowedFields); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("app_key", cidaas.SerializeAppKey(response.Data.AppKey)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("always_ask_mfa", response.Data.AlwaysAskMfa); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("smart_mfa", response.Data.SmartMfa); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("allowed_mfa", response.Data.AllowedMfa); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("captcha_ref", response.Data.CaptchaRef); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("captcha_refs", response.Data.CaptchaRefs); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("consent_refs", response.Data.ConsentRefs); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("communication_medium_verification", response.Data.CommunicationMediumVerification); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("email_verification_required", response.Data.EmailVerificationRequired); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("mobile_number_verification_required", response.Data.MobileNumberVerificationRequired); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("allowed_roles", response.Data.AllowedRoles); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("default_roles", response.Data.DefaultRoles); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("enable_classical_provider", response.Data.EnableClassicalProvider); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("is_remember_me_selected", response.Data.IsRememberMeSelected); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("bot_provider", response.Data.BotProvider); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("allow_guest_login_groups", cidaas.FlattenAllowedGroups(&response.Data.AllowGuestLoginGroups)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("is_login_success_page_enabled", response.Data.IsLoginSuccessPageEnabled); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("is_register_success_page_enabled", response.Data.IsRegisterSuccessPageEnabled); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("group_ids", response.Data.GroupIds); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("admin_client", response.Data.AdminClient); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("is_group_login_selection_enabled", response.Data.IsGroupLoginSelectionEnabled); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("group_selection", cidaas.FlattenGroupSelection(response.Data.GroupSelection)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("group_types", response.Data.GroupTypes); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("backchannel_logout_uri", response.Data.BackchannelLogoutUri); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("post_logout_redirect_uris", response.Data.PostLogoutRedirectUris); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("logo_align", response.Data.LogoAlign); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("mfa", cidaas.FlattenMfa(response.Data.Mfa)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("push_config", cidaas.FlattenPushConfig(response.Data.PushConfig)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("webfinger", response.Data.Webfinger); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("application_type", response.Data.ApplicationType); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("logo_uri", response.Data.LogoUri); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("initiate_login_uri", response.Data.InitiateLoginUri); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("client_secret_expires_at", response.Data.ClientSecretExpiresAt); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("client_id_issued_at", response.Data.ClientIdIssuedAt); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("registration_client_uri", response.Data.RegistrationClientUri); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("registration_access_token", response.Data.RegistrationAccessToken); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("client_uri", response.Data.ClientUri); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("jwks_uri", response.Data.JwksUri); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("jwks", response.Data.Jwks); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("sector_identifier_uri", response.Data.SectorIdentifierUri); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("subject_type", response.Data.SubjectType); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("id_token_signed_response_alg", response.Data.IdTokenEncryptedResponseAlg); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("id_token_encrypted_response_alg", response.Data.IdTokenEncryptedResponseAlg); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("id_token_encrypted_response_enc", response.Data.IdTokenEncryptedResponseEnc); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("userinfo_signed_response_alg", response.Data.UserinfoSignedResponseAlg); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("userinfo_encrypted_response_alg", response.Data.UserinfoEncryptedResponseAlg); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("userinfo_encrypted_response_enc", response.Data.UserinfoEncryptedResponseEnc); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("request_object_signing_alg", response.Data.RequestObjectSigningAlg); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("request_object_encryption_alg", response.Data.RequestObjectEncryptionAlg); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("request_object_encryption_enc", response.Data.RequestObjectEncryptionEnc); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("request_uris", response.Data.RequestUris); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", response.Data.Description); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("default_scopes", response.Data.DefaultScopes); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("pending_scopes", response.Data.PendingScopes); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("consent_page_group", response.Data.ConsentPageGroup); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("password_policy_ref", response.Data.PasswordPolicyRef); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("blocking_mechanism_ref", response.Data.BlockingMechanismRef); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("sub", response.Data.Sub); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("role", response.Data.Role); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("mfa_configuration", response.Data.MfaConfiguration); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("suggest_mfa", response.Data.SuggestMfa); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("login_spi", cidaas.FlattenLoginSpi(response.Data.LoginSpi)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("background_uri", response.Data.BackgroundUri); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("video_url", response.Data.VideoUrl); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("bot_captcha_ref", response.Data.BotCaptchaRef); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("created_at", response.Data.CreatedTime); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("updated_at", response.Data.UpdatedTime); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("application_meta_data", response.Data.ApplicationMetaData); err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceAppUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	d.Set("client_id", d.Id())
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
	var appConfig cidaas.AppConfig
	appConfig.ClientId = client_id
	_, err := cidaas_client.DeleteApp(appConfig)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("failed to delete client %+v", client_id),
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
	appConfig.AllowDisposableEmail = d.Get("allow_disposable_email").(bool)
	appConfig.ValidatePhoneNumber = d.Get("validate_phone_number").(bool)
	appConfig.FdsEnabled = d.Get("fds_enabled").(bool)
	appConfig.HostedPageGroup = d.Get("hosted_page_group").(string)
	appConfig.ClientName = d.Get("client_name").(string)
	appConfig.ClientDisplayName = d.Get("client_display_name").(string)
	appConfig.CompanyName = d.Get("company_name").(string)
	appConfig.CompanyAddress = d.Get("company_address").(string)
	appConfig.CompanyWebsite = d.Get("company_website").(string)
	appConfig.AllowedScopes = util.InterfaceArray2StringArray(d.Get("allowed_scopes").([]interface{}))
	appConfig.ResponseTypes = util.InterfaceArray2StringArray(d.Get("response_types").([]interface{}))
	appConfig.GrantTypes = util.InterfaceArray2StringArray(d.Get("grant_types").([]interface{}))
	appConfig.LoginProviders = util.InterfaceArray2StringArray(d.Get("login_providers").([]interface{}))
	appConfig.AdditionalAccessTokenPayload = util.InterfaceArray2StringArray(d.Get("additional_access_token_payload").([]interface{}))
	appConfig.RequiredFields = util.InterfaceArray2StringArray(d.Get("required_fields").([]interface{}))
	appConfig.IsHybridApp = d.Get("is_hybrid_app").(bool)
	appConfig.AllowedWebOrigins = util.InterfaceArray2StringArray(d.Get("allowed_web_origins").([]interface{}))
	appConfig.AllowedOrigins = util.InterfaceArray2StringArray(d.Get("allowed_origins").([]interface{}))
	appConfig.MobileSettings = cidaas.SerializeMobileSettings(d.Get("mobile_settings").([]interface{}))
	appConfig.DefaultMaxAge = d.Get("default_max_age").(int)
	appConfig.TokenLifetimeInSeconds = d.Get("token_lifetime_in_seconds").(int)
	appConfig.IdTokenLifetimeInSeconds = d.Get("id_token_lifetime_in_seconds").(int)
	appConfig.RefreshTokenLifetimeInSeconds = d.Get("refresh_token_lifetime_in_seconds").(int)
	appConfig.TemplateGroupId = d.Get("template_group_id").(string)
	appConfig.ClientId = d.Get("client_id").(string)
	appConfig.ClientSecret = d.Get("client_secret").(string)
	appConfig.PolicyUri = d.Get("policy_uri").(string)
	appConfig.TosUri = d.Get("tos_uri").(string)
	appConfig.ImprintUri = d.Get("tos_uri").(string)
	appConfig.Contacts = util.InterfaceArray2StringArray(d.Get("contacts").([]interface{}))
	appConfig.TokenEndpointAuthMethod = d.Get("token_endpoint_auth_method").(string)
	appConfig.TokenEndpointAuthSigningAlg = d.Get("token_endpoint_auth_signing_alg").(string)
	appConfig.DefaultAcrValues = util.InterfaceArray2StringArray(d.Get("default_acr_values").([]interface{}))
	appConfig.Editable = d.Get("editable").(bool)
	appConfig.WebMessageUris = util.InterfaceArray2StringArray(d.Get("web_message_uris").([]interface{}))
	appConfig.SocialProviders = cidaas.SerializeSocialProviders(d.Get("social_providers").([]interface{}))
	appConfig.CustomProviders = cidaas.SerializeProviders(d.Get("custom_providers").([]interface{}))
	appConfig.SamlProviders = cidaas.SerializeProviders(d.Get("saml_providers").([]interface{}))
	appConfig.AdProviders = cidaas.SerializeProviders(d.Get("ad_providers").([]interface{}))
	appConfig.AppOwner = d.Get("app_owner").(string)
	appConfig.JweEnabled = d.Get("jwe_enabled").(bool)
	appConfig.UserConsent = d.Get("user_consent").(bool)
	appConfig.AllowedGroups = cidaas.SerializeAllowedGroups(d.Get("allowed_groups").([]interface{}))
	appConfig.OperationsAllowedGroups = cidaas.SerializeAllowedGroups(d.Get("operations_allowed_groups").([]interface{}))
	appConfig.Deleted = d.Get("deleted").(bool)
	appConfig.Enabled = d.Get("enabled").(bool)
	appConfig.AllowedFields = util.InterfaceArray2StringArray(d.Get("allowed_fields").([]interface{}))
	appConfig.AlwaysAskMfa = d.Get("always_ask_mfa").(bool)
	appConfig.SmartMfa = d.Get("smart_mfa").(bool)
	appConfig.AllowedMfa = util.InterfaceArray2StringArray(d.Get("allowed_mfa").([]interface{}))
	appConfig.CaptchaRef = d.Get("captcha_ref").(string)
	appConfig.CaptchaRefs = util.InterfaceArray2StringArray(d.Get("captcha_refs").([]interface{}))
	appConfig.ConsentRefs = util.InterfaceArray2StringArray(d.Get("consent_refs").([]interface{}))
	appConfig.CommunicationMediumVerification = d.Get("communication_medium_verification").(string)
	appConfig.EmailVerificationRequired = d.Get("email_verification_required").(bool)
	appConfig.MobileNumberVerificationRequired = d.Get("mobile_number_verification_required").(bool)
	appConfig.AllowedRoles = util.InterfaceArray2StringArray(d.Get("allowed_roles").([]interface{}))
	appConfig.DefaultRoles = util.InterfaceArray2StringArray(d.Get("default_roles").([]interface{}))
	appConfig.EnableClassicalProvider = d.Get("enable_classical_provider").(bool)
	appConfig.IsRememberMeSelected = d.Get("is_remember_me_selected").(bool)
	appConfig.BotProvider = d.Get("bot_provider").(string)
	appConfig.AllowGuestLoginGroups = cidaas.SerializeAllowedGroups(d.Get("allow_guest_login_groups").([]interface{}))
	appConfig.IsLoginSuccessPageEnabled = d.Get("is_login_success_page_enabled").(bool)
	appConfig.IsRegisterSuccessPageEnabled = d.Get("is_register_success_page_enabled").(bool)
	appConfig.GroupIds = util.InterfaceArray2StringArray(d.Get("group_ids").([]interface{}))
	appConfig.AdminClient = d.Get("admin_client").(bool)
	appConfig.IsGroupLoginSelectionEnabled = d.Get("is_group_login_selection_enabled").(bool)
	appConfig.GroupSelection = cidaas.SerializeGroupSelection(d.Get("group_selection").([]interface{}))
	appConfig.GroupTypes = util.InterfaceArray2StringArray(d.Get("group_types").([]interface{}))
	appConfig.BackchannelLogoutUri = d.Get("backchannel_logout_uri").(string)
	appConfig.PostLogoutRedirectUris = util.InterfaceArray2StringArray(d.Get("post_logout_redirect_uris").([]interface{}))
	appConfig.LogoAlign = d.Get("logo_align").(string)
	appConfig.Mfa = cidaas.SerializeMfaOption(d.Get("mfa").([]interface{}))
	appConfig.PushConfig = cidaas.SerializePushConfig(d.Get("push_config").([]interface{}))
	appConfig.Webfinger = d.Get("webfinger").(string)
	appConfig.ApplicationType = d.Get("application_type").(string)
	appConfig.LogoUri = d.Get("logo_uri").(string)
	appConfig.InitiateLoginUri = d.Get("initiate_login_uri").(string)
	appConfig.ClientSecretExpiresAt = d.Get("client_secret_expires_at").(int)
	appConfig.ClientIdIssuedAt = d.Get("client_id_issued_at").(int)
	appConfig.RegistrationClientUri = d.Get("registration_client_uri").(string)
	appConfig.RegistrationAccessToken = d.Get("registration_access_token").(string)
	appConfig.ClientUri = d.Get("client_uri").(string)
	appConfig.JwksUri = d.Get("jwks_uri").(string)
	appConfig.Jwks = d.Get("jwks").(string)
	appConfig.SectorIdentifierUri = d.Get("sector_identifier_uri").(string)
	appConfig.SubjectType = d.Get("subject_type").(string)
	appConfig.IdTokenSignedResponseAlg = d.Get("id_token_signed_response_alg").(string)
	appConfig.IdTokenEncryptedResponseAlg = d.Get("id_token_encrypted_response_alg").(string)
	appConfig.IdTokenEncryptedResponseEnc = d.Get("id_token_encrypted_response_enc").(string)
	appConfig.UserinfoSignedResponseAlg = d.Get("userinfo_signed_response_alg").(string)
	appConfig.UserinfoEncryptedResponseAlg = d.Get("userinfo_encrypted_response_alg").(string)
	appConfig.UserinfoEncryptedResponseEnc = d.Get("userinfo_encrypted_response_enc").(string)
	appConfig.RequestObjectSigningAlg = d.Get("request_object_signing_alg").(string)
	appConfig.RequestObjectEncryptionAlg = d.Get("request_object_encryption_alg").(string)
	appConfig.RequestObjectEncryptionEnc = d.Get("request_object_encryption_enc").(string)
	appConfig.RequestUris = util.InterfaceArray2StringArray(d.Get("request_uris").([]interface{}))
	appConfig.Description = d.Get("description").(string)
	appConfig.DefaultScopes = util.InterfaceArray2StringArray(d.Get("default_scopes").([]interface{}))
	appConfig.PendingScopes = util.InterfaceArray2StringArray(d.Get("pending_scopes").([]interface{}))
	appConfig.ConsentPageGroup = d.Get("consent_page_group").(string)
	appConfig.PasswordPolicyRef = d.Get("password_policy_ref").(string)
	appConfig.BlockingMechanismRef = d.Get("blocking_mechanism_ref").(string)
	appConfig.Sub = d.Get("sub").(string)
	appConfig.Role = d.Get("role").(string)
	appConfig.MfaConfiguration = d.Get("mfa_configuration").(string)
	appConfig.SuggestMfa = util.InterfaceArray2StringArray(d.Get("suggest_mfa").([]interface{}))
	appConfig.LoginSpi = cidaas.SerializeLoginSpi(d.Get("login_spi").([]interface{}))
	appConfig.BackgroundUri = d.Get("background_uri").(string)
	appConfig.VideoUrl = d.Get("video_url").(string)
	appConfig.BotCaptchaRef = d.Get("bot_captcha_ref").(string)
	appConfig.ApplicationMetaData = d.Get("application_meta_data").(map[string]interface{})

	return appConfig
}
