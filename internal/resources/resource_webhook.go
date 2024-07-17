package resources

import (
	"context"
	"fmt"
	"regexp"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/cidaas"
	"github.com/Cidaas/terraform-provider-cidaas/helpers/util"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type WebhookResource struct {
	cidaasClient *cidaas.Client
}

type WebhookConfig struct {
	ID               types.String `tfsdk:"id"`
	AuthType         types.String `tfsdk:"auth_type"`
	URL              types.String `tfsdk:"url"`
	Events           types.Set    `tfsdk:"events"`
	APIKeyConfig     types.Object `tfsdk:"apikey_config"`
	TOTPConfig       types.Object `tfsdk:"totp_config"`
	CidaasAuthConfig types.Object `tfsdk:"cidaas_auth_config"`
	Disable          types.Bool   `tfsdk:"disable"`
	CreatedAt        types.String `tfsdk:"created_at"`
	UpdatedAt        types.String `tfsdk:"updated_at"`
	apiKeyConfig     *AuthConfig
	totpConfig       *AuthConfig
	cidaasAuthConfig *CidaasAuthConfig
}

type AuthConfig struct {
	Placeholder types.String `tfsdk:"placeholder"`
	Placement   types.String `tfsdk:"placement"`
	Key         types.String `tfsdk:"key"`
}

type CidaasAuthConfig struct {
	ClientID types.String `tfsdk:"client_id"`
}

func NewWebhookResource() resource.Resource {
	return &WebhookResource{}
}

func (w *WebhookConfig) extractAuthConfigs(ctx context.Context) diag.Diagnostics {
	var diags diag.Diagnostics
	if !w.APIKeyConfig.IsNull() {
		w.apiKeyConfig = &AuthConfig{}
		diags = w.APIKeyConfig.As(ctx, w.apiKeyConfig, basetypes.ObjectAsOptions{})
	}
	if !w.TOTPConfig.IsNull() {
		w.totpConfig = &AuthConfig{}
		diags = w.TOTPConfig.As(ctx, w.totpConfig, basetypes.ObjectAsOptions{})
	}
	if !w.CidaasAuthConfig.IsNull() {
		w.cidaasAuthConfig = &CidaasAuthConfig{}
		diags = w.CidaasAuthConfig.As(ctx, w.cidaasAuthConfig, basetypes.ObjectAsOptions{})
	}
	return diags
}

func (r *WebhookResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_webhook"
}

func (r *WebhookResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *WebhookResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The Webhook resource in the provider facilitates integration of webhooks in the Cidaas system." +
			" This resource allows you to configure webhooks with different authentication options." +
			"\n\n Ensure that the below scopes are assigned to the client with the specified `client_id`:" +
			"\n- cidaas:webhook_read" +
			"\n- cidaas:webhook_write" +
			"\n- cidaas:webhook_delete",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier of the webhook resource.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"auth_type": schema.StringAttribute{
				Required: true,
				Description: "The attribute auth_type is to define how this url is secured from your end." +
					"The allowed values are `APIKEY`, `TOTP` and `CIDAAS_OAUTH2`",
				Validators: []validator.String{
					stringvalidator.OneOf(cidaas.AllowedAuthType...),
				},
				PlanModifiers: []planmodifier.String{
					&configVerifier{},
				},
			},
			"url": schema.StringAttribute{
				Required:    true,
				Description: "The webhook url that needs to be called when an event occurs.",
			},
			"events": schema.SetAttribute{
				ElementType: types.StringType,
				Required:    true,
				Description: "A set of events that trigger the webhook.",
				Validators: []validator.Set{
					setvalidator.SizeAtLeast(1),
					setvalidator.ValueStringsAre(
						stringvalidator.OneOf(cidaas.AllowedEvents...),
					),
				},
			},
			"apikey_config": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "Configuration for API key-based authentication. It's a **required** parameter when the auth_type is APIKEY.",
				Attributes: map[string]schema.Attribute{
					"placeholder": schema.StringAttribute{
						Required:    true,
						Description: "The attribute is the placeholder for the key which need to be passed as a query parameter or in the request header.",
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[a-z]+$`),
								"must contain only lowercase alphabets",
							),
						},
					},
					"placement": schema.StringAttribute{
						Required: true,
						Description: "The placement of the API key in the request (e.g., query)." +
							"The allowed value are `header` and `query`.",
						Validators: []validator.String{
							stringvalidator.OneOf(cidaas.AllowedKeyPlacementValue...),
						},
					},
					"key": schema.StringAttribute{
						Required: true,
						Description: "The API key that will be used to authenticate the webhook request." +
							"The key that will be passed in the request header or in query param as configured in the attribute `placement`",
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
						},
					},
				},
			},
			"totp_config": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "Configuration for TOTP based authentication.  It's a **required** parameter when the auth_type is TOTP.",
				Attributes: map[string]schema.Attribute{
					"placeholder": schema.StringAttribute{
						Required:    true,
						Description: "A placeholder value for the TOTP.",
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[a-z]+$`),
								"must contain only lowercase alphabets",
							),
						},
					},
					"placement": schema.StringAttribute{
						Required: true,
						Description: "The placement of the TOTP in the request." +
							"The allowed value are `header` and `query`.",
						Validators: []validator.String{
							stringvalidator.OneOf(cidaas.AllowedKeyPlacementValue...),
						},
					},
					"key": schema.StringAttribute{
						Required:    true,
						Description: "The key used for TOTP authentication.",
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
						},
					},
				},
			},
			"cidaas_auth_config": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "Configuration for Cidaas authentication. It's a **required** parameter when the auth_type is CIDAAS_OAUTH2.",
				Attributes: map[string]schema.Attribute{
					"client_id": schema.StringAttribute{
						Required:    true,
						Description: "The client ID for Cidaas authentication.",
					},
				},
			},
			"disable": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Flag to disable the webhook.",
				Default:     booldefault.StaticBool(false),
			},
			"created_at": schema.StringAttribute{
				Computed:    true,
				Description: "The timestamp when the webhook was created.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"updated_at": schema.StringAttribute{
				Computed:    true,
				Description: "The timestamp when the webhook was last updated.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *WebhookResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) { //nolint:dupl
	var plan WebhookConfig
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(plan.extractAuthConfigs(ctx)...)
	wbModel, diags := prepareWebhookModel(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.cidaasClient.Webhook.Upsert(*wbModel)
	if err != nil {
		resp.Diagnostics.AddError("failed to create group type", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	plan.ID = util.StringValueOrNull(&res.Data.ID)
	plan.CreatedAt = util.StringValueOrNull(&res.Data.CreatedTime)
	plan.UpdatedAt = util.StringValueOrNull(&res.Data.UpdatedTime)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *WebhookResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state WebhookConfig
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	res, err := r.cidaasClient.Webhook.Get(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to read webhook", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	state.ID = util.StringValueOrNull(&res.Data.ID)
	state.AuthType = util.StringValueOrNull(&res.Data.AuthType)
	state.URL = util.StringValueOrNull(&res.Data.URL)
	state.Disable = util.BoolValueOrNull(&res.Data.Disable)
	state.ID = util.StringValueOrNull(&res.Data.ID)
	state.CreatedAt = util.StringValueOrNull(&res.Data.CreatedTime)
	state.UpdatedAt = util.StringValueOrNull(&res.Data.UpdatedTime)
	state.Events = util.SetValueOrNull(res.Data.Events)

	authConfig := types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"placeholder": types.StringType,
			"placement":   types.StringType,
			"key":         types.StringType,
		},
	}

	if res.Data.APIKeyDetails.Apikey != "" {
		apiKeyConfig, diags := types.ObjectValue(authConfig.AttrTypes, map[string]attr.Value{
			"placeholder": util.StringValueOrNull(&res.Data.APIKeyDetails.ApikeyPlaceholder),
			"placement":   util.StringValueOrNull(&res.Data.APIKeyDetails.ApikeyPlacement),
			"key":         util.StringValueOrNull(&res.Data.APIKeyDetails.Apikey),
		})
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		state.APIKeyConfig = apiKeyConfig
	}
	if res.Data.TotpDetails.TotpKey != "" {
		totpConfig, diags := types.ObjectValue(authConfig.AttrTypes, map[string]attr.Value{
			"placeholder": util.StringValueOrNull(&res.Data.TotpDetails.TotpPlaceholder),
			"placement":   util.StringValueOrNull(&res.Data.TotpDetails.TotpPlacement),
			"key":         util.StringValueOrNull(&res.Data.TotpDetails.TotpKey),
		})
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		state.TOTPConfig = totpConfig
	}

	if res.Data.CidaasAuthDetails.ClientID != "" {
		oauthConfig, diags := types.ObjectValue(map[string]attr.Type{
			"client_id": types.StringType,
		}, map[string]attr.Value{
			"client_id": util.StringValueOrNull(&res.Data.CidaasAuthDetails.ClientID),
		})
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		state.CidaasAuthConfig = oauthConfig
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *WebhookResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state WebhookConfig
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(plan.extractAuthConfigs(ctx)...)
	wbModel, diags := prepareWebhookModel(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.cidaasClient.Webhook.Upsert(*wbModel)
	if err != nil {
		resp.Diagnostics.AddError("failed to update webhook", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *WebhookResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state WebhookConfig
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	err := r.cidaasClient.Webhook.Delete(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to delete webhook", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
}

func (r *WebhookResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

var _ planmodifier.String = configVerifier{}

type configVerifier struct{}

func (v configVerifier) Description(_ context.Context) string {
	return "Verifies the availability of config details for the provided auth_type."
}

func (v configVerifier) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v configVerifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.ConfigValue.IsUnknown() ||
		req.ConfigValue.IsNull() ||
		!util.StringInSlice(req.ConfigValue.ValueString(), cidaas.AllowedAuthType) {
		return
	}

	var tempConfig types.Object
	configAttr := "apikey_config"
	authType := "APIKEY"

	if req.ConfigValue.ValueString() == "TOTP" {
		configAttr = "totp_config"
		authType = "TOTP"
	}
	if req.ConfigValue.ValueString() == "CIDAAS_OAUTH2" {
		configAttr = "cidaas_auth_config"
		authType = "CIDAAS_OAUTH2"
	}
	diags := req.Config.GetAttribute(ctx, path.Root(configAttr), &tempConfig)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	if tempConfig.IsNull() {
		resp.Diagnostics.AddError("Unexpected Resource Configuration",
			fmt.Sprintf("The attribute %s cannot be empty when the auth_type is %s", configAttr, authType))
	}
}

func prepareWebhookModel(ctx context.Context, plan WebhookConfig) (*cidaas.WebhookModel, diag.Diagnostics) {
	wb := cidaas.WebhookModel{
		ID:       plan.ID.ValueString(),
		AuthType: plan.AuthType.ValueString(),
		URL:      plan.URL.ValueString(),
		Disable:  plan.Disable.ValueBool(),
	}
	if !plan.APIKeyConfig.IsNull() {
		wb.APIKeyDetails = cidaas.APIKeyDetails{
			ApikeyPlaceholder: plan.apiKeyConfig.Placeholder.ValueString(),
			ApikeyPlacement:   plan.apiKeyConfig.Placement.ValueString(),
			Apikey:            plan.apiKeyConfig.Key.ValueString(),
		}
	}
	if !plan.TOTPConfig.IsNull() {
		wb.TotpDetails = cidaas.TotpDetails{
			TotpPlaceholder: plan.totpConfig.Placeholder.ValueString(),
			TotpPlacement:   plan.totpConfig.Placement.ValueString(),
			TotpKey:         plan.totpConfig.Key.ValueString(),
		}
	}
	if !plan.CidaasAuthConfig.IsNull() {
		wb.CidaasAuthDetails = cidaas.AuthDetails{
			ClientID: plan.cidaasAuthConfig.ClientID.ValueString(),
		}
	}
	diags := plan.Events.ElementsAs(ctx, &wb.Events, false)
	if diags.HasError() {
		return nil, diags
	}
	return &wb, nil
}
