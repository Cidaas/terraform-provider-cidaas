package resources

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/cidaas"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type CustomProvider struct {
	cidaasClient *cidaas.Client
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
	Scopes         types.Set    `tfsdk:"scopes"`
	UserinfoFields types.Object `tfsdk:"userinfo_fields"`
	scopes         []*Scope
	userinfoFields *UserInfoField
}

type Scope struct {
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

func NewCustomProvider() resource.Resource {
	return &CustomProvider{}
}

func (pc *ProviderConfig) extract(ctx context.Context) diag.Diagnostics {
	var diags diag.Diagnostics
	if !pc.UserinfoFields.IsNull() {
		pc.userinfoFields = &UserInfoField{}
		diags = pc.UserinfoFields.As(ctx, pc.userinfoFields, basetypes.ObjectAsOptions{})
	}
	if !pc.Scopes.IsNull() {
		pc.scopes = make([]*Scope, 0, len(pc.Scopes.Elements()))
		diags = pc.Scopes.ElementsAs(ctx, &pc.scopes, false)
	}
	return diags
}

func (r *CustomProvider) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_custom_provider"
}

func (r *CustomProvider) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *CustomProvider) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"provider_name": schema.StringAttribute{
				Required: true,
			},
			"display_name": schema.StringAttribute{
				Required: true,
			},
			"logo_url": schema.StringAttribute{
				Optional: true,
			},
			"standard_type": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"OPENID_CONNECT", "OAUTH2"}...),
				},
			},
			"client_id": schema.StringAttribute{
				Required: true,
			},
			"client_secret": schema.StringAttribute{
				Required:  true,
				Sensitive: true,
				// PlanModifiers: []planmodifier.String{
				// 	stringplanmodifier.UseStateForUnknown(),
				// },
			},
			"authorization_endpoint": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^https://.+$`),
						"must be a valid URL starting with https://",
					),
				},
			},
			"token_endpoint": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^https://.+$`),
						"must be a valid URL starting with https://",
					),
				},
			},
			"userinfo_endpoint": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^https://.+$`),
						"must be a valid URL starting with https://",
					),
				},
			},
			"scopes": schema.SetNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"scope_name": schema.StringAttribute{
							Optional: true,
						},
						"required": schema.BoolAttribute{
							Optional: true,
							// Default:  booldefault.StaticBool(true),
						},
						"recommended": schema.BoolAttribute{
							Optional: true,
						},
					},
				},
				Optional: true,
			},
			"scope_display_label": schema.StringAttribute{
				Required: true,
			},
			"userinfo_fields": schema.SingleNestedAttribute{
				Description: "A binding associates a set of principals to a role.",
				Optional:    true,
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
						Required:    true,
					},
				},
			},
			"domains": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
		},
	}
}

func (r *CustomProvider) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ProviderConfig
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(plan.extract(ctx)...)
	customProvider := prepareCpRequestPayload(plan)
	resp.Diagnostics.Append(plan.Domains.ElementsAs(ctx, &customProvider.Domains, false)...)
	// customFields can be unmarshalled the below way but we need to add cusomFields prefix to all the key so done separately
	// resp.Diagnostics.Append(plan.UserinfoFields.CustomFields.ElementsAs(ctx, &customProvider.UserinfoFields.CustomFields, false)...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.cidaasClient.CustomProvider.CreateCustomProvider(customProvider)
	if err != nil {
		// move fmt.Sprintf("Error: %s", err.Error()) in all to a util function
		resp.Diagnostics.AddError("failed to create custom provider", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	plan.ID = types.StringValue(res.Data.ProviderName)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *CustomProvider) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ProviderConfig
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	res, err := r.cidaasClient.CustomProvider.GetCustomProvider(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to read custom provider", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	state.StandardType = types.StringValue(res.Data.StandardType)
	state.AuthorizationEndpoint = types.StringValue(res.Data.AuthorizationEndpoint)
	state.TokenEndpoint = types.StringValue(res.Data.TokenEndpoint)
	state.ProviderName = types.StringValue(res.Data.ProviderName)
	state.DisplayName = types.StringValue(res.Data.DisplayName)
	state.LogoURL = types.StringValue(res.Data.LogoURL)
	state.UserinfoEndpoint = types.StringValue(res.Data.UserinfoEndpoint)
	state.ID = types.StringValue(res.Data.ProviderName)
	state.ScopeDisplayLabel = types.StringValue(res.Data.Scopes.DisplayLabel)
	state.ClientID = types.StringValue(res.Data.ClientID)
	state.ClientSecret = types.StringValue(res.Data.ClientSecret)

	domainsSetValue, d := types.SetValueFrom(ctx, types.StringType, res.Data.Domains)
	resp.Diagnostics.Append(d...)
	state.Domains, d = domainsSetValue.ToSetValue(ctx)
	resp.Diagnostics.Append(d...)
	if resp.Diagnostics.HasError() {
		return
	}

	// alternate way to assign the domains without diag
	// state.Domains = types.SetValueMust(state.Domains.ElementType(ctx), []attr.Value{
	//  run a for loop here
	// 	types.StringValue(res.Data.Domains[0]),
	// 	types.StringValue(res.Data.Domains[1]),
	// })

	var objectValues []attr.Value
	scopeObjectType := types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"scope_name":  types.StringType,
			"required":    types.BoolType,
			"recommended": types.BoolType,
		},
	}

	for _, sc := range res.Data.Scopes.Scopes {
		objValue := types.ObjectValueMust(scopeObjectType.AttrTypes, map[string]attr.Value{
			"scope_name":  types.StringValue(sc.ScopeName),
			"required":    types.BoolValue(sc.Required),
			"recommended": types.BoolValue(sc.Recommended),
		})
		objectValues = append(objectValues, objValue)
	}

	state.Scopes, d = types.SetValueFrom(ctx, scopeObjectType, objectValues)
	resp.Diagnostics.Append(d...)
	if resp.Diagnostics.HasError() {
		return
	}

	metadataAttributeTypes := map[string]attr.Type{}
	metadataAttributes := map[string]attr.Value{}
	customFields := map[string]attr.Value{}

	for key, value := range res.Data.UserinfoFields {
		if strings.HasPrefix(key, "customFields.") {
			customFields[strings.TrimPrefix(key, "customFields.")] = types.StringValue(value)
		} else {
			metadataAttributeTypes[key] = types.StringType
			metadataAttributes[key] = types.StringValue(value)
		}
	}
	metadataAttributeTypes["custom_fields"] = types.MapType{ElemType: types.StringType}
	metadataAttributes["custom_fields"] = types.MapValueMust(types.StringType, customFields)
	state.UserinfoFields, d = types.ObjectValue(metadataAttributeTypes, metadataAttributes)
	resp.Diagnostics.Append(d...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *CustomProvider) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ProviderConfig
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
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
}

func (r *CustomProvider) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func prepareCpRequestPayload(pc ProviderConfig) *cidaas.CustomProviderModel {
	var cp cidaas.CustomProviderModel

	cp.StandardType = pc.StandardType.ValueString()
	cp.AuthorizationEndpoint = pc.AuthorizationEndpoint.ValueString()
	cp.TokenEndpoint = pc.TokenEndpoint.ValueString()
	cp.ProviderName = pc.ProviderName.ValueString()
	cp.DisplayName = pc.DisplayName.ValueString()
	cp.LogoURL = pc.LogoURL.ValueString()
	cp.UserinfoEndpoint = pc.UserinfoEndpoint.ValueString()
	cp.ClientID = pc.ClientID.ValueString()
	cp.ClientSecret = pc.ClientSecret.ValueString()

	if !pc.ID.IsNull() {
		cp.ID = pc.ID.ValueString()
	}

	var childScopes []cidaas.ScopeChild
	for _, v := range pc.scopes {
		childScopes = append(childScopes, cidaas.ScopeChild{
			ScopeName:   v.ScopeName.ValueString(),
			Required:    v.Required.ValueBool(),
			Recommended: v.Recommended.ValueBool(),
		})
	}
	cp.Scopes.Scopes = childScopes
	cp.Scopes.DisplayLabel = pc.ScopeDisplayLabel.ValueString()

	userInfoFields := map[string]string{}

	userInfoFields["name"] = pc.userinfoFields.Name.ValueString()
	userInfoFields["family_name"] = pc.userinfoFields.FamilyName.ValueString()
	userInfoFields["given_name"] = pc.userinfoFields.GivenName.ValueString()
	userInfoFields["middle_name"] = pc.userinfoFields.MiddleName.ValueString()
	userInfoFields["nickname"] = pc.userinfoFields.Nickname.ValueString()
	userInfoFields["preferred_username"] = pc.userinfoFields.PreferredUsername.ValueString()
	userInfoFields["profile"] = pc.userinfoFields.Profile.ValueString()
	userInfoFields["picture"] = pc.userinfoFields.Picture.ValueString()
	userInfoFields["website"] = pc.userinfoFields.Website.ValueString()
	userInfoFields["gender"] = pc.userinfoFields.Gender.ValueString()
	userInfoFields["birthdate"] = pc.userinfoFields.Birthdate.ValueString()
	userInfoFields["zoneinfo"] = pc.userinfoFields.Zoneinfo.ValueString()
	userInfoFields["locale"] = pc.userinfoFields.Locale.ValueString()
	userInfoFields["updated_at"] = pc.userinfoFields.UpdatedAt.ValueString()
	userInfoFields["email"] = pc.userinfoFields.Email.ValueString()
	userInfoFields["email_verified"] = pc.userinfoFields.EmailVerified.ValueString()
	userInfoFields["phone_number"] = pc.userinfoFields.PhoneNumber.ValueString()
	userInfoFields["mobile_number"] = pc.userinfoFields.MobileNumber.ValueString()
	userInfoFields["address"] = pc.userinfoFields.Address.ValueString()
	userInfoFields["sub"] = pc.userinfoFields.Sub.ValueString()

	for k, v := range pc.userinfoFields.CustomFields.Elements() {
		userInfoFields["customFields."+k] = v.String()
	}

	cp.UserinfoFields = userInfoFields
	return &cp
}
