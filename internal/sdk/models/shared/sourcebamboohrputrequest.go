// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourceBambooHrPutRequest struct {
	Name          string               `json:"name"`
	WorkspaceID   string               `json:"workspaceId"`
	Configuration SourceBambooHrUpdate `json:"configuration"`
}

func (o *SourceBambooHrPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceBambooHrPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}

func (o *SourceBambooHrPutRequest) GetConfiguration() SourceBambooHrUpdate {
	if o == nil {
		return SourceBambooHrUpdate{}
	}
	return o.Configuration
}
