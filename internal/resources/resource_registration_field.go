package resources

import (
	"context"
	"fmt"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/cidaas"
	"github.com/Cidaas/terraform-provider-cidaas/helpers/util"
	"github.com/Cidaas/terraform-provider-cidaas/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var allowedDataTypes = []string{"TEXT", "NUMBER", "SELECT", "MULTISELECT", "RADIO", "CHECKBOX", "PASSWORD", "DATE", "URL", "EMAIL",
	"TEXTAREA", "MOBILE", "CONSENT", "JSON_STRING", "USERNAME", "ARRAY", "GROUPING", "DAYDATE"}

type RegFieldResource struct {
	cidaasClient *cidaas.Client
}

type RegFieldConfig struct{}

func NewRegFieldResource() resource.Resource {
	return &RegFieldResource{}
}

func (r *RegFieldResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_registration_page_field"
}

func (r *RegFieldResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *RegFieldResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"required": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
			},
			"internal": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
			},
			"claimable": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(true),
			},
			"is_searchable": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(true),
			},
			"enabled": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
			},
			"unique": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
			},
			"overwrite_with_null_from_social_provider": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(true),
			},
			"read_only": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
			},
			"is_group": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
			},
			"is_list": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
			},
			// optional: Order of the Field in the UI
			"order": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				Default:  int64default.StaticInt64(1),
			},
			"scopes": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"local_texts": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"locale": schema.StringAttribute{
							Optional: true,
							Computed: true,
							Default:  stringdefault.StaticString("en"),
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
						"name": schema.StringAttribute{
							Required: true,
						},
						"max_length_msg": schema.StringAttribute{
							Optional: true,
						},
						"min_length_msg": schema.StringAttribute{
							Optional: true,
						},
						"required_msg": schema.StringAttribute{
							Optional: true,
						},
						// optional: in case of datatype is PASSWORD and passwords dont match the message is localised here
						"matchWith": schema.StringAttribute{
							Optional: true,
						},
						// optional: in case of datatype is RADIO, SELECT, MULTISELECT, etc. the localised attribute values are specified here
						"attributes": schema.StringAttribute{
							Optional: true,
						},
					},
				},
			},
			"parent_group_id": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("DEFAULT"),
			},
			"field_type": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("CUSTOM"),
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"CUSTOM", "SYSTEM"}...),
				},
			},
			"data_type": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.OneOf(allowedDataTypes...),
				},
			},
			"field_key": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
				PlanModifiers: []planmodifier.String{
					&validators.UniqueIdentifier{},
				},
			},
			"field_definition": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"maxLength": schema.Int64Attribute{
						Optional: true,
						Validators: []validator.Int64{
							int64validator.AtLeast(1),
						},
					},
					"minLength": schema.Int64Attribute{
						Optional: true,
						Validators: []validator.Int64{
							int64validator.AtLeast(1),
						},
					},
					"minDate": schema.StringAttribute{
						Optional: true,
					},
					"maxDate": schema.StringAttribute{
						Optional: true,
					},
					"initialDateView": schema.StringAttribute{
						Optional: true,
					},
					"initialDate": schema.StringAttribute{
						Optional: true,
					},
					"matchWith": schema.StringAttribute{
						Optional: true,
					},
					// optional: in case of datatype is RADIO, SELECT, MULTISELECT, etc. the keys are specified here
					"attributesKeys": schema.SetAttribute{
						ElementType: types.StringType,
						Optional:    true,
					},
				},
			},
			"base_data_type": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
					stringvalidator.OneOf([]string{"string", "double", "datetime", "bool", "array"}...),
				},
				PlanModifiers: []planmodifier.String{
					&validators.UniqueIdentifier{},
				},
			},
			"created_at": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *RegFieldResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RegFieldConfig
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RegFieldResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RegFieldConfig
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *RegFieldResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) { //nolint:dupl
	var plan, state RegFieldConfig
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RegFieldResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RegFieldConfig
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *RegFieldResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
