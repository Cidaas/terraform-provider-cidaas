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

type RegistrationFieldsDataSource struct {
	BaseDataSource
}

type RegistrationFieldsFilterModel struct {
	BaseModel
	RegistrationField []RegistrationField `tfsdk:"registration_field"`
}

type RegistrationField struct {
	ID            types.String `tfsdk:"id"`
	FieldKey      types.String `tfsdk:"field_key"`
	DataType      types.String `tfsdk:"data_type"`
	Internal      types.Bool   `tfsdk:"internal"`
	ReadOnly      types.Bool   `tfsdk:"read_only"`
	Required      types.Bool   `tfsdk:"required"`
	Enabled       types.Bool   `tfsdk:"enabled"`
	IsGroup       types.Bool   `tfsdk:"is_group"`
	ParentGroupID types.String `tfsdk:"parent_group_id"`
	FieldType     types.String `tfsdk:"field_type"`
	Order         types.Int32  `tfsdk:"order"`
}

var registrationFieldsFilter = FilterConfig{
	"parent_group_id": {TypeFunc: FilterTypeString},
	"field_type":      {TypeFunc: FilterTypeString},
	"data_type":       {TypeFunc: FilterTypeString},
	"field_key":       {TypeFunc: FilterTypeString},
	"required":        {TypeFunc: FilterTypeBool},
	"internal":        {TypeFunc: FilterTypeBool},
	"read_only":       {TypeFunc: FilterTypeBool},
	"is_group":        {TypeFunc: FilterTypeBool},
	"enabled":         {TypeFunc: FilterTypeBool},
}

var fieldSchema = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The unique identifier of the group type.",
	},
	"parent_group_id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The ID of the parent registration group.",
	},
	"field_type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Specifies whether the field type is `SYSTEM` or `CUSTOM`.",
	},
	"data_type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The data type of the field.",
	},
	"field_key": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The unique name of the registration field.",
	},
	"required": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Flag to identify if a field is required in registration.",
	},
	"internal": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Flag to identify if a field is internal.",
	},
	"enabled": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Flag to identify if a field is enabled.",
	},
	"read_only": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Flag to identify if a field is read only.",
	},
	"is_group": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Flag to identify if a field is group field.",
	},
	"order": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The order of the Field in the UI.",
	},
}

var fieldDataSourceSchema = schema.Schema{
	MarkdownDescription: fmt.Sprintf("The data source `%s` returns a list of registration fields available in your Cidaas instance."+
		"\nYou can apply filters using the `filter` block in your Terraform configuration.", REG_FIELD_DATASOURCE),
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Description: "The data source's unique ID.",
			Computed:    true,
		},
	},
	Blocks: map[string]schema.Block{
		"filter": registrationFieldsFilter.Schema(),
		"registration_field": schema.ListNestedBlock{
			Description: "The returned list of registration fields.",
			NestedObject: schema.NestedBlockObject{
				Attributes: fieldSchema,
			},
		},
	},
}

func NewRegistrationField() datasource.DataSource {
	return &RegistrationFieldsDataSource{
		BaseDataSource: NewBaseDataSource(
			BaseDataSourceConfig{
				Name:   REG_FIELD_DATASOURCE,
				Schema: &fieldDataSourceSchema,
			},
		),
	}
}

func (d *RegistrationFieldsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data RegistrationFieldsFilterModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.ID = types.StringValue(uuid.New().String())
	result, diag := registrationFieldsFilter.GetAndFilter(d.Client, data.Filters, listRegistrationFieldss)
	if diag != nil {
		resp.Diagnostics.Append(diag)
		return
	}

	data.RegistrationField = parseModel(
		AnySliceToTyped[cidaas.RegistrationFieldConfig](result),
		parseRegistrationField,
	)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func listRegistrationFieldss(client *cidaas.Client) ([]any, error) {
	rfs, err := client.RegField.GetAll()
	if err != nil {
		return nil, err
	}
	return TypedSliceToAny(rfs), nil
}

func parseRegistrationField(rf cidaas.RegistrationFieldConfig) (result RegistrationField) {
	result.ID = types.StringValue(rf.ID)
	result.FieldKey = types.StringValue(rf.FieldKey)
	result.DataType = types.StringValue(rf.DataType)
	result.FieldType = types.StringValue(rf.FieldType)
	result.Enabled = types.BoolValue(rf.Enabled)
	result.IsGroup = types.BoolValue(rf.IsGroup)
	result.Required = types.BoolValue(rf.Required)
	result.ReadOnly = types.BoolValue(rf.ReadOnly)
	result.Internal = types.BoolValue(rf.Internal)
	result.ParentGroupID = types.StringValue(rf.ParentGroupID)
	return result
}
