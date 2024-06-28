package validators

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

var _ planmodifier.String = UniqueIdentifier{}

type UniqueIdentifier struct{}

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

	if !req.ConfigValue.Equal(req.StateValue) {
		resp.Diagnostics.AddError("Unexpected Resource Configuration",
			fmt.Sprintf("Attribute %s can't be modified. Existing value %s, got %s", req.Path.String(), req.StateValue.ValueString(), req.ConfigValue.ValueString()))
	}
}
