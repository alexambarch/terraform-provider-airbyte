// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceK6CloudResourceModel) ToSharedSourceK6CloudCreateRequest() *shared.SourceK6CloudCreateRequest {
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

	configuration := shared.SourceK6Cloud{
		APIToken: apiToken,
	}
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	out := shared.SourceK6CloudCreateRequest{
		Name:          name,
		DefinitionID:  definitionID,
		WorkspaceID:   workspaceID,
		Configuration: configuration,
		SecretID:      secretID,
	}
	return &out
}

func (r *SourceK6CloudResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.CreatedAt = types.Int64Value(resp.CreatedAt)
		r.DefinitionID = types.StringValue(resp.DefinitionID)
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceK6CloudResourceModel) ToSharedSourceK6CloudPutRequest() *shared.SourceK6CloudPutRequest {
	var name string
	name = r.Name.ValueString()

	var workspaceID string
	workspaceID = r.WorkspaceID.ValueString()

	var apiToken string
	apiToken = r.Configuration.APIToken.ValueString()

	configuration := shared.SourceK6CloudUpdate{
		APIToken: apiToken,
	}
	out := shared.SourceK6CloudPutRequest{
		Name:          name,
		WorkspaceID:   workspaceID,
		Configuration: configuration,
	}
	return &out
}
