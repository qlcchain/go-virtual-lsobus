// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProtoContractPrivacyParam proto contract privacy param
//
// swagger:model protoContractPrivacyParam
type ProtoContractPrivacyParam struct {

	// private for
	PrivateFor []string `json:"privateFor"`

	// private from
	PrivateFrom string `json:"privateFrom,omitempty"`

	// private group ID
	PrivateGroupID string `json:"privateGroupID,omitempty"`
}

// Validate validates this proto contract privacy param
func (m *ProtoContractPrivacyParam) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProtoContractPrivacyParam) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProtoContractPrivacyParam) UnmarshalBinary(b []byte) error {
	var res ProtoContractPrivacyParam
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}