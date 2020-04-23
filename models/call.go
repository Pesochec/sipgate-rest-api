// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Call call
//
// swagger:model Call
type Call struct {

	// The unique call identifier
	// Required: true
	CallID *string `json:"callId"`

	// The call is on hold
	// Required: true
	Hold bool `json:"hold"`

	// The call is muted
	// Required: true
	Muted bool `json:"muted"`

	// participants
	// Required: true
	Participants []*Participant `json:"participants"`

	// The call is being recorded
	// Required: true
	Recording bool `json:"recording"`
}

// Validate validates this call
func (m *Call) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCallID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHold(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMuted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParticipants(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecording(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Call) validateCallID(formats strfmt.Registry) error {

	if err := validate.Required("callId", "body", m.CallID); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateHold(formats strfmt.Registry) error {

	if err := validate.Required("hold", "body", bool(m.Hold)); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateMuted(formats strfmt.Registry) error {

	if err := validate.Required("muted", "body", bool(m.Muted)); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateParticipants(formats strfmt.Registry) error {

	if err := validate.Required("participants", "body", m.Participants); err != nil {
		return err
	}

	for i := 0; i < len(m.Participants); i++ {
		if swag.IsZero(m.Participants[i]) { // not required
			continue
		}

		if m.Participants[i] != nil {
			if err := m.Participants[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("participants" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Call) validateRecording(formats strfmt.Registry) error {

	if err := validate.Required("recording", "body", bool(m.Recording)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Call) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Call) UnmarshalBinary(b []byte) error {
	var res Call
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}