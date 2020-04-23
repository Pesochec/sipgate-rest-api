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

// SendFaxRequest send fax request
//
// swagger:model SendFaxRequest
type SendFaxRequest struct {

	// base64 content
	// Required: true
	// Max Length: 28330000
	// Min Length: 0
	Base64Content *string `json:"base64Content"`

	// faxline Id
	// Required: true
	FaxlineID *string `json:"faxlineId"`

	// filename
	// Required: true
	Filename *string `json:"filename"`

	// recipient
	// Required: true
	Recipient *string `json:"recipient"`
}

// Validate validates this send fax request
func (m *SendFaxRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBase64Content(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFaxlineID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilename(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipient(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SendFaxRequest) validateBase64Content(formats strfmt.Registry) error {

	if err := validate.Required("base64Content", "body", m.Base64Content); err != nil {
		return err
	}

	if err := validate.MinLength("base64Content", "body", string(*m.Base64Content), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("base64Content", "body", string(*m.Base64Content), 28330000); err != nil {
		return err
	}

	return nil
}

func (m *SendFaxRequest) validateFaxlineID(formats strfmt.Registry) error {

	if err := validate.Required("faxlineId", "body", m.FaxlineID); err != nil {
		return err
	}

	return nil
}

func (m *SendFaxRequest) validateFilename(formats strfmt.Registry) error {

	if err := validate.Required("filename", "body", m.Filename); err != nil {
		return err
	}

	return nil
}

func (m *SendFaxRequest) validateRecipient(formats strfmt.Registry) error {

	if err := validate.Required("recipient", "body", m.Recipient); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SendFaxRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SendFaxRequest) UnmarshalBinary(b []byte) error {
	var res SendFaxRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}