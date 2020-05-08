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

// ProductOfferingQualificationItemStateType POQ item states - The specific states are managed by the Seller based on its processing and/or based on Buyer's action.
//
// swagger:model ProductOfferingQualificationItemStateType
type ProductOfferingQualificationItemStateType string

const (

	// ProductOfferingQualificationItemStateTypeDone captures enum value "done"
	ProductOfferingQualificationItemStateTypeDone ProductOfferingQualificationItemStateType = "done"

	// ProductOfferingQualificationItemStateTypeTerminatedWithErrorUnableToProvide captures enum value "terminatedWithError.unableToProvide"
	ProductOfferingQualificationItemStateTypeTerminatedWithErrorUnableToProvide ProductOfferingQualificationItemStateType = "terminatedWithError.unableToProvide"

	// ProductOfferingQualificationItemStateTypeTerminatedWithErrorInsufficientInformationProvided captures enum value "terminatedWithError.insufficientInformationProvided"
	ProductOfferingQualificationItemStateTypeTerminatedWithErrorInsufficientInformationProvided ProductOfferingQualificationItemStateType = "terminatedWithError.insufficientInformationProvided"

	// ProductOfferingQualificationItemStateTypeInProgress captures enum value "inProgress"
	ProductOfferingQualificationItemStateTypeInProgress ProductOfferingQualificationItemStateType = "inProgress"
)

// for schema
var productOfferingQualificationItemStateTypeEnum []interface{}

func init() {
	var res []ProductOfferingQualificationItemStateType
	if err := json.Unmarshal([]byte(`["done","terminatedWithError.unableToProvide","terminatedWithError.insufficientInformationProvided","inProgress"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		productOfferingQualificationItemStateTypeEnum = append(productOfferingQualificationItemStateTypeEnum, v)
	}
}

func (m ProductOfferingQualificationItemStateType) validateProductOfferingQualificationItemStateTypeEnum(path, location string, value ProductOfferingQualificationItemStateType) error {
	if err := validate.Enum(path, location, value, productOfferingQualificationItemStateTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this product offering qualification item state type
func (m ProductOfferingQualificationItemStateType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateProductOfferingQualificationItemStateTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
