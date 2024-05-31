package resources

import (
	"context"
	"fmt"

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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ScopeResource struct {
	cidaasClient *cidaas.Client
}

type ScopeConfig struct {
	ID                    types.String `tfsdk:"id"`
	SecurityLevel         types.String `tfsdk:"security_level"`
	ScopeKey              types.String `tfsdk:"scope_key"`
	GroupName             types.Set    `tfsdk:"group_name"`
	RequiredUserConsent   types.Bool   `tfsdk:"required_user_consent"`
	LocalizedDescription  types.List   `tfsdk:"localized_descriptions"`
	localizedDescriptions []*LocalDescription
	ScopeOwner            types.String `tfsdk:"scope_owner"`
}

type LocalDescription struct {
	Locale      types.String `tfsdk:"locale"`
	Title       types.String `tfsdk:"title"`
	Description types.String `tfsdk:"description"`
}

func NewScopeResource() resource.Resource {
	return &ScopeResource{}
}

func (sc *ScopeConfig) extractLocalizedDescription(ctx context.Context) diag.Diagnostics {
	var diags diag.Diagnostics

	if !sc.LocalizedDescription.IsNull() {
		sc.localizedDescriptions = make([]*LocalDescription, 0, len(sc.LocalizedDescription.Elements()))
		diags = sc.LocalizedDescription.ElementsAs(ctx, &sc.localizedDescriptions, false)
	}
	return diags
}

func (r *ScopeResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_scope"
}

func (r *ScopeResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ScopeResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"security_level": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"PUBLIC", "CONFIDENTIAL"}...),
				},
				Default: stringdefault.StaticString("PUBLIC"),
			},
			"scope_key": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					// the destroy will throw the error too
					validators.UniqueIdentifier{},
				},
			},
			"group_name": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
				// group_name validator can be added by fetching the group_names using api
			},
			"required_user_consent": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
			},
			"scope_owner": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"localized_descriptions": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"locale": schema.StringAttribute{
							Optional: true,
							Computed: true,
							Default:  stringdefault.StaticString("en-US"),
							Validators: []validator.String{
								stringvalidator.OneOf(
									func() []string {
										var validLocals = make([]string, len(util.Locals)) //nolint:gofumpt
										for i, locale := range util.Locals {
											validLocals[i] = locale.LocaleString
										}
										return validLocals
									}()...),
							},
						},
						"title": schema.StringAttribute{
							Optional: true,
						},
						"description": schema.StringAttribute{
							Optional: true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(0, 256),
							},
						},
					},
				},
				Optional: true,
			},
		},
	}
}

func (r *ScopeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ScopeConfig
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(plan.extractLocalizedDescription(ctx)...)
	scopePayload, d := generateScopeModel(ctx, plan)
	resp.Diagnostics.Append(d...)
	if resp.Diagnostics.HasError() {
		return
	}
	response, err := r.cidaasClient.Scope.Upsert(*scopePayload)
	if err != nil {
		resp.Diagnostics.AddError("failed to create scope", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	plan.ScopeOwner = types.StringValue(response.Data.ScopeOwner)
	plan.ID = types.StringValue(response.Data.ID)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ScopeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ScopeConfig
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	res, err := r.cidaasClient.Scope.Get(state.ScopeKey.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to read scope", fmt.Sprintf("Error: %s", err.Error()))
		return
	}

	state.ID = types.StringValue(res.Data.ID)
	state.ScopeKey = types.StringValue(res.Data.ScopeKey)
	state.SecurityLevel = types.StringValue(res.Data.SecurityLevel)
	state.RequiredUserConsent = types.BoolValue(res.Data.RequiredUserConsent)
	state.ScopeOwner = types.StringValue(res.Data.ScopeOwner)

	var d diag.Diagnostics
	if len(res.Data.GroupName) > 0 {
		state.GroupName, d = types.SetValueFrom(ctx, types.StringType, res.Data.GroupName)
		resp.Diagnostics.Append(d...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	var objectValues []attr.Value
	localeDescription := types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"locale":      types.StringType,
			"title":       types.StringType,
			"description": types.StringType,
		},
	}

	for _, sc := range res.Data.LocaleWiseDescription {
		objValue := types.ObjectValueMust(localeDescription.AttrTypes, map[string]attr.Value{
			"locale":      types.StringValue(sc.Locale),
			"title":       types.StringValue(sc.Title),
			"description": types.StringValue(sc.Description),
		})
		objectValues = append(objectValues, objValue)
	}

	state.LocalizedDescription, d = types.ListValueFrom(ctx, localeDescription, objectValues)
	resp.Diagnostics.Append(d...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *ScopeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ScopeConfig
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(plan.extractLocalizedDescription(ctx)...)
	if resp.Diagnostics.HasError() {
		return
	}
	scopePayload, d := generateScopeModel(ctx, plan)
	resp.Diagnostics.Append(d...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.cidaasClient.Scope.Upsert(*scopePayload)
	if err != nil {
		resp.Diagnostics.AddError("failed to update scope", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ScopeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ScopeConfig
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	err := r.cidaasClient.Scope.Delete(state.ScopeKey.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to delete scope", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
}

func (r *ScopeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("scope_key"), req, resp)
}

func generateScopeModel(ctx context.Context, plan ScopeConfig) (*cidaas.ScopeModel, diag.Diagnostics) {
	scope := cidaas.ScopeModel{
		SecurityLevel:       plan.SecurityLevel.ValueString(),
		ScopeKey:            plan.ScopeKey.ValueString(),
		RequiredUserConsent: plan.RequiredUserConsent.ValueBool(),
	}
	diag := plan.GroupName.ElementsAs(ctx, &scope.GroupName, false)
	if diag.HasError() {
		return nil, diag
	}

	for _, ld := range plan.localizedDescriptions {
		scope.LocaleWiseDescription = append(scope.LocaleWiseDescription, cidaas.ScopeLocalDescription{
			Locale:      ld.Locale.ValueString(),
			Language:    getLanguageForLocale(ld.Locale.ValueString()),
			Title:       ld.Title.ValueString(),
			Description: ld.Description.ValueString(),
		})
	}
	return &scope, nil
}

func getLanguageForLocale(locale string) string {
	for _, v := range util.Locals {
		if v.LocaleString == locale {
			return v.Language
		}
	}
	return "en"
}
