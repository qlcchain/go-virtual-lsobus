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

// ReferencedAddress A globally unique identifier controlled by a generally accepted independent administrative authority that specifies a fixed geographical location.
//
// swagger:model ReferencedAddress
type ReferencedAddress struct {

	// A reference to an address by id; this would include such things as CLLI Codes, or Seller-assigned address identifiers
	// Required: true
	ReferenceID *string `json:"referenceId"`

	// The type of the reference. For North American providers this would normally be CLLI
	// Required: true
	ReferenceType *string `json:"referenceType"`
}

// Validate validates this referenced address
func (m *ReferencedAddress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReferenceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReferenceType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReferencedAddress) validateReferenceID(formats strfmt.Registry) error {

	if err := validate.Required("referenceId", "body", m.ReferenceID); err != nil {
		return err
	}

	return nil
}

func (m *ReferencedAddress) validateReferenceType(formats strfmt.Registry) error {

	if err := validate.Required("referenceType", "body", m.ReferenceType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReferencedAddress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReferencedAddress) UnmarshalBinary(b []byte) error {
	var res ReferencedAddress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}