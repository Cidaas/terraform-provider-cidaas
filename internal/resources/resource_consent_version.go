package resources

import (
	"context"
	"fmt"
	"net/http"
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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const (
	SCOPES = "SCOPES"
	URL    = "URL"
)

type ConsentVersionResource struct {
	BaseResource
}

func NewConsentVersionResource() resource.Resource {
	return &ConsentVersionResource{
		BaseResource: NewBaseResource(
			BaseResourceConfig{
				Name:   RESOURCE_CONSENT_VERSION,
				Schema: &consentversionSchema,
			},
		),
	}
}

type ConsentVersionConfig struct {
	ID             types.String  `tfsdk:"id"`
	Version        types.Float64 `tfsdk:"version"`
	ConsentID      types.String  `tfsdk:"consent_id"`
	ConsentType    types.String  `tfsdk:"consent_type"`
	Scopes         types.Set     `tfsdk:"scopes"`
	RequiredFields types.Set     `tfsdk:"required_fields"`
	ConsentLocales types.Set     `tfsdk:"consent_locales"`

	consentLocale []*ConsentLocale
}

type ConsentLocale struct {
	Content types.String `tfsdk:"content"`
	Locale  types.String `tfsdk:"locale"`
	URL     types.String `tfsdk:"url"`
}

func (r *ConsentVersionResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, res *resource.ValidateConfigResponse) {
	var config ConsentVersionConfig
	res.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	res.Diagnostics.Append(config.extract(ctx)...)

	localeMap := make(map[string]bool)

	if !config.ConsentLocales.IsNull() && !config.ConsentLocales.IsUnknown() && len(config.consentLocale) > 0 {
		for _, loc := range config.consentLocale {
			locale := loc.Locale.ValueString()
			if _, exists := localeMap[locale]; exists {
				res.Diagnostics.AddError("Duplicate locale not allowed", fmt.Sprintf("Duplicate locale '%s' found in consent_locales", locale))
			}
			if config.ConsentType.ValueString() == SCOPES && !loc.URL.IsNull() && !loc.URL.IsUnknown() {
				res.Diagnostics.AddError("Unsupported attribute", "attribute 'consent_locales.url' not supported when consent_type is 'SCOPES'")
			}
			if config.ConsentType.ValueString() == URL && (loc.URL.IsNull() || loc.URL.ValueString() == "") {
				res.Diagnostics.AddError("Missing required attribute", "attribute 'consent_locales.url' is required or can't be empty when consent_type is 'URL'")
			}
			localeMap[locale] = true
		}
	}
	if config.ConsentType.ValueString() == SCOPES && config.Scopes.IsNull() {
		res.Diagnostics.AddError("Missing required attribute", "attribute 'scopes' is required when consent_type is 'SCOPES'")
	}
	if config.ConsentType.ValueString() == SCOPES && config.RequiredFields.IsNull() {
		res.Diagnostics.AddError("Missing required attribute", "attribute 'required_fields' is required when consent_type is 'SCOPES'")
	}
	if config.ConsentType.ValueString() == URL && !config.Scopes.IsNull() && !config.Scopes.IsUnknown() {
		res.Diagnostics.AddError("Unsupported attribute", "attribute 'scopes' not supported when consent_type is 'URL'")
	}
	if config.ConsentType.ValueString() == URL && !config.RequiredFields.IsNull() && !config.RequiredFields.IsUnknown() {
		res.Diagnostics.AddError("Unsupported attribute", "attribute 'required_fields' not supported when consent_type is 'URL'")
	}
}

var consentversionSchema = schema.Schema{
	MarkdownDescription: "The Consent Version resource in the provider allows you to manage different versions of a specific consent in Cidaas." +
		"\n This resource also supports managing consent versions across multiple locales enabling different configurations such as URLs and content for each locale." +
		"\n\n Ensure that the below scopes are assigned to the client with the specified `client_id`:" +
		"\n- cidaas:tenant_consent_read" +
		"\n- cidaas:tenant_consent_write" +
		"\n- cidaas:tenant_consent_delete",
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The unique identifier of the consent version.",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"consent_type": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: "Specifies the type of consent. The allowed values are `SCOPES` or `URL`. It can not be updated for a specific consent version.",
			Validators: []validator.String{
				stringvalidator.OneOf(SCOPES, URL),
			},
			PlanModifiers: []planmodifier.String{
				validators.UniqueIdentifier{},
			},
		},
		"consent_id": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The `consent_id` for which the consent version is created. It can not be updated for a specific consent version.",
			PlanModifiers: []planmodifier.String{
				validators.UniqueIdentifier{},
			},
		},
		"version": schema.Float64Attribute{
			Required:            true,
			MarkdownDescription: "The version number of the consent. It can not be updated for a specific consent version.",
			PlanModifiers: []planmodifier.Float64{
				validators.ImmutableInt64Identifier{},
			},
		},
		"scopes": schema.SetAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: "A set of scopes related to the consent. It can not be updated for a specific consent version." +
				"\nNote that the attribute `scopes` is required only if the `consent_type` is set to **SCOPES**.",
			PlanModifiers: []planmodifier.Set{
				validators.ImmutableSetIdentifier{},
			},
		},
		"required_fields": schema.SetAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: "A set of fields that are required for the consent. It can not be updated for a specific consent version." +
				"\nNote that the attribute `required_fields` is required only if the `consent_type` is set to **SCOPES**.",
			PlanModifiers: []planmodifier.Set{
				validators.ImmutableSetIdentifier{},
			},
		},
		"consent_locales": schema.SetNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"content": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "The content of the consent version associated with a specific locale.",
					},
					"locale": schema.StringAttribute{
						Required:            true,
						MarkdownDescription: "The locale for which the consent version is created. e.g. `en-us`, `de`.",
						Validators: []validator.String{
							stringvalidator.OneOf(
								func() []string {
									validLocals := make([]string, len(util.Locals)) //nolint:gofumpt
									for i, locale := range util.Locals {
										validLocals[i] = strings.ToLower(locale.LocaleString)
									}
									return validLocals
								}()...),
						},
					},
					"url": schema.StringAttribute{
						Optional: true,
						MarkdownDescription: "The url to the consent page of the created consent version." +
							"\nNote that the attribute `url` is required only if the `consent_type` is set to **URL**.",
					},
				},
			},
			Required: true,
		},
	},
}

func (cv *ConsentVersionConfig) extract(ctx context.Context) diag.Diagnostics {
	var diags diag.Diagnostics
	if !cv.ConsentLocales.IsNull() && !cv.ConsentLocales.IsUnknown() {
		cv.consentLocale = make([]*ConsentLocale, 0, len(cv.ConsentLocales.Elements()))
		diags = cv.ConsentLocales.ElementsAs(ctx, &cv.consentLocale, false)
	}
	return diags
}

func (r *ConsentVersionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ConsentVersionConfig
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(plan.extract(ctx)...)
	consent := cidaas.ConsentVersionModel{
		Version:     plan.Version.ValueFloat64(),
		ConsentID:   plan.ConsentID.ValueString(),
		ConsentType: plan.ConsentType.ValueString(),
	}

	if consent.ConsentType == SCOPES {
		if len(plan.Scopes.Elements()) > 0 {
			resp.Diagnostics.Append(plan.Scopes.ElementsAs(ctx, &consent.Scopes, false)...)
		}
		if len(plan.RequiredFields.Elements()) > 0 {
			resp.Diagnostics.Append(plan.RequiredFields.ElementsAs(ctx, &consent.RequiredFields, false)...)
		}
	}

	var restLocals []*ConsentLocale
	if !plan.ConsentLocales.IsNull() && !plan.ConsentLocales.IsUnknown() && len(plan.consentLocale) > 0 {
		// consent version is created with the first element from the consent_local slice,
		// as the API requires the consent_local field to be present in the request payload.
		// later rest consent locals are processed and locals are created separately with the consent local api
		consent.ConsentLocale.Locale = plan.consentLocale[0].Locale.ValueString()
		if !plan.consentLocale[0].Content.IsNull() {
			consent.ConsentLocale.Content = plan.consentLocale[0].Content.ValueString()
		}
		if consent.ConsentType == URL && !plan.consentLocale[0].URL.IsNull() {
			consent.ConsentLocale.URL = plan.consentLocale[0].URL.ValueString()
		}
		restLocals = plan.consentLocale[1:]
	}

	res, err := r.cidaasClient.ConsentVersion.Upsert(consent)
	if err != nil {
		resp.Diagnostics.AddError("failed to create consent version", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	plan.ID = util.StringValueOrNull(&res.Data.ID)
	for _, pcl := range restLocals {
		consentLocal := cidaas.ConsentLocalModel{
			ConsentID:        plan.ConsentID.ValueString(),
			ConsentVersionID: plan.ID.ValueString(),
			Content:          pcl.Content.ValueString(),
			Locale:           pcl.Locale.ValueString(),
		}
		if plan.ConsentType.ValueString() == URL {
			consentLocal.URL = pcl.URL.ValueString()
		}
		_, err := r.cidaasClient.ConsentVersion.UpsertLocal(consentLocal)
		if err != nil {
			resp.Diagnostics.AddError("failed to create consent locale", fmt.Sprintf("Error: %s", err.Error()))
			return
		}
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ConsentVersionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ConsentVersionConfig
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(state.extract(ctx)...)
	res, err := r.cidaasClient.ConsentVersion.Get(state.ConsentID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Failed to read consent version", fmt.Sprintf("Error: %s ", err.Error()))
		return
	}
	if res.Success && res.Status == http.StatusOK && len(res.Data) == 0 {
		resp.Diagnostics.AddError("Invalid consent_id", fmt.Sprintf("No consent version found for the provided consent_id %+v", state.ConsentID.String()))
		return
	}
	isAvailable := false
	if len(res.Data) > 0 {
		for _, version := range res.Data {
			if version.ID == state.ID.ValueString() {
				isAvailable = true
				consentType := version.ConsentType
				state.Version = types.Float64Value(version.Version)
				state.ConsentType = util.StringValueOrNull(&consentType)
			}
		}
	}
	var diag diag.Diagnostics
	var objectValues []attr.Value
	consentLocalObjectType := types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"content": types.StringType,
			"locale":  types.StringType,
			"url":     types.StringType,
		},
	}

	for _, cl := range state.consentLocale {
		res, err := r.cidaasClient.ConsentVersion.GetLocal(state.ID.ValueString(), cl.Locale.ValueString())
		if err != nil {
			resp.Diagnostics.AddError("Failed to read consent version locale", fmt.Sprintf("Error: %s ", err.Error()))
			return
		}
		if !res.Success && res.Status == http.StatusNoContent {
			resp.Diagnostics.AddError(
				"Consent Version Local not found",
				fmt.Sprintf("No consent version locale found for the combination of consent_version_id %s and locale %s.", state.ID.String(), cl.Locale.String()),
			)
			return
		}
		state.ConsentID = util.StringValueOrNull(&res.Data.ConsentID)
		if state.ConsentType.ValueString() == SCOPES {
			state.Scopes = util.SetValueOrNull(res.Data.Scopes)
			state.RequiredFields = util.SetValueOrNull(res.Data.RequiredFields)
		}

		objValue := types.ObjectValueMust(
			consentLocalObjectType.AttrTypes,
			map[string]attr.Value{
				"content": util.StringValueOrNull(&res.Data.Content),
				"locale":  util.StringValueOrNull(&res.Data.Locale),
				"url":     util.StringValueOrNull(&res.Data.URL),
			})
		objectValues = append(objectValues, objValue)

	}
	state.ConsentLocales, diag = types.SetValueFrom(ctx, consentLocalObjectType, objectValues)
	resp.Diagnostics.Append(diag...)
	if resp.Diagnostics.HasError() {
		return
	}
	if !isAvailable {
		resp.Diagnostics.AddError(
			"Consent Version not found",
			fmt.Sprintf("Consent Version with ID %s not found for the provided consent_id %s", state.ID.String(), state.ConsentID.String()),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *ConsentVersionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ConsentVersionConfig

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(plan.extract(ctx)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	for _, pcl := range plan.consentLocale {
		consentLocal := cidaas.ConsentLocalModel{
			ConsentID:        state.ConsentID.ValueString(),
			ConsentVersionID: state.ID.ValueString(),
			Content:          pcl.Content.ValueString(),
			Locale:           pcl.Locale.ValueString(),
		}
		if plan.ConsentType.ValueString() == URL {
			consentLocal.URL = pcl.URL.ValueString()
		}
		_, err := r.cidaasClient.ConsentVersion.UpsertLocal(consentLocal)
		if err != nil {
			resp.Diagnostics.AddError("Failed to update consent locale", fmt.Sprintf("Error: %s", err.Error()))
			return
		}
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ConsentVersionResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
}

func (r *ConsentVersionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	id := req.ID
	parts := strings.Split(id, ":")
	if len(parts) < 3 {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: 'consent_id:id:locale', got: %s", id),
		)
		return
	}
	locals := parts[2:]
	var objectValues []attr.Value
	consentLocalObjectType := types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"content": types.StringType,
			"locale":  types.StringType,
			"url":     types.StringType,
		},
	}

	for _, locale := range locals {
		objValue := types.ObjectValueMust(
			consentLocalObjectType.AttrTypes,
			map[string]attr.Value{
				"content": types.StringNull(),
				"locale":  types.StringValue(locale),
				"url":     types.StringNull(),
			})
		objectValues = append(objectValues, objValue)
	}
	resp.State.SetAttribute(ctx, path.Root("consent_id"), parts[0])
	resp.State.SetAttribute(ctx, path.Root("id"), parts[1])
	resp.State.SetAttribute(ctx, path.Root("consent_locales"), types.SetValueMust(consentLocalObjectType, objectValues))
}
