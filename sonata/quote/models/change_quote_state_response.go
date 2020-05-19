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

// ChangeQuoteStateResponse Structure to response for a quote cancellation/rejection request
//
// swagger:model ChangeQuoteStateResponse
type ChangeQuoteStateResponse struct {

	// External Id of the quote (buyer quote id). If provided by seller this information is send back in the response.
	ExternalID string `json:"externalId,omitempty"`

	// id of the quote to be cancelled
	// Required: true
	ID *string `json:"id"`

	// Cancellation or rejection date (effective)
	// Required: true
	// Format: date-time
	QuoteEffectiveStateChangeDate *strfmt.DateTime `json:"quoteEffectiveStateChangeDate"`

	// state
	// Required: true
	State QuoteStateType `json:"state"`
}

// Validate validates this change quote state response
func (m *ChangeQuoteStateResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuoteEffectiveStateChangeDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChangeQuoteStateResponse) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ChangeQuoteStateResponse) validateQuoteEffectiveStateChangeDate(formats strfmt.Registry) error {

	if err := validate.Required("quoteEffectiveStateChangeDate", "body", m.QuoteEffectiveStateChangeDate); err != nil {
		return err
	}

	if err := validate.FormatOf("quoteEffectiveStateChangeDate", "body", "date-time", m.QuoteEffectiveStateChangeDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ChangeQuoteStateResponse) validateState(formats strfmt.Registry) error {

	if err := m.State.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChangeQuoteStateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChangeQuoteStateResponse) UnmarshalBinary(b []byte) error {
	var res ChangeQuoteStateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}