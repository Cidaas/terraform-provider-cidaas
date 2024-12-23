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
}

type CpScope struct {
	ScopeName   types.String `tfsdk:"scope_name"`
	Required    types.Bool   `tfsdk:"required"`
	Recommended types.Bool   `tfsdk:"recommended"`
}

type UserInfoField struct {
	Name              types.String `tfsdk:"name"`
	FamilyName        types.String `tfsdk:"family_name"`
	GivenName         types.String `tfsdk:"given_name"`
	MiddleName        types.String `tfsdk:"middle_name"`
	Nickname          types.String `tfsdk:"nickname"`
	PreferredUsername types.String `tfsdk:"preferred_username"`
	Profile           types.String `tfsdk:"profile"`
	Picture           types.String `tfsdk:"picture"`
	Website           types.String `tfsdk:"website"`
	Gender            types.String `tfsdk:"gender"`
	Birthdate         types.String `tfsdk:"birthdate"`
	Zoneinfo          types.String `tfsdk:"zoneinfo"`
	Locale            types.String `tfsdk:"locale"`
	UpdatedAt         types.String `tfsdk:"updated_at"`
	Email             types.String `tfsdk:"email"`
	EmailVerified     types.String `tfsdk:"email_verified"`
	PhoneNumber       types.String `tfsdk:"phone_number"`
	MobileNumber      types.String `tfsdk:"mobile_number"`
	Address           types.String `tfsdk:"address"`
	Sub               types.String `tfsdk:"sub"`
	CustomFields      types.Map    `tfsdk:"custom_fields"`
}

func (pc *ProviderConfig) extract(ctx context.Context) diag.Diagnostics {
	var diags diag.Diagnostics
	if !pc.UserinfoFields.IsNull() {
		pc.userinfoFields = &UserInfoField{}
		diags = pc.UserinfoFields.As(ctx, pc.userinfoFields, basetypes.ObjectAsOptions{})
	}
	if !pc.Scopes.IsNull() {
		pc.scopes = make([]*CpScope, 0, len(pc.Scopes.Elements()))
		diags = pc.Scopes.ElementsAs(ctx, &pc.scopes, false)
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
			Required:            true,
			MarkdownDescription: "Display label for the scope of the provider.",
		},
		"userinfo_fields": schema.SingleNestedAttribute{
			Optional: true,
			Computed: true,
			MarkdownDescription: "Object containing various user information fields with their values." +
				" The userinfo_fields section includes specific fields such as name, family_name, address, etc., along with custom_fields allowing additional user information customization",
			Attributes: map[string]schema.Attribute{
				"name": schema.StringAttribute{
					Optional: true,
				},
				"family_name": schema.StringAttribute{
					Optional: true,
				},
				"given_name": schema.StringAttribute{
					Optional: true,
				},
				"middle_name": schema.StringAttribute{
					Optional: true,
				},
				"nickname": schema.StringAttribute{
					Optional: true,
				},
				"preferred_username": schema.StringAttribute{
					Optional: true,
				},
				"profile": schema.StringAttribute{
					Optional: true,
				},
				"picture": schema.StringAttribute{
					Optional: true,
				},
				"website": schema.StringAttribute{
					Optional: true,
				},
				"gender": schema.StringAttribute{
					Optional: true,
				},
				"birthdate": schema.StringAttribute{
					Optional: true,
				},
				"zoneinfo": schema.StringAttribute{
					Optional: true,
				},
				"locale": schema.StringAttribute{
					Optional: true,
				},
				"updated_at": schema.StringAttribute{
					Optional: true,
				},
				"email": schema.StringAttribute{
					Optional: true,
				},
				"email_verified": schema.StringAttribute{
					Optional: true,
				},
				"phone_number": schema.StringAttribute{
					Optional: true,
				},
				"mobile_number": schema.StringAttribute{
					Optional: true,
				},
				"address": schema.StringAttribute{
					Optional: true,
				},
				"sub": schema.StringAttribute{
					Optional: true,
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

	metadataAttributeTypes := map[string]attr.Type{
		"name":               types.StringType,
		"family_name":        types.StringType,
		"given_name":         types.StringType,
		"middle_name":        types.StringType,
		"nickname":           types.StringType,
		"preferred_username": types.StringType,
		"profile":            types.StringType,
		"picture":            types.StringType,
		"website":            types.StringType,
		"gender":             types.StringType,
		"birthdate":          types.StringType,
		"zoneinfo":           types.StringType,
		"locale":             types.StringType,
		"updated_at":         types.StringType,
		"email":              types.StringType,
		"email_verified":     types.StringType,
		"phone_number":       types.StringType,
		"mobile_number":      types.StringType,
		"address":            types.StringType,
		"sub":                types.StringType,
	}

	metadataAttributes := map[string]attr.Value{
		"name":               types.StringNull(),
		"family_name":        types.StringNull(),
		"given_name":         types.StringNull(),
		"middle_name":        types.StringNull(),
		"nickname":           types.StringNull(),
		"preferred_username": types.StringNull(),
		"profile":            types.StringNull(),
		"picture":            types.StringNull(),
		"website":            types.StringNull(),
		"gender":             types.StringNull(),
		"birthdate":          types.StringNull(),
		"zoneinfo":           types.StringNull(),
		"locale":             types.StringNull(),
		"updated_at":         types.StringNull(),
		"email":              types.StringNull(),
		"email_verified":     types.StringNull(),
		"phone_number":       types.StringNull(),
		"mobile_number":      types.StringNull(),
		"address":            types.StringNull(),
		"sub":                types.StringNull(),
	}

	customFields := map[string]attr.Value{}
	hasCustomfield := false

	for key, field := range res.Data.UserinfoFields {
		if strings.HasPrefix(key, "customFields.") {
			customFieldKey := strings.TrimPrefix(key, "customFields.")
			customFields[customFieldKey] = types.StringValue(field.ExtFieldKey)
			hasCustomfield = true
		} else if _, exists := metadataAttributes[key]; exists {
			metadataAttributes[key] = types.StringValue(field.ExtFieldKey)
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

	cp.StandardType = plan.StandardType.ValueString()
	cp.AuthorizationEndpoint = plan.AuthorizationEndpoint.ValueString()
	cp.TokenEndpoint = plan.TokenEndpoint.ValueString()
	cp.ProviderName = plan.ProviderName.ValueString()
	cp.DisplayName = plan.DisplayName.ValueString()
	cp.LogoURL = plan.LogoURL.ValueString()
	cp.UserinfoEndpoint = plan.UserinfoEndpoint.ValueString()
	cp.ClientID = plan.ClientID.ValueString()
	cp.ClientSecret = plan.ClientSecret.ValueString()

	diag := plan.Domains.ElementsAs(ctx, &cp.Domains, false)
	if diag.HasError() {
		return nil, diag
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
		// Initialize the map
		cp.UserinfoFields = make(map[string]*cidaas.UserInfoField)

		// Add standard fields
		if !plan.userinfoFields.Name.IsNull() {
			cp.UserinfoFields["name"] = &cidaas.UserInfoField{ExtFieldKey: plan.userinfoFields.Name.ValueString()}
		}
		if !plan.userinfoFields.FamilyName.IsNull() {
			cp.UserinfoFields["family_name"] = &cidaas.UserInfoField{ExtFieldKey: plan.userinfoFields.FamilyName.ValueString()}
		}
		if !plan.userinfoFields.GivenName.IsNull() {
			cp.UserinfoFields["given_name"] = &cidaas.UserInfoField{ExtFieldKey: plan.userinfoFields.GivenName.ValueString()}
		}
		if !plan.userinfoFields.MiddleName.IsNull() {
			cp.UserinfoFields["middle_name"] = &cidaas.UserInfoField{ExtFieldKey: plan.userinfoFields.MiddleName.ValueString()}
		}
		if !plan.userinfoFields.Nickname.IsNull() {
			cp.UserinfoFields["nickname"] = &cidaas.UserInfoField{ExtFieldKey: plan.userinfoFields.Nickname.ValueString()}
		}
		if !plan.userinfoFields.PreferredUsername.IsNull() {
			cp.UserinfoFields["preferred_username"] = &cidaas.UserInfoField{ExtFieldKey: plan.userinfoFields.PreferredUsername.ValueString()}
		}
		if !plan.userinfoFields.Profile.IsNull() {
			cp.UserinfoFields["profile"] = &cidaas.UserInfoField{ExtFieldKey: plan.userinfoFields.Profile.ValueString()}
		}
		if !plan.userinfoFields.Picture.IsNull() {
			cp.UserinfoFields["picture"] = &cidaas.UserInfoField{ExtFieldKey: plan.userinfoFields.Picture.ValueString()}
		}
		if !plan.userinfoFields.Website.IsNull() {
			cp.UserinfoFields["website"] = &cidaas.UserInfoField{ExtFieldKey: plan.userinfoFields.Website.ValueString()}
		}
		if !plan.userinfoFields.Gender.IsNull() {
			cp.UserinfoFields["gender"] = &cidaas.UserInfoField{ExtFieldKey: plan.userinfoFields.Gender.ValueString()}
		}
		if !plan.userinfoFields.Birthdate.IsNull() {
			cp.UserinfoFields["birthdate"] = &cidaas.UserInfoField{ExtFieldKey: plan.userinfoFields.Birthdate.ValueString()}
		}
		if !plan.userinfoFields.Zoneinfo.IsNull() {
			cp.UserinfoFields["zoneinfo"] = &cidaas.UserInfoField{ExtFieldKey: plan.userinfoFields.Zoneinfo.ValueString()}
		}
		if !plan.userinfoFields.Locale.IsNull() {
			cp.UserinfoFields["locale"] = &cidaas.UserInfoField{ExtFieldKey: plan.userinfoFields.Locale.ValueString()}
		}
		if !plan.userinfoFields.UpdatedAt.IsNull() {
			cp.UserinfoFields["updated_at"] = &cidaas.UserInfoField{ExtFieldKey: plan.userinfoFields.UpdatedAt.ValueString()}
		}
		if !plan.userinfoFields.Email.IsNull() {
			cp.UserinfoFields["email"] = &cidaas.UserInfoField{ExtFieldKey: plan.userinfoFields.Email.ValueString()}
		}
		if !plan.userinfoFields.EmailVerified.IsNull() {
			cp.UserinfoFields["email_verified"] = &cidaas.UserInfoField{ExtFieldKey: plan.userinfoFields.EmailVerified.ValueString()}
		}
		if !plan.userinfoFields.PhoneNumber.IsNull() {
			cp.UserinfoFields["phone_number"] = &cidaas.UserInfoField{ExtFieldKey: plan.userinfoFields.PhoneNumber.ValueString()}
		}
		if !plan.userinfoFields.MobileNumber.IsNull() {
			cp.UserinfoFields["mobile_number"] = &cidaas.UserInfoField{ExtFieldKey: plan.userinfoFields.MobileNumber.ValueString()}
		}
		if !plan.userinfoFields.Address.IsNull() {
			cp.UserinfoFields["address"] = &cidaas.UserInfoField{ExtFieldKey: plan.userinfoFields.Address.ValueString()}
		}
		if !plan.userinfoFields.Sub.IsNull() {
			cp.UserinfoFields["sub"] = &cidaas.UserInfoField{ExtFieldKey: plan.userinfoFields.Sub.ValueString()}
		}

		// Handle custom fields
		if !plan.userinfoFields.CustomFields.IsNull() {
			var cfMap map[string]string
			diag = plan.userinfoFields.CustomFields.ElementsAs(ctx, &cfMap, false)
			if diag.HasError() {
				return nil, diag
			}

			// Add custom fields with "customFields." prefix
			for key, value := range cfMap {
				cp.UserinfoFields["customFields."+key] = &cidaas.UserInfoField{ExtFieldKey: value}
			}
		}
	}

	return &cp, nil
}

func userInfoDefaultValue() basetypes.ObjectValue {
	return types.ObjectValueMust(
		map[string]attr.Type{
			"name":               types.StringType,
			"family_name":        types.StringType,
			"given_name":         types.StringType,
			"middle_name":        types.StringType,
			"nickname":           types.StringType,
			"preferred_username": types.StringType,
			"profile":            types.StringType,
			"picture":            types.StringType,
			"website":            types.StringType,
			"gender":             types.StringType,
			"birthdate":          types.StringType,
			"zoneinfo":           types.StringType,
			"locale":             types.StringType,
			"updated_at":         types.StringType,
			"email":              types.StringType,
			"email_verified":     types.StringType,
			"phone_number":       types.StringType,
			"mobile_number":      types.StringType,
			"address":            types.StringType,
			"sub":                types.StringType,
			"custom_fields":      types.MapType{ElemType: types.StringType},
		},
		map[string]attr.Value{
			"name":               types.StringValue("name"),
			"family_name":        types.StringValue("family_name"),
			"given_name":         types.StringValue("given_name"),
			"middle_name":        types.StringValue("middle_name"),
			"nickname":           types.StringValue("nickname"),
			"preferred_username": types.StringValue("preferred_username"),
			"profile":            types.StringValue("profile"),
			"picture":            types.StringValue("picture"),
			"website":            types.StringValue("website"),
			"gender":             types.StringValue("gender"),
			"birthdate":          types.StringValue("birthdate"),
			"zoneinfo":           types.StringValue("zoneinfo"),
			"locale":             types.StringValue("locale"),
			"updated_at":         types.StringValue("updated_at"),
			"email":              types.StringValue("email"),
			"email_verified":     types.StringValue("email_verified"),
			"phone_number":       types.StringValue("phone_number"),
			"mobile_number":      types.StringValue("mobile_number"),
			"address":            types.StringValue("address"),
			"sub":                types.StringValue("sub"),
			"custom_fields":      types.MapNull(types.StringType),
		})
}
