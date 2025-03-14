// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourcePingdomResourceModel) ToSharedSourcePingdomCreateRequest() *shared.SourcePingdomCreateRequest {
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

	probes := new(string)
	if !r.Configuration.Probes.IsUnknown() && !r.Configuration.Probes.IsNull() {
		*probes = r.Configuration.Probes.ValueString()
	} else {
		probes = nil
	}
	var apiKey string
	apiKey = r.Configuration.APIKey.ValueString()

	resolution := new(shared.Resolution)
	if !r.Configuration.Resolution.IsUnknown() && !r.Configuration.Resolution.IsNull() {
		*resolution = shared.Resolution(r.Configuration.Resolution.ValueString())
	} else {
		resolution = nil
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourcePingdom{
		Probes:     probes,
		APIKey:     apiKey,
		Resolution: resolution,
		StartDate:  startDate,
	}
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	out := shared.SourcePingdomCreateRequest{
		Name:          name,
		DefinitionID:  definitionID,
		WorkspaceID:   workspaceID,
		Configuration: configuration,
		SecretID:      secretID,
	}
	return &out
}

func (r *SourcePingdomResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.CreatedAt = types.Int64Value(resp.CreatedAt)
		r.DefinitionID = types.StringValue(resp.DefinitionID)
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourcePingdomResourceModel) ToSharedSourcePingdomPutRequest() *shared.SourcePingdomPutRequest {
	var name string
	name = r.Name.ValueString()

	var workspaceID string
	workspaceID = r.WorkspaceID.ValueString()

	probes := new(string)
	if !r.Configuration.Probes.IsUnknown() && !r.Configuration.Probes.IsNull() {
		*probes = r.Configuration.Probes.ValueString()
	} else {
		probes = nil
	}
	var apiKey string
	apiKey = r.Configuration.APIKey.ValueString()

	resolution := new(shared.SourcePingdomUpdateResolution)
	if !r.Configuration.Resolution.IsUnknown() && !r.Configuration.Resolution.IsNull() {
		*resolution = shared.SourcePingdomUpdateResolution(r.Configuration.Resolution.ValueString())
	} else {
		resolution = nil
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourcePingdomUpdate{
		Probes:     probes,
		APIKey:     apiKey,
		Resolution: resolution,
		StartDate:  startDate,
	}
	out := shared.SourcePingdomPutRequest{
		Name:          name,
		WorkspaceID:   workspaceID,
		Configuration: configuration,
	}
	return &out
}
