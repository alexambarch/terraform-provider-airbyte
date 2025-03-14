// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourceShortioUpdate struct {
	DomainID string `json:"domain_id"`
	// Short.io Secret Key
	SecretKey string `json:"secret_key"`
	// UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated.
	StartDate string `json:"start_date"`
}

func (o *SourceShortioUpdate) GetDomainID() string {
	if o == nil {
		return ""
	}
	return o.DomainID
}

func (o *SourceShortioUpdate) GetSecretKey() string {
	if o == nil {
		return ""
	}
	return o.SecretKey
}

func (o *SourceShortioUpdate) GetStartDate() string {
	if o == nil {
		return ""
	}
	return o.StartDate
}
