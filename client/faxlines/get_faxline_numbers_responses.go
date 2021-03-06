// Code generated by go-swagger; DO NOT EDIT.

package faxlines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kben/sipgate-rest-api/models"
)

// GetFaxlineNumbersReader is a Reader for the GetFaxlineNumbers structure.
type GetFaxlineNumbersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFaxlineNumbersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFaxlineNumbersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFaxlineNumbersOK creates a GetFaxlineNumbersOK with default headers values
func NewGetFaxlineNumbersOK() *GetFaxlineNumbersOK {
	return &GetFaxlineNumbersOK{}
}

/*GetFaxlineNumbersOK handles this case with default header values.

successful operation
*/
type GetFaxlineNumbersOK struct {
	Payload *models.EndpointNumbersResponse
}

func (o *GetFaxlineNumbersOK) Error() string {
	return fmt.Sprintf("[GET /{userId}/faxlines/{faxlineId}/numbers][%d] getFaxlineNumbersOK  %+v", 200, o.Payload)
}

func (o *GetFaxlineNumbersOK) GetPayload() *models.EndpointNumbersResponse {
	return o.Payload
}

func (o *GetFaxlineNumbersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EndpointNumbersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
