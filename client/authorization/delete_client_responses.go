// Code generated by go-swagger; DO NOT EDIT.

package authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteClientReader is a Reader for the DeleteClient structure.
type DeleteClientReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteClientReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteClientAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteClientAccepted creates a DeleteClientAccepted with default headers values
func NewDeleteClientAccepted() *DeleteClientAccepted {
	return &DeleteClientAccepted{}
}

/*DeleteClientAccepted handles this case with default header values.

Deletion request accepted
*/
type DeleteClientAccepted struct {
}

func (o *DeleteClientAccepted) Error() string {
	return fmt.Sprintf("[DELETE /authorization/oauth2/clients/{clientId}][%d] deleteClientAccepted ", 202)
}

func (o *DeleteClientAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
