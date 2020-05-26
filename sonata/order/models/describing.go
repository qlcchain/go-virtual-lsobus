// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"

	cmnmod "github.com/qlcchain/go-lsobus/sonata/common/models"
)

// Describing Polymorphic class to describe request product attribute
//
// swagger:discriminator Describing @type
type Describing interface {
	runtime.Validatable

	// URL targeting where product description is stored
	AtSchemaLocation() string
	SetAtSchemaLocation(string)

	// Type of the product
	AtType() string
	SetAtType(string)

	// AdditionalProperties in base type shoud be handled just like regular properties
	// At this moment, the base type property is pushed down to the subtype
}

type describing struct {
	atSchemaLocationField string

	atTypeField string
}

// AtSchemaLocation gets the at schema location of this polymorphic type
func (m *describing) AtSchemaLocation() string {
	return m.atSchemaLocationField
}

// SetAtSchemaLocation sets the at schema location of this polymorphic type
func (m *describing) SetAtSchemaLocation(val string) {
	m.atSchemaLocationField = val
}

// AtType gets the at type of this polymorphic type
func (m *describing) AtType() string {
	return "Describing"
}

// SetAtType sets the at type of this polymorphic type
func (m *describing) SetAtType(val string) {
}

// UnmarshalDescribingSlice unmarshals polymorphic slices of Describing
func UnmarshalDescribingSlice(reader io.Reader, consumer runtime.Consumer) ([]Describing, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []Describing
	for _, element := range elements {
		obj, err := unmarshalDescribing(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalDescribing unmarshals polymorphic Describing
func UnmarshalDescribing(reader io.Reader, consumer runtime.Consumer) (Describing, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalDescribing(data, consumer)
}

func unmarshalDescribing(data []byte, consumer runtime.Consumer) (Describing, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the @type property.
	var getType struct {
		AtType string `json:"@type"`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("@type", "body", getType.AtType); err != nil {
		return nil, err
	}

	// The value of @type is used to determine which type to create and unmarshal the data into
	switch getType.AtType {
	case "Describing":
		var result describing
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "ELineSpec":
		var result cmnmod.ELineSpec
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "PCCWConnSpec":
		var result cmnmod.PCCWConnSpec
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "UNISpec":
		var result cmnmod.UNISpec
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, errors.New(422, "invalid @type value: %q", getType.AtType)
}

// Validate validates this describing
func (m *describing) Validate(formats strfmt.Registry) error {
	return nil
}
