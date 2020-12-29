// Code generated by go-swagger; DO NOT EDIT.

package notification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/qlcchain/go-lsobus/sonata/notify/order/models"
)

// NewNotificationProductOrderInformationRequiredNotificationParams creates a new NotificationProductOrderInformationRequiredNotificationParams object
// no default values defined in spec.
func NewNotificationProductOrderInformationRequiredNotificationParams() NotificationProductOrderInformationRequiredNotificationParams {

	return NotificationProductOrderInformationRequiredNotificationParams{}
}

// NotificationProductOrderInformationRequiredNotificationParams contains all the bound params for the notification product order information required notification operation
// typically these are obtained from a http.Request
//
// swagger:parameters notificationProductOrderInformationRequiredNotification
type NotificationProductOrderInformationRequiredNotificationParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	ProductOrderInformationRequired *models.EventPlus
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewNotificationProductOrderInformationRequiredNotificationParams() beforehand.
func (o *NotificationProductOrderInformationRequiredNotificationParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.EventPlus
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("productOrderInformationRequired", "body", ""))
			} else {
				res = append(res, errors.NewParseError("productOrderInformationRequired", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.ProductOrderInformationRequired = &body
			}
		}
	} else {
		res = append(res, errors.Required("productOrderInformationRequired", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
