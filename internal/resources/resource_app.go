package resources

import (
	"context"
	"fmt"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/cidaas"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type AppResource struct {
	cidaasClient *cidaas.Client
}

func NewAppResource() resource.Resource {
	return &AppResource{}
}

func (r *AppResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_app"
}

func (r *AppResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *AppResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan, config AppConfig
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	resp.Diagnostics.Append(plan.ExtractAppConfigs(ctx)...)
	resp.Diagnostics.Append(config.ExtractAppConfigs(ctx)...)

	appModel, diags := prepareAppModel(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.cidaasClient.App.Create(*appModel)
	if err != nil {
		resp.Diagnostics.AddError("failed to create app", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	resp.Diagnostics.Append(updateStateModel(ctx, res, &plan, &config).Diagnostics...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *AppResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state AppConfig
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(state.ExtractAppConfigs(ctx)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.cidaasClient.App.Get(state.ClientID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to read app", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	resp.Diagnostics.Append(updateStateModel(ctx, res, &state, &state).Diagnostics...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *AppResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state, config AppConfig
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	resp.Diagnostics.Append(plan.ExtractAppConfigs(ctx)...)
	appModel, diags := prepareAppModel(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	appModel.ID = state.ID.ValueString()
	err := r.cidaasClient.App.Update(*appModel)
	if err != nil {
		resp.Diagnostics.AddError("failed to update app", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *AppResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state AppConfig
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	err := r.cidaasClient.App.Delete(state.ClientID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to delete app", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
}

func (r *AppResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("client_id"), req, resp)
}

func (r *AppResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {

	if req.Plan.Raw.IsNull() {
		return
	}
	var plan, config AppConfig
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	resp.Diagnostics.Append(config.ExtractAppConfigs(ctx)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(plan.ExtractAppConfigs(ctx)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !config.CommonConfigs.IsNull() {
		if config.AccentColor.IsNull() || plan.AccentColor.Equal(types.StringValue("#ef4923")) {
			if !config.commonConfigs.AccentColor.IsNull() {
				resp.Plan.SetAttribute(ctx, path.Root("accent_color"), config.commonConfigs.AccentColor)
			}
		}
		if config.PrimaryColor.IsNull() || plan.PrimaryColor.Equal(types.StringValue("#f7941d")) {
			if !config.commonConfigs.PrimaryColor.IsNull() {
				resp.Plan.SetAttribute(ctx, path.Root("primary_color"), config.commonConfigs.PrimaryColor)
			}
		}
		if config.MediaType.IsNull() || plan.MediaType.Equal(types.StringValue("IMAGE")) {
			if !config.commonConfigs.MediaType.IsNull() {
				resp.Plan.SetAttribute(ctx, path.Root("media_type"), config.commonConfigs.MediaType)
			}
		}
		if config.HostedPageGroup.IsNull() || plan.HostedPageGroup.Equal(types.StringValue("default")) {
			if !config.commonConfigs.HostedPageGroup.IsNull() {
				resp.Plan.SetAttribute(ctx, path.Root("hosted_page_group"), config.commonConfigs.HostedPageGroup)
			}
		}
		if config.TemplateGroupID.IsNull() || plan.TemplateGroupID.Equal(types.StringValue("default")) {
			if !config.commonConfigs.TemplateGroupID.IsNull() {
				resp.Plan.SetAttribute(ctx, path.Root("template_group_id"), config.commonConfigs.TemplateGroupID)
			}
		}
		if config.BotProvider.IsNull() || plan.BotProvider.Equal(types.StringValue("CIDAAS")) {
			if !config.commonConfigs.BotProvider.IsNull() {
				resp.Plan.SetAttribute(ctx, path.Root("bot_provider"), config.commonConfigs.BotProvider)
			}
		}
		if config.LogoAlign.IsNull() || plan.LogoAlign.Equal(types.StringValue("CENTER")) {
			if !config.commonConfigs.LogoAlign.IsNull() {
				resp.Plan.SetAttribute(ctx, path.Root("logo_align"), config.commonConfigs.LogoAlign)
			}
		}
		if config.Webfinger.IsNull() || plan.Webfinger.Equal(types.StringValue("no_redirection")) {
			if !config.commonConfigs.Webfinger.IsNull() {
				resp.Plan.SetAttribute(ctx, path.Root("webfinger"), config.commonConfigs.Webfinger)
			}
		}
		if config.DefaultMaxAge.IsNull() || plan.DefaultMaxAge.Equal(types.Int64Value(86400)) {
			if !config.commonConfigs.DefaultMaxAge.IsNull() {
				resp.Plan.SetAttribute(ctx, path.Root("default_max_age"), config.commonConfigs.DefaultMaxAge)
			}
		}
		if config.TokenLifetimeInSeconds.IsNull() || plan.TokenLifetimeInSeconds.Equal(types.Int64Value(86400)) {
			if !config.commonConfigs.TokenLifetimeInSeconds.IsNull() {
				resp.Plan.SetAttribute(ctx, path.Root("token_lifetime_in_seconds"), config.commonConfigs.TokenLifetimeInSeconds)
			}
		}
		if config.IDTokenLifetimeInSeconds.IsNull() || plan.IDTokenLifetimeInSeconds.Equal(types.Int64Value(86400)) {
			if !config.commonConfigs.IDTokenLifetimeInSeconds.IsNull() {
				resp.Plan.SetAttribute(ctx, path.Root("id_token_lifetime_in_seconds"), config.commonConfigs.IDTokenLifetimeInSeconds)
			}
		}
		if config.RefreshTokenLifetimeInSeconds.IsNull() || plan.RefreshTokenLifetimeInSeconds.Equal(types.Int64Value(15780000)) {
			if !config.commonConfigs.RefreshTokenLifetimeInSeconds.IsNull() {
				resp.Plan.SetAttribute(ctx, path.Root("refresh_token_lifetime_in_seconds"), config.commonConfigs.RefreshTokenLifetimeInSeconds)
			}
		}
		if config.AllowGuestLogin.IsNull() || plan.AllowGuestLogin.Equal(types.BoolValue(false)) {
			if !config.commonConfigs.AllowGuestLogin.IsNull() {
				resp.Plan.SetAttribute(ctx, path.Root("allow_guest_login"), config.commonConfigs.AllowGuestLogin)
			}
		}
		if config.EnableDeduplication.IsNull() || plan.EnableDeduplication.Equal(types.BoolValue(false)) {
			if !config.commonConfigs.EnableDeduplication.IsNull() {
				resp.Plan.SetAttribute(ctx, path.Root("enable_deduplication"), config.commonConfigs.EnableDeduplication)
			}
		}
		if config.AutoLoginAfterRegister.IsNull() || plan.AutoLoginAfterRegister.Equal(types.BoolValue(false)) {
			if !config.commonConfigs.AutoLoginAfterRegister.IsNull() {
				resp.Plan.SetAttribute(ctx, path.Root("auto_login_after_register"), config.commonConfigs.AutoLoginAfterRegister)
			}
		}
		if config.EnablePasswordlessAuth.IsNull() || plan.EnablePasswordlessAuth.Equal(types.BoolValue(true)) {
			if !config.commonConfigs.EnablePasswordlessAuth.IsNull() {
				resp.Plan.SetAttribute(ctx, path.Root("enable_passwordless_auth"), config.commonConfigs.EnablePasswordlessAuth)
			}
		}
		if config.RegisterWithLoginInformation.IsNull() || plan.RegisterWithLoginInformation.Equal(types.BoolValue(false)) {
			if !config.commonConfigs.RegisterWithLoginInformation.IsNull() {
				resp.Plan.SetAttribute(ctx, path.Root("register_with_login_information"), config.commonConfigs.RegisterWithLoginInformation)
			}
		}
		if config.FdsEnabled.IsNull() || plan.FdsEnabled.Equal(types.BoolValue(true)) {
			if !config.commonConfigs.FdsEnabled.IsNull() {
				resp.Plan.SetAttribute(ctx, path.Root("fds_enabled"), config.commonConfigs.FdsEnabled)
			}
		}
		if config.IsHybridApp.IsNull() || plan.IsHybridApp.Equal(types.BoolValue(false)) {
			if !config.commonConfigs.IsHybridApp.IsNull() {
				resp.Plan.SetAttribute(ctx, path.Root("is_hybrid_app"), config.commonConfigs.IsHybridApp)
			}
		}
		if config.Editable.IsNull() || plan.Editable.Equal(types.BoolValue(true)) {
			if !config.commonConfigs.Editable.IsNull() {
				resp.Plan.SetAttribute(ctx, path.Root("editable"), config.commonConfigs.Editable)
			}
		}
		if config.Enabled.IsNull() || plan.Enabled.Equal(types.BoolValue(true)) {
			if !config.commonConfigs.Enabled.IsNull() {
				resp.Plan.SetAttribute(ctx, path.Root("enabled"), config.commonConfigs.Enabled)
			}
		}
		if config.AlwaysAskMfa.IsNull() || plan.AlwaysAskMfa.Equal(types.BoolValue(false)) {
			if !config.commonConfigs.AlwaysAskMfa.IsNull() {
				resp.Plan.SetAttribute(ctx, path.Root("always_ask_mfa"), config.commonConfigs.AlwaysAskMfa)
			}
		}
		if config.EmailVerificationRequired.IsNull() || plan.EmailVerificationRequired.Equal(types.BoolValue(true)) {
			if !config.commonConfigs.EmailVerificationRequired.IsNull() {
				resp.Plan.SetAttribute(ctx, path.Root("email_verification_required"), config.commonConfigs.EmailVerificationRequired)
			}
		}
		if config.EnableClassicalProvider.IsNull() || plan.EnableClassicalProvider.Equal(types.BoolValue(true)) {
			if !config.commonConfigs.EnableClassicalProvider.IsNull() {
				resp.Plan.SetAttribute(ctx, path.Root("enable_classical_provider"), config.commonConfigs.EnableClassicalProvider)
			}
		}
		if config.IsRememberMeSelected.IsNull() || plan.IsRememberMeSelected.Equal(types.BoolValue(true)) {
			if !config.commonConfigs.IsRememberMeSelected.IsNull() {
				resp.Plan.SetAttribute(ctx, path.Root("is_remember_me_selected"), config.commonConfigs.IsRememberMeSelected)
			}
		}
		if config.ResponseTypes.IsNull() || plan.ResponseTypes.Equal(
			basetypes.NewSetValueMust(types.StringType, []attr.Value{
				types.StringValue("code"), types.StringValue("token"), types.StringValue("id_token"),
			})) {
			if !config.commonConfigs.ResponseTypes.IsNull() {
				resp.Plan.SetAttribute(ctx, path.Root("response_types"), config.commonConfigs.ResponseTypes)
			}
		}
		if config.GrantTypes.IsNull() || plan.GrantTypes.Equal(
			basetypes.NewSetValueMust(types.StringType, []attr.Value{
				types.StringValue("implicit"), types.StringValue("authorization_code"), types.StringValue("password"), types.StringValue("refresh_token"),
			})) {
			if !config.commonConfigs.GrantTypes.IsNull() {
				resp.Plan.SetAttribute(ctx, path.Root("grant_types"), config.commonConfigs.GrantTypes)
			}
		}
		if config.AllowLoginWith.IsNull() || plan.AllowLoginWith.Equal(
			basetypes.NewSetValueMust(types.StringType, []attr.Value{
				types.StringValue("EMAIL"), types.StringValue("MOBILE"), types.StringValue("USER_NAME"),
			})) {
			if !config.commonConfigs.AllowLoginWith.IsNull() {
				resp.Plan.SetAttribute(ctx, path.Root("allow_login_with"), config.commonConfigs.AllowLoginWith)
			}
		}
		if config.Mfa.IsNull() || plan.Mfa.Equal(types.ObjectValueMust(
			map[string]attr.Type{
				"setting":                  types.StringType,
				"time_interval_in_seconds": types.Int64Type,
				"allowed_methods": types.SetType{
					ElemType: types.StringType,
				},
			},
			map[string]attr.Value{
				"setting":                  types.StringValue("OFF"),
				"time_interval_in_seconds": types.Int64Null(),
				"allowed_methods":          types.SetNull(types.StringType),
			})) {
			if !config.commonConfigs.Mfa.IsNull() {
				resp.Plan.SetAttribute(ctx, path.Root("mfa"), config.commonConfigs.Mfa)
			}
		}
	}
}
