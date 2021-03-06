// Code generated by go-swagger; DO NOT EDIT.

package history

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteHistoryEntriesReader is a Reader for the DeleteHistoryEntries structure.
type DeleteHistoryEntriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteHistoryEntriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewDeleteHistoryEntriesDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewDeleteHistoryEntriesDefault creates a DeleteHistoryEntriesDefault with default headers values
func NewDeleteHistoryEntriesDefault(code int) *DeleteHistoryEntriesDefault {
	return &DeleteHistoryEntriesDefault{
		_statusCode: code,
	}
}

/*DeleteHistoryEntriesDefault handles this case with default header values.

successful operation
*/
type DeleteHistoryEntriesDefault struct {
	_statusCode int
}

// Code gets the status code for the delete history entries default response
func (o *DeleteHistoryEntriesDefault) Code() int {
	return o._statusCode
}

func (o *DeleteHistoryEntriesDefault) Error() string {
	return fmt.Sprintf("[DELETE /history][%d] deleteHistoryEntries default ", o._statusCode)
}

func (o *DeleteHistoryEntriesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
