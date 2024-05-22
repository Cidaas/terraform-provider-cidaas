package cidaas

import (
	"encoding/json"
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
	HTTPClient     util.HTTPClient
}

type ClientConfig struct {
	ClientID     string
	ClientSecret string
	BaseURL      string
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
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token %s", err.Error())
	}
	defer res.Body.Close()
	var response TokenResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode token response %s", err.Error())
	}
	re := regexp.MustCompile(`/*$`)
	host := re.ReplaceAllString(config.BaseURL, "")
	ht := util.HTTPClient{Token: response.AccessToken, Host: host}
	client := &Client{
		Config:         config,
		Role:           NewRole(&ht),
		CustomProvider: NewCustomProvider(&ht),
	}
	return client, nil
}
