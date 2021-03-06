// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ContingentResponse contingent response
//
// swagger:model ContingentResponse
type ContingentResponse struct {

	// amount
	Amount int32 `json:"amount,omitempty"`

	// expires
	Expires string `json:"expires,omitempty"`

	// left
	Left int32 `json:"left,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// unit
	Unit string `json:"unit,omitempty"`

	// used
	Used int32 `json:"used,omitempty"`
}

// Validate validates this contingent response
func (m *ContingentResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ContingentResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContingentResponse) UnmarshalBinary(b []byte) error {
	var res ContingentResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
