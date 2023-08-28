package cidaas

type TokenData struct {
	AccessToken string
	TokenType   string
	Sub         string
	ExpiresIn   int
}

type CidaasClient struct {
	ClientId     string
	ClientSecret string
	AuthUrl      string
	GrantType    string
	RedirectURI  string
	TokenData    TokenData
	AppUrl       string
	BaseUrl      string
	ProvideUrl   string
}
