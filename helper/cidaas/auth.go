package cidaas

import (
	"encoding/json"
	"fmt"
	"terraform-provider-cidaas/helper/util"
)

type TokenRequestPayload struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectURI  string `json:"redirect_uri"`
	GrantType    string `json:"grant_type"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

func InitializeAuth(cidaas_client *CidaasClient) error {
	payload := &TokenRequestPayload{
		ClientId:     cidaas_client.ClientId,
		ClientSecret: cidaas_client.ClientSecret,
		GrantType:    cidaas_client.GrantType,
	}

	httpClient := util.HttpClient{}
	res, err := httpClient.Post(cidaas_client.AuthUrl, payload)
	if err != nil {
		return fmt.Errorf("failed to get token: %+v", err)
	}
	defer res.Body.Close()
	var response TokenResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return fmt.Errorf("failed to decode token response: %+v", err)
	}
	cidaas_client.TokenData.AccessToken = response.AccessToken
	return nil
}
