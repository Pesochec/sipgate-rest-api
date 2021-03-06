// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FaxSmsTarget fax sms target
//
// swagger:model FaxSmsTarget
type FaxSmsTarget struct {

	// direction
	// Enum: [INCOMING OUTGOING]
	Direction string `json:"direction,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// number
	Number string `json:"number,omitempty"`
}

// Validate validates this fax sms target
func (m *FaxSmsTarget) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var faxSmsTargetTypeDirectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["INCOMING","OUTGOING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		faxSmsTargetTypeDirectionPropEnum = append(faxSmsTargetTypeDirectionPropEnum, v)
	}
}

const (

	// FaxSmsTargetDirectionINCOMING captures enum value "INCOMING"
	FaxSmsTargetDirectionINCOMING string = "INCOMING"

	// FaxSmsTargetDirectionOUTGOING captures enum value "OUTGOING"
	FaxSmsTargetDirectionOUTGOING string = "OUTGOING"
)

// prop value enum
func (m *FaxSmsTarget) validateDirectionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, faxSmsTargetTypeDirectionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *FaxSmsTarget) validateDirection(formats strfmt.Registry) error {

	if swag.IsZero(m.Direction) { // not required
		return nil
	}

	// value enum
	if err := m.validateDirectionEnum("direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FaxSmsTarget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FaxSmsTarget) UnmarshalBinary(b []byte) error {
	var res FaxSmsTarget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
