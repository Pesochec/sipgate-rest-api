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

// NewCreateEmailNotificationForCallParams creates a new CreateEmailNotificationForCallParams object
// with the default values initialized.
func NewCreateEmailNotificationForCallParams() *CreateEmailNotificationForCallParams {
	var ()
	return &CreateEmailNotificationForCallParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateEmailNotificationForCallParamsWithTimeout creates a new CreateEmailNotificationForCallParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateEmailNotificationForCallParamsWithTimeout(timeout time.Duration) *CreateEmailNotificationForCallParams {
	var ()
	return &CreateEmailNotificationForCallParams{

		timeout: timeout,
	}
}

// NewCreateEmailNotificationForCallParamsWithContext creates a new CreateEmailNotificationForCallParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateEmailNotificationForCallParamsWithContext(ctx context.Context) *CreateEmailNotificationForCallParams {
	var ()
	return &CreateEmailNotificationForCallParams{

		Context: ctx,
	}
}

// NewCreateEmailNotificationForCallParamsWithHTTPClient creates a new CreateEmailNotificationForCallParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateEmailNotificationForCallParamsWithHTTPClient(client *http.Client) *CreateEmailNotificationForCallParams {
	var ()
	return &CreateEmailNotificationForCallParams{
		HTTPClient: client,
	}
}

/*CreateEmailNotificationForCallParams contains all the parameters to send to the API endpoint
for the create email notification for call operation typically these are written to a http.Request
*/
type CreateEmailNotificationForCallParams struct {

	/*Body*/
	Body *models.AddCallEmailNotificationRequest
	/*UserID
	  The unique user identifier

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create email notification for call params
func (o *CreateEmailNotificationForCallParams) WithTimeout(timeout time.Duration) *CreateEmailNotificationForCallParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create email notification for call params
func (o *CreateEmailNotificationForCallParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create email notification for call params
func (o *CreateEmailNotificationForCallParams) WithContext(ctx context.Context) *CreateEmailNotificationForCallParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create email notification for call params
func (o *CreateEmailNotificationForCallParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create email notification for call params
func (o *CreateEmailNotificationForCallParams) WithHTTPClient(client *http.Client) *CreateEmailNotificationForCallParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create email notification for call params
func (o *CreateEmailNotificationForCallParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create email notification for call params
func (o *CreateEmailNotificationForCallParams) WithBody(body *models.AddCallEmailNotificationRequest) *CreateEmailNotificationForCallParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create email notification for call params
func (o *CreateEmailNotificationForCallParams) SetBody(body *models.AddCallEmailNotificationRequest) {
	o.Body = body
}

// WithUserID adds the userID to the create email notification for call params
func (o *CreateEmailNotificationForCallParams) WithUserID(userID string) *CreateEmailNotificationForCallParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the create email notification for call params
func (o *CreateEmailNotificationForCallParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateEmailNotificationForCallParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
