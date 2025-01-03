package cidaas

import (
	"fmt"
	"net/http"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/util"
)

type AppResponse struct {
	Success bool     `json:"success,omitempty"`
	Status  int64    `json:"status,omitempty"`
	Data    AppModel `json:"data,omitempty"`
}

// to add
// suggestVerificationMethods
type AppModel struct {
	ID                               string                      `json:"_id,omitempty"`
	ClientType                       string                      `json:"client_type,omitempty"`
	AccentColor                      string                      `json:"accentColor,omitempty"`
	PrimaryColor                     string                      `json:"primaryColor,omitempty"`
	MediaType                        string                      `json:"mediaType,omitempty"`
	ContentAlign                     string                      `json:"contentAlign,omitempty"`
	AllowLoginWith                   []string                    `json:"allow_login_with,omitempty"`
	RedirectURIS                     []string                    `json:"redirect_uris,omitempty"`
	AllowedLogoutUrls                []string                    `json:"allowed_logout_urls,omitempty"`
	EnableDeduplication              *bool                       `json:"enable_deduplication"`
	AutoLoginAfterRegister           *bool                       `json:"auto_login_after_register,omitempty"`
	EnablePasswordlessAuth           *bool                       `json:"enable_passwordless_auth,omitempty"`
	RegisterWithLoginInformation     *bool                       `json:"register_with_login_information"`
	AllowDisposableEmail             *bool                       `json:"allow_disposable_email,omitempty"`
	ValidatePhoneNumber              *bool                       `json:"validate_phone_number,omitempty"`
	FdsEnabled                       *bool                       `json:"fds_enabled,omitempty"`
	HostedPageGroup                  string                      `json:"hosted_page_group,omitempty"`
	ClientName                       string                      `json:"client_name,omitempty"`
	ClientDisplayName                string                      `json:"client_display_name,omitempty"`
	CompanyName                      string                      `json:"company_name,omitempty"`
	CompanyAddress                   string                      `json:"company_address,omitempty"`
	CompanyWebsite                   string                      `json:"company_website,omitempty"`
	AllowedScopes                    []string                    `json:"allowed_scopes,omitempty"`
	ResponseTypes                    []string                    `json:"response_types,omitempty"`
	GrantTypes                       []string                    `json:"grant_types,omitempty"`
	LoginProviders                   []string                    `json:"login_providers,omitempty"`
	AdditionalAccessTokenPayload     []string                    `json:"additional_access_token_payload,omitempty"`
	RequiredFields                   []string                    `json:"required_fields,omitempty"`
	IsHybridApp                      *bool                       `json:"is_hybrid_app,omitempty"`
	AllowedWebOrigins                []string                    `json:"allowed_web_origins,omitempty"`
	AllowedOrigins                   []string                    `json:"allowed_origins,omitempty"`
	MobileSettings                   *IAppMobileSettings         `json:"mobile_settings,omitempty"`
	DefaultMaxAge                    *int64                      `json:"default_max_age,omitempty"`
	TokenLifetimeInSeconds           *int64                      `json:"token_lifetime_in_seconds,omitempty"`
	IDTokenLifetimeInSeconds         *int64                      `json:"id_token_lifetime_in_seconds,omitempty"`
	RefreshTokenLifetimeInSeconds    *int64                      `json:"refresh_token_lifetime_in_seconds,omitempty"`
	TemplateGroupID                  string                      `json:"template_group_id,omitempty"`
	ClientID                         string                      `json:"client_id,omitempty"`
	ClientSecret                     string                      `json:"client_secret,omitempty"`
	PolicyURI                        string                      `json:"policy_uri,omitempty"`
	TosURI                           string                      `json:"tos_uri,omitempty"`
	ImprintURI                       string                      `json:"imprint_uri,omitempty"`
	Contacts                         []string                    `json:"contacts,omitempty"`
	TokenEndpointAuthMethod          string                      `json:"token_endpoint_auth_method,omitempty"`
	TokenEndpointAuthSigningAlg      string                      `json:"token_endpoint_auth_signing_alg,omitempty"`
	DefaultAcrValues                 []string                    `json:"default_acr_values,omitempty"`
	Editable                         *bool                       `json:"editable,omitempty"`
	WebMessageUris                   []string                    `json:"web_message_uris,omitempty"`
	SocialProviders                  []ISocialProviderData       `json:"social_providers,omitempty"`
	CustomProviders                  []IProviderMetadData        `json:"custom_providers,omitempty"`
	SamlProviders                    []IProviderMetadData        `json:"saml_providers,omitempty"`
	AdProviders                      []IProviderMetadData        `json:"ad_providers,omitempty"`
	JweEnabled                       *bool                       `json:"jwe_enabled"`
	UserConsent                      *bool                       `json:"user_consent,omitempty"`
	AllowedGroups                    []IAllowedGroups            `json:"allowed_groups,omitempty"`
	OperationsAllowedGroups          []IAllowedGroups            `json:"operations_allowed_groups,omitempty"`
	Enabled                          *bool                       `json:"enabled"`
	AllowedFields                    []string                    `json:"allowed_fields,omitempty"`
	AlwaysAskMfa                     *bool                       `json:"always_ask_mfa,omitempty"`
	SmartMfa                         *bool                       `json:"smart_mfa,omitempty"`
	AllowedMfa                       []string                    `json:"allowed_mfa,omitempty"`
	CaptchaRef                       string                      `json:"captcha_ref,omitempty"`
	CaptchaRefs                      []string                    `json:"captcha_refs,omitempty"`
	ConsentRefs                      []string                    `json:"consent_refs"`
	CommunicationMediumVerification  string                      `json:"communication_medium_verification,omitempty"`
	EmailVerificationRequired        *bool                       `json:"email_verification_required,omitempty"`
	MobileNumberVerificationRequired *bool                       `json:"mobile_number_verification_required,omitempty"`
	AllowedRoles                     []string                    `json:"allowed_roles,omitempty"`
	DefaultRoles                     []string                    `json:"default_roles,omitempty"`
	EnableClassicalProvider          *bool                       `json:"enable_classical_provider,omitempty"`
	IsRememberMeSelected             *bool                       `json:"is_remember_me_selected"`
	EnableBotDetection               *bool                       `json:"enable_bot_detection,omitempty"`
	BotProvider                      string                      `json:"bot_provider,omitempty"`
	AllowGuestLoginGroups            []IAllowedGroups            `json:"allow_guest_login_groups,omitempty"`
	IsLoginSuccessPageEnabled        *bool                       `json:"is_login_success_page_enabled,omitempty"`
	IsRegisterSuccessPageEnabled     *bool                       `json:"is_register_success_page_enabled,omitempty"`
	GroupIDs                         []string                    `json:"groupIds,omitempty"`
	IsGroupLoginSelectionEnabled     *bool                       `json:"isGroupLoginSelectionEnabled,omitempty"`
	GroupSelection                   *IGroupSelection            `json:"groupSelection,omitempty"`
	GroupTypes                       []string                    `json:"groupTypes,omitempty"`
	BackchannelLogoutURI             string                      `json:"backchannel_logout_uri,omitempty"`
	PostLogoutRedirectUris           []string                    `json:"post_logout_redirect_uris,omitempty"`
	LogoAlign                        string                      `json:"logoAlign,omitempty"`
	Mfa                              *IMfaOption                 `json:"mfa,omitempty"`
	Webfinger                        string                      `json:"webfinger,omitempty"`
	ApplicationType                  string                      `json:"application_type,omitempty"`
	LogoURI                          string                      `json:"logo_uri,omitempty"`
	InitiateLoginURI                 string                      `json:"initiate_login_uri,omitempty"`
	RegistrationClientURI            string                      `json:"registration_client_uri,omitempty"`
	RegistrationAccessToken          string                      `json:"registration_access_token,omitempty"`
	ClientURI                        string                      `json:"client_uri,omitempty"`
	JwksURI                          string                      `json:"jwks_uri,omitempty"`
	Jwks                             string                      `json:"jwks,omitempty"`
	SectorIdentifierURI              string                      `json:"sector_identifier_uri,omitempty"`
	SubjectType                      string                      `json:"subject_type,omitempty"`
	IDTokenSignedResponseAlg         string                      `json:"id_token_signed_response_alg,omitempty"`
	IDTokenEncryptedResponseAlg      string                      `json:"id_token_encrypted_response_alg,omitempty"`
	IDTokenEncryptedResponseEnc      string                      `json:"id_token_encrypted_response_enc,omitempty"`
	UserinfoSignedResponseAlg        string                      `json:"userinfo_signed_response_alg,omitempty"`
	UserinfoEncryptedResponseAlg     string                      `json:"userinfo_encrypted_response_alg,omitempty"`
	UserinfoEncryptedResponseEnc     string                      `json:"userinfo_encrypted_response_enc,omitempty"`
	RequestObjectSigningAlg          string                      `json:"request_object_signing_alg,omitempty"`
	RequestObjectEncryptionAlg       string                      `json:"request_object_encryption_alg,omitempty"`
	RequestObjectEncryptionEnc       string                      `json:"request_object_encryption_enc,omitempty"`
	RequestUris                      []string                    `json:"request_uris,omitempty"`
	Description                      string                      `json:"description,omitempty"`
	DefaultScopes                    []string                    `json:"default_scopes,omitempty"`
	PendingScopes                    []string                    `json:"pending_scopes,omitempty"`
	ConsentPageGroup                 string                      `json:"consent_page_group,omitempty"`
	PasswordPolicyRef                string                      `json:"password_policy_ref"` // omitempty removed as we need to pass "" when configured "" or null
	BlockingMechanismRef             string                      `json:"blocking_mechanism_ref,omitempty"`
	Sub                              string                      `json:"sub,omitempty"`
	Role                             string                      `json:"role,omitempty"`
	MfaConfiguration                 string                      `json:"mfa_configuration,omitempty"`
	SuggestMfa                       []string                    `json:"suggest_mfa,omitempty"`
	AllowGuestLogin                  *bool                       `json:"allow_guest_login,omitempty"`
	LoginSpi                         *ILoginSPI                  `json:"login_spi,omitempty"`
	BackgroundURI                    string                      `json:"backgroundUri,omitempty"`
	VideoURL                         string                      `json:"videoUrl,omitempty"`
	BotCaptchaRef                    string                      `json:"bot_captcha_ref,omitempty"`
	ApplicationMetaData              map[string]string           `json:"application_meta_data,omitempty"`
	SuggestVerificationMethods       *SuggestVerificationMethods `json:"suggestVerificationMethods,omitempty"`
	GroupRoleRestriction             *GroupRoleRestriction       `json:"groupRoleRestriction,omitempty"`
	BasicSettings                    *BasicSettings              `json:"basic_settings,omitempty"`

	// attributes not available in resource app schema
	RequireAuthTime                  *bool    `json:"require_auth_time,omitempty"`
	BackchannelLogoutSessionRequired *bool    `json:"backchannel_logout_session_required,omitempty"`
	TappID                           string   `json:"tapp_id,omitempty"`
	ClientGroupID                    string   `json:"client_group_id,omitempty"`
	LegalEntity                      string   `json:"legal_entity,omitempty"`
	Tenant                           string   `json:"tenant,omitempty"`
	DeviceCode                       *bool    `json:"device_code,omitempty"`
	TestEmails                       []string `json:"test_emails,omitempty"`
	Active                           *bool    `json:"active,omitempty"`
	EnableLoginSpi                   *bool    `json:"enable_login_spi,omitempty"`
	AcceptRolesInTheRegistration     *bool    `json:"accept_roles_in_the_registration,omitempty"`

	// app_owner removed from schema but assigned while preparing the model
	AppOwner string `json:"app_owner,omitempty"`

	// removed from schema
	// ClientSecretExpiresAt int64       `json:"client_secret_expires_at,omitempty"`
	// ClientIDIssuedAt      int64       `json:"client_id_issued_at,omitempty"`
	// PushConfig            IPushConfig `json:"push_config,omitempty"`
	// CreatedTime           string      `json:"createdTime,omitempty"`
	// UpdatedTime           string      `json:"UpdatedTime,omitempty"`
	// Deleted               bool        `json:"deleted"`
	// AdminClient           bool        `json:"adminClient,omitempty"`
}

type IAllowedGroups struct {
	ID           string   `json:"id,omitempty"`
	SecondaryID  string   `json:"_id,omitempty"`
	GroupID      string   `json:"groupId,omitempty"`
	Roles        []string `json:"roles,omitempty"`
	DefaultRoles []string `json:"default_roles,omitempty"`
}
type ILoginSPI struct {
	OauthClientID string `json:"oauth_client_id,omitempty"`
	SpiURL        string `json:"spi_url,omitempty"`
}
type IAppMobileSettings struct {
	TeamID      string `json:"teamId,omitempty"`
	BundleID    string `json:"bundleId,omitempty"`
	PackageName string `json:"packageName,omitempty"`
	KeyHash     string `json:"keyHash,omitempty"`
}

type IGroupSelection struct {
	AlwaysShowGroupSelection *bool    `json:"alwaysShowGroupSelection"`
	SelectableGroups         []string `json:"selectableGroups,omitempty"`
	SelectableGroupTypes     []string `json:"selectableGroupTypes,omitempty"`
}

type IMfaOption struct {
	Setting               string   `json:"setting,omitempty"`
	TimeIntervalInSeconds *int64   `json:"time_interval_in_seconds,omitempty"`
	AllowedMethods        []string `json:"allowed_methods,omitempty"`
}

type IPushConfig struct {
	ID          string `json:"id,omitempty"`
	SecondaryID string `json:"_id,omitempty"`
	TenantKey   string `json:"tenantKey,omitempty"`
	Name        string `json:"name,omitempty"`
	Vendor      string `json:"vendor,omitempty"`
	Key         string `json:"key,omitempty"`
	Secret      string `json:"secret,omitempty"`
	Active      bool   `json:"active,omitempty"`
	Owner       string `json:"owner,omitempty"`
	ClassName   string `json:"className,omitempty"`
	CreatedTime string `json:"createdTime,omitempty"`
	UpdatedTime string `json:"updatedTime,omitempty"`
}

type ISocialProviderData struct {
	ProviderName string `json:"provider_name,omitempty"`
	SocialID     string `json:"social_id,omitempty"`
}
type IProviderMetadData struct {
	LogoURL           string   `json:"logo_url,omitempty"`
	ProviderName      string   `json:"provider_name,omitempty"`
	DisplayName       string   `json:"display_name,omitempty"`
	Type              string   `json:"type,omitempty"`
	IsProviderVisible *bool    `json:"isProviderVisible,omitempty"`
	Domains           []string `json:"domains,omitempty"`
}

type SuggestVerificationMethods struct {
	MandatoryConfig    MandatoryConfig `json:"mandatoryConfig"`
	OptionalConfig     OptionalConfig  `json:"optionalConfig,omitempty"`
	SkipDurationInDays int32           `json:"skipDurationInDays,omitempty"`
}

type MandatoryConfig struct {
	OptionalConfig
	SkipUntil string `json:"skipUntil,omitempty"`
	Range     string `json:"range,omitempty"`
}

type OptionalConfig struct {
	Methods []string `json:"methods,omitempty"`
}

type GroupRoleRestriction struct {
	MatchCondition string             `json:"matchCondition,omitempty"`
	Filters        []GroupRoleFilters `json:"filters,omitempty"`
}

type GroupRoleFilters struct {
	GroupID    string     `json:"groupId,omitempty"`
	RoleFilter RoleFilter `json:"roleFilter,omitempty"`
}

type RoleFilter struct {
	MatchCondition string   `json:"matchCondition,omitempty"`
	Roles          []string `json:"roles,omitempty"`
}

type BasicSettings struct {
	ClientID                string         `json:"client_id,omitempty"`
	ClientIssuedAt          int64          `json:"client_issued_at,omitempty"`
	TokenEndpointAuthMethod string         `json:"token_endpoint_auth_method,omitempty"`
	RedirectURIs            []string       `json:"redirect_uris,omitempty"`
	AllowedLogoutUrls       []string       `json:"allowed_logout_urls,omitempty"`
	AppOwner                string         `json:"app_owner,omitempty"`
	AllowedScopes           []string       `json:"allowed_scopes,omitempty"`
	HostedPageGroup         string         `json:"hosted_page_group,omitempty"`
	ClientSecrets           []ClientSecret `json:"client_secrets,omitempty"`
}

type ClientSecret struct {
	ClientSecret          string `json:"client_secret,omitempty"`
	ClientSecretExpiresAt int64  `json:"client_secret_expires_at,omitempty"`
	ClientSecretIssuedAt  int64  `json:"client_secret_issued_at,omitempty"`
}

var _ AppService = &App{}

type App struct {
	ClientConfig
}
type AppService interface {
	Create(app AppModel) (AppResponse, error)
	Get(clientID string) (*AppResponse, error)
	Update(app AppModel) (AppResponse, error)
	Delete(clientID string) error
}

func NewApp(clientConfig ClientConfig) AppService {
	return &App{clientConfig}
}

func (a *App) Create(app AppModel) (AppResponse, error) {
	var response AppResponse
	url := fmt.Sprintf("%s/%s", a.BaseURL, "apps-srv/clients")
	httpClient := util.NewHTTPClient(url, http.MethodPost, a.AccessToken)

	res, err := httpClient.MakeRequest(app)
	if err = util.HandleResponseError(res, err); err != nil {
		return response, err
	}
	defer res.Body.Close()

	if err = util.ProcessResponse(res, &response); err != nil {
		return response, err
	}
	return response, nil
}

func (a *App) Get(clientID string) (*AppResponse, error) {
	var response AppResponse
	url := fmt.Sprintf("%s/%s/%s", a.BaseURL, "apps-srv/clients", clientID)
	httpClient := util.NewHTTPClient(url, http.MethodGet, a.AccessToken)

	res, err := httpClient.MakeRequest(nil)
	if err = util.HandleResponseError(res, err); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err = util.ProcessResponse(res, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (a *App) Update(app AppModel) (AppResponse, error) {
	var response AppResponse
	url := fmt.Sprintf("%s/%s", a.BaseURL, "apps-srv/clients")
	httpClient := util.NewHTTPClient(url, http.MethodPut, a.AccessToken)

	res, err := httpClient.MakeRequest(app)
	if err = util.HandleResponseError(res, err); err != nil {
		return response, err
	}
	defer res.Body.Close()

	if err = util.ProcessResponse(res, &response); err != nil {
		return response, err
	}
	return response, nil
}

func (a *App) Delete(clientID string) error {
	url := fmt.Sprintf("%s/%s/%s", a.BaseURL, "apps-srv/clients", clientID)
	httpClient := util.NewHTTPClient(url, http.MethodDelete, a.AccessToken)

	res, err := httpClient.MakeRequest(nil)
	if err = util.HandleResponseError(res, err); err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}
