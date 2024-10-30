package resources

import (
	"context"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/util"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

type AppResource struct {
	BaseResource
}

func NewAppResource() resource.Resource {
	return &AppResource{
		BaseResource: NewBaseResource(
			BaseResourceConfig{
				Name:   RESOURCE_APP,
				Schema: &resourceAppSchema,
			},
		),
	}
}

func (r *AppResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan, config AppConfig
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(plan.ExtractAppConfigs(ctx)...)
	appModel, diags := prepareAppModel(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.cidaasClient.App.Create(*appModel)
	if err != nil {
		resp.Diagnostics.AddError("failed to create app", util.FormatErrorMessage(err))
		return
	}
	resp.Diagnostics.Append(updateStateModel(res, &plan, &config, CREATE)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *AppResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state, config AppConfig
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(state.ExtractAppConfigs(ctx)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.cidaasClient.App.Get(state.ClientID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to read app", util.FormatErrorMessage(err))
		return
	}
	operation := IMPORT
	if !state.ID.IsNull() {
		operation = READ
	}
	resp.Diagnostics.Append(updateStateModel(*res, &state, &config, operation)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *AppResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state, config AppConfig

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(plan.ExtractAppConfigs(ctx)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	resp.Diagnostics.Append(config.ExtractAppConfigs(ctx)...)

	appModel, diags := prepareAppModel(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	appModel.ID = state.ID.ValueString()
	res, err := r.cidaasClient.App.Update(*appModel)
	if err != nil {
		resp.Diagnostics.AddError("failed to update app", util.FormatErrorMessage(err))
		return
	}
	resp.Diagnostics.Append(updateStateModel(res, &plan, &config, UPDATE)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *AppResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state AppConfig
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	err := r.cidaasClient.App.Delete(state.ClientID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to delete app", util.FormatErrorMessage(err))
		return
	}
}

func (r *AppResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("client_id"), req, resp)
}
