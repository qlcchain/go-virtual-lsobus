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

// Quote A quote which can be used to negotiate service and product acquisition or modification between
// a customer and a service provider
//
// swagger:model Quote
type Quote struct {

	// Indicates the base (class) type of the quote.
	AtBaseType string `json:"@baseType,omitempty"`

	// Link to the schema describing the REST resource.
	AtSchemaLocation string `json:"@schemaLocation,omitempty"`

	// Indicates the (class) type of the quote.
	AtType string `json:"@type,omitempty"`

	// agreement
	Agreement []*AgreementRef `json:"agreement"`

	// Description of the quote
	Description string `json:"description,omitempty"`

	// Date when the quoted was Cancelled or Rejected or Accepted
	// Format: date-time
	EffectiveQuoteCompletionDate strfmt.DateTime `json:"effectiveQuoteCompletionDate,omitempty"`

	// This is the date wished by the requester to have the requested quote item(s) delivered
	// Format: date
	ExpectedFulfillmentStartDate strfmt.Date `json:"expectedFulfillmentStartDate,omitempty"`

	// This is the date filled by the seller to indicate expected quote completion date.
	// Format: date
	ExpectedQuoteCompletionDate strfmt.Date `json:"expectedQuoteCompletionDate,omitempty"`

	// ID given by the consumer and only understandable by him (to facilitate his searches afterwards)
	ExternalID string `json:"externalId,omitempty"`

	// Hyperlink to access the quote
	Href string `json:"href,omitempty"`

	// Unique (within the quoting domain) identifier for the quote, as attributed by the quoting system
	ID string `json:"id,omitempty"`

	// If this flag is set to Yes, Buyer requests to have instant quoting to be provided in operation POST response
	// Required: true
	InstantSyncQuoting bool `json:"instantSyncQuoting"`

	// note
	Note []*Note `json:"note"`

	// This value MAY be assigned by the Buyer/Seller to identify a project the quoting request is associated with.
	ProjectID string `json:"projectId,omitempty"`

	// Date when the quote was created
	// Format: date-time
	QuoteDate strfmt.DateTime `json:"quoteDate,omitempty"`

	// quote item
	// Required: true
	QuoteItem []*QuoteItem `json:"quoteItem"`

	// quote level
	// Required: true
	QuoteLevel QuoteLevel `json:"quoteLevel"`

	// related party
	// Required: true
	RelatedParty []*RelatedParty `json:"relatedParty"`

	// This is the date wished by the requester to have the quote completed (meaning priced)
	// Required: true
	// Format: date-time
	RequestedQuoteCompletionDate *strfmt.DateTime `json:"requestedQuoteCompletionDate"`

	// state
	// Required: true
	State QuoteStateType `json:"state"`

	// valid for
	ValidFor *TimePeriod `json:"validFor,omitempty"`
}

// Validate validates this quote
func (m *Quote) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAgreement(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEffectiveQuoteCompletionDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpectedFulfillmentStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpectedQuoteCompletionDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstantSyncQuoting(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNote(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuoteDate(formats); err != nil {
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

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidFor(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Quote) validateAgreement(formats strfmt.Registry) error {

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

func (m *Quote) validateEffectiveQuoteCompletionDate(formats strfmt.Registry) error {

	if swag.IsZero(m.EffectiveQuoteCompletionDate) { // not required
		return nil
	}

	if err := validate.FormatOf("effectiveQuoteCompletionDate", "body", "date-time", m.EffectiveQuoteCompletionDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Quote) validateExpectedFulfillmentStartDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ExpectedFulfillmentStartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("expectedFulfillmentStartDate", "body", "date", m.ExpectedFulfillmentStartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Quote) validateExpectedQuoteCompletionDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ExpectedQuoteCompletionDate) { // not required
		return nil
	}

	if err := validate.FormatOf("expectedQuoteCompletionDate", "body", "date", m.ExpectedQuoteCompletionDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Quote) validateInstantSyncQuoting(formats strfmt.Registry) error {

	if err := validate.Required("instantSyncQuoting", "body", bool(m.InstantSyncQuoting)); err != nil {
		return err
	}

	return nil
}

func (m *Quote) validateNote(formats strfmt.Registry) error {

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

func (m *Quote) validateQuoteDate(formats strfmt.Registry) error {

	if swag.IsZero(m.QuoteDate) { // not required
		return nil
	}

	if err := validate.FormatOf("quoteDate", "body", "date-time", m.QuoteDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Quote) validateQuoteItem(formats strfmt.Registry) error {

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

func (m *Quote) validateQuoteLevel(formats strfmt.Registry) error {

	if err := m.QuoteLevel.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("quoteLevel")
		}
		return err
	}

	return nil
}

func (m *Quote) validateRelatedParty(formats strfmt.Registry) error {

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

func (m *Quote) validateRequestedQuoteCompletionDate(formats strfmt.Registry) error {

	if err := validate.Required("requestedQuoteCompletionDate", "body", m.RequestedQuoteCompletionDate); err != nil {
		return err
	}

	if err := validate.FormatOf("requestedQuoteCompletionDate", "body", "date-time", m.RequestedQuoteCompletionDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Quote) validateState(formats strfmt.Registry) error {

	if err := m.State.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state")
		}
		return err
	}

	return nil
}

func (m *Quote) validateValidFor(formats strfmt.Registry) error {

	if swag.IsZero(m.ValidFor) { // not required
		return nil
	}

	if m.ValidFor != nil {
		if err := m.ValidFor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("validFor")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Quote) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Quote) UnmarshalBinary(b []byte) error {
	var res Quote
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}