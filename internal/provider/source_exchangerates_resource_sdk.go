// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	customTypes "github.com/airbytehq/terraform-provider-airbyte/internal/sdk/types"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceExchangeRatesResourceModel) ToSharedSourceExchangeRatesCreateRequest() *shared.SourceExchangeRatesCreateRequest {
	var name string
	name = r.Name.ValueString()

	definitionID := new(string)
	if !r.DefinitionID.IsUnknown() && !r.DefinitionID.IsNull() {
		*definitionID = r.DefinitionID.ValueString()
	} else {
		definitionID = nil
	}
	var workspaceID string
	workspaceID = r.WorkspaceID.ValueString()

	startDate := customTypes.MustDateFromString(r.Configuration.StartDate.ValueString())
	var accessKey string
	accessKey = r.Configuration.AccessKey.ValueString()

	base := new(string)
	if !r.Configuration.Base.IsUnknown() && !r.Configuration.Base.IsNull() {
		*base = r.Configuration.Base.ValueString()
	} else {
		base = nil
	}
	ignoreWeekends := new(bool)
	if !r.Configuration.IgnoreWeekends.IsUnknown() && !r.Configuration.IgnoreWeekends.IsNull() {
		*ignoreWeekends = r.Configuration.IgnoreWeekends.ValueBool()
	} else {
		ignoreWeekends = nil
	}
	configuration := shared.SourceExchangeRates{
		StartDate:      startDate,
		AccessKey:      accessKey,
		Base:           base,
		IgnoreWeekends: ignoreWeekends,
	}
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	out := shared.SourceExchangeRatesCreateRequest{
		Name:          name,
		DefinitionID:  definitionID,
		WorkspaceID:   workspaceID,
		Configuration: configuration,
		SecretID:      secretID,
	}
	return &out
}

func (r *SourceExchangeRatesResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.CreatedAt = types.Int64Value(resp.CreatedAt)
		r.DefinitionID = types.StringValue(resp.DefinitionID)
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceExchangeRatesResourceModel) ToSharedSourceExchangeRatesPutRequest() *shared.SourceExchangeRatesPutRequest {
	var name string
	name = r.Name.ValueString()

	var workspaceID string
	workspaceID = r.WorkspaceID.ValueString()

	startDate := customTypes.MustDateFromString(r.Configuration.StartDate.ValueString())
	var accessKey string
	accessKey = r.Configuration.AccessKey.ValueString()

	base := new(string)
	if !r.Configuration.Base.IsUnknown() && !r.Configuration.Base.IsNull() {
		*base = r.Configuration.Base.ValueString()
	} else {
		base = nil
	}
	ignoreWeekends := new(bool)
	if !r.Configuration.IgnoreWeekends.IsUnknown() && !r.Configuration.IgnoreWeekends.IsNull() {
		*ignoreWeekends = r.Configuration.IgnoreWeekends.ValueBool()
	} else {
		ignoreWeekends = nil
	}
	configuration := shared.SourceExchangeRatesUpdate{
		StartDate:      startDate,
		AccessKey:      accessKey,
		Base:           base,
		IgnoreWeekends: ignoreWeekends,
	}
	out := shared.SourceExchangeRatesPutRequest{
		Name:          name,
		WorkspaceID:   workspaceID,
		Configuration: configuration,
	}
	return &out
}
