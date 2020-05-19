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

// CancelProductOrder Structure to allow Seller to answer to a cancel order request
//
// swagger:model CancelProductOrder
type CancelProductOrder struct {

	// Technical attribute to extend this class.
	AtSchemaLocation string `json:"@schemaLocation,omitempty"`

	// Technical attribute to extend this class.
	AtType string `json:"@type,omitempty"`

	// If seller denied cancellation request he could here provide reason for this denial
	CancellationDeniedReason string `json:"cancellationDeniedReason,omitempty"`

	// An optional free-form text field for the Seller to provide additional information regarding the reason for the cancellation.
	CancellationReason string `json:"cancellationReason,omitempty"`

	// Hyperlink to access order cancellation request.This is not a product order href.
	Href string `json:"href,omitempty"`

	// Unique identifier for the order cancellation request that is generated by the Seller when the order cancellation is accepted via an API. This is not the order id.
	// Required: true
	ID *string `json:"id"`

	// product order
	// Required: true
	ProductOrder *ProductOrderRefCancel `json:"productOrder"`

	// Identifies the date the Seller cancelled the Order.
	// Required: true
	// Format: date-time
	RequestedCancellationDate *strfmt.DateTime `json:"requestedCancellationDate"`

	// state
	// Required: true
	State TaskStateType `json:"state"`
}

// Validate validates this cancel product order
func (m *CancelProductOrder) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductOrder(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedCancellationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CancelProductOrder) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *CancelProductOrder) validateProductOrder(formats strfmt.Registry) error {

	if err := validate.Required("productOrder", "body", m.ProductOrder); err != nil {
		return err
	}

	if m.ProductOrder != nil {
		if err := m.ProductOrder.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("productOrder")
			}
			return err
		}
	}

	return nil
}

func (m *CancelProductOrder) validateRequestedCancellationDate(formats strfmt.Registry) error {

	if err := validate.Required("requestedCancellationDate", "body", m.RequestedCancellationDate); err != nil {
		return err
	}

	if err := validate.FormatOf("requestedCancellationDate", "body", "date-time", m.RequestedCancellationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CancelProductOrder) validateState(formats strfmt.Registry) error {

	if err := m.State.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CancelProductOrder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CancelProductOrder) UnmarshalBinary(b []byte) error {
	var res CancelProductOrder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}