package cidaas_sdk

type TokenData struct {
	AccessToken      string
	IdToken          string
	TokenType        string
	Sub              string
	ExpiresIn        int
	IdTokenExpiresIn int
	RefreshToken     string
	IdentityId       string
}

type CidaasClient struct {
	ClientId     string
	ClientSecret string
	AuthUrl      string
	GrantType    string
	RedirectURI  string
	TokenData    TokenData // generated
	AppUrl       string
	BaseUrl      string
	ProvideUrl   string
}

func ClientBuilder(
	cidaas_client *CidaasClient,
	client_id string,
	client_secret string,
	redirect_uri string,
	grant_type string,
	auth_url string,
	app_url string,
	base_url string,
	provider_url string) {

	cidaas_client.ClientId = client_id
	cidaas_client.ClientSecret = client_secret
	cidaas_client.RedirectURI = redirect_uri
	cidaas_client.GrantType = grant_type
	cidaas_client.AuthUrl = auth_url
	cidaas_client.AppUrl = app_url
	cidaas_client.BaseUrl = base_url
	cidaas_client.ProvideUrl = provider_url
}
