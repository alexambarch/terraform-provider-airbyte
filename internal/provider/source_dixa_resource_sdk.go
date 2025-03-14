// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceDixaResourceModel) ToSharedSourceDixaCreateRequest() *shared.SourceDixaCreateRequest {
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

	batchSize := new(int64)
	if !r.Configuration.BatchSize.IsUnknown() && !r.Configuration.BatchSize.IsNull() {
		*batchSize = r.Configuration.BatchSize.ValueInt64()
	} else {
		batchSize = nil
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceDixa{
		APIToken:  apiToken,
		BatchSize: batchSize,
		StartDate: startDate,
	}
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	out := shared.SourceDixaCreateRequest{
		Name:          name,
		DefinitionID:  definitionID,
		WorkspaceID:   workspaceID,
		Configuration: configuration,
		SecretID:      secretID,
	}
	return &out
}

func (r *SourceDixaResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.CreatedAt = types.Int64Value(resp.CreatedAt)
		r.DefinitionID = types.StringValue(resp.DefinitionID)
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceDixaResourceModel) ToSharedSourceDixaPutRequest() *shared.SourceDixaPutRequest {
	var name string
	name = r.Name.ValueString()

	var workspaceID string
	workspaceID = r.WorkspaceID.ValueString()

	var apiToken string
	apiToken = r.Configuration.APIToken.ValueString()

	batchSize := new(int64)
	if !r.Configuration.BatchSize.IsUnknown() && !r.Configuration.BatchSize.IsNull() {
		*batchSize = r.Configuration.BatchSize.ValueInt64()
	} else {
		batchSize = nil
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceDixaUpdate{
		APIToken:  apiToken,
		BatchSize: batchSize,
		StartDate: startDate,
	}
	out := shared.SourceDixaPutRequest{
		Name:          name,
		WorkspaceID:   workspaceID,
		Configuration: configuration,
	}
	return &out
}
