// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourceCareQualityCommissionUpdate struct {
	// Your CQC Primary Key. See https://www.cqc.org.uk/about-us/transparency/using-cqc-data#api for steps to generate one.
	APIKey string `json:"api_key"`
}

func (o *SourceCareQualityCommissionUpdate) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}
