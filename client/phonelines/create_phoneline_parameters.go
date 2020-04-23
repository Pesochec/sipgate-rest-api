// Code generated by go-swagger; DO NOT EDIT.

package phonelines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewCreatePhonelineParams creates a new CreatePhonelineParams object
// with the default values initialized.
func NewCreatePhonelineParams() *CreatePhonelineParams {
	var ()
	return &CreatePhonelineParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePhonelineParamsWithTimeout creates a new CreatePhonelineParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreatePhonelineParamsWithTimeout(timeout time.Duration) *CreatePhonelineParams {
	var ()
	return &CreatePhonelineParams{

		timeout: timeout,
	}
}

// NewCreatePhonelineParamsWithContext creates a new CreatePhonelineParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreatePhonelineParamsWithContext(ctx context.Context) *CreatePhonelineParams {
	var ()
	return &CreatePhonelineParams{

		Context: ctx,
	}
}

// NewCreatePhonelineParamsWithHTTPClient creates a new CreatePhonelineParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreatePhonelineParamsWithHTTPClient(client *http.Client) *CreatePhonelineParams {
	var ()
	return &CreatePhonelineParams{
		HTTPClient: client,
	}
}

/*CreatePhonelineParams contains all the parameters to send to the API endpoint
for the create phoneline operation typically these are written to a http.Request
*/
type CreatePhonelineParams struct {

	/*UserID
	  The unique user identifier

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create phoneline params
func (o *CreatePhonelineParams) WithTimeout(timeout time.Duration) *CreatePhonelineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create phoneline params
func (o *CreatePhonelineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create phoneline params
func (o *CreatePhonelineParams) WithContext(ctx context.Context) *CreatePhonelineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create phoneline params
func (o *CreatePhonelineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create phoneline params
func (o *CreatePhonelineParams) WithHTTPClient(client *http.Client) *CreatePhonelineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create phoneline params
func (o *CreatePhonelineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the create phoneline params
func (o *CreatePhonelineParams) WithUserID(userID string) *CreatePhonelineParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the create phoneline params
func (o *CreatePhonelineParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePhonelineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}