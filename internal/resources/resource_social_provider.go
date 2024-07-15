package resources

import (
	"context"
	"fmt"
	"strings"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/cidaas"
	"github.com/Cidaas/terraform-provider-cidaas/helpers/util"
	"github.com/Cidaas/terraform-provider-cidaas/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// var allowedClaims = []string{
// 	"name", "given_name", "family_name", "nickname", "preferred_username", "profile",
// 	"picture", "website", "email", "email_verified", "gender", "middle_name", "birthdate", "zoneinfo", "locale",
// 	"phone_number", "phone_number_verified", "formatted", "street_address", "locality", "region", "postal_code", "country",
// }

var allowedProviders = []string{
	"google", "facebook", "linkedin", "amazon", "foursquare", "github",
	"instagram", "yammer", "wordpress", "microsoft", "yahoo", "officeenterprise", "salesforce", "paypal_sandbox", "paypal",
	"apple", "twitter", "netid", "netid_qa",
}

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
	Claims                types.Object `tfsdk:"claims"`
	EnabledForAdminPortal types.Bool   `tfsdk:"enabled_for_admin_portal"`
	UserInfoFields        types.List   `tfsdk:"userinfo_fields"`
	Scopes                types.Set    `tfsdk:"scopes"`

	claims         *Claims
	userInfoFields []*UserInfoFields
}

type Claims struct {
	RequiredClaims types.Object `tfsdk:"required_claims"`
	OptionalClaims types.Object `tfsdk:"optional_claims"`
	requiredClaims *ClaimConfigs
	optionalClaims *ClaimConfigs
}

type ClaimConfigs struct {
	UserInfo types.Set `tfsdk:"user_info"`
	IDToken  types.Set `tfsdk:"id_token"`
}

type UserInfoFields struct {
	InnerKey      types.String `tfsdk:"inner_key"`
	ExternalKey   types.String `tfsdk:"external_key"`
	IsCustomField types.Bool   `tfsdk:"is_custom_field"`
	IsSystemField types.Bool   `tfsdk:"is_system_field"`
}

func NewSocialProvider() resource.Resource {
	return &SocialProvider{}
}

func (r *SocialProvider) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, res *resource.ValidateConfigResponse) {
	var config SocialProviderConfig
	res.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	res.Diagnostics.Append(config.extract(ctx)...)
	if !config.UserInfoFields.IsNull() {
		innerKeysCustom := make(map[string]bool)
		innerKeysSystem := make(map[string]bool)
		for _, v := range config.userInfoFields {
			if v.IsCustomField.Equal(v.IsSystemField) {
				res.Diagnostics.AddError(
					"Unexpected Resource Configure Type",
					fmt.Sprintf("For inner_key \033[1m%s\033[0m the fields is_custom_field and is_system_field cannot have the same value.", v.InnerKey.ValueString()),
				)
			}
			if v.IsCustomField.ValueBool() {
				if innerKeysCustom[v.InnerKey.ValueString()] {
					res.Diagnostics.AddError(
						"Duplicate Custom Field Key",
						fmt.Sprintf("The custom field with the inner_key \033[1m%s\033[0m is repeated. Each key must be unique.", v.InnerKey.ValueString()),
					)
				}
				innerKeysCustom[v.InnerKey.ValueString()] = true
			}
			if v.IsSystemField.ValueBool() {
				if innerKeysSystem[v.InnerKey.ValueString()] {
					res.Diagnostics.AddError(
						"Duplicate System Field Key",
						fmt.Sprintf("The system field with the inner_key \033[1m%s\033[0m  is repeated. Each key must be unique.", v.InnerKey.ValueString()),
					)
				}
				innerKeysSystem[v.InnerKey.ValueString()] = true
			}
		}
	}
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
		MarkdownDescription: "The `cidaas_social_provider` resource allows you to configure and manage social login providers within Cidaas." +
			"\n Social login providers enable users to authenticate using their existing accounts from popular social platforms such as Google, Facebook, LinkedIn and others." +
			"\n\n Ensure that the below scopes are assigned to the client:" +
			"\n- cidaas:providers_read" +
			"\n- cidaas:providers_write",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "The unique identifier of the social provider",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					&validators.UniqueIdentifier{},
				},
				MarkdownDescription: "The name of the social provider configuration. This should be unique within your Cidaas environment.",
			},
			"provider_name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					&validators.UniqueIdentifier{},
				},
				Validators: []validator.String{
					stringvalidator.OneOf(allowedProviders...),
				},
				MarkdownDescription: "The name of the social provider. Supported values include `google`, `facebook`, `linkedin` etc.",
			},
			"enabled": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
				MarkdownDescription: "A flag to enable or disable the social provider configuration. Set to `true` to enable and `false` to disable.",
			},
			"client_id": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The client ID provided by the social provider. This is used to authenticate your application with the social provider.",
			},
			"client_secret": schema.StringAttribute{
				Required:            true,
				Sensitive:           true,
				MarkdownDescription: "The client secret provided by the social provider. This is used alongside the client ID to authenticate your application with the social provider.",
			},
			"scopes": schema.SetAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				MarkdownDescription: "A list of scopes of the social provider.",
			},
			"claims": schema.SingleNestedAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "A map defining required and optional claims to be requested from the social provider.",
				Attributes: map[string]schema.Attribute{
					"required_claims": schema.SingleNestedAttribute{
						Optional:            true,
						MarkdownDescription: "Defines the claims that are required from the social provider.",
						Attributes: map[string]schema.Attribute{
							"user_info": schema.SetAttribute{
								ElementType:         types.StringType,
								Optional:            true,
								MarkdownDescription: "A list of user information claims that are required.",
							},
							"id_token": schema.SetAttribute{
								ElementType:         types.StringType,
								Optional:            true,
								MarkdownDescription: "A list of ID token claims that are required.",
							},
						},
					},
					"optional_claims": schema.SingleNestedAttribute{
						Optional:            true,
						MarkdownDescription: "Defines the claims that are optional from the social provider.",
						Attributes: map[string]schema.Attribute{
							"user_info": schema.SetAttribute{
								ElementType:         types.StringType,
								Optional:            true,
								MarkdownDescription: "A list of user information claims that are optional.",
							},
							"id_token": schema.SetAttribute{
								ElementType:         types.StringType,
								Optional:            true,
								MarkdownDescription: "A list of ID token claims that are optional.",
							},
						},
					},
				},
				Default: objectdefault.StaticValue(
					types.ObjectValueMust(
						map[string]attr.Type{
							"required_claims": types.ObjectType{
								AttrTypes: map[string]attr.Type{
									"user_info": types.SetType{ElemType: types.StringType},
									"id_token":  types.SetType{ElemType: types.StringType},
								},
							},
							"optional_claims": types.ObjectType{
								AttrTypes: map[string]attr.Type{
									"user_info": types.SetType{ElemType: types.StringType},
									"id_token":  types.SetType{ElemType: types.StringType},
								},
							},
						},
						map[string]attr.Value{
							"required_claims": types.ObjectValueMust(map[string]attr.Type{
								"user_info": types.SetType{ElemType: types.StringType},
								"id_token":  types.SetType{ElemType: types.StringType},
							},
								map[string]attr.Value{
									"user_info": types.SetNull(types.StringType),
									"id_token":  types.SetNull(types.StringType),
								}),
							"optional_claims": types.ObjectValueMust(map[string]attr.Type{
								"user_info": types.SetType{ElemType: types.StringType},
								"id_token":  types.SetType{ElemType: types.StringType},
							},
								map[string]attr.Value{
									"user_info": types.SetNull(types.StringType),
									"id_token":  types.SetNull(types.StringType),
								}),
						}),
				),
			},
			"userinfo_fields": schema.ListNestedAttribute{
				Optional:            true,
				MarkdownDescription: "A list of user info fields to be mapped between the social provider and Cidaas.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"inner_key": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "The internal key used by Cidaas.",
						},
						"external_key": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "The external key used by the social provider.",
						},
						"is_custom_field": schema.BoolAttribute{
							Required:            true,
							MarkdownDescription: "A flag indicating whether the field is a custom field. Set to `true` if it is a custom field.",
						},
						"is_system_field": schema.BoolAttribute{
							Required:            true,
							MarkdownDescription: "A flag indicating whether the field is a system field. Set to `true` if it is a system field.",
						},
					},
				},
			},
			"enabled_for_admin_portal": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
				MarkdownDescription: "A flag to enable or disable the social provider for the admin portal. Set to `true` to enable and `false` to disable.",
			},
		},
	}
}

func (r *SocialProvider) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SocialProviderConfig
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(plan.extract(ctx)...)
	model, diag := prepareSocialProviderModel(ctx, plan)
	if diag.HasError() {
		resp.Diagnostics.AddError("error preparing social provider payload ", fmt.Sprintf("Error: %+v ", diag.Errors()))
		return
	}
	res, err := r.cidaasClient.SocialProvider.Upsert(model)
	if err != nil {
		resp.Diagnostics.AddError("failed to create social provider", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	plan.ID = util.StringValueOrNull(&res.Data.ID)
	resp.Diagnostics.Append(setClaimsInfo(&plan, res.Data.Claims)...)
	resp.Diagnostics.Append(setUserInfoFields(&plan, res.Data.UserInfoFields)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SocialProvider) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SocialProviderConfig
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.cidaasClient.SocialProvider.Get(state.ProviderName.ValueString(), state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to read social provider", fmt.Sprintf("Error: %s", err.Error()))
		return
	}

	state.ID = util.StringValueOrNull(&res.Data.ID)
	state.Name = util.StringValueOrNull(&res.Data.Name)
	state.ProviderName = util.StringValueOrNull(&res.Data.ProviderName)
	state.ClientID = util.StringValueOrNull(&res.Data.ClientID)
	state.ClientSecret = util.StringValueOrNull(&res.Data.ClientSecret)
	state.Enabled = util.BoolValueOrNull(&res.Data.Enabled)
	state.EnabledForAdminPortal = util.BoolValueOrNull(&res.Data.EnabledForAdminPortal)
	state.Scopes = util.SetValueOrNull(res.Data.Scopes)

	resp.Diagnostics.Append(setClaimsInfo(&state, res.Data.Claims)...)
	resp.Diagnostics.Append(setUserInfoFields(&state, res.Data.UserInfoFields)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *SocialProvider) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state SocialProviderConfig
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(plan.extract(ctx)...)
	model, diag := prepareSocialProviderModel(ctx, plan)
	if diag.HasError() {
		resp.Diagnostics.AddError("error preparing social provider payload ", fmt.Sprintf("Error: %+v ", diag.Errors()))
		return
	}
	model.ID = state.ID.ValueString()
	_, err := r.cidaasClient.SocialProvider.Upsert(model)
	if err != nil {
		resp.Diagnostics.AddError("failed to update social provider", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SocialProvider) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SocialProviderConfig
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	err := r.cidaasClient.SocialProvider.Delete(state.ProviderName.ValueString(), state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to delete social provider", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
}

func (r *SocialProvider) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	id := req.ID
	parts := strings.Split(id, ":")
	if len(parts) != 2 {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: 'provider_name:provider_id', got: %s", id),
		)
		return
	}
	providerName := parts[0]
	providerID := parts[1]
	if !util.StringInSlice(providerName, allowedProviders) {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Invalid provider_name provided in import identifier. Valid provider_names %+v, got: %s", allowedProviders, providerName),
		)
		return
	}
	resp.State.SetAttribute(ctx, path.Root("provider_name"), providerName)
	resp.State.SetAttribute(ctx, path.Root("id"), providerID)
}

func (sp *SocialProviderConfig) extract(ctx context.Context) diag.Diagnostics {
	var diags diag.Diagnostics
	if !sp.Claims.IsNull() {
		sp.claims = &Claims{}
		diags = sp.Claims.As(ctx, sp.claims, basetypes.ObjectAsOptions{})
		if !sp.claims.OptionalClaims.IsNull() {
			sp.claims.optionalClaims = &ClaimConfigs{}
			diags = sp.claims.OptionalClaims.As(ctx, &sp.claims.optionalClaims, basetypes.ObjectAsOptions{})
		}
		if !sp.claims.RequiredClaims.IsNull() {
			sp.claims.requiredClaims = &ClaimConfigs{}
			diags = sp.claims.RequiredClaims.As(ctx, &sp.claims.requiredClaims, basetypes.ObjectAsOptions{})
		}
	}
	if !sp.UserInfoFields.IsNull() {
		sp.userInfoFields = make([]*UserInfoFields, 0, len(sp.UserInfoFields.Elements()))
		diags = sp.UserInfoFields.ElementsAs(ctx, &sp.userInfoFields, false)
	}
	return diags
}

func prepareSocialProviderModel(ctx context.Context, plan SocialProviderConfig) (*cidaas.SocialProviderModel, diag.Diagnostics) {
	var sp cidaas.SocialProviderModel

	sp.Name = plan.Name.ValueString()
	sp.ProviderName = plan.ProviderName.ValueString()
	sp.ClientID = plan.ClientID.ValueString()
	sp.ClientSecret = plan.ClientSecret.ValueString()
	sp.Enabled = plan.Enabled.ValueBool()
	sp.EnabledForAdminPortal = plan.EnabledForAdminPortal.ValueBool()

	sp.Claims = &cidaas.ClaimsModel{}
	sp.Claims.RequiredClaims = cidaas.RequiredClaimsModel{}
	sp.Claims.OptionalClaims = cidaas.OptionalClaimsModel{}

	var diags diag.Diagnostics
	if !plan.Claims.IsNull() && !plan.claims.RequiredClaims.IsNull() {
		if !plan.claims.requiredClaims.UserInfo.IsNull() {
			diags = plan.claims.requiredClaims.UserInfo.ElementsAs(ctx, &sp.Claims.RequiredClaims.UserInfo, false)
			if diags.HasError() {
				return nil, diags
			}
		}
		if !plan.claims.requiredClaims.IDToken.IsNull() {
			diags = plan.claims.requiredClaims.IDToken.ElementsAs(ctx, &sp.Claims.RequiredClaims.IDToken, false)
			if diags.HasError() {
				return nil, diags
			}
		}
	}

	if !plan.Claims.IsNull() && !plan.claims.OptionalClaims.IsNull() {
		if !plan.claims.optionalClaims.UserInfo.IsNull() {
			diags = plan.claims.optionalClaims.UserInfo.ElementsAs(ctx, &sp.Claims.OptionalClaims.UserInfo, false)
			if diags.HasError() {
				return nil, diags
			}
		}
		if !plan.claims.optionalClaims.IDToken.IsNull() {
			diags = plan.claims.optionalClaims.IDToken.ElementsAs(ctx, &sp.Claims.OptionalClaims.IDToken, false)
			if diags.HasError() {
				return nil, diags
			}
		}
	}

	diags = plan.Scopes.ElementsAs(ctx, &sp.Scopes, false)
	if diags.HasError() {
		return nil, diags
	}

	if !plan.UserInfoFields.IsNull() {
		var userInfoModels []cidaas.UserInfoFieldsModel
		for _, v := range plan.userInfoFields {
			userInfoModels = append(userInfoModels, cidaas.UserInfoFieldsModel{
				IsCustomField: v.IsCustomField.ValueBool(),
				IsSystemField: v.IsSystemField.ValueBool(),
				InnerKey:      v.InnerKey.ValueString(),
				ExternalKey:   v.ExternalKey.ValueString(),
			})
		}
		sp.UserInfoFields = userInfoModels
	}
	return &sp, diags
}

func setClaimsInfo(state *SocialProviderConfig, claims *cidaas.ClaimsModel) diag.Diagnostics {
	var diags diag.Diagnostics
	if claims != nil {
		state.Claims, diags = types.ObjectValue(
			map[string]attr.Type{
				"required_claims": types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"user_info": types.SetType{ElemType: types.StringType},
						"id_token":  types.SetType{ElemType: types.StringType},
					},
				},
				"optional_claims": types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"user_info": types.SetType{ElemType: types.StringType},
						"id_token":  types.SetType{ElemType: types.StringType},
					},
				},
			},
			map[string]attr.Value{
				"required_claims": types.ObjectValueMust(map[string]attr.Type{
					"user_info": types.SetType{ElemType: types.StringType},
					"id_token":  types.SetType{ElemType: types.StringType},
				},
					map[string]attr.Value{
						"user_info": util.SetValueOrNull(claims.RequiredClaims.UserInfo),
						"id_token":  util.SetValueOrNull(claims.RequiredClaims.IDToken),
					}),
				"optional_claims": types.ObjectValueMust(map[string]attr.Type{
					"user_info": types.SetType{ElemType: types.StringType},
					"id_token":  types.SetType{ElemType: types.StringType},
				},
					map[string]attr.Value{
						"user_info": util.SetValueOrNull(claims.OptionalClaims.UserInfo),
						"id_token":  util.SetValueOrNull(claims.OptionalClaims.IDToken),
					}),
			})
	}
	return diags
}

func setUserInfoFields(state *SocialProviderConfig, ufs []cidaas.UserInfoFieldsModel) diag.Diagnostics {
	var diags diag.Diagnostics
	ufObjectType := types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"inner_key":       types.StringType,
			"external_key":    types.StringType,
			"is_custom_field": types.BoolType,
			"is_system_field": types.BoolType,
		},
	}
	if len(ufs) < 1 {
		state.UserInfoFields = types.ListNull(ufObjectType)
	} else {
		var ufObjectValues []attr.Value
		for _, uf := range ufs {
			objValue := types.ObjectValueMust(
				ufObjectType.AttrTypes,
				map[string]attr.Value{
					"inner_key":       util.StringValueOrNull(&uf.InnerKey),
					"external_key":    util.StringValueOrNull(&uf.ExternalKey),
					"is_custom_field": util.BoolValueOrNull(&uf.IsCustomField),
					"is_system_field": util.BoolValueOrNull(&uf.IsSystemField),
				})
			ufObjectValues = append(ufObjectValues, objValue)
		}
		state.UserInfoFields, diags = types.ListValue(ufObjectType, ufObjectValues)
	}
	return diags
}
