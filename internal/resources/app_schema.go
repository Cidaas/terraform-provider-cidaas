package resources

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func (r *AppResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The App resource allows creation and management of clients in Cidaas system." +
			" When creating a client with a custom `client_id` and `client_secret` you can include the configuration in the resource." +
			" If not provided, Cidaas will generate a set for you. `client_secret` is sensitive data." +
			" Refer to the article [Terraform Sensitive Variables](https://developer.hashicorp.com/terraform/tutorials/configuration-language/sensitive-variables) to properly handle sensitive information." +
			"\n\n Ensure that the below scopes are assigned to the client with the specified `client_id`:" +
			"\n- cidaas:apps_read" +
			"\n- cidaas:apps_write" +
			"\n- cidaas:apps_delete",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The ID of the resource.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"client_type": schema.StringAttribute{
				Optional: true,
				MarkdownDescription: "The type of the client. The allowed values are " +
					"SINGLE_PAGE, REGULAR_WEB, NON_INTERACTIVE" +
					"IOS, ANDROID, WINDOWS_MOBILE, DESKTOP, MOBILE, DEVICE and THIRD_PARTY",
				PlanModifiers: []planmodifier.String{
					&stringCustomRequired{},
				},
				Validators: []validator.String{
					stringvalidator.OneOf([]string{
						"SINGLE_PAGE", "REGULAR_WEB", "NON_INTERACTIVE",
						"IOS", "ANDROID", "WINDOWS_MOBILE", "DESKTOP", "MOBILE", "DEVICE", "THIRD_PARTY",
					}...),
				},
			},
			"accent_color": schema.StringAttribute{
				Optional: true,
				Computed: true,
				MarkdownDescription: "The accent color of the client. e.g., `#f7941d`. The value must be a valid hex color" +
					"The default is set to `#ef4923`.",
				Default: stringdefault.StaticString("#ef4923"),
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^#([a-fA-F0-9]{6}|[a-fA-F0-9]{3})$`),
						"accent_color must be a valid hex color",
					),
				},
			},
			"primary_color": schema.StringAttribute{
				Optional: true,
				Computed: true,
				MarkdownDescription: "The primary color of the client. e.g., `#ef4923`. The value must be a valid hex color" +
					"The default is set to `#f7941d`.",
				Default: stringdefault.StaticString("#f7941d"),
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^#([a-fA-F0-9]{6}|[a-fA-F0-9]{3})$`),
						"must be a valid hex color",
					),
				},
			},
			"media_type": schema.StringAttribute{
				Optional: true,
				Computed: true,
				MarkdownDescription: "The media type of the client. e.g., `IMAGE`. Allowed values are VIDEO and IMAGE" +
					"The default is set to `IMAGE`.",
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"VIDEO", "IMAGE"}...),
				},
				Default: stringdefault.StaticString("IMAGE"),
			},
			"content_align": schema.StringAttribute{
				Optional: true,
				Computed: true,
				MarkdownDescription: "The alignment of the content of the client. e.g., `CENTER`. Allowed values are CENTER, LEFT and RIGHT" +
					"The default is set to `CENTER`.",
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"CENTER", "LEFT", "RIGHT"}...),
				},
				Default: stringdefault.StaticString("CENTER"),
			},
			"allow_login_with": schema.SetAttribute{
				ElementType: types.StringType,
				Computed:    true,
				Optional:    true,
				MarkdownDescription: "allow_login_with is used to specify the preferred methods of login allowed for a client. Allowed values are EMAIL, MOBILE and USER_NAME" +
					"The default is set to `['EMAIL', 'MOBILE', 'USER_NAME']`.",
				Validators: []validator.Set{
					setvalidator.ValueStringsAre(
						stringvalidator.OneOf([]string{"EMAIL", "MOBILE", "USER_NAME"}...),
					),
				},
				Default: setdefault.StaticValue(basetypes.NewSetValueMust(types.StringType, []attr.Value{
					types.StringValue("EMAIL"), types.StringValue("MOBILE"), types.StringValue("USER_NAME"),
				})),
			},
			"redirect_uris": schema.SetAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				MarkdownDescription: "Redirect URIs for OAuth2 client.",
				Validators: []validator.Set{
					setvalidator.ValueStringsAre(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^https://.+$`),
							"must be a valid URL starting with https://",
						),
					),
				},
				PlanModifiers: []planmodifier.Set{
					&setCustomRequired{},
				},
			},
			"allowed_logout_urls": schema.SetAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				MarkdownDescription: "Allowed logout URLs for OAuth2 client.",
				Validators: []validator.Set{
					setvalidator.ValueStringsAre(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^https://.+$`),
							"must be a valid URL starting with https://",
						),
					),
				},
				PlanModifiers: []planmodifier.Set{
					&setCustomRequired{},
				},
			},
			"enable_deduplication": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Enable deduplication.",
				Default:             booldefault.StaticBool(false),
			},
			"auto_login_after_register": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Automatically login after registration. Default is set to `false` while creating an app.",
				Default:             booldefault.StaticBool(false),
			},
			"enable_passwordless_auth": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Enable passwordless authentication. Default is set to `true` while creating an app.",
				Default:             booldefault.StaticBool(true),
			},
			"register_with_login_information": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Register with login information. Default is set to `false` while creating an app.",
				Default:             booldefault.StaticBool(false),
			},
			"allow_disposable_email": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Allow disposable email addresses. Default is set to `false` while creating an app.",
				Default:             booldefault.StaticBool(false),
			},
			"validate_phone_number": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "if enabled, phone number is validaed. Default is set to `false` while creating an app.",
				Default:             booldefault.StaticBool(false),
			},
			"fds_enabled": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Flag to enable or disable fraud detection system. By default, it is enabled when a client is created",
				Default:             booldefault.StaticBool(true),
			},
			"hosted_page_group": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Hosted page group.",
				Default:             stringdefault.StaticString("default"),
			},
			"client_name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Name of the client.",
			},
			"client_display_name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The display name of the client.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"company_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "The name of the company that the client belongs to.",
				PlanModifiers: []planmodifier.String{
					&stringCustomRequired{},
				},
			},
			"company_address": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "The company address.",
				PlanModifiers: []planmodifier.String{
					&stringCustomRequired{},
				},
			},
			"company_website": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "The URL of the company website.",
				PlanModifiers: []planmodifier.String{
					&stringCustomRequired{},
				},
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^https://.+$`),
						"must be a valid URL starting with https://",
					),
				},
			},
			"allowed_scopes": schema.SetAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				MarkdownDescription: "The URL of the company website. allowed_scopes is a required attribute. It must be provided in the main config or common_config",
				PlanModifiers: []planmodifier.Set{
					&setCustomRequired{},
				},
			},
			"response_types": schema.SetAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The response types of the client. The default value is set to `['code','token', 'id_token']`",
				Default: setdefault.StaticValue(basetypes.NewSetValueMust(types.StringType, []attr.Value{
					types.StringValue("code"), types.StringValue("token"), types.StringValue("id_token"),
				})),
			},
			"grant_types": schema.SetAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The grant types of the client. The default value is set to `['implicit','authorization_code', 'password', 'refresh_token']`",
				Default: setdefault.StaticValue(basetypes.NewSetValueMust(types.StringType, []attr.Value{
					types.StringValue("implicit"), types.StringValue("authorization_code"), types.StringValue("password"), types.StringValue("refresh_token"),
				})),
			},
			"login_providers": schema.SetAttribute{
				ElementType:         types.StringType,
				MarkdownDescription: "With this attribute one can setup login provider to the client.",
				Optional:            true,
			},
			"additional_access_token_payload": schema.SetAttribute{
				ElementType:         types.StringType,
				MarkdownDescription: "Access token payload definition.",
				Optional:            true,
			},
			"required_fields": schema.SetAttribute{
				ElementType:         types.StringType,
				MarkdownDescription: "The required fields while registering to the client.",
				Optional:            true,
			},
			"is_hybrid_app": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Flag to set if your app is hybrid or not. Default is set to `false`. Set to `true` to make your app hybrid.",
				Default:             booldefault.StaticBool(false),
			},
			"allowed_web_origins": schema.SetAttribute{
				ElementType:         types.StringType,
				MarkdownDescription: "List of the web origins allowed to access the client.",
				Optional:            true,
			},
			"allowed_origins": schema.SetAttribute{
				ElementType:         types.StringType,
				MarkdownDescription: "List of the origins allowed to access the client.",
				Optional:            true,
			},
			"mobile_settings": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"team_id": schema.StringAttribute{
						Optional: true,
					},
					"bundle_id": schema.StringAttribute{
						Optional: true,
					},
					"package_name": schema.StringAttribute{
						Optional: true,
					},
					"key_hash": schema.StringAttribute{
						Optional: true,
					},
				},
				Default: objectdefault.StaticValue(types.ObjectValueMust(
					map[string]attr.Type{
						"team_id":      types.StringType,
						"bundle_id":    types.StringType,
						"package_name": types.StringType,
						"key_hash":     types.StringType,
					},
					map[string]attr.Value{
						"team_id":      types.StringNull(),
						"bundle_id":    types.StringNull(),
						"package_name": types.StringNull(),
						"key_hash":     types.StringNull(),
					})),
			},
			"default_max_age": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The default maximum age for the token in seconds. Default is 86400 seconds (24 hours).",
				Default:             int64default.StaticInt64(86400),
			},
			"token_lifetime_in_seconds": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The lifetime of the token in seconds. Default is 86400 seconds (24 hours).",
				Default:             int64default.StaticInt64(86400),
			},
			"id_token_lifetime_in_seconds": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The lifetime of the id_token in seconds. Default is 86400 seconds (24 hours).",
				Default:             int64default.StaticInt64(86400),
			},
			"refresh_token_lifetime_in_seconds": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The lifetime of the refresh token in seconds. Default is 15780000 seconds.",
				Default:             int64default.StaticInt64(15780000),
			},
			"template_group_id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The id of the template group to be configured for commenication. Default is set to the system default group.",
				Default:             stringdefault.StaticString("default"),
			},
			"client_id": schema.StringAttribute{
				Optional: true,
				Computed: true,
				MarkdownDescription: "The client_id is the unqique identifier of the app. It's an optional attribute." +
					" If not provided, cidaas will gererate one for you and the state will be updated with the same",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"client_secret": schema.StringAttribute{
				Optional:  true,
				Computed:  true,
				Sensitive: true,
				MarkdownDescription: "The client_id is the unqique identifier of the app. It's an optional attribute." +
					" If not provided, cidaas will gererate one for you and the state will be updated with the same",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"policy_uri": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "The URL to the policy of a client.",
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^https://.+$`),
						"must be a valid URL starting with https://",
					),
				},
			},
			"tos_uri": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "The URL to the TOS of a client.",
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^https://.+$`),
						"must be a valid URL starting with https://",
					),
				},
			},
			"imprint_uri": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "The URL to the imprint page.",
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^https://.+$`),
						"must be a valid URL starting with https://",
					),
				},
			},
			"contacts": schema.SetAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				MarkdownDescription: "The contacts of the client.",
			},
			"token_endpoint_auth_method": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("client_secret_post"),
			},
			"token_endpoint_auth_signing_alg": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("RS256"),
			},
			"default_acr_values": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"editable": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
				MarkdownDescription: "Flag to define if your client is editable or not. Default is `true`.",
			},
			"web_message_uris": schema.SetAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				MarkdownDescription: "A list of URLs for web messages used.",
				Validators: []validator.Set{
					setvalidator.ValueStringsAre(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^https://.+$`),
							"must be a valid URL starting with https://",
						),
					),
				},
			},
			"social_providers": schema.ListNestedAttribute{
				Optional:            true,
				MarkdownDescription: "A list of social identity providers that users can authenticate with. Examples: Google, Facebook etc...",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"provider_name": schema.StringAttribute{
							Required: true,
						},
						"social_id": schema.StringAttribute{
							Required: true,
						},
					},
				},
			},
			"custom_providers": schema.ListNestedAttribute{
				Optional:            true,
				MarkdownDescription: "A list of custom identity providers that users can authenticate with. A custom provider can be created with the help of the resource cidaas_custom_provider.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"provider_name": schema.StringAttribute{
							Required: true,
						},
						"display_name": schema.StringAttribute{
							Optional: true,
						},
						"logo_url": schema.StringAttribute{
							Optional: true,
						},
						"type": schema.StringAttribute{
							Required: true,
						},
						"is_provider_visible": schema.BoolAttribute{
							Optional: true,
							Computed: true,
							Default:  booldefault.StaticBool(false),
						},
						"domains": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
					},
				},
			},
			"saml_providers": schema.ListNestedAttribute{
				Optional:            true,
				MarkdownDescription: "A list of SAML identity providers that users can authenticate with.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"provider_name": schema.StringAttribute{
							Required: true,
						},
						"display_name": schema.StringAttribute{
							Optional: true,
						},
						"logo_url": schema.StringAttribute{
							Optional: true,
						},
						"type": schema.StringAttribute{
							Required: true,
						},
						"is_provider_visible": schema.BoolAttribute{
							Optional: true,
							Computed: true,
							Default:  booldefault.StaticBool(false),
						},
						"domains": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
					},
				},
			},
			"ad_providers": schema.ListNestedAttribute{
				Optional:            true,
				MarkdownDescription: "A list of Active Directory identity providers that users can authenticate with.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"provider_name": schema.StringAttribute{
							Required: true,
						},
						"display_name": schema.StringAttribute{
							Optional: true,
						},
						"logo_url": schema.StringAttribute{
							Optional: true,
						},
						"type": schema.StringAttribute{
							Required: true,
						},
						"is_provider_visible": schema.BoolAttribute{
							Optional: true,
							Computed: true,
							Default:  booldefault.StaticBool(false),
						},
						"domains": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
					},
				},
			},
			"jwe_enabled": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Flag to specify whether JSON Web Encryption (JWE) should be enabled for encrypting data.",
				Default:             booldefault.StaticBool(false),
			},
			"user_consent": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Specifies whether user consent is required or not. Default is `false`",
				Default:             booldefault.StaticBool(false),
			},
			"allowed_groups": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"group_id": schema.StringAttribute{
							Required: true,
						},
						"roles": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"default_roles": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
					},
				},
			},
			"operations_allowed_groups": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"group_id": schema.StringAttribute{
							Required: true,
						},
						"roles": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"default_roles": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
					},
				},
			},
			"enabled": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(true),
			},
			"allowed_fields": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"always_ask_mfa": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
			},
			"smart_mfa": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
			},
			"allowed_mfa": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"captcha_ref": schema.StringAttribute{
				Optional: true,
			},
			"captcha_refs": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"consent_refs": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"communication_medium_verification": schema.StringAttribute{
				Optional: true,
			},
			"email_verification_required": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(true),
			},
			"mobile_number_verification_required": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
			},
			"allowed_roles": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"default_roles": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"enable_classical_provider": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(true),
			},
			"is_remember_me_selected": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(true),
			},
			"enable_bot_detection": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
			},
			"bot_provider": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("CIDAAS"),
			},
			"allow_guest_login_groups": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"group_id": schema.StringAttribute{
							Required: true,
						},
						"roles": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"default_roles": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
					},
				},
			},
			"is_login_success_page_enabled": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
			},
			"is_register_success_page_enabled": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
			},
			"group_ids": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"is_group_login_selection_enabled": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
			},
			"group_selection": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"always_show_group_selection": schema.BoolAttribute{
						Optional: true,
					},
					"selectable_groups": schema.SetAttribute{
						ElementType: types.StringType,
						Optional:    true,
					},
					"selectable_group_types": schema.SetAttribute{
						ElementType: types.StringType,
						Optional:    true,
					},
				},
			},
			"group_types": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"backchannel_logout_uri": schema.StringAttribute{
				Optional: true,
			},
			"post_logout_redirect_uris": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"logo_align": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"CENTER", "LEFT", "RIGHT"}...),
				},
				Default: stringdefault.StaticString("CENTER"),
			},
			"mfa": schema.SingleNestedAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Configuration settings for Multi-Factor Authentication (MFA).",
				Attributes: map[string]schema.Attribute{
					"setting": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Specifies the Multi-Factor Authentication (MFA) setting. Allowed values are 'OFF', 'ALWAYS', 'SMART', 'TIME_BASED' and 'SMART_PLUS_TIME_BASED'.",
						Validators: []validator.String{
							stringvalidator.OneOf([]string{"OFF", "ALWAYS", "SMART", "TIME_BASED", "SMART_PLUS_TIME_BASED"}...),
						},
					},
					"time_interval_in_seconds": schema.Int64Attribute{
						Optional:            true,
						MarkdownDescription: "Optional time interval in seconds for time-based Multi-Factor Authentication.",
					},
					"allowed_methods": schema.SetAttribute{
						ElementType:         types.StringType,
						Optional:            true,
						MarkdownDescription: "Optional set of allowed MFA methods.",
					},
				},
				Default: objectdefault.StaticValue(types.ObjectValueMust(
					map[string]attr.Type{
						"setting":                  types.StringType,
						"time_interval_in_seconds": types.Int64Type,
						"allowed_methods": types.SetType{
							ElemType: types.StringType,
						},
					},
					map[string]attr.Value{
						"setting":                  types.StringValue("OFF"),
						"time_interval_in_seconds": types.Int64Null(),
						"allowed_methods":          types.SetNull(types.StringType),
					})),
			},
			"webfinger": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("no_redirection"),
			},
			"logo_uri": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^https://.+$`),
						"must be a valid URL starting with https://",
					),
				},
			},
			"initiate_login_uri": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^https://.+$`),
						"must be a valid URL starting with https://",
					),
				},
			},
			"registration_client_uri": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^https://.+$`),
						"must be a valid URL starting with https://",
					),
				},
			},
			"registration_access_token": schema.StringAttribute{
				Optional: true,
			},
			"client_uri": schema.StringAttribute{
				Optional: true,
			},
			"jwks_uri": schema.StringAttribute{
				Optional: true,
			},
			"jwks": schema.StringAttribute{
				Optional: true,
			},
			"sector_identifier_uri": schema.StringAttribute{
				Optional: true,
			},
			"subject_type": schema.StringAttribute{
				Optional: true,
			},
			"id_token_signed_response_alg": schema.StringAttribute{
				Optional: true,
			},
			"id_token_encrypted_response_alg": schema.StringAttribute{
				Optional: true,
			},
			"id_token_encrypted_response_enc": schema.StringAttribute{
				Optional: true,
			},
			"userinfo_signed_response_alg": schema.StringAttribute{
				Optional: true,
			},
			"userinfo_encrypted_response_alg": schema.StringAttribute{
				Optional: true,
			},
			"userinfo_encrypted_response_enc": schema.StringAttribute{
				Optional: true,
			},
			"request_object_signing_alg": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"request_object_encryption_alg": schema.StringAttribute{
				Optional: true,
			},
			"request_object_encryption_enc": schema.StringAttribute{
				Optional: true,
			},
			"request_uris": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"description": schema.StringAttribute{
				Optional: true,
			},
			"default_scopes": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"pending_scopes": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"consent_page_group": schema.StringAttribute{
				Optional: true,
			},
			"password_policy_ref": schema.StringAttribute{
				Optional: true,
			},
			"blocking_mechanism_ref": schema.StringAttribute{
				Optional: true,
			},
			"sub": schema.StringAttribute{
				Optional: true,
			},
			"role": schema.StringAttribute{
				Optional: true,
			},
			"mfa_configuration": schema.StringAttribute{
				Optional: true,
			},
			"suggest_mfa": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"login_spi": schema.SingleNestedAttribute{
				Optional:            true,
				MarkdownDescription: "A map defining the Login SPI configuration.",
				Attributes: map[string]schema.Attribute{
					"oauth_client_id": schema.StringAttribute{
						Optional: true,
					},
					"spi_url": schema.StringAttribute{
						Optional: true,
					},
				},
			},
			"background_uri": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "The URL to the background image of the client.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^https://.+$`),
						"must be a valid URL starting with https://",
					),
				},
			},
			"video_url": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "The URL to the video of the client.",
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^https://.+$`),
						"must be a valid URL starting with https://",
					),
				},
			},
			"bot_captcha_ref": schema.StringAttribute{
				Optional: true,
			},
			"application_meta_data": schema.MapAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				MarkdownDescription: "A map to add metadata of a client.",
			},
			"allow_guest_login": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				MarkdownDescription: "Flag to specify whether guest users are allowed to access functionalities of the client." +
					" Default is set to `false`",
				Default: booldefault.StaticBool(false),
			},
			"common_configs": getCommonConfig(),
		},
	}
}

func getCommonConfig() schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Optional: true,
		MarkdownDescription: "The `common_configs` attribute is used for sharing the same configuration across multiple cidaas_app resources." +
			" It is a map of some attributes from the main configuration." +
			" Please check the list of the attributes that it supports in the common_confis section." +
			" if an attribute is available both common_config and main config then attribute from the main config will be considered to create an app",
		PlanModifiers: []planmodifier.Object{
			&commonConfigConflictVerifier{},
		},
		Attributes: map[string]schema.Attribute{
			"company_name": schema.StringAttribute{
				Optional: true,
			},
			"company_website": schema.StringAttribute{
				Optional: true,
			},
			"client_type": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.OneOf([]string{
						"SINGLE_PAGE", "REGULAR_WEB", "NON_INTERACTIVE",
						"IOS", "ANDROID", "WINDOWS_MOBILE", "DESKTOP", "MOBILE", "DEVICE", "THIRD_PARTY",
					}...),
				},
			},
			"company_address": schema.StringAttribute{
				Optional: true,
			},
			"allowed_scopes": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"redirect_uris": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Validators: []validator.Set{
					setvalidator.ValueStringsAre(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^https://.+$`),
							"must be a valid URL starting with https://",
						),
					),
				},
			},
			"allowed_logout_urls": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Validators: []validator.Set{
					setvalidator.ValueStringsAre(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^https://.+$`),
							"must be a valid URL starting with https://",
						),
					),
				},
			},
			"allowed_web_origins": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"allowed_origins": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"login_providers": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"default_scopes": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"pending_scopes": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"allowed_mfa": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"allowed_roles": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"default_roles": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"social_providers": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"provider_name": schema.StringAttribute{
							Required: true,
						},
						"social_id": schema.StringAttribute{
							Required: true,
						},
					},
				},
			},
			"custom_providers": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"provider_name": schema.StringAttribute{
							Required: true,
						},
						"display_name": schema.StringAttribute{
							Optional: true,
						},
						"logo_url": schema.StringAttribute{
							Optional: true,
						},
						"type": schema.StringAttribute{
							Required: true,
						},
						"is_provider_visible": schema.BoolAttribute{
							Optional: true,
							Computed: true,
							Default:  booldefault.StaticBool(false),
						},
						"domains": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
					},
				},
			},
			"saml_providers": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"provider_name": schema.StringAttribute{
							Required: true,
						},
						"display_name": schema.StringAttribute{
							Optional: true,
						},
						"logo_url": schema.StringAttribute{
							Optional: true,
						},
						"type": schema.StringAttribute{
							Required: true,
						},
						"is_provider_visible": schema.BoolAttribute{
							Optional: true,
							Computed: true,
							Default:  booldefault.StaticBool(false),
						},
						"domains": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
					},
				},
			},
			"ad_providers": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"provider_name": schema.StringAttribute{
							Required: true,
						},
						"display_name": schema.StringAttribute{
							Optional: true,
						},
						"logo_url": schema.StringAttribute{
							Optional: true,
						},
						"type": schema.StringAttribute{
							Required: true,
						},
						"is_provider_visible": schema.BoolAttribute{
							Optional: true,
							Computed: true,
							Default:  booldefault.StaticBool(false),
						},
						"domains": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
					},
				},
			},
			"allowed_groups": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"group_id": schema.StringAttribute{
							Required: true,
						},
						"roles": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"default_roles": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
					},
				},
			},
			"operations_allowed_groups": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"group_id": schema.StringAttribute{
							Required: true,
						},
						"roles": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"default_roles": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
					},
				},
			},
			"accent_color": schema.StringAttribute{
				Optional: true,
			},
			"primary_color": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^#([a-fA-F0-9]{6}|[a-fA-F0-9]{3})$`),
						"must be a valid hex color",
					),
				},
			},
			"media_type": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"VIDEO", "IMAGE"}...),
				},
			},
			"hosted_page_group": schema.StringAttribute{
				Optional: true,
			},
			"template_group_id": schema.StringAttribute{
				Optional: true,
			},
			"bot_provider": schema.StringAttribute{
				Optional: true,
			},
			"logo_align": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"CENTER", "LEFT", "RIGHT"}...),
				},
			},
			"webfinger": schema.StringAttribute{
				Optional: true,
			},
			"default_max_age": schema.Int64Attribute{
				Optional: true,
			},
			"token_lifetime_in_seconds": schema.Int64Attribute{
				Optional: true,
			},
			"id_token_lifetime_in_seconds": schema.Int64Attribute{
				Optional: true,
			},
			"refresh_token_lifetime_in_seconds": schema.Int64Attribute{
				Optional: true,
			},
			"allow_guest_login": schema.BoolAttribute{
				Optional: true,
			},
			"enable_deduplication": schema.BoolAttribute{
				Optional: true,
			},
			"auto_login_after_register": schema.BoolAttribute{
				Optional: true,
			},
			"enable_passwordless_auth": schema.BoolAttribute{
				Optional: true,
			},
			"register_with_login_information": schema.BoolAttribute{
				Optional: true,
			},
			"fds_enabled": schema.BoolAttribute{
				Optional: true,
			},
			"is_hybrid_app": schema.BoolAttribute{
				Optional: true,
			},
			"editable": schema.BoolAttribute{
				Optional: true,
			},
			"enabled": schema.BoolAttribute{
				Optional: true,
			},
			"always_ask_mfa": schema.BoolAttribute{
				Optional: true,
			},
			"email_verification_required": schema.BoolAttribute{
				Optional: true,
			},
			"enable_classical_provider": schema.BoolAttribute{
				Optional: true,
			},
			"is_remember_me_selected": schema.BoolAttribute{
				Optional: true,
			},
			"response_types": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"grant_types": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"allow_login_with": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Validators: []validator.Set{
					setvalidator.ValueStringsAre(
						stringvalidator.OneOf([]string{"EMAIL", "MOBILE", "USER_NAME"}...),
					),
				},
			},
			"mfa": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"setting": schema.StringAttribute{
						Optional: true,
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf([]string{"OFF", "ALWAYS", "SMART", "TIME_BASED", "SMART_PLUS_TIME_BASED"}...),
						},
						Default: stringdefault.StaticString("OFF"),
					},
					"time_interval_in_seconds": schema.Int64Attribute{
						Optional: true,
					},
					"allowed_methods": schema.SetAttribute{
						ElementType: types.StringType,
						Optional:    true,
					},
				},
			},
		},
	}
}
