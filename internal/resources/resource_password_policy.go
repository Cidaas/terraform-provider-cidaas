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
	cidaasClient *cidaas.Client
}

type PasswordPolicyConfig struct {
	ID                types.String `tfsdk:"id"`
	MaximumLength     types.Int64  `tfsdk:"maximum_length"`
	MinimumLength     types.Int64  `tfsdk:"minimum_length"`
	NoOfSpecialChars  types.Int64  `tfsdk:"no_of_special_chars"`
	NoOfDigits        types.Int64  `tfsdk:"no_of_digits"`
	LowerAndUppercase types.Bool   `tfsdk:"lower_and_uppercase"`
	ReuseLimit        types.Int64  `tfsdk:"reuse_limit"`
	ExpirationInDays  types.Int64  `tfsdk:"expiration_in_days"`
	NoOfDaysToRemind  types.Int64  `tfsdk:"no_of_days_to_remind"`
}

func NewPasswordPolicy() resource.Resource {
	return &PasswordPolicy{}
}

func (r *PasswordPolicy) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_password_policy"
}

func (r *PasswordPolicy) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *PasswordPolicy) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) { //nolint:gocognit
	if req.Plan.Raw.IsNull() {
		return
	}
	var plan PasswordPolicyConfig
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if plan.ID.IsUnknown() {
		resp.Diagnostics.AddError(
			"Resource Creation Not Allowed",
			"Creating this resource using 'terraform apply' is not allowed. You must first import the existing resource using 'terraform import'. After the import, you can perform updates as needed.",
		)
	}
}

func (r *PasswordPolicy) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The Password Policy resource in the provider allows you to manage the password policy within the Cidaas." +
			"\nNote that resource creation is not allowed, only updates are permitted after the resource has been imported." +
			"\n\n Ensure that the below scopes are assigned to the client with the specified `client_id`:" +
			"\n- cidaas:password_policy_read" +
			"\n- cidaas:password_policy_write",
		Attributes: map[string]schema.Attribute{
			// The id is used to determine the operation type.
			// If id is unknown, it indicates a create operation. Otherwise, it indicates an update operation.
			// Note: Creation is not allowed, only updates are permitted after importing the resource.
			"id": schema.StringAttribute{
				Computed: true,
				MarkdownDescription: "Unique identifier of the password policy. This will be set to the same value as the import identifier." +
					"\nWhile the cidaas API does not require an identifier to import password policy, Terraform's import command does. Therefore, you can provide any arbitrary string as the identifier.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"maximum_length": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "The maximum length allowed for the password. The `maximum_length` must be greater than `minimum_length`",
				Validators: []validator.Int64{
					int64validator.AtLeastSumOf(path.MatchRoot("minimum_length")),
				},
			},
			"minimum_length": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "The minimum length required for the password. The `minimum_length` must be greater than or equal to the sum of `no_of_special_chars`, `no_of_digits`, and `lowercase/uppercase` characters.",
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
			"reuse_limit": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "The number of previous passwords that cannot be reused. This number cannot exceed 5.",
				Validators: []validator.Int64{
					int64validator.AtMost(5),
				},
			},
			"expiration_in_days": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "The number of days after which the password expires.",
			},
			"no_of_days_to_remind": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "The number of days before the password expiration to remind the user to change their password.",
			},
		},
	}
}

func (r *PasswordPolicy) Create(_ context.Context, _ resource.CreateRequest, resp *resource.CreateResponse) {
	resp.Diagnostics.AddError(
		"Resource Creation Not Allowed",
		"Creating this resource using 'terraform apply' is not allowed. You must first import the existing resource using 'terraform import'. After the import, you can perform updates as needed.",
	)
}

func (r *PasswordPolicy) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state PasswordPolicyConfig
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	res, err := r.cidaasClient.PasswordPolicy.Get()
	if err != nil {
		resp.Diagnostics.AddError("failed to read password policy", fmt.Sprintf("Error: %s", err.Error()))
		return
	}

	state.MaximumLength = types.Int64Value(res.Data.MaximumLength)
	state.MinimumLength = types.Int64Value(res.Data.MinimumLength)
	state.NoOfSpecialChars = types.Int64Value(res.Data.NoOfSpecialChars)
	state.NoOfDigits = types.Int64Value(res.Data.NoOfDigits)
	state.LowerAndUppercase = types.BoolValue(res.Data.LowerAndUppercase)
	state.ReuseLimit = types.Int64Value(res.Data.ReuseLimit)
	state.ExpirationInDays = types.Int64Value(res.Data.ExpirationInDays)
	state.NoOfDaysToRemind = types.Int64Value(res.Data.NoOfDaysToRemind)

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
		MaximumLength:     plan.MaximumLength.ValueInt64(),
		MinimumLength:     plan.MinimumLength.ValueInt64(),
		NoOfSpecialChars:  plan.NoOfSpecialChars.ValueInt64(),
		NoOfDigits:        plan.NoOfDigits.ValueInt64(),
		LowerAndUppercase: plan.LowerAndUppercase.ValueBool(),
		ReuseLimit:        plan.ReuseLimit.ValueInt64(),
		ExpirationInDays:  plan.ExpirationInDays.ValueInt64(),
		NoOfDaysToRemind:  plan.NoOfDaysToRemind.ValueInt64(),
	}

	err := r.cidaasClient.PasswordPolicy.Update(payload)
	if err != nil {
		resp.Diagnostics.AddError("failed to update password policy", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *PasswordPolicy) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
}

func (r *PasswordPolicy) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
