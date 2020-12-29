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

	"github.com/qlcchain/go-lsobus/sonata/notify/quote/models"
)

// NewNotificationQuoteInformationRequiredNotificationParams creates a new NotificationQuoteInformationRequiredNotificationParams object
// no default values defined in spec.
func NewNotificationQuoteInformationRequiredNotificationParams() NotificationQuoteInformationRequiredNotificationParams {

	return NotificationQuoteInformationRequiredNotificationParams{}
}

// NotificationQuoteInformationRequiredNotificationParams contains all the bound params for the notification quote information required notification operation
// typically these are obtained from a http.Request
//
// swagger:parameters notificationQuoteInformationRequiredNotification
type NotificationQuoteInformationRequiredNotificationParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	QuoteInformationRequiredNotification []*models.EventPlus
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewNotificationQuoteInformationRequiredNotificationParams() beforehand.
func (o *NotificationQuoteInformationRequiredNotificationParams) BindRequest(
	r *http.Request, route *middleware.MatchedRoute,
) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body []*models.EventPlus
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("quoteInformationRequiredNotification", "body", nil))
			} else {
				res = append(res, errors.NewParseError("quoteInformationRequiredNotification", "body", "", err))
			}
		} else {
			// validate array of body objects
			for i := range body {
				if body[i] == nil {
					continue
				}
				if err := body[i].Validate(route.Formats); err != nil {
					res = append(res, err)
					break
				}
			}
			if len(res) == 0 {
				o.QuoteInformationRequiredNotification = body
			}
		}
	} else {
		res = append(res, errors.Required("quoteInformationRequiredNotification", "body", nil))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
