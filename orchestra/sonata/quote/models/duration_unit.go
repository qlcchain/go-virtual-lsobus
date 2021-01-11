// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// DurationUnit duration unit
//
// swagger:model DurationUnit
type DurationUnit string

const (

	// DurationUnitDAY captures enum value "DAY"
	DurationUnitDAY DurationUnit = "DAY"

	// DurationUnitWEEK captures enum value "WEEK"
	DurationUnitWEEK DurationUnit = "WEEK"

	// DurationUnitMONTH captures enum value "MONTH"
	DurationUnitMONTH DurationUnit = "MONTH"

	// DurationUnitYEAR captures enum value "YEAR"
	DurationUnitYEAR DurationUnit = "YEAR"
)

// for schema
var durationUnitEnum []interface{}

func init() {
	var res []DurationUnit
	if err := json.Unmarshal([]byte(`["DAY","WEEK","MONTH","YEAR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		durationUnitEnum = append(durationUnitEnum, v)
	}
}

func (m DurationUnit) validateDurationUnitEnum(path, location string, value DurationUnit) error {
	if err := validate.Enum(path, location, value, durationUnitEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this duration unit
func (m DurationUnit) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDurationUnitEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}