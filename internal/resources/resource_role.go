package resources

import (
	"context"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/cidaas"
	"github.com/Cidaas/terraform-provider-cidaas/helpers/util"
	"github.com/Cidaas/terraform-provider-cidaas/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RoleResource struct {
	BaseResource
}

func NewRoleResource() resource.Resource {
	return &RoleResource{
		BaseResource: NewBaseResource(
			BaseResourceConfig{
				Name:   RESOURCE_ROLE,
				Schema: &roleSchema,
			},
		),
	}
}

type Role struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Role        types.String `tfsdk:"role"`
}

var roleSchema = schema.Schema{
	MarkdownDescription: "The cidaas_role resource in Terraform facilitates the management of roles in Cidaas system." +
		" This resource allows you to configure and define custom roles to suit your application's specific access control requirements." +
		"\n\n Ensure that the below scopes are assigned to the client with the specified `client_id`:" +
		"\n- cidaas:roles_read" +
		"\n- cidaas:roles_write" +
		"\n- cidaas:roles_delete",
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:    true,
			Description: "The ID of the role resource.",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"role": schema.StringAttribute{
			Required:    true,
			Description: "The unique identifier of the role. The role name must be unique across the cidaas system and cannot be updated for an existing state.",
			PlanModifiers: []planmodifier.String{
				&validators.UniqueIdentifier{},
			},
		},
		"name": schema.StringAttribute{
			Optional:    true,
			Description: "The name of the role.",
		},
		"description": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: "The `description` attribute provides details about the role, explaining its purpose.",
		},
	},
}

func (r *RoleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan Role
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	role := cidaas.RoleModel{
		Name:        plan.Name.ValueString(),
		Role:        plan.Role.ValueString(),
		Description: plan.Description.ValueString(),
	}
	response, err := r.cidaasClient.Roles.UpsertRole(ctx, role)
	if err != nil {
		resp.Diagnostics.AddError("failed to create role", util.FormatErrorMessage(err))
		return
	}
	plan.ID = util.StringValueOrNull(&response.Data.Role)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state Role
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	response, err := r.cidaasClient.Roles.GetRole(ctx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to read role", util.FormatErrorMessage(err))
		return
	}
	state.Role = util.StringValueOrNull(&response.Data.Role)
	state.Description = util.StringValueOrNull(&response.Data.Description)
	state.Name = util.StringValueOrNull(&response.Data.Name)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *RoleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state Role
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	role := cidaas.RoleModel{
		Name:        plan.Name.ValueString(),
		Role:        plan.Role.ValueString(),
		Description: plan.Description.ValueString(),
	}
	response, err := r.cidaasClient.Roles.UpsertRole(ctx, role)
	if err != nil {
		resp.Diagnostics.AddError("failed to update role", util.FormatErrorMessage(err))
		return
	}
	plan.ID = util.StringValueOrNull(&response.Data.Role)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state Role
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	err := r.cidaasClient.Roles.DeleteRole(ctx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to delete role", util.FormatErrorMessage(err))
		return
	}
}

func (r *RoleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
