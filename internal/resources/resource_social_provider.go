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
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"strings"
)

type SocialProvider struct {
	cidaasClient *cidaas.Client
}

type SocialProviderConfig struct {
	ID                    types.String `tfsdk:"id"`
	Name                  types.String `tfsdk:"name"`
	ProviderName          types.String `tfsdk:"provider_name"`
	Enabled               types.Bool   `tfsdk:"enabled"`
	ClientID              types.String `tfsdk:"client_id"`
	ClientSecret          types.String `tfsdk:"client_secret"`
	ClaimsConfig          types.Object `tfsdk:"claims"`
	EnabledForAdminPortal types.Bool   `tfsdk:"enabled_for_admin_portal"`
	UserInfoFieldsConfigs types.List   `tfsdk:"userinfo_fields"`
	userInfoFieldsConfigs []*UserInfoFieldsConfig
	Scopes                types.Set `tfsdk:"scopes"`
	//CreatedAt             types.String `tfsdk:"created_at"`
	//UpdatedAt             types.String `tfsdk:"updated_at"`
	claimsConfig *ClaimsConfig
}

type ClaimsConfig struct {
	RequiredClaimsConfig types.Object `tfsdk:"required_claims"`
	OptionalClaimsConfig types.Object `tfsdk:"optional_claims"`
	requiredClaimsConfig *RequiredClaimsConfig
	optionalClaimsConfig *OptionalClaimsConfig
}

type RequiredClaimsConfig struct {
	UserInfo types.Set `tfsdk:"user_info"`
	IdToken  types.Set `tfsdk:"id_token"`
}

type OptionalClaimsConfig struct {
	UserInfo types.Set `tfsdk:"user_info"`
	IdToken  types.Set `tfsdk:"id_token"`
}

type UserInfoFieldsConfig struct {
	InnerKey      types.String `tfsdk:"inner_key"`
	ExternalKey   types.String `tfsdk:"external_key"`
	IsCustomField types.Bool   `tfsdk:"is_custom_field"`
	IsSystemField types.Bool   `tfsdk:"is_system_field"`
}

func NewSocialProvider() resource.Resource {
	return &SocialProvider{}
}

func (r *SocialProvider) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_social_provider"
}

func (r *SocialProvider) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *SocialProvider) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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
				PlanModifiers: []planmodifier.String{
					&validators.UniqueIdentifier{},
				},
			},
			"provider_name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					&validators.UniqueIdentifier{},
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						func() []string {
							var provider = make([]string, len(util.AllowedProviders)) //nolint:gofumpt
							for i, cType := range util.AllowedProviders {
								provider[i] = cType
							}
							return provider
						}()...),
				},
			},
			"enabled": schema.BoolAttribute{
				Required: true,
			},
			"client_id": schema.StringAttribute{
				Required: true,
			},
			"client_secret": schema.StringAttribute{
				Required:  true,
				Sensitive: true,
			},
			"scopes": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"claims": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"required_claims": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"user_info": schema.SetAttribute{
								ElementType: types.StringType,
								Optional:    true,
							},
							"id_token": schema.SetAttribute{
								ElementType: types.StringType,
								Optional:    true,
							},
						},
					},
					"optional_claims": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"user_info": schema.SetAttribute{
								ElementType: types.StringType,
								Optional:    true,
							},
							"id_token": schema.SetAttribute{
								ElementType: types.StringType,
								Optional:    true,
							},
						},
					},
				},
			},
			"userinfo_fields": schema.ListNestedAttribute{
				Computed: true,
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"inner_key": schema.StringAttribute{
							Required: true,
						},
						"external_key": schema.StringAttribute{
							Required: true,
						},
						"is_custom_field": schema.BoolAttribute{
							Required: true,
						},
						"is_system_field": schema.BoolAttribute{
							Required: true,
						},
					},
				},
			},
			"enabled_for_admin_portal": schema.BoolAttribute{
				Required: true,
			},
		},
	}
}

func (r *SocialProvider) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SocialProviderConfig
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(plan.extract(ctx)...)
	resp.Diagnostics.AddWarning("ClaimsConfig", fmt.Sprintf("SP: %+v ", plan.ClaimsConfig))
	model, diag := prepareSocialProviderModel(ctx, plan)
	if diag.HasError() {
		resp.Diagnostics.AddWarning("error preparing social provider payload ", fmt.Sprintf("Error: %+v ", diag.Errors()))
		return
	}
	//resp.Diagnostics.AddWarning("SocialProviderModel ", fmt.Sprintf("SP: %+v ", model))
	res, err := r.cidaasClient.SocialProvider.UpsertSocialProvider(model)
	if err != nil {
		resp.Diagnostics.AddError("failed to create social provider", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	resp.Diagnostics.AddWarning("SocialProviderModel Response ", fmt.Sprintf("SP: %+v ", res.Data))

	//plan.ID = types.StringValue(res.Data.ID)
	plan.ID = types.StringValue(res.Data.ProviderName + "_" + res.Data.ID)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SocialProvider) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SocialProviderConfig
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	params := strings.Split(state.ID.ValueString(), "_")
	if len(params) != 2 {
		resp.Diagnostics.AddError("invalid id. Valid Format: providerName_ID ", fmt.Sprintf("ID: %s", state.ID.ValueString()))
		return
	}
	//res, err := r.cidaasClient.SocialProvider.GetSocialProvider(state.ProviderName.ValueString(), state.ID.ValueString())
	res, err := r.cidaasClient.SocialProvider.GetSocialProvider(params[0], params[1])
	if err != nil {
		resp.Diagnostics.AddError("failed to read social provider", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	state.ClientID = types.StringValue(res.Data.ClientID)
	state.ClientSecret = types.StringValue(res.Data.ClientSecret)
	state.Enabled = types.BoolValue(res.Data.Enabled)
	state.EnabledForAdminPortal = types.BoolValue(res.Data.EnabledForAdminPortal)
	//
	//state.ID = types.StringValue(res.Data.ID)
	//state.ProviderName = types.StringValue(res.Data.ID)

	var d diag.Diagnostics
	if len(res.Data.Scopes) > 0 {
		state.Scopes, d = types.SetValueFrom(ctx, types.StringType, res.Data.Scopes)
		resp.Diagnostics.Append(d...)
		if resp.Diagnostics.HasError() {
			return
		}
	}
	ReadClaimsInfo(state, res.Data, ctx, resp)
	//resp.Diagnostics.AddWarning("State After reading.", fmt.Sprintf("state: %+v", state))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *SocialProvider) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state SocialProviderConfig
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(plan.extract(ctx)...)
	params := strings.Split(plan.ID.ValueString(), "_")
	if len(params) != 2 {
		resp.Diagnostics.AddError("invalid id. Valid Format: providerName_ID ", fmt.Sprintf("ID: %s", state.ID.ValueString()))
		return
	}
	plan.ID = types.StringValue(params[1])
	state.ID = types.StringValue(params[1])
	plan.ProviderName = types.StringValue(params[0])
	spModel, d := prepareSocialProviderModel(ctx, plan)
	//spModel.ID = params[1]
	resp.Diagnostics.Append(d...)
	if resp.Diagnostics.HasError() {
		return
	}
	spModel.ID = state.ID.ValueString()

	res, err := r.cidaasClient.SocialProvider.UpsertSocialProvider(spModel)
	if err != nil {
		resp.Diagnostics.AddError("failed to update social provider", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	plan.ID = types.StringValue(res.Data.ProviderName + "_" + res.Data.ID)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SocialProvider) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SocialProviderConfig
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	params := strings.Split(state.ID.ValueString(), "_")
	if len(params) != 2 {
		resp.Diagnostics.AddError("invalid id. Valid Format: providerName_ID ", fmt.Sprintf("ID: %s", state.ID.ValueString()))
	}
	//err := r.cidaasClient.SocialProvider.DeleteSocialProvider(state.ProviderName.ValueString(), state.ID.ValueString())
	err := r.cidaasClient.SocialProvider.DeleteSocialProvider(params[0], params[1])
	if err != nil {
		resp.Diagnostics.AddError("failed to delete social provider", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
}

func (r *SocialProvider) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	//resource.
}

func (sp *SocialProviderConfig) extract(ctx context.Context) diag.Diagnostics {
	var diags diag.Diagnostics
	if !sp.ClaimsConfig.IsNull() {
		sp.claimsConfig = &ClaimsConfig{}
		diags = sp.ClaimsConfig.As(ctx, sp.claimsConfig, basetypes.ObjectAsOptions{})
		sp.claimsConfig.optionalClaimsConfig = &OptionalClaimsConfig{}
		diags = sp.claimsConfig.OptionalClaimsConfig.As(ctx, sp.claimsConfig.optionalClaimsConfig, basetypes.ObjectAsOptions{})
		sp.claimsConfig.requiredClaimsConfig = &RequiredClaimsConfig{}
		diags = sp.claimsConfig.RequiredClaimsConfig.As(ctx, sp.claimsConfig.requiredClaimsConfig, basetypes.ObjectAsOptions{})
		//sp.claimsConfig.requiredClaimsConfig.UserInfo = make([]*string, 0, len(sp.claimsConfig.requiredClaimsConfig.UserInfo.Elements()))
		/*if !sp.claimsConfig.RequiredClaimsConfig.IsNull() {
			diags = sp.claimsConfig.requiredClaimsConfig.UserInfo.ElementsAs(ctx, &sp.claimsConfig.requiredClaimsConfig.UserInfo, false)
			diags = sp.claimsConfig.requiredClaimsConfig.IdToken.ElementsAs(ctx, &sp.claimsConfig.requiredClaimsConfig.IdToken, false)
		}
		if !sp.claimsConfig.OptionalClaimsConfig.IsNull() {
			diags = sp.claimsConfig.optionalClaimsConfig.UserInfo.ElementsAs(ctx, &sp.claimsConfig.optionalClaimsConfig.UserInfo, false)
			diags = sp.claimsConfig.optionalClaimsConfig.IdToken.ElementsAs(ctx, &sp.claimsConfig.optionalClaimsConfig.IdToken, false)
		}*/
		if !sp.UserInfoFieldsConfigs.IsNull() {
			sp.userInfoFieldsConfigs = make([]*UserInfoFieldsConfig, 0, len(sp.UserInfoFieldsConfigs.Elements()))
			diags = sp.UserInfoFieldsConfigs.ElementsAs(ctx, &sp.userInfoFieldsConfigs, false)
		}

	}
	/*if !sp.ClaimsConfig.IsNull() &&  sp.claimsConfig.{
		pc.scopes = make([]*CpScope, 0, len(pc.Scopes.Elements()))
		diags = pc.Scopes.ElementsAs(ctx, &pc.scopes, false)
	}*/

	return diags
}

func prepareSocialProviderModel(ctx context.Context, plan SocialProviderConfig) (*cidaas.SocialProviderModel, diag.Diagnostics) {
	var sp cidaas.SocialProviderModel
	sp.Name = plan.Name.ValueString()
	sp.ProviderName = plan.ProviderName.ValueString()
	sp.ClientID = plan.ClientID.ValueString()
	sp.ClientSecret = plan.ClientSecret.ValueString()
	sp.EnabledForAdminPortal = plan.EnabledForAdminPortal.ValueBool()
	sp.Enabled = plan.Enabled.ValueBool()
	if !plan.ClaimsConfig.IsNull() {
		diag := plan.claimsConfig.requiredClaimsConfig.UserInfo.ElementsAs(ctx, &sp.Claims.RequiredClaims.UserInfo, false)
		diag = plan.claimsConfig.requiredClaimsConfig.IdToken.ElementsAs(ctx, &sp.Claims.RequiredClaims.IdToken, false)
		diag = plan.claimsConfig.optionalClaimsConfig.IdToken.ElementsAs(ctx, &sp.Claims.OptionalClaims.IdToken, false)
		diag = plan.claimsConfig.optionalClaimsConfig.UserInfo.ElementsAs(ctx, &sp.Claims.OptionalClaims.UserInfo, false)
		if diag.HasError() {
			return nil, diag
		}
	}
	if !plan.UserInfoFieldsConfigs.IsNull() {
		var userInfoModels []cidaas.UserInfoFieldsModel
		for _, v := range plan.userInfoFieldsConfigs {
			userInfoModels = append(userInfoModels, cidaas.UserInfoFieldsModel{
				IsCustomField: v.IsCustomField.ValueBool(),
				IsSystemField: v.IsSystemField.ValueBool(),
				InnerKey:      v.InnerKey.ValueString(),
				ExternalKey:   v.ExternalKey.ValueString(),
			})
		}
		sp.UserInfoFields = userInfoModels
	}
	diag := plan.Scopes.ElementsAs(ctx, &sp.Scopes, false)
	if diag.HasError() {
		return nil, diag
	}
	return &sp, nil
}

func ReadClaimsInfo(spc SocialProviderConfig, model cidaas.SocialProviderModel, ctx context.Context, resp *resource.ReadResponse) {
	var setIdTokenHolder basetypes.SetValue
	var setUserInfoHolder basetypes.SetValue
	var diag diag.Diagnostics
	if len(model.Claims.RequiredClaims.IdToken) > 0 {
		setIdTokenHolder, diag = types.SetValueFrom(ctx, types.StringType, model.Claims.RequiredClaims.IdToken)
		resp.Diagnostics.Append(diag...)
	}
	if len(model.Claims.RequiredClaims.UserInfo) > 0 {
		setUserInfoHolder, diag = types.SetValueFrom(ctx, types.StringType, model.Claims.RequiredClaims.UserInfo)
		resp.Diagnostics.Append(diag...)
	}
	/*reqClaimsObjectType := types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"user_info": types.SetType{ElemType: types.StringType},
			"id_token":  types.SetType{ElemType: types.StringType},
		},
	}
	reqClaimsObjectValue := types.ObjectValueMust(
		reqClaimsObjectType.AttrTypes, map[string]attr.Value{
			"user_info": setUserInfoHolder,
			"id_token":  setIdTokenHolder,
		},
	)*/
	if len(model.Claims.OptionalClaims.IdToken) > 0 {
		setIdTokenHolder, diag = types.SetValueFrom(ctx, types.StringType, model.Claims.OptionalClaims.IdToken)
		resp.Diagnostics.Append(diag...)
	}
	if len(model.Claims.OptionalClaims.UserInfo) > 0 {
		setUserInfoHolder, diag = types.SetValueFrom(ctx, types.StringType, model.Claims.OptionalClaims.UserInfo)
		resp.Diagnostics.Append(diag...)
	}
	fmt.Println("sss", setUserInfoHolder, setIdTokenHolder)
	/*optionalClaimsObjectValue := types.ObjectValueMust(
		reqClaimsObjectType.AttrTypes, map[string]attr.Value{
			"user_info": setUserInfoHolder,
			"id_token":  setIdTokenHolder,
		},
	)*/
}
