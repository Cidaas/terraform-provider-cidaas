package cidaas

import (
	"context"
	"os"

	"terraform-provider-cidaas/helper_pkg/cidaas_sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"redirect_uri": {
				Type:     schema.TypeString,
				Required: true,
			},

			"base_url": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"cidaas_app":                     resourceApp(),
			"cidaas_registration_page_field": resourceRegistrationField(),
			"cidaas_custom_provider":         resourceCustomProvider(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"cidaas_app":             dataSourceApp(),
			"cidaas_custom_provider": dataSourceApp(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {

	redirect_uri := d.Get("redirect_uri").(string)
	base_url := d.Get("base_url").(string)

	grant_type := "client_credentials"
	auth_url := base_url + "/token-srv/token"
	provide_url := base_url + "/providers-srv/custom"
	app_url := base_url + "/apps-srv/clients"

	client_id := os.Getenv("TERRAFORM_PROVIDER_CIDAAS_CLIENT_ID")
	client_secret := os.Getenv("TERRAFORM_PROVIDER_CIDAAS_CLIENT_SECRET")

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	var cidaas_client cidaas_sdk.CidaasClient

	cidaas_sdk.ClientBuilder(&cidaas_client,
		client_id,
		client_secret,
		redirect_uri,
		grant_type,
		auth_url,
		app_url,
		base_url,
		provide_url)

	cidaas_sdk.InitializeAuth(&cidaas_client)

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Payload Warning Message Summary",
		Detail:   cidaas_client.TokenData.Sub,
	})

	return cidaas_client, diags
}
