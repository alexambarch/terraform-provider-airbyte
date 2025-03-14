// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"net/http"
)

type PutSourceCaptainDataRequest struct {
	SourceID                    string                              `pathParam:"style=simple,explode=false,name=sourceId"`
	SourceCaptainDataPutRequest *shared.SourceCaptainDataPutRequest `request:"mediaType=application/json"`
}

func (o *PutSourceCaptainDataRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

func (o *PutSourceCaptainDataRequest) GetSourceCaptainDataPutRequest() *shared.SourceCaptainDataPutRequest {
	if o == nil {
		return nil
	}
	return o.SourceCaptainDataPutRequest
}

type PutSourceCaptainDataResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSourceCaptainDataResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSourceCaptainDataResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSourceCaptainDataResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
