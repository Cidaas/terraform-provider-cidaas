package datasources

import (
	"context"
	"fmt"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/cidaas"
	"github.com/Cidaas/terraform-provider-cidaas/helpers/util"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// nolint:revive
const (
	CONSENT_DATASOURCE         = "cidaas_consent"                // nolint:stylecheck
	CUSTOM_PROVIDER_DATASOURCE = "cidaas_custom_provider"        // nolint:stylecheck
	GROUP_TYPE_DATASOURCE      = "cidaas_group_type"             // nolint:stylecheck
	REG_FIELD_DATASOURCE       = "cidaas_registration_field"     // nolint:stylecheck
	ROLE_DATASOURCE            = "cidaas_role"                   // nolint:stylecheck
	SCOPE_GRUOP_DATASOURCE     = "cidaas_scope_group"            // nolint:stylecheck
	SCOPE_DATASOURCE           = "cidaas_scope"                  // nolint:stylecheck
	SOCIAL_PROVIDER_DATASOURCE = "cidaas_social_provider"        // nolint:stylecheck
	SYSTEM_TEMPLATE_DATASOURCE = "cidaas_system_template_option" // nolint:stylecheck
)

type Parser[K, V any] func(K) V

type BaseDataSourceConfig struct {
	Name   string
	Schema *schema.Schema
}

type BaseDataSource struct {
	Config BaseDataSourceConfig
	Client *cidaas.Client
}

type BaseModel struct {
	ID      types.String     `tfsdk:"id"`
	Filters FiltersModelType `tfsdk:"filter"`
}

func NewBaseDataSource(cfg BaseDataSourceConfig) BaseDataSource {
	return BaseDataSource{
		Config: cfg,
	}
}

func (r *BaseDataSource) Configure(
	ctx context.Context,
	req datasource.ConfigureRequest,
	resp *datasource.ConfigureResponse,
) {
	// Prevent panic if the provider has not been configured
	if req.ProviderData == nil {
		return
	}

	r.Client = GetDataSourceMeta(ctx, req, resp)
	if resp.Diagnostics.HasError() {
		return
	}
}

func GetDataSourceMeta(
	ctx context.Context,
	req datasource.ConfigureRequest,
	resp *datasource.ConfigureResponse,
) *cidaas.Client {
	tflog.Debug(ctx, "starting data source configuration")
	client, ok := req.ProviderData.(*cidaas.Client)

	if !ok {
		tflog.Error(ctx, "Failed to configure data source - unexpected provider data type", util.H{
			"expected_type": "*cidaas.Client",
			"actual_type":   fmt.Sprintf("%T", req.ProviderData),
		})
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected cidaas.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return nil
	}
	tflog.Info(ctx, "data source configured successfully")
	return client
}

func (r *BaseDataSource) Metadata(
	_ context.Context,
	_ datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	resp.TypeName = r.Config.Name
}

func (r *BaseDataSource) Schema(
	_ context.Context,
	_ datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	if r.Config.Schema == nil {
		resp.Diagnostics.AddError(
			"Missing Schema",
			"Base data source was not provided a schema. "+
				"Please provide a Schema config attribute or implement, the Schema(...) function.",
		)
		return
	}
	resp.Schema = *r.Config.Schema
}

// parses the api response data & serialize to provider schema model
func parseModel[DataModel, ProviderModel any](
	data []DataModel,
	modelParseFunc Parser[DataModel, ProviderModel],
) []ProviderModel {
	result := make([]ProviderModel, len(data))
	for i := range data {
		result[i] = modelParseFunc(data[i])
	}
	return result
}
