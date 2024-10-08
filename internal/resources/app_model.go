package resources

import (
	"context"
	"fmt"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/cidaas"
	"github.com/Cidaas/terraform-provider-cidaas/helpers/util"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const (
	CREATE = 1
	READ   = 2
	IMPORT = 3
	UPDATE = 4
)

type AppConfig struct {
	ID                              types.String `tfsdk:"id"`
	ClientType                      types.String `tfsdk:"client_type"`
	AccentColor                     types.String `tfsdk:"accent_color"`
	PrimaryColor                    types.String `tfsdk:"primary_color"`
	MediaType                       types.String `tfsdk:"media_type"`
	ContentAlign                    types.String `tfsdk:"content_align"`
	HostedPageGroup                 types.String `tfsdk:"hosted_page_group"`
	ClientName                      types.String `tfsdk:"client_name"`
	ClientDisplayName               types.String `tfsdk:"client_display_name"`
	CompanyName                     types.String `tfsdk:"company_name"`
	CompanyAddress                  types.String `tfsdk:"company_address"`
	CompanyWebsite                  types.String `tfsdk:"company_website"`
	TemplateGroupID                 types.String `tfsdk:"template_group_id"`
	ClientID                        types.String `tfsdk:"client_id"`
	ClientSecret                    types.String `tfsdk:"client_secret"`
	PolicyURI                       types.String `tfsdk:"policy_uri"`
	TosURI                          types.String `tfsdk:"tos_uri"`
	ImprintURI                      types.String `tfsdk:"imprint_uri"`
	TokenEndpointAuthMethod         types.String `tfsdk:"token_endpoint_auth_method"`
	TokenEndpointAuthSigningAlg     types.String `tfsdk:"token_endpoint_auth_signing_alg"`
	CaptchaRef                      types.String `tfsdk:"captcha_ref"`
	CommunicationMediumVerification types.String `tfsdk:"communication_medium_verification"`
	BotProvider                     types.String `tfsdk:"bot_provider"`
	BackchannelLogoutURI            types.String `tfsdk:"backchannel_logout_uri"`
	LogoAlign                       types.String `tfsdk:"logo_align"`
	Webfinger                       types.String `tfsdk:"webfinger"`
	LogoURI                         types.String `tfsdk:"logo_uri"`
	InitiateLoginURI                types.String `tfsdk:"initiate_login_uri"`
	RegistrationClientURI           types.String `tfsdk:"registration_client_uri"`
	RegistrationAccessToken         types.String `tfsdk:"registration_access_token"`
	ClientURI                       types.String `tfsdk:"client_uri"`
	JwksURI                         types.String `tfsdk:"jwks_uri"`
	Jwks                            types.String `tfsdk:"jwks"`
	SectorIdentifierURI             types.String `tfsdk:"sector_identifier_uri"`
	SubjectType                     types.String `tfsdk:"subject_type"`
	IDTokenSignedResponseAlg        types.String `tfsdk:"id_token_signed_response_alg"`
	IDTokenEncryptedResponseAlg     types.String `tfsdk:"id_token_encrypted_response_alg"`
	IDTokenEncryptedResponseEnc     types.String `tfsdk:"id_token_encrypted_response_enc"`
	UserinfoSignedResponseAlg       types.String `tfsdk:"userinfo_signed_response_alg"`
	UserinfoEncryptedResponseAlg    types.String `tfsdk:"userinfo_encrypted_response_alg"`
	UserinfoEncryptedResponseEnc    types.String `tfsdk:"userinfo_encrypted_response_enc"`
	RequestObjectSigningAlg         types.String `tfsdk:"request_object_signing_alg"`
	RequestObjectEncryptionAlg      types.String `tfsdk:"request_object_encryption_alg"`
	RequestObjectEncryptionEnc      types.String `tfsdk:"request_object_encryption_enc"`
	Description                     types.String `tfsdk:"description"`
	ConsentPageGroup                types.String `tfsdk:"consent_page_group"`
	PasswordPolicyRef               types.String `tfsdk:"password_policy_ref"`
	BlockingMechanismRef            types.String `tfsdk:"blocking_mechanism_ref"`
	Sub                             types.String `tfsdk:"sub"`
	Role                            types.String `tfsdk:"role"`
	MfaConfiguration                types.String `tfsdk:"mfa_configuration"`
	BackgroundURI                   types.String `tfsdk:"background_uri"`
	VideoURL                        types.String `tfsdk:"video_url"`
	BotCaptchaRef                   types.String `tfsdk:"bot_captcha_ref"`

	EnableDeduplication              types.Bool `tfsdk:"enable_deduplication"`
	AutoLoginAfterRegister           types.Bool `tfsdk:"auto_login_after_register"`
	EnablePasswordlessAuth           types.Bool `tfsdk:"enable_passwordless_auth"`
	RegisterWithLoginInformation     types.Bool `tfsdk:"register_with_login_information"`
	AllowDisposableEmail             types.Bool `tfsdk:"allow_disposable_email"`
	ValidatePhoneNumber              types.Bool `tfsdk:"validate_phone_number"`
	FdsEnabled                       types.Bool `tfsdk:"fds_enabled"`
	IsHybridApp                      types.Bool `tfsdk:"is_hybrid_app"`
	Editable                         types.Bool `tfsdk:"editable"`
	JweEnabled                       types.Bool `tfsdk:"jwe_enabled"`
	UserConsent                      types.Bool `tfsdk:"user_consent"`
	Enabled                          types.Bool `tfsdk:"enabled"`
	AlwaysAskMfa                     types.Bool `tfsdk:"always_ask_mfa"`
	SmartMfa                         types.Bool `tfsdk:"smart_mfa"`
	EmailVerificationRequired        types.Bool `tfsdk:"email_verification_required"`
	MobileNumberVerificationRequired types.Bool `tfsdk:"mobile_number_verification_required"`
	EnableClassicalProvider          types.Bool `tfsdk:"enable_classical_provider"`
	IsRememberMeSelected             types.Bool `tfsdk:"is_remember_me_selected"`
	EnableBotDetection               types.Bool `tfsdk:"enable_bot_detection"`
	IsLoginSuccessPageEnabled        types.Bool `tfsdk:"is_login_success_page_enabled"`
	IsRegisterSuccessPageEnabled     types.Bool `tfsdk:"is_register_success_page_enabled"`
	IsGroupLoginSelectionEnabled     types.Bool `tfsdk:"is_group_login_selection_enabled"`
	AllowGuestLogin                  types.Bool `tfsdk:"allow_guest_login"`
	RequireAuthTime                  types.Bool `tfsdk:"require_auth_time"`
	EnableLoginSpi                   types.Bool `tfsdk:"enable_login_spi"`
	BackchannelLogoutSessionRequired types.Bool `tfsdk:"backchannel_logout_session_required"`

	DefaultMaxAge                 types.Int64 `tfsdk:"default_max_age"`
	TokenLifetimeInSeconds        types.Int64 `tfsdk:"token_lifetime_in_seconds"`
	IDTokenLifetimeInSeconds      types.Int64 `tfsdk:"id_token_lifetime_in_seconds"`
	RefreshTokenLifetimeInSeconds types.Int64 `tfsdk:"refresh_token_lifetime_in_seconds"`

	AllowLoginWith               types.Set `tfsdk:"allow_login_with"`
	RedirectURIS                 types.Set `tfsdk:"redirect_uris"`
	AllowedLogoutUrls            types.Set `tfsdk:"allowed_logout_urls"`
	AllowedScopes                types.Set `tfsdk:"allowed_scopes"`
	ResponseTypes                types.Set `tfsdk:"response_types"`
	GrantTypes                   types.Set `tfsdk:"grant_types"`
	LoginProviders               types.Set `tfsdk:"login_providers"`
	AdditionalAccessTokenPayload types.Set `tfsdk:"additional_access_token_payload"`
	RequiredFields               types.Set `tfsdk:"required_fields"`
	AllowedWebOrigins            types.Set `tfsdk:"allowed_web_origins"`
	AllowedOrigins               types.Set `tfsdk:"allowed_origins"`
	Contacts                     types.Set `tfsdk:"contacts"`
	DefaultAcrValues             types.Set `tfsdk:"default_acr_values"`
	WebMessageUris               types.Set `tfsdk:"web_message_uris"`
	AllowedFields                types.Set `tfsdk:"allowed_fields"`
	AllowedMfa                   types.Set `tfsdk:"allowed_mfa"`
	CaptchaRefs                  types.Set `tfsdk:"captcha_refs"`
	ConsentRefs                  types.Set `tfsdk:"consent_refs"`
	AllowedRoles                 types.Set `tfsdk:"allowed_roles"`
	DefaultRoles                 types.Set `tfsdk:"default_roles"`
	GroupIDs                     types.Set `tfsdk:"group_ids"`
	GroupTypes                   types.Set `tfsdk:"group_types"`
	PostLogoutRedirectUris       types.Set `tfsdk:"post_logout_redirect_uris"`
	RequestUris                  types.Set `tfsdk:"request_uris"`
	DefaultScopes                types.Set `tfsdk:"default_scopes"`
	PendingScopes                types.Set `tfsdk:"pending_scopes"`
	SuggestMfa                   types.Set `tfsdk:"suggest_mfa"`

	ApplicationMetaData types.Map `tfsdk:"application_meta_data"`

	SocialProviders         types.List `tfsdk:"social_providers"`
	CustomProviders         types.List `tfsdk:"custom_providers"`
	SamlProviders           types.List `tfsdk:"saml_providers"`
	AdProviders             types.List `tfsdk:"ad_providers"`
	AllowedGroups           types.List `tfsdk:"allowed_groups"`
	OperationsAllowedGroups types.List `tfsdk:"operations_allowed_groups"`
	AllowGuestLoginGroups   types.List `tfsdk:"allow_guest_login_groups"`

	socialProviders         []*SocialProviderData
	customProviders         []*ProviderMetadData
	samlProviders           []*ProviderMetadData
	adProviders             []*ProviderMetadData
	allowedGroups           []*AllowedGroups
	operationsAllowedGroups []*AllowedGroups
	allowGuestLoginGroups   []*AllowedGroups

	LoginSpi                   types.Object `tfsdk:"login_spi"`
	GroupSelection             types.Object `tfsdk:"group_selection"`
	Mfa                        types.Object `tfsdk:"mfa"`
	MobileSettings             types.Object `tfsdk:"mobile_settings"`
	SuggestVerificationMethods types.Object `tfsdk:"suggest_verification_methods"`
	GroupRoleRestriction       types.Object `tfsdk:"group_role_restriction"`
	BasicSettings              types.Object `tfsdk:"basic_settings"`

	loginSpi                   *LoginSPI
	groupSelection             *GroupSelection
	mfa                        *MfaOption
	mobileSettings             *AppMobileSettings
	suggestVerificationMethods *SuggestVerificationMethods
	groupRoleRestriction       *GroupRoleRestriction
	basicSettings              *BasicSettings

	CommonConfigs types.Object `tfsdk:"common_configs"`

	commonConfigs *CommonConfigs
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
type SuggestVerificationMethods struct {
	MandatoryConfig    types.Object `tfsdk:"mandatory_config"`
	OptionalConfig     types.Object `tfsdk:"optional_config"`
	SkipDurationInDays types.Int32  `tfsdk:"skip_duration_in_days"`
}

type MandatoryConfig struct {
	Methods   types.Set    `tfsdk:"methods"`
	SkipUntil types.String `tfsdk:"skip_until"`
	Range     types.String `tfsdk:"range"`
}

type OptionalConfig struct {
	Methods types.Set `tfsdk:"methods"`
}

type GroupRoleRestriction struct {
	MatchCondition types.String `tfsdk:"match_condition"`
	Filters        types.List   `tfsdk:"filters"`
}
type GroupRoleFilters struct {
	GroupID    types.String `tfsdk:"group_id"`
	RoleFilter types.Object `tfsdk:"role_filter"`
}

type RoleFilter struct {
	MatchCondition types.String `tfsdk:"match_condition"`
	Roles          types.Set    `tfsdk:"roles"`
}

type CommonConfigs struct {
	CompanyName    types.String `tfsdk:"company_name"`
	CompanyWebsite types.String `tfsdk:"company_website"`
	ClientType     types.String `tfsdk:"client_type"`
	CompanyAddress types.String `tfsdk:"company_address"`

	AllowedScopes     types.Set `tfsdk:"allowed_scopes"`
	RedirectUris      types.Set `tfsdk:"redirect_uris"`
	AllowedLogoutUrls types.Set `tfsdk:"allowed_logout_urls"`
	AllowedWebOrigins types.Set `tfsdk:"allowed_web_origins"`
	AllowedOrigins    types.Set `tfsdk:"allowed_origins"`
	LoginProviders    types.Set `tfsdk:"login_providers"`
	DefaultScopes     types.Set `tfsdk:"default_scopes"`
	PendingScopes     types.Set `tfsdk:"pending_scopes"`
	AllowedMfa        types.Set `tfsdk:"allowed_mfa"`
	AllowedRoles      types.Set `tfsdk:"allowed_roles"`
	DefaultRoles      types.Set `tfsdk:"default_roles"`

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

type BasicSettings struct {
	ClientID          types.String `tfsdk:"client_id"`
	RedirectURIs      types.Set    `tfsdk:"redirect_uris"`
	AllowedLogoutUrls types.Set    `tfsdk:"allowed_logout_urls"`
	AllowedScopes     types.Set    `tfsdk:"allowed_scopes"`
	ClientSecrets     types.List   `tfsdk:"client_secrets"`
}

type ClientSecret struct {
	ClientSecret          types.String `tfsdk:"client_secret"`
	ClientSecretExpiresAt types.Int64  `tfsdk:"client_secret_expires_at"`
}

func (w *AppConfig) ExtractAppConfigs(ctx context.Context) diag.Diagnostics {
	var diags diag.Diagnostics
	if !w.LoginSpi.IsNull() && !w.LoginSpi.IsUnknown() {
		w.loginSpi = &LoginSPI{}
		diags = w.LoginSpi.As(ctx, w.loginSpi, basetypes.ObjectAsOptions{})
	}
	if !w.GroupSelection.IsNull() && !w.GroupSelection.IsUnknown() {
		w.groupSelection = &GroupSelection{}
		diags = w.GroupSelection.As(ctx, w.groupSelection, basetypes.ObjectAsOptions{})
	}
	if !w.Mfa.IsNull() && !w.Mfa.IsUnknown() {
		w.mfa = &MfaOption{}
		diags = w.Mfa.As(ctx, w.mfa, basetypes.ObjectAsOptions{})
	}
	if !w.MobileSettings.IsNull() && !w.MobileSettings.IsUnknown() {
		w.mobileSettings = &AppMobileSettings{}
		diags = w.MobileSettings.As(ctx, w.mobileSettings, basetypes.ObjectAsOptions{})
	}
	if !w.CommonConfigs.IsNull() && !w.CommonConfigs.IsUnknown() {
		w.commonConfigs = &CommonConfigs{}
		diags = w.CommonConfigs.As(ctx, w.commonConfigs, basetypes.ObjectAsOptions{})
	}
	if !w.SocialProviders.IsNull() && !w.SocialProviders.IsUnknown() {
		w.socialProviders = make([]*SocialProviderData, 0, len(w.SocialProviders.Elements()))
		diags = w.SocialProviders.ElementsAs(ctx, &w.socialProviders, false)
	}
	if !w.CustomProviders.IsNull() && !w.CustomProviders.IsUnknown() {
		w.customProviders = make([]*ProviderMetadData, 0, len(w.CustomProviders.Elements()))
		diags = w.CustomProviders.ElementsAs(ctx, &w.customProviders, false)
	}
	if !w.SamlProviders.IsNull() && !w.SamlProviders.IsUnknown() {
		w.samlProviders = make([]*ProviderMetadData, 0, len(w.SamlProviders.Elements()))
		diags = w.SamlProviders.ElementsAs(ctx, &w.samlProviders, false)
	}
	if !w.AdProviders.IsNull() && !w.AdProviders.IsUnknown() {
		w.adProviders = make([]*ProviderMetadData, 0, len(w.AdProviders.Elements()))
		diags = w.AdProviders.ElementsAs(ctx, &w.adProviders, false)
	}
	if !w.AllowedGroups.IsNull() && !w.AllowedGroups.IsUnknown() {
		w.allowedGroups = make([]*AllowedGroups, 0, len(w.AllowedGroups.Elements()))
		diags = w.AllowedGroups.ElementsAs(ctx, &w.allowedGroups, false)
	}
	if !w.OperationsAllowedGroups.IsNull() && !w.OperationsAllowedGroups.IsUnknown() {
		w.operationsAllowedGroups = make([]*AllowedGroups, 0, len(w.OperationsAllowedGroups.Elements()))
		diags = w.OperationsAllowedGroups.ElementsAs(ctx, &w.operationsAllowedGroups, false)
	}
	if !w.AllowGuestLoginGroups.IsNull() && !w.AllowGuestLoginGroups.IsUnknown() {
		w.allowGuestLoginGroups = make([]*AllowedGroups, 0, len(w.AllowGuestLoginGroups.Elements()))
		diags = w.AllowGuestLoginGroups.ElementsAs(ctx, &w.allowGuestLoginGroups, false)
	}
	if !w.SuggestVerificationMethods.IsNull() && !w.SuggestVerificationMethods.IsUnknown() {
		w.suggestVerificationMethods = &SuggestVerificationMethods{}
		diags = w.SuggestVerificationMethods.As(ctx, w.suggestVerificationMethods, basetypes.ObjectAsOptions{})
	}
	if !w.GroupRoleRestriction.IsNull() && !w.GroupRoleRestriction.IsUnknown() {
		w.groupRoleRestriction = &GroupRoleRestriction{}
		diags = w.GroupRoleRestriction.As(ctx, w.groupRoleRestriction, basetypes.ObjectAsOptions{})
	}
	if !w.BasicSettings.IsNull() && !w.BasicSettings.IsUnknown() {
		w.basicSettings = &BasicSettings{}
		diags = w.BasicSettings.As(ctx, w.basicSettings, basetypes.ObjectAsOptions{})
	}
	return diags
}

func prepareAppModel(ctx context.Context, plan AppConfig) (*cidaas.AppModel, diag.Diagnostics) { //nolint:gocognit

	assignSetValues := func(ctx context.Context, planValue basetypes.SetValue, target *[]string) diag.Diagnostics {
		if !planValue.IsNull() && len(planValue.Elements()) > 0 {
			return planValue.ElementsAs(ctx, target, false)
		}
		return nil
	}

	app := cidaas.AppModel{
		AppOwner:                         "CLIENT",
		CompanyName:                      plan.CompanyName.ValueString(),
		CompanyAddress:                   plan.CompanyAddress.ValueString(),
		CompanyWebsite:                   plan.CompanyWebsite.ValueString(),
		ClientType:                       plan.ClientType.ValueString(),
		AccentColor:                      plan.AccentColor.ValueString(),
		PrimaryColor:                     plan.PrimaryColor.ValueString(),
		MediaType:                        plan.MediaType.ValueString(),
		ContentAlign:                     plan.ContentAlign.ValueString(),
		EnableDeduplication:              plan.EnableDeduplication.ValueBoolPointer(),
		AutoLoginAfterRegister:           plan.AutoLoginAfterRegister.ValueBoolPointer(),
		EnablePasswordlessAuth:           plan.EnablePasswordlessAuth.ValueBoolPointer(),
		RegisterWithLoginInformation:     plan.RegisterWithLoginInformation.ValueBoolPointer(),
		AllowDisposableEmail:             plan.AllowDisposableEmail.ValueBoolPointer(),
		ValidatePhoneNumber:              plan.ValidatePhoneNumber.ValueBoolPointer(),
		FdsEnabled:                       plan.FdsEnabled.ValueBoolPointer(),
		HostedPageGroup:                  plan.HostedPageGroup.ValueString(),
		ClientName:                       plan.ClientName.ValueString(),
		ClientDisplayName:                plan.ClientDisplayName.ValueString(),
		IsHybridApp:                      plan.IsHybridApp.ValueBoolPointer(),
		DefaultMaxAge:                    plan.DefaultMaxAge.ValueInt64Pointer(),
		TokenLifetimeInSeconds:           plan.TokenLifetimeInSeconds.ValueInt64Pointer(),
		IDTokenLifetimeInSeconds:         plan.IDTokenLifetimeInSeconds.ValueInt64Pointer(),
		RefreshTokenLifetimeInSeconds:    plan.RefreshTokenLifetimeInSeconds.ValueInt64Pointer(),
		TemplateGroupID:                  plan.TemplateGroupID.ValueString(),
		ClientID:                         plan.ClientID.ValueString(),
		ClientSecret:                     plan.ClientSecret.ValueString(),
		PolicyURI:                        plan.PolicyURI.ValueString(),
		TosURI:                           plan.TosURI.ValueString(),
		ImprintURI:                       plan.ImprintURI.ValueString(),
		TokenEndpointAuthMethod:          plan.TokenEndpointAuthMethod.ValueString(),
		TokenEndpointAuthSigningAlg:      plan.TokenEndpointAuthSigningAlg.ValueString(),
		Editable:                         plan.Editable.ValueBoolPointer(),
		JweEnabled:                       plan.JweEnabled.ValueBoolPointer(),
		UserConsent:                      plan.UserConsent.ValueBoolPointer(),
		Enabled:                          plan.Enabled.ValueBoolPointer(),
		AlwaysAskMfa:                     plan.AlwaysAskMfa.ValueBoolPointer(),
		SmartMfa:                         plan.SmartMfa.ValueBoolPointer(),
		CaptchaRef:                       plan.CaptchaRef.ValueString(),
		CommunicationMediumVerification:  plan.CommunicationMediumVerification.ValueString(),
		EmailVerificationRequired:        plan.EmailVerificationRequired.ValueBoolPointer(),
		MobileNumberVerificationRequired: plan.MobileNumberVerificationRequired.ValueBoolPointer(),
		EnableClassicalProvider:          plan.EnableClassicalProvider.ValueBoolPointer(),
		IsRememberMeSelected:             plan.IsRememberMeSelected.ValueBoolPointer(),
		EnableBotDetection:               plan.EnableBotDetection.ValueBoolPointer(),
		BotProvider:                      plan.BotProvider.ValueString(),
		IsLoginSuccessPageEnabled:        plan.IsLoginSuccessPageEnabled.ValueBoolPointer(),
		IsRegisterSuccessPageEnabled:     plan.IsRegisterSuccessPageEnabled.ValueBoolPointer(),
		IsGroupLoginSelectionEnabled:     plan.IsGroupLoginSelectionEnabled.ValueBoolPointer(),
		BackchannelLogoutURI:             plan.BackchannelLogoutURI.ValueString(),
		LogoAlign:                        plan.LogoAlign.ValueString(),
		Webfinger:                        plan.Webfinger.ValueString(),
		LogoURI:                          plan.LogoURI.ValueString(),
		InitiateLoginURI:                 plan.InitiateLoginURI.ValueString(),
		RegistrationClientURI:            plan.RegistrationClientURI.ValueString(),
		RegistrationAccessToken:          plan.RegistrationAccessToken.ValueString(),
		ClientURI:                        plan.ClientURI.ValueString(),
		JwksURI:                          plan.JwksURI.ValueString(),
		Jwks:                             plan.Jwks.ValueString(),
		SectorIdentifierURI:              plan.SectorIdentifierURI.ValueString(),
		SubjectType:                      plan.SubjectType.ValueString(),
		IDTokenSignedResponseAlg:         plan.IDTokenSignedResponseAlg.ValueString(),
		IDTokenEncryptedResponseAlg:      plan.IDTokenEncryptedResponseAlg.ValueString(),
		IDTokenEncryptedResponseEnc:      plan.IDTokenEncryptedResponseEnc.ValueString(),
		UserinfoSignedResponseAlg:        plan.UserinfoSignedResponseAlg.ValueString(),
		UserinfoEncryptedResponseAlg:     plan.UserinfoEncryptedResponseAlg.ValueString(),
		UserinfoEncryptedResponseEnc:     plan.UserinfoEncryptedResponseEnc.ValueString(),
		RequestObjectSigningAlg:          plan.RequestObjectSigningAlg.ValueString(),
		RequestObjectEncryptionAlg:       plan.RequestObjectEncryptionAlg.ValueString(),
		RequestObjectEncryptionEnc:       plan.RequestObjectEncryptionEnc.ValueString(),
		Description:                      plan.Description.ValueString(),
		ConsentPageGroup:                 plan.ConsentPageGroup.ValueString(),
		BlockingMechanismRef:             plan.BlockingMechanismRef.ValueString(),
		Sub:                              plan.Sub.ValueString(),
		Role:                             plan.Role.ValueString(),
		MfaConfiguration:                 plan.MfaConfiguration.ValueString(),
		AllowGuestLogin:                  plan.AllowGuestLogin.ValueBoolPointer(),
		BackgroundURI:                    plan.BackgroundURI.ValueString(),
		VideoURL:                         plan.VideoURL.ValueString(),
		BotCaptchaRef:                    plan.BotCaptchaRef.ValueString(),
		RequireAuthTime:                  plan.RequireAuthTime.ValueBoolPointer(),
		EnableLoginSpi:                   plan.EnableLoginSpi.ValueBoolPointer(),
		BackchannelLogoutSessionRequired: plan.BackchannelLogoutSessionRequired.ValueBoolPointer(),
	}

	if plan.PasswordPolicyRef.IsNull() {
		app.PasswordPolicyRef = ""
	} else {
		app.PasswordPolicyRef = plan.PasswordPolicyRef.ValueString()
	}

	var diags diag.Diagnostics

	diags.Append(assignSetValues(ctx, plan.AllowedScopes, &app.AllowedScopes)...)
	diags.Append(assignSetValues(ctx, plan.RedirectURIS, &app.RedirectURIS)...)
	diags.Append(assignSetValues(ctx, plan.AllowedLogoutUrls, &app.AllowedLogoutUrls)...)
	diags.Append(assignSetValues(ctx, plan.AllowedWebOrigins, &app.AllowedWebOrigins)...)
	diags.Append(assignSetValues(ctx, plan.AllowedOrigins, &app.AllowedOrigins)...)
	diags.Append(assignSetValues(ctx, plan.LoginProviders, &app.LoginProviders)...)
	diags.Append(assignSetValues(ctx, plan.DefaultScopes, &app.DefaultScopes)...)
	diags.Append(assignSetValues(ctx, plan.PendingScopes, &app.PendingScopes)...)
	diags.Append(assignSetValues(ctx, plan.AllowedMfa, &app.AllowedMfa)...)
	diags.Append(assignSetValues(ctx, plan.AllowedRoles, &app.AllowedRoles)...)
	diags.Append(assignSetValues(ctx, plan.DefaultRoles, &app.DefaultRoles)...)
	diags.Append(assignSetValues(ctx, plan.PostLogoutRedirectUris, &app.PostLogoutRedirectUris)...)
	diags.Append(assignSetValues(ctx, plan.GroupIDs, &app.GroupIDs)...)
	diags.Append(assignSetValues(ctx, plan.AllowedFields, &app.AllowedFields)...)
	diags.Append(assignSetValues(ctx, plan.WebMessageUris, &app.WebMessageUris)...)
	diags.Append(assignSetValues(ctx, plan.DefaultAcrValues, &app.DefaultAcrValues)...)
	diags.Append(assignSetValues(ctx, plan.GroupTypes, &app.GroupTypes)...)
	diags.Append(assignSetValues(ctx, plan.SuggestMfa, &app.SuggestMfa)...)
	diags.Append(assignSetValues(ctx, plan.CaptchaRefs, &app.CaptchaRefs)...)
	diags.Append(assignSetValues(ctx, plan.ConsentRefs, &app.ConsentRefs)...)
	diags.Append(assignSetValues(ctx, plan.AllowLoginWith, &app.AllowLoginWith)...)
	diags.Append(assignSetValues(ctx, plan.AllowedLogoutUrls, &app.AllowedLogoutUrls)...)
	diags.Append(assignSetValues(ctx, plan.ResponseTypes, &app.ResponseTypes)...)
	diags.Append(assignSetValues(ctx, plan.GrantTypes, &app.GrantTypes)...)
	diags.Append(assignSetValues(ctx, plan.AdditionalAccessTokenPayload, &app.AdditionalAccessTokenPayload)...)
	diags.Append(assignSetValues(ctx, plan.RequiredFields, &app.RequiredFields)...)
	diags.Append(assignSetValues(ctx, plan.Contacts, &app.Contacts)...)
	diags.Append(assignSetValues(ctx, plan.RequestUris, &app.RequestUris)...)

	// SocialProviders
	if !plan.SocialProviders.IsNull() && len(plan.socialProviders) > 0 {
		target := []cidaas.ISocialProviderData{}
		var objectValues []attr.Value
		spObjectType := types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"provider_name": types.StringType,
				"social_id":     types.StringType,
			},
		}
		objValue := types.ObjectValueMust(
			spObjectType.AttrTypes,
			map[string]attr.Value{
				"provider_name": types.StringNull(),
				"social_id":     types.StringNull(),
			})
		objectValues = append(objectValues, objValue)
		emptyProvider := types.ListValueMust(spObjectType, objectValues)
		if plan.SocialProviders.Equal(emptyProvider) {
			app.SocialProviders = target
		} else {
			for _, sp := range plan.socialProviders {
				target = append(target, cidaas.ISocialProviderData{
					ProviderName: sp.ProviderName.ValueString(),
					SocialID:     sp.SocialID.ValueString(),
				})
				app.SocialProviders = target
			}
		}
	}

	// helper function to get the []strings from SetValue
	getSetAsStrings := func(ctx context.Context, planValue basetypes.SetValue) ([]string, diag.Diagnostics) {
		var result []string
		diags := assignSetValues(ctx, planValue, &result)
		return result, diags
	}

	setProviders := func(ctx context.Context, source []*ProviderMetadData, target *[]cidaas.IProviderMetadData) diag.Diagnostics {
		for _, cp := range source {
			temp := cidaas.IProviderMetadData{
				ProviderName:      cp.ProviderName.ValueString(),
				Type:              cp.Type.ValueString(),
				DisplayName:       cp.DisplayName.ValueString(),
				LogoURL:           cp.LogoURL.ValueString(),
				IsProviderVisible: cp.IsProviderVisible.ValueBoolPointer(),
			}
			diags := cp.Domains.ElementsAs(ctx, &temp.Domains, false)
			if diags.HasError() {
				return diags
			}
			*target = append(*target, temp)
		}
		return nil
	}

	var objectValues []attr.Value
	spObjectType := types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"provider_name":       types.StringType,
			"display_name":        types.StringType,
			"logo_url":            types.StringType,
			"type":                types.StringType,
			"is_provider_visible": types.BoolType,
			"domains":             types.SetType{ElemType: types.StringType},
		},
	}
	objValue := types.ObjectValueMust(
		spObjectType.AttrTypes,
		map[string]attr.Value{
			"provider_name":       types.StringNull(),
			"display_name":        types.StringNull(),
			"logo_url":            types.StringNull(),
			"type":                types.StringNull(),
			"is_provider_visible": types.BoolNull(),
			"domains":             types.SetNull(types.StringType),
		})
	objectValues = append(objectValues, objValue)
	emptyProvider := types.ListValueMust(spObjectType, objectValues)

	// CustomProviders
	if !plan.CustomProviders.IsNull() && len(plan.customProviders) > 0 {
		if plan.CustomProviders.Equal(emptyProvider) {
			app.CustomProviders = []cidaas.IProviderMetadData{}
		} else {
			diags.Append(setProviders(ctx, plan.customProviders, &app.CustomProviders)...)
			if diags.HasError() {
				return nil, diags
			}
		}
	}

	// SamlProviders
	if !plan.SamlProviders.IsNull() && len(plan.samlProviders) > 0 {
		if plan.SamlProviders.Equal(emptyProvider) {
			app.SamlProviders = []cidaas.IProviderMetadData{}
		} else {
			diags.Append(setProviders(ctx, plan.samlProviders, &app.SamlProviders)...)
			if diags.HasError() {
				return nil, diags
			}
		}
	}
	// AdProviders
	if !plan.AdProviders.IsNull() && len(plan.adProviders) > 0 {
		if plan.AdProviders.Equal(emptyProvider) {
			app.AdProviders = []cidaas.IProviderMetadData{}
		} else {
			diags.Append(setProviders(ctx, plan.adProviders, &app.AdProviders)...)
			if diags.HasError() {
				return nil, diags
			}
		}
	}

	var allowedGroupObjectValues []attr.Value
	agObjectType := types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_id":      types.StringType,
			"roles":         types.SetType{ElemType: types.StringType},
			"default_roles": types.SetType{ElemType: types.StringType},
		},
	}
	value := types.ObjectValueMust(
		agObjectType.AttrTypes,
		map[string]attr.Value{
			"group_id":      types.StringNull(),
			"roles":         types.SetNull(types.StringType),
			"default_roles": types.SetNull(types.StringType),
		})
	allowedGroupObjectValues = append(allowedGroupObjectValues, value)
	emptyAllowedGroups := types.ListValueMust(agObjectType, allowedGroupObjectValues)

	// AllowedGroups
	if !plan.AllowedGroups.IsNull() && len(plan.allowedGroups) > 0 {
		if plan.AllowedGroups.Equal(emptyAllowedGroups) {
			app.AllowedGroups = []cidaas.IAllowedGroups{}
		} else {
			for _, ag := range plan.allowedGroups {
				temp := cidaas.IAllowedGroups{
					GroupID: ag.GroupID.ValueString(),
				}
				diags := ag.Roles.ElementsAs(ctx, &temp.Roles, false)
				if diags.HasError() {
					return nil, diags
				}
				diags = ag.DefaultRoles.ElementsAs(ctx, &temp.DefaultRoles, false)
				if diags.HasError() {
					return nil, diags
				}
				app.AllowedGroups = append(app.AllowedGroups, temp)
			}
		}
	}

	// OperationsAllowedGroups
	if !plan.OperationsAllowedGroups.IsNull() && len(plan.operationsAllowedGroups) > 0 {
		if plan.OperationsAllowedGroups.Equal(emptyAllowedGroups) {
			app.OperationsAllowedGroups = []cidaas.IAllowedGroups{}
		} else {
			for _, oag := range plan.operationsAllowedGroups {
				temp := cidaas.IAllowedGroups{
					GroupID: oag.GroupID.ValueString(),
				}
				diags := oag.Roles.ElementsAs(ctx, &temp.Roles, false)
				if diags.HasError() {
					return nil, diags
				}
				diags = oag.DefaultRoles.ElementsAs(ctx, &temp.DefaultRoles, false)
				if diags.HasError() {
					return nil, diags
				}
				app.OperationsAllowedGroups = append(app.OperationsAllowedGroups, temp)
			}
		}
	}

	// AllowGuestLoginGroups
	if !plan.AllowGuestLoginGroups.IsNull() && len(plan.allowGuestLoginGroups) > 0 {
		for _, aglg := range plan.allowGuestLoginGroups {
			temp := cidaas.IAllowedGroups{
				GroupID: aglg.GroupID.ValueString(),
			}
			diags := aglg.Roles.ElementsAs(ctx, &temp.Roles, false)
			if diags.HasError() {
				return nil, diags
			}
			diags = aglg.DefaultRoles.ElementsAs(ctx, &temp.DefaultRoles, false)
			if diags.HasError() {
				return nil, diags
			}
			app.AllowGuestLoginGroups = append(app.AllowGuestLoginGroups, temp)
		}
	}

	// LoginSpi
	if plan.loginSpi != nil {
		loginSpi := &cidaas.ILoginSPI{}
		if !plan.loginSpi.OauthClientID.IsNull() {
			loginSpi.OauthClientID = plan.loginSpi.OauthClientID.ValueString()
		}
		if !plan.loginSpi.SpiURL.IsNull() {
			loginSpi.SpiURL = plan.loginSpi.SpiURL.ValueString()
		}
		app.LoginSpi = loginSpi
	}
	// GroupSelection
	if plan.groupSelection != nil {
		groupSelection := &cidaas.IGroupSelection{}
		if !plan.groupSelection.AlwaysShowGroupSelection.IsNull() && !plan.groupSelection.AlwaysShowGroupSelection.IsUnknown() {
			groupSelection.AlwaysShowGroupSelection = plan.groupSelection.AlwaysShowGroupSelection.ValueBoolPointer()
		}
		if !plan.groupSelection.SelectableGroups.IsNull() && !plan.groupSelection.SelectableGroups.IsUnknown() {
			selectableGroups, diags := getSetAsStrings(ctx, plan.groupSelection.SelectableGroups)
			if diags.HasError() {
				return nil, diags
			}
			groupSelection.SelectableGroups = selectableGroups
		}

		if !plan.groupSelection.SelectableGroupTypes.IsNull() && !plan.groupSelection.SelectableGroupTypes.IsUnknown() {
			selectableGroupTypes, diags := getSetAsStrings(ctx, plan.groupSelection.SelectableGroupTypes)
			if diags.HasError() {
				return nil, diags
			}
			groupSelection.SelectableGroupTypes = selectableGroupTypes
		}
		app.GroupSelection = groupSelection
	}
	// Mfa
	if plan.mfa != nil {
		mfa := &cidaas.IMfaOption{}
		if !plan.mfa.Setting.IsNull() && !plan.mfa.Setting.IsUnknown() {
			mfa.Setting = plan.mfa.Setting.ValueString()
		}
		if !plan.mfa.TimeIntervalInSeconds.IsNull() && !plan.mfa.TimeIntervalInSeconds.IsUnknown() {
			mfa.TimeIntervalInSeconds = plan.mfa.TimeIntervalInSeconds.ValueInt64Pointer()
		}
		if !plan.mfa.AllowedMethods.IsNull() && !plan.mfa.AllowedMethods.IsUnknown() {
			allowedMethods, diags := getSetAsStrings(ctx, plan.mfa.AllowedMethods)
			if diags.HasError() {
				return nil, diags
			}
			mfa.AllowedMethods = allowedMethods
		}
		app.Mfa = mfa
	}
	// MobileSettings
	if plan.mobileSettings != nil {
		mobileSettings := &cidaas.IAppMobileSettings{}
		if !plan.mobileSettings.TeamID.IsNull() && !plan.mobileSettings.TeamID.IsUnknown() {
			mobileSettings.TeamID = plan.mobileSettings.TeamID.ValueString()
		}
		if !plan.mobileSettings.BundleID.IsNull() && !plan.mobileSettings.BundleID.IsUnknown() {
			mobileSettings.BundleID = plan.mobileSettings.BundleID.ValueString()
		}
		if !plan.mobileSettings.PackageName.IsNull() && !plan.mobileSettings.PackageName.IsUnknown() {
			mobileSettings.PackageName = plan.mobileSettings.PackageName.ValueString()
		}
		if !plan.mobileSettings.KeyHash.IsNull() && !plan.mobileSettings.KeyHash.IsUnknown() {
			mobileSettings.KeyHash = plan.mobileSettings.KeyHash.ValueString()
		}
		app.MobileSettings = mobileSettings
	}
	// ApplicationMetaData
	if len(plan.ApplicationMetaData.Elements()) > 0 {
		diags.Append(plan.ApplicationMetaData.ElementsAs(ctx, &app.ApplicationMetaData, false)...)
	}
	// SuggestVerificationMethods
	if plan.suggestVerificationMethods != nil {
		svm := &cidaas.SuggestVerificationMethods{}
		if !plan.suggestVerificationMethods.MandatoryConfig.IsNull() && !plan.suggestVerificationMethods.MandatoryConfig.IsUnknown() {
			mf := &MandatoryConfig{}
			diags = plan.suggestVerificationMethods.MandatoryConfig.As(ctx, mf, basetypes.ObjectAsOptions{})
			diags.Append(assignSetValues(ctx, mf.Methods, &svm.MandatoryConfig.Methods)...)
			svm.MandatoryConfig.Range = mf.Range.ValueString()
			svm.MandatoryConfig.SkipUntil = mf.SkipUntil.ValueString()
		}
		if !plan.suggestVerificationMethods.OptionalConfig.IsNull() && !plan.suggestVerificationMethods.OptionalConfig.IsUnknown() {
			of := &OptionalConfig{}
			diags = plan.suggestVerificationMethods.OptionalConfig.As(ctx, of, basetypes.ObjectAsOptions{})
			diags.Append(assignSetValues(ctx, of.Methods, &svm.OptionalConfig.Methods)...)
		}
		svm.SkipDurationInDays = plan.suggestVerificationMethods.SkipDurationInDays.ValueInt32()
		app.SuggestVerificationMethods = svm
	}
	// GroupRoleRestriction
	if plan.groupRoleRestriction != nil {
		grr := &cidaas.GroupRoleRestriction{}
		if !plan.groupRoleRestriction.Filters.IsNull() && !plan.groupRoleRestriction.Filters.IsUnknown() {
			filters := make([]GroupRoleFilters, 0, len(plan.groupRoleRestriction.Filters.Elements()))
			diags = plan.groupRoleRestriction.Filters.ElementsAs(ctx, &filters, false)
			target := []cidaas.GroupRoleFilters{}
			for _, f := range filters {
				rf := &cidaas.RoleFilter{}
				if !f.RoleFilter.IsNull() && !f.RoleFilter.IsUnknown() {
					tfrf := &RoleFilter{}
					diags = f.RoleFilter.As(ctx, tfrf, basetypes.ObjectAsOptions{})
					diags.Append(assignSetValues(ctx, tfrf.Roles, &rf.Roles)...)
					rf.MatchCondition = tfrf.MatchCondition.ValueString()
				}
				target = append(target, cidaas.GroupRoleFilters{
					GroupID:    f.GroupID.ValueString(),
					RoleFilter: *rf,
				})
				grr.Filters = target
			}
		}
		grr.MatchCondition = plan.groupRoleRestriction.MatchCondition.ValueString()
		app.GroupRoleRestriction = grr
	}
	if plan.basicSettings != nil {
		basicSettings := &cidaas.BasicSettings{}
		if !plan.RedirectURIS.IsNull() {
			diags.Append(assignSetValues(ctx, plan.RedirectURIS, &basicSettings.RedirectURIs)...)
		}
		if !plan.AllowedLogoutUrls.IsNull() {
			diags.Append(assignSetValues(ctx, plan.AllowedLogoutUrls, &basicSettings.AllowedLogoutUrls)...)
		}
		if !plan.AllowedScopes.IsNull() {
			diags.Append(assignSetValues(ctx, plan.AllowedScopes, &basicSettings.AllowedScopes)...)
		}
		if !plan.basicSettings.ClientSecrets.IsNull() {
			clientSecrets := make([]ClientSecret, 0, len(plan.basicSettings.ClientSecrets.Elements()))
			diags = plan.basicSettings.ClientSecrets.ElementsAs(ctx, &clientSecrets, false)
			target := []cidaas.ClientSecret{}

			for _, value := range clientSecrets {
				target = append(target, cidaas.ClientSecret{
					ClientSecret:          value.ClientSecret.ValueString(),
					ClientSecretExpiresAt: value.ClientSecretExpiresAt.ValueInt64(),
				})
			}
			basicSettings.ClientSecrets = target
		}
		app.BasicSettings = basicSettings
	}
	return &app, diags
}

func updateStateModel(res cidaas.AppResponse, state, config *AppConfig, operation int) diag.Diagnostics { //nolint:gocognit
	var diags diag.Diagnostics

	// computed always
	state.ID = util.StringValueOrNull(&res.Data.ID)
	state.ClientID = util.StringValueOrNull(&res.Data.ClientID)
	state.ClientSecret = util.StringValueOrNull(&res.Data.ClientSecret)
	// list ends here

	diags.Append(StringValueOrNullWithPlanValue(&res.Data.CompanyName, &config.CompanyName, &state.CompanyName, "company_name", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.CompanyWebsite, &config.CompanyWebsite, &state.CompanyWebsite, "company_website", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.ClientType, &config.ClientType, &state.ClientType, "client_type", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.CompanyAddress, &config.CompanyAddress, &state.CompanyAddress, "company_address", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.AccentColor, &config.AccentColor, &state.AccentColor, "accent_color", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.PrimaryColor, &config.PrimaryColor, &state.PrimaryColor, "primary_color", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.MediaType, &config.MediaType, &state.MediaType, "media_type", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.HostedPageGroup, &config.HostedPageGroup, &state.HostedPageGroup, "hosted_page_group", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.TemplateGroupID, &config.TemplateGroupID, &state.TemplateGroupID, "template_group_id", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.LogoAlign, &config.LogoAlign, &state.LogoAlign, "logo_align", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.Webfinger, &config.Webfinger, &state.Webfinger, "webfinger", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.ContentAlign, &config.ContentAlign, &state.ContentAlign, "content_align", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.ClientName, &config.ClientName, &state.ClientName, "client_name", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.ClientDisplayName, &config.ClientDisplayName, &state.ClientDisplayName, "client_display_name", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.PolicyURI, &config.PolicyURI, &state.PolicyURI, "policy_uri", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.TosURI, &config.TosURI, &state.TosURI, "tos_uri", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.ImprintURI, &config.ImprintURI, &state.ImprintURI, "imprint_uri", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.TokenEndpointAuthMethod, &config.TokenEndpointAuthMethod, &state.TokenEndpointAuthMethod, "token_endpoint_auth_method", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.TokenEndpointAuthSigningAlg, &config.TokenEndpointAuthSigningAlg, &state.TokenEndpointAuthSigningAlg, "token_endpoint_auth_signing_alg", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.CaptchaRef, &config.CaptchaRef, &state.CaptchaRef, "captcha_ref", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.CommunicationMediumVerification, &config.CommunicationMediumVerification, &state.CommunicationMediumVerification, "communication_medium_verification", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.BackchannelLogoutURI, &config.BackchannelLogoutURI, &state.BackchannelLogoutURI, "backchannel_logout_uri", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.LogoURI, &config.LogoURI, &state.LogoURI, "logo_uri", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.InitiateLoginURI, &config.InitiateLoginURI, &state.InitiateLoginURI, "initiate_login_uri", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.RegistrationClientURI, &config.RegistrationClientURI, &state.RegistrationClientURI, "registration_client_uri", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.RegistrationAccessToken, &config.RegistrationAccessToken, &state.RegistrationAccessToken, "registration_access_token", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.ClientURI, &config.ClientURI, &state.ClientURI, "client_uri", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.JwksURI, &config.JwksURI, &state.JwksURI, "jwks_uri", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.Jwks, &config.Jwks, &state.Jwks, "jwks", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.SectorIdentifierURI, &config.SectorIdentifierURI, &state.SectorIdentifierURI, "sector_identifier_uri", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.SubjectType, &config.SubjectType, &state.SubjectType, "subject_type", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.IDTokenSignedResponseAlg, &config.IDTokenSignedResponseAlg, &state.IDTokenSignedResponseAlg, "id_token_signed_response_alg", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.IDTokenEncryptedResponseAlg, &config.IDTokenEncryptedResponseAlg, &state.IDTokenEncryptedResponseAlg, "id_token_encrypted_response_alg", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.IDTokenEncryptedResponseEnc, &config.IDTokenEncryptedResponseEnc, &state.IDTokenEncryptedResponseEnc, "id_token_encrypted_response_enc", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.UserinfoSignedResponseAlg, &config.UserinfoSignedResponseAlg, &state.UserinfoSignedResponseAlg, "userinfo_signed_response_alg", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.UserinfoEncryptedResponseAlg, &config.UserinfoEncryptedResponseAlg, &state.UserinfoEncryptedResponseAlg, "userinfo_encrypted_response_alg", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.UserinfoEncryptedResponseEnc, &config.UserinfoEncryptedResponseEnc, &state.UserinfoEncryptedResponseEnc, "userinfo_encrypted_response_enc", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.RequestObjectSigningAlg, &config.RequestObjectSigningAlg, &state.RequestObjectSigningAlg, "request_object_signing_alg", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.RequestObjectEncryptionAlg, &config.RequestObjectEncryptionAlg, &state.RequestObjectEncryptionAlg, "request_object_encryption_alg", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.RequestObjectEncryptionEnc, &config.RequestObjectEncryptionEnc, &state.RequestObjectEncryptionEnc, "request_object_encryption_enc", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.Description, &config.Description, &state.Description, "description", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.ConsentPageGroup, &config.ConsentPageGroup, &state.ConsentPageGroup, "consent_page_group", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.PasswordPolicyRef, &config.PasswordPolicyRef, &state.PasswordPolicyRef, "password_policy_ref", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.BlockingMechanismRef, &config.BlockingMechanismRef, &state.BlockingMechanismRef, "blocking_mechanism_ref", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.Sub, &config.Sub, &state.Sub, "sub", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.Role, &config.Role, &state.Role, "role", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.MfaConfiguration, &config.MfaConfiguration, &state.MfaConfiguration, "mfa_configuration", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.BackgroundURI, &config.BackgroundURI, &state.BackgroundURI, "background_uri", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.VideoURL, &config.VideoURL, &state.VideoURL, "video_url", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.BotCaptchaRef, &config.BotCaptchaRef, &state.BotCaptchaRef, "bot_captcha_ref", operation)...)
	diags.Append(StringValueOrNullWithPlanValue(&res.Data.BotProvider, &config.BotProvider, &state.BotProvider, "bot_provider", operation)...)

	diags.Append(BoolValueOrNullWithPlanValue(res.Data.AllowGuestLogin, &state.AllowGuestLogin, "allow_guest_login", operation, false)...)
	diags.Append(BoolValueOrNullWithPlanValue(res.Data.EnableDeduplication, &state.EnableDeduplication, "enable_deduplication", operation, false)...)
	diags.Append(BoolValueOrNullWithPlanValue(res.Data.AutoLoginAfterRegister, &state.AutoLoginAfterRegister, "auto_login_after_register", operation, false)...)
	diags.Append(BoolValueOrNullWithPlanValue(res.Data.RegisterWithLoginInformation, &state.RegisterWithLoginInformation, "register_with_login_information", operation, false)...)
	diags.Append(BoolValueOrNullWithPlanValue(res.Data.FdsEnabled, &state.FdsEnabled, "fds_enabled", operation, true)...)
	diags.Append(BoolValueOrNullWithPlanValue(res.Data.IsHybridApp, &state.IsHybridApp, "is_hybrid_app", operation, false)...)
	diags.Append(BoolValueOrNullWithPlanValue(res.Data.Enabled, &state.Enabled, "enabled", operation, true)...)
	diags.Append(BoolValueOrNullWithPlanValue(res.Data.EmailVerificationRequired, &state.EmailVerificationRequired, "email_verification_required", operation, true)...)
	diags.Append(BoolValueOrNullWithPlanValue(res.Data.IsRememberMeSelected, &state.IsRememberMeSelected, "is_remember_me_selected", operation, true)...)
	diags.Append(BoolValueOrNullWithPlanValue(res.Data.AllowDisposableEmail, &state.AllowDisposableEmail, "allow_disposable_email", operation, false)...)
	diags.Append(BoolValueOrNullWithPlanValue(res.Data.ValidatePhoneNumber, &state.ValidatePhoneNumber, "validate_phone_number", operation, false)...)
	diags.Append(BoolValueOrNullWithPlanValue(res.Data.SmartMfa, &state.SmartMfa, "smart_mfa", operation, false)...)
	diags.Append(BoolValueOrNullWithPlanValue(res.Data.MobileNumberVerificationRequired, &state.MobileNumberVerificationRequired, "mobile_number_verification_required", operation, false)...)
	diags.Append(BoolValueOrNullWithPlanValue(res.Data.EnableBotDetection, &state.EnableBotDetection, "enable_bot_detection", operation, false)...)
	diags.Append(BoolValueOrNullWithPlanValue(res.Data.IsLoginSuccessPageEnabled, &state.IsLoginSuccessPageEnabled, "is_login_success_page_enabled", operation, false)...)
	diags.Append(BoolValueOrNullWithPlanValue(res.Data.IsRegisterSuccessPageEnabled, &state.IsRegisterSuccessPageEnabled, "is_register_success_page_enabled", operation, false)...)
	diags.Append(BoolValueOrNullWithPlanValue(res.Data.IsGroupLoginSelectionEnabled, &state.IsGroupLoginSelectionEnabled, "is_group_login_selection_enabled", operation, false)...)
	diags.Append(BoolValueOrNullWithPlanValue(res.Data.JweEnabled, &state.JweEnabled, "jwe_enabled", operation, false)...)
	diags.Append(BoolValueOrNullWithPlanValue(res.Data.Editable, &state.Editable, "editable", operation, true)...)
	diags.Append(BoolValueOrNullWithPlanValue(res.Data.UserConsent, &state.UserConsent, "user_consent", operation, false)...)
	diags.Append(BoolValueOrNullWithPlanValue(res.Data.EnableClassicalProvider, &state.EnableClassicalProvider, "enable_classical_provider", operation, true)...)
	diags.Append(BoolValueOrNullWithPlanValue(res.Data.EnablePasswordlessAuth, &state.EnablePasswordlessAuth, "enable_passwordless_auth", operation, true)...)
	diags.Append(BoolValueOrNullWithPlanValue(res.Data.AlwaysAskMfa, &state.AlwaysAskMfa, "always_ask_mfa", operation, false)...)

	diags.Append(BoolValueOrNullWithPlanValue(res.Data.RequireAuthTime, &state.RequireAuthTime, "require_auth_time", operation, false)...)
	diags.Append(BoolValueOrNullWithPlanValue(res.Data.EnableLoginSpi, &state.EnableLoginSpi, "enable_login_spi", operation, false)...)
	diags.Append(BoolValueOrNullWithPlanValue(res.Data.BackchannelLogoutSessionRequired, &state.BackchannelLogoutSessionRequired, "backchannel_logout_session_required", operation, false)...)

	state.DefaultMaxAge = util.Int64ValueOrNull(res.Data.DefaultMaxAge)
	state.TokenLifetimeInSeconds = util.Int64ValueOrNull(res.Data.TokenLifetimeInSeconds)
	state.IDTokenLifetimeInSeconds = util.Int64ValueOrNull(res.Data.IDTokenLifetimeInSeconds)
	state.RefreshTokenLifetimeInSeconds = util.Int64ValueOrNull(res.Data.RefreshTokenLifetimeInSeconds)

	diags.Append(SetValueOrNullWithPlanValue(res.Data.AllowedScopes, &state.AllowedScopes, "allowed_scopes", operation)...)
	diags.Append(SetValueOrNullWithPlanValue(res.Data.RedirectURIS, &state.RedirectURIS, "redirect_uris", operation)...)
	diags.Append(SetValueOrNullWithPlanValue(res.Data.AllowedLogoutUrls, &state.AllowedLogoutUrls, "allowed_logout_urls", operation)...)
	diags.Append(SetValueOrNullWithPlanValue(res.Data.AllowedWebOrigins, &state.AllowedWebOrigins, "allowed_web_origins", operation)...)
	diags.Append(SetValueOrNullWithPlanValue(res.Data.DefaultScopes, &state.DefaultScopes, "default_scopes", operation)...)
	diags.Append(SetValueOrNullWithPlanValue(res.Data.AllowedRoles, &state.AllowedRoles, "allowed_roles", operation)...)
	diags.Append(SetValueOrNullWithPlanValue(res.Data.ResponseTypes, &state.ResponseTypes, "response_types", operation)...)
	diags.Append(SetValueOrNullWithPlanValue(res.Data.GrantTypes, &state.GrantTypes, "grant_types", operation)...)
	diags.Append(SetValueOrNullWithPlanValue(res.Data.AllowLoginWith, &state.AllowLoginWith, "allow_login_with", operation)...)
	diags.Append(SetValueOrNullWithPlanValue(res.Data.PostLogoutRedirectUris, &state.PostLogoutRedirectUris, "post_logout_redirect_uris", operation)...)
	diags.Append(SetValueOrNullWithPlanValue(res.Data.GroupIDs, &state.GroupIDs, "group_ids", operation)...)
	diags.Append(SetValueOrNullWithPlanValue(res.Data.AllowedFields, &state.AllowedFields, "allowed_fields", operation)...)
	diags.Append(SetValueOrNullWithPlanValue(res.Data.GroupTypes, &state.GroupTypes, "group_types", operation)...)
	diags.Append(SetValueOrNullWithPlanValue(res.Data.AdditionalAccessTokenPayload, &state.AdditionalAccessTokenPayload, "additional_access_token_payload", operation)...)
	diags.Append(SetValueOrNullWithPlanValue(res.Data.RequiredFields, &state.RequiredFields, "required_fields", operation)...)
	diags.Append(SetValueOrNullWithPlanValue(res.Data.AllowedWebOrigins, &state.AllowedWebOrigins, "allowed_web_origins", operation)...)
	diags.Append(SetValueOrNullWithPlanValue(res.Data.Contacts, &state.Contacts, "contacts", operation)...)
	diags.Append(SetValueOrNullWithPlanValue(res.Data.WebMessageUris, &state.WebMessageUris, "web_message_uris", operation)...)
	diags.Append(SetValueOrNullWithPlanValue(res.Data.AllowedOrigins, &state.AllowedOrigins, "allowed_origins", operation)...)
	diags.Append(SetValueOrNullWithPlanValue(res.Data.LoginProviders, &state.LoginProviders, "login_providers", operation)...)
	diags.Append(SetValueOrNullWithPlanValue(res.Data.PendingScopes, &state.PendingScopes, "pending_scopes", operation)...)
	diags.Append(SetValueOrNullWithPlanValue(res.Data.AllowedMfa, &state.AllowedMfa, "allowed_mfa", operation)...)
	diags.Append(SetValueOrNullWithPlanValue(res.Data.DefaultRoles, &state.DefaultRoles, "default_roles", operation)...)
	diags.Append(SetValueOrNullWithPlanValue(res.Data.DefaultAcrValues, &state.DefaultAcrValues, "default_acr_values", operation)...)
	diags.Append(SetValueOrNullWithPlanValue(res.Data.SuggestMfa, &state.SuggestMfa, "suggest_mfa", operation)...)
	diags.Append(SetValueOrNullWithPlanValue(res.Data.CaptchaRefs, &state.CaptchaRefs, "captcha_refs", operation)...)
	diags.Append(SetValueOrNullWithPlanValue(res.Data.ConsentRefs, &state.ConsentRefs, "consent_refs", operation)...)
	diags.Append(SetValueOrNullWithPlanValue(res.Data.RequestUris, &state.RequestUris, "request_uris", operation)...)

	var spObjectValues []attr.Value
	spObjectType := types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"provider_name": types.StringType,
			"social_id":     types.StringType,
		},
	}

	if (operation == IMPORT || operation == READ) || len(state.CustomProviders.Elements()) > 0 {
		if len(res.Data.SocialProviders) > 0 {
			for _, sp := range res.Data.SocialProviders {
				providerName := sp.ProviderName
				socialID := sp.SocialID
				objValue := types.ObjectValueMust(
					spObjectType.AttrTypes,
					map[string]attr.Value{
						"provider_name": util.StringValueOrNull(&providerName),
						"social_id":     util.StringValueOrNull(&socialID),
					})
				spObjectValues = append(spObjectValues, objValue)
			}
			state.SocialProviders = types.ListValueMust(spObjectType, spObjectValues)
		}
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

	createProviderMetaObjectValues := func(providers []cidaas.IProviderMetadData) []attr.Value {
		var providerMetaObjectValues []attr.Value
		if len(providers) > 0 {
			for _, provider := range providers {
				providerName := provider.ProviderName
				displayName := provider.DisplayName
				logoURL := provider.LogoURL
				providerType := provider.Type
				objValue := types.ObjectValueMust(
					providerObjectType.AttrTypes,
					map[string]attr.Value{
						"provider_name":       util.StringValueOrNull(&providerName),
						"display_name":        util.StringValueOrNull(&displayName),
						"logo_url":            util.StringValueOrNull(&logoURL),
						"type":                util.StringValueOrNull(&providerType),
						"is_provider_visible": util.BoolValueOrNull(provider.IsProviderVisible),
						"domains":             util.SetValueOrNull(provider.Domains),
					})
				providerMetaObjectValues = append(providerMetaObjectValues, objValue)
			}
		}
		return providerMetaObjectValues
	}

	if len(res.Data.CustomProviders) > 0 && ((operation == IMPORT || operation == READ) || len(state.CustomProviders.Elements()) > 0) {
		state.CustomProviders = types.ListValueMust(providerObjectType, createProviderMetaObjectValues(res.Data.CustomProviders))
	}
	if len(res.Data.SamlProviders) > 0 && ((operation == IMPORT || operation == READ) || len(state.SamlProviders.Elements()) > 0) {
		state.SamlProviders = types.ListValueMust(providerObjectType, createProviderMetaObjectValues(res.Data.SamlProviders))
	}
	if len(res.Data.AdProviders) > 0 && ((operation == IMPORT || operation == READ) || len(state.AdProviders.Elements()) > 0) {
		state.AdProviders = types.ListValueMust(providerObjectType, createProviderMetaObjectValues(res.Data.AdProviders))
	}

	allowedGroupsObjectType := types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_id":      types.StringType,
			"roles":         types.SetType{ElemType: types.StringType},
			"default_roles": types.SetType{ElemType: types.StringType},
		},
	}

	createAllowedGroupsObjectValues := func(groups []cidaas.IAllowedGroups) []attr.Value {
		var allowedGroupObjectValues []attr.Value
		if len(groups) > 0 {
			for _, group := range groups {
				groupID := group.GroupID
				objValue := types.ObjectValueMust(
					allowedGroupsObjectType.AttrTypes,
					map[string]attr.Value{
						"group_id":      util.StringValueOrNull(&groupID),
						"roles":         util.SetValueOrNull(group.Roles),
						"default_roles": util.SetValueOrNull(group.DefaultRoles),
					})
				allowedGroupObjectValues = append(allowedGroupObjectValues, objValue)
			}
		}
		return allowedGroupObjectValues
	}

	if len(res.Data.AllowedGroups) > 0 && ((operation == IMPORT || operation == READ) || len(state.AllowedGroups.Elements()) > 0) {
		state.AllowedGroups = types.ListValueMust(allowedGroupsObjectType, createAllowedGroupsObjectValues(res.Data.AllowedGroups))
	}
	if len(res.Data.OperationsAllowedGroups) > 0 && ((operation == IMPORT || operation == READ) || len(state.OperationsAllowedGroups.Elements()) > 0) {
		state.OperationsAllowedGroups = types.ListValueMust(allowedGroupsObjectType, createAllowedGroupsObjectValues(res.Data.OperationsAllowedGroups))
	}

	if (operation == IMPORT || operation == READ) || len(state.AllowGuestLoginGroups.Elements()) > 0 {
		if len(res.Data.AllowGuestLoginGroups) > 0 {
			var allowedGroupObjectValues []attr.Value
			for _, group := range res.Data.AllowGuestLoginGroups {
				groupID := group.GroupID
				objValue := types.ObjectValueMust(
					allowedGroupsObjectType.AttrTypes,
					map[string]attr.Value{
						"group_id":      util.StringValueOrNull(&groupID),
						"roles":         util.SetValueOrNull(group.Roles),
						"default_roles": util.SetValueOrNull(group.DefaultRoles),
					})
				allowedGroupObjectValues = append(allowedGroupObjectValues, objValue)
			}
			state.AllowGuestLoginGroups = types.ListValueMust(allowedGroupsObjectType, allowedGroupObjectValues)
		}
	}

	metaData := map[string]attr.Value{}
	for key, value := range res.Data.ApplicationMetaData {
		val := value
		metaData[key] = util.StringValueOrNull(&val)
	}

	if len(res.Data.ApplicationMetaData) > 0 {
		appMetadata := types.MapValueMust(types.StringType, metaData)
		state.ApplicationMetaData = appMetadata
	}

	checkAndUpdateStringAttribute := func(attributeName string, stateValue types.String, plannedValue *types.String) {
		if !plannedValue.IsNull() {
			if !stateValue.IsNull() && stateValue != *plannedValue && operation != READ && operation != IMPORT {
				msg := fmt.Sprintf(
					"The attribute '%s' cannot be updated in Cidaas according to the provided configuration. "+
						"This might be due to the Cidaas API no longer supporting the attribute or because the attribute value "+
						"is computed based on other attributes or dependencies. The state is updated with the value '%s' "+
						"computed by Cidaas to prevent unnecessary errors in the provider. To prevent this warning, please update your configuration accordingly.",
					attributeName, plannedValue.ValueString())
				summaryMsg := fmt.Sprintf("Failed to Set %s", attributeName)
				diags.AddWarning(summaryMsg, msg)
			}
		} else if !stateValue.IsNull() {
			*plannedValue = stateValue
			if operation != READ && operation != IMPORT {
				msg := fmt.Sprintf(
					"The attribute '%s' cannot be updated in Cidaas according to the provided configuration. "+
						"This might be due to the Cidaas API no longer supporting the attribute or because the attribute value "+
						"is computed based on other attributes or dependencies. The state is updated with the planned value '%s' "+
						"to prevent unnecessary errors in the provider. To prevent this warning, please update your configuration accordingly.",
					attributeName, stateValue.ValueString())
				summaryMsg := fmt.Sprintf("Failed to Set %s", attributeName)
				diags.AddWarning(summaryMsg, msg)
			}
		}
	}

	checkAndUpdateIntAttribute := func(attributeName string, stateValue types.Int64, plannedValue *types.Int64) {
		if !plannedValue.IsNull() {
			if !stateValue.IsNull() && !stateValue.Equal(*plannedValue) && operation != READ && operation != IMPORT {
				msg := fmt.Sprintf(
					"The attribute '%s' cannot be updated in Cidaas according to the provided configuration. "+
						"This might be due to the Cidaas API no longer supporting the attribute or because the attribute value "+
						"is computed based on other attributes or dependencies. The state is updated with the value '%+v' "+
						"computed by Cidaas to prevent unnecessary errors in the provider. To prevent this warning, please update your configuration accordingly.",
					attributeName, plannedValue.ValueInt64())
				summaryMsg := fmt.Sprintf("Failed to Set %s", attributeName)
				diags.AddWarning(summaryMsg, msg)
			}
		} else if !stateValue.IsNull() {
			*plannedValue = stateValue
			if operation != READ && operation != IMPORT {
				msg := fmt.Sprintf(
					"The attribute '%s' cannot be updated in Cidaas according to the provided configuration. "+
						"This might be due to the Cidaas API no longer supporting the attribute or because the attribute value "+
						"is computed based on other attributes or dependencies. The state is updated with the planned value '%+v' "+
						"to prevent unnecessary errors in the provider. To prevent this warning, please update your configuration accordingly.",
					attributeName, stateValue.ValueInt64())
				summaryMsg := fmt.Sprintf("Failed to Set %s", attributeName)
				diags.AddWarning(summaryMsg, msg)
			}
		}
	}

	checkAndUpdateBoolAttribute := func(attributeName string, stateValue types.Bool, plannedValue *types.Bool) {
		if !plannedValue.IsNull() {
			if !stateValue.IsNull() && !stateValue.Equal(*plannedValue) && operation != READ && operation != IMPORT {
				msg := fmt.Sprintf(
					"The attribute '%s' cannot be updated in Cidaas according to the provided configuration. "+
						"This might be due to the Cidaas API no longer supporting the attribute or because the attribute value "+
						"is computed based on other attributes or dependencies. The state is updated with the value '%+v' "+
						"computed by Cidaas to prevent unnecessary errors in the provider. To prevent this warning, please update your configuration accordingly.",
					attributeName, plannedValue.ValueBool())
				summaryMsg := fmt.Sprintf("Failed to Set %s", attributeName)
				diags.AddWarning(summaryMsg, msg)
			}
		} else if !stateValue.IsNull() {
			*plannedValue = stateValue
			if operation != READ && operation != IMPORT {
				msg := fmt.Sprintf(
					"The attribute '%s' cannot be updated in Cidaas according to the provided configuration. "+
						"This might be due to the Cidaas API no longer supporting the attribute or because the attribute value "+
						"is computed based on other attributes or dependencies. The state is updated with the planned value '%+v' "+
						"to prevent unnecessary errors in the provider. To prevent this warning, please update your configuration accordingly.",
					attributeName, stateValue.ValueBool())
				summaryMsg := fmt.Sprintf("Failed to Set %s", attributeName)
				diags.AddWarning(summaryMsg, msg)
			}
		}
	}

	checkAndUpdateSetAttribute := func(attributeName string, stateValue types.Set, plannedValue *types.Set) {
		if !plannedValue.IsNull() {
			if !stateValue.IsNull() && !stateValue.Equal(*plannedValue) && operation != READ && operation != IMPORT {
				msg := fmt.Sprintf(
					"The attribute '%s' cannot be updated in Cidaas according to the provided configuration. "+
						"This might be due to the Cidaas API no longer supporting the attribute or because the attribute value "+
						"is computed based on other attributes or dependencies. The state is updated with the value '%+v' "+
						"computed by Cidaas to prevent unnecessary errors in the provider. To prevent this warning, please update your configuration accordingly.",
					attributeName, plannedValue.Elements())
				summaryMsg := fmt.Sprintf("Failed to Set %s", attributeName)
				diags.AddWarning(summaryMsg, msg)
			}
		} else if !stateValue.IsNull() {
			*plannedValue = stateValue
			if operation != READ && operation != IMPORT {
				msg := fmt.Sprintf(
					"The attribute '%s' cannot be updated in Cidaas according to the provided configuration. "+
						"This might be due to the Cidaas API no longer supporting the attribute or because the attribute value "+
						"is computed based on other attributes or dependencies. The state is updated with the planned value '%+v' "+
						"to prevent unnecessary errors in the provider. To prevent this warning, please update your configuration accordingly.",
					attributeName, stateValue.Elements())
				summaryMsg := fmt.Sprintf("Failed to Set %s", attributeName)
				diags.AddWarning(summaryMsg, msg)
			}
		}
	}

	if res.Data.LoginSpi != nil && (((operation == CREATE || operation == UPDATE) && !config.LoginSpi.IsNull()) ||
		operation == IMPORT || operation == READ) {
		oauthClientID := util.StringValueOrNull(&res.Data.LoginSpi.OauthClientID)
		spiURL := util.StringValueOrNull(&res.Data.LoginSpi.SpiURL)
		if !state.LoginSpi.IsNull() && !state.LoginSpi.IsUnknown() && state.loginSpi != nil {
			checkAndUpdateStringAttribute("login_spi.oauth_client_id", state.loginSpi.OauthClientID, &oauthClientID)
			checkAndUpdateStringAttribute("login_spi.spi_url", state.loginSpi.SpiURL, &spiURL)
		}

		loginSpi := types.ObjectValueMust(
			map[string]attr.Type{
				"oauth_client_id": types.StringType,
				"spi_url":         types.StringType,
			},
			map[string]attr.Value{
				"oauth_client_id": &oauthClientID,
				"spi_url":         &spiURL,
			})
		if !state.LoginSpi.IsNull() && !state.LoginSpi.IsUnknown() && !loginSpi.Equal(state.LoginSpi) && operation != READ && operation != IMPORT {
			msg := fmt.Sprintf(
				"The attribute 'login_spi' can not be updated in Cidaas according to the provided configuration."+
					"This could be due to the Cidaas API no longer supporting the attribute or because the attribute value might be computed based on other attributes or dependencies."+
					"The state is updated with the value '%+v' to prevent unnecessary errors in the provider.",
				loginSpi)
			diags.AddWarning("Failed to Set login_spi", msg)
		}
		state.LoginSpi = loginSpi
	}

	if res.Data.GroupSelection != nil &&
		(((operation == CREATE || operation == UPDATE) && !config.GroupSelection.IsNull()) ||
			operation == IMPORT || operation == READ) {
		alwaysShowGroupSelection := util.BoolValueOrNull(res.Data.GroupSelection.AlwaysShowGroupSelection)
		selectableGroups := util.SetValueOrNull(res.Data.GroupSelection.SelectableGroups)
		selectableGroupTypes := util.SetValueOrNull(res.Data.GroupSelection.SelectableGroupTypes)

		if !state.GroupSelection.IsNull() && !state.GroupSelection.IsUnknown() && state.groupSelection != nil {
			checkAndUpdateBoolAttribute("group_selection.always_show_group_selection", state.groupSelection.AlwaysShowGroupSelection, &alwaysShowGroupSelection)
			checkAndUpdateSetAttribute("group_selection.selectable_groups", state.groupSelection.SelectableGroups, &selectableGroups)
			checkAndUpdateSetAttribute("group_selection.selectable_group_types", state.groupSelection.SelectableGroupTypes, &selectableGroupTypes)
		}

		groupSelection := types.ObjectValueMust(
			map[string]attr.Type{
				"always_show_group_selection": types.BoolType,
				"selectable_groups":           types.SetType{ElemType: types.StringType},
				"selectable_group_types":      types.SetType{ElemType: types.StringType},
			},
			map[string]attr.Value{
				"always_show_group_selection": &alwaysShowGroupSelection,
				"selectable_groups":           &selectableGroups,
				"selectable_group_types":      &selectableGroupTypes,
			})
		if !state.GroupSelection.IsNull() && !state.GroupSelection.IsUnknown() && !groupSelection.Equal(state.GroupSelection) && operation != READ && operation != IMPORT {
			msg := fmt.Sprintf(
				"The attribute 'group_selection' can not be updated in Cidaas according to the provided configuration."+
					"This could be due to the Cidaas API no longer supporting the attribute or because the attribute value might be computed based on other attributes or dependencies."+
					"The state is updated with the value '%+v' to prevent unnecessary errors in the provider.",
				groupSelection)
			diags.AddWarning("Failed to Set group_selection", msg)
		}
		state.GroupSelection = groupSelection
	}

	if res.Data.Mfa != nil {
		setting := util.StringValueOrNull(&res.Data.Mfa.Setting)
		timeInterval := util.Int64ValueOrNull(res.Data.Mfa.TimeIntervalInSeconds)
		allowedMethods := util.SetValueOrNull(res.Data.Mfa.AllowedMethods)

		if !state.Mfa.IsNull() && !state.Mfa.IsUnknown() && state.mfa != nil {
			checkAndUpdateStringAttribute("mfa.setting", state.mfa.Setting, &setting)
			checkAndUpdateIntAttribute("mfa.time_interval_in_seconds", state.mfa.TimeIntervalInSeconds, &timeInterval)
			checkAndUpdateSetAttribute("mfa.allowed_methods", state.mfa.AllowedMethods, &allowedMethods)
		}

		mfa := types.ObjectValueMust(
			map[string]attr.Type{
				"setting":                  types.StringType,
				"time_interval_in_seconds": types.Int64Type,
				"allowed_methods":          types.SetType{ElemType: types.StringType},
			},
			map[string]attr.Value{
				"setting":                  &setting,
				"time_interval_in_seconds": &timeInterval,
				"allowed_methods":          &allowedMethods,
			},
		)
		if !state.Mfa.IsNull() && !state.Mfa.IsUnknown() && !mfa.Equal(state.Mfa) && operation != READ && operation != IMPORT {
			msg := fmt.Sprintf(
				"The attribute 'mfa' can not be updated in Cidaas according to the provided configuration."+
					"This could be due to the Cidaas API no longer supporting the attribute or because the attribute value might be computed based on other attributes or dependencies."+
					"The state is updated with the planned value '%+v' to prevent unnecessary errors in the provider.",
				mfa)
			diags.AddWarning("Failed to Set mfa", msg)
		}
		state.Mfa = mfa
	}

	if res.Data.MobileSettings != nil && (((operation == CREATE || operation == UPDATE) && !config.MobileSettings.IsNull()) ||
		operation == IMPORT || operation == READ) {
		teamID := util.StringValueOrNull(&res.Data.MobileSettings.TeamID)
		bundleID := util.StringValueOrNull(&res.Data.MobileSettings.BundleID)
		packageName := util.StringValueOrNull(&res.Data.MobileSettings.PackageName)
		keyHash := util.StringValueOrNull(&res.Data.MobileSettings.KeyHash)

		if !state.MobileSettings.IsNull() && !state.MobileSettings.IsUnknown() && state.mobileSettings != nil {
			checkAndUpdateStringAttribute("mobile_settings.team_id", state.mobileSettings.TeamID, &teamID)
			checkAndUpdateStringAttribute("mobile_settings.bundle_id", state.mobileSettings.BundleID, &bundleID)
			checkAndUpdateStringAttribute("mobile_settings.package_name", state.mobileSettings.PackageName, &packageName)
			checkAndUpdateStringAttribute("mobile_settings.key_hash", state.mobileSettings.KeyHash, &keyHash)
		}
		mobileSettings := types.ObjectValueMust(
			map[string]attr.Type{
				"team_id":      types.StringType,
				"bundle_id":    types.StringType,
				"package_name": types.StringType,
				"key_hash":     types.StringType,
			},
			map[string]attr.Value{
				"team_id":      util.StringValueOrNull(&res.Data.MobileSettings.TeamID),
				"bundle_id":    util.StringValueOrNull(&res.Data.MobileSettings.BundleID),
				"package_name": util.StringValueOrNull(&res.Data.MobileSettings.PackageName),
				"key_hash":     util.StringValueOrNull(&res.Data.MobileSettings.KeyHash),
			})
		if !state.MobileSettings.IsNull() && !state.MobileSettings.IsUnknown() && !mobileSettings.Equal(state.MobileSettings) && operation != READ && operation != IMPORT {
			msg := fmt.Sprintf(
				"The attribute 'mobile_settings' can not be updated in Cidaas according to the provided configuration."+
					"This could be due to the Cidaas API no longer supporting the attribute or because the attribute value might be computed based on other attributes or dependencies."+
					"The state is updated with the planned value '%+v' to prevent unnecessary errors in the provider. The computed value by cidaas is '%+v'. To prevent this warning, please update your configuration accordingly.",
				state.MobileSettings, mobileSettings)
			diags.AddWarning("Failed to Set mobile_settings", msg)
		}
		state.MobileSettings = mobileSettings
	}

	if res.Data.SuggestVerificationMethods != nil {
		skipDurationInDays := types.Int32Value(res.Data.SuggestVerificationMethods.SkipDurationInDays)
		skipUntil := util.StringValueOrNull(&res.Data.SuggestVerificationMethods.MandatoryConfig.SkipUntil)
		mandatatoryRange := util.StringValueOrNull(&res.Data.SuggestVerificationMethods.MandatoryConfig.Range)
		mandatoryMethods := util.SetValueOrNull(res.Data.SuggestVerificationMethods.MandatoryConfig.Methods)
		optionalMethods := util.SetValueOrNull(res.Data.SuggestVerificationMethods.OptionalConfig.Methods)

		mandateConfigType := map[string]attr.Type{
			"methods":    types.SetType{ElemType: types.StringType},
			"range":      types.StringType,
			"skip_until": types.StringType,
		}
		mandatoryConfig := types.ObjectValueMust(
			mandateConfigType,
			map[string]attr.Value{
				"methods":    mandatoryMethods,
				"range":      mandatatoryRange,
				"skip_until": skipUntil,
			},
		)

		optionalConfigType := map[string]attr.Type{
			"methods": types.SetType{ElemType: types.StringType},
		}
		optionalConfig := types.ObjectValueMust(
			optionalConfigType,
			map[string]attr.Value{
				"methods": optionalMethods,
			},
		)

		obj := types.ObjectValueMust(
			map[string]attr.Type{
				"mandatory_config":      types.ObjectType{AttrTypes: mandateConfigType},
				"optional_config":       types.ObjectType{AttrTypes: optionalConfigType},
				"skip_duration_in_days": types.Int32Type,
			},
			map[string]attr.Value{
				"mandatory_config":      mandatoryConfig,
				"optional_config":       optionalConfig,
				"skip_duration_in_days": skipDurationInDays,
			},
		)
		state.SuggestVerificationMethods = obj
	}

	if res.Data.GroupRoleRestriction != nil && (((operation == CREATE || operation == UPDATE) && !config.GroupRoleRestriction.IsNull()) ||
		operation == IMPORT || operation == READ) {
		roleFilterType := map[string]attr.Type{
			"match_condition": types.StringType,
			"roles":           types.SetType{ElemType: types.StringType},
		}
		filterType := map[string]attr.Type{
			"group_id":    types.StringType,
			"role_filter": types.ObjectType{AttrTypes: roleFilterType},
		}

		filterObjectType := types.ObjectType{
			AttrTypes: filterType,
		}
		if !(operation == READ && config.GroupRoleRestriction.IsNull()) {
			var filters basetypes.ListValue
			var filterObjectValues []attr.Value

			parentMatchCondition := util.StringValueOrNull(&res.Data.GroupRoleRestriction.MatchCondition)
			if len(res.Data.GroupRoleRestriction.Filters) > 0 {
				for _, grr := range res.Data.GroupRoleRestriction.Filters {
					groupID := grr.GroupID
					matchCondition := grr.RoleFilter.MatchCondition
					roles := grr.RoleFilter.Roles
					objValue := types.ObjectValueMust(
						filterType,
						map[string]attr.Value{
							"group_id": util.StringValueOrNull(&groupID),
							"role_filter": types.ObjectValueMust(
								roleFilterType,
								map[string]attr.Value{
									"match_condition": util.StringValueOrNull(&matchCondition),
									"roles":           util.SetValueOrNull(roles),
								},
							),
						})
					filterObjectValues = append(filterObjectValues, objValue)
				}
				filters = types.ListValueMust(filterObjectType, filterObjectValues)
			}

			obj := types.ObjectValueMust(
				map[string]attr.Type{
					"match_condition": types.StringType,
					"filters":         types.ListType{ElemType: types.ObjectType{AttrTypes: filterType}},
				},
				map[string]attr.Value{
					"match_condition": parentMatchCondition,
					"filters":         filters,
				},
			)
			state.GroupRoleRestriction = obj
		}
	}

	if res.Data.BasicSettings != nil {
		clientID := util.StringValueOrNull(&res.Data.ClientID)
		redirectURIs := util.SetValueOrNull(res.Data.BasicSettings.RedirectURIs)
		allowedLogoutUrls := util.SetValueOrNull(res.Data.BasicSettings.AllowedLogoutUrls)
		allowedScopes := util.SetValueOrNull(res.Data.BasicSettings.AllowedScopes)

		var clientSecrets []attr.Value
		for _, cs := range res.Data.BasicSettings.ClientSecrets {
			secret := cs.ClientSecret
			csExpiresAt := cs.ClientSecretExpiresAt
			clientSecret := types.ObjectValueMust(
				map[string]attr.Type{
					"client_secret":            types.StringType,
					"client_secret_expires_at": types.Int64Type,
				},
				map[string]attr.Value{
					"client_secret":            util.StringValueOrNull(&secret),
					"client_secret_expires_at": util.Int64ValueOrNull(&csExpiresAt),
				},
			)
			clientSecrets = append(clientSecrets, clientSecret)
		}

		basicSettings := types.ObjectValueMust(
			map[string]attr.Type{
				"client_id":           types.StringType,
				"redirect_uris":       types.SetType{ElemType: types.StringType},
				"allowed_logout_urls": types.SetType{ElemType: types.StringType},
				"allowed_scopes":      types.SetType{ElemType: types.StringType},
				"client_secrets": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
					"client_secret":            types.StringType,
					"client_secret_expires_at": types.Int64Type,
				}}},
			},
			map[string]attr.Value{
				"client_id":           clientID,
				"redirect_uris":       redirectURIs,
				"allowed_logout_urls": allowedLogoutUrls,
				"allowed_scopes":      allowedScopes,
				"client_secrets": types.ListValueMust(types.ObjectType{AttrTypes: map[string]attr.Type{
					"client_secret":            types.StringType,
					"client_secret_expires_at": types.Int64Type,
				}}, clientSecrets),
			},
		)
		state.BasicSettings = basicSettings
	}
	return diags
}
