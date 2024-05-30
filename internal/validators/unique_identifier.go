package validators

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

var _ planmodifier.String = UniqueIdentifier{}

type UniqueIdentifier struct {
	AttrName string
}

func (v UniqueIdentifier) Description(_ context.Context) string {
	return "Checks if a unique identifier has been changed"
}

func (v UniqueIdentifier) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v UniqueIdentifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.StateValue.IsNull() || req.PlanValue.IsUnknown() || req.ConfigValue.IsUnknown() {
		return
	}

	var attributeName string
	for attrName, attrConfig := range req.Plan.Schema.GetAttributes() {
		if attrConfig.IsComputed() || attrConfig.IsOptional() || !attrConfig.GetType().Equal(req.ConfigValue.Type(ctx)) {
			continue
		}
		var temp string
		diags := req.Config.GetAttribute(ctx, path.Root(attrName), &temp)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		if temp == req.ConfigValue.ValueString() {
			attributeName = attrName
			break
		}
	}

	if !req.ConfigValue.Equal(req.StateValue) {
		resp.Diagnostics.AddError("Unexpected Resource Configuration",
			fmt.Sprintf("Attribute %s can't be modified.", attributeName))
	}
}
