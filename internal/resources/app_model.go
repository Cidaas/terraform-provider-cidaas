package resources

import (
	"context"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/cidaas"
	"github.com/Cidaas/terraform-provider-cidaas/helpers/util"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
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
	ApplicationType                  types.String `tfsdk:"application_type"`
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
	DisplayName  types.String `tfsdk:"display_name"`
}

type ProviderMetadData struct {
	LogoURL           types.String `tfsdk:"logo_url"`
	ProviderName      types.String `tfsdk:"provider_name"`
	DisplayName       types.String `tfsdk:"display_name"`
	Type              types.String `tfsdk:"type"`
	IsProviderVisible types.Bool   `tfsdk:"is_provider_visible"`
	Domains           types.Set    `tfsdk:"domains"`
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
	app := cidaas.AppModel{
		ClientType:                       plan.ClientType.ValueString(),
		AccentColor:                      plan.AccentColor.ValueString(),
		PrimaryColor:                     plan.PrimaryColor.ValueString(),
		MediaType:                        plan.MediaType.ValueString(),
		ContentAlign:                     plan.ContentAlign.ValueString(),
		EnableDeduplication:              plan.EnableDeduplication.ValueBool(),
		AutoLoginAfterRegister:           plan.AutoLoginAfterRegister.ValueBool(),
		EnablePasswordlessAuth:           plan.EnablePasswordlessAuth.ValueBool(),
		RegisterWithLoginInformation:     plan.RegisterWithLoginInformation.ValueBool(),
		AllowDisposableEmail:             plan.AllowDisposableEmail.ValueBool(),
		ValidatePhoneNumber:              plan.ValidatePhoneNumber.ValueBool(),
		FdsEnabled:                       plan.FdsEnabled.ValueBool(),
		HostedPageGroup:                  plan.HostedPageGroup.ValueString(),
		ClientName:                       plan.ClientName.ValueString(),
		ClientDisplayName:                plan.ClientDisplayName.ValueString(),
		CompanyName:                      plan.CompanyName.ValueString(),
		CompanyAddress:                   plan.CompanyAddress.ValueString(),
		CompanyWebsite:                   plan.CompanyWebsite.ValueString(),
		IsHybridApp:                      plan.IsHybridApp.ValueBool(),
		DefaultMaxAge:                    plan.DefaultMaxAge.ValueInt64(),
		TokenLifetimeInSeconds:           plan.TokenLifetimeInSeconds.ValueInt64(),
		IDTokenLifetimeInSeconds:         plan.IDTokenLifetimeInSeconds.ValueInt64(),
		RefreshTokenLifetimeInSeconds:    plan.RefreshTokenLifetimeInSeconds.ValueInt64(),
		TemplateGroupID:                  plan.TemplateGroupID.ValueString(),
		ClientID:                         plan.ClientID.ValueString(),
		ClientSecret:                     plan.ClientSecret.ValueString(),
		PolicyURI:                        plan.PolicyURI.ValueString(),
		TosURI:                           plan.TosURI.ValueString(),
		ImprintURI:                       plan.ImprintURI.ValueString(),
		TokenEndpointAuthMethod:          plan.TokenEndpointAuthMethod.ValueString(),
		TokenEndpointAuthSigningAlg:      plan.TokenEndpointAuthSigningAlg.ValueString(),
		Editable:                         plan.Editable.ValueBool(),
		AppOwner:                         plan.AppOwner.ValueString(),
		JweEnabled:                       plan.JweEnabled.ValueBool(),
		UserConsent:                      plan.UserConsent.ValueBool(),
		Deleted:                          plan.Deleted.ValueBool(),
		Enabled:                          plan.Enabled.ValueBool(),
		AlwaysAskMfa:                     plan.AlwaysAskMfa.ValueBool(),
		SmartMfa:                         plan.SmartMfa.ValueBool(),
		CaptchaRef:                       plan.CaptchaRef.ValueString(),
		CommunicationMediumVerification:  plan.CommunicationMediumVerification.ValueString(),
		EmailVerificationRequired:        plan.EmailVerificationRequired.ValueBool(),
		MobileNumberVerificationRequired: plan.MobileNumberVerificationRequired.ValueBool(),
		EnableClassicalProvider:          plan.EnableClassicalProvider.ValueBool(),
		IsRememberMeSelected:             plan.IsRememberMeSelected.ValueBool(),
		EnableBotDetection:               plan.EnableBotDetection.ValueBool(),
		BotProvider:                      plan.BotProvider.ValueString(),
		IsLoginSuccessPageEnabled:        plan.IsLoginSuccessPageEnabled.ValueBool(),
		IsRegisterSuccessPageEnabled:     plan.IsRegisterSuccessPageEnabled.ValueBool(),
		AdminClient:                      plan.AdminClient.ValueBool(),
		IsGroupLoginSelectionEnabled:     plan.IsGroupLoginSelectionEnabled.ValueBool(),
		BackchannelLogoutURI:             plan.BackchannelLogoutURI.ValueString(),
		LogoAlign:                        plan.LogoAlign.ValueString(),
		Webfinger:                        plan.Webfinger.ValueString(),
		ApplicationType:                  plan.ApplicationType.ValueString(),
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
		PasswordPolicyRef:                plan.PasswordPolicyRef.ValueString(),
		BlockingMechanismRef:             plan.BlockingMechanismRef.ValueString(),
		Sub:                              plan.Sub.ValueString(),
		Role:                             plan.Role.ValueString(),
		MfaConfiguration:                 plan.MfaConfiguration.ValueString(),
		AllowGuestLogin:                  plan.AllowGuestLogin.ValueBool(),
		BackgroundURI:                    plan.BackgroundURI.ValueString(),
		VideoURL:                         plan.VideoURL.ValueString(),
		BotCaptchaRef:                    plan.BotCaptchaRef.ValueString(),
		CreatedTime:                      plan.CreatedAt.ValueString(),
		UpdatedTime:                      plan.UpdatedAt.ValueString(),
	}

	diags := updateSetValues(ctx, plan.PostLogoutRedirectUris, &app.PostLogoutRedirectUris)
	if diags.HasError() {
		return nil, diags
	}
	diags = updateSetValues(ctx, plan.GroupIDs, &app.GroupIDs)
	if diags.HasError() {
		return nil, diags
	}
	diags = updateSetValues(ctx, plan.AllowedMfa, &app.AllowedMfa)
	if diags.HasError() {
		return nil, diags
	}
	diags = updateSetValues(ctx, plan.AllowedFields, &app.AllowedFields)
	if diags.HasError() {
		return nil, diags
	}
	diags = updateSetValues(ctx, plan.WebMessageUris, &app.WebMessageUris)
	if diags.HasError() {
		return nil, diags
	}
	diags = updateSetValues(ctx, plan.DefaultAcrValues, &app.DefaultAcrValues)
	if diags.HasError() {
		return nil, diags
	}
	diags = updateSetValues(ctx, plan.GroupTypes, &app.GroupTypes)
	if diags.HasError() {
		return nil, diags
	}
	diags = updateSetValues(ctx, plan.AllowedRoles, &app.AllowedRoles)
	if diags.HasError() {
		return nil, diags
	}
	diags = updateSetValues(ctx, plan.DefaultRoles, &app.DefaultRoles)
	if diags.HasError() {
		return nil, diags
	}
	diags = updateSetValues(ctx, plan.SuggestMfa, &app.SuggestMfa)
	if diags.HasError() {
		return nil, diags
	}
	diags = updateSetValues(ctx, plan.CaptchaRefs, &app.CaptchaRefs)
	if diags.HasError() {
		return nil, diags
	}
	diags = updateSetValues(ctx, plan.ConsentRefs, &app.ConsentRefs)
	if diags.HasError() {
		return nil, diags
	}
	diags = updateSetValues(ctx, plan.DefaultScopes, &app.DefaultScopes)
	if diags.HasError() {
		return nil, diags
	}
	diags = updateSetValues(ctx, plan.PendingScopes, &app.PendingScopes)
	if diags.HasError() {
		return nil, diags
	}
	diags = updateSetValues(ctx, plan.AllowLoginWith, &app.AllowLoginWith)
	if diags.HasError() {
		return nil, diags
	}
	diags = updateSetValues(ctx, plan.RedirectURIS, &app.RedirectURIS)
	if diags.HasError() {
		return nil, diags
	}
	diags = updateSetValues(ctx, plan.AllowedLogoutUrls, &app.AllowedLogoutUrls)
	if diags.HasError() {
		return nil, diags
	}
	diags = updateSetValues(ctx, plan.AllowedScopes, &app.AllowedScopes)
	if diags.HasError() {
		return nil, diags
	}
	diags = updateSetValues(ctx, plan.ResponseTypes, &app.ResponseTypes)
	if diags.HasError() {
		return nil, diags
	}
	diags = updateSetValues(ctx, plan.GrantTypes, &app.GrantTypes)
	if diags.HasError() {
		return nil, diags
	}
	diags = updateSetValues(ctx, plan.LoginProviders, &app.LoginProviders)
	if diags.HasError() {
		return nil, diags
	}
	diags = updateSetValues(ctx, plan.AdditionalAccessTokenPayload, &app.AdditionalAccessTokenPayload)
	if diags.HasError() {
		return nil, diags
	}
	diags = updateSetValues(ctx, plan.RequiredFields, &app.RequiredFields)
	if diags.HasError() {
		return nil, diags
	}
	diags = updateSetValues(ctx, plan.AllowedWebOrigins, &app.AllowedWebOrigins)
	if diags.HasError() {
		return nil, diags
	}
	diags = updateSetValues(ctx, plan.AllowedOrigins, &app.AllowedOrigins)
	if diags.HasError() {
		return nil, diags
	}
	diags = updateSetValues(ctx, plan.Contacts, &app.Contacts)
	if diags.HasError() {
		return nil, diags
	}
	diags = updateSetValues(ctx, plan.RequestUris, &app.RequestUris)
	if diags.HasError() {
		return nil, diags
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

	for _, sp := range plan.socialProviders {
		app.SocialProviders = append(app.SocialProviders, &cidaas.ISocialProviderData{
			ProviderName: sp.ProviderName.ValueString(),
			SocialID:     sp.SocialID.ValueString(),
			DisplayName:  sp.DisplayName.ValueString(),
		})
	}
	for _, cp := range plan.customProviders {
		temp := &cidaas.IProviderMetadData{
			ProviderName:      cp.ProviderName.ValueString(),
			Type:              cp.Type.ValueString(),
			DisplayName:       cp.DisplayName.ValueString(),
			LogoURL:           cp.LogoURL.ValueString(),
			IsProviderVisible: cp.IsProviderVisible.ValueBool(),
		}
		diags := cp.Domains.ElementsAs(ctx, temp.Domains, false)
		if diags.HasError() {
			return nil, diags
		}
		app.CustomProviders = append(app.CustomProviders, temp)
	}
	for _, sp := range plan.samlProviders {
		temp := &cidaas.IProviderMetadData{
			ProviderName:      sp.ProviderName.ValueString(),
			Type:              sp.Type.ValueString(),
			DisplayName:       sp.DisplayName.ValueString(),
			LogoURL:           sp.LogoURL.ValueString(),
			IsProviderVisible: sp.IsProviderVisible.ValueBool(),
		}
		diags := sp.Domains.ElementsAs(ctx, temp.Domains, false)
		if diags.HasError() {
			return nil, diags
		}
		app.SamlProviders = append(app.SamlProviders, temp)
	}

	for _, ap := range plan.adProviders {
		temp := &cidaas.IProviderMetadData{
			ProviderName:      ap.ProviderName.ValueString(),
			Type:              ap.Type.ValueString(),
			DisplayName:       ap.DisplayName.ValueString(),
			LogoURL:           ap.LogoURL.ValueString(),
			IsProviderVisible: ap.IsProviderVisible.ValueBool(),
		}
		diags := ap.Domains.ElementsAs(ctx, temp.Domains, false)
		if diags.HasError() {
			return nil, diags
		}
		app.AdProviders = append(app.AdProviders, temp)
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
	return &app, nil
}

func updateSetValues(ctx context.Context, src basetypes.SetValue, dest *[]string) diag.Diagnostics {
	if len(src.Elements()) > 0 {
		return src.ElementsAs(ctx, dest, false)
	}
	return nil
}

func updateStateModel(ctx context.Context, res *cidaas.AppResponse, state *AppConfig) resource.ReadResponse {
	resp := resource.ReadResponse{}

	state.ID = util.StringValueOrNull(&res.Data.ID)
	state.ClientType = util.StringValueOrNull(&res.Data.ClientType)
	state.AccentColor = util.StringValueOrNull(&res.Data.AccentColor)
	state.PrimaryColor = util.StringValueOrNull(&res.Data.PrimaryColor)
	state.MediaType = util.StringValueOrNull(&res.Data.MediaType)
	state.ContentAlign = util.StringValueOrNull(&res.Data.ContentAlign)
	state.EnableDeduplication = types.BoolValue(res.Data.EnableDeduplication)
	state.AutoLoginAfterRegister = types.BoolValue(res.Data.AutoLoginAfterRegister)
	state.EnablePasswordlessAuth = types.BoolValue(res.Data.EnablePasswordlessAuth)
	state.RegisterWithLoginInformation = types.BoolValue(res.Data.RegisterWithLoginInformation)
	state.AllowDisposableEmail = types.BoolValue(res.Data.AllowDisposableEmail)
	state.ValidatePhoneNumber = types.BoolValue(res.Data.ValidatePhoneNumber)
	state.FdsEnabled = types.BoolValue(res.Data.FdsEnabled)
	state.HostedPageGroup = util.StringValueOrNull(&res.Data.HostedPageGroup)
	state.ClientName = util.StringValueOrNull(&res.Data.ClientName)
	state.ClientDisplayName = util.StringValueOrNull(&res.Data.ClientDisplayName)
	state.CompanyName = util.StringValueOrNull(&res.Data.CompanyName)
	state.CompanyAddress = util.StringValueOrNull(&res.Data.CompanyAddress)
	state.CompanyWebsite = util.StringValueOrNull(&res.Data.CompanyWebsite)
	state.IsHybridApp = types.BoolValue(res.Data.IsHybridApp)
	state.DefaultMaxAge = types.Int64Value(res.Data.DefaultMaxAge)
	state.TokenLifetimeInSeconds = types.Int64Value(res.Data.TokenLifetimeInSeconds)
	state.IDTokenLifetimeInSeconds = types.Int64Value(res.Data.IDTokenLifetimeInSeconds)
	state.RefreshTokenLifetimeInSeconds = types.Int64Value(res.Data.RefreshTokenLifetimeInSeconds)
	state.TemplateGroupID = util.StringValueOrNull(&res.Data.TemplateGroupID)
	state.ClientID = util.StringValueOrNull(&res.Data.ClientID)
	state.ClientSecret = util.StringValueOrNull(&res.Data.ClientSecret)
	state.PolicyURI = util.StringValueOrNull(&res.Data.PolicyURI)
	state.TosURI = util.StringValueOrNull(&res.Data.TosURI)
	state.ImprintURI = util.StringValueOrNull(&res.Data.ImprintURI)
	state.TokenEndpointAuthMethod = util.StringValueOrNull(&res.Data.TokenEndpointAuthMethod)
	state.TokenEndpointAuthSigningAlg = util.StringValueOrNull(&res.Data.TokenEndpointAuthSigningAlg)
	state.Editable = types.BoolValue(res.Data.Editable)
	state.AppOwner = util.StringValueOrNull(&res.Data.AppOwner)
	state.JweEnabled = types.BoolValue(res.Data.JweEnabled)
	state.UserConsent = types.BoolValue(res.Data.UserConsent)
	state.Deleted = types.BoolValue(res.Data.Deleted)
	state.Enabled = types.BoolValue(res.Data.Enabled)
	state.AlwaysAskMfa = types.BoolValue(res.Data.AlwaysAskMfa)
	state.SmartMfa = types.BoolValue(res.Data.SmartMfa)
	state.CaptchaRef = util.StringValueOrNull(&res.Data.CaptchaRef)
	state.CommunicationMediumVerification = util.StringValueOrNull(&res.Data.CommunicationMediumVerification)
	state.EmailVerificationRequired = types.BoolValue(res.Data.EmailVerificationRequired)
	state.MobileNumberVerificationRequired = types.BoolValue(res.Data.MobileNumberVerificationRequired)
	state.EnableClassicalProvider = types.BoolValue(res.Data.EnableClassicalProvider)
	state.IsRememberMeSelected = types.BoolValue(res.Data.IsRememberMeSelected)
	state.EnableBotDetection = types.BoolValue(res.Data.EnableBotDetection)
	state.BotProvider = util.StringValueOrNull(&res.Data.BotProvider)
	state.IsLoginSuccessPageEnabled = types.BoolValue(res.Data.IsLoginSuccessPageEnabled)
	state.IsRegisterSuccessPageEnabled = types.BoolValue(res.Data.IsRegisterSuccessPageEnabled)
	state.AdminClient = types.BoolValue(res.Data.AdminClient)
	state.IsGroupLoginSelectionEnabled = types.BoolValue(res.Data.IsGroupLoginSelectionEnabled)
	state.BackchannelLogoutURI = util.StringValueOrNull(&res.Data.BackchannelLogoutURI)
	state.LogoAlign = util.StringValueOrNull(&res.Data.LogoAlign)
	state.Webfinger = util.StringValueOrNull(&res.Data.Webfinger)
	state.ApplicationType = util.StringValueOrNull(&res.Data.ApplicationType)
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
	state.AllowGuestLogin = types.BoolValue(res.Data.AllowGuestLogin)
	state.BackgroundURI = util.StringValueOrNull(&res.Data.BackgroundURI)
	state.VideoURL = util.StringValueOrNull(&res.Data.VideoURL)
	state.BotCaptchaRef = util.StringValueOrNull(&res.Data.BotCaptchaRef)
	state.CreatedAt = util.StringValueOrNull(&res.Data.CreatedTime)
	state.UpdatedAt = util.StringValueOrNull(&res.Data.UpdatedTime)

	var d diag.Diagnostics
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
	resp.Diagnostics.Append(extractSetValues(ctx, &state.AllowLoginWith, res.Data.AllowLoginWith)...)
	resp.Diagnostics.Append(extractSetValues(ctx, &state.RedirectURIS, res.Data.RedirectURIS)...)
	resp.Diagnostics.Append(extractSetValues(ctx, &state.AllowedLogoutUrls, res.Data.AllowedLogoutUrls)...)
	resp.Diagnostics.Append(extractSetValues(ctx, &state.AllowedScopes, res.Data.AllowedScopes)...)
	resp.Diagnostics.Append(extractSetValues(ctx, &state.ResponseTypes, res.Data.ResponseTypes)...)
	resp.Diagnostics.Append(extractSetValues(ctx, &state.GrantTypes, res.Data.GrantTypes)...)
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

	var spObjectValues []attr.Value
	spObjectType := types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"provider_name": types.StringType,
			"social_id":     types.StringType,
			"display_name":  types.StringType,
		},
	}
	for _, sp := range res.Data.SocialProviders {
		objValue := types.ObjectValueMust(
			spObjectType.AttrTypes,
			map[string]attr.Value{
				"provider_name": util.StringValueOrNull(&sp.ProviderName),
				"social_id":     util.StringValueOrNull(&sp.SocialID),
				"display_name":  util.StringValueOrNull(&sp.DisplayName),
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
				"domains": types.SetValueMust(types.StringType, func() []attr.Value {
					var temp []attr.Value
					for _, role := range cp.Domains {
						temp = append(temp, util.StringValueOrNull(&role))
					}
					return temp
				}()),
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
				"domains": types.SetValueMust(types.StringType, func() []attr.Value {
					var temp []attr.Value
					for _, role := range saml.Domains {
						temp = append(temp, util.StringValueOrNull(&role))
					}
					return temp
				}()),
			})
		samlProviderMetaObjectValues = append(samlProviderMetaObjectValues, objValue)
	}

	state.SamlProviders, d = types.ListValueFrom(ctx, providerObjectType, samlProviderMetaObjectValues)
	resp.Diagnostics.Append(d...)
	if resp.Diagnostics.HasError() {
		return resp
	}

	var adProviderMetaObjectValues []attr.Value
	for _, ad := range res.Data.SamlProviders {
		objValue := types.ObjectValueMust(
			providerObjectType.AttrTypes,
			map[string]attr.Value{
				"provider_name":       util.StringValueOrNull(&ad.ProviderName),
				"display_name":        util.StringValueOrNull(&ad.DisplayName),
				"logo_url":            util.StringValueOrNull(&ad.LogoURL),
				"type":                util.StringValueOrNull(&ad.Type),
				"is_provider_visible": types.BoolValue(ad.IsProviderVisible),
				"domains": types.SetValueMust(types.StringType, func() []attr.Value {
					var temp []attr.Value
					for _, role := range ad.Domains {
						temp = append(temp, util.StringValueOrNull(&role))
					}
					return temp
				}()),
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
