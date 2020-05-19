// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Product One or more services sold to a Buyer by a Seller.  A particular Product Offering defines the technical and commercial attributes and behav-iors of a Product.
//
// swagger:model Product
type Product struct {

	// Reference of the product (url)
	Href string `json:"href,omitempty"`

	// Unique identifier of the product
	// Required: true
	ID *string `json:"id"`

	placeField []RelatedPlaceRefOrValue

	// product relationship
	ProductRelationship []*ProductRelationship `json:"productRelationship"`

	// product specification
	ProductSpecification *ProductSpecificationRef `json:"productSpecification,omitempty"`
}

// Place gets the place of this base type
func (m *Product) Place() []RelatedPlaceRefOrValue {
	return m.placeField
}

// SetPlace sets the place of this base type
func (m *Product) SetPlace(val []RelatedPlaceRefOrValue) {
	m.placeField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *Product) UnmarshalJSON(raw []byte) error {
	var data struct {
		Href string `json:"href,omitempty"`

		ID *string `json:"id"`

		Place json.RawMessage `json:"place"`

		ProductRelationship []*ProductRelationship `json:"productRelationship"`

		ProductSpecification *ProductSpecificationRef `json:"productSpecification,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var propPlace []RelatedPlaceRefOrValue
	if string(data.Place) != "null" {
		place, err := UnmarshalRelatedPlaceRefOrValueSlice(bytes.NewBuffer(data.Place), runtime.JSONConsumer())
		if err != nil && err != io.EOF {
			return err
		}
		propPlace = place
	}

	var result Product

	// href
	result.Href = data.Href

	// id
	result.ID = data.ID

	// place
	result.placeField = propPlace

	// productRelationship
	result.ProductRelationship = data.ProductRelationship

	// productSpecification
	result.ProductSpecification = data.ProductSpecification

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m Product) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		Href string `json:"href,omitempty"`

		ID *string `json:"id"`

		ProductRelationship []*ProductRelationship `json:"productRelationship"`

		ProductSpecification *ProductSpecificationRef `json:"productSpecification,omitempty"`
	}{

		Href: m.Href,

		ID: m.ID,

		ProductRelationship: m.ProductRelationship,

		ProductSpecification: m.ProductSpecification,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Place []RelatedPlaceRefOrValue `json:"place"`
	}{

		Place: m.placeField,
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this product
func (m *Product) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductRelationship(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductSpecification(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Product) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Product) validatePlace(formats strfmt.Registry) error {

	if swag.IsZero(m.Place()) { // not required
		return nil
	}

	for i := 0; i < len(m.Place()); i++ {

		if err := m.placeField[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("place" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *Product) validateProductRelationship(formats strfmt.Registry) error {

	if swag.IsZero(m.ProductRelationship) { // not required
		return nil
	}

	for i := 0; i < len(m.ProductRelationship); i++ {
		if swag.IsZero(m.ProductRelationship[i]) { // not required
			continue
		}

		if m.ProductRelationship[i] != nil {
			if err := m.ProductRelationship[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("productRelationship" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Product) validateProductSpecification(formats strfmt.Registry) error {

	if swag.IsZero(m.ProductSpecification) { // not required
		return nil
	}

	if m.ProductSpecification != nil {
		if err := m.ProductSpecification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("productSpecification")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Product) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Product) UnmarshalBinary(b []byte) error {
	var res Product
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}