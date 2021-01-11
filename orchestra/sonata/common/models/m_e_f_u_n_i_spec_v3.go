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

// MEFUNISpecV3 UNISpec
//
// Description of UNISpec for Ordering- source 57.1
//
// swagger:model mEFUNISpecV3
type MEFUNISpecV3 struct {

	// All-to-One Bundling can be either Enabled or Disabled. A value of True would equate to Enabled
	AllToOneBundling *bool `json:"allToOneBundling,omitempty"`

	// Link OAM enables the operator to monitor and troubleshoot a single Ethernet link.
	LinkOamEnabled *bool `json:"linkOamEnabled,omitempty"`

	// Indicates the maximum service frame size that can be reliably processed at the UNI level
	// Minimum: 1522
	MaxServiceFrameSize int32 `json:"maxServiceFrameSize,omitempty"`

	// A UNI can be implemented with one or more physical links. This attribute specifies the number of such links.
	NumberOfLinks int32 `json:"numberOfLinks,omitempty"`

	// This attribute is a list of physical layers, one for each physical link implementing the UNI
	// Required: true
	PhysicalLayer []PhysicalLayer `json:"physicalLayer"`

	// The Synchronous Mode Service Attribute is a list with one item for each of the physical links implementing the UNI. When the value of an item is "Enabled," the bits transmitted from the CEN to the CE on the physical link corresponding to the item can be used by the CE as a bit clock reference
	SynchronousMode []*SynchronousMode `json:"synchronousMode"`

	// Identifies whether a given UNI is capable of sharing tokens across Bandwidth Profile Flows in an Envelope.  The default is disabled ("false").
	TokenShareEnabled *bool `json:"tokenShareEnabled,omitempty"`

	// Indicates the desire for management of the customer interface. For ordering, this is an optional parameter that may be requested
	UniEthernetLmiEnabled *bool `json:"uniEthernetLmiEnabled,omitempty"`

	// uni l2cp address set
	UniL2cpAddressSet L2CP `json:"uniL2cpAddressSet,omitempty"`

	// Can be either an empty list, or a list of entries identifying protocols to be peered where each entry consists of {Destination Address, Protocol Identifier} or {Destination Address, Protocol Identifier, Link Identifier}.
	UniL2cpPeering []*UniL2cpPeering `json:"uniL2cpPeering"`

	// Request for monitoring of a UNI.
	UniMegEnabled *bool `json:"uniMegEnabled,omitempty"`

	// uni resiliency
	UniResiliency Resiliency `json:"uniResiliency,omitempty"`
}

// Validate validates this m e f u n i spec v3
func (m *MEFUNISpecV3) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMaxServiceFrameSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhysicalLayer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSynchronousMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUniL2cpAddressSet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUniL2cpPeering(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUniResiliency(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MEFUNISpecV3) validateMaxServiceFrameSize(formats strfmt.Registry) error {

	if swag.IsZero(m.MaxServiceFrameSize) { // not required
		return nil
	}

	if err := validate.MinimumInt("maxServiceFrameSize", "body", int64(m.MaxServiceFrameSize), 1522, false); err != nil {
		return err
	}

	return nil
}

func (m *MEFUNISpecV3) validatePhysicalLayer(formats strfmt.Registry) error {

	if err := validate.Required("physicalLayer", "body", m.PhysicalLayer); err != nil {
		return err
	}

	for i := 0; i < len(m.PhysicalLayer); i++ {

		if err := m.PhysicalLayer[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("physicalLayer" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *MEFUNISpecV3) validateSynchronousMode(formats strfmt.Registry) error {

	if swag.IsZero(m.SynchronousMode) { // not required
		return nil
	}

	for i := 0; i < len(m.SynchronousMode); i++ {
		if swag.IsZero(m.SynchronousMode[i]) { // not required
			continue
		}

		if m.SynchronousMode[i] != nil {
			if err := m.SynchronousMode[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("synchronousMode" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MEFUNISpecV3) validateUniL2cpAddressSet(formats strfmt.Registry) error {

	if swag.IsZero(m.UniL2cpAddressSet) { // not required
		return nil
	}

	if err := m.UniL2cpAddressSet.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("uniL2cpAddressSet")
		}
		return err
	}

	return nil
}

func (m *MEFUNISpecV3) validateUniL2cpPeering(formats strfmt.Registry) error {

	if swag.IsZero(m.UniL2cpPeering) { // not required
		return nil
	}

	for i := 0; i < len(m.UniL2cpPeering); i++ {
		if swag.IsZero(m.UniL2cpPeering[i]) { // not required
			continue
		}

		if m.UniL2cpPeering[i] != nil {
			if err := m.UniL2cpPeering[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("uniL2cpPeering" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MEFUNISpecV3) validateUniResiliency(formats strfmt.Registry) error {

	if swag.IsZero(m.UniResiliency) { // not required
		return nil
	}

	if err := m.UniResiliency.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("uniResiliency")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MEFUNISpecV3) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MEFUNISpecV3) UnmarshalBinary(b []byte) error {
	var res MEFUNISpecV3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}