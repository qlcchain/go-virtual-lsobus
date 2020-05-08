// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// RelationshipType Indicates the type of relationship between products.
//
// swagger:model RelationshipType
type RelationshipType string

const (

	// RelationshipTypeReliesOn captures enum value "reliesOn"
	RelationshipTypeReliesOn RelationshipType = "reliesOn"

	// RelationshipTypeBundled captures enum value "bundled"
	RelationshipTypeBundled RelationshipType = "bundled"

	// RelationshipTypeComesFrom captures enum value "comesFrom"
	RelationshipTypeComesFrom RelationshipType = "comesFrom"
)

// for schema
var relationshipTypeEnum []interface{}

func init() {
	var res []RelationshipType
	if err := json.Unmarshal([]byte(`["reliesOn","bundled","comesFrom"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		relationshipTypeEnum = append(relationshipTypeEnum, v)
	}
}

func (m RelationshipType) validateRelationshipTypeEnum(path, location string, value RelationshipType) error {
	if err := validate.Enum(path, location, value, relationshipTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this relationship type
func (m RelationshipType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateRelationshipTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
