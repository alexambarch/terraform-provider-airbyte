// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"net/http"
)

type PutSourceShipstationRequest struct {
	SourceID                    string                              `pathParam:"style=simple,explode=false,name=sourceId"`
	SourceShipstationPutRequest *shared.SourceShipstationPutRequest `request:"mediaType=application/json"`
}

func (o *PutSourceShipstationRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

func (o *PutSourceShipstationRequest) GetSourceShipstationPutRequest() *shared.SourceShipstationPutRequest {
	if o == nil {
		return nil
	}
	return o.SourceShipstationPutRequest
}

type PutSourceShipstationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSourceShipstationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSourceShipstationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSourceShipstationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
