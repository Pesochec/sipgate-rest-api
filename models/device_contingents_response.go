// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeviceContingentsResponse device contingents response
//
// swagger:model DeviceContingentsResponse
type DeviceContingentsResponse struct {

	// contingents
	Contingents []*ContingentResponse `json:"contingents"`
}

// Validate validates this device contingents response
func (m *DeviceContingentsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContingents(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceContingentsResponse) validateContingents(formats strfmt.Registry) error {

	if swag.IsZero(m.Contingents) { // not required
		return nil
	}

	for i := 0; i < len(m.Contingents); i++ {
		if swag.IsZero(m.Contingents[i]) { // not required
			continue
		}

		if m.Contingents[i] != nil {
			if err := m.Contingents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("contingents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceContingentsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceContingentsResponse) UnmarshalBinary(b []byte) error {
	var res DeviceContingentsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}