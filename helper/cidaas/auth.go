package cidaas

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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

func InitializeAuth(cidaas_client *CidaasClient) {
	client := &http.Client{}

	requestPayload := &TokenRequestPayload{
		ClientId:     cidaas_client.ClientId,
		ClientSecret: cidaas_client.ClientSecret,
		RedirectURI:  cidaas_client.RedirectURI,
		GrantType:    cidaas_client.GrantType,
	}

	json_payload, err := json.Marshal(requestPayload)
	if err != nil {
		fmt.Println(err)
		return
	}
	payload_string := string(json_payload)
	payload := strings.NewReader(payload_string)
	req, err := http.NewRequest("POST", cidaas_client.AuthUrl, payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var response TokenResponse
	json.Unmarshal([]byte(body), &response)
	cidaas_client.TokenData.AccessToken = response.AccessToken
}
