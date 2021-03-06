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

// AddCallSmsNotificationRequest add call sms notification request
//
// swagger:model AddCallSmsNotificationRequest
type AddCallSmsNotificationRequest struct {

	// cause
	// Required: true
	Cause *string `json:"cause"`

	// direction
	// Required: true
	Direction *string `json:"direction"`

	// endpoint Id
	// Required: true
	EndpointID *string `json:"endpointId"`

	// number
	// Required: true
	Number *string `json:"number"`
}

// Validate validates this add call sms notification request
func (m *AddCallSmsNotificationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCause(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndpointID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumber(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddCallSmsNotificationRequest) validateCause(formats strfmt.Registry) error {

	if err := validate.Required("cause", "body", m.Cause); err != nil {
		return err
	}

	return nil
}

func (m *AddCallSmsNotificationRequest) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *AddCallSmsNotificationRequest) validateEndpointID(formats strfmt.Registry) error {

	if err := validate.Required("endpointId", "body", m.EndpointID); err != nil {
		return err
	}

	return nil
}

func (m *AddCallSmsNotificationRequest) validateNumber(formats strfmt.Registry) error {

	if err := validate.Required("number", "body", m.Number); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AddCallSmsNotificationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddCallSmsNotificationRequest) UnmarshalBinary(b []byte) error {
	var res AddCallSmsNotificationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
