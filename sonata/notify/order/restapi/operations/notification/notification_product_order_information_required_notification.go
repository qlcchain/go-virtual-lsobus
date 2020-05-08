// Code generated by go-swagger; DO NOT EDIT.

package notification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// NotificationProductOrderInformationRequiredNotificationHandlerFunc turns a function with the right signature into a notification product order information required notification handler
type NotificationProductOrderInformationRequiredNotificationHandlerFunc func(NotificationProductOrderInformationRequiredNotificationParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NotificationProductOrderInformationRequiredNotificationHandlerFunc) Handle(params NotificationProductOrderInformationRequiredNotificationParams) middleware.Responder {
	return fn(params)
}

// NotificationProductOrderInformationRequiredNotificationHandler interface for that can handle valid notification product order information required notification params
type NotificationProductOrderInformationRequiredNotificationHandler interface {
	Handle(NotificationProductOrderInformationRequiredNotificationParams) middleware.Responder
}

// NewNotificationProductOrderInformationRequiredNotification creates a new http.Handler for the notification product order information required notification operation
func NewNotificationProductOrderInformationRequiredNotification(ctx *middleware.Context, handler NotificationProductOrderInformationRequiredNotificationHandler) *NotificationProductOrderInformationRequiredNotification {
	return &NotificationProductOrderInformationRequiredNotification{Context: ctx, Handler: handler}
}

/*NotificationProductOrderInformationRequiredNotification swagger:route POST /notification/productOrderInformationRequiredNotification Notification notificationProductOrderInformationRequiredNotification

Product Order information required structure

Product Order information required structure description

*/
type NotificationProductOrderInformationRequiredNotification struct {
	Context *middleware.Context
	Handler NotificationProductOrderInformationRequiredNotificationHandler
}

func (o *NotificationProductOrderInformationRequiredNotification) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNotificationProductOrderInformationRequiredNotificationParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
