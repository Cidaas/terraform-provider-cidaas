package datasources

import (
	"context"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/cidaas"
	"github.com/Cidaas/terraform-provider-cidaas/helpers/util"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ScopesDataSource struct {
	BaseDataSource
}

type ScopesFilterModel struct {
	BaseModel
	Scope []Scope `tfsdk:"scope"`
}

type Scope struct {
	ID                   types.String `tfsdk:"id"`
	SecurityLevel        types.String `tfsdk:"security_level"`
	ScopeKey             types.String `tfsdk:"scope_key"`
	GroupName            types.Set    `tfsdk:"group_name"`
	RequiredUserConsent  types.Bool   `tfsdk:"required_user_consent"`
	LocalizedDescription types.List   `tfsdk:"localized_descriptions"`
	ScopeOwner           types.String `tfsdk:"scope_owner"`
}

type LocalDescription struct {
	Locale      types.String `tfsdk:"locale"`
	Title       types.String `tfsdk:"title"`
	Description types.String `tfsdk:"description"`
}

var scopeFilter = FilterConfig{
	"scope_key":             {TypeFunc: FilterTypeString},
	"security_level":        {TypeFunc: FilterTypeString},
	"group_name":            {TypeFunc: FilterTypeString},
	"required_user_consent": {TypeFunc: FilterTypeBool},
}

var scopeSchema = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:    true,
		Description: "The ID of the scope.",
	},
	"security_level": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The security level of the scope, `PUBLIC` or `CONFIDENTIAL`.",
	},
	"scope_key": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Unique identifier(name) for the scope.",
	},
	"group_name": schema.SetAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "List of scope_groups associated with the scope.",
	},
	"required_user_consent": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Indicates whether user consent is required for the scope.",
	},
	"scope_owner": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The owner of the scope. e.g. `ADMIN`.",
	},
	"localized_descriptions": schema.ListNestedAttribute{
		Optional: true,
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"locale": schema.StringAttribute{
					Computed:            true,
					MarkdownDescription: "The locale for the scope, e.g., `en-US`.",
				},
				"title": schema.StringAttribute{
					Computed:            true,
					MarkdownDescription: "The title of the scope in the configured locale.",
				},
				"description": schema.StringAttribute{
					Computed:            true,
					MarkdownDescription: "The description of the scope in the configured locale.",
				},
			},
		},
	},
}

var scopesDataSourceSchema = schema.Schema{
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Description: "The data source's unique ID.",
			Computed:    true,
		},
	},
	Blocks: map[string]schema.Block{
		"filter": scopeFilter.Schema(),
		"scope": schema.ListNestedBlock{
			Description: "The returned list of scopes.",
			NestedObject: schema.NestedBlockObject{
				Attributes: scopeSchema,
			},
		},
	},
}

func NewScope() datasource.DataSource {
	return &ScopesDataSource{
		BaseDataSource: NewBaseDataSource(
			BaseDataSourceConfig{
				Name:   "cidaas_scope",
				Schema: &scopesDataSourceSchema,
			},
		),
	}
}

func (d *ScopesDataSource) Read(
	ctx context.Context,
	req datasource.ReadRequest,
	resp *datasource.ReadResponse,
) {
	var data ScopesFilterModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.ID = types.StringValue(uuid.New().String())
	result, diag := scopeFilter.GetAndFilter(d.Client, data.Filters, listScopes)
	if diag != nil {
		resp.Diagnostics.Append(diag)
		return
	}

	data.Scope = parseModel(AnySliceToTyped[cidaas.ScopeModel](result), parseScope)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func listScopes(client *cidaas.Client) ([]any, error) {
	scopes, err := client.Scope.GetAll()
	if err != nil {
		return nil, err
	}
	return TypedSliceToAny(scopes), nil
}

func parseScope(scope cidaas.ScopeModel) (r Scope) {
	r.ID = types.StringValue(scope.ID)
	r.ScopeKey = types.StringValue(scope.ScopeKey)
	r.GroupName = util.SetValueOrNull(scope.GroupName)
	r.SecurityLevel = types.StringValue(scope.SecurityLevel)
	r.RequiredUserConsent = types.BoolValue(scope.RequiredUserConsent)

	var objectValues []attr.Value
	localeDescription := types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"locale":      types.StringType,
			"title":       types.StringType,
			"description": types.StringType,
		},
	}

	for _, sc := range scope.LocaleWiseDescription {
		local := sc.Locale
		title := sc.Title
		description := sc.Description
		objValue := types.ObjectValueMust(localeDescription.AttrTypes, map[string]attr.Value{
			"locale":      util.StringValueOrNull(&local),
			"title":       util.StringValueOrNull(&title),
			"description": util.StringValueOrNull(&description),
		})
		objectValues = append(objectValues, objValue)
	}
	r.LocalizedDescription = types.ListValueMust(localeDescription, objectValues)
	return r
}
