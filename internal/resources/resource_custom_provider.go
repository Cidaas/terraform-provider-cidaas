package resources

import (
	"context"
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

type CustomProvider struct {
	BaseResource
}

func NewCustomProvider() resource.Resource {
	return &CustomProvider{
		BaseResource: NewBaseResource(
			BaseResourceConfig{
				Name:   RESOURCE_CUSTOM_PROVIDER,
				Schema: &customProviderSchema,
			},
		),
	}
}

type ProviderConfig struct {
	ID                    types.String `tfsdk:"id"`
	ProviderName          types.String `tfsdk:"provider_name"`
	DisplayName           types.String `tfsdk:"display_name"`
	LogoURL               types.String `tfsdk:"logo_url"`
	StandardType          types.String `tfsdk:"standard_type"`
	ClientID              types.String `tfsdk:"client_id"`
	ClientSecret          types.String `tfsdk:"client_secret"`
	AuthorizationEndpoint types.String `tfsdk:"authorization_endpoint"`
	TokenEndpoint         types.String `tfsdk:"token_endpoint"`
	UserinfoEndpoint      types.String `tfsdk:"userinfo_endpoint"`
	ScopeDisplayLabel     types.String `tfsdk:"scope_display_label"`
	Domains               types.Set    `tfsdk:"domains"`
	// Scopes                []Scope       `tfsdk:"scopes"`
	// UserinfoFields        UserInfoField `tfsdk:"userinfo_fields"`
	Scopes         types.List   `tfsdk:"scopes"`
	UserinfoFields types.Object `tfsdk:"userinfo_fields"`
	scopes         []*CpScope
	userinfoFields *UserInfoField
	AmrConfig      types.List   `tfsdk:"amr_config"`
	UserinfoSource types.String `tfsdk:"userinfo_source"`
}

type CpScope struct {
	ScopeName   types.String `tfsdk:"scope_name"`
	Required    types.Bool   `tfsdk:"required"`
	Recommended types.Bool   `tfsdk:"recommended"`
}

type UfNestedObject struct {
	ExtFieldKey types.String `tfsdk:"ext_field_key"`
	Default     types.String `tfsdk:"default"`
}

type UfEmailVerifiedNestedObject struct {
	ExtFieldKey types.String `tfsdk:"ext_field_key"`
	Default     types.Bool   `tfsdk:"default"`
}

type UserInfoField struct {
	Name              types.Object `tfsdk:"name"`
	FamilyName        types.Object `tfsdk:"family_name"`
	GivenName         types.Object `tfsdk:"given_name"`
	MiddleName        types.Object `tfsdk:"middle_name"`
	Nickname          types.Object `tfsdk:"nickname"`
	PreferredUsername types.Object `tfsdk:"preferred_username"`
	Profile           types.Object `tfsdk:"profile"`
	Picture           types.Object `tfsdk:"picture"`
	Website           types.Object `tfsdk:"website"`
	Gender            types.Object `tfsdk:"gender"`
	Birthdate         types.Object `tfsdk:"birthdate"`
	Zoneinfo          types.Object `tfsdk:"zoneinfo"`
	Locale            types.Object `tfsdk:"locale"`
	UpdatedAt         types.Object `tfsdk:"updated_at"`
	Email             types.Object `tfsdk:"email"`
	EmailVerified     types.Object `tfsdk:"email_verified"`
	PhoneNumber       types.Object `tfsdk:"phone_number"`
	MobileNumber      types.Object `tfsdk:"mobile_number"`
	Address           types.Object `tfsdk:"address"`
	Sub               types.Object `tfsdk:"sub"`

	name              *UfNestedObject
	familyName        *UfNestedObject
	givenName         *UfNestedObject
	middleName        *UfNestedObject
	nickname          *UfNestedObject
	preferredUsername *UfNestedObject
	profile           *UfNestedObject
	picture           *UfNestedObject
	website           *UfNestedObject
	gender            *UfNestedObject
	birthdate         *UfNestedObject
	zoneinfo          *UfNestedObject
	locale            *UfNestedObject
	updatedAt         *UfNestedObject
	email             *UfNestedObject
	emailVerified     *UfEmailVerifiedNestedObject
	phoneNumber       *UfNestedObject
	mobileNumber      *UfNestedObject
	address           *UfNestedObject
	sub               *UfNestedObject

	CustomFields types.Map `tfsdk:"custom_fields"`
}

type AmrConfig struct {
	AmrValue    types.String `tfsdk:"amr_value"`
	ExtAmrValue types.String `tfsdk:"ext_amr_value"`
}

func (pc *ProviderConfig) extract(ctx context.Context) diag.Diagnostics {
	var diags diag.Diagnostics
	if !pc.UserinfoFields.IsNull() {
		pc.userinfoFields = &UserInfoField{}
		diags = pc.UserinfoFields.As(ctx, pc.userinfoFields, basetypes.ObjectAsOptions{})

		extractField := func(field types.Object, target **UfNestedObject) {
			if !field.IsNull() && !field.IsUnknown() {
				*target = &UfNestedObject{}
				diags.Append(field.As(ctx, *target, basetypes.ObjectAsOptions{})...)
			}
		}

		extractField(pc.userinfoFields.Name, &pc.userinfoFields.name)
		extractField(pc.userinfoFields.FamilyName, &pc.userinfoFields.familyName)
		extractField(pc.userinfoFields.GivenName, &pc.userinfoFields.givenName)
		extractField(pc.userinfoFields.MiddleName, &pc.userinfoFields.middleName)
		extractField(pc.userinfoFields.Nickname, &pc.userinfoFields.nickname)
		extractField(pc.userinfoFields.PreferredUsername, &pc.userinfoFields.preferredUsername)
		extractField(pc.userinfoFields.Profile, &pc.userinfoFields.profile)
		extractField(pc.userinfoFields.Picture, &pc.userinfoFields.picture)
		extractField(pc.userinfoFields.Website, &pc.userinfoFields.website)
		extractField(pc.userinfoFields.Gender, &pc.userinfoFields.gender)
		extractField(pc.userinfoFields.Birthdate, &pc.userinfoFields.birthdate)
		extractField(pc.userinfoFields.Zoneinfo, &pc.userinfoFields.zoneinfo)
		extractField(pc.userinfoFields.Locale, &pc.userinfoFields.locale)
		extractField(pc.userinfoFields.UpdatedAt, &pc.userinfoFields.updatedAt)
		extractField(pc.userinfoFields.Email, &pc.userinfoFields.email)
		extractField(pc.userinfoFields.PhoneNumber, &pc.userinfoFields.phoneNumber)
		extractField(pc.userinfoFields.MobileNumber, &pc.userinfoFields.mobileNumber)
		extractField(pc.userinfoFields.Address, &pc.userinfoFields.address)
		extractField(pc.userinfoFields.Sub, &pc.userinfoFields.sub)

		if !pc.userinfoFields.EmailVerified.IsNull() && !pc.userinfoFields.EmailVerified.IsUnknown() {
			pc.userinfoFields.emailVerified = &UfEmailVerifiedNestedObject{}
			diags.Append(pc.userinfoFields.EmailVerified.As(ctx, &pc.userinfoFields.emailVerified, basetypes.ObjectAsOptions{})...)
		}
	}

	if !pc.Scopes.IsNull() {
		pc.scopes = make([]*CpScope, 0, len(pc.Scopes.Elements()))
		diags.Append(pc.Scopes.ElementsAs(ctx, &pc.scopes, false)...)
	}
	return diags
}

var customProviderSchema = schema.Schema{
	MarkdownDescription: "This example demonstrates the configuration of a custom provider resource for interacting with Cidaas." +
		"\n\n Ensure that the below scopes are assigned to the client with the specified `client_id`:" +
		"\n- cidaas:providers_read" +
		"\n- cidaas:providers_write" +
		"\n- cidaas:providers_delete",
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The ID of the resource.",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"provider_name": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The unique identifier of the custom provider. This cannot be updated for an existing state.",
			PlanModifiers: []planmodifier.String{
				&validators.UniqueIdentifier{},
			},
		},
		"display_name": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The display name of the provider.",
		},
		"logo_url": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: "The URL for the provider's logo.",
		},
		"standard_type": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: "Type of standard. Allowed values `OAUTH2` and `OPENID_CONNECT`.",
			Validators: []validator.String{
				stringvalidator.OneOf([]string{"OPENID_CONNECT", "OAUTH2"}...),
			},
		},
		"client_id": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The client ID of the provider.",
		},
		"client_secret": schema.StringAttribute{
			Required:            true,
			Sensitive:           true,
			MarkdownDescription: "The client secret of the provider.",
		},
		"authorization_endpoint": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The URL for authorization of the provider.",
		},
		"token_endpoint": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The URL to generate token with this provider.",
		},
		"userinfo_endpoint": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The URL to fetch user details using this provider.",
		},
		// In plan Set deletes an existing record and create a whole new one, so preferred list. However, to allow only unique values use set
		"scopes": schema.ListNestedAttribute{
			MarkdownDescription: "List of scopes of the provider with details",
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"scope_name": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "The name of the scope, e.g., `openid`, `profile`.",
					},
					"required": schema.BoolAttribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Indicates if the scope is required.",
						Default:             booldefault.StaticBool(false),
					},
					"recommended": schema.BoolAttribute{
						MarkdownDescription: "Indicates if the scope is recommended.",
						Optional:            true,
						Computed:            true,
						Default:             booldefault.StaticBool(false),
					},
				},
			},
			Optional: true,
		},
		"scope_display_label": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: "Display label for the scope of the provider.",
		},
		"userinfo_fields": schema.SingleNestedAttribute{
			Optional: true,
			Computed: true,
			MarkdownDescription: "Object containing various user information fields with their values." +
				" The userinfo_fields section includes specific fields such as name, family_name, address, etc., along with custom_fields allowing additional user information customization",
			Attributes: map[string]schema.Attribute{
				"name":               createStandardNestedAttribute(),
				"family_name":        createStandardNestedAttribute(),
				"given_name":         createStandardNestedAttribute(),
				"middle_name":        createStandardNestedAttribute(),
				"nickname":           createStandardNestedAttribute(),
				"preferred_username": createStandardNestedAttribute(),
				"profile":            createStandardNestedAttribute(),
				"picture":            createStandardNestedAttribute(),
				"website":            createStandardNestedAttribute(),
				"gender":             createStandardNestedAttribute(),
				"birthdate":          createStandardNestedAttribute(),
				"zoneinfo":           createStandardNestedAttribute(),
				"locale":             createStandardNestedAttribute(),
				"updated_at":         createStandardNestedAttribute(),
				"email":              createStandardNestedAttribute(),
				"phone_number":       createStandardNestedAttribute(),
				"mobile_number":      createStandardNestedAttribute(),
				"address":            createStandardNestedAttribute(),
				"sub":                createStandardNestedAttribute(),

				"email_verified": schema.SingleNestedAttribute{
					Optional: true,
					Attributes: map[string]schema.Attribute{
						"ext_field_key": schema.StringAttribute{
							Optional: true,
						},
						"default": schema.BoolAttribute{
							Optional: true,
							Computed: true,
							Default:  booldefault.StaticBool(true),
						},
					},
				},

				"custom_fields": schema.MapAttribute{
					ElementType: types.StringType,
					Optional:    true,
				},
			},
			Default: objectdefault.StaticValue(userInfoDefaultValue()),
		},
		"domains": schema.SetAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The domains of the provider.",
			Optional:            true,
		},
		"amr_config": schema.ListNestedAttribute{
			Optional:            true,
			MarkdownDescription: "AMR configuration mapping.",
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"amr_value": schema.StringAttribute{
						Required: true,
					},
					"ext_amr_value": schema.StringAttribute{
						Required: true,
					},
				},
			},
		},
		"userinfo_source": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: "Source of userinfo. Allowed values are `IDTOKEN` and `USERINFOENDPOINT`.",
			Validators: []validator.String{
				stringvalidator.OneOf([]string{"IDTOKEN", "USERINFOENDPOINT"}...),
			},
		},
	},
}

func (r *CustomProvider) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ProviderConfig
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(plan.extract(ctx)...)
	cp, d := prepareCpRequestPayload(ctx, plan)
	resp.Diagnostics.Append(d...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.cidaasClient.CustomProvider.CreateCustomProvider(cp)
	if err != nil {
		resp.Diagnostics.AddError("failed to create custom provider", util.FormatErrorMessage(err))
		return
	}
	plan.ID = util.StringValueOrNull(&res.Data.ID)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *CustomProvider) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ProviderConfig
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	res, err := r.cidaasClient.CustomProvider.GetCustomProvider(state.ProviderName.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to read custom provider", util.FormatErrorMessage(err))
		return
	}
	state.StandardType = util.StringValueOrNull(&res.Data.StandardType)
	state.AuthorizationEndpoint = util.StringValueOrNull(&res.Data.AuthorizationEndpoint)
	state.TokenEndpoint = util.StringValueOrNull(&res.Data.TokenEndpoint)
	state.ProviderName = util.StringValueOrNull(&res.Data.ProviderName)
	state.DisplayName = util.StringValueOrNull(&res.Data.DisplayName)
	state.LogoURL = util.StringValueOrNull(&res.Data.LogoURL)
	state.UserinfoEndpoint = util.StringValueOrNull(&res.Data.UserinfoEndpoint)
	state.ID = util.StringValueOrNull(&res.Data.ID)
	state.ScopeDisplayLabel = util.StringValueOrNull(&res.Data.Scopes.DisplayLabel)
	state.ClientID = util.StringValueOrNull(&res.Data.ClientID)
	state.ClientSecret = util.StringValueOrNull(&res.Data.ClientSecret)
	state.Domains = util.SetValueOrNull(res.Data.Domains)

	var diag diag.Diagnostics
	var objectValues []attr.Value
	scopeObjectType := types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"scope_name":  types.StringType,
			"required":    types.BoolType,
			"recommended": types.BoolType,
		},
	}

	for _, sc := range res.Data.Scopes.Scopes {
		scopeName := sc.ScopeName
		required := sc.Required
		recommended := sc.Recommended
		objValue := types.ObjectValueMust(scopeObjectType.AttrTypes, map[string]attr.Value{
			"scope_name":  util.StringValueOrNull(&scopeName),
			"required":    util.BoolValueOrNull(&required),
			"recommended": util.BoolValueOrNull(&recommended),
		})
		objectValues = append(objectValues, objValue)
	}

	state.Scopes, diag = types.ListValueFrom(ctx, scopeObjectType, objectValues)
	resp.Diagnostics.Append(diag...)
	if resp.Diagnostics.HasError() {
		return
	}

	standardFields := []string{
		"name", "family_name", "given_name", "middle_name", "nickname",
		"preferred_username", "profile", "picture", "website", "gender",
		"birthdate", "zoneinfo", "locale", "updated_at", "email",
		"phone_number", "mobile_number", "address", "sub",
	}

	metadataAttributeTypes := make(map[string]attr.Type)
	metadataAttributes := make(map[string]attr.Value)

	standardNestedType := types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ext_field_key": types.StringType,
			"default":       types.StringType,
		},
	}

	emailVerifiedType := types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ext_field_key": types.StringType,
			"default":       types.BoolType,
		},
	}

	for _, field := range standardFields {
		metadataAttributeTypes[field] = standardNestedType
		metadataAttributes[field] = types.ObjectNull(standardNestedType.AttrTypes)
	}

	metadataAttributeTypes["email_verified"] = emailVerifiedType
	metadataAttributes["email_verified"] = types.ObjectNull(emailVerifiedType.AttrTypes)

	customFields := map[string]attr.Value{}
	hasCustomfield := false

	for key, fieldInterface := range res.Data.UserinfoFields {
		if strings.HasPrefix(key, "customFields.") {
			if fieldMap, ok := fieldInterface.(map[string]interface{}); ok {
				customFieldKey := strings.TrimPrefix(key, "customFields.")
				if extFieldKey, exists := fieldMap["extFieldKey"].(string); exists {
					customFields[customFieldKey] = types.StringValue(extFieldKey)
					hasCustomfield = true
				}
			}
			continue
		}

		if _, exists := metadataAttributeTypes[key]; exists {
			if key == "email_verified" {
				if fieldMap, ok := fieldInterface.(map[string]interface{}); ok {
					extFieldKey, _ := fieldMap["extFieldKey"].(string)
					defaultValue, _ := fieldMap["default"].(bool)
					metadataAttributes[key] = types.ObjectValueMust(
						metadataAttributeTypes[key].(types.ObjectType).AttrTypes,
						map[string]attr.Value{
							"ext_field_key": types.StringValue(extFieldKey),
							"default":       types.BoolValue(defaultValue),
						},
					)
				}
			} else {
				if fieldMap, ok := fieldInterface.(map[string]interface{}); ok {
					extFieldKey, _ := fieldMap["extFieldKey"].(string)
					defaultValue, hasDefault := fieldMap["default"].(string)

					var defaultAttrValue attr.Value
					if hasDefault && defaultValue != "" {
						defaultAttrValue = types.StringValue(defaultValue)
					} else {
						defaultAttrValue = types.StringNull()
					}

					metadataAttributes[key] = types.ObjectValueMust(
						metadataAttributeTypes[key].(types.ObjectType).AttrTypes,
						map[string]attr.Value{
							"ext_field_key": types.StringValue(extFieldKey),
							"default":       defaultAttrValue,
						},
					)
				}
			}
		}
	}

	metadataAttributeTypes["custom_fields"] = types.MapType{ElemType: types.StringType}
	if hasCustomfield {
		metadataAttributes["custom_fields"] = types.MapValueMust(types.StringType, customFields)
	} else {
		metadataAttributes["custom_fields"] = types.MapNull(types.StringType)
	}

	state.UserinfoFields, diag = types.ObjectValue(metadataAttributeTypes, metadataAttributes)
	resp.Diagnostics.Append(diag...)
	if resp.Diagnostics.HasError() {
		return
	}

	if len(res.Data.AmrConfig) > 0 {
		amrObjectType := types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"amr_value":     types.StringType,
				"ext_amr_value": types.StringType,
			},
		}

		amrValues := make([]attr.Value, 0, len(res.Data.AmrConfig))
		for _, amrConfig := range res.Data.AmrConfig {
			amrValues = append(amrValues, types.ObjectValueMust(
				amrObjectType.AttrTypes,
				map[string]attr.Value{
					"amr_value":     types.StringValue(amrConfig.AmrValue),
					"ext_amr_value": types.StringValue(amrConfig.ExtAmrValue),
				},
			))
		}

		state.AmrConfig, diag = types.ListValueFrom(ctx, amrObjectType, amrValues)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		state.AmrConfig = types.ListNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"amr_value":     types.StringType,
				"ext_amr_value": types.StringType,
			},
		})
	}

	state.UserinfoSource = util.StringValueOrNull(&res.Data.UserInfoSource)

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *CustomProvider) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ProviderConfig

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(plan.extract(ctx)...)

	cp, d := prepareCpRequestPayload(ctx, plan)
	resp.Diagnostics.Append(d...)
	if resp.Diagnostics.HasError() {
		return
	}
	cp.ID = state.ID.ValueString()
	err := r.cidaasClient.CustomProvider.UpdateCustomProvider(cp)
	if err != nil {
		resp.Diagnostics.AddError("failed to update custom provider", util.FormatErrorMessage(err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *CustomProvider) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ProviderConfig
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	err := r.cidaasClient.CustomProvider.DeleteCustomProvider(state.ProviderName.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to delete custom provier", util.FormatErrorMessage(err))
		return
	}
}

func (r *CustomProvider) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("provider_name"), req, resp)
}

func prepareCpRequestPayload(ctx context.Context, plan ProviderConfig) (*cidaas.CustomProviderModel, diag.Diagnostics) {
	var cp cidaas.CustomProviderModel
	var diags diag.Diagnostics

	cp.StandardType = plan.StandardType.ValueString()
	cp.AuthorizationEndpoint = plan.AuthorizationEndpoint.ValueString()
	cp.TokenEndpoint = plan.TokenEndpoint.ValueString()
	cp.ProviderName = plan.ProviderName.ValueString()
	cp.DisplayName = plan.DisplayName.ValueString()
	cp.LogoURL = plan.LogoURL.ValueString()
	cp.UserinfoEndpoint = plan.UserinfoEndpoint.ValueString()
	cp.ClientID = plan.ClientID.ValueString()
	cp.ClientSecret = plan.ClientSecret.ValueString()

	diags = plan.Domains.ElementsAs(ctx, &cp.Domains, false)
	if diags.HasError() {
		return nil, diags
	}
	var childScopes []cidaas.ScopeChild
	for _, v := range plan.scopes {
		childScopes = append(childScopes, cidaas.ScopeChild{
			ScopeName:   v.ScopeName.ValueString(),
			Required:    v.Required.ValueBool(),
			Recommended: v.Recommended.ValueBool(),
		})
	}
	cp.Scopes.Scopes = childScopes
	cp.Scopes.DisplayLabel = plan.ScopeDisplayLabel.ValueString()

	if !plan.UserinfoFields.IsNull() {
		cp.UserinfoFields = make(map[string]interface{})

		addUserInfoField := func(fieldName string, obj types.Object, nestedObj *UfNestedObject) {
			if !obj.IsNull() && nestedObj != nil && !nestedObj.ExtFieldKey.IsNull() {
				cp.UserinfoFields[fieldName] = &cidaas.UserInfoField{
					ExtFieldKey: nestedObj.ExtFieldKey.ValueString(),
					Default:     nestedObj.Default.ValueString(),
				}
			}
		}

		addUserInfoField("name", plan.userinfoFields.Name, plan.userinfoFields.name)
		addUserInfoField("family_name", plan.userinfoFields.FamilyName, plan.userinfoFields.familyName)
		addUserInfoField("given_name", plan.userinfoFields.GivenName, plan.userinfoFields.givenName)
		addUserInfoField("middle_name", plan.userinfoFields.MiddleName, plan.userinfoFields.middleName)
		addUserInfoField("nickname", plan.userinfoFields.Nickname, plan.userinfoFields.nickname)
		addUserInfoField("preferred_username", plan.userinfoFields.PreferredUsername, plan.userinfoFields.preferredUsername)
		addUserInfoField("profile", plan.userinfoFields.Profile, plan.userinfoFields.profile)
		addUserInfoField("picture", plan.userinfoFields.Picture, plan.userinfoFields.picture)
		addUserInfoField("website", plan.userinfoFields.Website, plan.userinfoFields.website)
		addUserInfoField("gender", plan.userinfoFields.Gender, plan.userinfoFields.gender)
		addUserInfoField("birthdate", plan.userinfoFields.Birthdate, plan.userinfoFields.birthdate)
		addUserInfoField("zoneinfo", plan.userinfoFields.Zoneinfo, plan.userinfoFields.zoneinfo)
		addUserInfoField("locale", plan.userinfoFields.Locale, plan.userinfoFields.locale)
		addUserInfoField("updated_at", plan.userinfoFields.UpdatedAt, plan.userinfoFields.updatedAt)
		addUserInfoField("email", plan.userinfoFields.Email, plan.userinfoFields.email)
		addUserInfoField("phone_number", plan.userinfoFields.PhoneNumber, plan.userinfoFields.phoneNumber)
		addUserInfoField("mobile_number", plan.userinfoFields.MobileNumber, plan.userinfoFields.mobileNumber)
		addUserInfoField("address", plan.userinfoFields.Address, plan.userinfoFields.address)
		addUserInfoField("sub", plan.userinfoFields.Sub, plan.userinfoFields.sub)

		if !plan.userinfoFields.EmailVerified.IsNull() &&
			plan.userinfoFields.emailVerified != nil &&
			!plan.userinfoFields.emailVerified.ExtFieldKey.IsNull() {
			cp.UserinfoFields["email_verified"] = &cidaas.UserInfoFieldBoolean{
				ExtFieldKey: plan.userinfoFields.emailVerified.ExtFieldKey.ValueString(),
				Default:     plan.userinfoFields.emailVerified.Default.ValueBool(),
			}
		}

		if !plan.userinfoFields.CustomFields.IsNull() {
			var cfMap map[string]string
			diags.Append(plan.userinfoFields.CustomFields.ElementsAs(ctx, &cfMap, false)...)
			if diags.HasError() {
				return nil, diags
			}

			for key, value := range cfMap {
				cp.UserinfoFields["customFields."+key] = &cidaas.UserInfoField{
					ExtFieldKey: value,
				}
			}
		}
	}

	if !plan.AmrConfig.IsNull() {
		var amrConfigs []AmrConfig
		diags.Append(plan.AmrConfig.ElementsAs(ctx, &amrConfigs, false)...)
		if diags.HasError() {
			return nil, diags
		}

		cp.AmrConfig = make([]cidaas.AmrConfig, len(amrConfigs))
		for i, config := range amrConfigs {
			cp.AmrConfig[i] = cidaas.AmrConfig{
				AmrValue:    config.AmrValue.ValueString(),
				ExtAmrValue: config.ExtAmrValue.ValueString(),
			}
		}
	}

	if !plan.UserinfoSource.IsNull() {
		cp.UserInfoSource = plan.UserinfoSource.ValueString()
	}

	return &cp, diags
}

func userInfoDefaultValue() basetypes.ObjectValue {
	standardFieldType := types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ext_field_key": types.StringType,
			"default":       types.StringType,
		},
	}

	emailVerifiedType := types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ext_field_key": types.StringType,
			"default":       types.BoolType,
		},
	}

	attrTypes := map[string]attr.Type{
		"name":               standardFieldType,
		"family_name":        standardFieldType,
		"given_name":         standardFieldType,
		"middle_name":        standardFieldType,
		"nickname":           standardFieldType,
		"preferred_username": standardFieldType,
		"profile":            standardFieldType,
		"picture":            standardFieldType,
		"website":            standardFieldType,
		"gender":             standardFieldType,
		"birthdate":          standardFieldType,
		"zoneinfo":           standardFieldType,
		"locale":             standardFieldType,
		"updated_at":         standardFieldType,
		"email":              standardFieldType,
		"email_verified":     emailVerifiedType,
		"phone_number":       standardFieldType,
		"mobile_number":      standardFieldType,
		"address":            standardFieldType,
		"sub":                standardFieldType,
		"custom_fields":      types.MapType{ElemType: types.StringType},
	}

	assignedValue := func(value string) basetypes.ObjectValue {
		standardFieldValue := types.ObjectValueMust(
			standardFieldType.AttrTypes,
			map[string]attr.Value{
				"ext_field_key": types.StringValue(value),
				"default":       types.StringNull(),
			},
		)
		return standardFieldValue
	}

	attrValues := map[string]attr.Value{
		"name":               assignedValue("name"),
		"family_name":        assignedValue("family_name"),
		"given_name":         assignedValue("given_name"),
		"middle_name":        assignedValue("middle_name"),
		"nickname":           assignedValue("nickname"),
		"preferred_username": assignedValue("preferred_username"),
		"profile":            assignedValue("profile"),
		"picture":            assignedValue("picture"),
		"website":            assignedValue("website"),
		"gender":             assignedValue("gender"),
		"birthdate":          assignedValue("birthdate"),
		"zoneinfo":           assignedValue("zoneinfo"),
		"locale":             assignedValue("locale"),
		"updated_at":         assignedValue("updated_at"),
		"email":              assignedValue("email"),
		"email_verified": types.ObjectValueMust(
			emailVerifiedType.AttrTypes,
			map[string]attr.Value{
				"ext_field_key": types.StringValue("email_verified"),
				"default":       types.BoolValue(true),
			},
		),
		"phone_number":  assignedValue("phone_number"),
		"mobile_number": assignedValue("mobile_number"),
		"address":       assignedValue("address"),
		"sub":           assignedValue("sub"),
		"custom_fields": types.MapNull(types.StringType),
	}

	return types.ObjectValueMust(attrTypes, attrValues)
}

func createStandardNestedAttribute() schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Optional: true,
		Attributes: map[string]schema.Attribute{
			"ext_field_key": schema.StringAttribute{
				Required: true,
			},
			"default": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}
