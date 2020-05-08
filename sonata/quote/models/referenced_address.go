// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ReferencedAddress A globally unique identifier controlled by a generally accepted independent administrative authority that specifies a fixed geographical location.
//
// swagger:model ReferencedAddress
type ReferencedAddress struct {
	hrefField string

	idField string

	roleField *string

	// A reference to an address by id; this would include such things as CLLI Codes, or Seller-assigned address identifiers
	// Required: true
	ReferenceID *string `json:"referenceId"`

	// The type of the reference. For North American providers this would normally be CLLI
	// Required: true
	ReferenceType *string `json:"referenceType"`
}

// AtType gets the at type of this subtype
func (m *ReferencedAddress) AtType() string {
	return "ReferencedAddress"
}

// SetAtType sets the at type of this subtype
func (m *ReferencedAddress) SetAtType(val string) {
}

// Href gets the href of this subtype
func (m *ReferencedAddress) Href() string {
	return m.hrefField
}

// SetHref sets the href of this subtype
func (m *ReferencedAddress) SetHref(val string) {
	m.hrefField = val
}

// ID gets the id of this subtype
func (m *ReferencedAddress) ID() string {
	return m.idField
}

// SetID sets the id of this subtype
func (m *ReferencedAddress) SetID(val string) {
	m.idField = val
}

// Role gets the role of this subtype
func (m *ReferencedAddress) Role() *string {
	return m.roleField
}

// SetRole sets the role of this subtype
func (m *ReferencedAddress) SetRole(val *string) {
	m.roleField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *ReferencedAddress) UnmarshalJSON(raw []byte) error {
	var data struct {

		// A reference to an address by id; this would include such things as CLLI Codes, or Seller-assigned address identifiers
		// Required: true
		ReferenceID *string `json:"referenceId"`

		// The type of the reference. For North American providers this would normally be CLLI
		// Required: true
		ReferenceType *string `json:"referenceType"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		AtType string `json:"@type,omitempty"`

		Href string `json:"href,omitempty"`

		ID string `json:"id,omitempty"`

		Role *string `json:"role"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result ReferencedAddress

	if base.AtType != result.AtType() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid @type value: %q", base.AtType)
	}
	result.hrefField = base.Href

	result.idField = base.ID

	result.roleField = base.Role

	result.ReferenceID = data.ReferenceID
	result.ReferenceType = data.ReferenceType

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m ReferencedAddress) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// A reference to an address by id; this would include such things as CLLI Codes, or Seller-assigned address identifiers
		// Required: true
		ReferenceID *string `json:"referenceId"`

		// The type of the reference. For North American providers this would normally be CLLI
		// Required: true
		ReferenceType *string `json:"referenceType"`
	}{

		ReferenceID: m.ReferenceID,

		ReferenceType: m.ReferenceType,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		AtType string `json:"@type,omitempty"`

		Href string `json:"href,omitempty"`

		ID string `json:"id,omitempty"`

		Role *string `json:"role"`
	}{

		AtType: m.AtType(),

		Href: m.Href(),

		ID: m.ID(),

		Role: m.Role(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this referenced address
func (m *ReferencedAddress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

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

func (m *ReferencedAddress) validateRole(formats strfmt.Registry) error {

	if err := validate.Required("role", "body", m.Role()); err != nil {
		return err
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
