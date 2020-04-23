// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AddFaxReportEmailNotificationRequest add fax report email notification request
//
// swagger:model AddFaxReportEmailNotificationRequest
type AddFaxReportEmailNotificationRequest struct {

	// email
	// Required: true
	Email *string `json:"email"`

	// faxline Id
	// Required: true
	FaxlineID *string `json:"faxlineId"`
}

// Validate validates this add fax report email notification request
func (m *AddFaxReportEmailNotificationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFaxlineID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddFaxReportEmailNotificationRequest) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	return nil
}

func (m *AddFaxReportEmailNotificationRequest) validateFaxlineID(formats strfmt.Registry) error {

	if err := validate.Required("faxlineId", "body", m.FaxlineID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AddFaxReportEmailNotificationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddFaxReportEmailNotificationRequest) UnmarshalBinary(b []byte) error {
	var res AddFaxReportEmailNotificationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}