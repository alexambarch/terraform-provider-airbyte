// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type Glassfrog string

const (
	GlassfrogGlassfrog Glassfrog = "glassfrog"
)

func (e Glassfrog) ToPointer() *Glassfrog {
	return &e
}
func (e *Glassfrog) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "glassfrog":
		*e = Glassfrog(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Glassfrog: %v", v)
	}
}

type SourceGlassfrog struct {
	// API key provided by Glassfrog
	APIKey     string    `json:"api_key"`
	sourceType Glassfrog `const:"glassfrog" json:"sourceType"`
}

func (s SourceGlassfrog) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceGlassfrog) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceGlassfrog) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceGlassfrog) GetSourceType() Glassfrog {
	return GlassfrogGlassfrog
}
