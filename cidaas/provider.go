package cidaas

import (
	"context"
	"os"

	"terraform-provider-cidaas/helper/cidaas"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"base_url": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"cidaas_app":                     resourceApp(),
			"cidaas_registration_page_field": resourceRegistrationField(),
			"cidaas_custom_provider":         resourceCustomProvider(),
			"cidaas_scope":                   resourceScope(),
			"cidaas_scope_group":             resourceScopeGroup(),
			"cidaas_role":                    resourceRole(),
			"cidaas_webhook":                 resourceWebhook(),
			"cidaas_hosted_page":             resourceHostedPage(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	base_url := d.Get("base_url").(string)
	client_id := os.Getenv("TERRAFORM_PROVIDER_CIDAAS_CLIENT_ID")
	client_secret := os.Getenv("TERRAFORM_PROVIDER_CIDAAS_CLIENT_SECRET")

	var diags diag.Diagnostics

	if client_id == "" || client_secret == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "missing environment variables",
			Detail:   `env variable TERRAFORM_PROVIDER_CIDAAS_CLIENT_ID or TERRAFORM_PROVIDER_CIDAAS_CLIENT_SECRET missing. please check the document https://registry.terraform.io/providers/Cidaas/cidaas/latest/docs`,
		})
		return nil, diags
	}

	cidaas_client := cidaas.CidaasClient{
		ClientId:     client_id,
		ClientSecret: client_secret,
		GrantType:    "client_credentials",
		AuthUrl:      base_url + "/token-srv/token",
		AppUrl:       base_url + "/apps-srv/clients",
		BaseUrl:      base_url,
		ProvideUrl:   base_url + "/providers-srv/custom",
	}
	cidaas.InitializeAuth(&cidaas_client)
	return cidaas_client, diags
}
