// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// QuoteCreate This structure is used only in the POST operation for the request.
//
// swagger:model Quote_Create
type QuoteCreate struct {

	// Indicates the base (class) type of the quote.
	AtBaseType string `json:"@baseType,omitempty"`

	// Link to the schema describing the REST resource.
	// Technical attribute to extend this class.
	AtSchemaLocation string `json:"@schemaLocation,omitempty"`

	// Indicates the (class) type of the quote.
	// Technical attribute to extend this class.
	AtType string `json:"@type,omitempty"`

	// agreement
	Agreement []*AgreementRef `json:"agreement"`

	// Description of the quote
	Description string `json:"description,omitempty"`

	// This is the date wished by the requester to have the requested quote item(s) delivered
	// Format: date
	ExpectedFulfillmentStartDate strfmt.Date `json:"expectedFulfillmentStartDate,omitempty"`

	// ID given by the consumer and only understandable by him (to facilitate his searches afterwards)
	ExternalID string `json:"externalId,omitempty"`

	// If this flag is set to Yes, Buyer requests to have instant quoting to be provided in operation POST response
	InstantSyncQuoting *bool `json:"instantSyncQuoting,omitempty"`

	// note
	Note []*Note `json:"note"`

	// This value MAY be assigned by the Buyer/Seller to identify a project the quoting request is associated with.
	ProjectID string `json:"projectId,omitempty"`

	// quote item
	// Required: true
	QuoteItem []*QuoteItemCreate `json:"quoteItem"`

	// quote level
	QuoteLevel QuoteLevel `json:"quoteLevel,omitempty"`

	// related party
	// Required: true
	RelatedParty []*RelatedParty `json:"relatedParty"`

	// This is the date wished by the requester to have the quote completed (meaning priced).
	// This attribute is not considered when instantSyncQuoting is set to Yes.
	// Required: true
	// Format: date-time
	RequestedQuoteCompletionDate *strfmt.DateTime `json:"requestedQuoteCompletionDate"`
}

// Validate validates this quote create
func (m *QuoteCreate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAgreement(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpectedFulfillmentStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNote(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuoteItem(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuoteLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelatedParty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedQuoteCompletionDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuoteCreate) validateAgreement(formats strfmt.Registry) error {

	if swag.IsZero(m.Agreement) { // not required
		return nil
	}

	for i := 0; i < len(m.Agreement); i++ {
		if swag.IsZero(m.Agreement[i]) { // not required
			continue
		}

		if m.Agreement[i] != nil {
			if err := m.Agreement[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("agreement" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *QuoteCreate) validateExpectedFulfillmentStartDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ExpectedFulfillmentStartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("expectedFulfillmentStartDate", "body", "date", m.ExpectedFulfillmentStartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *QuoteCreate) validateNote(formats strfmt.Registry) error {

	if swag.IsZero(m.Note) { // not required
		return nil
	}

	for i := 0; i < len(m.Note); i++ {
		if swag.IsZero(m.Note[i]) { // not required
			continue
		}

		if m.Note[i] != nil {
			if err := m.Note[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("note" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *QuoteCreate) validateQuoteItem(formats strfmt.Registry) error {

	if err := validate.Required("quoteItem", "body", m.QuoteItem); err != nil {
		return err
	}

	for i := 0; i < len(m.QuoteItem); i++ {
		if swag.IsZero(m.QuoteItem[i]) { // not required
			continue
		}

		if m.QuoteItem[i] != nil {
			if err := m.QuoteItem[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("quoteItem" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *QuoteCreate) validateQuoteLevel(formats strfmt.Registry) error {

	if swag.IsZero(m.QuoteLevel) { // not required
		return nil
	}

	if err := m.QuoteLevel.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("quoteLevel")
		}
		return err
	}

	return nil
}

func (m *QuoteCreate) validateRelatedParty(formats strfmt.Registry) error {

	if err := validate.Required("relatedParty", "body", m.RelatedParty); err != nil {
		return err
	}

	for i := 0; i < len(m.RelatedParty); i++ {
		if swag.IsZero(m.RelatedParty[i]) { // not required
			continue
		}

		if m.RelatedParty[i] != nil {
			if err := m.RelatedParty[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("relatedParty" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *QuoteCreate) validateRequestedQuoteCompletionDate(formats strfmt.Registry) error {

	if err := validate.Required("requestedQuoteCompletionDate", "body", m.RequestedQuoteCompletionDate); err != nil {
		return err
	}

	if err := validate.FormatOf("requestedQuoteCompletionDate", "body", "date-time", m.RequestedQuoteCompletionDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QuoteCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuoteCreate) UnmarshalBinary(b []byte) error {
	var res QuoteCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
