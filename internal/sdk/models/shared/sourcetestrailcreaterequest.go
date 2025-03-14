// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourceTestrailCreateRequest struct {
	// Name of the source e.g. dev-mysql-instance.
	Name string `json:"name"`
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided.
	DefinitionID  *string        `json:"definitionId,omitempty"`
	WorkspaceID   string         `json:"workspaceId"`
	Configuration SourceTestrail `json:"configuration"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretID *string `json:"secretId,omitempty"`
}

func (o *SourceTestrailCreateRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceTestrailCreateRequest) GetDefinitionID() *string {
	if o == nil {
		return nil
	}
	return o.DefinitionID
}

func (o *SourceTestrailCreateRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}

func (o *SourceTestrailCreateRequest) GetConfiguration() SourceTestrail {
	if o == nil {
		return SourceTestrail{}
	}
	return o.Configuration
}

func (o *SourceTestrailCreateRequest) GetSecretID() *string {
	if o == nil {
		return nil
	}
	return o.SecretID
}
