// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationFirestoreResourceModel) ToSharedDestinationFirestoreCreateRequest() *shared.DestinationFirestoreCreateRequest {
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

	var projectID string
	projectID = r.Configuration.ProjectID.ValueString()

	credentialsJSON := new(string)
	if !r.Configuration.CredentialsJSON.IsUnknown() && !r.Configuration.CredentialsJSON.IsNull() {
		*credentialsJSON = r.Configuration.CredentialsJSON.ValueString()
	} else {
		credentialsJSON = nil
	}
	configuration := shared.DestinationFirestore{
		ProjectID:       projectID,
		CredentialsJSON: credentialsJSON,
	}
	out := shared.DestinationFirestoreCreateRequest{
		Name:          name,
		DefinitionID:  definitionID,
		WorkspaceID:   workspaceID,
		Configuration: configuration,
	}
	return &out
}

func (r *DestinationFirestoreResourceModel) RefreshFromSharedDestinationResponse(resp *shared.DestinationResponse) {
	if resp != nil {
		r.CreatedAt = types.Int64Value(resp.CreatedAt)
		r.DefinitionID = types.StringValue(resp.DefinitionID)
		r.DestinationID = types.StringValue(resp.DestinationID)
		r.DestinationType = types.StringValue(resp.DestinationType)
		r.Name = types.StringValue(resp.Name)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *DestinationFirestoreResourceModel) ToSharedDestinationFirestorePutRequest() *shared.DestinationFirestorePutRequest {
	var name string
	name = r.Name.ValueString()

	var workspaceID string
	workspaceID = r.WorkspaceID.ValueString()

	var projectID string
	projectID = r.Configuration.ProjectID.ValueString()

	credentialsJSON := new(string)
	if !r.Configuration.CredentialsJSON.IsUnknown() && !r.Configuration.CredentialsJSON.IsNull() {
		*credentialsJSON = r.Configuration.CredentialsJSON.ValueString()
	} else {
		credentialsJSON = nil
	}
	configuration := shared.DestinationFirestoreUpdate{
		ProjectID:       projectID,
		CredentialsJSON: credentialsJSON,
	}
	out := shared.DestinationFirestorePutRequest{
		Name:          name,
		WorkspaceID:   workspaceID,
		Configuration: configuration,
	}
	return &out
}
