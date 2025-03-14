// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type Lokalise string

const (
	LokaliseLokalise Lokalise = "lokalise"
)

func (e Lokalise) ToPointer() *Lokalise {
	return &e
}
func (e *Lokalise) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "lokalise":
		*e = Lokalise(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Lokalise: %v", v)
	}
}

type SourceLokalise struct {
	// Lokalise API Key with read-access. Available at Profile settings > API tokens. See <a href="https://docs.lokalise.com/en/articles/1929556-api-tokens">here</a>.
	APIKey string `json:"api_key"`
	// Lokalise project ID. Available at Project Settings > General.
	ProjectID  string   `json:"project_id"`
	sourceType Lokalise `const:"lokalise" json:"sourceType"`
}

func (s SourceLokalise) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceLokalise) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceLokalise) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceLokalise) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

func (o *SourceLokalise) GetSourceType() Lokalise {
	return LokaliseLokalise
}
