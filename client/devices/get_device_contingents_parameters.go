// Code generated by go-swagger; DO NOT EDIT.

package devices

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

// NewGetDeviceContingentsParams creates a new GetDeviceContingentsParams object
// with the default values initialized.
func NewGetDeviceContingentsParams() *GetDeviceContingentsParams {
	var ()
	return &GetDeviceContingentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceContingentsParamsWithTimeout creates a new GetDeviceContingentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDeviceContingentsParamsWithTimeout(timeout time.Duration) *GetDeviceContingentsParams {
	var ()
	return &GetDeviceContingentsParams{

		timeout: timeout,
	}
}

// NewGetDeviceContingentsParamsWithContext creates a new GetDeviceContingentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDeviceContingentsParamsWithContext(ctx context.Context) *GetDeviceContingentsParams {
	var ()
	return &GetDeviceContingentsParams{

		Context: ctx,
	}
}

// NewGetDeviceContingentsParamsWithHTTPClient creates a new GetDeviceContingentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDeviceContingentsParamsWithHTTPClient(client *http.Client) *GetDeviceContingentsParams {
	var ()
	return &GetDeviceContingentsParams{
		HTTPClient: client,
	}
}

/*GetDeviceContingentsParams contains all the parameters to send to the API endpoint
for the get device contingents operation typically these are written to a http.Request
*/
type GetDeviceContingentsParams struct {

	/*DeviceID
	  The unique mobile device identifier

	*/
	DeviceID string
	/*UserID
	  The unique user identifier

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get device contingents params
func (o *GetDeviceContingentsParams) WithTimeout(timeout time.Duration) *GetDeviceContingentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device contingents params
func (o *GetDeviceContingentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device contingents params
func (o *GetDeviceContingentsParams) WithContext(ctx context.Context) *GetDeviceContingentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device contingents params
func (o *GetDeviceContingentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device contingents params
func (o *GetDeviceContingentsParams) WithHTTPClient(client *http.Client) *GetDeviceContingentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device contingents params
func (o *GetDeviceContingentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the get device contingents params
func (o *GetDeviceContingentsParams) WithDeviceID(deviceID string) *GetDeviceContingentsParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get device contingents params
func (o *GetDeviceContingentsParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WithUserID adds the userID to the get device contingents params
func (o *GetDeviceContingentsParams) WithUserID(userID string) *GetDeviceContingentsParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get device contingents params
func (o *GetDeviceContingentsParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceContingentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deviceId
	if err := r.SetPathParam("deviceId", o.DeviceID); err != nil {
		return err
	}

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}