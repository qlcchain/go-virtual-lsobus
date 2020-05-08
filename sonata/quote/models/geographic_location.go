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

// GeographicLocation A set of coordinates (typically including latitude and longitude) that describes a particular loca-tion on earth.
//
// swagger:model GeographicLocation
type GeographicLocation struct {
	hrefField string

	idField string

	roleField *string

	// geographic point
	// Required: true
	GeographicPoint *GeographicPoint `json:"geographicPoint"`

	// The spatial reference system used to determine the coordinates
	// Required: true
	SpatialRef *string `json:"spatialRef"`
}

// AtType gets the at type of this subtype
func (m *GeographicLocation) AtType() string {
	return "GeographicLocation"
}

// SetAtType sets the at type of this subtype
func (m *GeographicLocation) SetAtType(val string) {
}

// Href gets the href of this subtype
func (m *GeographicLocation) Href() string {
	return m.hrefField
}

// SetHref sets the href of this subtype
func (m *GeographicLocation) SetHref(val string) {
	m.hrefField = val
}

// ID gets the id of this subtype
func (m *GeographicLocation) ID() string {
	return m.idField
}

// SetID sets the id of this subtype
func (m *GeographicLocation) SetID(val string) {
	m.idField = val
}

// Role gets the role of this subtype
func (m *GeographicLocation) Role() *string {
	return m.roleField
}

// SetRole sets the role of this subtype
func (m *GeographicLocation) SetRole(val *string) {
	m.roleField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *GeographicLocation) UnmarshalJSON(raw []byte) error {
	var data struct {

		// geographic point
		// Required: true
		GeographicPoint *GeographicPoint `json:"geographicPoint"`

		// The spatial reference system used to determine the coordinates
		// Required: true
		SpatialRef *string `json:"spatialRef"`
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

	var result GeographicLocation

	if base.AtType != result.AtType() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid @type value: %q", base.AtType)
	}
	result.hrefField = base.Href

	result.idField = base.ID

	result.roleField = base.Role

	result.GeographicPoint = data.GeographicPoint
	result.SpatialRef = data.SpatialRef

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m GeographicLocation) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// geographic point
		// Required: true
		GeographicPoint *GeographicPoint `json:"geographicPoint"`

		// The spatial reference system used to determine the coordinates
		// Required: true
		SpatialRef *string `json:"spatialRef"`
	}{

		GeographicPoint: m.GeographicPoint,

		SpatialRef: m.SpatialRef,
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

// Validate validates this geographic location
func (m *GeographicLocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGeographicPoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpatialRef(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GeographicLocation) validateRole(formats strfmt.Registry) error {

	if err := validate.Required("role", "body", m.Role()); err != nil {
		return err
	}

	return nil
}

func (m *GeographicLocation) validateGeographicPoint(formats strfmt.Registry) error {

	if err := validate.Required("geographicPoint", "body", m.GeographicPoint); err != nil {
		return err
	}

	if m.GeographicPoint != nil {
		if err := m.GeographicPoint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("geographicPoint")
			}
			return err
		}
	}

	return nil
}

func (m *GeographicLocation) validateSpatialRef(formats strfmt.Registry) error {

	if err := validate.Required("spatialRef", "body", m.SpatialRef); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GeographicLocation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GeographicLocation) UnmarshalBinary(b []byte) error {
	var res GeographicLocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
