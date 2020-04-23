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
	"github.com/go-openapi/swag"
)

// NewGetNumbersParams creates a new GetNumbersParams object
// with the default values initialized.
func NewGetNumbersParams() *GetNumbersParams {
	var ()
	return &GetNumbersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNumbersParamsWithTimeout creates a new GetNumbersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNumbersParamsWithTimeout(timeout time.Duration) *GetNumbersParams {
	var ()
	return &GetNumbersParams{

		timeout: timeout,
	}
}

// NewGetNumbersParamsWithContext creates a new GetNumbersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNumbersParamsWithContext(ctx context.Context) *GetNumbersParams {
	var ()
	return &GetNumbersParams{

		Context: ctx,
	}
}

// NewGetNumbersParamsWithHTTPClient creates a new GetNumbersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNumbersParamsWithHTTPClient(client *http.Client) *GetNumbersParams {
	var ()
	return &GetNumbersParams{
		HTTPClient: client,
	}
}

/*GetNumbersParams contains all the parameters to send to the API endpoint
for the get numbers operation typically these are written to a http.Request
*/
type GetNumbersParams struct {

	/*Limit*/
	Limit *int32
	/*Offset*/
	Offset *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get numbers params
func (o *GetNumbersParams) WithTimeout(timeout time.Duration) *GetNumbersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get numbers params
func (o *GetNumbersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get numbers params
func (o *GetNumbersParams) WithContext(ctx context.Context) *GetNumbersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get numbers params
func (o *GetNumbersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get numbers params
func (o *GetNumbersParams) WithHTTPClient(client *http.Client) *GetNumbersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get numbers params
func (o *GetNumbersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get numbers params
func (o *GetNumbersParams) WithLimit(limit *int32) *GetNumbersParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get numbers params
func (o *GetNumbersParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the get numbers params
func (o *GetNumbersParams) WithOffset(offset *int32) *GetNumbersParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get numbers params
func (o *GetNumbersParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *GetNumbersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}