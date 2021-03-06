// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SimpleGroupResponse simple group response
//
// swagger:model SimpleGroupResponse
type SimpleGroupResponse struct {

	// alias
	Alias string `json:"alias,omitempty"`

	// department Id
	DepartmentID int64 `json:"departmentId,omitempty"`

	// group numbers Url
	GroupNumbersURL string `json:"groupNumbersUrl,omitempty"`

	// group users Url
	GroupUsersURL string `json:"groupUsersUrl,omitempty"`

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this simple group response
func (m *SimpleGroupResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SimpleGroupResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SimpleGroupResponse) UnmarshalBinary(b []byte) error {
	var res SimpleGroupResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
