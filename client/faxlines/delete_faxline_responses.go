// Code generated by go-swagger; DO NOT EDIT.

package faxlines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteFaxlineReader is a Reader for the DeleteFaxline structure.
type DeleteFaxlineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFaxlineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewDeleteFaxlineDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewDeleteFaxlineDefault creates a DeleteFaxlineDefault with default headers values
func NewDeleteFaxlineDefault(code int) *DeleteFaxlineDefault {
	return &DeleteFaxlineDefault{
		_statusCode: code,
	}
}

/*DeleteFaxlineDefault handles this case with default header values.

successful operation
*/
type DeleteFaxlineDefault struct {
	_statusCode int
}

// Code gets the status code for the delete faxline default response
func (o *DeleteFaxlineDefault) Code() int {
	return o._statusCode
}

func (o *DeleteFaxlineDefault) Error() string {
	return fmt.Sprintf("[DELETE /{userId}/faxlines/{faxlineId}][%d] deleteFaxline default ", o._statusCode)
}

func (o *DeleteFaxlineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
