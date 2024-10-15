package datasources

import (
	"context"
	"fmt"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/cidaas"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ScopeGroupDataSource struct {
	BaseDataSource
}

type ScopeGroupFilterModel struct {
	BaseModel
	ScopeGroup []ScopeGroup `tfsdk:"scope_group"`
}

type ScopeGroup struct {
	ID          types.String `tfsdk:"id"`
	GroupName   types.String `tfsdk:"group_name"`
	Description types.String `tfsdk:"description"`
}

var scopeGroupFilter = FilterConfig{
	"group_name": {TypeFunc: FilterTypeString},
}

var scopeGroupSchema = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The ID of th resource.",
	},
	"group_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the group.",
	},
	"description": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The `description` attribute provides details about the scope of the group, explaining its purpose.",
	},
}

var scopeGroupDataSourceSchema = schema.Schema{
	MarkdownDescription: fmt.Sprintf("The data source `%s` returns a list of scope groups available in your Cidaas instance."+
		"\nYou can apply filters using the `filter` block in your Terraform configuration.", SCOPE_GRUOP_DATASOURCE),
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Description: "The data source's unique ID.",
			Computed:    true,
		},
	},
	Blocks: map[string]schema.Block{
		"filter": scopeGroupFilter.Schema(),
		"scope_group": schema.ListNestedBlock{
			Description: "The returned list of scope groups",
			NestedObject: schema.NestedBlockObject{
				Attributes: scopeGroupSchema,
			},
		},
	},
}

func NewScopeGroup() datasource.DataSource {
	return &ScopeGroupDataSource{
		BaseDataSource: NewBaseDataSource(
			BaseDataSourceConfig{
				Name:   SCOPE_GRUOP_DATASOURCE,
				Schema: &scopeGroupDataSourceSchema,
			},
		),
	}
}

func (d *ScopeGroupDataSource) Read(
	ctx context.Context,
	req datasource.ReadRequest,
	resp *datasource.ReadResponse,
) {
	var data ScopeGroupFilterModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.ID = types.StringValue(uuid.New().String())
	result, diag := scopeGroupFilter.GetAndFilter(d.Client, data.Filters, listScopeGroups)
	if diag != nil {
		resp.Diagnostics.Append(diag)
		return
	}

	data.ScopeGroup = parseModel(AnySliceToTyped[cidaas.ScopeGroupConfig](result), parseScopeGroup)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func listScopeGroups(client *cidaas.Client) ([]any, error) {
	sgs, err := client.ScopeGroup.GetAll()
	if err != nil {
		return nil, err
	}
	return TypedSliceToAny(sgs), nil
}

func parseScopeGroup(scope cidaas.ScopeGroupConfig) (result ScopeGroup) {
	result.ID = types.StringValue(scope.ID)
	result.GroupName = types.StringValue(scope.GroupName)
	result.Description = types.StringValue(scope.Description)
	return result
}
