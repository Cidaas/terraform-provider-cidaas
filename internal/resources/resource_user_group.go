package resources

import (
	"context"
	"fmt"
	"regexp"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/cidaas"
	"github.com/Cidaas/terraform-provider-cidaas/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type UserGroupResource struct {
	cidaasClient *cidaas.Client
}

type UserGroupConfig struct {
	ID                          types.String `tfsdk:"id"`
	GroupType                   types.String `tfsdk:"group_type"`
	GroupID                     types.String `tfsdk:"group_id"`
	GroupName                   types.String `tfsdk:"group_name"`
	ParentID                    types.String `tfsdk:"parent_id"`
	LogoURL                     types.String `tfsdk:"logo_url"`
	Description                 types.String `tfsdk:"description"`
	MakeFirstUserAdmin          types.Bool   `tfsdk:"make_first_user_admin"`
	MemberProfileVisibility     types.String `tfsdk:"member_profile_visibility"`
	NoneMemberProfileVisibility types.String `tfsdk:"none_member_profile_visibility"`
	CustomFields                types.Map    `tfsdk:"custom_fields"`
	CreatedAt                   types.String `tfsdk:"created_at"`
	UpdatedAt                   types.String `tfsdk:"updated_at"`
}

func NewUserGroupResource() resource.Resource {
	return &UserGroupResource{}
}

func (r *UserGroupResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user_groups"
}

func (r *UserGroupResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *UserGroupResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"group_type": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"group_id": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
				// ValidateFunc: validation.StringDoesNotContainAny(" "),
				PlanModifiers: []planmodifier.String{
					&validators.UniqueIdentifier{},
				},
			},
			"group_name": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"parent_id": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("root"),
			},
			"logo_url": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^https://.+$`),
						"must be a valid URL starting with https://",
					),
				},
			},
			"description": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.LengthAtMost(256),
				},
			},
			"make_first_user_admin": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
			},
			"member_profile_visibility": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"public", "full"}...),
				},
				Default: stringdefault.StaticString("public"),
			},
			"none_member_profile_visibility": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"none", "public"}...),
				},
				Default: stringdefault.StaticString("none"),
			},
			"custom_fields": schema.MapAttribute{
				ElementType: types.StringType,
				Optional:    true,
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

func (r *UserGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan UserGroupConfig
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	userGroup, diags := prepareUserGroupPayload(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.cidaasClient.UserGroup.Create(*userGroup)
	if err != nil {
		resp.Diagnostics.AddError("failed to create user group", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	plan.ID = types.StringValue(res.Data.ID)
	plan.CreatedAt = types.StringValue(res.Data.CreatedTime)
	plan.UpdatedAt = types.StringValue(res.Data.UpdatedTime)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *UserGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state UserGroupConfig
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.cidaasClient.UserGroup.Get(state.GroupID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to read user group", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	state.ID = types.StringValue(res.Data.ID)
	state.GroupType = types.StringValue(res.Data.GroupType)
	state.GroupID = types.StringValue(res.Data.GroupID)
	state.GroupName = types.StringValue(res.Data.GroupName)
	state.ParentID = types.StringValue(res.Data.ParentID)
	state.LogoURL = types.StringValue(res.Data.LogoURL)
	state.Description = types.StringValue(res.Data.Description)
	state.MakeFirstUserAdmin = types.BoolValue(res.Data.MakeFirstUserAdmin)
	state.MemberProfileVisibility = types.StringValue(res.Data.MemberProfileVisibility)
	state.NoneMemberProfileVisibility = types.StringValue(res.Data.NoneMemberProfileVisibility)
	state.CreatedAt = types.StringValue(res.Data.CreatedTime)
	state.UpdatedAt = types.StringValue(res.Data.UpdatedTime)

	cfAttributes := map[string]attr.Value{}
	for key, value := range res.Data.CustomFields {
		cfAttributes[key] = types.StringValue(value)
	}
	cf, d := types.MapValue(types.StringType, cfAttributes)
	resp.Diagnostics.Append(d...)
	if resp.Diagnostics.HasError() {
		return
	}
	state.CustomFields = cf
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *UserGroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state UserGroupConfig
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	userGroup, diags := prepareUserGroupPayload(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	err := r.cidaasClient.UserGroup.Update(*userGroup)
	if err != nil {
		resp.Diagnostics.AddError("failed to update user group", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *UserGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state UserGroupConfig
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	subGroups, err := r.cidaasClient.UserGroup.GetSubGroups(state.GroupID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to check sub user groups", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	if len(subGroups) > 0 {
		resp.Diagnostics.AddError("Invalid Request", "The group contains sub user groups and cannot be deleted.")
		return
	}
	err = r.cidaasClient.UserGroup.Delete(state.GroupID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to delete user group", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
}

func (r *UserGroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("group_id"), req, resp)
}

func prepareUserGroupPayload(ctx context.Context, plan UserGroupConfig) (*cidaas.UserGroupData, diag.Diagnostics) {
	userGroup := cidaas.UserGroupData{
		GroupType:                   plan.GroupType.ValueString(),
		GroupID:                     plan.GroupID.ValueString(),
		GroupName:                   plan.GroupName.ValueString(),
		ParentID:                    plan.ParentID.ValueString(),
		LogoURL:                     plan.LogoURL.ValueString(),
		Description:                 plan.Description.ValueString(),
		MakeFirstUserAdmin:          plan.MakeFirstUserAdmin.ValueBool(),
		MemberProfileVisibility:     plan.MemberProfileVisibility.ValueString(),
		NoneMemberProfileVisibility: plan.NoneMemberProfileVisibility.ValueString(),
	}
	diags := plan.CustomFields.ElementsAs(ctx, &userGroup.CustomFields, false)
	if diags.HasError() {
		return nil, diags
	}
	return &userGroup, nil
}
