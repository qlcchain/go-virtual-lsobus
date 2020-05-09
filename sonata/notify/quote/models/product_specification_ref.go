// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProductSpecificationRef A ProductSpecification is a detailed description of a tangible or intangible object made available externally in the form of a ProductOffering to customers or other parties playing a party role.
//
// swagger:model ProductSpecificationRef
type ProductSpecificationRef struct {
	describingField Describing

	// Reference of the product specification
	Href string `json:"href,omitempty"`

	// Unique identifier of the product specification
	ID string `json:"id,omitempty"`

	// Name of the product specification
	Name string `json:"name,omitempty"`

	// Version of the product specification
	Version string `json:"version,omitempty"`
}

// Describing gets the describing of this base type
func (m *ProductSpecificationRef) Describing() Describing {
	return m.describingField
}

// SetDescribing sets the describing of this base type
func (m *ProductSpecificationRef) SetDescribing(val Describing) {
	m.describingField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *ProductSpecificationRef) UnmarshalJSON(raw []byte) error {
	var data struct {
		Describing json.RawMessage `json:"describing,omitempty"`

		Href string `json:"href,omitempty"`

		ID string `json:"id,omitempty"`

		Name string `json:"name,omitempty"`

		Version string `json:"version,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var propDescribing Describing
	if string(data.Describing) != "null" {
		describing, err := UnmarshalDescribing(bytes.NewBuffer(data.Describing), runtime.JSONConsumer())
		if err != nil && err != io.EOF {
			return err
		}
		propDescribing = describing
	}

	var result ProductSpecificationRef

	// describing
	result.describingField = propDescribing

	// href
	result.Href = data.Href

	// id
	result.ID = data.ID

	// name
	result.Name = data.Name

	// version
	result.Version = data.Version

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m ProductSpecificationRef) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		Href string `json:"href,omitempty"`

		ID string `json:"id,omitempty"`

		Name string `json:"name,omitempty"`

		Version string `json:"version,omitempty"`
	}{

		Href: m.Href,

		ID: m.ID,

		Name: m.Name,

		Version: m.Version,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Describing Describing `json:"describing,omitempty"`
	}{

		Describing: m.describingField,
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this product specification ref
func (m *ProductSpecificationRef) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescribing(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProductSpecificationRef) validateDescribing(formats strfmt.Registry) error {

	if swag.IsZero(m.Describing()) { // not required
		return nil
	}

	if err := m.Describing().Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("describing")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProductSpecificationRef) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProductSpecificationRef) UnmarshalBinary(b []byte) error {
	var res ProductSpecificationRef
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
