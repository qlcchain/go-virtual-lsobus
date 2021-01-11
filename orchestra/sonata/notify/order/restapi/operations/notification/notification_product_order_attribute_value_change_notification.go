// Code generated by go-swagger; DO NOT EDIT.

package notification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// NotificationProductOrderAttributeValueChangeNotificationHandlerFunc turns a function with the right signature into a notification product order attribute value change notification handler
type NotificationProductOrderAttributeValueChangeNotificationHandlerFunc func(NotificationProductOrderAttributeValueChangeNotificationParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NotificationProductOrderAttributeValueChangeNotificationHandlerFunc) Handle(params NotificationProductOrderAttributeValueChangeNotificationParams) middleware.Responder {
	return fn(params)
}

// NotificationProductOrderAttributeValueChangeNotificationHandler interface for that can handle valid notification product order attribute value change notification params
type NotificationProductOrderAttributeValueChangeNotificationHandler interface {
	Handle(NotificationProductOrderAttributeValueChangeNotificationParams) middleware.Responder
}

// NewNotificationProductOrderAttributeValueChangeNotification creates a new http.Handler for the notification product order attribute value change notification operation
func NewNotificationProductOrderAttributeValueChangeNotification(ctx *middleware.Context, handler NotificationProductOrderAttributeValueChangeNotificationHandler) *NotificationProductOrderAttributeValueChangeNotification {
	return &NotificationProductOrderAttributeValueChangeNotification{Context: ctx, Handler: handler}
}

/*NotificationProductOrderAttributeValueChangeNotification swagger:route POST /notification/productOrderAttributeValueChangeNotification Notification notificationProductOrderAttributeValueChangeNotification

Product Order attribute value change structure

Product Order attribute value change structure description

*/
type NotificationProductOrderAttributeValueChangeNotification struct {
	Context *middleware.Context
	Handler NotificationProductOrderAttributeValueChangeNotificationHandler
}

func (o *NotificationProductOrderAttributeValueChangeNotification) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNotificationProductOrderAttributeValueChangeNotificationParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}