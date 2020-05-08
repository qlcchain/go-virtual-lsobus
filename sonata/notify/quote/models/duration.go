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

// Duration Duration of a term (value + unit)
//
// swagger:model Duration
type Duration struct {

	// duration
	// Required: true
	Duration DurationUnit `json:"duration"`

	// unit
	// Required: true
	Unit *int32 `json:"unit"`
}

// Validate validates this duration
func (m *Duration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Duration) validateDuration(formats strfmt.Registry) error {

	if err := m.Duration.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("duration")
		}
		return err
	}

	return nil
}

func (m *Duration) validateUnit(formats strfmt.Registry) error {

	if err := validate.Required("unit", "body", m.Unit); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Duration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Duration) UnmarshalBinary(b []byte) error {
	var res Duration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
