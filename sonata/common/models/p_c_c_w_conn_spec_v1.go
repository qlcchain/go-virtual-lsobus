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

// PCCWConnSpecV1 ConnSpec
//
// Description of PCCW ConnSpec for Ordering
//
// swagger:model pCCWConnSpecV1
type PCCWConnSpecV1 struct {

	// The bandwidth for this service.
	Bandwidth int32 `json:"bandwidth,omitempty"`

	// class of service
	ClassOfService string `json:"classOfService,omitempty"`

	// dest company Id
	DestCompanyID string `json:"destCompanyId,omitempty"`

	// The destination location ID for this service.
	DestLocationID string `json:"destLocationId,omitempty"`

	// dest metro Id
	DestMetroID string `json:"destMetroId,omitempty"`

	// dest port Id
	DestPortID string `json:"destPortId,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// The source location ID for this service.
	SrcLocationID string `json:"srcLocationId,omitempty"`

	// src port Id
	SrcPortID string `json:"srcPortId,omitempty"`

	// started at
	// Format: date-time
	StartedAt strfmt.DateTime `json:"startedAt,omitempty"`

	// terminated at
	// Format: date-time
	TerminatedAt strfmt.DateTime `json:"terminatedAt,omitempty"`
}

// Validate validates this p c c w conn spec v1
func (m *PCCWConnSpecV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStartedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerminatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PCCWConnSpecV1) validateStartedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.StartedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("startedAt", "body", "date-time", m.StartedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PCCWConnSpecV1) validateTerminatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.TerminatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("terminatedAt", "body", "date-time", m.TerminatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PCCWConnSpecV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PCCWConnSpecV1) UnmarshalBinary(b []byte) error {
	var res PCCWConnSpecV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
