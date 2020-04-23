// Code generated by go-swagger; DO NOT EDIT.

package notifications

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

// NewAddVoicemailSmsNotificationParams creates a new AddVoicemailSmsNotificationParams object
// with the default values initialized.
func NewAddVoicemailSmsNotificationParams() *AddVoicemailSmsNotificationParams {
	var ()
	return &AddVoicemailSmsNotificationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddVoicemailSmsNotificationParamsWithTimeout creates a new AddVoicemailSmsNotificationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddVoicemailSmsNotificationParamsWithTimeout(timeout time.Duration) *AddVoicemailSmsNotificationParams {
	var ()
	return &AddVoicemailSmsNotificationParams{

		timeout: timeout,
	}
}

// NewAddVoicemailSmsNotificationParamsWithContext creates a new AddVoicemailSmsNotificationParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddVoicemailSmsNotificationParamsWithContext(ctx context.Context) *AddVoicemailSmsNotificationParams {
	var ()
	return &AddVoicemailSmsNotificationParams{

		Context: ctx,
	}
}

// NewAddVoicemailSmsNotificationParamsWithHTTPClient creates a new AddVoicemailSmsNotificationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddVoicemailSmsNotificationParamsWithHTTPClient(client *http.Client) *AddVoicemailSmsNotificationParams {
	var ()
	return &AddVoicemailSmsNotificationParams{
		HTTPClient: client,
	}
}

/*AddVoicemailSmsNotificationParams contains all the parameters to send to the API endpoint
for the add voicemail sms notification operation typically these are written to a http.Request
*/
type AddVoicemailSmsNotificationParams struct {

	/*Body*/
	Body *models.AddVoicemailSmslVoicemailNotificationRequest
	/*UserID
	  The unique user identifier

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add voicemail sms notification params
func (o *AddVoicemailSmsNotificationParams) WithTimeout(timeout time.Duration) *AddVoicemailSmsNotificationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add voicemail sms notification params
func (o *AddVoicemailSmsNotificationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add voicemail sms notification params
func (o *AddVoicemailSmsNotificationParams) WithContext(ctx context.Context) *AddVoicemailSmsNotificationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add voicemail sms notification params
func (o *AddVoicemailSmsNotificationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add voicemail sms notification params
func (o *AddVoicemailSmsNotificationParams) WithHTTPClient(client *http.Client) *AddVoicemailSmsNotificationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add voicemail sms notification params
func (o *AddVoicemailSmsNotificationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add voicemail sms notification params
func (o *AddVoicemailSmsNotificationParams) WithBody(body *models.AddVoicemailSmslVoicemailNotificationRequest) *AddVoicemailSmsNotificationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add voicemail sms notification params
func (o *AddVoicemailSmsNotificationParams) SetBody(body *models.AddVoicemailSmslVoicemailNotificationRequest) {
	o.Body = body
}

// WithUserID adds the userID to the add voicemail sms notification params
func (o *AddVoicemailSmsNotificationParams) WithUserID(userID string) *AddVoicemailSmsNotificationParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the add voicemail sms notification params
func (o *AddVoicemailSmsNotificationParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *AddVoicemailSmsNotificationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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