package datasources

import (
	"context"
	"fmt"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/cidaas"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
	_ context.Context,
	req datasource.ConfigureRequest,
	resp *datasource.ConfigureResponse,
) {
	// Prevent panic if the provider has not been configured
	if req.ProviderData == nil {
		return
	}

	r.Client = GetDataSourceMeta(req, resp)
	if resp.Diagnostics.HasError() {
		return
	}
}

func GetDataSourceMeta(
	req datasource.ConfigureRequest,
	resp *datasource.ConfigureResponse,
) *cidaas.Client {
	client, ok := req.ProviderData.(*cidaas.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected cidaas.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return nil
	}
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
