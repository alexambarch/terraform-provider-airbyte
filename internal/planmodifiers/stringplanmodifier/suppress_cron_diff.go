package stringplanmodifier

import (
	"context"

	"github.com/airbytehq/terraform-provider-airbyte/internal/planmodifiers/utils"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

// SuppressDiff returns a plan modifier that propagates a state value into the planned value, when it is Known, and the Plan Value is Unknown
func SuppressCronDiff() planmodifier.String {
	return suppressCronDiff{}
}

// suppressDiff implements the plan modifier.
type suppressCronDiff struct {
}

// Description returns a human-readable description of the plan modifier.
func (m suppressCronDiff) Description(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// MarkdownDescription returns a markdown description of the plan modifier.
func (m suppressCronDiff) MarkdownDescription(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// PlanModifyString implements the plan modification logic.
func (m suppressCronDiff) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	// Do nothing if there is an unknown configuration value
	if req.ConfigValue.IsUnknown() {
		return
	}

	if utils.IsAllStateUnknown(ctx, req.State) {
		return
	}

	planValue := req.PlanValue.ValueString()
	stateValue := req.StateValue.ValueString()

	// Strip the timezone suffix from the state value
	if planValue == stateValue[:len(stateValue)-4] {
		resp.PlanValue = req.StateValue
	}
}
