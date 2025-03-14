// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourceCallrailPutRequest struct {
	Name          string               `json:"name"`
	WorkspaceID   string               `json:"workspaceId"`
	Configuration SourceCallrailUpdate `json:"configuration"`
}

func (o *SourceCallrailPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceCallrailPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}

func (o *SourceCallrailPutRequest) GetConfiguration() SourceCallrailUpdate {
	if o == nil {
		return SourceCallrailUpdate{}
	}
	return o.Configuration
}
