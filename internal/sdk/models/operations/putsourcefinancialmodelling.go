// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"net/http"
)

type PutSourceFinancialModellingRequest struct {
	SourceID                           string                                     `pathParam:"style=simple,explode=false,name=sourceId"`
	SourceFinancialModellingPutRequest *shared.SourceFinancialModellingPutRequest `request:"mediaType=application/json"`
}

func (o *PutSourceFinancialModellingRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

func (o *PutSourceFinancialModellingRequest) GetSourceFinancialModellingPutRequest() *shared.SourceFinancialModellingPutRequest {
	if o == nil {
		return nil
	}
	return o.SourceFinancialModellingPutRequest
}

type PutSourceFinancialModellingResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSourceFinancialModellingResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSourceFinancialModellingResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSourceFinancialModellingResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
