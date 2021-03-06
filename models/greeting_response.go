// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GreetingResponse greeting response
//
// swagger:model GreetingResponse
type GreetingResponse struct {

	// active
	Active *bool `json:"active,omitempty"`

	// alias
	Alias string `json:"alias,omitempty"`

	// greeting Url
	GreetingURL string `json:"greetingUrl,omitempty"`

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this greeting response
func (m *GreetingResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GreetingResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GreetingResponse) UnmarshalBinary(b []byte) error {
	var res GreetingResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
