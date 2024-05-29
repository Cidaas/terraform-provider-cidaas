// user_group_category is changed to group_type
package resources

import (
	"context"
	"fmt"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/cidaas"
	"github.com/Cidaas/terraform-provider-cidaas/helpers/util"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type GroupTypeResource struct {
	cidaasClient *cidaas.Client
}

type GroupTypeConfig struct {
	ID           types.String `tfsdk:"id"`
	RoleMode     types.String `tfsdk:"role_mode"`
	Description  types.String `tfsdk:"description"`
	GroupType    types.String `tfsdk:"group_type"`
	AllowedRoles types.Set    `tfsdk:"allowed_roles"`
	CreatedAt    types.String `tfsdk:"created_at"`
	UpdatedAt    types.String `tfsdk:"updated_at"`
}

type allowedRolesValidator struct{}

func NewGroupTypeResource() resource.Resource {
	return &GroupTypeResource{}
}

func (r *GroupTypeResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_group_type"
}

func (r *GroupTypeResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *GroupTypeResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"role_mode": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"any_roles", "no_roles", "roles_required", "allowed_roles"}...),
				},
			},
			"description": schema.StringAttribute{
				Optional: true,
			},
			"group_type": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"allowed_roles": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Validators: []validator.Set{
					&allowedRolesValidator{},
				},
			},
			"created_at": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *GroupTypeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan GroupTypeConfig
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	groupType := cidaas.GroupTypeData{
		RoleMode:    plan.RoleMode.ValueString(),
		GroupType:   plan.GroupType.ValueString(),
		Description: plan.Description.ValueString(),
	}
	diag := plan.AllowedRoles.ElementsAs(ctx, &groupType.AllowedRoles, false)
	resp.Diagnostics.Append(diag...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.cidaasClient.GroupType.Create(groupType)
	if err != nil {
		resp.Diagnostics.AddError("failed to create group type", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	plan.ID = types.StringValue(res.Data.ID)
	plan.CreatedAt = types.StringValue(res.Data.CreatedTime)
	plan.UpdatedAt = types.StringValue(res.Data.UpdatedTime)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *GroupTypeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state GroupTypeConfig
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	res, err := r.cidaasClient.GroupType.Get(state.GroupType.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to read group type", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	state.ID = types.StringValue(res.Data.ID)
	state.CreatedAt = types.StringValue(res.Data.CreatedTime)
	state.UpdatedAt = types.StringValue(res.Data.UpdatedTime)
	state.RoleMode = types.StringValue(res.Data.RoleMode)
	state.GroupType = types.StringValue(res.Data.GroupType)
	state.Description = types.StringValue(res.Data.Description)

	if len(res.Data.AllowedRoles) > 0 {
		allowedRoles, diag := types.SetValueFrom(ctx, types.StringType, res.Data.AllowedRoles)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
		state.AllowedRoles = allowedRoles
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *GroupTypeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state GroupTypeConfig
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if !plan.GroupType.Equal(state.GroupType) {
		resp.Diagnostics.AddError("Unexpected Resource Configuration",
			fmt.Sprintf("Attribute group_type can't be modified. Expected %s, got: %s", state.GroupType, plan.GroupType))
		return
	}
	groupType := cidaas.GroupTypeData{
		RoleMode:    plan.RoleMode.ValueString(),
		GroupType:   plan.GroupType.ValueString(),
		Description: plan.Description.ValueString(),
	}
	diag := plan.AllowedRoles.ElementsAs(ctx, &groupType.AllowedRoles, false)
	resp.Diagnostics.Append(diag...)
	if resp.Diagnostics.HasError() {
		return
	}
	err := r.cidaasClient.GroupType.Update(groupType)
	if err != nil {
		resp.Diagnostics.AddError("failed to update group type", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *GroupTypeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state GroupTypeConfig
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	err := r.cidaasClient.GroupType.Delete(state.GroupType.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to delete scope group", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
}

func (r *GroupTypeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("group_type"), req, resp)
}

func (v allowedRolesValidator) Description(_ context.Context) string {
	return "Validates allowed_roles if required"
}

func (v allowedRolesValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v allowedRolesValidator) ValidateSet(ctx context.Context, req validator.SetRequest, resp *validator.SetResponse) {
	if req.Config.Raw.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	var (
		roleMode     string
		allowedRoles []string
		diags        diag.Diagnostics
	)

	diags = req.ConfigValue.ElementsAs(ctx, &allowedRoles, false)
	resp.Diagnostics.Append(diags...)
	diags = req.Config.GetAttribute(ctx, path.Root("role_mode"), &roleMode)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	allowedRolesRequired := util.StringInSlice(roleMode, []string{"roles_required", "allowed_roles"})
	if allowedRolesRequired && len(allowedRoles) == 0 {
		resp.Diagnostics.AddError("Unexpected Resource Configuration",
			fmt.Sprintf("The attribute allowed_roles cannot be empty when role_mode is set to %s", roleMode))
	} else if !allowedRolesRequired && len(allowedRoles) > 0 {
		resp.Diagnostics.AddError("Unexpected Resource Configuration",
			fmt.Sprintf("The attribute allowed_roles must be empty or omitted when role_mode is set to %s", roleMode))
	}
}
