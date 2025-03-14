// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type DestinationS3DataLakeCreateRequest struct {
	// Name of the destination e.g. dev-mysql-instance.
	Name string `json:"name"`
	// The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided.
	DefinitionID *string `json:"definitionId,omitempty"`
	WorkspaceID  string  `json:"workspaceId"`
	// Defines the configurations required to connect to an Iceberg catalog, including warehouse location, main branch name, and catalog type specifics.
	Configuration DestinationS3DataLake `json:"configuration"`
}

func (o *DestinationS3DataLakeCreateRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *DestinationS3DataLakeCreateRequest) GetDefinitionID() *string {
	if o == nil {
		return nil
	}
	return o.DefinitionID
}

func (o *DestinationS3DataLakeCreateRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}

func (o *DestinationS3DataLakeCreateRequest) GetConfiguration() DestinationS3DataLake {
	if o == nil {
		return DestinationS3DataLake{}
	}
	return o.Configuration
}
