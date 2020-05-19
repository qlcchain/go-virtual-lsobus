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

// OrderItemRelationShip This class allows the ability to associate one order item to another order item.
//
// swagger:model OrderItemRelationShip
type OrderItemRelationShip struct {

	// The id of the targeted order item by the relationship
	// Required: true
	ID *string `json:"id"`

	// Id of another product order if this relationship is between two distinct orders.
	ProductOrderID string `json:"productOrderId,omitempty"`

	// Indicates the type of order item relationship
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this order item relation ship
func (m *OrderItemRelationShip) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
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

func (m *OrderItemRelationShip) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *OrderItemRelationShip) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrderItemRelationShip) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderItemRelationShip) UnmarshalBinary(b []byte) error {
	var res OrderItemRelationShip
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}