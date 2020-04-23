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

// UpdateHistoryEntriesRequest update history entries request
//
// swagger:model UpdateHistoryEntriesRequest
type UpdateHistoryEntriesRequest struct {

	// archived
	Archived *bool `json:"archived,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// read
	Read *bool `json:"read,omitempty"`

	// starred
	Starred *bool `json:"starred,omitempty"`
}

// Validate validates this update history entries request
func (m *UpdateHistoryEntriesRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateHistoryEntriesRequest) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateHistoryEntriesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateHistoryEntriesRequest) UnmarshalBinary(b []byte) error {
	var res UpdateHistoryEntriesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}