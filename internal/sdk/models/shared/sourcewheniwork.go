// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type WhenIWork string

const (
	WhenIWorkWhenIWork WhenIWork = "when-i-work"
)

func (e WhenIWork) ToPointer() *WhenIWork {
	return &e
}
func (e *WhenIWork) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "when-i-work":
		*e = WhenIWork(v)
		return nil
	default:
		return fmt.Errorf("invalid value for WhenIWork: %v", v)
	}
}

type SourceWhenIWork struct {
	// Email of your when-i-work account
	Email string `json:"email"`
	// Password for your when-i-work account
	Password   string    `json:"password"`
	sourceType WhenIWork `const:"when-i-work" json:"sourceType"`
}

func (s SourceWhenIWork) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceWhenIWork) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceWhenIWork) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *SourceWhenIWork) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *SourceWhenIWork) GetSourceType() WhenIWork {
	return WhenIWorkWhenIWork
}
