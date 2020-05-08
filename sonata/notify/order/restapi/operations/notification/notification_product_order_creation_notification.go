// Code generated by go-swagger; DO NOT EDIT.

package notification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// NotificationProductOrderCreationNotificationHandlerFunc turns a function with the right signature into a notification product order creation notification handler
type NotificationProductOrderCreationNotificationHandlerFunc func(NotificationProductOrderCreationNotificationParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NotificationProductOrderCreationNotificationHandlerFunc) Handle(params NotificationProductOrderCreationNotificationParams) middleware.Responder {
	return fn(params)
}

// NotificationProductOrderCreationNotificationHandler interface for that can handle valid notification product order creation notification params
type NotificationProductOrderCreationNotificationHandler interface {
	Handle(NotificationProductOrderCreationNotificationParams) middleware.Responder
}

// NewNotificationProductOrderCreationNotification creates a new http.Handler for the notification product order creation notification operation
func NewNotificationProductOrderCreationNotification(ctx *middleware.Context, handler NotificationProductOrderCreationNotificationHandler) *NotificationProductOrderCreationNotification {
	return &NotificationProductOrderCreationNotification{Context: ctx, Handler: handler}
}

/*NotificationProductOrderCreationNotification swagger:route POST /notification/productOrderCreationNotification Notification notificationProductOrderCreationNotification

Product order creation notification structure

Product order creation notification structure description

*/
type NotificationProductOrderCreationNotification struct {
	Context *middleware.Context
	Handler NotificationProductOrderCreationNotificationHandler
}

func (o *NotificationProductOrderCreationNotification) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNotificationProductOrderCreationNotificationParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
