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

// ProductRef Targeted existing product used in product relationship description.
//
// swagger:model ProductRef
type ProductRef struct {

	// Targeted Buyer product id - Informative.
	BuyerProductID string `json:"buyerProductId,omitempty"`

	// Hyperlink to the product in the inventory
	Href string `json:"href,omitempty"`

	// Targeted Seller product id
	// Required: true
	ID *string `json:"id"`
}

// Validate validates this product ref
func (m *ProductRef) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProductRef) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProductRef) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProductRef) UnmarshalBinary(b []byte) error {
	var res ProductRef
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
