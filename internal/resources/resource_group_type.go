// user_group_category is changed to group_type
package resources

import (
	"context"
	"fmt"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/cidaas"
	"github.com/Cidaas/terraform-provider-cidaas/helpers/util"
	"github.com/Cidaas/terraform-provider-cidaas/internal/validators"
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
		MarkdownDescription: "The Group Type, managed through the `cidaas_group_type` resource in the provider defines and configures categories for user groups within the Cidaas system." +
			"\n\n Ensure that the below scopes are assigned to the client with the specified `client_id`:" +
			"\n- cidaas:group_type_read" +
			"\n- cidaas:group_type_write" +
			"\n- cidaas:group_type_delete",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the resource.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"role_mode": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Determines the role mode for the user group type. Allowed values are `any_roles`, `no_roles`, `roles_required` and `allowed_roles`",
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"any_roles", "no_roles", "roles_required", "allowed_roles"}...),
				},
			},
			// TODO: description is a required parameter in admin ui
			"description": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "The `description` attribute provides details about the group type, explaining its purpose.",
			},
			"group_type": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The unique identifier of the group type. This cannot be updated for an existing state.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
				PlanModifiers: []planmodifier.String{
					&validators.UniqueIdentifier{},
				},
			},
			"allowed_roles": schema.SetAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				MarkdownDescription: "List of allowed roles in this group type.",
				Validators: []validator.Set{
					&allowedRolesValidator{},
				},
			},
			"created_at": schema.StringAttribute{
				Computed:    true,
				Description: "The timestamp when the resource was created.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"updated_at": schema.StringAttribute{
				Computed:    true,
				Description: "The timestamp when the resource was last updated.",
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
		resp.Diagnostics.AddError("failed to create group type", util.FormatErrorMessage(err))
		return
	}
	plan.ID = util.StringValueOrNull(&res.Data.ID)
	plan.CreatedAt = util.StringValueOrNull(&res.Data.CreatedTime)
	plan.UpdatedAt = util.StringValueOrNull(&res.Data.UpdatedTime)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *GroupTypeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state GroupTypeConfig
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	res, err := r.cidaasClient.GroupType.Get(state.GroupType.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to read group type", util.FormatErrorMessage(err))
		return
	}
	state.ID = util.StringValueOrNull(&res.Data.ID)
	state.CreatedAt = util.StringValueOrNull(&res.Data.CreatedTime)
	state.UpdatedAt = util.StringValueOrNull(&res.Data.UpdatedTime)
	state.RoleMode = util.StringValueOrNull(&res.Data.RoleMode)
	state.GroupType = util.StringValueOrNull(&res.Data.GroupType)
	state.Description = util.StringValueOrNull(&res.Data.Description)
	state.AllowedRoles = util.SetValueOrNull(res.Data.AllowedRoles)

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *GroupTypeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state GroupTypeConfig
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
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
		resp.Diagnostics.AddError("failed to update group type", util.FormatErrorMessage(err))
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
		resp.Diagnostics.AddError("failed to delete group type", util.FormatErrorMessage(err))
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
