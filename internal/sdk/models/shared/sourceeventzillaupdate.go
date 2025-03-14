// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourceEventzillaUpdate struct {
	// API key to use. Generate it by creating a new application within your Eventzilla account settings under Settings > App Management.
	XAPIKey string `json:"x-api-key"`
}

func (o *SourceEventzillaUpdate) GetXAPIKey() string {
	if o == nil {
		return ""
	}
	return o.XAPIKey
}
