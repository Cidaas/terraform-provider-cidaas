package resources

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/cidaas"
	"github.com/Cidaas/terraform-provider-cidaas/helpers/util"
	"github.com/Cidaas/terraform-provider-cidaas/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const layout = "2006-01-02T15:04:05Z"

var allowedDataTypes = []string{"TEXT", "NUMBER", "SELECT", "MULTISELECT", "RADIO", "CHECKBOX", "PASSWORD", "DATE", "URL", "EMAIL",
	"TEXTAREA", "MOBILE", "CONSENT", "JSON_STRING", "USERNAME", "ARRAY", "GROUPING", "DAYDATE"}

type RegFieldConfig struct {
	ID                                  types.String `tfsdk:"id"`
	BaseDataType                        types.String `tfsdk:"base_data_type"`
	ParentGroupId                       types.String `tfsdk:"parent_group_id"`
	FieldType                           types.String `tfsdk:"field_type"`
	DataType                            types.String `tfsdk:"data_type"`
	FieldKey                            types.String `tfsdk:"field_key"`
	Required                            types.Bool   `tfsdk:"required"`
	Internal                            types.Bool   `tfsdk:"internal"`
	Claimable                           types.Bool   `tfsdk:"claimable"`
	IsSearchable                        types.Bool   `tfsdk:"is_searchable"`
	Enabled                             types.Bool   `tfsdk:"enabled"`
	Unique                              types.Bool   `tfsdk:"unique"`
	OverwriteWithNullFromSocialProvider types.Bool   `tfsdk:"overwrite_with_null_value_from_social_provider"`
	ReadOnly                            types.Bool   `tfsdk:"read_only"`
	IsGroup                             types.Bool   `tfsdk:"is_group"`
	IsList                              types.Bool   `tfsdk:"is_list"`
	Order                               types.Int64  `tfsdk:"order"`
	Scopes                              types.Set    `tfsdk:"scopes"`
	ConsentRefs                         types.Set    `tfsdk:"consent_refs"`
	LocalTexts                          types.List   `tfsdk:"local_texts"`
	FieldDefinition                     types.Object `tfsdk:"field_definition"`

	localTexts      []*LocalTexts
	fieldDefinition *FieldDefinition
}

type LocalTexts struct {
	Locale       types.String  `tfsdk:"locale"`
	Name         types.String  `tfsdk:"name"`
	MaxLengthMsg types.String  `tfsdk:"max_length_msg"`
	MinLengthMsg types.String  `tfsdk:"min_length_msg"`
	RequiredMsg  types.String  `tfsdk:"required_msg"`
	Attributes   []*Attributes `tfsdk:"attributes"`
	ConsentLabel types.Object  `tfsdk:"consent_label"`
}

type Attributes struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

type Consent struct {
	Label     types.String `tfsdk:"label"`
	LabelText types.String `tfsdk:"label_text"`
}

type FieldDefinition struct {
	MaxLength       types.Int64  `tfsdk:"max_length"`
	MinLength       types.Int64  `tfsdk:"min_length"`
	MinDate         types.String `tfsdk:"min_date"`
	MaxDate         types.String `tfsdk:"max_date"`
	InitialDateView types.String `tfsdk:"initial_date_view"`
	InitialDate     types.String `tfsdk:"initial_date"`
}

type RegFieldResource struct {
	cidaasClient *cidaas.Client
}

func NewRegFieldResource() resource.Resource {
	return &RegFieldResource{}
}

func (r *RegFieldResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_registration_field"
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
		MarkdownDescription: "The `cidaas_registration_page_field` in the provider allows management of registration fields in the Cidaas system." +
			" This resource enables you to configure and customize the fields displayed during user registration." +
			"\n\n Ensure that the below scopes are assigned to the client with the specified `client_id`:" +
			"\n- cidaas:field_setup_read" +
			"\n- cidaas:field_setup_write" +
			"\n- cidaas:field_setup_delete\n",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The ID of the resource",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			// "string", "double", "datetime", "bool", "array"
			"base_data_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The base data type of the field. This is computed property.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"parent_group_id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The ID of the parent registration group. Defaults to `DEFAULT` if not provided.",
				Default:             stringdefault.StaticString("DEFAULT"),
			},
			"field_type": schema.StringAttribute{
				Optional: true,
				Computed: true,
				MarkdownDescription: "Specifies whether the field type is `SYSTEM` or `CUSTOM`. Defaults to `CUSTOM`." +
					" This cannot be modified for an existing resource. `SYSTEM` fields cannot be created but can be modified. To modify an existing field import it first and then update.",
				Default: stringdefault.StaticString("CUSTOM"),
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"CUSTOM", "SYSTEM"}...),
				},
				PlanModifiers: []planmodifier.String{
					&validators.UniqueIdentifier{},
					&fieldTypeModifier{},
				},
			},
			"data_type": schema.StringAttribute{
				Required: true,
				MarkdownDescription: "The data type of the field. This cannot be modified for an existing resource." +
					fmt.Sprintf(" Allowed values are %s", func() string {
						var temp string
						for _, v := range allowedDataTypes {
							temp += fmt.Sprintf("`%s`,", v)
						}
						return temp
					}()),
				Validators: []validator.String{
					stringvalidator.OneOf(allowedDataTypes...),
					&dataTypeValidator{},
				},
				PlanModifiers: []planmodifier.String{
					&validators.UniqueIdentifier{},
				},
			},
			"field_key": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The unique identifier of the registration field. This cannot be modified for an existing resource.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
				PlanModifiers: []planmodifier.String{
					&validators.UniqueIdentifier{},
				},
			},
			"required": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Flag to mark if a field is required in registration. Defaults set to `false`",
				Default:             booldefault.StaticBool(false),
				Validators: []validator.Bool{
					&validateIsRequiredMsgAvailable{},
				},
			},
			"internal": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Flag to mark if a field is internal. Defaults set to `false`",
				Default:             booldefault.StaticBool(false),
			},
			"claimable": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Flag to mark if a field is claimable. Defaults set to `true`",
				Default:             booldefault.StaticBool(true),
			},
			"is_searchable": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Flag to mark if a field is searchable. Defaults set to `true`",
				Default:             booldefault.StaticBool(true),
			},
			"enabled": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Flag to mark if a field is enabled. Defaults set to `true`",
				Default:             booldefault.StaticBool(true),
			},
			"unique": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Flag to mark if a field is unique. Defaults set to `false`",
				Default:             booldefault.StaticBool(false),
			},
			// set to true if you want the value should be reset by identity provider
			"overwrite_with_null_value_from_social_provider": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Set to true if you want the value should be reset by identity provider. Defaults set to `false`",
				Default:             booldefault.StaticBool(false),
			},
			"read_only": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Flag to mark if a field is read only. Defaults set to `false`",
				Default:             booldefault.StaticBool(false),
			},
			"is_group": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				MarkdownDescription: "Setting is_group to `true` creates a registration field group. Defaults set to `false`" +
					" The data_type attribute must be set to TEXT when is_group is true. ",
				Default: booldefault.StaticBool(false),
				Validators: []validator.Bool{
					&isGroupValidator{},
				},
			},
			"is_list": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
			},
			// optional: Order of the Field in the UI
			"order": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The attribute order is used to set the order of the Field in the UI. Defaults set to `1`",
				Default:             int64default.StaticInt64(1),
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"scopes": schema.SetAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				MarkdownDescription: "The scopes of the registration field.",
			},
			"consent_refs": schema.SetAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				MarkdownDescription: "List of consents(the ids of the consent in cidaas must be passed) in registration. The data type must be `CONSENT` in this case",
			},
			"local_texts": schema.ListNestedAttribute{
				Required:            true,
				MarkdownDescription: "The localized detail of the registration field.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"locale": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "The locale of the field. example: de-DE.",
							Default:             stringdefault.StaticString("en"),
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
							Required:            true,
							MarkdownDescription: "Then name of the field in the local configured. for exmaple: in **en-US** the name is `Sample Field` in de-DE `Beispielfeld`.",
						},
						"max_length_msg": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "warning/error msg to show to the user when user exceeds the maximum character configured. This is applicable only for the attributes of base_data_type string.",
						},
						"min_length_msg": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "warning/error msg to show to the user when user don't provide the minimum character required. This is applicable only for the attributes of base_data_type string.",
						},
						"required_msg": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "When the flag required is set to true the required_msg must be provided. required_msg is shown if user does not provide a required field.",
						},
						// optional: in case of datatype is RADIO, SELECT, MULTISELECT, etc. the localised attribute values are specified here
						"attributes": schema.ListNestedAttribute{
							Optional:            true,
							MarkdownDescription: "The field attributes must be provided for the data_type SELECT, MULTISELECT and RADIO. it's an array of key value pairs. Example provided in the example section.",
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Required: true,
									},
									"value": schema.StringAttribute{
										Required: true,
									},
								},
							},
						},
						"consent_label": schema.SingleNestedAttribute{
							Optional:            true,
							MarkdownDescription: "required when data_type is CONSENT. Example provided in the example section.",
							Attributes: map[string]schema.Attribute{
								"label": schema.StringAttribute{
									Required: true,
								},
								"label_text": schema.StringAttribute{
									Required: true,
								},
							},
						},
					},
				},
			},
			"field_definition": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"max_length": schema.Int64Attribute{
						Optional:            true,
						MarkdownDescription: "The maximum length of a string type attribute.",
						Validators: []validator.Int64{
							&validateIsMaxMinMsgAvailable{},
						},
					},
					"min_length": schema.Int64Attribute{
						Optional:            true,
						MarkdownDescription: "The minimum length of a string type attribute",
						Validators: []validator.Int64{
							&validateIsMaxMinMsgAvailable{},
						},
					},
					"min_date": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "The earliest date a user can select. Applicable only for DATE attributes. Example format: `2024-06-28T18:30:00Z`.",
						Validators: []validator.String{
							&dateTypeValidator{},
							&dateValidator{},
						},
					},
					"max_date": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "The maximum date a user can select. Applicable only for DATE attributes. Example format: `2024-06-28T18:30:00Z`.",
						Validators: []validator.String{
							&dateTypeValidator{},
							&dateValidator{},
						},
					},
					"initial_date_view": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "The view of the calender. Applicable only for DATE attributes. Allowed values: `month`, `year` and `multi-year`",
						Validators: []validator.String{
							&dateTypeValidator{},
							stringvalidator.OneOf("month", "year", "multi-year"),
						},
					},
					"initial_date": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "The initial date. Applicable only for DATE attributes. Example format: `2024-06-28T18:30:00Z`.",
						Validators: []validator.String{
							&dateTypeValidator{},
							&dateValidator{},
						},
					},
				},
				Default: objectdefault.StaticValue(types.ObjectValueMust(
					map[string]attr.Type{
						"max_length":        types.Int64Type,
						"min_length":        types.Int64Type,
						"min_date":          types.StringType,
						"max_date":          types.StringType,
						"initial_date_view": types.StringType,
						"initial_date":      types.StringType,
					},
					map[string]attr.Value{
						"max_length":        types.Int64Null(),
						"min_length":        types.Int64Null(),
						"min_date":          types.StringNull(),
						"max_date":          types.StringNull(),
						"initial_date_view": types.StringNull(),
						"initial_date":      types.StringNull(),
					})),
			},
		},
	}
}

func (rfc *RegFieldConfig) ExtractConfigs(ctx context.Context) diag.Diagnostics {
	var diags diag.Diagnostics
	if !rfc.FieldDefinition.IsNull() {
		rfc.fieldDefinition = &FieldDefinition{}
		diags = rfc.FieldDefinition.As(ctx, rfc.fieldDefinition, basetypes.ObjectAsOptions{})
	}
	if !rfc.LocalTexts.IsNull() {
		rfc.localTexts = make([]*LocalTexts, 0, len(rfc.LocalTexts.Elements()))
		diags = rfc.LocalTexts.ElementsAs(ctx, &rfc.localTexts, false)
	}
	return diags
}

func (r *RegFieldResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RegFieldConfig
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(plan.ExtractConfigs(ctx)...)
	rfModel, diags := prepareRegFieldModel(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.cidaasClient.RegField.Upsert(*rfModel)
	if err != nil {
		resp.Diagnostics.AddError("failed to create registration field", fmt.Sprintf("Error: %+v", err.Error()))
		return
	}
	plan.ID = types.StringValue(res.Data.ID)
	plan.BaseDataType = types.StringValue(res.Data.BaseDataType)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RegFieldResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RegFieldConfig
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	res, err := r.cidaasClient.RegField.Get(state.FieldKey.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to read role", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	state.ID = util.StringValueOrNull(&res.Data.ID)
	state.BaseDataType = util.StringValueOrNull(&res.Data.BaseDataType)
	state.ParentGroupId = util.StringValueOrNull(&res.Data.ParentGroupID)
	state.FieldType = util.StringValueOrNull(&res.Data.FieldType)
	state.DataType = util.StringValueOrNull(&res.Data.DataType)
	state.FieldKey = util.StringValueOrNull(&res.Data.FieldKey)
	state.Required = util.BoolValueOrNull(&res.Data.Required)
	state.Internal = util.BoolValueOrNull(&res.Data.Internal)
	state.Claimable = util.BoolValueOrNull(&res.Data.Claimable)
	state.IsSearchable = util.BoolValueOrNull(&res.Data.IsSearchable)
	state.Enabled = util.BoolValueOrNull(&res.Data.Enabled)
	state.Unique = util.BoolValueOrNull(&res.Data.Unique)
	state.OverwriteWithNullFromSocialProvider = util.BoolValueOrNull(&res.Data.OverwriteWithNullValueFromSocialProvider)
	state.ReadOnly = util.BoolValueOrNull(&res.Data.ReadOnly)
	state.IsGroup = util.BoolValueOrNull(&res.Data.IsGroup)
	state.IsList = util.BoolValueOrNull(&res.Data.IsList)
	state.Scopes = util.SetValueOrNull(res.Data.Scopes)
	state.ConsentRefs = util.SetValueOrNull(res.Data.ConsentRefs)
	state.Order = util.Int64ValueOrNull(&res.Data.Order)

	var localTextsObjectValues []attr.Value
	typesOfAttribute := map[string]attr.Type{
		"key":   types.StringType,
		"value": types.StringType,
	}

	typesOfConsentLabel := map[string]attr.Type{
		"label":      types.StringType,
		"label_text": types.StringType,
	}

	localTextObjectType := types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"locale":         types.StringType,
			"name":           types.StringType,
			"max_length_msg": types.StringType,
			"min_length_msg": types.StringType,
			"required_msg":   types.StringType,
			"attributes":     types.ListType{ElemType: types.ObjectType{AttrTypes: typesOfAttribute}},
			"consent_label":  types.ObjectType{AttrTypes: typesOfConsentLabel},
		},
	}
	for _, lt := range res.Data.LocaleTexts {
		var attributeValues []attr.Value
		for _, v := range lt.Attributes {
			attributeValue := types.ObjectValueMust(
				typesOfAttribute,
				map[string]attr.Value{
					"key":   util.StringValueOrNull(&v.Key),
					"value": util.StringValueOrNull(&v.Value),
				})
			attributeValues = append(attributeValues, attributeValue)
		}
		objValue := types.ObjectValueMust(
			localTextObjectType.AttrTypes,
			map[string]attr.Value{
				"locale":         util.StringValueOrNull(&lt.Locale),
				"name":           util.StringValueOrNull(&lt.Name),
				"max_length_msg": util.StringValueOrNull(&lt.MaxLengthErrorMsg),
				"min_length_msg": util.StringValueOrNull(&lt.MinLengthErrorMsg),
				"required_msg":   util.StringValueOrNull(&lt.RequiredMsg),
				"attributes": func() types.List {
					if !(len(lt.Attributes) > 0) {
						return types.ListNull(types.ObjectType{AttrTypes: typesOfAttribute})
					}
					return types.ListValueMust(
						types.ObjectType{AttrTypes: typesOfAttribute},
						attributeValues,
					)
				}(),
				"consent_label": func() types.Object {
					if lt.ConsentLabel == nil {
						return types.ObjectNull(typesOfConsentLabel)
					}
					return types.ObjectValueMust(
						typesOfConsentLabel,
						map[string]attr.Value{
							"label":      util.StringValueOrNull(&lt.ConsentLabel.Label),
							"label_text": util.StringValueOrNull(&lt.ConsentLabel.LabelText),
						},
					)
				}(),
			})
		localTextsObjectValues = append(localTextsObjectValues, objValue)
	}
	var diags diag.Diagnostics
	state.LocalTexts, diags = types.ListValueFrom(ctx, localTextObjectType, localTextsObjectValues)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if res.Data.FieldDefinition != nil {
		fd, diags := types.ObjectValue(
			map[string]attr.Type{
				"max_length":        types.Int64Type,
				"min_length":        types.Int64Type,
				"min_date":          types.StringType,
				"max_date":          types.StringType,
				"initial_date_view": types.StringType,
				"initial_date":      types.StringType,
			},
			map[string]attr.Value{
				"max_length":        util.Int64ValueOrNull(res.Data.FieldDefinition.MaxLength),
				"min_length":        util.Int64ValueOrNull(res.Data.FieldDefinition.MinLength),
				"min_date":          util.TimeValueOrNull(res.Data.FieldDefinition.MinDate),
				"max_date":          util.TimeValueOrNull(res.Data.FieldDefinition.MaxDate),
				"initial_date_view": util.StringValueOrNull(&res.Data.FieldDefinition.InitialDateView),
				"initial_date":      util.TimeValueOrNull(res.Data.FieldDefinition.InitialDate),
			})
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		state.FieldDefinition = fd
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *RegFieldResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) { //nolint:dupl
	var plan, state RegFieldConfig
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(plan.ExtractConfigs(ctx)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	fieldModel, diags := prepareRegFieldModel(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	fieldModel.ID = state.ID.ValueString()
	_, err := r.cidaasClient.RegField.Upsert(*fieldModel)
	if err != nil {
		resp.Diagnostics.AddError("failed to update registration field", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RegFieldResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RegFieldConfig
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	err := r.cidaasClient.RegField.Delete(state.FieldKey.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to registration field", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
}

func (r *RegFieldResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("field_key"), req, resp)
}

func prepareRegFieldModel(ctx context.Context, plan RegFieldConfig) (*cidaas.RegistrationFieldConfig, diag.Diagnostics) {
	var regConfig cidaas.RegistrationFieldConfig
	regConfig.Internal = plan.Internal.ValueBool()
	regConfig.ReadOnly = plan.ReadOnly.ValueBool()
	regConfig.Claimable = plan.Claimable.ValueBool()
	regConfig.Required = plan.Required.ValueBool()
	regConfig.Unique = plan.Unique.ValueBool()
	regConfig.IsSearchable = plan.IsSearchable.ValueBool()
	regConfig.OverwriteWithNullValueFromSocialProvider = plan.OverwriteWithNullFromSocialProvider.ValueBool()
	regConfig.Enabled = plan.Enabled.ValueBool()
	regConfig.IsGroup = plan.IsGroup.ValueBool()
	regConfig.IsList = plan.IsList.ValueBool()
	regConfig.ParentGroupID = plan.ParentGroupId.ValueString()
	regConfig.FieldType = plan.FieldType.ValueString()
	regConfig.FieldKey = plan.FieldKey.ValueString()
	regConfig.DataType = plan.DataType.ValueString()
	regConfig.Order = plan.Order.ValueInt64()

	className := "FieldSetup"
	if regConfig.FieldType == "SYSTEM" {
		className = "de.cidaas.core.db.RegistrationFieldSetup"
	}
	regConfig.ClassName = className

	diag := plan.Scopes.ElementsAs(ctx, &regConfig.Scopes, false)
	if diag.HasError() {
		return nil, diag
	}

	if plan.DataType.ValueString() == "CONSENT" {
		diag = plan.ConsentRefs.ElementsAs(ctx, &regConfig.ConsentRefs, false)
		if diag.HasError() {
			return nil, diag
		}
	}

	var attrKeys []string
	setLocalTexts := func(source []*LocalTexts, target *[]*cidaas.LocaleText) error {
		for _, s := range source {
			tempLocalText := &cidaas.LocaleText{
				Locale:            s.Locale.ValueString(),
				Language:          util.GetLanguageForLocale(s.Locale.ValueString()),
				Name:              s.Name.ValueString(),
				MaxLengthErrorMsg: s.MaxLengthMsg.ValueString(),
				MinLengthErrorMsg: s.MinLengthMsg.ValueString(),
				RequiredMsg:       s.RequiredMsg.ValueString(),
			}
			cidaasAttribues := []*cidaas.Attribute{}
			for _, v := range s.Attributes {
				cidaasAttribues = append(cidaasAttribues, &cidaas.Attribute{
					Key:   v.Key.ValueString(),
					Value: v.Value.ValueString(),
				})
				attrKeys = append(attrKeys, v.Key.ValueString())
			}
			if len(s.Attributes) > 0 {
				tempLocalText.Attributes = cidaasAttribues
			}
			if !s.ConsentLabel.IsNull() {
				consent := Consent{}
				diag = s.ConsentLabel.As(ctx, &consent, basetypes.ObjectAsOptions{})
				if diag.HasError() {
					return errors.New("failed to parse consent_label")
				}
				tempLocalText.ConsentLabel = &cidaas.Consent{
					Label:     consent.Label.ValueString(),
					LabelText: consent.LabelText.ValueString(),
				}
			}
			*target = append(*target, tempLocalText)
		}
		return nil
	}
	if len(plan.localTexts) > 0 {
		err := setLocalTexts(plan.localTexts, &regConfig.LocaleTexts)
		if err != nil {
			diag.AddError("Failed ti create registration field request payload", err.Error())
			return nil, diag
		}
	}

	if !plan.FieldDefinition.IsNull() {
		regConfig.FieldDefinition = &cidaas.FieldDefinition{
			MinLength:       plan.fieldDefinition.MinLength.ValueInt64Pointer(),
			MaxLength:       plan.fieldDefinition.MaxLength.ValueInt64Pointer(),
			InitialDateView: plan.fieldDefinition.InitialDateView.ValueString(),
		}
		if len(attrKeys) > 0 {
			regConfig.FieldDefinition.AttributesKeys = attrKeys
		}
		if !plan.fieldDefinition.MinDate.IsNull() {
			minDate, err := time.Parse(layout, plan.fieldDefinition.MinDate.ValueString())
			if err != nil {
				diag.AddError("Parse Error", "failed to parse min_date configured")
			}
			regConfig.FieldDefinition.MinDate = &minDate
		}
		if !plan.fieldDefinition.MaxDate.IsNull() {
			maxDate, err := time.Parse(layout, plan.fieldDefinition.MaxDate.ValueString())
			if err != nil {
				diag.AddError("Parse Error", "failed to parse max_date configured")
			}
			regConfig.FieldDefinition.MaxDate = &maxDate
		}
		if !plan.fieldDefinition.InitialDate.IsNull() {
			initialDate, err := time.Parse(layout, plan.fieldDefinition.InitialDate.ValueString())
			if err != nil {
				diag.AddError("Parse Error", "failed to parse initial_date configured")
			}
			regConfig.FieldDefinition.InitialDate = &initialDate
		}
	}
	return &regConfig, nil
}

// custom validations
var (
	_ validator.Bool      = validateIsRequiredMsgAvailable{}
	_ validator.Int64     = validateIsMaxMinMsgAvailable{}
	_ planmodifier.String = fieldTypeModifier{}
	_ validator.String    = dateTypeValidator{}
	_ validator.String    = dateValidator{}
	_ validator.String    = dataTypeValidator{}
	_ validator.Bool      = isGroupValidator{}
)

type (
	validateIsRequiredMsgAvailable struct{}
	validateIsMaxMinMsgAvailable   struct{}
	fieldTypeModifier              struct{}
	dateTypeValidator              struct{}
	dateValidator                  struct{}
	dataTypeValidator              struct{}
	isGroupValidator               struct{}
)

func (v validateIsRequiredMsgAvailable) Description(_ context.Context) string {
	return "msg is required when enabled is true"
}

func (v validateIsRequiredMsgAvailable) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v validateIsRequiredMsgAvailable) ValidateBool(ctx context.Context, req validator.BoolRequest, resp *validator.BoolResponse) {
	var config RegFieldConfig
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	resp.Diagnostics.Append(config.ExtractConfigs(ctx)...)
	if !req.ConfigValue.IsNull() && req.ConfigValue.ValueBool() {
		if len(config.localTexts) > 0 {
			for _, v := range config.localTexts {
				if v.RequiredMsg.IsNull() || v.RequiredMsg.ValueString() == "" {
					resp.Diagnostics.AddError(
						"Validation Error",
						fmt.Sprintf("The attribute local_texts.required_msg is required when %s is set to true", req.Path.String()),
					)
					return
				}
			}
		}
	} else {
		if len(config.localTexts) > 0 {
			for _, v := range config.localTexts {
				if !v.RequiredMsg.IsNull() {
					resp.Diagnostics.AddError(
						"Validation Error",
						fmt.Sprintf("The attribute local_texts.required_msg is not allowed in config when attribute %s is set to false", req.Path.String()),
					)
					return
				}
			}
		}
	}
}

func (v validateIsMaxMinMsgAvailable) Description(_ context.Context) string {
	return "max_length & min_length validation"
}

func (v validateIsMaxMinMsgAvailable) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v validateIsMaxMinMsgAvailable) ValidateInt64(ctx context.Context, req validator.Int64Request, resp *validator.Int64Response) {

	var config RegFieldConfig
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	resp.Diagnostics.Append(config.ExtractConfigs(ctx)...)

	if !req.ConfigValue.IsNull() {
		if req.ConfigValue.ValueInt64() < 1 {
			resp.Diagnostics.AddError(
				"Validation Error",
				fmt.Sprintf("The attribute %s must be greater than 0", req.Path.String()),
			)
			return
		}
		if req.Path.String() == "field_definition.min_length" {
			for _, v := range config.localTexts {
				if v.MinLengthMsg.IsNull() || v.MinLengthMsg.ValueString() == "" {
					resp.Diagnostics.AddError(
						"Validation Error",
						fmt.Sprintf("The attribute local_texts.min_length_msg can not be empty when %s is set", req.Path.String()),
					)
					return
				}
			}

			if config.fieldDefinition.MaxLength.IsNull() || config.fieldDefinition.MaxLength.ValueInt64() <= 0 {
				resp.Diagnostics.AddError(
					"Validation Error",
					fmt.Sprintf("The attribute field_definition.max_length can not be empty when %s is set", req.Path.String()),
				)
				return
			}

			if config.fieldDefinition.MaxLength.ValueInt64() < config.fieldDefinition.MinLength.ValueInt64() {
				resp.Diagnostics.AddError(
					"Validation Error",
					fmt.Sprintf("The attribute field_definition.max_length can not be less than %s", req.Path.String()),
				)
				return
			}
		}
		if req.Path.String() == "field_definition.max_length" {
			for _, v := range config.localTexts {
				if v.MaxLengthMsg.IsNull() || v.MaxLengthMsg.ValueString() == "" {
					resp.Diagnostics.AddError(
						"Validation Error",
						fmt.Sprintf("The attribute local_texts.max_length_msg can not be empty when %s is set", req.Path.String()),
					)
					return
				}
			}
		}
	} else {
		if req.Path.String() == "field_definition.min_length" {
			for _, v := range config.localTexts {
				if !v.MinLengthMsg.IsNull() {
					resp.Diagnostics.AddError(
						"Validation Error",
						fmt.Sprintf("The attribute local_texts.min_length_msg is not allowed in config when %s is not set", req.Path.String()),
					)
					return
				}
			}
		}
		if req.Path.String() == "field_definition.max_length" {
			for _, v := range config.localTexts {
				if !v.MaxLengthMsg.IsNull() {
					resp.Diagnostics.AddError(
						"Validation Error",
						fmt.Sprintf("The attribute local_texts.max_length_msg is not allowed in config when %s is not set", req.Path.String()),
					)
					return
				}
			}
		}
	}
}

func (v fieldTypeModifier) Description(_ context.Context) string {
	return "Checks if field_type is SYSTEM while creating a field"
}

func (v fieldTypeModifier) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v fieldTypeModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.StateValue.IsNull() && req.ConfigValue.Equal(types.StringValue("SYSTEM")) {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configuration",
			"field with SYSYTEM field_type cannot be created. SYSTEM fields can only be updated. To update an existing field please import first",
		)
	}
}

func (v dateTypeValidator) Description(_ context.Context) string {
	return "Checks min_date, max_date, initial_date_view and initiate_date"
}

func (v dateTypeValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v dateTypeValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if !req.ConfigValue.IsNull() {
		var config RegFieldConfig
		resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
		if config.DataType.ValueString() != "DATE" {
			resp.Diagnostics.AddError(
				"Validation Error",
				fmt.Sprintf("The attribute %s is only allowed when data_type is DATE", req.Path.String()),
			)
			return
		}
	}
}

func (v dateValidator) Description(_ context.Context) string {
	return "Validates that the value is a valid date."
}

func (v dateValidator) MarkdownDescription(_ context.Context) string {
	return v.Description(context.Background())
}

func (v dateValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	_, err := time.Parse(layout, req.ConfigValue.ValueString())
	if err != nil {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid Date Format",
			fmt.Sprintf("Attribute %s expected to be a valid ISO 8601 date in the format %s.", req.Path.String(), layout),
		)
	}
}

func (v dataTypeValidator) Description(_ context.Context) string {
	return "Validates that the value is a valid date."
}

func (v dataTypeValidator) MarkdownDescription(_ context.Context) string {
	return v.Description(context.Background())
}

func (v dataTypeValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	var config RegFieldConfig
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	resp.Diagnostics.Append(config.ExtractConfigs(ctx)...)

	if req.ConfigValue.ValueString() == "DATE" &&
		(config.FieldDefinition.IsNull() ||
			config.fieldDefinition.MinDate.IsNull() ||
			config.fieldDefinition.MaxDate.IsNull() ||
			config.fieldDefinition.InitialDate.IsNull()) {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configuration",
			"Attributes min_date, max_date, initial_date and initial_date_view can not be empty when data_type is DATE.",
		)
	}

	if req.ConfigValue.ValueString() != "CONSENT" &&
		!config.ConsentRefs.IsNull() {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configuration",
			"Attribute consent_refs is only allowed when data_type is set to CONSENT.",
		)
	}

	attrKeysRequiredDataTypes := []string{"SELECT", "RADIO", "MULTISELECT"}
	if util.StringInSlice(req.ConfigValue.ValueString(), attrKeysRequiredDataTypes) {
		for _, v := range config.localTexts {
			if !(len(v.Attributes) > 0) {
				resp.Diagnostics.AddError(
					"Unexpected Resource Configuration",
					fmt.Sprintf("Attributes local_texts[i].attributes can not be empty when data_type is %s.", req.ConfigValue.ValueString()),
				)
			}
		}
	}

	noMaxMinLengthDataTypes := []string{"CHECKBOX", "CONSENT", "JSON_STRING", "ARRAY", "NUMBER", "SELECT", "RADIO", "MULTISELECT", "MOBILE", "JSON_STRING"}
	if util.StringInSlice(req.ConfigValue.ValueString(), noMaxMinLengthDataTypes) {
		if config.FieldDefinition.IsNull() {
			return
		}
		if !config.fieldDefinition.MinLength.IsNull() || !config.fieldDefinition.MaxLength.IsNull() {
			resp.Diagnostics.AddError(
				"Unexpected Resource Configuration",
				fmt.Sprintf("Attributes min_length, max_length are not allowed in config when the data_type is %s.", req.ConfigValue.ValueString()),
			)
		}
	}

	noAttributesDataTypes := []string{"TEXT", "NUMBER", "CHECKBOX", "PASSWORD", "DATE", "URL", "EMAIL",
		"TEXTAREA", "MOBILE", "CONSENT", "JSON_STRING", "USERNAME", "ARRAY", "GROUPING", "DAYDATE"}
	if util.StringInSlice(req.ConfigValue.ValueString(), noAttributesDataTypes) {
		for _, v := range config.localTexts {
			if len(v.Attributes) > 0 {
				resp.Diagnostics.AddError(
					"Unexpected Resource Configuration",
					fmt.Sprintf("param local_texts[i].attributes not allowed in config when the data_type is %s.", req.ConfigValue.ValueString()),
				)
			}
		}
	}
}

func (v isGroupValidator) Description(_ context.Context) string {
	return "Validates a registration field group"
}

func (v isGroupValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v isGroupValidator) ValidateBool(ctx context.Context, req validator.BoolRequest, resp *validator.BoolResponse) {
	if !req.ConfigValue.IsNull() && req.ConfigValue.ValueBool() {
		var config RegFieldConfig
		resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
		if config.DataType.ValueString() != "TEXT" {
			resp.Diagnostics.AddError(
				"Unexpected Resource Configuration",
				"The data_type attribute must be set to TEXT when is_group is true. Setting is_group to true creates a registration field group.",
			)
		}
	}
}
