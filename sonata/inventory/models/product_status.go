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

// ProductStatus The state of the product in accordance with the product lifecycle.
//
// swagger:model ProductStatus
type ProductStatus string

const (

	// ProductStatusActive captures enum value "active"
	ProductStatusActive ProductStatus = "active"

	// ProductStatusSuspended captures enum value "suspended"
	ProductStatusSuspended ProductStatus = "suspended"

	// ProductStatusActivePendingTerminate captures enum value "activePendingTerminate"
	ProductStatusActivePendingTerminate ProductStatus = "activePendingTerminate"

	// ProductStatusTerminated captures enum value "terminated"
	ProductStatusTerminated ProductStatus = "terminated"

	// ProductStatusPendingActive captures enum value "pendingActive"
	ProductStatusPendingActive ProductStatus = "pendingActive"

	// ProductStatusSuspendedPendingTerminate captures enum value "suspendedPendingTerminate"
	ProductStatusSuspendedPendingTerminate ProductStatus = "suspendedPendingTerminate"
)

// for schema
var productStatusEnum []interface{}

func init() {
	var res []ProductStatus
	if err := json.Unmarshal([]byte(`["active","suspended","activePendingTerminate","terminated","pendingActive","suspendedPendingTerminate"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		productStatusEnum = append(productStatusEnum, v)
	}
}

func (m ProductStatus) validateProductStatusEnum(path, location string, value ProductStatus) error {
	if err := validate.Enum(path, location, value, productStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this product status
func (m ProductStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateProductStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
