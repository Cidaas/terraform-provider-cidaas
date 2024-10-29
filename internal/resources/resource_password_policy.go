package resources

import (
	"context"
	"fmt"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/cidaas"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PasswordPolicy struct {
	BaseResource
}

func NewPasswordPolicy() resource.Resource {
	return &PasswordPolicy{
		BaseResource: NewBaseResource(
			BaseResourceConfig{
				Name:   RESOURCE_PASSWORD_POLICY,
				Schema: &passwordPolicySchema,
			},
		),
	}
}

type PasswordPolicyConfig struct {
	ID                types.String `tfsdk:"id"`
	PolicyName        types.String `tfsdk:"policy_name"`
	MaximumLength     types.Int64  `tfsdk:"maximum_length"`
	MinimumLength     types.Int64  `tfsdk:"minimum_length"`
	NoOfSpecialChars  types.Int64  `tfsdk:"no_of_special_chars"`
	NoOfDigits        types.Int64  `tfsdk:"no_of_digits"`
	LowerAndUppercase types.Bool   `tfsdk:"lower_and_uppercase"`
}

var passwordPolicySchema = schema.Schema{
	MarkdownDescription: "The Password Policy resource in the provider allows you to manage the password policy within the Cidaas." +
		"\n\n Ensure that the below scopes are assigned to the client with the specified `client_id`:" +
		"\n- cidaas:password_policy_read" +
		"\n- cidaas:password_policy_write" +
		"\n- cidaas:password_policy_delete",
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Unique identifier of the password policy.",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"policy_name": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The name of the password policy.",
		},
		"maximum_length": schema.Int64Attribute{
			Required:            true,
			MarkdownDescription: "The maximum length allowed for the password. The `maximum_length` must be at least sum of `minimum_length`, `no_of_special_chars`, `no_of_digits` and `lower_and_uppercase(1)` ",
			Validators: []validator.Int64{
				int64validator.AtLeastSumOf(
					path.MatchRoot("minimum_length"),
					path.MatchRoot("no_of_special_chars"),
					path.MatchRoot("no_of_digits"),
				),
			},
		},
		"minimum_length": schema.Int64Attribute{
			Required:            true,
			MarkdownDescription: "The minimum length required for the password. The `minimum_length` must be greater than or equal to 5.",
			Validators: []validator.Int64{
				int64validator.AtLeast(5),
			},
		},
		"no_of_special_chars": schema.Int64Attribute{
			Required:            true,
			MarkdownDescription: "The required number of special characters in the password.",
		},
		"no_of_digits": schema.Int64Attribute{
			Required:            true,
			MarkdownDescription: "The required number of digits in the password.",
		},
		"lower_and_uppercase": schema.BoolAttribute{
			Required:            true,
			MarkdownDescription: "Specifies whether the password must contain both lowercase and uppercase letters.",
		},
	},
}

func (r *PasswordPolicy) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan PasswordPolicyConfig
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	payload := cidaas.PasswordPolicyModel{
		PolicyName:        plan.PolicyName.ValueString(),
		MaximumLength:     plan.MaximumLength.ValueInt64(),
		MinimumLength:     plan.MinimumLength.ValueInt64(),
		NoOfSpecialChars:  plan.NoOfSpecialChars.ValueInt64(),
		NoOfDigits:        plan.NoOfDigits.ValueInt64(),
		LowerAndUppercase: plan.LowerAndUppercase.ValueBool(),
	}

	res, err := r.cidaasClient.PasswordPolicy.Upsert(payload)
	if err != nil {
		resp.Diagnostics.AddError("failed to create password policy", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	plan.ID = types.StringValue(res.Data.ID)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *PasswordPolicy) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state PasswordPolicyConfig
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	res, err := r.cidaasClient.PasswordPolicy.Get(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to read password policy", fmt.Sprintf("Error: %s", err.Error()))
		return
	}

	state.PolicyName = types.StringValue(res.Data.PolicyName)
	state.MaximumLength = types.Int64Value(res.Data.MaximumLength)
	state.MinimumLength = types.Int64Value(res.Data.MinimumLength)
	state.NoOfSpecialChars = types.Int64Value(res.Data.NoOfSpecialChars)
	state.NoOfDigits = types.Int64Value(res.Data.NoOfDigits)
	state.LowerAndUppercase = types.BoolValue(res.Data.LowerAndUppercase)

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *PasswordPolicy) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state PasswordPolicyConfig
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	payload := cidaas.PasswordPolicyModel{
		ID:                state.ID.ValueString(),
		PolicyName:        plan.PolicyName.ValueString(),
		MaximumLength:     plan.MaximumLength.ValueInt64(),
		MinimumLength:     plan.MinimumLength.ValueInt64(),
		NoOfSpecialChars:  plan.NoOfSpecialChars.ValueInt64(),
		NoOfDigits:        plan.NoOfDigits.ValueInt64(),
		LowerAndUppercase: plan.LowerAndUppercase.ValueBool(),
	}

	_, err := r.cidaasClient.PasswordPolicy.Upsert(payload)
	if err != nil {
		resp.Diagnostics.AddError("failed to update password policy", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *PasswordPolicy) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state PasswordPolicyConfig
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.cidaasClient.PasswordPolicy.Delete(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to delete password policy", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
}

func (r *PasswordPolicy) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
