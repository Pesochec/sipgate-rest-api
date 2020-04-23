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

// SendWebSmsRequest send web sms request
//
// swagger:model SendWebSmsRequest
type SendWebSmsRequest struct {

	// message
	// Required: true
	Message *string `json:"message"`

	// recipient
	// Required: true
	Recipient *string `json:"recipient"`

	// send at
	SendAt int64 `json:"sendAt,omitempty"`

	// sms Id
	// Required: true
	SmsID *string `json:"smsId"`
}

// Validate validates this send web sms request
func (m *SendWebSmsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipient(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSmsID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SendWebSmsRequest) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

func (m *SendWebSmsRequest) validateRecipient(formats strfmt.Registry) error {

	if err := validate.Required("recipient", "body", m.Recipient); err != nil {
		return err
	}

	return nil
}

func (m *SendWebSmsRequest) validateSmsID(formats strfmt.Registry) error {

	if err := validate.Required("smsId", "body", m.SmsID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SendWebSmsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SendWebSmsRequest) UnmarshalBinary(b []byte) error {
	var res SendWebSmsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}