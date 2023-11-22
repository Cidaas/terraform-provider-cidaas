package cidaas

import (
	"encoding/json"
	"fmt"
	"terraform-provider-cidaas/helper/util"
)

type AppResponse struct {
	Success bool      `json:"success,omitempty"`
	Status  int       `json:"status,omitempty"`
	Data    AppConfig `json:"data,omitempty"`
}

type AppConfig struct {
	ClientType                       string                `json:"client_type,omitempty"`
	AccentColor                      string                `json:"accentColor,omitempty"`
	PrimaryColor                     string                `json:"primaryColor,omitempty"`
	MediaType                        string                `json:"mediaType,omitempty"`
	ContentAlign                     string                `json:"contentAlign,omitempty"`
	AllowLoginWith                   []string              `json:"allow_login_with,omitempty"`
	RedirectURIS                     []string              `json:"redirect_uris,omitempty"`
	AllowedLogoutUrls                []string              `json:"allowed_logout_urls,omitempty"`
	EnableDeduplication              bool                  `json:"enable_deduplication,omitempty"`
	AutoLoginAfterRegister           bool                  `json:"auto_login_after_register,omitempty"`
	EnablePasswordlessAuth           bool                  `json:"enable_passwordless_auth,omitempty"`
	RegisterWithLoginInformation     bool                  `json:"register_with_login_information,omitempty"`
	AllowDisposableEmail             bool                  `json:"allow_disposable_email,omitempty"`
	ValidatePhoneNumber              bool                  `json:"validate_phone_number,omitempty"`
	FdsEnabled                       bool                  `json:"fds_enabled,omitempty"`
	HostedPageGroup                  string                `json:"hosted_page_group,omitempty"`
	ClientName                       string                `json:"client_name,omitempty"`
	ClientDisplayName                string                `json:"client_display_name,omitempty"`
	CompanyName                      string                `json:"company_name,omitempty"`
	CompanyAddress                   string                `json:"company_address,omitempty"`
	CompanyWebsite                   string                `json:"company_website,omitempty"`
	AllowedScopes                    []string              `json:"allowed_scopes,omitempty"`
	ResponseTypes                    []string              `json:"response_types,omitempty"`
	GrantTypes                       []string              `json:"grant_types,omitempty"`
	LoginProviders                   []string              `json:"login_providers,omitempty"`
	AdditionalAccessTokenPayload     []string              `json:"additional_access_token_payload,omitempty"`
	RequiredFields                   []string              `json:"required_fields,omitempty"`
	IsHybridApp                      bool                  `json:"is_hybrid_app,omitempty"`
	AllowedWebOrigins                []string              `json:"allowed_web_origins,omitempty"`
	AllowedOrigins                   []string              `json:"allowed_origins,omitempty"`
	MobileSettings                   IAppMobileSettings    `json:"mobile_settings,omitempty"`
	DefaultMaxAge                    int                   `json:"default_max_age,omitempty"`
	TokenLifetimeInSeconds           int                   `json:"token_lifetime_in_seconds,omitempty"`
	IdTokenLifetimeInSeconds         int                   `json:"id_token_lifetime_in_seconds,omitempty"`
	RefreshTokenLifetimeInSeconds    int                   `json:"refresh_token_lifetime_in_seconds,omitempty"`
	TemplateGroupId                  string                `json:"template_group_id,omitempty"`
	ClientId                         string                `json:"client_id,omitempty"`
	ClientSecret                     string                `json:"client_secret,omitempty"`
	PolicyUri                        string                `json:"policy_uri,omitempty"`
	TosUri                           string                `json:"tos_uri,omitempty"`
	ImprintUri                       string                `json:"imprint_uri,omitempty"`
	Contacts                         []string              `json:"contacts,omitempty"`
	TokenEndpointAuthMethod          string                `json:"token_endpoint_auth_method,omitempty"`
	TokenEndpointAuthSigningAlg      string                `json:"token_endpoint_auth_signing_alg,omitempty"`
	DefaultAcrValues                 []string              `json:"default_acr_values,omitempty"`
	Editable                         bool                  `json:"editable,omitempty"`
	WebMessageUris                   []string              `json:"web_message_uris,omitempty"`
	SocialProviders                  []ISocialProviderData `json:"social_providers,omitempty"`
	CustomProviders                  []IProviderMetadData  `json:"custom_providers,omitempty"`
	SamlProviders                    []IProviderMetadData  `json:"saml_providers,omitempty"`
	AdProviders                      []IProviderMetadData  `json:"ad_providers,omitempty"`
	AppOwner                         string                `json:"app_owner,omitempty"`
	JweEnabled                       bool                  `json:"jwe_enabled,omitempty"`
	UserConsent                      bool                  `json:"user_consent,omitempty"`
	AllowedGroups                    []IAllowedGroups      `json:"allowed_groups,omitempty"`
	OperationsAllowedGroups          []IAllowedGroups      `json:"operations_allowed_groups,omitempty"`
	Deleted                          bool                  `json:"deleted,omitempty"`
	Enabled                          bool                  `json:"enabled,omitempty"`
	AllowedFields                    []string              `json:"allowed_fields,omitempty"`
	AppKey                           IAppKeySettings       `json:"appKey,omitempty"`
	AlwaysAskMfa                     bool                  `json:"always_ask_mfa,omitempty"`
	SmartMfa                         bool                  `json:"smart_mfa,omitempty"`
	AllowedMfa                       []string              `json:"allowed_mfa,omitempty"`
	CaptchaRef                       string                `json:"captcha_ref,omitempty"`
	CaptchaRefs                      []string              `json:"captcha_refs,omitempty"`
	ConsentRefs                      []string              `json:"consent_refs,omitempty"`
	CommunicationMediumVerification  string                `json:"communication_medium_verification,omitempty"`
	EmailVerificationRequired        bool                  `json:"email_verification_required,omitempty"`
	MobileNumberVerificationRequired bool                  `json:"mobile_number_verification_required,omitempty"`
	AllowedRoles                     []string              `json:"allowed_roles,omitempty"`
	DefaultRoles                     []string              `json:"default_roles,omitempty"`
	EnableClassicalProvider          bool                  `json:"enable_classical_provider,omitempty"`
	IsRememberMeSelected             bool                  `json:"is_remember_me_selected,omitempty"`
	EnableBotDetection               bool                  `json:"enable_bot_detection,omitempty"`
	BotProvider                      string                `json:"bot_provider,omitempty"`
	AllowGuestLoginGroups            []IAllowedGroups      `json:"allow_guest_login_groups,omitempty"`
	IsLoginSuccessPageEnabled        bool                  `json:"is_login_success_page_enabled,omitempty"`
	IsRegisterSuccessPageEnabled     bool                  `json:"is_register_success_page_enabled,omitempty"`
	GroupIds                         []string              `json:"groupIds,omitempty"`
	AdminClient                      bool                  `json:"adminClient,omitempty"`
	IsGroupLoginSelectionEnabled     bool                  `json:"isGroupLoginSelectionEnabled,omitempty"`
	GroupSelection                   IGroupSelection       `json:"groupSelection,omitempty"`
	GroupTypes                       []string              `json:"groupTypes,omitempty"`
	BackchannelLogoutUri             string                `json:"backchannel_logout_uri,omitempty"`
	PostLogoutRedirectUris           []string              `json:"post_logout_redirect_uris,omitempty"`
	LogoAlign                        string                `json:"logoAlign,omitempty"`
	Mfa                              IMfaOption            `json:"mfa,omitempty"`
	PushConfig                       IPushConfig           `json:"push_config,omitempty"`
	Webfinger                        string                `json:"webfinger,omitempty"`
	ApplicationType                  string                `json:"application_type,omitempty"`
	LogoUri                          string                `json:"logo_uri,omitempty"`
	InitiateLoginUri                 string                `json:"initiate_login_uri,omitempty"`
	ClientSecretExpiresAt            int                   `json:"client_secret_expires_at,omitempty"`
	ClientIdIssuedAt                 int                   `json:"client_id_issued_at,omitempty"`
	RegistrationClientUri            string                `json:"registration_client_uri,omitempty"`
	RegistrationAccessToken          string                `json:"registration_access_token,omitempty"`
	ClientUri                        string                `json:"client_uri,omitempty"`
	JwksUri                          string                `json:"jwks_uri,omitempty"`
	Jwks                             string                `json:"jwks,omitempty"`
	SectorIdentifierUri              string                `json:"sector_identifier_uri,omitempty"`
	SubjectType                      string                `json:"subject_type,omitempty"`
	IdTokenSignedResponseAlg         string                `json:"id_token_signed_response_alg,omitempty"`
	IdTokenEncryptedResponseAlg      string                `json:"id_token_encrypted_response_alg,omitempty"`
	IdTokenEncryptedResponseEnc      string                `json:"id_token_encrypted_response_enc,omitempty"`
	UserinfoSignedResponseAlg        string                `json:"userinfo_signed_response_alg,omitempty"`
	UserinfoEncryptedResponseAlg     string                `json:"userinfo_encrypted_response_alg,omitempty"`
	UserinfoEncryptedResponseEnc     string                `json:"userinfo_encrypted_response_enc,omitempty"`
	RequestObjectSigningAlg          string                `json:"request_object_signing_alg,omitempty"`
	RequestObjectEncryptionAlg       string                `json:"request_object_encryption_alg,omitempty"`
	RequestObjectEncryptionEnc       string                `json:"request_object_encryption_enc,omitempty"`
	RequestUris                      []string              `json:"request_uris,omitempty"`
	RequireAuthTime                  bool                  `json:"require_auth_time,omitempty"`
	BackchannelLogoutSessionRequired bool                  `json:"backchannel_logout_session_required,omitempty"`
	TappId                           string                `json:"tapp_id,omitempty"`
	ClientGroupId                    string                `json:"client_group_id,omitempty"`
	LegalEntity                      string                `json:"legal_entity,omitempty"`
	Tenant                           string                `json:"tenant,omitempty"`
	DeviceCode                       bool                  `json:"device_code,omitempty"`
	TestEmails                       []string              `json:"test_emails,omitempty"`
	Active                           bool                  `json:"active,omitempty"`
	Description                      string                `json:"description,omitempty"`
	DefaultScopes                    []string              `json:"default_scopes,omitempty"`
	PendingScopes                    []string              `json:"pending_scopes,omitempty"`
	ConsentPageGroup                 string                `json:"consent_page_group,omitempty"`
	PasswordPolicyRef                string                `json:"password_policy_ref,omitempty"`
	BlockingMechanismRef             string                `json:"blocking_mechanism_ref,omitempty"`
	AcceptRolesInTheRegistration     bool                  `json:"accept_roles_in_the_registration,omitempty"`
	Sub                              string                `json:"sub,omitempty"`
	Role                             string                `json:"role,omitempty"`
	MfaConfiguration                 string                `json:"mfa_configuration,omitempty"`
	SuggestMfa                       []string              `json:"suggest_mfa,omitempty"`
	AllowGuestLogin                  bool                  `json:"allow_guest_login,omitempty"`
	LoginSpi                         ILoginSPI             `json:"login_spi,omitempty"`
	EnableLoginSpi                   bool                  `json:"enable_login_spi,omitempty"`
	BackgroundUri                    string                `json:"backgroundUri,omitempty"`
	VideoUrl                         string                `json:"videoUrl,omitempty"`
	BotCaptchaRef                    string                `json:"bot_captcha_ref,omitempty"`
	CreatedTime                      string                `json:"createdTime,omitempty"`
	UpdatedTime                      string                `json:"UpdatedTime,omitempty"`
}

type IAllowedGroups struct {
	Id           string   `json:"id,omitempty"`
	SecondaryId  string   `json:"_id,omitempty"`
	GroupId      string   `json:"groupId,omitempty"`
	Roles        []string `json:"roles,omitempty"`
	DefaultRoles []string `json:"default_roles,omitempty"`
}
type ILoginSPI struct {
	OauthClientId string `json:"oauth_client_id,omitempty"`
	SpiUrl        string `json:"spi_url,omitempty"`
}
type IAppMobileSettings struct {
	Id          string `json:"id,omitempty"`
	SecondaryId string `json:"_id,omitempty"`
	TeamId      string `json:"teamId,omitempty"`
	BundleId    string `json:"bundleId,omitempty"`
	PackageName string `json:"packageName,omitempty"`
	KeyHash     string `json:"keyHash,omitempty"`
}
type IAppKeySettings struct {
	Id           string      `json:"_id,omitempty"`
	KeyType      string      `json:"keyType,omitempty"`
	PublicKey    string      `json:"publicKey,omitempty"`
	PublicKeyJWK interface{} `json:"publicKeyJWK,omitempty"`
	CreatedTime  string      `json:"createdTime,omitempty"`
}
type IGroupSelection struct {
	AlwaysShowGroupSelection bool     `json:"alwaysShowGroupSelection,omitempty"`
	SelectableGroups         []string `json:"selectableGroups,omitempty"`
	SelectableGroupTypes     []string `json:"selectableGroupTypes,omitempty"`
}

type IMfaOption struct {
	Setting               string   `json:"setting,omitempty"`
	TimeIntervalInSeconds int      `json:"time_interval_in_seconds,omitempty"`
	AllowedMethods        []string `json:"allowed_methods,omitempty"`
}

type IPushConfig struct {
	Id          string `json:"id,omitempty"`
	SecondaryId string `json:"_id,omitempty"`
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
	SocialId     string `json:"social_id,omitempty"`
	DisplayName  string `json:"display_name,omitempty"`
}
type IProviderMetadData struct {
	LogoUrl      string `json:"logo_url,omitempty"`
	ProviderName string `json:"provider_name,omitempty"`
	DisplayName  string `json:"display_name,omitempty"`
	Type         string `json:"type,omitempty"`
}

func (c *CidaasClient) CreateApp(app AppConfig) (response *AppResponse, err error) {
	h := util.HttpClient{
		Token: c.TokenData.AccessToken,
	}
	res, err := h.Post(c.AppUrl, app)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %v", err)
	}
	return response, nil
}

func (c *CidaasClient) UpdateApp(app AppConfig) (response *AppResponse, err error) {
	h := util.HttpClient{
		Token: c.TokenData.AccessToken,
	}
	res, err := h.Put(c.AppUrl, app)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %v", err)
	}
	return response, nil
}

func (c *CidaasClient) DeleteApp(app AppConfig) (response *AppResponse, err error) {
	h := util.HttpClient{
		Token: c.TokenData.AccessToken,
	}
	url := c.BaseUrl + "/apps-srv/clients/" + app.ClientId
	res, err := h.Delete(url)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %v", err)
	}
	return response, nil
}

func (c *CidaasClient) GetApp(app AppConfig) (response *AppResponse, err error) {
	h := util.HttpClient{
		Token: c.TokenData.AccessToken,
	}
	url := c.BaseUrl + "/apps-srv/clients/" + app.ClientId
	res, err := h.Get(url)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %v", err)
	}
	return response, nil
}

func SerializeMobileSettings(ms []interface{}) (resp IAppMobileSettings) {
	for _, value := range ms {
		if value != nil {
			temp := value.(map[string]interface{})
			if temp["team_id"] != nil {
				resp.TeamId = temp["team_id"].(string)
			}
			if temp["bundle_id"] != nil {
				resp.BundleId = temp["bundle_id"].(string)
			}
			if temp["package_name"] != nil {
				resp.PackageName = temp["package_name"].(string)
			}
			if temp["key_hash"] != nil {
				resp.KeyHash = temp["key_hash"].(string)
			}
		}
	}
	return resp
}

func SerializeSocialProviders(sp []interface{}) (resp []ISocialProviderData) {
	for _, value := range sp {
		if value != nil {
			temp := value.(map[string]interface{})
			var socialProvider ISocialProviderData
			if temp["social_id"] != nil {
				socialProvider.SocialId = temp["social_id"].(string)
			}
			if temp["provider_name"] != nil {
				socialProvider.ProviderName = temp["provider_name"].(string)
			}
			if temp["display_name"] != nil {
				socialProvider.DisplayName = temp["display_name"].(string)
			}
			resp = append(resp, socialProvider)
		}
	}
	return resp
}

func SerializeProviders(providers []interface{}) (resp []IProviderMetadData) {
	for _, value := range providers {
		if value != nil {
			temp := value.(map[string]interface{})
			var provider IProviderMetadData
			if temp["logo_url"] != nil {
				provider.LogoUrl = temp["logo_url"].(string)
			}
			if temp["provider_name"] != nil {
				provider.ProviderName = temp["provider_name"].(string)
			}
			if temp["display_name"] != nil {
				provider.DisplayName = temp["display_name"].(string)
			}
			if temp["type"] != nil {
				provider.Type = temp["type"].(string)
			}
			resp = append(resp, provider)
		}
	}
	return resp
}

func SerializeAllowedGroups(groups []interface{}) (resp []IAllowedGroups) {
	for _, value := range groups {
		if value != nil {
			temp := value.(map[string]interface{})
			var group IAllowedGroups
			if temp["group_id"] != nil {
				group.GroupId = temp["group_id"].(string)
			}
			if temp["roles"] != nil {
				group.Roles = util.InterfaceArray2StringArray(temp["roles"].([]interface{}))
			}
			if temp["default_roles"] != nil {
				group.DefaultRoles = util.InterfaceArray2StringArray(temp["default_roles"].([]interface{}))
			}
			resp = append(resp, group)
		}
	}
	return resp
}

func SerializeGroupSelection(groups []interface{}) (resp IGroupSelection) {
	for _, value := range groups {
		if value != nil {
			temp := value.(map[string]interface{})
			if temp["always_show_group_selection"] != nil {
				resp.AlwaysShowGroupSelection = temp["always_show_group_selection"].(bool)
			}
			if temp["selectable_groups"] != nil {
				resp.SelectableGroupTypes = util.InterfaceArray2StringArray(temp["selectable_groups"].([]interface{}))
			}
			if temp["selectable_group_types"] != nil {
				resp.SelectableGroupTypes = util.InterfaceArray2StringArray(temp["selectable_group_types"].([]interface{}))
			}
		}
	}
	return resp
}

func SerializeMfaOption(options []interface{}) (resp IMfaOption) {
	for _, value := range options {
		if value != nil {
			temp := value.(map[string]interface{})
			if temp["setting"] != nil {
				resp.Setting = temp["setting"].(string)
			}
			if temp["time_interval_in_seconds"] != nil {
				resp.TimeIntervalInSeconds = temp["time_interval_in_seconds"].(int)
			}
			if temp["allowed_methods"] != nil {
				resp.AllowedMethods = util.InterfaceArray2StringArray(temp["allowed_methods"].([]interface{}))
			}
		}
	}
	return resp
}

func SerializePushConfig(configs []interface{}) (resp IPushConfig) {
	for _, value := range configs {
		if value != nil {
			temp := value.(map[string]interface{})
			if temp["tenant_key"] != nil {
				resp.TenantKey = temp["tenant_key"].(string)
			}
			if temp["name"] != nil {
				resp.Name = temp["name"].(string)
			}
			if temp["vendor"] != nil {
				resp.Vendor = temp["vendor"].(string)
			}
			if temp["key"] != nil {
				resp.Key = temp["key"].(string)
			}
			if temp["secret"] != nil {
				resp.Secret = temp["secret"].(string)
			}
			if temp["active"] != nil {
				resp.Active = temp["active"].(bool)
			}
			if temp["owner"] != nil {
				resp.Owner = temp["owner"].(string)
			}
		}
	}
	return resp
}

func SerializeLoginSpi(spi []interface{}) (resp ILoginSPI) {
	for _, value := range spi {
		if value != nil {
			temp := value.(map[string]interface{})
			if temp["oauth_client_id"] != nil {
				resp.OauthClientId = temp["oauth_client_id"].(string)
			}
			if temp["spi_url"] != nil {
				resp.SpiUrl = temp["spi_url"].(string)
			}
		}
	}
	return resp
}

func SerializeAppKey(aks IAppKeySettings) []interface{} {
	fields := make(map[string]interface{})
	fields["_id"] = aks.Id
	fields["key_type"] = aks.KeyType
	fields["public_key"] = aks.PublicKey
	fields["public_key_jwk"] = aks.PublicKeyJWK
	fields["created_time"] = aks.CreatedTime
	return []interface{}{fields}
}

func FlattenMobileSettings(mbs IAppMobileSettings) []interface{} {
	fields := make(map[string]interface{})
	fields["id"] = mbs.Id
	fields["_id"] = mbs.SecondaryId
	fields["team_id"] = mbs.TeamId
	fields["bundle_id"] = mbs.BundleId
	fields["package_name"] = mbs.PackageName
	fields["key_hash"] = mbs.KeyHash
	return []interface{}{fields}
}

func FlattenSocialProvider(sps *[]ISocialProviderData) []interface{} {
	if sps != nil {
		tempSps := make([]interface{}, len(*sps), len(*sps))
		for i, sp := range *sps {
			temp := make(map[string]interface{})
			temp["provider_name"] = sp.ProviderName
			temp["social_id"] = sp.SocialId
			temp["display_name"] = sp.DisplayName
			tempSps[i] = temp
		}
		return tempSps
	}
	return make([]interface{}, 0)
}

func FlattenProviders(pmds *[]IProviderMetadData) []interface{} {
	if pmds != nil {
		tempSps := make([]interface{}, len(*pmds), len(*pmds))
		for i, pmd := range *pmds {
			temp := make(map[string]interface{})
			temp["logo_url"] = pmd.LogoUrl
			temp["provider_name"] = pmd.ProviderName
			temp["display_name"] = pmd.DisplayName
			temp["type"] = pmd.Type
			tempSps[i] = temp
		}
		return tempSps
	}
	return make([]interface{}, 0)
}

func FlattenAllowedGroups(ags *[]IAllowedGroups) []interface{} {
	if ags != nil {
		tempAgs := make([]interface{}, len(*ags), len(*ags))
		for i, pmd := range *ags {
			temp := make(map[string]interface{})
			temp["id"] = pmd.Id
			temp["_id"] = pmd.SecondaryId
			temp["group_id"] = pmd.GroupId
			temp["roles"] = pmd.Roles
			temp["default_roles"] = pmd.DefaultRoles
			tempAgs[i] = temp
		}
		return tempAgs
	}
	return make([]interface{}, 0)
}

func FlattenGroupSelection(gss IGroupSelection) []interface{} {
	fields := make(map[string]interface{})
	fields["always_show_group_selection"] = gss.AlwaysShowGroupSelection
	fields["selectable_groups"] = gss.SelectableGroups
	fields["selectable_group_types"] = gss.SelectableGroupTypes
	return []interface{}{fields}
}

func FlattenMfa(mo IMfaOption) []interface{} {
	fields := make(map[string]interface{})
	fields["setting"] = mo.Setting
	fields["time_interval_in_seconds"] = mo.TimeIntervalInSeconds
	fields["allowed_methods"] = mo.AllowedMethods
	return []interface{}{fields}
}

func FlattenPushConfig(pc IPushConfig) []interface{} {
	fields := make(map[string]interface{})
	fields["_id"] = pc.SecondaryId
	fields["id"] = pc.Id
	fields["tenant_key"] = pc.TenantKey
	fields["name"] = pc.Name
	fields["vendor"] = pc.Vendor
	fields["key"] = pc.Key
	fields["secret"] = pc.Secret
	fields["owner"] = pc.Owner
	return []interface{}{fields}
}

func FlattenLoginSpi(ls ILoginSPI) []interface{} {
	fields := make(map[string]interface{})
	fields["oauth_client_id"] = ls.OauthClientId
	fields["spi_url"] = ls.SpiUrl
	return []interface{}{fields}
}
