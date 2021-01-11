// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GeographicSiteFindResp Technical structure to provide site list retrieve in the GET operation
//
// swagger:model GeographicSiteFindResp
type GeographicSiteFindResp struct {

	// geographic address
	GeographicAddress *GeographicAddressFindResp `json:"geographicAddress,omitempty"`

	// id of the site. This is coud be used in other API as a place
	ID string `json:"id,omitempty"`

	// The name of the company that is the administrative authority (e.g. controls access) for this Service Site. (For example, the building owner)
	SiteCompanyName string `json:"siteCompanyName,omitempty"`

	// Name of the relatedParty which has role 'Site Contact' in the Site record.
	SiteContactName string `json:"siteContactName,omitempty"`

	// The name of the company that is the administrative authority for the space within this Service Site. (For example, the company leasing space in a multi-tenant building).
	SiteCustomerName string `json:"siteCustomerName,omitempty"`

	// A textual description of the Service Site.
	SiteDescription string `json:"siteDescription,omitempty"`

	// A name commonly used by people to refer to this Service Site.
	SiteName string `json:"siteName,omitempty"`

	// This defines whether a Service Site is public or private.
	SiteType string `json:"siteType,omitempty"`

	// status
	Status Status `json:"status,omitempty"`
}

// Validate validates this geographic site find resp
func (m *GeographicSiteFindResp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGeographicAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GeographicSiteFindResp) validateGeographicAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.GeographicAddress) { // not required
		return nil
	}

	if m.GeographicAddress != nil {
		if err := m.GeographicAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("geographicAddress")
			}
			return err
		}
	}

	return nil
}

func (m *GeographicSiteFindResp) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GeographicSiteFindResp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GeographicSiteFindResp) UnmarshalBinary(b []byte) error {
	var res GeographicSiteFindResp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}