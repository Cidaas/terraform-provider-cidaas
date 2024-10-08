package cidaas

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/util"
)

type Client struct {
	Role           RoleService
	CustomProvider CustomProvideService
	SocialProvider SocialProviderService
	Scope          ScopeService
	ScopeGroup     ScopeGroupService
	ConsentGroup   ConsentGroupService
	GroupType      GroupTypeService
	UserGroup      UserGroupService
	HostedPage     HostedPageService
	Webhook        WebhookService
	App            AppService
	RegField       RegFieldService
	TemplateGroup  TemplateGroupService
	Template       TemplateService
	PasswordPolicy PasswordPolicyService
	Consent        ConsentService
	ConsentVersion ConsentVersionService
}

type ClientConfig struct {
	ClientID     string
	ClientSecret string
	BaseURL      string
	AccessToken  string
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

func NewClient(config ClientConfig) (*Client, error) {
	re := regexp.MustCompile(`/*$`)
	config.BaseURL = re.ReplaceAllString(config.BaseURL, "")
	tokenURL, err := url.JoinPath(config.BaseURL, "token-srv/token")
	if err != nil {
		return nil, fmt.Errorf("failed to create token url %s", err.Error())
	}
	httpClient := util.NewHTTPClient(tokenURL, http.MethodPost)
	payload := map[string]string{
		"client_id":     config.ClientID,
		"client_secret": config.ClientSecret,
		"grant_type":    "client_credentials",
	}
	res, err := httpClient.MakeRequest(payload)
	if err = util.HandleResponseError(res, err); err != nil {
		return nil, fmt.Errorf("failed to generate access token %s", err.Error())
	}
	defer res.Body.Close()
	var response TokenResponse
	if err = util.ProcessResponse(res, &response); err != nil {
		return nil, fmt.Errorf("failed to generate access token %s", err.Error())
	}
	config.AccessToken = response.AccessToken
	client := &Client{
		Role:           NewRole(config),
		CustomProvider: NewCustomProvider(config),
		Scope:          NewScope(config),
		ScopeGroup:     NewScopeGroup(config),
		GroupType:      NewGroupType(config),
		UserGroup:      NewUserGroup(config),
		HostedPage:     NewHostedPage(config),
		Webhook:        NewWebhook(config),
		App:            NewApp(config),
		RegField:       NewRegField(config),
		TemplateGroup:  NewTemplateGroup(config),
		Template:       NewTemplate(config),
		SocialProvider: NewSocialProvider(config),
		PasswordPolicy: NewPasswordPolicy(config),
		ConsentGroup:   NewConsentGroup(config),
		Consent:        NewConsent(config),
		ConsentVersion: NewConsentVersion(config),
	}
	return client, nil
}
