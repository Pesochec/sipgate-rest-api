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

// NewAddQuickDialNumberParams creates a new AddQuickDialNumberParams object
// with the default values initialized.
func NewAddQuickDialNumberParams() *AddQuickDialNumberParams {
	var ()
	return &AddQuickDialNumberParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddQuickDialNumberParamsWithTimeout creates a new AddQuickDialNumberParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddQuickDialNumberParamsWithTimeout(timeout time.Duration) *AddQuickDialNumberParams {
	var ()
	return &AddQuickDialNumberParams{

		timeout: timeout,
	}
}

// NewAddQuickDialNumberParamsWithContext creates a new AddQuickDialNumberParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddQuickDialNumberParamsWithContext(ctx context.Context) *AddQuickDialNumberParams {
	var ()
	return &AddQuickDialNumberParams{

		Context: ctx,
	}
}

// NewAddQuickDialNumberParamsWithHTTPClient creates a new AddQuickDialNumberParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddQuickDialNumberParamsWithHTTPClient(client *http.Client) *AddQuickDialNumberParams {
	var ()
	return &AddQuickDialNumberParams{
		HTTPClient: client,
	}
}

/*AddQuickDialNumberParams contains all the parameters to send to the API endpoint
for the add quick dial number operation typically these are written to a http.Request
*/
type AddQuickDialNumberParams struct {

	/*Body*/
	Body *models.AddDirectDialRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add quick dial number params
func (o *AddQuickDialNumberParams) WithTimeout(timeout time.Duration) *AddQuickDialNumberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add quick dial number params
func (o *AddQuickDialNumberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add quick dial number params
func (o *AddQuickDialNumberParams) WithContext(ctx context.Context) *AddQuickDialNumberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add quick dial number params
func (o *AddQuickDialNumberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add quick dial number params
func (o *AddQuickDialNumberParams) WithHTTPClient(client *http.Client) *AddQuickDialNumberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add quick dial number params
func (o *AddQuickDialNumberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add quick dial number params
func (o *AddQuickDialNumberParams) WithBody(body *models.AddDirectDialRequest) *AddQuickDialNumberParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add quick dial number params
func (o *AddQuickDialNumberParams) SetBody(body *models.AddDirectDialRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddQuickDialNumberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
