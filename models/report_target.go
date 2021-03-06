// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ReportTarget report target
//
// swagger:model ReportTarget
type ReportTarget struct {

	// email
	Email string `json:"email,omitempty"`

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this report target
func (m *ReportTarget) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReportTarget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportTarget) UnmarshalBinary(b []byte) error {
	var res ReportTarget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
