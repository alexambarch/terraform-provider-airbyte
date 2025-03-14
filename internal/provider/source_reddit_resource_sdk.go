// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceRedditResourceModel) ToSharedSourceRedditCreateRequest() *shared.SourceRedditCreateRequest {
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

	query := new(string)
	if !r.Configuration.Query.IsUnknown() && !r.Configuration.Query.IsNull() {
		*query = r.Configuration.Query.ValueString()
	} else {
		query = nil
	}
	includeOver18 := new(bool)
	if !r.Configuration.IncludeOver18.IsUnknown() && !r.Configuration.IncludeOver18.IsNull() {
		*includeOver18 = r.Configuration.IncludeOver18.ValueBool()
	} else {
		includeOver18 = nil
	}
	exact := new(bool)
	if !r.Configuration.Exact.IsUnknown() && !r.Configuration.Exact.IsNull() {
		*exact = r.Configuration.Exact.ValueBool()
	} else {
		exact = nil
	}
	limit := new(float64)
	if !r.Configuration.Limit.IsUnknown() && !r.Configuration.Limit.IsNull() {
		*limit, _ = r.Configuration.Limit.ValueBigFloat().Float64()
	} else {
		limit = nil
	}
	var subreddits []interface{} = []interface{}{}
	for _, subredditsItem := range r.Configuration.Subreddits {
		var subredditsTmp interface{}
		_ = json.Unmarshal([]byte(subredditsItem.ValueString()), &subredditsTmp)
		subreddits = append(subreddits, subredditsTmp)
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceReddit{
		APIKey:        apiKey,
		Query:         query,
		IncludeOver18: includeOver18,
		Exact:         exact,
		Limit:         limit,
		Subreddits:    subreddits,
		StartDate:     startDate,
	}
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	out := shared.SourceRedditCreateRequest{
		Name:          name,
		DefinitionID:  definitionID,
		WorkspaceID:   workspaceID,
		Configuration: configuration,
		SecretID:      secretID,
	}
	return &out
}

func (r *SourceRedditResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.CreatedAt = types.Int64Value(resp.CreatedAt)
		r.DefinitionID = types.StringValue(resp.DefinitionID)
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceRedditResourceModel) ToSharedSourceRedditPutRequest() *shared.SourceRedditPutRequest {
	var name string
	name = r.Name.ValueString()

	var workspaceID string
	workspaceID = r.WorkspaceID.ValueString()

	var apiKey string
	apiKey = r.Configuration.APIKey.ValueString()

	query := new(string)
	if !r.Configuration.Query.IsUnknown() && !r.Configuration.Query.IsNull() {
		*query = r.Configuration.Query.ValueString()
	} else {
		query = nil
	}
	includeOver18 := new(bool)
	if !r.Configuration.IncludeOver18.IsUnknown() && !r.Configuration.IncludeOver18.IsNull() {
		*includeOver18 = r.Configuration.IncludeOver18.ValueBool()
	} else {
		includeOver18 = nil
	}
	exact := new(bool)
	if !r.Configuration.Exact.IsUnknown() && !r.Configuration.Exact.IsNull() {
		*exact = r.Configuration.Exact.ValueBool()
	} else {
		exact = nil
	}
	limit := new(float64)
	if !r.Configuration.Limit.IsUnknown() && !r.Configuration.Limit.IsNull() {
		*limit, _ = r.Configuration.Limit.ValueBigFloat().Float64()
	} else {
		limit = nil
	}
	var subreddits []interface{} = []interface{}{}
	for _, subredditsItem := range r.Configuration.Subreddits {
		var subredditsTmp interface{}
		_ = json.Unmarshal([]byte(subredditsItem.ValueString()), &subredditsTmp)
		subreddits = append(subreddits, subredditsTmp)
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceRedditUpdate{
		APIKey:        apiKey,
		Query:         query,
		IncludeOver18: includeOver18,
		Exact:         exact,
		Limit:         limit,
		Subreddits:    subreddits,
		StartDate:     startDate,
	}
	out := shared.SourceRedditPutRequest{
		Name:          name,
		WorkspaceID:   workspaceID,
		Configuration: configuration,
	}
	return &out
}
