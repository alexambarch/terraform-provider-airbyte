// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourceZohoBiginCreateRequest struct {
	// Name of the source e.g. dev-mysql-instance.
	Name string `json:"name"`
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided.
	DefinitionID  *string         `json:"definitionId,omitempty"`
	WorkspaceID   string          `json:"workspaceId"`
	Configuration SourceZohoBigin `json:"configuration"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretID *string `json:"secretId,omitempty"`
}

func (o *SourceZohoBiginCreateRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceZohoBiginCreateRequest) GetDefinitionID() *string {
	if o == nil {
		return nil
	}
	return o.DefinitionID
}

func (o *SourceZohoBiginCreateRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}

func (o *SourceZohoBiginCreateRequest) GetConfiguration() SourceZohoBigin {
	if o == nil {
		return SourceZohoBigin{}
	}
	return o.Configuration
}

func (o *SourceZohoBiginCreateRequest) GetSecretID() *string {
	if o == nil {
		return nil
	}
	return o.SecretID
}
