// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"net/http"
)

type PutSourceAuth0Request struct {
	SourceID              string                        `pathParam:"style=simple,explode=false,name=sourceId"`
	SourceAuth0PutRequest *shared.SourceAuth0PutRequest `request:"mediaType=application/json"`
}

func (o *PutSourceAuth0Request) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

func (o *PutSourceAuth0Request) GetSourceAuth0PutRequest() *shared.SourceAuth0PutRequest {
	if o == nil {
		return nil
	}
	return o.SourceAuth0PutRequest
}

type PutSourceAuth0Response struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSourceAuth0Response) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSourceAuth0Response) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSourceAuth0Response) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
