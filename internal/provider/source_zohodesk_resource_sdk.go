// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceZohoDeskResourceModel) ToSharedSourceZohoDeskCreateRequest() *shared.SourceZohoDeskCreateRequest {
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

	var clientID string
	clientID = r.Configuration.ClientID.ValueString()

	var clientSecret string
	clientSecret = r.Configuration.ClientSecret.ValueString()

	var tokenRefreshEndpoint string
	tokenRefreshEndpoint = r.Configuration.TokenRefreshEndpoint.ValueString()

	var refreshToken string
	refreshToken = r.Configuration.RefreshToken.ValueString()

	includeCustomDomain := new(bool)
	if !r.Configuration.IncludeCustomDomain.IsUnknown() && !r.Configuration.IncludeCustomDomain.IsNull() {
		*includeCustomDomain = r.Configuration.IncludeCustomDomain.ValueBool()
	} else {
		includeCustomDomain = nil
	}
	configuration := shared.SourceZohoDesk{
		ClientID:             clientID,
		ClientSecret:         clientSecret,
		TokenRefreshEndpoint: tokenRefreshEndpoint,
		RefreshToken:         refreshToken,
		IncludeCustomDomain:  includeCustomDomain,
	}
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	out := shared.SourceZohoDeskCreateRequest{
		Name:          name,
		DefinitionID:  definitionID,
		WorkspaceID:   workspaceID,
		Configuration: configuration,
		SecretID:      secretID,
	}
	return &out
}

func (r *SourceZohoDeskResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.CreatedAt = types.Int64Value(resp.CreatedAt)
		r.DefinitionID = types.StringValue(resp.DefinitionID)
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceZohoDeskResourceModel) ToSharedSourceZohoDeskPutRequest() *shared.SourceZohoDeskPutRequest {
	var name string
	name = r.Name.ValueString()

	var workspaceID string
	workspaceID = r.WorkspaceID.ValueString()

	var clientID string
	clientID = r.Configuration.ClientID.ValueString()

	var clientSecret string
	clientSecret = r.Configuration.ClientSecret.ValueString()

	var tokenRefreshEndpoint string
	tokenRefreshEndpoint = r.Configuration.TokenRefreshEndpoint.ValueString()

	var refreshToken string
	refreshToken = r.Configuration.RefreshToken.ValueString()

	includeCustomDomain := new(bool)
	if !r.Configuration.IncludeCustomDomain.IsUnknown() && !r.Configuration.IncludeCustomDomain.IsNull() {
		*includeCustomDomain = r.Configuration.IncludeCustomDomain.ValueBool()
	} else {
		includeCustomDomain = nil
	}
	configuration := shared.SourceZohoDeskUpdate{
		ClientID:             clientID,
		ClientSecret:         clientSecret,
		TokenRefreshEndpoint: tokenRefreshEndpoint,
		RefreshToken:         refreshToken,
		IncludeCustomDomain:  includeCustomDomain,
	}
	out := shared.SourceZohoDeskPutRequest{
		Name:          name,
		WorkspaceID:   workspaceID,
		Configuration: configuration,
	}
	return &out
}
