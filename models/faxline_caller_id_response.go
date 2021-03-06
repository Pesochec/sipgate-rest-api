// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FaxlineCallerIDResponse faxline caller Id response
//
// swagger:model FaxlineCallerIdResponse
type FaxlineCallerIDResponse struct {

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this faxline caller Id response
func (m *FaxlineCallerIDResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FaxlineCallerIDResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FaxlineCallerIDResponse) UnmarshalBinary(b []byte) error {
	var res FaxlineCallerIDResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
