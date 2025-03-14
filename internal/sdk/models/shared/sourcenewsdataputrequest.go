// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourceNewsdataPutRequest struct {
	Name          string               `json:"name"`
	WorkspaceID   string               `json:"workspaceId"`
	Configuration SourceNewsdataUpdate `json:"configuration"`
}

func (o *SourceNewsdataPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceNewsdataPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}

func (o *SourceNewsdataPutRequest) GetConfiguration() SourceNewsdataUpdate {
	if o == nil {
		return SourceNewsdataUpdate{}
	}
	return o.Configuration
}
