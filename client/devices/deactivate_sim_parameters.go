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

// NewDeactivateSimParams creates a new DeactivateSimParams object
// with the default values initialized.
func NewDeactivateSimParams() *DeactivateSimParams {
	var ()
	return &DeactivateSimParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeactivateSimParamsWithTimeout creates a new DeactivateSimParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeactivateSimParamsWithTimeout(timeout time.Duration) *DeactivateSimParams {
	var ()
	return &DeactivateSimParams{

		timeout: timeout,
	}
}

// NewDeactivateSimParamsWithContext creates a new DeactivateSimParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeactivateSimParamsWithContext(ctx context.Context) *DeactivateSimParams {
	var ()
	return &DeactivateSimParams{

		Context: ctx,
	}
}

// NewDeactivateSimParamsWithHTTPClient creates a new DeactivateSimParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeactivateSimParamsWithHTTPClient(client *http.Client) *DeactivateSimParams {
	var ()
	return &DeactivateSimParams{
		HTTPClient: client,
	}
}

/*DeactivateSimParams contains all the parameters to send to the API endpoint
for the deactivate sim operation typically these are written to a http.Request
*/
type DeactivateSimParams struct {

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

// WithTimeout adds the timeout to the deactivate sim params
func (o *DeactivateSimParams) WithTimeout(timeout time.Duration) *DeactivateSimParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the deactivate sim params
func (o *DeactivateSimParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the deactivate sim params
func (o *DeactivateSimParams) WithContext(ctx context.Context) *DeactivateSimParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the deactivate sim params
func (o *DeactivateSimParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the deactivate sim params
func (o *DeactivateSimParams) WithHTTPClient(client *http.Client) *DeactivateSimParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the deactivate sim params
func (o *DeactivateSimParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the deactivate sim params
func (o *DeactivateSimParams) WithDeviceID(deviceID string) *DeactivateSimParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the deactivate sim params
func (o *DeactivateSimParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WithUserID adds the userID to the deactivate sim params
func (o *DeactivateSimParams) WithUserID(userID string) *DeactivateSimParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the deactivate sim params
func (o *DeactivateSimParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *DeactivateSimParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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