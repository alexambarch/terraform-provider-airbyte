// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourceRecruiteePutRequest struct {
	Name          string                `json:"name"`
	WorkspaceID   string                `json:"workspaceId"`
	Configuration SourceRecruiteeUpdate `json:"configuration"`
}

func (o *SourceRecruiteePutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceRecruiteePutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}

func (o *SourceRecruiteePutRequest) GetConfiguration() SourceRecruiteeUpdate {
	if o == nil {
		return SourceRecruiteeUpdate{}
	}
	return o.Configuration
}
