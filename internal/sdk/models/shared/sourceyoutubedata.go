// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type YoutubeData string

const (
	YoutubeDataYoutubeData YoutubeData = "youtube-data"
)

func (e YoutubeData) ToPointer() *YoutubeData {
	return &e
}
func (e *YoutubeData) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "youtube-data":
		*e = YoutubeData(v)
		return nil
	default:
		return fmt.Errorf("invalid value for YoutubeData: %v", v)
	}
}

type SourceYoutubeData struct {
	APIKey     string      `json:"api_key"`
	ChannelIds []any       `json:"channel_ids"`
	sourceType YoutubeData `const:"youtube-data" json:"sourceType"`
}

func (s SourceYoutubeData) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceYoutubeData) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceYoutubeData) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceYoutubeData) GetChannelIds() []any {
	if o == nil {
		return []any{}
	}
	return o.ChannelIds
}

func (o *SourceYoutubeData) GetSourceType() YoutubeData {
	return YoutubeDataYoutubeData
}
