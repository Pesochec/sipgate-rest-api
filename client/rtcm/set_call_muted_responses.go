// Code generated by go-swagger; DO NOT EDIT.

package rtcm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SetCallMutedReader is a Reader for the SetCallMuted structure.
type SetCallMutedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetCallMutedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewSetCallMutedNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewSetCallMutedNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetCallMutedNoContent creates a SetCallMutedNoContent with default headers values
func NewSetCallMutedNoContent() *SetCallMutedNoContent {
	return &SetCallMutedNoContent{}
}

/*SetCallMutedNoContent handles this case with default header values.

Success
*/
type SetCallMutedNoContent struct {
}

func (o *SetCallMutedNoContent) Error() string {
	return fmt.Sprintf("[PUT /calls/{callId}/muted][%d] setCallMutedNoContent ", 204)
}

func (o *SetCallMutedNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSetCallMutedNotFound creates a SetCallMutedNotFound with default headers values
func NewSetCallMutedNotFound() *SetCallMutedNotFound {
	return &SetCallMutedNotFound{}
}

/*SetCallMutedNotFound handles this case with default header values.

Call not found
*/
type SetCallMutedNotFound struct {
}

func (o *SetCallMutedNotFound) Error() string {
	return fmt.Sprintf("[PUT /calls/{callId}/muted][%d] setCallMutedNotFound ", 404)
}

func (o *SetCallMutedNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}