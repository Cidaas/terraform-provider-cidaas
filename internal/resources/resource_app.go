package resources

import (
	"context"
	"fmt"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/cidaas"
	"github.com/Cidaas/terraform-provider-cidaas/helpers/util"
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
	resp.Diagnostics.Append(plan.ExtractAppConfigs(ctx)...)
	appModel, diags := prepareAppModel(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.cidaasClient.App.Create(*appModel)
	if err != nil {
		resp.Diagnostics.AddError("failed to create app", fmt.Sprintf("Error: %+v", err.Error()))
		return
	}
	resp.Diagnostics.Append(updateStateModel(res, &plan, &config, CREATE)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *AppResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state, config AppConfig
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
	operation := IMPORT
	if !state.ID.IsNull() {
		operation = READ
	}
	resp.Diagnostics.Append(updateStateModel(*res, &state, &config, operation)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *AppResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state, config AppConfig

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(plan.ExtractAppConfigs(ctx)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	resp.Diagnostics.Append(config.ExtractAppConfigs(ctx)...)

	appModel, diags := prepareAppModel(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	appModel.ID = state.ID.ValueString()
	res, err := r.cidaasClient.App.Update(*appModel)
	if err != nil {
		resp.Diagnostics.AddError("failed to update app", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	resp.Diagnostics.Append(updateStateModel(res, &plan, &config, UPDATE)...)
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

func (r *AppResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) { //nolint:gocognit
	var config AppConfig
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	resp.Diagnostics.Append(config.ExtractAppConfigs(ctx)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if config.CommonConfigs.IsNull() || config.CommonConfigs.IsUnknown() {
		return
	}

	isClientTypeNull := false
	isAllowedScopesNull := false

	// required params
	if config.ClientType.IsNull() && (config.CommonConfigs.IsNull() || config.commonConfigs.ClientType.IsNull()) {
		isClientTypeNull = true
		resp.Diagnostics.AddError("Missing required argument",
			"The 'client_type' argument is required but has not been configured. You can set the 'client_type' either directly in the main configuration or within the 'common_configs' attribute.",
		)
	}
	if config.CompanyName.IsNull() && (config.CommonConfigs.IsNull() || config.commonConfigs.CompanyName.IsNull()) {
		resp.Diagnostics.AddError("Missing required argument",
			"The 'company_name' argument is required but has not been configured. You can set the 'company_name' either directly in the main configuration or within the 'common_configs' attribute.",
		)
	}
	if config.CompanyAddress.IsNull() && (config.CommonConfigs.IsNull() || config.commonConfigs.CompanyAddress.IsNull()) {
		resp.Diagnostics.AddError("Missing required argument",
			"The 'company_address' argument is required but has not been configured. You can set the 'company_address' either directly in the main configuration or within the 'common_configs' attribute.",
		)
	}
	if config.CompanyWebsite.IsNull() && (config.CommonConfigs.IsNull() || config.commonConfigs.CompanyWebsite.IsNull()) {
		resp.Diagnostics.AddError("Missing required argument",
			"The 'company_website' argument is required but has not been configured. You can set the 'company_website' either directly in the main configuration or within the 'common_configs' attribute.",
		)
	}

	if config.AllowedScopes.IsNull() && (config.CommonConfigs.IsNull() || config.commonConfigs.AllowedScopes.IsNull()) {
		isAllowedScopesNull = true
		resp.Diagnostics.AddError("Missing required argument",
			"The 'allowed_scopes' argument is required but has not been configured. You can set the 'allowed_scopes' either directly in the main configuration or within the 'common_configs' attribute.",
		)
	}

	if !isAllowedScopesNull {
		allowedScopes := []string{}
		if !config.AllowedScopes.IsNull() {
			for _, v := range config.AllowedScopes.Elements() {
				allowedScopes = append(allowedScopes, v.String())
			}
		} else if !config.commonConfigs.AllowedScopes.IsNull() {
			for _, v := range config.commonConfigs.AllowedScopes.Elements() {
				allowedScopes = append(allowedScopes, v.String())
			}
		}
		if (!config.AllowedScopes.IsNull() || !config.commonConfigs.AllowedScopes.IsNull()) && len(allowedScopes) < 1 {
			resp.Diagnostics.AddError("Unexpected Resource Configuration",
				"The 'allowed_scopes' argument must contain at least one value.",
			)
		}
	}

	if !isClientTypeNull {
		isRedirectURIReqired := false
		validClientTypes := []string{"SINGLE_PAGE", "REGULAR_WEB", "THIRD_PARTY"}
		if !config.ClientType.IsNull() {
			isRedirectURIReqired = util.StringInSlice(config.ClientType.ValueString(), validClientTypes)
		} else if !config.CommonConfigs.IsNull() && !config.commonConfigs.ClientType.IsNull() {
			isRedirectURIReqired = util.StringInSlice(config.ClientType.ValueString(), validClientTypes)
		}
		if isRedirectURIReqired {
			if config.RedirectURIS.IsNull() && (config.CommonConfigs.IsNull() || config.commonConfigs.RedirectUris.IsNull()) {
				resp.Diagnostics.AddError("Missing required argument",
					"The 'redirect_uris' argument is required but has not been configured. You can set the 'redirect_uris' either directly in the main configuration or within the 'common_configs' attribute.",
				)
			}

			if config.AllowedLogoutUrls.IsNull() && (config.CommonConfigs.IsNull() || config.commonConfigs.AllowedLogoutUrls.IsNull()) {
				resp.Diagnostics.AddError("Missing required argument",
					"The 'allowed_logout_urls' argument is required but has not been configured. You can set the 'allowed_logout_urls' either directly in the main configuration or within the 'common_configs' attribute.",
				)
			}

			redirectUrls := []string{}
			if !config.RedirectURIS.IsNull() {
				for _, v := range config.RedirectURIS.Elements() {
					redirectUrls = append(redirectUrls, v.String())
				}
			} else if !config.CommonConfigs.IsNull() && !config.commonConfigs.RedirectUris.IsNull() {
				for _, v := range config.commonConfigs.RedirectUris.Elements() {
					redirectUrls = append(redirectUrls, v.String())
				}
			}
			if (!config.RedirectURIS.IsNull() || (!config.CommonConfigs.IsNull() && !config.commonConfigs.RedirectUris.IsNull())) && len(redirectUrls) < 1 {
				resp.Diagnostics.AddError("Unexpected Resource Configuration",
					"The 'redirect_uris' argument must contain at least one URI.",
				)
			}

			logoutUrls := []string{}
			if !config.AllowedLogoutUrls.IsNull() {
				for _, v := range config.AllowedLogoutUrls.Elements() {
					logoutUrls = append(logoutUrls, v.String())
				}
			} else if !config.CommonConfigs.IsNull() && !config.commonConfigs.AllowedLogoutUrls.IsNull() {
				for _, v := range config.commonConfigs.AllowedLogoutUrls.Elements() {
					logoutUrls = append(logoutUrls, v.String())
				}
			}
			if (!config.AllowedLogoutUrls.IsNull() || (!config.CommonConfigs.IsNull() && !config.commonConfigs.AllowedLogoutUrls.IsNull())) && len(logoutUrls) < 1 {
				resp.Diagnostics.AddError("Unexpected Resource Configuration",
					"The 'allowed_logout_urls' argument must contain at least one URI.",
				)
			}
		}
	}
}

func (r *AppResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) { //nolint:gocognit
	if req.Plan.Raw.IsNull() {
		return
	}

	var config, plan AppConfig
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	resp.Diagnostics.Append(config.ExtractAppConfigs(ctx)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(plan.ExtractAppConfigs(ctx)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if plan.CommonConfigs.IsNull() || plan.CommonConfigs.IsUnknown() || plan.commonConfigs == nil {
		return
	}

	if !config.CommonConfigs.IsNull() {
		message := "\033[1mRecommendations:\033[0m\n\n" +
			"  - Utilize common_configs only when there are shared configuration attributes for multiple resources.\n" +
			"  - If you need to override any specific attribute for a particular resource, you can supply the main configuration attribute directly within the resource block.\n" +
			"  - If your configuration involves a single resource or if the common configuration attributes are not shared across multiple resources we do not suggest using common_configs."

		resp.Diagnostics.AddWarning(
			`Use of common_configs`,
			message,
		)
	}

	updatePlanStringAttribute := func(planAttr, commonConfigAttr, configAttr *types.String) {
		if (planAttr.IsNull() || planAttr.IsUnknown()) && !commonConfigAttr.IsNull() {
			*planAttr = *commonConfigAttr
		}
		if !planAttr.IsNull() && !planAttr.IsUnknown() && configAttr.IsNull() && !commonConfigAttr.IsNull() && (planAttr.ValueString() == "" || planAttr.ValueString() != commonConfigAttr.ValueString()) {
			*planAttr = *commonConfigAttr
		}
	}

	updateBoolAttribute := func(configAttr, planAttr, commonConfigAttr *types.Bool) {
		if configAttr.IsNull() && !commonConfigAttr.IsNull() && !commonConfigAttr.Equal(planAttr) {
			*planAttr = *commonConfigAttr
		}
	}
	updatePlanStringAttribute(&plan.CompanyName, &plan.commonConfigs.CompanyName, &config.CompanyName)
	updatePlanStringAttribute(&plan.CompanyWebsite, &plan.commonConfigs.CompanyWebsite, &config.CompanyWebsite)
	updatePlanStringAttribute(&plan.CompanyAddress, &plan.commonConfigs.CompanyAddress, &config.CompanyAddress)
	updatePlanStringAttribute(&plan.ClientType, &plan.commonConfigs.ClientType, &config.ClientType)
	updatePlanStringAttribute(&plan.AccentColor, &plan.commonConfigs.AccentColor, &config.AccentColor)
	updatePlanStringAttribute(&plan.PrimaryColor, &plan.commonConfigs.PrimaryColor, &config.PrimaryColor)
	updatePlanStringAttribute(&plan.MediaType, &plan.commonConfigs.MediaType, &config.MediaType)
	updatePlanStringAttribute(&plan.HostedPageGroup, &plan.commonConfigs.HostedPageGroup, &config.HostedPageGroup)
	updatePlanStringAttribute(&plan.TemplateGroupID, &plan.commonConfigs.TemplateGroupID, &config.TemplateGroupID)
	updatePlanStringAttribute(&plan.BotProvider, &plan.commonConfigs.BotProvider, &config.BotProvider)
	updatePlanStringAttribute(&plan.LogoAlign, &plan.commonConfigs.LogoAlign, &config.LogoAlign)
	updatePlanStringAttribute(&plan.Webfinger, &plan.commonConfigs.Webfinger, &config.Webfinger)

	// default int check
	if config.DefaultMaxAge.IsNull() && !plan.commonConfigs.DefaultMaxAge.IsNull() && !plan.commonConfigs.DefaultMaxAge.Equal(plan.DefaultMaxAge) {
		plan.DefaultMaxAge = plan.commonConfigs.DefaultMaxAge
	}
	if config.TokenLifetimeInSeconds.IsNull() && !plan.commonConfigs.TokenLifetimeInSeconds.IsNull() && !plan.commonConfigs.TokenLifetimeInSeconds.Equal(plan.TokenLifetimeInSeconds) {
		plan.TokenLifetimeInSeconds = plan.commonConfigs.TokenLifetimeInSeconds
	}
	if config.IDTokenLifetimeInSeconds.IsNull() && !plan.commonConfigs.IDTokenLifetimeInSeconds.IsNull() && !plan.commonConfigs.IDTokenLifetimeInSeconds.Equal(plan.IDTokenLifetimeInSeconds) {
		plan.IDTokenLifetimeInSeconds = plan.commonConfigs.IDTokenLifetimeInSeconds
	}
	if config.RefreshTokenLifetimeInSeconds.IsNull() && !plan.commonConfigs.RefreshTokenLifetimeInSeconds.IsNull() && !plan.commonConfigs.RefreshTokenLifetimeInSeconds.Equal(plan.RefreshTokenLifetimeInSeconds) {
		plan.RefreshTokenLifetimeInSeconds = plan.commonConfigs.RefreshTokenLifetimeInSeconds
	}

	// default bool check
	updateBoolAttribute(&config.AllowGuestLogin, &plan.AllowGuestLogin, &plan.commonConfigs.AllowGuestLogin)
	updateBoolAttribute(&config.EnableDeduplication, &plan.EnableDeduplication, &plan.commonConfigs.EnableDeduplication)
	updateBoolAttribute(&config.AutoLoginAfterRegister, &plan.AutoLoginAfterRegister, &plan.commonConfigs.AutoLoginAfterRegister)
	updateBoolAttribute(&config.EnablePasswordlessAuth, &plan.EnablePasswordlessAuth, &plan.commonConfigs.EnablePasswordlessAuth)
	updateBoolAttribute(&config.RegisterWithLoginInformation, &plan.RegisterWithLoginInformation, &plan.commonConfigs.RegisterWithLoginInformation)
	updateBoolAttribute(&config.FdsEnabled, &plan.FdsEnabled, &plan.commonConfigs.FdsEnabled)
	updateBoolAttribute(&config.IsHybridApp, &plan.IsHybridApp, &plan.commonConfigs.IsHybridApp)
	updateBoolAttribute(&config.Editable, &plan.Editable, &plan.commonConfigs.Editable)
	updateBoolAttribute(&config.Enabled, &plan.Enabled, &plan.commonConfigs.Enabled)
	updateBoolAttribute(&config.AlwaysAskMfa, &plan.AlwaysAskMfa, &plan.commonConfigs.AlwaysAskMfa)
	updateBoolAttribute(&config.EmailVerificationRequired, &plan.EmailVerificationRequired, &plan.commonConfigs.EmailVerificationRequired)
	updateBoolAttribute(&config.EnableClassicalProvider, &plan.EnableClassicalProvider, &plan.commonConfigs.EnableClassicalProvider)
	updateBoolAttribute(&config.IsRememberMeSelected, &plan.IsRememberMeSelected, &plan.commonConfigs.IsRememberMeSelected)

	// set attributes
	updateSetAttributes := func(configValue, commonConfigValue, planValue basetypes.SetValue, planField *basetypes.SetValue) {
		if configValue.IsNull() {
			if !commonConfigValue.IsNull() {
				*planField = commonConfigValue
			} else if !planValue.IsNull() && !planValue.IsUnknown() {
				*planField = types.SetValueMust(types.StringType, []attr.Value{})
			}
		}
	}

	updateSetAttributes(config.AllowedScopes, config.commonConfigs.AllowedScopes, plan.AllowedScopes, &plan.AllowedScopes)
	updateSetAttributes(config.RedirectURIS, config.commonConfigs.RedirectUris, plan.RedirectURIS, &plan.RedirectURIS)
	updateSetAttributes(config.AllowedLogoutUrls, config.commonConfigs.AllowedLogoutUrls, plan.AllowedLogoutUrls, &plan.AllowedLogoutUrls)
	updateSetAttributes(config.AllowedWebOrigins, config.commonConfigs.AllowedWebOrigins, plan.AllowedWebOrigins, &plan.AllowedWebOrigins)
	updateSetAttributes(config.AllowedOrigins, config.commonConfigs.AllowedOrigins, plan.AllowedOrigins, &plan.AllowedOrigins)
	updateSetAttributes(config.LoginProviders, config.commonConfigs.LoginProviders, plan.LoginProviders, &plan.LoginProviders)
	updateSetAttributes(config.DefaultScopes, config.commonConfigs.DefaultScopes, plan.DefaultScopes, &plan.DefaultScopes)
	updateSetAttributes(config.PendingScopes, config.commonConfigs.PendingScopes, plan.PendingScopes, &plan.PendingScopes)
	updateSetAttributes(config.AllowedMfa, config.commonConfigs.AllowedMfa, plan.AllowedMfa, &plan.AllowedMfa)
	updateSetAttributes(config.AllowedRoles, config.commonConfigs.AllowedRoles, plan.AllowedRoles, &plan.AllowedRoles)
	updateSetAttributes(config.DefaultRoles, config.commonConfigs.DefaultRoles, plan.DefaultRoles, &plan.DefaultRoles)

	// has default set value
	if config.ResponseTypes.IsNull() {
		if !config.commonConfigs.ResponseTypes.IsNull() {
			plan.ResponseTypes = config.commonConfigs.ResponseTypes
		} else if !plan.ResponseTypes.IsNull() && !plan.ResponseTypes.IsUnknown() {
			plan.ResponseTypes = basetypes.NewSetValueMust(types.StringType, []attr.Value{
				types.StringValue("code"), types.StringValue("token"), types.StringValue("id_token"),
			})
		}
	}
	if config.GrantTypes.IsNull() {
		if !config.commonConfigs.GrantTypes.IsNull() {
			plan.GrantTypes = config.commonConfigs.GrantTypes
		} else if !plan.GrantTypes.IsNull() && !plan.GrantTypes.IsUnknown() {
			plan.GrantTypes = basetypes.NewSetValueMust(types.StringType, []attr.Value{
				types.StringValue("implicit"), types.StringValue("authorization_code"), types.StringValue("password"), types.StringValue("refresh_token"),
			})
		}
	}

	if config.AllowLoginWith.IsNull() {
		if !config.commonConfigs.AllowLoginWith.IsNull() {
			plan.AllowLoginWith = config.commonConfigs.AllowLoginWith
		} else if !plan.AllowLoginWith.IsNull() && !plan.AllowLoginWith.IsUnknown() {
			plan.AllowLoginWith = basetypes.NewSetValueMust(types.StringType, []attr.Value{
				types.StringValue("EMAIL"), types.StringValue("MOBILE"), types.StringValue("USER_NAME"),
			})
		}
	}

	if config.Mfa.IsNull() {
		if !config.commonConfigs.Mfa.IsNull() {
			plan.Mfa = plan.commonConfigs.Mfa
		} else if !plan.Mfa.IsNull() && !plan.Mfa.IsUnknown() {
			mfa := types.ObjectValueMust(
				map[string]attr.Type{
					"setting":                  types.StringType,
					"time_interval_in_seconds": types.Int64Type,
					"allowed_methods":          types.SetType{ElemType: types.StringType},
				},
				map[string]attr.Value{
					"setting":                  types.StringValue("OFF"),
					"time_interval_in_seconds": types.Int64Null(),
					"allowed_methods":          types.SetNull(types.StringType),
				},
			)
			plan.Mfa = mfa
		}
	}

	if config.SocialProviders.IsNull() {
		if !config.commonConfigs.SocialProviders.IsNull() {
			plan.SocialProviders = plan.commonConfigs.SocialProviders
		} else if !plan.SocialProviders.IsNull() && !plan.SocialProviders.IsUnknown() {
			var spObjectValues []attr.Value
			spObjectType := types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"provider_name": types.StringType,
					"social_id":     types.StringType,
				},
			}
			objValue := types.ObjectValueMust(
				spObjectType.AttrTypes,
				map[string]attr.Value{
					"provider_name": types.StringNull(),
					"social_id":     types.StringNull(),
				})
			spObjectValues = append(spObjectValues, objValue)
			plan.SocialProviders = types.ListValueMust(spObjectType, spObjectValues)
		}
	}

	if config.CustomProviders.IsNull() {
		if !config.commonConfigs.CustomProviders.IsNull() {
			plan.CustomProviders = plan.commonConfigs.CustomProviders
		} else if !plan.CustomProviders.IsNull() && !plan.CustomProviders.IsUnknown() {
			var objectValues []attr.Value
			spObjectType := types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"provider_name":       types.StringType,
					"display_name":        types.StringType,
					"logo_url":            types.StringType,
					"type":                types.StringType,
					"is_provider_visible": types.BoolType,
					"domains":             types.SetType{ElemType: types.StringType},
				},
			}
			objValue := types.ObjectValueMust(
				spObjectType.AttrTypes,
				map[string]attr.Value{
					"provider_name":       types.StringNull(),
					"display_name":        types.StringNull(),
					"logo_url":            types.StringNull(),
					"type":                types.StringNull(),
					"is_provider_visible": types.BoolValue(false),
					"domains":             types.SetNull(types.StringType),
				})
			objectValues = append(objectValues, objValue)
			plan.CustomProviders = types.ListValueMust(spObjectType, objectValues)
		}
	}

	if config.SamlProviders.IsNull() {
		if !config.commonConfigs.SamlProviders.IsNull() {
			plan.SamlProviders = plan.commonConfigs.SamlProviders
		} else if !plan.SamlProviders.IsNull() && !plan.SamlProviders.IsUnknown() {
			var objectValues []attr.Value
			spObjectType := types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"provider_name":       types.StringType,
					"display_name":        types.StringType,
					"logo_url":            types.StringType,
					"type":                types.StringType,
					"is_provider_visible": types.BoolType,
					"domains":             types.SetType{ElemType: types.StringType},
				},
			}
			objValue := types.ObjectValueMust(
				spObjectType.AttrTypes,
				map[string]attr.Value{
					"provider_name":       types.StringNull(),
					"display_name":        types.StringNull(),
					"logo_url":            types.StringNull(),
					"type":                types.StringNull(),
					"is_provider_visible": types.BoolValue(false),
					"domains":             types.SetNull(types.StringType),
				})
			objectValues = append(objectValues, objValue)
			plan.SamlProviders = types.ListValueMust(spObjectType, objectValues)
		}
	}

	if config.AdProviders.IsNull() {
		if !config.commonConfigs.AdProviders.IsNull() {
			plan.AdProviders = plan.commonConfigs.AdProviders
		} else if !plan.AdProviders.IsNull() && !plan.AdProviders.IsUnknown() {
			var objectValues []attr.Value
			spObjectType := types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"provider_name":       types.StringType,
					"display_name":        types.StringType,
					"logo_url":            types.StringType,
					"type":                types.StringType,
					"is_provider_visible": types.BoolType,
					"domains":             types.SetType{ElemType: types.StringType},
				},
			}
			objValue := types.ObjectValueMust(
				spObjectType.AttrTypes,
				map[string]attr.Value{
					"provider_name":       types.StringNull(),
					"display_name":        types.StringNull(),
					"logo_url":            types.StringNull(),
					"type":                types.StringNull(),
					"is_provider_visible": types.BoolValue(false),
					"domains":             types.SetNull(types.StringType),
				})
			objectValues = append(objectValues, objValue)
			plan.AdProviders = types.ListValueMust(spObjectType, objectValues)
		}
	}

	if config.AllowedGroups.IsNull() {
		if !config.commonConfigs.AllowedGroups.IsNull() {
			plan.AllowedGroups = plan.commonConfigs.AllowedGroups
		} else if !plan.AllowedGroups.IsNull() && !plan.AllowedGroups.IsUnknown() {
			var objectValues []attr.Value
			spObjectType := types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"group_id":      types.StringType,
					"roles":         types.SetType{ElemType: types.StringType},
					"default_roles": types.SetType{ElemType: types.StringType},
				},
			}
			objValue := types.ObjectValueMust(
				spObjectType.AttrTypes,
				map[string]attr.Value{
					"group_id":      types.StringNull(),
					"roles":         types.SetNull(types.StringType),
					"default_roles": types.SetNull(types.StringType),
				})
			objectValues = append(objectValues, objValue)
			plan.AllowedGroups = types.ListValueMust(spObjectType, objectValues)
		}
	}

	if config.OperationsAllowedGroups.IsNull() {
		if !config.commonConfigs.OperationsAllowedGroups.IsNull() {
			plan.OperationsAllowedGroups = plan.commonConfigs.OperationsAllowedGroups
		} else if !plan.OperationsAllowedGroups.IsNull() && !plan.OperationsAllowedGroups.IsUnknown() {
			var objectValues []attr.Value
			spObjectType := types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"group_id":      types.StringType,
					"roles":         types.SetType{ElemType: types.StringType},
					"default_roles": types.SetType{ElemType: types.StringType},
				},
			}
			objValue := types.ObjectValueMust(
				spObjectType.AttrTypes,
				map[string]attr.Value{
					"group_id":      types.StringNull(),
					"roles":         types.SetNull(types.StringType),
					"default_roles": types.SetNull(types.StringType),
				})
			objectValues = append(objectValues, objValue)
			plan.OperationsAllowedGroups = types.ListValueMust(spObjectType, objectValues)
		}
	}
	resp.Plan.Set(ctx, plan)
}
