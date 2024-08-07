package resources

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/cidaas"
	"github.com/Cidaas/terraform-provider-cidaas/helpers/util"
	"github.com/Cidaas/terraform-provider-cidaas/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ConsentResource struct {
	cidaasClient *cidaas.Client
}

type ConsentConfig struct {
	ID             types.String `tfsdk:"id"`
	ConsentGroupID types.String `tfsdk:"consent_group_id"`
	Name           types.String `tfsdk:"name"`
	Enabled        types.Bool   `tfsdk:"enabled"`
	CreatedAt      types.String `tfsdk:"created_at"`
	UpdatedAt      types.String `tfsdk:"updated_at"`
}

func NewConsentResource() resource.Resource {
	return &ConsentResource{}
}

func (r *ConsentResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_consent"
}

func (r *ConsentResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(*cidaas.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected cidaas.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}
	r.cidaasClient = client
}

func (r *ConsentResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
				PlanModifiers: []planmodifier.String{
					&validators.UniqueIdentifier{},
				},
			},
			"consent_group_id": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
				PlanModifiers: []planmodifier.String{
					&validators.UniqueIdentifier{},
				},
			},
			"enabled": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(true),
			},
			"created_at": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *ConsentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ConsentConfig
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	consent := cidaas.ConsentModel{
		ConsentName:    plan.Name.ValueString(),
		ConsentGroupID: plan.ConsentGroupID.ValueString(),
		Enabled:        plan.Enabled.ValueBool(),
	}
	res, err := r.cidaasClient.Consent.Upsert(consent)
	if err != nil {
		resp.Diagnostics.AddError("failed to create consent", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	plan.ID = util.StringValueOrNull(&res.Data.ID)
	plan.CreatedAt = util.StringValueOrNull(&res.Data.CreatedTime)
	plan.UpdatedAt = util.StringValueOrNull(&res.Data.UpdatedTime)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ConsentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ConsentConfig
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	res, err := r.cidaasClient.Consent.GetConsentInstances(state.ConsentGroupID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Failed to read consent", fmt.Sprintf("Error: %s ", err.Error()))
		return
	}
	if !res.Success && res.Status == http.StatusNoContent {
		resp.Diagnostics.AddError("Invalid consent_group_id", fmt.Sprintf("No consent group found for the provided consent_group_id %+v", state.ConsentGroupID.String()))
		return
	}
	isAvailable := false
	if len(res.Data) > 0 {
		for _, instance := range res.Data {
			if instance.ConsentName == state.Name.ValueString() {
				isAvailable = true
				state.ID = util.StringValueOrNull(&instance.ID)
				state.Name = util.StringValueOrNull(&instance.ConsentName)
				state.Enabled = types.BoolValue(instance.Enabled)
				state.CreatedAt = util.StringValueOrNull(&instance.CreatedTime)
				state.UpdatedAt = util.StringValueOrNull(&instance.UpdatedTime)
			}
		}
	}

	if !isAvailable {
		resp.Diagnostics.AddError("Consent not found", fmt.Sprintf("consent %s not found for the provided consent_group_id %s", state.Name.String(), state.ConsentGroupID.String()))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *ConsentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ConsentConfig
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	consent := cidaas.ConsentModel{
		ID:             state.ID.ValueString(),
		ConsentName:    plan.Name.ValueString(),
		ConsentGroupID: plan.ConsentGroupID.ValueString(),
		Enabled:        plan.Enabled.ValueBool(),
	}
	res, err := r.cidaasClient.Consent.Upsert(consent)
	if err != nil {
		resp.Diagnostics.AddError("failed to update consent", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	plan.UpdatedAt = util.StringValueOrNull(&res.Data.UpdatedTime)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ConsentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ConsentConfig
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	err := r.cidaasClient.Consent.Delete(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to delete consent", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
}

func (r *ConsentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	id := req.ID
	parts := strings.Split(id, ":")
	if len(parts) != 2 {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: 'consent_group_id:name', got: %s", id),
		)
		return
	}
	resp.State.SetAttribute(ctx, path.Root("consent_group_id"), parts[0])
	resp.State.SetAttribute(ctx, path.Root("name"), parts[1])
}
