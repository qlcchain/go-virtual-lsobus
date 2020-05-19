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

// BillingAccountRef References the billing  arrangement that a buyer has with a seller that provides products to the customer.
//
// swagger:model BillingAccountRef
type BillingAccountRef struct {

	// billing contact
	BillingContact *Contact `json:"billingContact,omitempty"`

	// Identifies the buyer’s billing account to which the recurring and non-recurring charges for this order or order item will be billed.
	// If the value ‘NEW’ is provided it means that buyer request a new BAN.
	// Required: true
	ID *string `json:"id"`
}

// Validate validates this billing account ref
func (m *BillingAccountRef) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillingContact(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BillingAccountRef) validateBillingContact(formats strfmt.Registry) error {

	if swag.IsZero(m.BillingContact) { // not required
		return nil
	}

	if m.BillingContact != nil {
		if err := m.BillingContact.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("billingContact")
			}
			return err
		}
	}

	return nil
}

func (m *BillingAccountRef) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BillingAccountRef) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BillingAccountRef) UnmarshalBinary(b []byte) error {
	var res BillingAccountRef
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}