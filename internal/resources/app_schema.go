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
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"client_type": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					&stringCustomRequired{},
				},
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"SINGLE_PAGE", "REGULAR_WEB", "NON_INTERACTIVE",
						"IOS", "ANDROID", "WINDOWS_MOBILE", "DESKTOP", "MOBILE", "DEVICE", "THIRD_PARTY"}...),
				},
			},
			"accent_color": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("#ef4923"),
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
				Default:  stringdefault.StaticString("#f7941d"),
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
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"VIDEO", "IMAGE"}...),
				},
				Default: stringdefault.StaticString("IMAGE"),
			},
			"content_align": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"CENTER", "LEFT", "RIGHT"}...),
				},
				Default: stringdefault.StaticString("CENTER"),
			},
			"allow_login_with": schema.SetAttribute{
				ElementType: types.StringType,
				Computed:    true,
				Optional:    true,
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
				PlanModifiers: []planmodifier.Set{
					&setCustomRequired{},
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
				PlanModifiers: []planmodifier.Set{
					&setCustomRequired{},
				},
			},
			"enable_deduplication": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
			},
			"auto_login_after_register": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
			},
			"enable_passwordless_auth": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(true),
			},
			"register_with_login_information": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
			},
			"allow_disposable_email": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
			},
			"validate_phone_number": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
			},
			"fds_enabled": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(true),
			},
			"hosted_page_group": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("default"),
			},
			"client_name": schema.StringAttribute{
				Required: true,
			},
			"client_display_name": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"company_name": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					&stringCustomRequired{},
				},
			},
			"company_address": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					&stringCustomRequired{},
				},
			},
			"company_website": schema.StringAttribute{
				Optional: true,
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
				ElementType: types.StringType,
				Optional:    true,
				PlanModifiers: []planmodifier.Set{
					&setCustomRequired{},
				},
			},
			"response_types": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Default: setdefault.StaticValue(basetypes.NewSetValueMust(types.StringType, []attr.Value{
					types.StringValue("code"), types.StringValue("token"), types.StringValue("id_token"),
				})),
			},
			"grant_types": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Default: setdefault.StaticValue(basetypes.NewSetValueMust(types.StringType, []attr.Value{
					types.StringValue("implicit"), types.StringValue("authorization_code"), types.StringValue("password"), types.StringValue("refresh_token"),
				})),
			},
			"login_providers": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"additional_access_token_payload": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"required_fields": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"is_hybrid_app": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
			},
			"allowed_web_origins": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"allowed_origins": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
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
				Optional: true,
				Computed: true,
				Default:  int64default.StaticInt64(86400),
			},
			"token_lifetime_in_seconds": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				Default:  int64default.StaticInt64(86400),
			},
			"id_token_lifetime_in_seconds": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				Default:  int64default.StaticInt64(86400),
			},
			"refresh_token_lifetime_in_seconds": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				Default:  int64default.StaticInt64(15780000),
			},
			"template_group_id": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("default"),
			},
			"client_id": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"client_secret": schema.StringAttribute{
				Optional:  true,
				Computed:  true,
				Sensitive: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"policy_uri": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^https://.+$`),
						"must be a valid URL starting with https://",
					),
				},
			},
			"tos_uri": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^https://.+$`),
						"must be a valid URL starting with https://",
					),
				},
			},
			"imprint_uri": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^https://.+$`),
						"must be a valid URL starting with https://",
					),
				},
			},
			"contacts": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
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
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(true),
			},
			"web_message_uris": schema.SetAttribute{
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
						"display_name": schema.StringAttribute{
							Optional: true,
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
			"app_owner": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"jwe_enabled": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
			},
			"user_consent": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
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
			"deleted": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
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
			"admin_client": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
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
				Optional: true,
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"setting": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							stringvalidator.OneOf([]string{"OFF", "ALWAYS", "SMART", "TIME_BASED", "SMART_PLUS_TIME_BASED"}...),
						},
					},
					"time_interval_in_seconds": schema.Int64Attribute{
						Optional: true,
					},
					"allowed_methods": schema.SetAttribute{
						ElementType: types.StringType,
						Optional:    true,
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
				Optional: true,
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
				Optional: true,
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
				Optional: true,
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
				ElementType: types.StringType,
				Optional:    true,
			},
			"allow_guest_login": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
			},
			// "common_configs": getCommonConfig(),
			"created_at": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func getCommonConfig() schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Optional: true,
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
					stringvalidator.OneOf([]string{"SINGLE_PAGE", "REGULAR_WEB", "NON_INTERACTIVE",
						"IOS", "ANDROID", "WINDOWS_MOBILE", "DESKTOP", "MOBILE", "DEVICE", "THIRD_PARTY"}...),
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
						"display_name": schema.StringAttribute{
							Optional: true,
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
