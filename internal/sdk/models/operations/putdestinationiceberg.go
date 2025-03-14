// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"net/http"
)

type PutDestinationIcebergRequest struct {
	DestinationID                string                               `pathParam:"style=simple,explode=false,name=destinationId"`
	DestinationIcebergPutRequest *shared.DestinationIcebergPutRequest `request:"mediaType=application/json"`
}

func (o *PutDestinationIcebergRequest) GetDestinationID() string {
	if o == nil {
		return ""
	}
	return o.DestinationID
}

func (o *PutDestinationIcebergRequest) GetDestinationIcebergPutRequest() *shared.DestinationIcebergPutRequest {
	if o == nil {
		return nil
	}
	return o.DestinationIcebergPutRequest
}

type PutDestinationIcebergResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutDestinationIcebergResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutDestinationIcebergResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutDestinationIcebergResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
