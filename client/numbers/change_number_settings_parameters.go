// Code generated by go-swagger; DO NOT EDIT.

package numbers

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

	"github.com/kben/sipgate-rest-api/models"
)

// NewChangeNumberSettingsParams creates a new ChangeNumberSettingsParams object
// with the default values initialized.
func NewChangeNumberSettingsParams() *ChangeNumberSettingsParams {
	var ()
	return &ChangeNumberSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewChangeNumberSettingsParamsWithTimeout creates a new ChangeNumberSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChangeNumberSettingsParamsWithTimeout(timeout time.Duration) *ChangeNumberSettingsParams {
	var ()
	return &ChangeNumberSettingsParams{

		timeout: timeout,
	}
}

// NewChangeNumberSettingsParamsWithContext creates a new ChangeNumberSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewChangeNumberSettingsParamsWithContext(ctx context.Context) *ChangeNumberSettingsParams {
	var ()
	return &ChangeNumberSettingsParams{

		Context: ctx,
	}
}

// NewChangeNumberSettingsParamsWithHTTPClient creates a new ChangeNumberSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChangeNumberSettingsParamsWithHTTPClient(client *http.Client) *ChangeNumberSettingsParams {
	var ()
	return &ChangeNumberSettingsParams{
		HTTPClient: client,
	}
}

/*ChangeNumberSettingsParams contains all the parameters to send to the API endpoint
for the change number settings operation typically these are written to a http.Request
*/
type ChangeNumberSettingsParams struct {

	/*Body*/
	Body *models.ChangeNumberSettingsRequest
	/*NumberID
	  The unique phone number identifier

	*/
	NumberID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the change number settings params
func (o *ChangeNumberSettingsParams) WithTimeout(timeout time.Duration) *ChangeNumberSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change number settings params
func (o *ChangeNumberSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change number settings params
func (o *ChangeNumberSettingsParams) WithContext(ctx context.Context) *ChangeNumberSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change number settings params
func (o *ChangeNumberSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change number settings params
func (o *ChangeNumberSettingsParams) WithHTTPClient(client *http.Client) *ChangeNumberSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change number settings params
func (o *ChangeNumberSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the change number settings params
func (o *ChangeNumberSettingsParams) WithBody(body *models.ChangeNumberSettingsRequest) *ChangeNumberSettingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the change number settings params
func (o *ChangeNumberSettingsParams) SetBody(body *models.ChangeNumberSettingsRequest) {
	o.Body = body
}

// WithNumberID adds the numberID to the change number settings params
func (o *ChangeNumberSettingsParams) WithNumberID(numberID string) *ChangeNumberSettingsParams {
	o.SetNumberID(numberID)
	return o
}

// SetNumberID adds the numberId to the change number settings params
func (o *ChangeNumberSettingsParams) SetNumberID(numberID string) {
	o.NumberID = numberID
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeNumberSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param numberId
	if err := r.SetPathParam("numberId", o.NumberID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}