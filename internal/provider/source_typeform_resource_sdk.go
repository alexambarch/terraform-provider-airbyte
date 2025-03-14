// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceTypeformResourceModel) ToSharedSourceTypeformCreateRequest() *shared.SourceTypeformCreateRequest {
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

	var credentials shared.SourceTypeformAuthorizationMethod
	var sourceTypeformOAuth20 *shared.SourceTypeformOAuth20
	if r.Configuration.Credentials.OAuth20 != nil {
		var clientID string
		clientID = r.Configuration.Credentials.OAuth20.ClientID.ValueString()

		var clientSecret string
		clientSecret = r.Configuration.Credentials.OAuth20.ClientSecret.ValueString()

		var accessToken string
		accessToken = r.Configuration.Credentials.OAuth20.AccessToken.ValueString()

		tokenExpiryDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.Credentials.OAuth20.TokenExpiryDate.ValueString())
		var refreshToken string
		refreshToken = r.Configuration.Credentials.OAuth20.RefreshToken.ValueString()

		sourceTypeformOAuth20 = &shared.SourceTypeformOAuth20{
			ClientID:        clientID,
			ClientSecret:    clientSecret,
			AccessToken:     accessToken,
			TokenExpiryDate: tokenExpiryDate,
			RefreshToken:    refreshToken,
		}
	}
	if sourceTypeformOAuth20 != nil {
		credentials = shared.SourceTypeformAuthorizationMethod{
			SourceTypeformOAuth20: sourceTypeformOAuth20,
		}
	}
	var sourceTypeformPrivateToken *shared.SourceTypeformPrivateToken
	if r.Configuration.Credentials.PrivateToken != nil {
		var accessToken1 string
		accessToken1 = r.Configuration.Credentials.PrivateToken.AccessToken.ValueString()

		sourceTypeformPrivateToken = &shared.SourceTypeformPrivateToken{
			AccessToken: accessToken1,
		}
	}
	if sourceTypeformPrivateToken != nil {
		credentials = shared.SourceTypeformAuthorizationMethod{
			SourceTypeformPrivateToken: sourceTypeformPrivateToken,
		}
	}
	startDate := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	var formIds []string = []string{}
	for _, formIdsItem := range r.Configuration.FormIds {
		formIds = append(formIds, formIdsItem.ValueString())
	}
	configuration := shared.SourceTypeform{
		Credentials: credentials,
		StartDate:   startDate,
		FormIds:     formIds,
	}
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	out := shared.SourceTypeformCreateRequest{
		Name:          name,
		DefinitionID:  definitionID,
		WorkspaceID:   workspaceID,
		Configuration: configuration,
		SecretID:      secretID,
	}
	return &out
}

func (r *SourceTypeformResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.CreatedAt = types.Int64Value(resp.CreatedAt)
		r.DefinitionID = types.StringValue(resp.DefinitionID)
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceTypeformResourceModel) ToSharedSourceTypeformPutRequest() *shared.SourceTypeformPutRequest {
	var name string
	name = r.Name.ValueString()

	var workspaceID string
	workspaceID = r.WorkspaceID.ValueString()

	var credentials shared.SourceTypeformUpdateAuthorizationMethod
	var sourceTypeformUpdateOAuth20 *shared.SourceTypeformUpdateOAuth20
	if r.Configuration.Credentials.OAuth20 != nil {
		var clientID string
		clientID = r.Configuration.Credentials.OAuth20.ClientID.ValueString()

		var clientSecret string
		clientSecret = r.Configuration.Credentials.OAuth20.ClientSecret.ValueString()

		var accessToken string
		accessToken = r.Configuration.Credentials.OAuth20.AccessToken.ValueString()

		tokenExpiryDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.Credentials.OAuth20.TokenExpiryDate.ValueString())
		var refreshToken string
		refreshToken = r.Configuration.Credentials.OAuth20.RefreshToken.ValueString()

		sourceTypeformUpdateOAuth20 = &shared.SourceTypeformUpdateOAuth20{
			ClientID:        clientID,
			ClientSecret:    clientSecret,
			AccessToken:     accessToken,
			TokenExpiryDate: tokenExpiryDate,
			RefreshToken:    refreshToken,
		}
	}
	if sourceTypeformUpdateOAuth20 != nil {
		credentials = shared.SourceTypeformUpdateAuthorizationMethod{
			SourceTypeformUpdateOAuth20: sourceTypeformUpdateOAuth20,
		}
	}
	var sourceTypeformUpdatePrivateToken *shared.SourceTypeformUpdatePrivateToken
	if r.Configuration.Credentials.PrivateToken != nil {
		var accessToken1 string
		accessToken1 = r.Configuration.Credentials.PrivateToken.AccessToken.ValueString()

		sourceTypeformUpdatePrivateToken = &shared.SourceTypeformUpdatePrivateToken{
			AccessToken: accessToken1,
		}
	}
	if sourceTypeformUpdatePrivateToken != nil {
		credentials = shared.SourceTypeformUpdateAuthorizationMethod{
			SourceTypeformUpdatePrivateToken: sourceTypeformUpdatePrivateToken,
		}
	}
	startDate := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	var formIds []string = []string{}
	for _, formIdsItem := range r.Configuration.FormIds {
		formIds = append(formIds, formIdsItem.ValueString())
	}
	configuration := shared.SourceTypeformUpdate{
		Credentials: credentials,
		StartDate:   startDate,
		FormIds:     formIds,
	}
	out := shared.SourceTypeformPutRequest{
		Name:          name,
		WorkspaceID:   workspaceID,
		Configuration: configuration,
	}
	return &out
}
