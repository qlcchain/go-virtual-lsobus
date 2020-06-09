// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProtoTerminateConnectionParam proto terminate connection param
//
// swagger:model protoTerminateConnectionParam
type ProtoTerminateConnectionParam struct {

	// dynamic param
	DynamicParam *ProtoConnectionDynamicParam `json:"dynamicParam,omitempty"`

	// product Id
	ProductID string `json:"productId,omitempty"`
}

// Validate validates this proto terminate connection param
func (m *ProtoTerminateConnectionParam) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDynamicParam(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProtoTerminateConnectionParam) validateDynamicParam(formats strfmt.Registry) error {

	if swag.IsZero(m.DynamicParam) { // not required
		return nil
	}

	if m.DynamicParam != nil {
		if err := m.DynamicParam.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dynamicParam")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProtoTerminateConnectionParam) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProtoTerminateConnectionParam) UnmarshalBinary(b []byte) error {
	var res ProtoTerminateConnectionParam
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}