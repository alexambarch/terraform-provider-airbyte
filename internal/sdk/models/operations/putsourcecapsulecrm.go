// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"net/http"
)

type PutSourceCapsuleCrmRequest struct {
	SourceID                   string                             `pathParam:"style=simple,explode=false,name=sourceId"`
	SourceCapsuleCrmPutRequest *shared.SourceCapsuleCrmPutRequest `request:"mediaType=application/json"`
}

func (o *PutSourceCapsuleCrmRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

func (o *PutSourceCapsuleCrmRequest) GetSourceCapsuleCrmPutRequest() *shared.SourceCapsuleCrmPutRequest {
	if o == nil {
		return nil
	}
	return o.SourceCapsuleCrmPutRequest
}

type PutSourceCapsuleCrmResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSourceCapsuleCrmResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSourceCapsuleCrmResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSourceCapsuleCrmResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
