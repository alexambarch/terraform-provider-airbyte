// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceSigmaComputingResourceModel) ToSharedSourceSigmaComputingCreateRequest() *shared.SourceSigmaComputingCreateRequest {
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

	var clientRefreshToken string
	clientRefreshToken = r.Configuration.ClientRefreshToken.ValueString()

	oauthAccessToken := new(string)
	if !r.Configuration.OauthAccessToken.IsUnknown() && !r.Configuration.OauthAccessToken.IsNull() {
		*oauthAccessToken = r.Configuration.OauthAccessToken.ValueString()
	} else {
		oauthAccessToken = nil
	}
	oauthTokenExpiryDate := new(time.Time)
	if !r.Configuration.OauthTokenExpiryDate.IsUnknown() && !r.Configuration.OauthTokenExpiryDate.IsNull() {
		*oauthTokenExpiryDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.OauthTokenExpiryDate.ValueString())
	} else {
		oauthTokenExpiryDate = nil
	}
	var baseURL string
	baseURL = r.Configuration.BaseURL.ValueString()

	configuration := shared.SourceSigmaComputing{
		ClientID:             clientID,
		ClientSecret:         clientSecret,
		ClientRefreshToken:   clientRefreshToken,
		OauthAccessToken:     oauthAccessToken,
		OauthTokenExpiryDate: oauthTokenExpiryDate,
		BaseURL:              baseURL,
	}
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	out := shared.SourceSigmaComputingCreateRequest{
		Name:          name,
		DefinitionID:  definitionID,
		WorkspaceID:   workspaceID,
		Configuration: configuration,
		SecretID:      secretID,
	}
	return &out
}

func (r *SourceSigmaComputingResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.CreatedAt = types.Int64Value(resp.CreatedAt)
		r.DefinitionID = types.StringValue(resp.DefinitionID)
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceSigmaComputingResourceModel) ToSharedSourceSigmaComputingPutRequest() *shared.SourceSigmaComputingPutRequest {
	var name string
	name = r.Name.ValueString()

	var workspaceID string
	workspaceID = r.WorkspaceID.ValueString()

	var clientID string
	clientID = r.Configuration.ClientID.ValueString()

	var clientSecret string
	clientSecret = r.Configuration.ClientSecret.ValueString()

	var clientRefreshToken string
	clientRefreshToken = r.Configuration.ClientRefreshToken.ValueString()

	oauthAccessToken := new(string)
	if !r.Configuration.OauthAccessToken.IsUnknown() && !r.Configuration.OauthAccessToken.IsNull() {
		*oauthAccessToken = r.Configuration.OauthAccessToken.ValueString()
	} else {
		oauthAccessToken = nil
	}
	oauthTokenExpiryDate := new(time.Time)
	if !r.Configuration.OauthTokenExpiryDate.IsUnknown() && !r.Configuration.OauthTokenExpiryDate.IsNull() {
		*oauthTokenExpiryDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.OauthTokenExpiryDate.ValueString())
	} else {
		oauthTokenExpiryDate = nil
	}
	var baseURL string
	baseURL = r.Configuration.BaseURL.ValueString()

	configuration := shared.SourceSigmaComputingUpdate{
		ClientID:             clientID,
		ClientSecret:         clientSecret,
		ClientRefreshToken:   clientRefreshToken,
		OauthAccessToken:     oauthAccessToken,
		OauthTokenExpiryDate: oauthTokenExpiryDate,
		BaseURL:              baseURL,
	}
	out := shared.SourceSigmaComputingPutRequest{
		Name:          name,
		WorkspaceID:   workspaceID,
		Configuration: configuration,
	}
	return &out
}
