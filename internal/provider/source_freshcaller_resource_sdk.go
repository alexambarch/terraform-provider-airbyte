// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceFreshcallerResourceModel) ToSharedSourceFreshcallerCreateRequest() *shared.SourceFreshcallerCreateRequest {
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

	var domain string
	domain = r.Configuration.Domain.ValueString()

	var apiKey string
	apiKey = r.Configuration.APIKey.ValueString()

	requestsPerMinute := new(int64)
	if !r.Configuration.RequestsPerMinute.IsUnknown() && !r.Configuration.RequestsPerMinute.IsNull() {
		*requestsPerMinute = r.Configuration.RequestsPerMinute.ValueInt64()
	} else {
		requestsPerMinute = nil
	}
	startDate := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	syncLagMinutes := new(int64)
	if !r.Configuration.SyncLagMinutes.IsUnknown() && !r.Configuration.SyncLagMinutes.IsNull() {
		*syncLagMinutes = r.Configuration.SyncLagMinutes.ValueInt64()
	} else {
		syncLagMinutes = nil
	}
	configuration := shared.SourceFreshcaller{
		Domain:            domain,
		APIKey:            apiKey,
		RequestsPerMinute: requestsPerMinute,
		StartDate:         startDate,
		SyncLagMinutes:    syncLagMinutes,
	}
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	out := shared.SourceFreshcallerCreateRequest{
		Name:          name,
		DefinitionID:  definitionID,
		WorkspaceID:   workspaceID,
		Configuration: configuration,
		SecretID:      secretID,
	}
	return &out
}

func (r *SourceFreshcallerResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.CreatedAt = types.Int64Value(resp.CreatedAt)
		r.DefinitionID = types.StringValue(resp.DefinitionID)
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceFreshcallerResourceModel) ToSharedSourceFreshcallerPutRequest() *shared.SourceFreshcallerPutRequest {
	var name string
	name = r.Name.ValueString()

	var workspaceID string
	workspaceID = r.WorkspaceID.ValueString()

	var domain string
	domain = r.Configuration.Domain.ValueString()

	var apiKey string
	apiKey = r.Configuration.APIKey.ValueString()

	requestsPerMinute := new(int64)
	if !r.Configuration.RequestsPerMinute.IsUnknown() && !r.Configuration.RequestsPerMinute.IsNull() {
		*requestsPerMinute = r.Configuration.RequestsPerMinute.ValueInt64()
	} else {
		requestsPerMinute = nil
	}
	startDate := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	syncLagMinutes := new(int64)
	if !r.Configuration.SyncLagMinutes.IsUnknown() && !r.Configuration.SyncLagMinutes.IsNull() {
		*syncLagMinutes = r.Configuration.SyncLagMinutes.ValueInt64()
	} else {
		syncLagMinutes = nil
	}
	configuration := shared.SourceFreshcallerUpdate{
		Domain:            domain,
		APIKey:            apiKey,
		RequestsPerMinute: requestsPerMinute,
		StartDate:         startDate,
		SyncLagMinutes:    syncLagMinutes,
	}
	out := shared.SourceFreshcallerPutRequest{
		Name:          name,
		WorkspaceID:   workspaceID,
		Configuration: configuration,
	}
	return &out
}
