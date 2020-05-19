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

// QuoteSummaryView Quote Summary view is provided in the response of the GET(LIST) quote. Only a subset of information are provided.
//
// swagger:model QuoteSummaryView
type QuoteSummaryView struct {

	// Indicates the base (class) type of the quote.
	AtBaseType string `json:"@baseType,omitempty"`

	// Link to the schema describing the REST resource.
	AtSchemaLocation string `json:"@schemaLocation,omitempty"`

	// Indicates the (class) type of the quote.
	AtType string `json:"@type,omitempty"`

	// Used to categorize the quote from a business perspective that can be useful for the CRM system (e.g. “enterprise”, “residential”, ...)
	Category string `json:"category,omitempty"`

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

	// ID attributed by quoting system
	ID string `json:"id,omitempty"`

	// Date when the quote was created
	// Format: date-time
	QuoteDate strfmt.DateTime `json:"quoteDate,omitempty"`

	// quote item
	// Required: true
	QuoteItem []*QuoteItem `json:"quoteItem"`

	// quote level
	QuoteLevel QuoteLevel `json:"quoteLevel,omitempty"`

	// related party role
	// Required: true
	RelatedPartyRole []*RelatedPartyRole `json:"relatedPartyRole"`

	// This is the date wished by the requester to have the quote completed (meaning priced)
	// Required: true
	// Format: date-time
	RequestedQuoteCompletionDate *strfmt.DateTime `json:"requestedQuoteCompletionDate"`

	// state
	State QuoteState `json:"state,omitempty"`

	// valid for
	ValidFor *TimePeriod `json:"validFor,omitempty"`
}

// Validate validates this quote summary view
func (m *QuoteSummaryView) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEffectiveQuoteCompletionDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpectedFulfillmentStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpectedQuoteCompletionDate(formats); err != nil {
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

	if err := m.validateRelatedPartyRole(formats); err != nil {
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

func (m *QuoteSummaryView) validateEffectiveQuoteCompletionDate(formats strfmt.Registry) error {

	if swag.IsZero(m.EffectiveQuoteCompletionDate) { // not required
		return nil
	}

	if err := validate.FormatOf("effectiveQuoteCompletionDate", "body", "date-time", m.EffectiveQuoteCompletionDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *QuoteSummaryView) validateExpectedFulfillmentStartDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ExpectedFulfillmentStartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("expectedFulfillmentStartDate", "body", "date", m.ExpectedFulfillmentStartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *QuoteSummaryView) validateExpectedQuoteCompletionDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ExpectedQuoteCompletionDate) { // not required
		return nil
	}

	if err := validate.FormatOf("expectedQuoteCompletionDate", "body", "date", m.ExpectedQuoteCompletionDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *QuoteSummaryView) validateQuoteDate(formats strfmt.Registry) error {

	if swag.IsZero(m.QuoteDate) { // not required
		return nil
	}

	if err := validate.FormatOf("quoteDate", "body", "date-time", m.QuoteDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *QuoteSummaryView) validateQuoteItem(formats strfmt.Registry) error {

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

func (m *QuoteSummaryView) validateQuoteLevel(formats strfmt.Registry) error {

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

func (m *QuoteSummaryView) validateRelatedPartyRole(formats strfmt.Registry) error {

	if err := validate.Required("relatedPartyRole", "body", m.RelatedPartyRole); err != nil {
		return err
	}

	for i := 0; i < len(m.RelatedPartyRole); i++ {
		if swag.IsZero(m.RelatedPartyRole[i]) { // not required
			continue
		}

		if m.RelatedPartyRole[i] != nil {
			if err := m.RelatedPartyRole[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("relatedPartyRole" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *QuoteSummaryView) validateRequestedQuoteCompletionDate(formats strfmt.Registry) error {

	if err := validate.Required("requestedQuoteCompletionDate", "body", m.RequestedQuoteCompletionDate); err != nil {
		return err
	}

	if err := validate.FormatOf("requestedQuoteCompletionDate", "body", "date-time", m.RequestedQuoteCompletionDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *QuoteSummaryView) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	if err := m.State.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state")
		}
		return err
	}

	return nil
}

func (m *QuoteSummaryView) validateValidFor(formats strfmt.Registry) error {

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
func (m *QuoteSummaryView) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuoteSummaryView) UnmarshalBinary(b []byte) error {
	var res QuoteSummaryView
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}