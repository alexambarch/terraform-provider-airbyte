// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceHeightResourceModel) ToSharedSourceHeightCreateRequest() *shared.SourceHeightCreateRequest {
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

	var apiKey string
	apiKey = r.Configuration.APIKey.ValueString()

	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	searchQuery := new(string)
	if !r.Configuration.SearchQuery.IsUnknown() && !r.Configuration.SearchQuery.IsNull() {
		*searchQuery = r.Configuration.SearchQuery.ValueString()
	} else {
		searchQuery = nil
	}
	configuration := shared.SourceHeight{
		APIKey:      apiKey,
		StartDate:   startDate,
		SearchQuery: searchQuery,
	}
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	out := shared.SourceHeightCreateRequest{
		Name:          name,
		DefinitionID:  definitionID,
		WorkspaceID:   workspaceID,
		Configuration: configuration,
		SecretID:      secretID,
	}
	return &out
}

func (r *SourceHeightResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.CreatedAt = types.Int64Value(resp.CreatedAt)
		r.DefinitionID = types.StringValue(resp.DefinitionID)
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceHeightResourceModel) ToSharedSourceHeightPutRequest() *shared.SourceHeightPutRequest {
	var name string
	name = r.Name.ValueString()

	var workspaceID string
	workspaceID = r.WorkspaceID.ValueString()

	var apiKey string
	apiKey = r.Configuration.APIKey.ValueString()

	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	searchQuery := new(string)
	if !r.Configuration.SearchQuery.IsUnknown() && !r.Configuration.SearchQuery.IsNull() {
		*searchQuery = r.Configuration.SearchQuery.ValueString()
	} else {
		searchQuery = nil
	}
	configuration := shared.SourceHeightUpdate{
		APIKey:      apiKey,
		StartDate:   startDate,
		SearchQuery: searchQuery,
	}
	out := shared.SourceHeightPutRequest{
		Name:          name,
		WorkspaceID:   workspaceID,
		Configuration: configuration,
	}
	return &out
}
