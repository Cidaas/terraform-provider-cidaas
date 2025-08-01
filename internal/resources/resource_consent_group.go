package resources

import (
	"context"
	"fmt"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/cidaas"
	"github.com/Cidaas/terraform-provider-cidaas/helpers/util"
	"github.com/Cidaas/terraform-provider-cidaas/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ConsentGroupResource struct {
	BaseResource
}

func NewConsentGroupResource() resource.Resource {
	return &ConsentGroupResource{
		BaseResource: NewBaseResource(
			BaseResourceConfig{
				Name:   RESOURCE_CONSENT_GROUP,
				Schema: &consentGroupSchema,
			},
		),
	}
}

type ConsentGroupConfig struct {
	ID          types.String `tfsdk:"id"`
	GroupName   types.String `tfsdk:"group_name"`
	Description types.String `tfsdk:"description"`
	CreatedAt   types.String `tfsdk:"created_at"`
	UpdatedAt   types.String `tfsdk:"updated_at"`
}

var consentGroupSchema = schema.Schema{
	MarkdownDescription: "The Consent Group resource in the provider allows you to define and manage consent groups in Cidaas." +
		"\n Consent Groups are useful to organize and manage consents by grouping related consent items together." +
		"\n\n Ensure that the below scopes are assigned to the client with the specified `client_id`:" +
		"\n- cidaas:tenant_consent_read" +
		"\n- cidaas:tenant_consent_write" +
		"\n- cidaas:tenant_consent_delete",
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The unique identifier of the consent group.",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"group_name": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The name of the consent group.",
			Validators: []validator.String{
				stringvalidator.LengthAtLeast(1),
			},
			PlanModifiers: []planmodifier.String{
				&validators.UniqueIdentifier{},
			},
		},
		"description": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: "Description of the consent group.",
		},
		"created_at": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The timestamp when the consent group was created.",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"updated_at": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The timestamp when the consent group was last updated.",
		},
	},
}

func (r *ConsentGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ScopeGroupConfig
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	consentGroup := cidaas.ConsentGroupConfig{
		GroupName:   plan.GroupName.ValueString(),
		Description: plan.Description.ValueString(),
	}
	res, err := r.cidaasClient.ConsentGroup.Upsert(ctx, consentGroup)
	if err != nil {
		resp.Diagnostics.AddError("failed to create consent group", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	plan.ID = util.StringValueOrNull(&res.Data.ID)
	plan.GroupName = util.StringValueOrNull(&res.Data.GroupName)
	plan.CreatedAt = util.StringValueOrNull(&res.Data.CreatedTime)
	plan.UpdatedAt = util.StringValueOrNull(&res.Data.UpdatedTime)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ConsentGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) { //nolint:dupl
	var state ConsentGroupConfig
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	res, err := r.cidaasClient.ConsentGroup.Get(ctx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to read consent group", fmt.Sprintf("Error: %s ", err.Error()))
		return
	}
	state.ID = util.StringValueOrNull(&res.Data.ID)
	state.GroupName = util.StringValueOrNull(&res.Data.GroupName)
	state.CreatedAt = util.StringValueOrNull(&res.Data.CreatedTime)
	state.UpdatedAt = util.StringValueOrNull(&res.Data.UpdatedTime)
	state.Description = util.StringValueOrNull(&res.Data.Description)

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *ConsentGroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ConsentGroupConfig
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	consentGroup := cidaas.ConsentGroupConfig{
		GroupName:   plan.GroupName.ValueString(),
		Description: plan.Description.ValueString(),
	}
	res, err := r.cidaasClient.ConsentGroup.Upsert(ctx, consentGroup)
	if err != nil {
		resp.Diagnostics.AddError("failed to update consent group", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	plan.UpdatedAt = util.StringValueOrNull(&res.Data.UpdatedTime)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ConsentGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ConsentGroupConfig
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	err := r.cidaasClient.ConsentGroup.Delete(ctx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to delete consent group", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
}

func (r *ConsentGroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
