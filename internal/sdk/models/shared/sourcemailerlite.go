// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type Mailerlite string

const (
	MailerliteMailerlite Mailerlite = "mailerlite"
)

func (e Mailerlite) ToPointer() *Mailerlite {
	return &e
}
func (e *Mailerlite) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "mailerlite":
		*e = Mailerlite(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Mailerlite: %v", v)
	}
}

type SourceMailerlite struct {
	// Your API Token. See <a href="https://developers.mailerlite.com/docs/#authentication">here</a>.
	APIToken   string     `json:"api_token"`
	sourceType Mailerlite `const:"mailerlite" json:"sourceType"`
}

func (s SourceMailerlite) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceMailerlite) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceMailerlite) GetAPIToken() string {
	if o == nil {
		return ""
	}
	return o.APIToken
}

func (o *SourceMailerlite) GetSourceType() Mailerlite {
	return MailerliteMailerlite
}
