package datasources

import (
	"context"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/cidaas"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ConsentDataSource struct {
	BaseDataSource
}

type ConsentFilterModel struct {
	BaseModel
	Consent []Consent `tfsdk:"consent"`
}

type Consent struct {
	ID          types.String `tfsdk:"id"`
	ConsentName types.String `tfsdk:"consent_name"`
}

var consentFilter = FilterConfig{
	"consent_name": {TypeFunc: FilterTypeString},
}

func NewConsent() datasource.DataSource {
	return &ConsentDataSource{
		BaseDataSource: NewBaseDataSource(
			BaseDataSourceConfig{
				Name:   "cidaas_consent",
				Schema: &consentSchema,
			},
		),
	}
}

var consentSchema = schema.Schema{
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Description: "The data source's unique ID.",
			Computed:    true,
		},
	},
	Blocks: map[string]schema.Block{
		"filter": consentFilter.Schema(),
		"consent": schema.ListNestedBlock{
			Description: "The returned list of consents.",
			NestedObject: schema.NestedBlockObject{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "The unique identifier of the consent.",
					},
					"consent_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "The name of the consent.",
					},
				},
			},
		},
	},
}

func (d *ConsentDataSource) Read(
	ctx context.Context,
	req datasource.ReadRequest,
	resp *datasource.ReadResponse,
) {
	var data ConsentFilterModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.ID = types.StringValue(uuid.New().String())
	result, diag := consentFilter.GetAndFilter(d.Client, data.Filters, listConsents)
	if diag != nil {
		resp.Diagnostics.Append(diag)
		return
	}

	data.Consent = parseModel[cidaas.ConsentModel, Consent](AnySliceToTyped[cidaas.ConsentModel](result), parseConsent)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func listConsents(client *cidaas.Client) ([]any, error) {
	consents, err := client.Consent.GetAll()
	if err != nil {
		return nil, err
	}
	return TypedSliceToAny(consents), nil
}

func parseConsent(c cidaas.ConsentModel) (result Consent) {
	result.ID = types.StringValue(c.ID)
	result.ConsentName = types.StringValue(c.ConsentName)
	return result
}
