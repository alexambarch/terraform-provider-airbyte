// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourceOmnisendUpdate struct {
	// API Key
	APIKey string `json:"api_key"`
}

func (o *SourceOmnisendUpdate) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}
