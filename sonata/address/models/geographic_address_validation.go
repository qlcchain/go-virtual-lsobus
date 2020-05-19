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

// GeographicAddressValidation Resource used to support a request for a validation address : check if a geographic address described by values attributes exists)
// If exist, id is provided
// if not, alternate address(es) could be provided
//
// swagger:model GeographicAddressValidation
type GeographicAddressValidation struct {

	// Unique identifier of the Address Validation (Not an address id !!)
	ID string `json:"id,omitempty"`

	// requested address
	// Required: true
	RequestedAddress *GeographicAddressRequestBuyerInput `json:"requestedAddress"`

	// Date when the address validation is performed
	// Format: date-time
	ValidationDate strfmt.DateTime `json:"validationDate,omitempty"`

	// validation result
	ValidationResult ValidationResult `json:"validationResult,omitempty"`

	// verified address
	VerifiedAddress []*GeographicAddressSellerResponse `json:"verifiedAddress"`
}

// Validate validates this geographic address validation
func (m *GeographicAddressValidation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRequestedAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidationResult(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVerifiedAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GeographicAddressValidation) validateRequestedAddress(formats strfmt.Registry) error {

	if err := validate.Required("requestedAddress", "body", m.RequestedAddress); err != nil {
		return err
	}

	if m.RequestedAddress != nil {
		if err := m.RequestedAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requestedAddress")
			}
			return err
		}
	}

	return nil
}

func (m *GeographicAddressValidation) validateValidationDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ValidationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("validationDate", "body", "date-time", m.ValidationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GeographicAddressValidation) validateValidationResult(formats strfmt.Registry) error {

	if swag.IsZero(m.ValidationResult) { // not required
		return nil
	}

	if err := m.ValidationResult.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("validationResult")
		}
		return err
	}

	return nil
}

func (m *GeographicAddressValidation) validateVerifiedAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.VerifiedAddress) { // not required
		return nil
	}

	for i := 0; i < len(m.VerifiedAddress); i++ {
		if swag.IsZero(m.VerifiedAddress[i]) { // not required
			continue
		}

		if m.VerifiedAddress[i] != nil {
			if err := m.VerifiedAddress[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("verifiedAddress" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GeographicAddressValidation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GeographicAddressValidation) UnmarshalBinary(b []byte) error {
	var res GeographicAddressValidation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}