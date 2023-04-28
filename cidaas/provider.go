package cidaas

import (
	"context"

	"terraform-provider-cidaas/helper_pkg/cidaas_sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"default_app_client_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			"default_app_client_secret": {
				Type:     schema.TypeString,
				Required: true,
			},

			"default_app_redirect_uri": {
				Type:     schema.TypeString,
				Required: true,
			},

			"default_app_auth_url": {
				Type:     schema.TypeString,
				Required: true,
			},

			"default_app_app_url": {
				Type:     schema.TypeString,
				Required: true,
			},

			"default_app_base_url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"default_app_provider_url": {
				Type:     schema.TypeString,
				Optional: true,
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
	client_id := d.Get("default_app_client_id").(string)
	client_secret := d.Get("default_app_client_secret").(string)
	redirect_uri := d.Get("default_app_redirect_uri").(string)
	grant_type := "client_credentials"
	auth_url := d.Get("default_app_auth_url").(string)
	app_url := d.Get("default_app_app_url").(string)
	base_url := d.Get("default_app_base_url").(string)
	provide_url := d.Get("default_app_provider_url").(string)

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
