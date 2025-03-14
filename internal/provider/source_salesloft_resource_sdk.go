// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceSalesloftResourceModel) ToSharedSourceSalesloftCreateRequest() *shared.SourceSalesloftCreateRequest {
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

	var credentials shared.SourceSalesloftCredentials
	var authenticateViaOAuth *shared.AuthenticateViaOAuth
	if r.Configuration.Credentials.AuthenticateViaOAuth != nil {
		var clientID string
		clientID = r.Configuration.Credentials.AuthenticateViaOAuth.ClientID.ValueString()

		var accessToken string
		accessToken = r.Configuration.Credentials.AuthenticateViaOAuth.AccessToken.ValueString()

		var clientSecret string
		clientSecret = r.Configuration.Credentials.AuthenticateViaOAuth.ClientSecret.ValueString()

		var refreshToken string
		refreshToken = r.Configuration.Credentials.AuthenticateViaOAuth.RefreshToken.ValueString()

		tokenExpiryDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.Credentials.AuthenticateViaOAuth.TokenExpiryDate.ValueString())
		authenticateViaOAuth = &shared.AuthenticateViaOAuth{
			ClientID:        clientID,
			AccessToken:     accessToken,
			ClientSecret:    clientSecret,
			RefreshToken:    refreshToken,
			TokenExpiryDate: tokenExpiryDate,
		}
	}
	if authenticateViaOAuth != nil {
		credentials = shared.SourceSalesloftCredentials{
			AuthenticateViaOAuth: authenticateViaOAuth,
		}
	}
	var authenticateViaAPIKey *shared.AuthenticateViaAPIKey
	if r.Configuration.Credentials.AuthenticateViaAPIKey != nil {
		var apiKey string
		apiKey = r.Configuration.Credentials.AuthenticateViaAPIKey.APIKey.ValueString()

		authenticateViaAPIKey = &shared.AuthenticateViaAPIKey{
			APIKey: apiKey,
		}
	}
	if authenticateViaAPIKey != nil {
		credentials = shared.SourceSalesloftCredentials{
			AuthenticateViaAPIKey: authenticateViaAPIKey,
		}
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceSalesloft{
		Credentials: credentials,
		StartDate:   startDate,
	}
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	out := shared.SourceSalesloftCreateRequest{
		Name:          name,
		DefinitionID:  definitionID,
		WorkspaceID:   workspaceID,
		Configuration: configuration,
		SecretID:      secretID,
	}
	return &out
}

func (r *SourceSalesloftResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.CreatedAt = types.Int64Value(resp.CreatedAt)
		r.DefinitionID = types.StringValue(resp.DefinitionID)
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceSalesloftResourceModel) ToSharedSourceSalesloftPutRequest() *shared.SourceSalesloftPutRequest {
	var name string
	name = r.Name.ValueString()

	var workspaceID string
	workspaceID = r.WorkspaceID.ValueString()

	var credentials shared.SourceSalesloftUpdateCredentials
	var sourceSalesloftUpdateAuthenticateViaOAuth *shared.SourceSalesloftUpdateAuthenticateViaOAuth
	if r.Configuration.Credentials.AuthenticateViaOAuth != nil {
		var clientID string
		clientID = r.Configuration.Credentials.AuthenticateViaOAuth.ClientID.ValueString()

		var accessToken string
		accessToken = r.Configuration.Credentials.AuthenticateViaOAuth.AccessToken.ValueString()

		var clientSecret string
		clientSecret = r.Configuration.Credentials.AuthenticateViaOAuth.ClientSecret.ValueString()

		var refreshToken string
		refreshToken = r.Configuration.Credentials.AuthenticateViaOAuth.RefreshToken.ValueString()

		tokenExpiryDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.Credentials.AuthenticateViaOAuth.TokenExpiryDate.ValueString())
		sourceSalesloftUpdateAuthenticateViaOAuth = &shared.SourceSalesloftUpdateAuthenticateViaOAuth{
			ClientID:        clientID,
			AccessToken:     accessToken,
			ClientSecret:    clientSecret,
			RefreshToken:    refreshToken,
			TokenExpiryDate: tokenExpiryDate,
		}
	}
	if sourceSalesloftUpdateAuthenticateViaOAuth != nil {
		credentials = shared.SourceSalesloftUpdateCredentials{
			SourceSalesloftUpdateAuthenticateViaOAuth: sourceSalesloftUpdateAuthenticateViaOAuth,
		}
	}
	var sourceSalesloftUpdateAuthenticateViaAPIKey *shared.SourceSalesloftUpdateAuthenticateViaAPIKey
	if r.Configuration.Credentials.AuthenticateViaAPIKey != nil {
		var apiKey string
		apiKey = r.Configuration.Credentials.AuthenticateViaAPIKey.APIKey.ValueString()

		sourceSalesloftUpdateAuthenticateViaAPIKey = &shared.SourceSalesloftUpdateAuthenticateViaAPIKey{
			APIKey: apiKey,
		}
	}
	if sourceSalesloftUpdateAuthenticateViaAPIKey != nil {
		credentials = shared.SourceSalesloftUpdateCredentials{
			SourceSalesloftUpdateAuthenticateViaAPIKey: sourceSalesloftUpdateAuthenticateViaAPIKey,
		}
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceSalesloftUpdate{
		Credentials: credentials,
		StartDate:   startDate,
	}
	out := shared.SourceSalesloftPutRequest{
		Name:          name,
		WorkspaceID:   workspaceID,
		Configuration: configuration,
	}
	return &out
}
