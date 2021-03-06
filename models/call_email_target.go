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

// CallEmailTarget call email target
//
// swagger:model CallEmailTarget
type CallEmailTarget struct {

	// cause
	// Enum: [MISSED SUCCESSFUL]
	Cause string `json:"cause,omitempty"`

	// direction
	// Enum: [INCOMING OUTGOING]
	Direction string `json:"direction,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this call email target
func (m *CallEmailTarget) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCause(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var callEmailTargetTypeCausePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["MISSED","SUCCESSFUL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		callEmailTargetTypeCausePropEnum = append(callEmailTargetTypeCausePropEnum, v)
	}
}

const (

	// CallEmailTargetCauseMISSED captures enum value "MISSED"
	CallEmailTargetCauseMISSED string = "MISSED"

	// CallEmailTargetCauseSUCCESSFUL captures enum value "SUCCESSFUL"
	CallEmailTargetCauseSUCCESSFUL string = "SUCCESSFUL"
)

// prop value enum
func (m *CallEmailTarget) validateCauseEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, callEmailTargetTypeCausePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CallEmailTarget) validateCause(formats strfmt.Registry) error {

	if swag.IsZero(m.Cause) { // not required
		return nil
	}

	// value enum
	if err := m.validateCauseEnum("cause", "body", m.Cause); err != nil {
		return err
	}

	return nil
}

var callEmailTargetTypeDirectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["INCOMING","OUTGOING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		callEmailTargetTypeDirectionPropEnum = append(callEmailTargetTypeDirectionPropEnum, v)
	}
}

const (

	// CallEmailTargetDirectionINCOMING captures enum value "INCOMING"
	CallEmailTargetDirectionINCOMING string = "INCOMING"

	// CallEmailTargetDirectionOUTGOING captures enum value "OUTGOING"
	CallEmailTargetDirectionOUTGOING string = "OUTGOING"
)

// prop value enum
func (m *CallEmailTarget) validateDirectionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, callEmailTargetTypeDirectionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CallEmailTarget) validateDirection(formats strfmt.Registry) error {

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
func (m *CallEmailTarget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CallEmailTarget) UnmarshalBinary(b []byte) error {
	var res CallEmailTarget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
