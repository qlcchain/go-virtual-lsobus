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

// TaskStateType task state type
//
// swagger:model TaskStateType
type TaskStateType string

const (

	// TaskStateTypeAcknowledged captures enum value "acknowledged"
	TaskStateTypeAcknowledged TaskStateType = "acknowledged"

	// TaskStateTypeInProgress captures enum value "inProgress"
	TaskStateTypeInProgress TaskStateType = "inProgress"

	// TaskStateTypeDone captures enum value "done"
	TaskStateTypeDone TaskStateType = "done"

	// TaskStateTypeTerminatedWithError captures enum value "terminatedWithError"
	TaskStateTypeTerminatedWithError TaskStateType = "terminatedWithError"
)

// for schema
var taskStateTypeEnum []interface{}

func init() {
	var res []TaskStateType
	if err := json.Unmarshal([]byte(`["acknowledged","inProgress","done","terminatedWithError"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		taskStateTypeEnum = append(taskStateTypeEnum, v)
	}
}

func (m TaskStateType) validateTaskStateTypeEnum(path, location string, value TaskStateType) error {
	if err := validate.Enum(path, location, value, taskStateTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this task state type
func (m TaskStateType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTaskStateTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
