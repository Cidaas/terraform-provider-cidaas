package resources

import (
	"context"
	"fmt"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/cidaas"
	"github.com/Cidaas/terraform-provider-cidaas/helpers/util"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type AppConfig struct {
	ID                               types.String `tfsdk:"id"`
	ClientType                       types.String `tfsdk:"client_type"`
	AccentColor                      types.String `tfsdk:"accent_color"`
	PrimaryColor                     types.String `tfsdk:"primary_color"`
	MediaType                        types.String `tfsdk:"media_type"`
	ContentAlign                     types.String `tfsdk:"content_align"`
	AllowLoginWith                   types.Set    `tfsdk:"allow_login_with"`
	RedirectURIS                     types.Set    `tfsdk:"redirect_uris"`
	AllowedLogoutUrls                types.Set    `tfsdk:"allowed_logout_urls"`
	EnableDeduplication              types.Bool   `tfsdk:"enable_deduplication"`
	AutoLoginAfterRegister           types.Bool   `tfsdk:"auto_login_after_register"`
	EnablePasswordlessAuth           types.Bool   `tfsdk:"enable_passwordless_auth"`
	RegisterWithLoginInformation     types.Bool   `tfsdk:"register_with_login_information"`
	AllowDisposableEmail             types.Bool   `tfsdk:"allow_disposable_email"`
	ValidatePhoneNumber              types.Bool   `tfsdk:"validate_phone_number"`
	FdsEnabled                       types.Bool   `tfsdk:"fds_enabled"`
	HostedPageGroup                  types.String `tfsdk:"hosted_page_group"`
	ClientName                       types.String `tfsdk:"client_name"`
	ClientDisplayName                types.String `tfsdk:"client_display_name"`
	CompanyName                      types.String `tfsdk:"company_name"`
	CompanyAddress                   types.String `tfsdk:"company_address"`
	CompanyWebsite                   types.String `tfsdk:"company_website"`
	AllowedScopes                    types.Set    `tfsdk:"allowed_scopes"`
	ResponseTypes                    types.Set    `tfsdk:"response_types"`
	GrantTypes                       types.Set    `tfsdk:"grant_types"`
	LoginProviders                   types.Set    `tfsdk:"login_providers"`
	AdditionalAccessTokenPayload     types.Set    `tfsdk:"additional_access_token_payload"`
	RequiredFields                   types.Set    `tfsdk:"required_fields"`
	IsHybridApp                      types.Bool   `tfsdk:"is_hybrid_app"`
	AllowedWebOrigins                types.Set    `tfsdk:"allowed_web_origins"`
	AllowedOrigins                   types.Set    `tfsdk:"allowed_origins"`
	DefaultMaxAge                    types.Int64  `tfsdk:"default_max_age"`
	TokenLifetimeInSeconds           types.Int64  `tfsdk:"token_lifetime_in_seconds"`
	IDTokenLifetimeInSeconds         types.Int64  `tfsdk:"id_token_lifetime_in_seconds"`
	RefreshTokenLifetimeInSeconds    types.Int64  `tfsdk:"refresh_token_lifetime_in_seconds"`
	TemplateGroupID                  types.String `tfsdk:"template_group_id"`
	ClientID                         types.String `tfsdk:"client_id"`
	ClientSecret                     types.String `tfsdk:"client_secret"`
	PolicyURI                        types.String `tfsdk:"policy_uri"`
	TosURI                           types.String `tfsdk:"tos_uri"`
	ImprintURI                       types.String `tfsdk:"imprint_uri"`
	Contacts                         types.Set    `tfsdk:"contacts"`
	TokenEndpointAuthMethod          types.String `tfsdk:"token_endpoint_auth_method"`
	TokenEndpointAuthSigningAlg      types.String `tfsdk:"token_endpoint_auth_signing_alg"`
	DefaultAcrValues                 types.Set    `tfsdk:"default_acr_values"`
	Editable                         types.Bool   `tfsdk:"editable"`
	WebMessageUris                   types.Set    `tfsdk:"web_message_uris"`
	AppOwner                         types.String `tfsdk:"app_owner"`
	JweEnabled                       types.Bool   `tfsdk:"jwe_enabled"`
	UserConsent                      types.Bool   `tfsdk:"user_consent"`
	Deleted                          types.Bool   `tfsdk:"deleted"`
	Enabled                          types.Bool   `tfsdk:"enabled"`
	AllowedFields                    types.Set    `tfsdk:"allowed_fields"`
	AlwaysAskMfa                     types.Bool   `tfsdk:"always_ask_mfa"`
	SmartMfa                         types.Bool   `tfsdk:"smart_mfa"`
	AllowedMfa                       types.Set    `tfsdk:"allowed_mfa"`
	CaptchaRef                       types.String `tfsdk:"captcha_ref"`
	CaptchaRefs                      types.Set    `tfsdk:"captcha_refs"`
	ConsentRefs                      types.Set    `tfsdk:"consent_refs"`
	CommunicationMediumVerification  types.String `tfsdk:"communication_medium_verification"`
	EmailVerificationRequired        types.Bool   `tfsdk:"email_verification_required"`
	MobileNumberVerificationRequired types.Bool   `tfsdk:"mobile_number_verification_required"`
	AllowedRoles                     types.Set    `tfsdk:"allowed_roles"`
	DefaultRoles                     types.Set    `tfsdk:"default_roles"`
	EnableClassicalProvider          types.Bool   `tfsdk:"enable_classical_provider"`
	IsRememberMeSelected             types.Bool   `tfsdk:"is_remember_me_selected"`
	EnableBotDetection               types.Bool   `tfsdk:"enable_bot_detection"`
	BotProvider                      types.String `tfsdk:"bot_provider"`
	IsLoginSuccessPageEnabled        types.Bool   `tfsdk:"is_login_success_page_enabled"`
	IsRegisterSuccessPageEnabled     types.Bool   `tfsdk:"is_register_success_page_enabled"`
	GroupIDs                         types.Set    `tfsdk:"group_ids"`
	AdminClient                      types.Bool   `tfsdk:"admin_client"`
	IsGroupLoginSelectionEnabled     types.Bool   `tfsdk:"is_group_login_selection_enabled"`
	GroupTypes                       types.Set    `tfsdk:"group_types"`
	BackchannelLogoutURI             types.String `tfsdk:"backchannel_logout_uri"`
	PostLogoutRedirectUris           types.Set    `tfsdk:"post_logout_redirect_uris"`
	LogoAlign                        types.String `tfsdk:"logo_align"`
	Webfinger                        types.String `tfsdk:"webfinger"`
	LogoURI                          types.String `tfsdk:"logo_uri"`
	InitiateLoginURI                 types.String `tfsdk:"initiate_login_uri"`
	RegistrationClientURI            types.String `tfsdk:"registration_client_uri"`
	RegistrationAccessToken          types.String `tfsdk:"registration_access_token"`
	ClientURI                        types.String `tfsdk:"client_uri"`
	JwksURI                          types.String `tfsdk:"jwks_uri"`
	Jwks                             types.String `tfsdk:"jwks"`
	SectorIdentifierURI              types.String `tfsdk:"sector_identifier_uri"`
	SubjectType                      types.String `tfsdk:"subject_type"`
	IDTokenSignedResponseAlg         types.String `tfsdk:"id_token_signed_response_alg"`
	IDTokenEncryptedResponseAlg      types.String `tfsdk:"id_token_encrypted_response_alg"`
	IDTokenEncryptedResponseEnc      types.String `tfsdk:"id_token_encrypted_response_enc"`
	UserinfoSignedResponseAlg        types.String `tfsdk:"userinfo_signed_response_alg"`
	UserinfoEncryptedResponseAlg     types.String `tfsdk:"userinfo_encrypted_response_alg"`
	UserinfoEncryptedResponseEnc     types.String `tfsdk:"userinfo_encrypted_response_enc"`
	RequestObjectSigningAlg          types.String `tfsdk:"request_object_signing_alg"`
	RequestObjectEncryptionAlg       types.String `tfsdk:"request_object_encryption_alg"`
	RequestObjectEncryptionEnc       types.String `tfsdk:"request_object_encryption_enc"`
	RequestUris                      types.Set    `tfsdk:"request_uris"`
	Description                      types.String `tfsdk:"description"`
	DefaultScopes                    types.Set    `tfsdk:"default_scopes"`
	PendingScopes                    types.Set    `tfsdk:"pending_scopes"`
	ConsentPageGroup                 types.String `tfsdk:"consent_page_group"`
	PasswordPolicyRef                types.String `tfsdk:"password_policy_ref"`
	BlockingMechanismRef             types.String `tfsdk:"blocking_mechanism_ref"`
	Sub                              types.String `tfsdk:"sub"`
	Role                             types.String `tfsdk:"role"`
	MfaConfiguration                 types.String `tfsdk:"mfa_configuration"`
	SuggestMfa                       types.Set    `tfsdk:"suggest_mfa"`
	AllowGuestLogin                  types.Bool   `tfsdk:"allow_guest_login"`
	BackgroundURI                    types.String `tfsdk:"background_uri"`
	VideoURL                         types.String `tfsdk:"video_url"`
	BotCaptchaRef                    types.String `tfsdk:"bot_captcha_ref"`
	ApplicationMetaData              types.Map    `tfsdk:"application_meta_data"`
	CreatedAt                        types.String `tfsdk:"created_at"`
	UpdatedAt                        types.String `tfsdk:"updated_at"`

	SocialProviders         types.List   `tfsdk:"social_providers"`
	CustomProviders         types.List   `tfsdk:"custom_providers"`
	SamlProviders           types.List   `tfsdk:"saml_providers"`
	AdProviders             types.List   `tfsdk:"ad_providers"`
	AllowedGroups           types.List   `tfsdk:"allowed_groups"`
	OperationsAllowedGroups types.List   `tfsdk:"operations_allowed_groups"`
	AllowGuestLoginGroups   types.List   `tfsdk:"allow_guest_login_groups"`
	LoginSpi                types.Object `tfsdk:"login_spi"`
	GroupSelection          types.Object `tfsdk:"group_selection"`
	Mfa                     types.Object `tfsdk:"mfa"`
	MobileSettings          types.Object `tfsdk:"mobile_settings"`
	CommonConfigs           types.Object `tfsdk:"common_configs"`

	socialProviders         []*SocialProviderData
	customProviders         []*ProviderMetadData
	samlProviders           []*ProviderMetadData
	adProviders             []*ProviderMetadData
	allowedGroups           []*AllowedGroups
	operationsAllowedGroups []*AllowedGroups
	allowGuestLoginGroups   []*AllowedGroups

	loginSpi       *LoginSPI
	groupSelection *GroupSelection
	mfa            *MfaOption
	mobileSettings *AppMobileSettings
	commonConfigs  *CommonConfigs
}

type AllowedGroups struct {
	GroupID      types.String `tfsdk:"group_id"`
	Roles        types.Set    `tfsdk:"roles"`
	DefaultRoles types.Set    `tfsdk:"default_roles"`
}

type LoginSPI struct {
	OauthClientID types.String `tfsdk:"oauth_client_id"`
	SpiURL        types.String `tfsdk:"spi_url"`
}

type AppMobileSettings struct {
	TeamID      types.String `tfsdk:"team_id"`
	BundleID    types.String `tfsdk:"bundle_id"`
	PackageName types.String `tfsdk:"package_name"`
	KeyHash     types.String `tfsdk:"key_hash"`
}

type GroupSelection struct {
	AlwaysShowGroupSelection types.Bool `tfsdk:"always_show_group_selection"`
	SelectableGroups         types.Set  `tfsdk:"selectable_groups"`
	SelectableGroupTypes     types.Set  `tfsdk:"selectable_group_types"`
}

type MfaOption struct {
	Setting               types.String `tfsdk:"setting"`
	TimeIntervalInSeconds types.Int64  `tfsdk:"time_interval_in_seconds"`
	AllowedMethods        types.Set    `tfsdk:"allowed_methods"`
}

type SocialProviderData struct {
	ProviderName types.String `tfsdk:"provider_name"`
	SocialID     types.String `tfsdk:"social_id"`
}

type ProviderMetadData struct {
	LogoURL           types.String `tfsdk:"logo_url"`
	ProviderName      types.String `tfsdk:"provider_name"`
	DisplayName       types.String `tfsdk:"display_name"`
	Type              types.String `tfsdk:"type"`
	IsProviderVisible types.Bool   `tfsdk:"is_provider_visible"`
	Domains           types.Set    `tfsdk:"domains"`
}

type CommonConfigs struct {
	CompanyName    types.String `tfsdk:"company_name"`
	CompanyWebsite types.String `tfsdk:"company_website"`
	ClientType     types.String `tfsdk:"client_type"`
	CompanyAddress types.String `tfsdk:"company_address"`

	AllowedScopes           types.Set  `tfsdk:"allowed_scopes"`
	RedirectUris            types.Set  `tfsdk:"redirect_uris"`
	AllowedLogoutUrls       types.Set  `tfsdk:"allowed_logout_urls"`
	AllowedWebOrigins       types.Set  `tfsdk:"allowed_web_origins"`
	AllowedOrigins          types.Set  `tfsdk:"allowed_origins"`
	LoginProviders          types.Set  `tfsdk:"login_providers"`
	DefaultScopes           types.Set  `tfsdk:"default_scopes"`
	PendingScopes           types.Set  `tfsdk:"pending_scopes"`
	AllowedMfa              types.Set  `tfsdk:"allowed_mfa"`
	AllowedRoles            types.Set  `tfsdk:"allowed_roles"`
	DefaultRoles            types.Set  `tfsdk:"default_roles"`
	SocialProviders         types.List `tfsdk:"social_providers"`
	CustomProviders         types.List `tfsdk:"custom_providers"`
	SamlProviders           types.List `tfsdk:"saml_providers"`
	AdProviders             types.List `tfsdk:"ad_providers"`
	AllowedGroups           types.List `tfsdk:"allowed_groups"`
	OperationsAllowedGroups types.List `tfsdk:"operations_allowed_groups"`

	// attributes with default value
	AccentColor                   types.String `tfsdk:"accent_color"`
	PrimaryColor                  types.String `tfsdk:"primary_color"`
	MediaType                     types.String `tfsdk:"media_type"`
	HostedPageGroup               types.String `tfsdk:"hosted_page_group"`
	TemplateGroupID               types.String `tfsdk:"template_group_id"`
	BotProvider                   types.String `tfsdk:"bot_provider"`
	LogoAlign                     types.String `tfsdk:"logo_align"`
	Webfinger                     types.String `tfsdk:"webfinger"`
	DefaultMaxAge                 types.Int64  `tfsdk:"default_max_age"`
	TokenLifetimeInSeconds        types.Int64  `tfsdk:"token_lifetime_in_seconds"`
	IDTokenLifetimeInSeconds      types.Int64  `tfsdk:"id_token_lifetime_in_seconds"`
	RefreshTokenLifetimeInSeconds types.Int64  `tfsdk:"refresh_token_lifetime_in_seconds"`
	AllowGuestLogin               types.Bool   `tfsdk:"allow_guest_login"`
	EnableDeduplication           types.Bool   `tfsdk:"enable_deduplication"`
	AutoLoginAfterRegister        types.Bool   `tfsdk:"auto_login_after_register"`
	EnablePasswordlessAuth        types.Bool   `tfsdk:"enable_passwordless_auth"`
	RegisterWithLoginInformation  types.Bool   `tfsdk:"register_with_login_information"`
	FdsEnabled                    types.Bool   `tfsdk:"fds_enabled"`
	IsHybridApp                   types.Bool   `tfsdk:"is_hybrid_app"`
	Editable                      types.Bool   `tfsdk:"editable"`
	Enabled                       types.Bool   `tfsdk:"enabled"`
	AlwaysAskMfa                  types.Bool   `tfsdk:"always_ask_mfa"`
	EmailVerificationRequired     types.Bool   `tfsdk:"email_verification_required"`
	EnableClassicalProvider       types.Bool   `tfsdk:"enable_classical_provider"`
	IsRememberMeSelected          types.Bool   `tfsdk:"is_remember_me_selected"`
	ResponseTypes                 types.Set    `tfsdk:"response_types"`
	GrantTypes                    types.Set    `tfsdk:"grant_types"`
	AllowLoginWith                types.Set    `tfsdk:"allow_login_with"`
	Mfa                           types.Object `tfsdk:"mfa"`
}

func (w *AppConfig) ExtractAppConfigs(ctx context.Context) diag.Diagnostics {
	var diags diag.Diagnostics
	if !w.LoginSpi.IsNull() {
		w.loginSpi = &LoginSPI{}
		diags = w.LoginSpi.As(ctx, w.loginSpi, basetypes.ObjectAsOptions{})
	}
	if !w.GroupSelection.IsNull() {
		w.groupSelection = &GroupSelection{}
		diags = w.GroupSelection.As(ctx, w.groupSelection, basetypes.ObjectAsOptions{})
	}
	if !w.Mfa.IsNull() {
		w.mfa = &MfaOption{}
		diags = w.Mfa.As(ctx, w.mfa, basetypes.ObjectAsOptions{})
	}
	if !w.MobileSettings.IsNull() {
		w.mobileSettings = &AppMobileSettings{}
		diags = w.MobileSettings.As(ctx, w.mobileSettings, basetypes.ObjectAsOptions{})
	}
	if !w.CommonConfigs.IsNull() {
		w.commonConfigs = &CommonConfigs{}
		diags = w.CommonConfigs.As(ctx, w.commonConfigs, basetypes.ObjectAsOptions{})
	}
	if !w.SocialProviders.IsNull() {
		w.socialProviders = make([]*SocialProviderData, 0, len(w.SocialProviders.Elements()))
		diags = w.SocialProviders.ElementsAs(ctx, &w.socialProviders, false)
	}
	if !w.CustomProviders.IsNull() {
		w.customProviders = make([]*ProviderMetadData, 0, len(w.CustomProviders.Elements()))
		diags = w.CustomProviders.ElementsAs(ctx, &w.customProviders, false)
	}
	if !w.SamlProviders.IsNull() {
		w.samlProviders = make([]*ProviderMetadData, 0, len(w.SamlProviders.Elements()))
		diags = w.SamlProviders.ElementsAs(ctx, &w.samlProviders, false)
	}
	if !w.AdProviders.IsNull() {
		w.adProviders = make([]*ProviderMetadData, 0, len(w.AdProviders.Elements()))
		diags = w.AdProviders.ElementsAs(ctx, &w.adProviders, false)
	}
	if !w.AllowedGroups.IsNull() {
		w.allowedGroups = make([]*AllowedGroups, 0, len(w.AllowedGroups.Elements()))
		diags = w.AllowedGroups.ElementsAs(ctx, &w.allowedGroups, false)
	}
	if !w.OperationsAllowedGroups.IsNull() {
		w.operationsAllowedGroups = make([]*AllowedGroups, 0, len(w.OperationsAllowedGroups.Elements()))
		diags = w.OperationsAllowedGroups.ElementsAs(ctx, &w.operationsAllowedGroups, false)
	}
	if !w.AllowGuestLoginGroups.IsNull() {
		w.allowGuestLoginGroups = make([]*AllowedGroups, 0, len(w.AllowGuestLoginGroups.Elements()))
		diags = w.AllowGuestLoginGroups.ElementsAs(ctx, &w.allowGuestLoginGroups, false)
	}
	return diags
}

func extractSetValues(ctx context.Context, dest *basetypes.SetValue, src []string) diag.Diagnostics {
	if len(src) > 0 {
		result, diags := types.SetValueFrom(ctx, types.StringType, src)
		if diags.HasError() {
			return diags
		}
		*dest = result
	}
	return nil
}

func prepareAppModel(ctx context.Context, plan AppConfig) (*cidaas.AppModel, diag.Diagnostics) {

	setStringValue := func(value, commonConfigValue types.String, target *string) {
		if !value.IsNull() {
			*target = value.ValueString()
		} else if !commonConfigValue.IsNull() {
			*target = commonConfigValue.ValueString()
		}
	}

	setBoolValue := func(value, commonConfigValue types.Bool, target *bool) {
		if !value.IsNull() {
			*target = value.ValueBool()
		} else if !commonConfigValue.IsNull() {
			*target = commonConfigValue.ValueBool()
		}
	}

	setInt64Value := func(value, commonConfigValue types.Int64, target *int64) {
		if !value.IsNull() {
			*target = value.ValueInt64()
		} else if !commonConfigValue.IsNull() {
			*target = commonConfigValue.ValueInt64()
		}
	}

	setSetValues := func(ctx context.Context, planValue, commonConfigValue basetypes.SetValue, target *[]string) diag.Diagnostics {
		if !planValue.IsNull() {
			if len(planValue.Elements()) > 0 {
				return planValue.ElementsAs(ctx, target, false)
			}
		} else if !commonConfigValue.IsNull() {
			if len(commonConfigValue.Elements()) > 0 {
				return commonConfigValue.ElementsAs(ctx, target, false)
			}
		}
		return nil
	}

	commonConfigs := &CommonConfigs{}
	if plan.commonConfigs != nil {
		commonConfigs = plan.commonConfigs
	}

	app := cidaas.AppModel{}

	setStringValue(plan.CompanyName, commonConfigs.CompanyName, &app.CompanyName)
	setStringValue(plan.CompanyAddress, commonConfigs.CompanyAddress, &app.CompanyAddress)
	setStringValue(plan.CompanyWebsite, commonConfigs.CompanyWebsite, &app.CompanyWebsite)
	setStringValue(plan.ClientType, commonConfigs.ClientType, &app.ClientType)
	setStringValue(plan.AccentColor, commonConfigs.AccentColor, &app.AccentColor)
	setStringValue(plan.PrimaryColor, commonConfigs.PrimaryColor, &app.PrimaryColor)
	setStringValue(plan.MediaType, commonConfigs.MediaType, &app.MediaType)
	setStringValue(plan.ContentAlign, types.StringNull(), &app.ContentAlign)
	setBoolValue(plan.EnableDeduplication, commonConfigs.EnableDeduplication, &app.EnableDeduplication)
	setBoolValue(plan.AutoLoginAfterRegister, commonConfigs.AutoLoginAfterRegister, &app.AutoLoginAfterRegister)
	setBoolValue(plan.EnablePasswordlessAuth, commonConfigs.EnablePasswordlessAuth, &app.EnablePasswordlessAuth)
	setBoolValue(plan.RegisterWithLoginInformation, commonConfigs.RegisterWithLoginInformation, &app.RegisterWithLoginInformation)
	setBoolValue(plan.AllowDisposableEmail, types.BoolNull(), &app.AllowDisposableEmail)
	setBoolValue(plan.ValidatePhoneNumber, types.BoolNull(), &app.ValidatePhoneNumber)
	setBoolValue(plan.FdsEnabled, commonConfigs.FdsEnabled, &app.FdsEnabled)
	setStringValue(plan.HostedPageGroup, commonConfigs.HostedPageGroup, &app.HostedPageGroup)
	setStringValue(plan.ClientName, types.StringNull(), &app.ClientName)
	setStringValue(plan.ClientDisplayName, types.StringNull(), &app.ClientDisplayName)
	setBoolValue(plan.IsHybridApp, commonConfigs.IsHybridApp, &app.IsHybridApp)
	setInt64Value(plan.DefaultMaxAge, commonConfigs.DefaultMaxAge, &app.DefaultMaxAge)
	setInt64Value(plan.TokenLifetimeInSeconds, commonConfigs.TokenLifetimeInSeconds, &app.TokenLifetimeInSeconds)
	setInt64Value(plan.IDTokenLifetimeInSeconds, commonConfigs.IDTokenLifetimeInSeconds, &app.IDTokenLifetimeInSeconds)
	setInt64Value(plan.RefreshTokenLifetimeInSeconds, commonConfigs.RefreshTokenLifetimeInSeconds, &app.RefreshTokenLifetimeInSeconds)
	setStringValue(plan.TemplateGroupID, commonConfigs.TemplateGroupID, &app.TemplateGroupID)
	setStringValue(plan.ClientID, types.StringNull(), &app.ClientID)
	setStringValue(plan.ClientSecret, types.StringNull(), &app.ClientSecret)
	setStringValue(plan.PolicyURI, types.StringNull(), &app.PolicyURI)
	setStringValue(plan.TosURI, types.StringNull(), &app.TosURI)
	setStringValue(plan.ImprintURI, types.StringNull(), &app.ImprintURI)
	setStringValue(plan.TokenEndpointAuthMethod, types.StringNull(), &app.TokenEndpointAuthMethod)
	setStringValue(plan.TokenEndpointAuthSigningAlg, types.StringNull(), &app.TokenEndpointAuthSigningAlg)
	setBoolValue(plan.Editable, commonConfigs.Editable, &app.Editable)
	setStringValue(plan.AppOwner, types.StringNull(), &app.AppOwner)
	setBoolValue(plan.JweEnabled, types.BoolNull(), &app.JweEnabled)
	setBoolValue(plan.UserConsent, types.BoolNull(), &app.UserConsent)
	setBoolValue(plan.Deleted, types.BoolNull(), &app.Deleted)
	setBoolValue(plan.Enabled, commonConfigs.Enabled, &app.Enabled)
	setBoolValue(plan.AlwaysAskMfa, commonConfigs.AlwaysAskMfa, &app.AlwaysAskMfa)
	setBoolValue(plan.SmartMfa, types.BoolNull(), &app.SmartMfa)
	setStringValue(plan.CaptchaRef, types.StringNull(), &app.CaptchaRef)
	setStringValue(plan.CommunicationMediumVerification, types.StringNull(), &app.CommunicationMediumVerification)
	setBoolValue(plan.EmailVerificationRequired, commonConfigs.EmailVerificationRequired, &app.EmailVerificationRequired)
	setBoolValue(plan.MobileNumberVerificationRequired, types.BoolNull(), &app.MobileNumberVerificationRequired)
	setBoolValue(plan.EnableClassicalProvider, commonConfigs.EnableClassicalProvider, &app.EnableClassicalProvider)
	setBoolValue(plan.IsRememberMeSelected, commonConfigs.IsRememberMeSelected, &app.IsRememberMeSelected)
	setBoolValue(plan.EnableBotDetection, types.BoolNull(), &app.EnableBotDetection)
	setStringValue(plan.BotProvider, commonConfigs.BotProvider, &app.BotProvider)
	setBoolValue(plan.IsLoginSuccessPageEnabled, types.BoolNull(), &app.IsLoginSuccessPageEnabled)
	setBoolValue(plan.IsRegisterSuccessPageEnabled, types.BoolNull(), &app.IsRegisterSuccessPageEnabled)
	setBoolValue(plan.AdminClient, types.BoolNull(), &app.AdminClient)
	setBoolValue(plan.IsGroupLoginSelectionEnabled, types.BoolNull(), &app.IsGroupLoginSelectionEnabled)
	setStringValue(plan.BackchannelLogoutURI, types.StringNull(), &app.BackchannelLogoutURI)
	setStringValue(plan.LogoAlign, commonConfigs.LogoAlign, &app.LogoAlign)
	setStringValue(plan.Webfinger, commonConfigs.Webfinger, &app.Webfinger)
	setStringValue(plan.LogoURI, types.StringNull(), &app.LogoURI)
	setStringValue(plan.InitiateLoginURI, types.StringNull(), &app.InitiateLoginURI)
	setStringValue(plan.RegistrationClientURI, types.StringNull(), &app.RegistrationClientURI)
	setStringValue(plan.RegistrationAccessToken, types.StringNull(), &app.RegistrationAccessToken)
	setStringValue(plan.ClientURI, types.StringNull(), &app.ClientURI)
	setStringValue(plan.JwksURI, types.StringNull(), &app.JwksURI)
	setStringValue(plan.Jwks, types.StringNull(), &app.Jwks)
	setStringValue(plan.SectorIdentifierURI, types.StringNull(), &app.SectorIdentifierURI)
	setStringValue(plan.SubjectType, types.StringNull(), &app.SubjectType)
	setStringValue(plan.IDTokenSignedResponseAlg, types.StringNull(), &app.IDTokenSignedResponseAlg)
	setStringValue(plan.IDTokenEncryptedResponseAlg, types.StringNull(), &app.IDTokenEncryptedResponseAlg)
	setStringValue(plan.IDTokenEncryptedResponseEnc, types.StringNull(), &app.IDTokenEncryptedResponseEnc)
	setStringValue(plan.UserinfoSignedResponseAlg, types.StringNull(), &app.UserinfoSignedResponseAlg)
	setStringValue(plan.UserinfoEncryptedResponseAlg, types.StringNull(), &app.UserinfoEncryptedResponseAlg)
	setStringValue(plan.UserinfoEncryptedResponseEnc, types.StringNull(), &app.UserinfoEncryptedResponseEnc)
	setStringValue(plan.RequestObjectSigningAlg, types.StringNull(), &app.RequestObjectSigningAlg)
	setStringValue(plan.RequestObjectEncryptionAlg, types.StringNull(), &app.RequestObjectEncryptionAlg)
	setStringValue(plan.RequestObjectEncryptionEnc, types.StringNull(), &app.RequestObjectEncryptionEnc)
	setStringValue(plan.Description, types.StringNull(), &app.Description)
	setStringValue(plan.ConsentPageGroup, types.StringNull(), &app.ConsentPageGroup)
	setStringValue(plan.PasswordPolicyRef, types.StringNull(), &app.PasswordPolicyRef)
	setStringValue(plan.BlockingMechanismRef, types.StringNull(), &app.BlockingMechanismRef)
	setStringValue(plan.Sub, types.StringNull(), &app.Sub)
	setStringValue(plan.Role, types.StringNull(), &app.Role)
	setStringValue(plan.MfaConfiguration, types.StringNull(), &app.MfaConfiguration)
	setBoolValue(plan.AllowGuestLogin, types.BoolNull(), &app.AllowGuestLogin)
	setStringValue(plan.BackgroundURI, types.StringNull(), &app.BackgroundURI)
	setStringValue(plan.VideoURL, types.StringNull(), &app.VideoURL)
	setStringValue(plan.BotCaptchaRef, types.StringNull(), &app.BotCaptchaRef)
	setStringValue(plan.CreatedAt, types.StringNull(), &app.CreatedTime)
	setStringValue(plan.UpdatedAt, types.StringNull(), &app.UpdatedTime)

	var diags diag.Diagnostics

	diags.Append(setSetValues(ctx, plan.AllowedScopes, commonConfigs.AllowedScopes, &app.AllowedScopes)...)
	diags.Append(setSetValues(ctx, plan.RedirectURIS, commonConfigs.RedirectUris, &app.RedirectURIS)...)
	diags.Append(setSetValues(ctx, plan.AllowedLogoutUrls, commonConfigs.AllowedLogoutUrls, &app.AllowedLogoutUrls)...)
	diags.Append(setSetValues(ctx, plan.AllowedWebOrigins, commonConfigs.AllowedWebOrigins, &app.AllowedWebOrigins)...)
	diags.Append(setSetValues(ctx, plan.AllowedOrigins, commonConfigs.AllowedOrigins, &app.AllowedOrigins)...)
	diags.Append(setSetValues(ctx, plan.LoginProviders, commonConfigs.LoginProviders, &app.LoginProviders)...)
	diags.Append(setSetValues(ctx, plan.DefaultScopes, commonConfigs.DefaultScopes, &app.DefaultScopes)...)
	diags.Append(setSetValues(ctx, plan.PendingScopes, commonConfigs.PendingScopes, &app.PendingScopes)...)
	diags.Append(setSetValues(ctx, plan.AllowedMfa, commonConfigs.AllowedMfa, &app.AllowedMfa)...)
	diags.Append(setSetValues(ctx, plan.AllowedRoles, commonConfigs.AllowedRoles, &app.AllowedRoles)...)
	diags.Append(setSetValues(ctx, plan.DefaultRoles, commonConfigs.DefaultRoles, &app.DefaultRoles)...)
	diags.Append(setSetValues(ctx, plan.PostLogoutRedirectUris, types.SetNull(types.StringType), &app.PostLogoutRedirectUris)...)
	diags.Append(setSetValues(ctx, plan.GroupIDs, types.SetNull(types.StringType), &app.GroupIDs)...)
	diags.Append(setSetValues(ctx, plan.AllowedFields, types.SetNull(types.StringType), &app.AllowedFields)...)
	diags.Append(setSetValues(ctx, plan.WebMessageUris, types.SetNull(types.StringType), &app.WebMessageUris)...)
	diags.Append(setSetValues(ctx, plan.DefaultAcrValues, types.SetNull(types.StringType), &app.DefaultAcrValues)...)
	diags.Append(setSetValues(ctx, plan.GroupTypes, types.SetNull(types.StringType), &app.GroupTypes)...)
	diags.Append(setSetValues(ctx, plan.SuggestMfa, types.SetNull(types.StringType), &app.SuggestMfa)...)
	diags.Append(setSetValues(ctx, plan.CaptchaRefs, types.SetNull(types.StringType), &app.CaptchaRefs)...)
	diags.Append(setSetValues(ctx, plan.ConsentRefs, types.SetNull(types.StringType), &app.ConsentRefs)...)
	diags.Append(setSetValues(ctx, plan.AllowLoginWith, types.SetNull(types.StringType), &app.AllowLoginWith)...)
	diags.Append(setSetValues(ctx, plan.AllowedLogoutUrls, types.SetNull(types.StringType), &app.AllowedLogoutUrls)...)
	diags.Append(setSetValues(ctx, plan.ResponseTypes, types.SetNull(types.StringType), &app.ResponseTypes)...)
	diags.Append(setSetValues(ctx, plan.GrantTypes, types.SetNull(types.StringType), &app.GrantTypes)...)
	diags.Append(setSetValues(ctx, plan.AdditionalAccessTokenPayload, types.SetNull(types.StringType), &app.AdditionalAccessTokenPayload)...)
	diags.Append(setSetValues(ctx, plan.RequiredFields, types.SetNull(types.StringType), &app.RequiredFields)...)
	diags.Append(setSetValues(ctx, plan.Contacts, types.SetNull(types.StringType), &app.Contacts)...)
	diags.Append(setSetValues(ctx, plan.RequestUris, types.SetNull(types.StringType), &app.RequestUris)...)

	setSocialProviders := func(source []*SocialProviderData, target *[]*cidaas.ISocialProviderData) {
		for _, sp := range source {
			*target = append(*target, &cidaas.ISocialProviderData{
				ProviderName: sp.ProviderName.ValueString(),
				SocialID:     sp.SocialID.ValueString(),
			})
		}
	}

	if len(plan.socialProviders) > 0 {
		setSocialProviders(plan.socialProviders, &app.SocialProviders)
	} else if len(commonConfigs.SocialProviders.Elements()) > 0 {
		sps := make([]*SocialProviderData, 0, len(commonConfigs.SocialProviders.Elements()))
		diags.Append(commonConfigs.SocialProviders.ElementsAs(ctx, &sps, false)...)
		setSocialProviders(sps, &app.SocialProviders)
	}

	setProviders := func(ctx context.Context, source []*ProviderMetadData, target *[]*cidaas.IProviderMetadData) diag.Diagnostics {
		for _, cp := range source {
			temp := &cidaas.IProviderMetadData{
				ProviderName:      cp.ProviderName.ValueString(),
				Type:              cp.Type.ValueString(),
				DisplayName:       cp.DisplayName.ValueString(),
				LogoURL:           cp.LogoURL.ValueString(),
				IsProviderVisible: cp.IsProviderVisible.ValueBool(),
			}
			diags := cp.Domains.ElementsAs(ctx, &temp.Domains, false)
			if diags.HasError() {
				return diags
			}
			*target = append(*target, temp)
		}
		return nil
	}

	if len(plan.customProviders) > 0 {
		diags.Append(setProviders(ctx, plan.customProviders, &app.CustomProviders)...)
		if diags.HasError() {
			return nil, diags
		}
	} else if len(commonConfigs.CustomProviders.Elements()) > 0 {
		pmd := make([]*ProviderMetadData, 0, len(commonConfigs.CustomProviders.Elements()))
		diags.Append(commonConfigs.CustomProviders.ElementsAs(ctx, &pmd, false)...)
		diags.Append(setProviders(ctx, pmd, &app.CustomProviders)...)
	}

	if len(plan.samlProviders) > 0 {
		diags.Append(setProviders(ctx, plan.samlProviders, &app.SamlProviders)...)
	} else if len(commonConfigs.SamlProviders.Elements()) > 0 {
		pmd := make([]*ProviderMetadData, 0, len(commonConfigs.SamlProviders.Elements()))
		diags.Append(commonConfigs.SamlProviders.ElementsAs(ctx, &pmd, false)...)
		diags.Append(setProviders(ctx, pmd, &app.SamlProviders)...)
	}

	if len(plan.adProviders) > 0 {
		diags.Append(setProviders(ctx, plan.adProviders, &app.AdProviders)...)
	} else if len(commonConfigs.AdProviders.Elements()) > 0 {
		pmd := make([]*ProviderMetadData, 0, len(commonConfigs.AdProviders.Elements()))
		diags.Append(commonConfigs.AdProviders.ElementsAs(ctx, &pmd, false)...)
		diags.Append(setProviders(ctx, pmd, &app.AdProviders)...)
	}

	for _, ag := range plan.allowedGroups {
		temp := &cidaas.IAllowedGroups{
			GroupID: ag.GroupID.ValueString(),
		}
		diags := ag.Roles.ElementsAs(ctx, temp.Roles, false)
		if diags.HasError() {
			return nil, diags
		}
		diags = ag.DefaultRoles.ElementsAs(ctx, temp.DefaultRoles, false)
		if diags.HasError() {
			return nil, diags
		}
		app.AllowedGroups = append(app.AllowedGroups, temp)
	}

	for _, oag := range plan.operationsAllowedGroups {
		temp := &cidaas.IAllowedGroups{
			GroupID: oag.GroupID.ValueString(),
		}
		diags := oag.Roles.ElementsAs(ctx, temp.Roles, false)
		if diags.HasError() {
			return nil, diags
		}
		diags = oag.DefaultRoles.ElementsAs(ctx, temp.DefaultRoles, false)
		if diags.HasError() {
			return nil, diags
		}
		app.OperationsAllowedGroups = append(app.OperationsAllowedGroups, temp)
	}

	for _, aglg := range plan.allowGuestLoginGroups {
		temp := &cidaas.IAllowedGroups{
			GroupID: aglg.GroupID.ValueString(),
		}
		diags := aglg.Roles.ElementsAs(ctx, temp.Roles, false)
		if diags.HasError() {
			return nil, diags
		}
		diags = aglg.DefaultRoles.ElementsAs(ctx, temp.DefaultRoles, false)
		if diags.HasError() {
			return nil, diags
		}
		app.AllowGuestLoginGroups = append(app.AllowGuestLoginGroups, temp)
	}

	if !plan.LoginSpi.IsNull() {
		app.LoginSpi = &cidaas.ILoginSPI{
			OauthClientID: plan.loginSpi.OauthClientID.ValueString(),
			SpiURL:        plan.loginSpi.SpiURL.ValueString(),
		}
	}

	if !plan.GroupSelection.IsNull() {
		app.GroupSelection = &cidaas.IGroupSelection{
			AlwaysShowGroupSelection: plan.groupSelection.AlwaysShowGroupSelection.ValueBool(),
		}
		diags := plan.groupSelection.SelectableGroups.ElementsAs(ctx, &app.GroupSelection.SelectableGroups, false)
		if diags.HasError() {
			return nil, diags
		}
		diags = plan.groupSelection.SelectableGroups.ElementsAs(ctx, &app.GroupSelection.SelectableGroupTypes, false)
		if diags.HasError() {
			return nil, diags
		}
	}

	if !plan.Mfa.IsNull() {
		app.Mfa = &cidaas.IMfaOption{
			Setting:               plan.mfa.Setting.ValueString(),
			TimeIntervalInSeconds: plan.mfa.TimeIntervalInSeconds.ValueInt64Pointer(),
		}
		diags := plan.mfa.AllowedMethods.ElementsAs(ctx, &app.Mfa.AllowedMethods, false)
		if diags.HasError() {
			return nil, diags
		}
	}

	if !plan.MobileSettings.IsNull() {
		app.MobileSettings = &cidaas.IAppMobileSettings{
			TeamID:      plan.mobileSettings.TeamID.ValueString(),
			BundleID:    plan.mobileSettings.BundleID.ValueString(),
			PackageName: plan.mobileSettings.PackageName.ValueString(),
			KeyHash:     plan.mobileSettings.KeyHash.ValueString(),
		}
	}
	return &app, diags
}

func updateStateModel(ctx context.Context, res *cidaas.AppResponse, state, config *AppConfig) resource.ReadResponse {
	var d diag.Diagnostics
	resp := resource.ReadResponse{}

	// only applicable for the required varaibles
	if config.CommonConfigs.IsNull() {
		state.CompanyName = util.StringValueOrNull(&res.Data.CompanyName)
		state.CompanyWebsite = util.StringValueOrNull(&res.Data.CompanyWebsite)
		state.ClientType = util.StringValueOrNull(&res.Data.ClientType)
		state.CompanyAddress = util.StringValueOrNull(&res.Data.CompanyAddress)

		state.AllowedScopes = util.SetValueOrNull(res.Data.AllowedScopes)
		state.RedirectURIS = util.SetValueOrNull(res.Data.RedirectURIS)
		state.AllowedLogoutUrls = util.SetValueOrNull(res.Data.AllowedLogoutUrls)
		state.AllowedWebOrigins = util.SetValueOrNull(res.Data.AllowedWebOrigins)
		state.AllowedOrigins = util.SetValueOrNull(res.Data.AllowedOrigins)
		state.LoginProviders = util.SetValueOrNull(res.Data.LoginProviders)
		state.DefaultScopes = util.SetValueOrNull(res.Data.DefaultScopes)
		state.PendingScopes = util.SetValueOrNull(res.Data.PendingScopes)
		state.AllowedMfa = util.SetValueOrNull(res.Data.AllowedMfa)
		state.AllowedRoles = util.SetValueOrNull(res.Data.AllowedRoles)
		state.DefaultRoles = util.SetValueOrNull(res.Data.DefaultRoles)

		state.AccentColor = util.StringValueOrNull(&res.Data.AccentColor)
		state.PrimaryColor = util.StringValueOrNull(&res.Data.PrimaryColor)
		state.MediaType = util.StringValueOrNull(&res.Data.MediaType)
		state.HostedPageGroup = util.StringValueOrNull(&res.Data.HostedPageGroup)
		state.TemplateGroupID = util.StringValueOrNull(&res.Data.TemplateGroupID)
		state.BotProvider = util.StringValueOrNull(&res.Data.BotProvider)
		state.LogoAlign = util.StringValueOrNull(&res.Data.LogoAlign)
		state.Webfinger = util.StringValueOrNull(&res.Data.Webfinger)
		state.DefaultMaxAge = types.Int64Value(res.Data.DefaultMaxAge)
		state.TokenLifetimeInSeconds = types.Int64Value(res.Data.TokenLifetimeInSeconds)
		state.IDTokenLifetimeInSeconds = types.Int64Value(res.Data.IDTokenLifetimeInSeconds)
		state.RefreshTokenLifetimeInSeconds = types.Int64Value(res.Data.RefreshTokenLifetimeInSeconds)
		state.AllowGuestLogin = types.BoolValue(res.Data.AllowGuestLogin)
		state.EnableDeduplication = types.BoolValue(res.Data.EnableDeduplication)
		state.AutoLoginAfterRegister = types.BoolValue(res.Data.AutoLoginAfterRegister)
		state.EnablePasswordlessAuth = types.BoolValue(res.Data.EnablePasswordlessAuth)
		state.RegisterWithLoginInformation = types.BoolValue(res.Data.RegisterWithLoginInformation)
		state.FdsEnabled = types.BoolValue(res.Data.FdsEnabled)
		state.IsHybridApp = types.BoolValue(res.Data.IsHybridApp)
		state.Editable = types.BoolValue(res.Data.Editable)
		state.Enabled = types.BoolValue(res.Data.Enabled)
		state.AlwaysAskMfa = types.BoolValue(res.Data.AlwaysAskMfa)
		state.EmailVerificationRequired = types.BoolValue(res.Data.EmailVerificationRequired)
		state.EnableClassicalProvider = types.BoolValue(res.Data.EnableClassicalProvider)
		state.IsRememberMeSelected = types.BoolValue(res.Data.IsRememberMeSelected)
		resp.Diagnostics.Append(extractSetValues(ctx, &state.ResponseTypes, res.Data.ResponseTypes)...)
		resp.Diagnostics.Append(extractSetValues(ctx, &state.GrantTypes, res.Data.GrantTypes)...)
		resp.Diagnostics.Append(extractSetValues(ctx, &state.AllowLoginWith, res.Data.AllowLoginWith)...)

		if res.Data.Mfa != nil {
			mfa, diags := types.ObjectValue(
				map[string]attr.Type{
					"setting":                  types.StringType,
					"time_interval_in_seconds": types.Int64Type,
					"allowed_methods":          types.SetType{ElemType: types.StringType},
				},
				map[string]attr.Value{
					"setting":                  util.StringValueOrNull(&res.Data.Mfa.Setting),
					"time_interval_in_seconds": util.Int64ValueOrNull(res.Data.Mfa.TimeIntervalInSeconds),
					"allowed_methods":          util.SetValueOrNull(res.Data.Mfa.AllowedMethods),
				})
			resp.Diagnostics.Append(diags...)
			if resp.Diagnostics.HasError() {
				return resp
			}
			state.Mfa = mfa
		}

		var spObjectValues []attr.Value
		spObjectType := types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"provider_name": types.StringType,
				"social_id":     types.StringType,
			},
		}
		for _, sp := range res.Data.SocialProviders {
			objValue := types.ObjectValueMust(
				spObjectType.AttrTypes,
				map[string]attr.Value{
					"provider_name": util.StringValueOrNull(&sp.ProviderName),
					"social_id":     util.StringValueOrNull(&sp.SocialID),
				})
			spObjectValues = append(spObjectValues, objValue)
		}

		state.SocialProviders, d = types.ListValueFrom(ctx, spObjectType, spObjectValues)
		resp.Diagnostics.Append(d...)
		if resp.Diagnostics.HasError() {
			return resp
		}

		providerObjectType := types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"provider_name":       types.StringType,
				"display_name":        types.StringType,
				"logo_url":            types.StringType,
				"type":                types.StringType,
				"is_provider_visible": types.BoolType,
				"domains":             types.SetType{ElemType: types.StringType},
			},
		}

		var customProviderMetaObjectValues []attr.Value
		for _, cp := range res.Data.CustomProviders {
			objValue := types.ObjectValueMust(
				providerObjectType.AttrTypes,
				map[string]attr.Value{
					"provider_name":       util.StringValueOrNull(&cp.ProviderName),
					"display_name":        util.StringValueOrNull(&cp.DisplayName),
					"logo_url":            util.StringValueOrNull(&cp.LogoURL),
					"type":                util.StringValueOrNull(&cp.Type),
					"is_provider_visible": types.BoolValue(cp.IsProviderVisible),
					"domains": func() basetypes.SetValue {
						if len(cp.Domains) == 0 {
							return types.SetNull(types.StringType)
						}
						return types.SetValueMust(types.StringType, func() []attr.Value {
							var temp []attr.Value
							for _, role := range cp.Domains {
								temp = append(temp, util.StringValueOrNull(&role))
							}
							return temp
						}())
					}(),
				})
			customProviderMetaObjectValues = append(customProviderMetaObjectValues, objValue)
		}
		state.CustomProviders, d = types.ListValueFrom(ctx, providerObjectType, customProviderMetaObjectValues)
		resp.Diagnostics.Append(d...)
		if resp.Diagnostics.HasError() {
			return resp
		}

		var samlProviderMetaObjectValues []attr.Value
		for _, saml := range res.Data.SamlProviders {
			objValue := types.ObjectValueMust(
				providerObjectType.AttrTypes,
				map[string]attr.Value{
					"provider_name":       util.StringValueOrNull(&saml.ProviderName),
					"display_name":        util.StringValueOrNull(&saml.DisplayName),
					"logo_url":            util.StringValueOrNull(&saml.LogoURL),
					"type":                util.StringValueOrNull(&saml.Type),
					"is_provider_visible": types.BoolValue(saml.IsProviderVisible),
					"domains": func() basetypes.SetValue {
						if len(saml.Domains) == 0 {
							return types.SetNull(types.StringType)
						}
						return types.SetValueMust(types.StringType, func() []attr.Value {
							var temp []attr.Value
							for _, role := range saml.Domains {
								temp = append(temp, util.StringValueOrNull(&role))
							}
							return temp
						}())
					}(),
				})
			samlProviderMetaObjectValues = append(samlProviderMetaObjectValues, objValue)
		}

		state.SamlProviders, d = types.ListValueFrom(ctx, providerObjectType, samlProviderMetaObjectValues)
		resp.Diagnostics.Append(d...)
		if resp.Diagnostics.HasError() {
			return resp
		}

		var adProviderMetaObjectValues []attr.Value
		for _, ad := range res.Data.AdProviders {
			objValue := types.ObjectValueMust(
				providerObjectType.AttrTypes,
				map[string]attr.Value{
					"provider_name":       util.StringValueOrNull(&ad.ProviderName),
					"display_name":        util.StringValueOrNull(&ad.DisplayName),
					"logo_url":            util.StringValueOrNull(&ad.LogoURL),
					"type":                util.StringValueOrNull(&ad.Type),
					"is_provider_visible": types.BoolValue(ad.IsProviderVisible),
					"domains": func() basetypes.SetValue {
						if len(ad.Domains) == 0 {
							return types.SetNull(types.StringType)
						}
						return types.SetValueMust(types.StringType, func() []attr.Value {
							var temp []attr.Value
							for _, role := range ad.Domains {
								temp = append(temp, util.StringValueOrNull(&role))
							}
							return temp
						}())
					}(),
				})
			adProviderMetaObjectValues = append(adProviderMetaObjectValues, objValue)
		}

		state.AdProviders, d = types.ListValueFrom(ctx, providerObjectType, adProviderMetaObjectValues)
		resp.Diagnostics.Append(d...)
		if resp.Diagnostics.HasError() {
			return resp
		}

		allowedGroupsObjectType := types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"group_id":      types.StringType,
				"roles":         types.SetType{ElemType: types.StringType},
				"default_roles": types.SetType{ElemType: types.StringType},
			},
		}

		var allowedGroupsObjectValues []attr.Value
		for _, ag := range res.Data.AllowedGroups {
			objValue := types.ObjectValueMust(
				allowedGroupsObjectType.AttrTypes,
				map[string]attr.Value{
					"group_id": util.StringValueOrNull(&ag.GroupID),
					"roles": types.SetValueMust(types.StringType, func() []attr.Value {
						var temp []attr.Value
						for _, role := range ag.Roles {
							temp = append(temp, util.StringValueOrNull(&role))
						}
						return temp
					}()),
					"default_roles": types.SetValueMust(types.StringType, func() []attr.Value {
						var temp []attr.Value
						for _, role := range ag.DefaultRoles {
							temp = append(temp, util.StringValueOrNull(&role))
						}
						return temp
					}()),
				})
			allowedGroupsObjectValues = append(allowedGroupsObjectValues, objValue)
		}

		state.AllowedGroups, d = types.ListValueFrom(ctx, allowedGroupsObjectType, allowedGroupsObjectValues)
		resp.Diagnostics.Append(d...)
		if resp.Diagnostics.HasError() {
			return resp
		}

		var opsAllowedGroupsObjectValues []attr.Value
		for _, oag := range res.Data.OperationsAllowedGroups {
			objValue := types.ObjectValueMust(
				allowedGroupsObjectType.AttrTypes,
				map[string]attr.Value{
					"group_id": util.StringValueOrNull(&oag.GroupID),
					"roles": types.SetValueMust(types.StringType, func() []attr.Value {
						var temp []attr.Value
						for _, role := range oag.Roles {
							temp = append(temp, util.StringValueOrNull(&role))
						}
						return temp
					}()),
					"default_roles": types.SetValueMust(types.StringType, func() []attr.Value {
						var temp []attr.Value
						for _, role := range oag.DefaultRoles {
							temp = append(temp, util.StringValueOrNull(&role))
						}
						return temp
					}()),
				})
			opsAllowedGroupsObjectValues = append(opsAllowedGroupsObjectValues, objValue)
		}

		state.OperationsAllowedGroups, d = types.ListValueFrom(ctx, allowedGroupsObjectType, opsAllowedGroupsObjectValues)
		resp.Diagnostics.Append(d...)
		if resp.Diagnostics.HasError() {
			return resp
		}
	} else {
		if !config.commonConfigs.CompanyName.IsNull() {
			state.commonConfigs.CompanyName = util.StringValueOrNull(&res.Data.CompanyName)
		}
		if !config.CompanyName.IsNull() {
			state.CompanyName = util.StringValueOrNull(&res.Data.CompanyName)
		}
		if !config.commonConfigs.CompanyWebsite.IsNull() {
			state.commonConfigs.CompanyWebsite = util.StringValueOrNull(&res.Data.CompanyWebsite)
		}
		if !config.CompanyWebsite.IsNull() {
			state.CompanyWebsite = util.StringValueOrNull(&res.Data.CompanyWebsite)
		}

		if !config.commonConfigs.ClientType.IsNull() {
			state.commonConfigs.ClientType = util.StringValueOrNull(&res.Data.ClientType)
		}
		if !config.ClientType.IsNull() {
			state.ClientType = util.StringValueOrNull(&res.Data.ClientType)
		}

		if !config.commonConfigs.CompanyAddress.IsNull() {
			state.commonConfigs.CompanyAddress = util.StringValueOrNull(&res.Data.CompanyAddress)
		}
		if !config.CompanyAddress.IsNull() {
			state.CompanyAddress = util.StringValueOrNull(&res.Data.CompanyAddress)
		}
		if !config.commonConfigs.AllowedScopes.IsNull() {
			state.commonConfigs.AllowedScopes = util.SetValueOrNull(res.Data.AllowedScopes)
		}
		if !config.AllowedScopes.IsNull() {
			state.AllowedScopes = util.SetValueOrNull(res.Data.AllowedScopes)
		}
		if !config.commonConfigs.RedirectUris.IsNull() {
			state.commonConfigs.RedirectUris = util.SetValueOrNull(res.Data.RedirectURIS)
		}
		if !config.RedirectURIS.IsNull() {
			state.RedirectURIS = util.SetValueOrNull(res.Data.RedirectURIS)
		}
		if !config.commonConfigs.AllowedLogoutUrls.IsNull() {
			state.commonConfigs.AllowedLogoutUrls = util.SetValueOrNull(res.Data.AllowedLogoutUrls)
		}
		if !config.AllowedLogoutUrls.IsNull() {
			state.AllowedLogoutUrls = util.SetValueOrNull(res.Data.AllowedLogoutUrls)
		}
		if !config.commonConfigs.AllowedWebOrigins.IsNull() {
			state.commonConfigs.AllowedWebOrigins = util.SetValueOrNull(res.Data.AllowedWebOrigins)
		}
		if !config.AllowedWebOrigins.IsNull() {
			state.AllowedWebOrigins = util.SetValueOrNull(res.Data.AllowedWebOrigins)
		}
		if !config.commonConfigs.AllowedOrigins.IsNull() {
			state.commonConfigs.AllowedOrigins = util.SetValueOrNull(res.Data.AllowedOrigins)
		}
		if !config.AllowedOrigins.IsNull() {
			state.AllowedOrigins = util.SetValueOrNull(res.Data.AllowedOrigins)
		}
		if !config.commonConfigs.LoginProviders.IsNull() {
			state.commonConfigs.LoginProviders = util.SetValueOrNull(res.Data.LoginProviders)
		}
		if !config.LoginProviders.IsNull() {
			state.LoginProviders = util.SetValueOrNull(res.Data.LoginProviders)
		}
		if !config.commonConfigs.DefaultScopes.IsNull() {
			state.commonConfigs.DefaultScopes = util.SetValueOrNull(res.Data.DefaultScopes)
		}
		if !config.DefaultScopes.IsNull() {
			state.DefaultScopes = util.SetValueOrNull(res.Data.DefaultScopes)
		}

		if !config.commonConfigs.PendingScopes.IsNull() {
			state.commonConfigs.PendingScopes = util.SetValueOrNull(res.Data.PendingScopes)
		}
		if !config.PendingScopes.IsNull() {
			state.PendingScopes = util.SetValueOrNull(res.Data.PendingScopes)
		}
		if !config.commonConfigs.AllowedMfa.IsNull() {
			state.commonConfigs.AllowedMfa = util.SetValueOrNull(res.Data.AllowedMfa)
		}
		if !config.AllowedMfa.IsNull() {
			state.AllowedMfa = util.SetValueOrNull(res.Data.AllowedMfa)
		}
		if !config.commonConfigs.AllowedRoles.IsNull() {
			state.commonConfigs.AllowedRoles = util.SetValueOrNull(res.Data.AllowedRoles)
		}
		if !config.AllowedRoles.IsNull() {
			state.AllowedRoles = util.SetValueOrNull(res.Data.AllowedRoles)
		}

		if !config.commonConfigs.DefaultRoles.IsNull() {
			state.commonConfigs.DefaultRoles = util.SetValueOrNull(res.Data.DefaultRoles)
		}
		if !config.DefaultRoles.IsNull() {
			state.DefaultRoles = util.SetValueOrNull(res.Data.DefaultRoles)
		}
		if !config.commonConfigs.AccentColor.IsNull() {
			state.commonConfigs.AccentColor = util.StringValueOrNull(&res.Data.AccentColor)
		}
		if !config.AccentColor.IsNull() {
			state.AccentColor = util.StringValueOrNull(&res.Data.AccentColor)
		}

		if !config.commonConfigs.PrimaryColor.IsNull() {
			state.commonConfigs.PrimaryColor = util.StringValueOrNull(&res.Data.PrimaryColor)
		}
		if !config.PrimaryColor.IsNull() {
			state.PrimaryColor = util.StringValueOrNull(&res.Data.PrimaryColor)
		}

		if !config.commonConfigs.MediaType.IsNull() {
			state.commonConfigs.MediaType = util.StringValueOrNull(&res.Data.MediaType)
		}
		if !config.MediaType.IsNull() {
			state.MediaType = util.StringValueOrNull(&res.Data.MediaType)
		}

		if !config.commonConfigs.HostedPageGroup.IsNull() {
			state.commonConfigs.HostedPageGroup = util.StringValueOrNull(&res.Data.HostedPageGroup)
		}
		if !config.HostedPageGroup.IsNull() {
			state.HostedPageGroup = util.StringValueOrNull(&res.Data.HostedPageGroup)
		}

		if !config.commonConfigs.TemplateGroupID.IsNull() {
			state.commonConfigs.TemplateGroupID = util.StringValueOrNull(&res.Data.TemplateGroupID)
		}
		if !config.TemplateGroupID.IsNull() {
			state.TemplateGroupID = util.StringValueOrNull(&res.Data.TemplateGroupID)
		}

		if !config.commonConfigs.BotProvider.IsNull() {
			state.commonConfigs.BotProvider = util.StringValueOrNull(&res.Data.BotProvider)
		}
		if !config.BotProvider.IsNull() {
			state.BotProvider = util.StringValueOrNull(&res.Data.BotProvider)
		}

		if !config.commonConfigs.LogoAlign.IsNull() {
			state.commonConfigs.LogoAlign = util.StringValueOrNull(&res.Data.LogoAlign)
		}
		if !config.LogoAlign.IsNull() {
			state.LogoAlign = util.StringValueOrNull(&res.Data.LogoAlign)
		}

		if !config.commonConfigs.Webfinger.IsNull() {
			state.commonConfigs.Webfinger = util.StringValueOrNull(&res.Data.Webfinger)
		}
		if !config.Webfinger.IsNull() {
			state.Webfinger = util.StringValueOrNull(&res.Data.Webfinger)
		}

		if !config.commonConfigs.DefaultMaxAge.IsNull() {
			state.commonConfigs.DefaultMaxAge = types.Int64Value(res.Data.DefaultMaxAge)
		}
		if !config.DefaultMaxAge.IsNull() {
			state.DefaultMaxAge = types.Int64Value(res.Data.DefaultMaxAge)
		}

		if !config.commonConfigs.TokenLifetimeInSeconds.IsNull() {
			state.commonConfigs.TokenLifetimeInSeconds = types.Int64Value(res.Data.TokenLifetimeInSeconds)
		}
		if !config.TokenLifetimeInSeconds.IsNull() {
			state.TokenLifetimeInSeconds = types.Int64Value(res.Data.TokenLifetimeInSeconds)
		}

		if !config.commonConfigs.IDTokenLifetimeInSeconds.IsNull() {
			state.commonConfigs.IDTokenLifetimeInSeconds = types.Int64Value(res.Data.IDTokenLifetimeInSeconds)
		}
		if !config.IDTokenLifetimeInSeconds.IsNull() {
			state.IDTokenLifetimeInSeconds = types.Int64Value(res.Data.IDTokenLifetimeInSeconds)
		}

		if !config.commonConfigs.RefreshTokenLifetimeInSeconds.IsNull() {
			state.commonConfigs.RefreshTokenLifetimeInSeconds = types.Int64Value(res.Data.RefreshTokenLifetimeInSeconds)
		}
		if !config.RefreshTokenLifetimeInSeconds.IsNull() {
			state.RefreshTokenLifetimeInSeconds = types.Int64Value(res.Data.RefreshTokenLifetimeInSeconds)
		}

		if !config.commonConfigs.AllowGuestLogin.IsNull() {
			state.commonConfigs.AllowGuestLogin = types.BoolValue(res.Data.AllowGuestLogin)
		}
		if !config.AllowGuestLogin.IsNull() {
			state.AllowGuestLogin = types.BoolValue(res.Data.AllowGuestLogin)
		}

		if !config.commonConfigs.EnableDeduplication.IsNull() {
			state.commonConfigs.EnableDeduplication = types.BoolValue(res.Data.EnableDeduplication)
		}
		if !config.EnableDeduplication.IsNull() {
			state.EnableDeduplication = types.BoolValue(res.Data.EnableDeduplication)
		}

		if !config.commonConfigs.AutoLoginAfterRegister.IsNull() {
			state.commonConfigs.AutoLoginAfterRegister = types.BoolValue(res.Data.AutoLoginAfterRegister)
		}
		if !config.AutoLoginAfterRegister.IsNull() {
			state.AutoLoginAfterRegister = types.BoolValue(res.Data.AutoLoginAfterRegister)
		}

		if !config.commonConfigs.EnablePasswordlessAuth.IsNull() {
			state.commonConfigs.EnablePasswordlessAuth = types.BoolValue(res.Data.EnablePasswordlessAuth)
		}
		if !config.EnablePasswordlessAuth.IsNull() {
			state.EnablePasswordlessAuth = types.BoolValue(res.Data.EnablePasswordlessAuth)
		}

		if !config.commonConfigs.RegisterWithLoginInformation.IsNull() {
			state.commonConfigs.RegisterWithLoginInformation = types.BoolValue(res.Data.RegisterWithLoginInformation)
		}
		if !config.RegisterWithLoginInformation.IsNull() {
			state.RegisterWithLoginInformation = types.BoolValue(res.Data.RegisterWithLoginInformation)
		}

		if !config.commonConfigs.FdsEnabled.IsNull() {
			state.commonConfigs.FdsEnabled = types.BoolValue(res.Data.FdsEnabled)
		}
		if !config.FdsEnabled.IsNull() {
			state.FdsEnabled = types.BoolValue(res.Data.FdsEnabled)
		}

		if !config.commonConfigs.IsHybridApp.IsNull() {
			state.commonConfigs.IsHybridApp = types.BoolValue(res.Data.IsHybridApp)
		}
		if !config.IsHybridApp.IsNull() {
			state.IsHybridApp = types.BoolValue(res.Data.IsHybridApp)
		}

		if !config.commonConfigs.Editable.IsNull() {
			state.commonConfigs.Editable = types.BoolValue(res.Data.Editable)
		}
		if !config.Editable.IsNull() {
			state.Editable = types.BoolValue(res.Data.Editable)
		}

		if !config.commonConfigs.Enabled.IsNull() {
			state.commonConfigs.Enabled = types.BoolValue(res.Data.Enabled)
		}
		if !config.Enabled.IsNull() {
			state.Enabled = types.BoolValue(res.Data.Enabled)
		}

		if !config.commonConfigs.AlwaysAskMfa.IsNull() {
			state.commonConfigs.AlwaysAskMfa = types.BoolValue(res.Data.AlwaysAskMfa)
		}
		if !config.AlwaysAskMfa.IsNull() {
			state.AlwaysAskMfa = types.BoolValue(res.Data.AlwaysAskMfa)
		}

		if !config.commonConfigs.EmailVerificationRequired.IsNull() {
			state.commonConfigs.EmailVerificationRequired = types.BoolValue(res.Data.EmailVerificationRequired)
		}
		if !config.EmailVerificationRequired.IsNull() {
			state.EmailVerificationRequired = types.BoolValue(res.Data.EmailVerificationRequired)
		}

		if !config.commonConfigs.EnableClassicalProvider.IsNull() {
			state.commonConfigs.EnableClassicalProvider = types.BoolValue(res.Data.EnableClassicalProvider)
		}
		if !config.EnableClassicalProvider.IsNull() {
			state.EnableClassicalProvider = types.BoolValue(res.Data.EnableClassicalProvider)
		}

		if !config.commonConfigs.IsRememberMeSelected.IsNull() {
			state.commonConfigs.IsRememberMeSelected = types.BoolValue(res.Data.IsRememberMeSelected)
		}
		if !config.IsRememberMeSelected.IsNull() {
			state.IsRememberMeSelected = types.BoolValue(res.Data.IsRememberMeSelected)
		}

		if !config.commonConfigs.ResponseTypes.IsNull() {
			resp.Diagnostics.Append(extractSetValues(ctx, &state.commonConfigs.ResponseTypes, res.Data.ResponseTypes)...)
		}
		if !config.ResponseTypes.IsNull() {
			resp.Diagnostics.Append(extractSetValues(ctx, &state.ResponseTypes, res.Data.ResponseTypes)...)
		}
		if !config.commonConfigs.GrantTypes.IsNull() {
			resp.Diagnostics.Append(extractSetValues(ctx, &state.commonConfigs.GrantTypes, res.Data.GrantTypes)...)
		}
		if !config.GrantTypes.IsNull() {
			resp.Diagnostics.Append(extractSetValues(ctx, &state.GrantTypes, res.Data.GrantTypes)...)
		}

		if !config.commonConfigs.AllowLoginWith.IsNull() {
			resp.Diagnostics.Append(extractSetValues(ctx, &state.commonConfigs.AllowLoginWith, res.Data.GrantTypes)...)
		}
		if !config.AllowLoginWith.IsNull() {
			resp.Diagnostics.Append(extractSetValues(ctx, &state.AllowLoginWith, res.Data.AllowLoginWith)...)
		}

		if res.Data.Mfa != nil {
			mfa, diags := types.ObjectValue(
				map[string]attr.Type{
					"setting":                  types.StringType,
					"time_interval_in_seconds": types.Int64Type,
					"allowed_methods":          types.SetType{ElemType: types.StringType},
				},
				map[string]attr.Value{
					"setting":                  util.StringValueOrNull(&res.Data.Mfa.Setting),
					"time_interval_in_seconds": util.Int64ValueOrNull(res.Data.Mfa.TimeIntervalInSeconds),
					"allowed_methods":          util.SetValueOrNull(res.Data.Mfa.AllowedMethods),
				})
			resp.Diagnostics.Append(diags...)
			if resp.Diagnostics.HasError() {
				return resp
			}
			if !config.commonConfigs.Mfa.IsNull() {
				state.commonConfigs.Mfa = mfa
			}
			if !config.Mfa.IsNull() {
				state.Mfa = mfa
			}
		}

		var spObjectValues []attr.Value
		spObjectType := types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"provider_name": types.StringType,
				"social_id":     types.StringType,
			},
		}
		for _, sp := range res.Data.SocialProviders {
			objValue := types.ObjectValueMust(
				spObjectType.AttrTypes,
				map[string]attr.Value{
					"provider_name": util.StringValueOrNull(&sp.ProviderName),
					"social_id":     util.StringValueOrNull(&sp.SocialID),
				})
			spObjectValues = append(spObjectValues, objValue)
		}

		socialProviders, d := types.ListValueFrom(ctx, spObjectType, spObjectValues)
		resp.Diagnostics.Append(d...)
		if resp.Diagnostics.HasError() {
			return resp
		}

		if !config.commonConfigs.SocialProviders.IsNull() {
			state.commonConfigs.SocialProviders = socialProviders
		}
		if !config.SocialProviders.IsNull() {
			state.SocialProviders = socialProviders
		}

		providerObjectType := types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"provider_name":       types.StringType,
				"display_name":        types.StringType,
				"logo_url":            types.StringType,
				"type":                types.StringType,
				"is_provider_visible": types.BoolType,
				"domains":             types.SetType{ElemType: types.StringType},
			},
		}

		var customProviderMetaObjectValues []attr.Value
		for _, cp := range res.Data.CustomProviders {
			objValue := types.ObjectValueMust(
				providerObjectType.AttrTypes,
				map[string]attr.Value{
					"provider_name":       util.StringValueOrNull(&cp.ProviderName),
					"display_name":        util.StringValueOrNull(&cp.DisplayName),
					"logo_url":            util.StringValueOrNull(&cp.LogoURL),
					"type":                util.StringValueOrNull(&cp.Type),
					"is_provider_visible": types.BoolValue(cp.IsProviderVisible),
					"domains":             util.SetValueOrNull(cp.Domains),
				})
			customProviderMetaObjectValues = append(customProviderMetaObjectValues, objValue)
		}
		customProviders, d := types.ListValueFrom(ctx, providerObjectType, customProviderMetaObjectValues)
		resp.Diagnostics.Append(d...)
		if resp.Diagnostics.HasError() {
			return resp
		}

		if !config.commonConfigs.CustomProviders.IsNull() {
			state.commonConfigs.CustomProviders = customProviders
		}
		if !config.CustomProviders.IsNull() {
			state.CustomProviders = customProviders
		}

		var samlProviderMetaObjectValues []attr.Value
		for _, saml := range res.Data.SamlProviders {
			objValue := types.ObjectValueMust(
				providerObjectType.AttrTypes,
				map[string]attr.Value{
					"provider_name":       util.StringValueOrNull(&saml.ProviderName),
					"display_name":        util.StringValueOrNull(&saml.DisplayName),
					"logo_url":            util.StringValueOrNull(&saml.LogoURL),
					"type":                util.StringValueOrNull(&saml.Type),
					"is_provider_visible": types.BoolValue(saml.IsProviderVisible),
					"domains":             util.SetValueOrNull(saml.Domains),
				})
			samlProviderMetaObjectValues = append(samlProviderMetaObjectValues, objValue)
		}

		samlProviders, d := types.ListValueFrom(ctx, providerObjectType, samlProviderMetaObjectValues)
		resp.Diagnostics.Append(d...)
		if resp.Diagnostics.HasError() {
			return resp
		}

		if !config.commonConfigs.SamlProviders.IsNull() {
			state.commonConfigs.SamlProviders = samlProviders
		}
		if !config.SamlProviders.IsNull() {
			state.SamlProviders = samlProviders
		}

		var adProviderMetaObjectValues []attr.Value
		for _, ad := range res.Data.AdProviders {
			objValue := types.ObjectValueMust(
				providerObjectType.AttrTypes,
				map[string]attr.Value{
					"provider_name":       util.StringValueOrNull(&ad.ProviderName),
					"display_name":        util.StringValueOrNull(&ad.DisplayName),
					"logo_url":            util.StringValueOrNull(&ad.LogoURL),
					"type":                util.StringValueOrNull(&ad.Type),
					"is_provider_visible": types.BoolValue(ad.IsProviderVisible),
					"domains":             util.SetValueOrNull(ad.Domains),
				})
			adProviderMetaObjectValues = append(adProviderMetaObjectValues, objValue)
		}

		adProviders, d := types.ListValueFrom(ctx, providerObjectType, adProviderMetaObjectValues)
		resp.Diagnostics.Append(d...)
		if resp.Diagnostics.HasError() {
			return resp
		}

		if !config.commonConfigs.AdProviders.IsNull() {
			state.commonConfigs.AdProviders = adProviders
		}
		if !config.AdProviders.IsNull() {
			state.AdProviders = adProviders
		}

		allowedGroupsObjectType := types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"group_id":      types.StringType,
				"roles":         types.SetType{ElemType: types.StringType},
				"default_roles": types.SetType{ElemType: types.StringType},
			},
		}

		var allowedGroupsObjectValues []attr.Value
		for _, ag := range res.Data.AllowedGroups {
			objValue := types.ObjectValueMust(
				allowedGroupsObjectType.AttrTypes,
				map[string]attr.Value{
					"group_id":      util.StringValueOrNull(&ag.GroupID),
					"roles":         util.SetValueOrNull(ag.Roles),
					"default_roles": util.SetValueOrNull(ag.DefaultRoles),
				})
			allowedGroupsObjectValues = append(allowedGroupsObjectValues, objValue)
		}

		allowedGroups, d := types.ListValueFrom(ctx, allowedGroupsObjectType, allowedGroupsObjectValues)
		resp.Diagnostics.Append(d...)
		if resp.Diagnostics.HasError() {
			return resp
		}

		if !config.commonConfigs.AllowedGroups.IsNull() {
			state.commonConfigs.AllowedGroups = allowedGroups
		}
		if !config.AllowedGroups.IsNull() {
			state.AllowedGroups = allowedGroups
		}

		var opsAllowedGroupsObjectValues []attr.Value
		for _, oag := range res.Data.OperationsAllowedGroups {
			objValue := types.ObjectValueMust(
				allowedGroupsObjectType.AttrTypes,
				map[string]attr.Value{
					"group_id":      util.StringValueOrNull(&oag.GroupID),
					"roles":         util.SetValueOrNull(oag.Roles),
					"default_roles": util.SetValueOrNull(oag.DefaultRoles),
				})
			opsAllowedGroupsObjectValues = append(opsAllowedGroupsObjectValues, objValue)
		}

		operationsAllowedGroups, d := types.ListValueFrom(ctx, allowedGroupsObjectType, opsAllowedGroupsObjectValues)
		resp.Diagnostics.Append(d...)
		if resp.Diagnostics.HasError() {
			return resp
		}
		if !config.commonConfigs.OperationsAllowedGroups.IsNull() {
			state.commonConfigs.OperationsAllowedGroups = operationsAllowedGroups
		}
		if !config.OperationsAllowedGroups.IsNull() {
			state.OperationsAllowedGroups = operationsAllowedGroups
		}
	}

	state.ID = util.StringValueOrNull(&res.Data.ID)
	state.ContentAlign = util.StringValueOrNull(&res.Data.ContentAlign)
	state.AllowDisposableEmail = types.BoolValue(res.Data.AllowDisposableEmail)
	state.ValidatePhoneNumber = types.BoolValue(res.Data.ValidatePhoneNumber)
	state.ClientName = util.StringValueOrNull(&res.Data.ClientName)
	state.ClientDisplayName = util.StringValueOrNull(&res.Data.ClientDisplayName)
	state.ClientID = util.StringValueOrNull(&res.Data.ClientID)
	state.ClientSecret = util.StringValueOrNull(&res.Data.ClientSecret)
	state.PolicyURI = util.StringValueOrNull(&res.Data.PolicyURI)
	state.TosURI = util.StringValueOrNull(&res.Data.TosURI)
	state.ImprintURI = util.StringValueOrNull(&res.Data.ImprintURI)
	state.TokenEndpointAuthMethod = util.StringValueOrNull(&res.Data.TokenEndpointAuthMethod)
	state.TokenEndpointAuthSigningAlg = util.StringValueOrNull(&res.Data.TokenEndpointAuthSigningAlg)
	state.AppOwner = util.StringValueOrNull(&res.Data.AppOwner)
	state.JweEnabled = types.BoolValue(res.Data.JweEnabled)
	state.UserConsent = types.BoolValue(res.Data.UserConsent)
	state.Deleted = types.BoolValue(res.Data.Deleted)
	state.SmartMfa = types.BoolValue(res.Data.SmartMfa)
	state.CaptchaRef = util.StringValueOrNull(&res.Data.CaptchaRef)
	state.CommunicationMediumVerification = util.StringValueOrNull(&res.Data.CommunicationMediumVerification)
	state.MobileNumberVerificationRequired = types.BoolValue(res.Data.MobileNumberVerificationRequired)
	state.EnableBotDetection = types.BoolValue(res.Data.EnableBotDetection)
	state.IsLoginSuccessPageEnabled = types.BoolValue(res.Data.IsLoginSuccessPageEnabled)
	state.IsRegisterSuccessPageEnabled = types.BoolValue(res.Data.IsRegisterSuccessPageEnabled)
	state.AdminClient = types.BoolValue(res.Data.AdminClient)
	state.IsGroupLoginSelectionEnabled = types.BoolValue(res.Data.IsGroupLoginSelectionEnabled)
	state.BackchannelLogoutURI = util.StringValueOrNull(&res.Data.BackchannelLogoutURI)
	state.LogoURI = util.StringValueOrNull(&res.Data.LogoURI)
	state.InitiateLoginURI = util.StringValueOrNull(&res.Data.InitiateLoginURI)
	state.RegistrationClientURI = util.StringValueOrNull(&res.Data.RegistrationClientURI)
	state.RegistrationAccessToken = util.StringValueOrNull(&res.Data.RegistrationAccessToken)
	state.ClientURI = util.StringValueOrNull(&res.Data.ClientURI)
	state.JwksURI = util.StringValueOrNull(&res.Data.JwksURI)
	state.Jwks = util.StringValueOrNull(&res.Data.Jwks)
	state.SectorIdentifierURI = util.StringValueOrNull(&res.Data.SectorIdentifierURI)
	state.SubjectType = util.StringValueOrNull(&res.Data.SubjectType)
	state.IDTokenSignedResponseAlg = util.StringValueOrNull(&res.Data.IDTokenSignedResponseAlg)
	state.IDTokenEncryptedResponseAlg = util.StringValueOrNull(&res.Data.IDTokenEncryptedResponseAlg)
	state.IDTokenEncryptedResponseEnc = util.StringValueOrNull(&res.Data.IDTokenEncryptedResponseEnc)
	state.UserinfoSignedResponseAlg = util.StringValueOrNull(&res.Data.UserinfoSignedResponseAlg)
	state.UserinfoEncryptedResponseAlg = util.StringValueOrNull(&res.Data.UserinfoEncryptedResponseAlg)
	state.UserinfoEncryptedResponseEnc = util.StringValueOrNull(&res.Data.UserinfoEncryptedResponseEnc)
	state.RequestObjectSigningAlg = util.StringValueOrNull(&res.Data.RequestObjectSigningAlg)
	state.RequestObjectEncryptionAlg = util.StringValueOrNull(&res.Data.RequestObjectEncryptionAlg)
	state.RequestObjectEncryptionEnc = util.StringValueOrNull(&res.Data.RequestObjectEncryptionEnc)
	state.Description = util.StringValueOrNull(&res.Data.Description)
	state.ConsentPageGroup = util.StringValueOrNull(&res.Data.ConsentPageGroup)
	state.PasswordPolicyRef = util.StringValueOrNull(&res.Data.PasswordPolicyRef)
	state.BlockingMechanismRef = util.StringValueOrNull(&res.Data.BlockingMechanismRef)
	state.Sub = util.StringValueOrNull(&res.Data.Sub)
	state.Role = util.StringValueOrNull(&res.Data.Role)
	state.MfaConfiguration = util.StringValueOrNull(&res.Data.MfaConfiguration)
	state.BackgroundURI = util.StringValueOrNull(&res.Data.BackgroundURI)
	state.VideoURL = util.StringValueOrNull(&res.Data.VideoURL)
	state.BotCaptchaRef = util.StringValueOrNull(&res.Data.BotCaptchaRef)
	state.CreatedAt = util.StringValueOrNull(&res.Data.CreatedTime)
	state.UpdatedAt = util.StringValueOrNull(&res.Data.UpdatedTime)

	resp.Diagnostics.Append(extractSetValues(ctx, &state.PostLogoutRedirectUris, res.Data.PostLogoutRedirectUris)...)
	resp.Diagnostics.Append(extractSetValues(ctx, &state.GroupIDs, res.Data.GroupIDs)...)
	resp.Diagnostics.Append(extractSetValues(ctx, &state.AllowedMfa, res.Data.AllowedMfa)...)
	resp.Diagnostics.Append(extractSetValues(ctx, &state.AllowedFields, res.Data.AllowedFields)...)
	resp.Diagnostics.Append(extractSetValues(ctx, &state.WebMessageUris, res.Data.WebMessageUris)...)
	resp.Diagnostics.Append(extractSetValues(ctx, &state.DefaultAcrValues, res.Data.DefaultAcrValues)...)
	resp.Diagnostics.Append(extractSetValues(ctx, &state.GroupTypes, res.Data.GroupTypes)...)
	resp.Diagnostics.Append(extractSetValues(ctx, &state.AllowedRoles, res.Data.AllowedRoles)...)
	resp.Diagnostics.Append(extractSetValues(ctx, &state.DefaultRoles, res.Data.DefaultRoles)...)
	resp.Diagnostics.Append(extractSetValues(ctx, &state.SuggestMfa, res.Data.SuggestMfa)...)
	resp.Diagnostics.Append(extractSetValues(ctx, &state.CaptchaRefs, res.Data.CaptchaRefs)...)
	resp.Diagnostics.Append(extractSetValues(ctx, &state.ConsentRefs, res.Data.ConsentRefs)...)
	resp.Diagnostics.Append(extractSetValues(ctx, &state.DefaultScopes, res.Data.DefaultScopes)...)
	resp.Diagnostics.Append(extractSetValues(ctx, &state.PendingScopes, res.Data.PendingScopes)...)
	resp.Diagnostics.Append(extractSetValues(ctx, &state.LoginProviders, res.Data.LoginProviders)...)
	resp.Diagnostics.Append(extractSetValues(ctx, &state.AdditionalAccessTokenPayload, res.Data.AdditionalAccessTokenPayload)...)
	resp.Diagnostics.Append(extractSetValues(ctx, &state.RequiredFields, res.Data.RequiredFields)...)
	resp.Diagnostics.Append(extractSetValues(ctx, &state.AllowedWebOrigins, res.Data.AllowedWebOrigins)...)
	resp.Diagnostics.Append(extractSetValues(ctx, &state.AllowedOrigins, res.Data.AllowedOrigins)...)
	resp.Diagnostics.Append(extractSetValues(ctx, &state.Contacts, res.Data.Contacts)...)
	resp.Diagnostics.Append(extractSetValues(ctx, &state.RequestUris, res.Data.RequestUris)...)

	if resp.Diagnostics.HasError() {
		return resp
	}

	metaData := map[string]attr.Value{}
	for key, value := range res.Data.ApplicationMetaData {
		metaData[key] = util.StringValueOrNull(&value)
	}

	if len(res.Data.ApplicationMetaData) > 0 {
		appMetadata, d := types.MapValue(types.StringType, metaData)
		resp.Diagnostics.Append(d...)
		if resp.Diagnostics.HasError() {
			return resp
		}
		state.ApplicationMetaData = appMetadata
	}

	if res.Data.LoginSpi != nil {
		loginSpi, diags := types.ObjectValue(
			map[string]attr.Type{
				"oauth_client_id": types.StringType,
				"spi_url":         types.StringType,
			},
			map[string]attr.Value{
				"oauth_client_id": util.StringValueOrNull(&res.Data.LoginSpi.OauthClientID),
				"spi_url":         util.StringValueOrNull(&res.Data.LoginSpi.SpiURL),
			})
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return resp
		}
		state.LoginSpi = loginSpi
	}

	if res.Data.GroupSelection != nil {
		groupSelection, diags := types.ObjectValue(
			map[string]attr.Type{
				"always_show_group_selection": types.BoolType,
				"selectable_groups":           types.SetType{ElemType: types.StringType},
				"selectable_group_types":      types.SetType{ElemType: types.StringType},
			},
			map[string]attr.Value{
				"always_show_group_selection": util.BoolValueOrNull(&res.Data.GroupSelection.AlwaysShowGroupSelection),
				"selectable_groups":           util.SetValueOrNull(res.Data.GroupSelection.SelectableGroups),
				"selectable_group_types":      util.SetValueOrNull(res.Data.GroupSelection.SelectableGroupTypes),
			})
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return resp
		}
		state.GroupSelection = groupSelection
	}

	if res.Data.MobileSettings != nil {
		mobileSettings, diags := types.ObjectValue(
			map[string]attr.Type{
				"team_id":      types.StringType,
				"bundle_id":    types.StringType,
				"package_name": types.StringType,
				"key_hash":     types.StringType,
			},
			map[string]attr.Value{
				"team_id":      util.StringValueOrNull(&res.Data.MobileSettings.BundleID),
				"bundle_id":    util.StringValueOrNull(&res.Data.MobileSettings.BundleID),
				"package_name": util.StringValueOrNull(&res.Data.MobileSettings.PackageName),
				"key_hash":     util.StringValueOrNull(&res.Data.MobileSettings.KeyHash),
			})
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return resp
		}
		state.MobileSettings = mobileSettings
	}

	allowedGroupsObjectType := types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_id":      types.StringType,
			"roles":         types.SetType{ElemType: types.StringType},
			"default_roles": types.SetType{ElemType: types.StringType},
		},
	}

	var aglgObjectValues []attr.Value
	for _, aglg := range res.Data.AllowGuestLoginGroups {
		objValue := types.ObjectValueMust(
			allowedGroupsObjectType.AttrTypes,
			map[string]attr.Value{
				"group_id": util.StringValueOrNull(&aglg.GroupID),
				"roles": types.SetValueMust(types.StringType, func() []attr.Value {
					var temp []attr.Value
					for _, role := range aglg.Roles {
						temp = append(temp, util.StringValueOrNull(&role))
					}
					return temp
				}()),
				"default_roles": types.SetValueMust(types.StringType, func() []attr.Value {
					var temp []attr.Value
					for _, role := range aglg.DefaultRoles {
						temp = append(temp, util.StringValueOrNull(&role))
					}
					return temp
				}()),
			})
		aglgObjectValues = append(aglgObjectValues, objValue)
	}

	state.AllowGuestLoginGroups, d = types.ListValueFrom(ctx, allowedGroupsObjectType, aglgObjectValues)
	resp.Diagnostics.Append(d...)
	if resp.Diagnostics.HasError() {
		return resp
	}
	return resp
}

var _ planmodifier.Object = commonConfigConflictVerifier{}
var _ planmodifier.String = stringCustomRequired{}
var _ planmodifier.Set = setCustomRequired{}

type commonConfigConflictVerifier struct{}
type stringCustomRequired struct{}
type setCustomRequired struct{}

func (v commonConfigConflictVerifier) Description(_ context.Context) string {
	return "Verifies the availability of config details for the provided auth_type."
}

func (v commonConfigConflictVerifier) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

// Example
/*
	module "shared_settings" {
		source = "./modules"
		common_configs = {
			company_name    = "Company One"
			company_website = "https://cidaas.com"
		}
	}

	variable "common_configs" {
		type = object({
			company_name    = string
			company_website = string
		})
	}

	resource "cidaas_app" "sample_one" {
		common_configs = var.common_configs
	}

	resource "cidaas_app" "sample_two" {
		common_configs = var.common_configs
	}

// Overriding an attribute but still using common_configs for other shared attributes

	resource "cidaas_app" "sample_three" {
		common_configs = var.common_configs
		company_name   = "Company Two"
	}
*/
func (v commonConfigConflictVerifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	if req.ConfigValue.IsUnknown() || req.ConfigValue.IsNull() {
		return
	}
	commonConfig := CommonConfigs{}
	diags := req.ConfigValue.As(ctx, &commonConfig, basetypes.ObjectAsOptions{})
	resp.Diagnostics.Append(diags...)

	var company_name types.String

	diags = req.Config.GetAttribute(ctx, path.Root("company_name"), &company_name)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	message := "\033[1mRecommendations:\033[0m\n\n" +
		"  - Utilize common_configs only when there are shared configuration attributes for multiple resources.\n" +
		"  - If you need to override any specific attribute for a particular resource, you can supply the main configuration attribute directly within the resource block.\n" +
		"  - If your configuration involves a single resource or if the common configuration attributes are not shared across multiple resources, we do not suggest using common_configs."

	if !commonConfig.CompanyName.IsNull() && company_name.ValueString() != "" {
		resp.Diagnostics.AddWarning(
			`Identical attributes found in both main config and common_configs in some cidaas_app resources.`+
				`The values from the main config will be used and the values in common_configs will be ignored for those identical attributes.`,
			message,
		)
		return
	}
}

func (v stringCustomRequired) Description(_ context.Context) string {
	return "Verifies if a string type attribute from common_configs required."
}

func (v stringCustomRequired) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

// with switch only one error will be returned at a time for multiple missing required attr
func (v stringCustomRequired) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if !req.ConfigValue.IsNull() {
		return
	}
	var commonConfig types.Object

	diags := req.Config.GetAttribute(ctx, path.Root("common_configs"), &commonConfig)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if commonConfig.IsNull() {
		resp.Diagnostics.AddError("Missing required argument",
			fmt.Sprintf("The argument %s is required, but no definition was found.", req.Path.String()))
		return
	}

	config := CommonConfigs{}
	diags = commonConfig.As(ctx, &config, basetypes.ObjectAsOptions{})
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	switch req.Path.String() {
	case "company_name":
		if config.CompanyName.IsNull() {
			resp.Diagnostics.AddError(missingRequiredCustomError(req.Path.String()))
		}
	case "company_address":
		if config.CompanyAddress.IsNull() {
			resp.Diagnostics.AddError(missingRequiredCustomError(req.Path.String()))
		}
	case "company_website":
		if config.CompanyWebsite.IsNull() {
			resp.Diagnostics.AddError(missingRequiredCustomError(req.Path.String()))
		}
	case "client_type":
		if config.ClientType.IsNull() {
			resp.Diagnostics.AddError(missingRequiredCustomError(req.Path.String()))
		}
	default:
		resp.Diagnostics.AddError("Invalid stringCustomRequired configuration ",
			fmt.Sprintf("The argument %s is configured with stringCustomRequired but not part of common_configs schema.", req.Path.String()))
	}
}

func missingRequiredCustomError(argName string) (string, string) {
	return "Missing required argument",
		fmt.Sprintf("The argument %s is required, but no definition was found.", argName)
}

func (v setCustomRequired) Description(_ context.Context) string {
	return "Verifies if a set type attribute from common_configs required."
}

func (v setCustomRequired) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v setCustomRequired) PlanModifySet(ctx context.Context, req planmodifier.SetRequest, resp *planmodifier.SetResponse) {
	if !req.ConfigValue.IsNull() {
		return
	}
	var commonConfig types.Object

	diags := req.Config.GetAttribute(ctx, path.Root("common_configs"), &commonConfig)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if commonConfig.IsNull() {
		resp.Diagnostics.AddError("Missing required argument",
			fmt.Sprintf("The argument %s is required, but no definition was found.", req.Path.String()))
		return
	}

	config := CommonConfigs{}
	diags = commonConfig.As(ctx, &config, basetypes.ObjectAsOptions{})
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// with switch only one error will be returned at a time for multiple missing required attr
	switch req.Path.String() {
	case "allowed_scopes":
		if config.AllowedScopes.IsNull() {
			resp.Diagnostics.AddError(missingRequiredCustomError(req.Path.String()))
		}
	case "redirect_uris":
		if config.RedirectUris.IsNull() {
			resp.Diagnostics.AddError(missingRequiredCustomError(req.Path.String()))
		}
	case "allowed_logout_urls":
		if config.AllowedLogoutUrls.IsNull() {
			resp.Diagnostics.AddError(missingRequiredCustomError(req.Path.String()))
		}
	default:
		resp.Diagnostics.AddError("Invalid setCustomRequired configuration ",
			fmt.Sprintf("The argument %s is configured with setCustomRequired but not part of common_configs schema.", req.Path.String()))
	}
}
