package cidaas

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/util"
)

type Client struct {
	Config         ClientConfig
	Role           RoleService
	CustomProvider CustomProvideService
	Scope          ScopeService
	ScopeGroup     ScopeGroupService
	GroupType      GroupTypeService
	UserGroup      UserGroupService
	HostedPage     HostedPageService
	Webhook        WebhookService
	App            AppService
	RegField       RegFieldService
	TemplateGroup  TemplateGroupService
	Template       TemplateService
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
	re := regexp.MustCompile(`/*$`)
	host := re.ReplaceAllString(config.BaseURL, "")
	ht := util.HTTPClient{Token: response.AccessToken, Host: host}
	config.AccessToken = response.AccessToken
	client := &Client{
		Config:         config,
		Role:           NewRole(config),
		CustomProvider: NewCustomProvider(&ht),
		Scope:          NewScope(&ht),
		ScopeGroup:     NewScopeGroup(&ht),
		GroupType:      NewGroupType(&ht),
		UserGroup:      NewUserGroup(&ht),
		HostedPage:     NewHostedPage(&ht),
		Webhook:        NewWebhook(&ht),
		App:            NewApp(&ht),
		RegField:       NewRegField(&ht),
		TemplateGroup:  NewTemplateGroup(&ht),
		Template:       NewTemplate(&ht),
	}
	return client, nil
}
