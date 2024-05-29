package provider

import (
	"context"
	"fmt"
	"os"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/cidaas"
	cidaasDataSources "github.com/Cidaas/terraform-provider-cidaas/internal/datasources"
	cidaasResource "github.com/Cidaas/terraform-provider-cidaas/internal/resources"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type cidaasProvider struct {
	version string
}

type Model struct {
	BaseURL types.String `tfsdk:"base_url"`
}

func Cidaas(version string) func() provider.Provider {
	return func() provider.Provider {
		return &cidaasProvider{
			version: version,
		}
	}
}

func (p *cidaasProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "cidaas"
	resp.Version = "dev"
}

func (p *cidaasProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"base_url": schema.StringAttribute{
				Required:    true,
				Description: "The base url of the Terraform client",
			},
		},
	}
}

func (p *cidaasProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		cidaasDataSources.NewRoleDataSource,
	}
}

func (p *cidaasProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		cidaasResource.NewRoleResource,
		cidaasResource.NewCustomProvider,
		cidaasResource.NewScopeResource,
		cidaasResource.NewScopeGroupResource,
		cidaasResource.NewGroupTypeResource,
	}
}

func (p *cidaasProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data Model
	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	clientID := os.Getenv("TERRAFORM_PROVIDER_CIDAAS_CLIENT_ID")
	clientSecret := os.Getenv("TERRAFORM_PROVIDER_CIDAAS_CLIENT_SECRET")

	if clientID == "" || clientSecret == "" {
		resp.Diagnostics.AddError(
			"missing environment variables",
			"env variable TERRAFORM_PROVIDER_CIDAAS_CLIENT_ID or TERRAFORM_PROVIDER_CIDAAS_CLIENT_SECRET missing "+
				"please check the document https://registry.terraform.io/providers/Cidaas/cidaas/latest/docs")
		return
	}

	clientConfig := cidaas.ClientConfig{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		BaseURL:      data.BaseURL.ValueString(),
	}

	client, err := cidaas.NewClient(clientConfig)
	if err != nil {
		resp.Diagnostics.AddError("provide configuration failed", fmt.Sprintf("failed to create cidaas client %s", err.Error()))
		return
	}
	resp.ResourceData = client
	resp.DataSourceData = client
	tflog.Info(ctx, "provider configured successfully")
}
