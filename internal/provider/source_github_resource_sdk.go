// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceGithubResourceModel) ToSharedSourceGithubCreateRequest() *shared.SourceGithubCreateRequest {
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

	var credentials shared.SourceGithubAuthentication
	var oAuth *shared.OAuth
	if r.Configuration.Credentials.OAuth != nil {
		var accessToken string
		accessToken = r.Configuration.Credentials.OAuth.AccessToken.ValueString()

		clientID := new(string)
		if !r.Configuration.Credentials.OAuth.ClientID.IsUnknown() && !r.Configuration.Credentials.OAuth.ClientID.IsNull() {
			*clientID = r.Configuration.Credentials.OAuth.ClientID.ValueString()
		} else {
			clientID = nil
		}
		clientSecret := new(string)
		if !r.Configuration.Credentials.OAuth.ClientSecret.IsUnknown() && !r.Configuration.Credentials.OAuth.ClientSecret.IsNull() {
			*clientSecret = r.Configuration.Credentials.OAuth.ClientSecret.ValueString()
		} else {
			clientSecret = nil
		}
		oAuth = &shared.OAuth{
			AccessToken:  accessToken,
			ClientID:     clientID,
			ClientSecret: clientSecret,
		}
	}
	if oAuth != nil {
		credentials = shared.SourceGithubAuthentication{
			OAuth: oAuth,
		}
	}
	var sourceGithubPersonalAccessToken *shared.SourceGithubPersonalAccessToken
	if r.Configuration.Credentials.PersonalAccessToken != nil {
		var personalAccessToken string
		personalAccessToken = r.Configuration.Credentials.PersonalAccessToken.PersonalAccessToken.ValueString()

		sourceGithubPersonalAccessToken = &shared.SourceGithubPersonalAccessToken{
			PersonalAccessToken: personalAccessToken,
		}
	}
	if sourceGithubPersonalAccessToken != nil {
		credentials = shared.SourceGithubAuthentication{
			SourceGithubPersonalAccessToken: sourceGithubPersonalAccessToken,
		}
	}
	var repositories []string = []string{}
	for _, repositoriesItem := range r.Configuration.Repositories {
		repositories = append(repositories, repositoriesItem.ValueString())
	}
	startDate := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	apiURL := new(string)
	if !r.Configuration.APIURL.IsUnknown() && !r.Configuration.APIURL.IsNull() {
		*apiURL = r.Configuration.APIURL.ValueString()
	} else {
		apiURL = nil
	}
	var branches []string = []string{}
	for _, branchesItem := range r.Configuration.Branches {
		branches = append(branches, branchesItem.ValueString())
	}
	maxWaitingTime := new(int64)
	if !r.Configuration.MaxWaitingTime.IsUnknown() && !r.Configuration.MaxWaitingTime.IsNull() {
		*maxWaitingTime = r.Configuration.MaxWaitingTime.ValueInt64()
	} else {
		maxWaitingTime = nil
	}
	configuration := shared.SourceGithub{
		Credentials:    credentials,
		Repositories:   repositories,
		StartDate:      startDate,
		APIURL:         apiURL,
		Branches:       branches,
		MaxWaitingTime: maxWaitingTime,
	}
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	out := shared.SourceGithubCreateRequest{
		Name:          name,
		DefinitionID:  definitionID,
		WorkspaceID:   workspaceID,
		Configuration: configuration,
		SecretID:      secretID,
	}
	return &out
}

func (r *SourceGithubResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.CreatedAt = types.Int64Value(resp.CreatedAt)
		r.DefinitionID = types.StringValue(resp.DefinitionID)
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceGithubResourceModel) ToSharedSourceGithubPutRequest() *shared.SourceGithubPutRequest {
	var name string
	name = r.Name.ValueString()

	var workspaceID string
	workspaceID = r.WorkspaceID.ValueString()

	var credentials shared.SourceGithubUpdateAuthentication
	var sourceGithubUpdateOAuth *shared.SourceGithubUpdateOAuth
	if r.Configuration.Credentials.OAuth != nil {
		var accessToken string
		accessToken = r.Configuration.Credentials.OAuth.AccessToken.ValueString()

		clientID := new(string)
		if !r.Configuration.Credentials.OAuth.ClientID.IsUnknown() && !r.Configuration.Credentials.OAuth.ClientID.IsNull() {
			*clientID = r.Configuration.Credentials.OAuth.ClientID.ValueString()
		} else {
			clientID = nil
		}
		clientSecret := new(string)
		if !r.Configuration.Credentials.OAuth.ClientSecret.IsUnknown() && !r.Configuration.Credentials.OAuth.ClientSecret.IsNull() {
			*clientSecret = r.Configuration.Credentials.OAuth.ClientSecret.ValueString()
		} else {
			clientSecret = nil
		}
		sourceGithubUpdateOAuth = &shared.SourceGithubUpdateOAuth{
			AccessToken:  accessToken,
			ClientID:     clientID,
			ClientSecret: clientSecret,
		}
	}
	if sourceGithubUpdateOAuth != nil {
		credentials = shared.SourceGithubUpdateAuthentication{
			SourceGithubUpdateOAuth: sourceGithubUpdateOAuth,
		}
	}
	var sourceGithubUpdatePersonalAccessToken *shared.SourceGithubUpdatePersonalAccessToken
	if r.Configuration.Credentials.PersonalAccessToken != nil {
		var personalAccessToken string
		personalAccessToken = r.Configuration.Credentials.PersonalAccessToken.PersonalAccessToken.ValueString()

		sourceGithubUpdatePersonalAccessToken = &shared.SourceGithubUpdatePersonalAccessToken{
			PersonalAccessToken: personalAccessToken,
		}
	}
	if sourceGithubUpdatePersonalAccessToken != nil {
		credentials = shared.SourceGithubUpdateAuthentication{
			SourceGithubUpdatePersonalAccessToken: sourceGithubUpdatePersonalAccessToken,
		}
	}
	var repositories []string = []string{}
	for _, repositoriesItem := range r.Configuration.Repositories {
		repositories = append(repositories, repositoriesItem.ValueString())
	}
	startDate := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	apiURL := new(string)
	if !r.Configuration.APIURL.IsUnknown() && !r.Configuration.APIURL.IsNull() {
		*apiURL = r.Configuration.APIURL.ValueString()
	} else {
		apiURL = nil
	}
	var branches []string = []string{}
	for _, branchesItem := range r.Configuration.Branches {
		branches = append(branches, branchesItem.ValueString())
	}
	maxWaitingTime := new(int64)
	if !r.Configuration.MaxWaitingTime.IsUnknown() && !r.Configuration.MaxWaitingTime.IsNull() {
		*maxWaitingTime = r.Configuration.MaxWaitingTime.ValueInt64()
	} else {
		maxWaitingTime = nil
	}
	configuration := shared.SourceGithubUpdate{
		Credentials:    credentials,
		Repositories:   repositories,
		StartDate:      startDate,
		APIURL:         apiURL,
		Branches:       branches,
		MaxWaitingTime: maxWaitingTime,
	}
	out := shared.SourceGithubPutRequest{
		Name:          name,
		WorkspaceID:   workspaceID,
		Configuration: configuration,
	}
	return &out
}
