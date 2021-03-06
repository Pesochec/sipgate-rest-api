// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ExternalDeviceResponse external device response
//
// swagger:model ExternalDeviceResponse
type ExternalDeviceResponse struct {

	// active groups
	ActiveGroups []*ActiveRouting `json:"activeGroups"`

	// active phonelines
	ActivePhonelines []*ActiveRouting `json:"activePhonelines"`

	// alias
	Alias string `json:"alias,omitempty"`

	// dnd
	Dnd *bool `json:"dnd,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// incoming call display state
	// Enum: [CALLED_NUMBER CALLER_NUMBER]
	IncomingCallDisplayState string `json:"incomingCallDisplayState,omitempty"`

	// number
	Number string `json:"number,omitempty"`

	// online
	Online *bool `json:"online,omitempty"`

	// type
	// Enum: [REGISTER MOBILE EXTERNAL]
	Type string `json:"type,omitempty"`
}

// Validate validates this external device response
func (m *ExternalDeviceResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActiveGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateActivePhonelines(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIncomingCallDisplayState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExternalDeviceResponse) validateActiveGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.ActiveGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.ActiveGroups); i++ {
		if swag.IsZero(m.ActiveGroups[i]) { // not required
			continue
		}

		if m.ActiveGroups[i] != nil {
			if err := m.ActiveGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("activeGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ExternalDeviceResponse) validateActivePhonelines(formats strfmt.Registry) error {

	if swag.IsZero(m.ActivePhonelines) { // not required
		return nil
	}

	for i := 0; i < len(m.ActivePhonelines); i++ {
		if swag.IsZero(m.ActivePhonelines[i]) { // not required
			continue
		}

		if m.ActivePhonelines[i] != nil {
			if err := m.ActivePhonelines[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("activePhonelines" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var externalDeviceResponseTypeIncomingCallDisplayStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CALLED_NUMBER","CALLER_NUMBER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		externalDeviceResponseTypeIncomingCallDisplayStatePropEnum = append(externalDeviceResponseTypeIncomingCallDisplayStatePropEnum, v)
	}
}

const (

	// ExternalDeviceResponseIncomingCallDisplayStateCALLEDNUMBER captures enum value "CALLED_NUMBER"
	ExternalDeviceResponseIncomingCallDisplayStateCALLEDNUMBER string = "CALLED_NUMBER"

	// ExternalDeviceResponseIncomingCallDisplayStateCALLERNUMBER captures enum value "CALLER_NUMBER"
	ExternalDeviceResponseIncomingCallDisplayStateCALLERNUMBER string = "CALLER_NUMBER"
)

// prop value enum
func (m *ExternalDeviceResponse) validateIncomingCallDisplayStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, externalDeviceResponseTypeIncomingCallDisplayStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ExternalDeviceResponse) validateIncomingCallDisplayState(formats strfmt.Registry) error {

	if swag.IsZero(m.IncomingCallDisplayState) { // not required
		return nil
	}

	// value enum
	if err := m.validateIncomingCallDisplayStateEnum("incomingCallDisplayState", "body", m.IncomingCallDisplayState); err != nil {
		return err
	}

	return nil
}

var externalDeviceResponseTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["REGISTER","MOBILE","EXTERNAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		externalDeviceResponseTypeTypePropEnum = append(externalDeviceResponseTypeTypePropEnum, v)
	}
}

const (

	// ExternalDeviceResponseTypeREGISTER captures enum value "REGISTER"
	ExternalDeviceResponseTypeREGISTER string = "REGISTER"

	// ExternalDeviceResponseTypeMOBILE captures enum value "MOBILE"
	ExternalDeviceResponseTypeMOBILE string = "MOBILE"

	// ExternalDeviceResponseTypeEXTERNAL captures enum value "EXTERNAL"
	ExternalDeviceResponseTypeEXTERNAL string = "EXTERNAL"
)

// prop value enum
func (m *ExternalDeviceResponse) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, externalDeviceResponseTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ExternalDeviceResponse) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExternalDeviceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExternalDeviceResponse) UnmarshalBinary(b []byte) error {
	var res ExternalDeviceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
