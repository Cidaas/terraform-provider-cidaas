package cidaas_sdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type LoginRequestPayload struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectURI  string `json:"redirect_uri"`
	GrantType    string `json:"grant_type"`
}

type LoginResponsePayload struct {
	AccessToken      string `json:"access_token"`
	TokenType        string `json:"token_type"`
	Sub              string `json:"sub"`
	ExpiresIn        int    `json:"expires_in"`
	IdToken          string `json:"id_token"`
	RefreshToken     string `json:"refresh_token"`
	IdentityId       string `json:"identity_id"`
	IdTokenExpiresIn int    `json:"id_token_expires_in"`
}

// Function to initialize client Auth
func InitializeAuth(cidaas_client *CidaasClient) {
	url := cidaas_client.AuthUrl
	method := "POST"

	loginrequestpayload := &LoginRequestPayload{
		ClientId:     cidaas_client.ClientId,
		ClientSecret: cidaas_client.ClientSecret,
		Username:     cidaas_client.Username,
		Password:     cidaas_client.Password,
		RedirectURI:  cidaas_client.RedirectURI,
		GrantType:    cidaas_client.GrantType,
	}

	json_payload, err := json.Marshal(loginrequestpayload)
	if err != nil {
		fmt.Println(err)
		return
	}

	payload_string := string(json_payload)
	payload := strings.NewReader(payload_string)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

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
	log.Println("BODY  ", string(body))

	var loginresponse LoginResponsePayload

	if err := json.Unmarshal([]byte(body), &loginresponse); err != nil {
		log.Printf("JSON UNMARSHAL")
		log.Printf("Error: %s", err.Error())
		return
	 }
	

	cidaas_client.TokenData.AccessToken = loginresponse.AccessToken
	cidaas_client.TokenData.TokenType = loginresponse.TokenType
	cidaas_client.TokenData.Sub = loginresponse.Sub
	cidaas_client.TokenData.ExpiresIn = loginresponse.ExpiresIn
	cidaas_client.TokenData.IdTokenExpiresIn = loginresponse.IdTokenExpiresIn
	cidaas_client.TokenData.IdToken = loginresponse.IdToken
	cidaas_client.TokenData.RefreshToken = loginresponse.RefreshToken
	cidaas_client.TokenData.IdentityId = loginresponse.IdentityId

	log.Println("CLIENT_SDK TOKEN TYPE ", cidaas_client.TokenData.TokenType)
	log.Println("CLIENT_SDK SUB ", cidaas_client.TokenData.Sub)
	log.Println("CLIENT_SDK  access token", cidaas_client.TokenData.AccessToken)


}
