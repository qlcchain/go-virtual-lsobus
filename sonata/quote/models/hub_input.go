// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HubInput This class is used to request a HUB creation - Used in the POST operation.
//
// swagger:model HubInput
type HubInput struct {

	// This attribute is the callback url where event will be sent when occurs, for instance an url http://yourdomain/listener/api/v1/event
	// Required: true
	Callback *string `json:"callback"`

	// This attribute is used to define notification type. Example: ","query":”eventType = QuoteStateChangeNotification”}
	// Required: true
	Query *string `json:"query"`
}

// Validate validates this hub input
func (m *HubInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCallback(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuery(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HubInput) validateCallback(formats strfmt.Registry) error {

	if err := validate.Required("callback", "body", m.Callback); err != nil {
		return err
	}

	return nil
}

func (m *HubInput) validateQuery(formats strfmt.Registry) error {

	if err := validate.Required("query", "body", m.Query); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HubInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HubInput) UnmarshalBinary(b []byte) error {
	var res HubInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
