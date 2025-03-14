// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceClickupAPIResourceModel) ToSharedSourceClickupAPICreateRequest() *shared.SourceClickupAPICreateRequest {
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

	var apiToken string
	apiToken = r.Configuration.APIToken.ValueString()

	includeClosedTasks := new(bool)
	if !r.Configuration.IncludeClosedTasks.IsUnknown() && !r.Configuration.IncludeClosedTasks.IsNull() {
		*includeClosedTasks = r.Configuration.IncludeClosedTasks.ValueBool()
	} else {
		includeClosedTasks = nil
	}
	configuration := shared.SourceClickupAPI{
		APIToken:           apiToken,
		IncludeClosedTasks: includeClosedTasks,
	}
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	out := shared.SourceClickupAPICreateRequest{
		Name:          name,
		DefinitionID:  definitionID,
		WorkspaceID:   workspaceID,
		Configuration: configuration,
		SecretID:      secretID,
	}
	return &out
}

func (r *SourceClickupAPIResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.CreatedAt = types.Int64Value(resp.CreatedAt)
		r.DefinitionID = types.StringValue(resp.DefinitionID)
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceClickupAPIResourceModel) ToSharedSourceClickupAPIPutRequest() *shared.SourceClickupAPIPutRequest {
	var name string
	name = r.Name.ValueString()

	var workspaceID string
	workspaceID = r.WorkspaceID.ValueString()

	var apiToken string
	apiToken = r.Configuration.APIToken.ValueString()

	includeClosedTasks := new(bool)
	if !r.Configuration.IncludeClosedTasks.IsUnknown() && !r.Configuration.IncludeClosedTasks.IsNull() {
		*includeClosedTasks = r.Configuration.IncludeClosedTasks.ValueBool()
	} else {
		includeClosedTasks = nil
	}
	configuration := shared.SourceClickupAPIUpdate{
		APIToken:           apiToken,
		IncludeClosedTasks: includeClosedTasks,
	}
	out := shared.SourceClickupAPIPutRequest{
		Name:          name,
		WorkspaceID:   workspaceID,
		Configuration: configuration,
	}
	return &out
}
