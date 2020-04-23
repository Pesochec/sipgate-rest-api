// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DetailedError detailed error
//
// swagger:model DetailedError
type DetailedError struct {

	// The detailed description of the error
	Detail string `json:"detail,omitempty"`

	// source
	Source *ErrorSource `json:"source,omitempty"`

	// The status code
	Status int32 `json:"status,omitempty"`

	// The title of the error
	Title string `json:"title,omitempty"`
}

// Validate validates this detailed error
func (m *DetailedError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DetailedError) validateSource(formats strfmt.Registry) error {

	if swag.IsZero(m.Source) { // not required
		return nil
	}

	if m.Source != nil {
		if err := m.Source.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DetailedError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DetailedError) UnmarshalBinary(b []byte) error {
	var res DetailedError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}