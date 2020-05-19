// Code generated by go-swagger; DO NOT EDIT.

package notification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/qlcchain/go-virtual-lsobus/sonata/notify/order/models"
)

// NotificationProductOrderCreationNotificationNoContentCode is the HTTP code returned for type NotificationProductOrderCreationNotificationNoContent
const NotificationProductOrderCreationNotificationNoContentCode int = 204

/*NotificationProductOrderCreationNotificationNoContent No Content

swagger:response notificationProductOrderCreationNotificationNoContent
*/
type NotificationProductOrderCreationNotificationNoContent struct {
}

// NewNotificationProductOrderCreationNotificationNoContent creates NotificationProductOrderCreationNotificationNoContent with default headers values
func NewNotificationProductOrderCreationNotificationNoContent() *NotificationProductOrderCreationNotificationNoContent {

	return &NotificationProductOrderCreationNotificationNoContent{}
}

// WriteResponse to the client
func (o *NotificationProductOrderCreationNotificationNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// NotificationProductOrderCreationNotificationBadRequestCode is the HTTP code returned for type NotificationProductOrderCreationNotificationBadRequest
const NotificationProductOrderCreationNotificationBadRequestCode int = 400

/*NotificationProductOrderCreationNotificationBadRequest Bad Request

List of supported error codes:
- 20: Invalid URL parameter value
- 21: Missing body
- 22: Invalid body
- 23: Missing body field
- 24: Invalid body field
- 25: Missing header
- 26: Invalid header value
- 27: Missing query-string parameter
- 28: Invalid query-string parameter value

swagger:response notificationProductOrderCreationNotificationBadRequest
*/
type NotificationProductOrderCreationNotificationBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewNotificationProductOrderCreationNotificationBadRequest creates NotificationProductOrderCreationNotificationBadRequest with default headers values
func NewNotificationProductOrderCreationNotificationBadRequest() *NotificationProductOrderCreationNotificationBadRequest {

	return &NotificationProductOrderCreationNotificationBadRequest{}
}

// WithPayload adds the payload to the notification product order creation notification bad request response
func (o *NotificationProductOrderCreationNotificationBadRequest) WithPayload(payload *models.ErrorRepresentation) *NotificationProductOrderCreationNotificationBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the notification product order creation notification bad request response
func (o *NotificationProductOrderCreationNotificationBadRequest) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NotificationProductOrderCreationNotificationBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NotificationProductOrderCreationNotificationUnauthorizedCode is the HTTP code returned for type NotificationProductOrderCreationNotificationUnauthorized
const NotificationProductOrderCreationNotificationUnauthorizedCode int = 401

/*NotificationProductOrderCreationNotificationUnauthorized Unauthorized

List of supported error codes:
- 40: Missing credentials
- 41: Invalid credentials
- 42: Expired credentials

swagger:response notificationProductOrderCreationNotificationUnauthorized
*/
type NotificationProductOrderCreationNotificationUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewNotificationProductOrderCreationNotificationUnauthorized creates NotificationProductOrderCreationNotificationUnauthorized with default headers values
func NewNotificationProductOrderCreationNotificationUnauthorized() *NotificationProductOrderCreationNotificationUnauthorized {

	return &NotificationProductOrderCreationNotificationUnauthorized{}
}

// WithPayload adds the payload to the notification product order creation notification unauthorized response
func (o *NotificationProductOrderCreationNotificationUnauthorized) WithPayload(payload *models.ErrorRepresentation) *NotificationProductOrderCreationNotificationUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the notification product order creation notification unauthorized response
func (o *NotificationProductOrderCreationNotificationUnauthorized) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NotificationProductOrderCreationNotificationUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NotificationProductOrderCreationNotificationForbiddenCode is the HTTP code returned for type NotificationProductOrderCreationNotificationForbidden
const NotificationProductOrderCreationNotificationForbiddenCode int = 403

/*NotificationProductOrderCreationNotificationForbidden Forbidden

List of supported error codes:
- 50: Access denied
- 51: Forbidden requester
- 52: Forbidden user
- 53: Too many requests

swagger:response notificationProductOrderCreationNotificationForbidden
*/
type NotificationProductOrderCreationNotificationForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewNotificationProductOrderCreationNotificationForbidden creates NotificationProductOrderCreationNotificationForbidden with default headers values
func NewNotificationProductOrderCreationNotificationForbidden() *NotificationProductOrderCreationNotificationForbidden {

	return &NotificationProductOrderCreationNotificationForbidden{}
}

// WithPayload adds the payload to the notification product order creation notification forbidden response
func (o *NotificationProductOrderCreationNotificationForbidden) WithPayload(payload *models.ErrorRepresentation) *NotificationProductOrderCreationNotificationForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the notification product order creation notification forbidden response
func (o *NotificationProductOrderCreationNotificationForbidden) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NotificationProductOrderCreationNotificationForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NotificationProductOrderCreationNotificationNotFoundCode is the HTTP code returned for type NotificationProductOrderCreationNotificationNotFound
const NotificationProductOrderCreationNotificationNotFoundCode int = 404

/*NotificationProductOrderCreationNotificationNotFound Not Found

List of supported error codes:
- 60: Resource not found

swagger:response notificationProductOrderCreationNotificationNotFound
*/
type NotificationProductOrderCreationNotificationNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewNotificationProductOrderCreationNotificationNotFound creates NotificationProductOrderCreationNotificationNotFound with default headers values
func NewNotificationProductOrderCreationNotificationNotFound() *NotificationProductOrderCreationNotificationNotFound {

	return &NotificationProductOrderCreationNotificationNotFound{}
}

// WithPayload adds the payload to the notification product order creation notification not found response
func (o *NotificationProductOrderCreationNotificationNotFound) WithPayload(payload *models.ErrorRepresentation) *NotificationProductOrderCreationNotificationNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the notification product order creation notification not found response
func (o *NotificationProductOrderCreationNotificationNotFound) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NotificationProductOrderCreationNotificationNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NotificationProductOrderCreationNotificationMethodNotAllowedCode is the HTTP code returned for type NotificationProductOrderCreationNotificationMethodNotAllowed
const NotificationProductOrderCreationNotificationMethodNotAllowedCode int = 405

/*NotificationProductOrderCreationNotificationMethodNotAllowed Method Not Allowed

List of supported error codes:
- 61: Method not allowed

swagger:response notificationProductOrderCreationNotificationMethodNotAllowed
*/
type NotificationProductOrderCreationNotificationMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewNotificationProductOrderCreationNotificationMethodNotAllowed creates NotificationProductOrderCreationNotificationMethodNotAllowed with default headers values
func NewNotificationProductOrderCreationNotificationMethodNotAllowed() *NotificationProductOrderCreationNotificationMethodNotAllowed {

	return &NotificationProductOrderCreationNotificationMethodNotAllowed{}
}

// WithPayload adds the payload to the notification product order creation notification method not allowed response
func (o *NotificationProductOrderCreationNotificationMethodNotAllowed) WithPayload(payload *models.ErrorRepresentation) *NotificationProductOrderCreationNotificationMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the notification product order creation notification method not allowed response
func (o *NotificationProductOrderCreationNotificationMethodNotAllowed) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NotificationProductOrderCreationNotificationMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NotificationProductOrderCreationNotificationRequestTimeoutCode is the HTTP code returned for type NotificationProductOrderCreationNotificationRequestTimeout
const NotificationProductOrderCreationNotificationRequestTimeoutCode int = 408

/*NotificationProductOrderCreationNotificationRequestTimeout Request Time-out

List of supported error codes:
- 63: Request time-out

swagger:response notificationProductOrderCreationNotificationRequestTimeout
*/
type NotificationProductOrderCreationNotificationRequestTimeout struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewNotificationProductOrderCreationNotificationRequestTimeout creates NotificationProductOrderCreationNotificationRequestTimeout with default headers values
func NewNotificationProductOrderCreationNotificationRequestTimeout() *NotificationProductOrderCreationNotificationRequestTimeout {

	return &NotificationProductOrderCreationNotificationRequestTimeout{}
}

// WithPayload adds the payload to the notification product order creation notification request timeout response
func (o *NotificationProductOrderCreationNotificationRequestTimeout) WithPayload(payload *models.ErrorRepresentation) *NotificationProductOrderCreationNotificationRequestTimeout {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the notification product order creation notification request timeout response
func (o *NotificationProductOrderCreationNotificationRequestTimeout) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NotificationProductOrderCreationNotificationRequestTimeout) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(408)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NotificationProductOrderCreationNotificationUnprocessableEntityCode is the HTTP code returned for type NotificationProductOrderCreationNotificationUnprocessableEntity
const NotificationProductOrderCreationNotificationUnprocessableEntityCode int = 422

/*NotificationProductOrderCreationNotificationUnprocessableEntity Unprocessable entity

Functional error

swagger:response notificationProductOrderCreationNotificationUnprocessableEntity
*/
type NotificationProductOrderCreationNotificationUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewNotificationProductOrderCreationNotificationUnprocessableEntity creates NotificationProductOrderCreationNotificationUnprocessableEntity with default headers values
func NewNotificationProductOrderCreationNotificationUnprocessableEntity() *NotificationProductOrderCreationNotificationUnprocessableEntity {

	return &NotificationProductOrderCreationNotificationUnprocessableEntity{}
}

// WithPayload adds the payload to the notification product order creation notification unprocessable entity response
func (o *NotificationProductOrderCreationNotificationUnprocessableEntity) WithPayload(payload *models.ErrorRepresentation) *NotificationProductOrderCreationNotificationUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the notification product order creation notification unprocessable entity response
func (o *NotificationProductOrderCreationNotificationUnprocessableEntity) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NotificationProductOrderCreationNotificationUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NotificationProductOrderCreationNotificationInternalServerErrorCode is the HTTP code returned for type NotificationProductOrderCreationNotificationInternalServerError
const NotificationProductOrderCreationNotificationInternalServerErrorCode int = 500

/*NotificationProductOrderCreationNotificationInternalServerError Internal Server Error

List of supported error codes:
- 1: Internal error

swagger:response notificationProductOrderCreationNotificationInternalServerError
*/
type NotificationProductOrderCreationNotificationInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewNotificationProductOrderCreationNotificationInternalServerError creates NotificationProductOrderCreationNotificationInternalServerError with default headers values
func NewNotificationProductOrderCreationNotificationInternalServerError() *NotificationProductOrderCreationNotificationInternalServerError {

	return &NotificationProductOrderCreationNotificationInternalServerError{}
}

// WithPayload adds the payload to the notification product order creation notification internal server error response
func (o *NotificationProductOrderCreationNotificationInternalServerError) WithPayload(payload *models.ErrorRepresentation) *NotificationProductOrderCreationNotificationInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the notification product order creation notification internal server error response
func (o *NotificationProductOrderCreationNotificationInternalServerError) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NotificationProductOrderCreationNotificationInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NotificationProductOrderCreationNotificationServiceUnavailableCode is the HTTP code returned for type NotificationProductOrderCreationNotificationServiceUnavailable
const NotificationProductOrderCreationNotificationServiceUnavailableCode int = 503

/*NotificationProductOrderCreationNotificationServiceUnavailable Service Unavailable



swagger:response notificationProductOrderCreationNotificationServiceUnavailable
*/
type NotificationProductOrderCreationNotificationServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewNotificationProductOrderCreationNotificationServiceUnavailable creates NotificationProductOrderCreationNotificationServiceUnavailable with default headers values
func NewNotificationProductOrderCreationNotificationServiceUnavailable() *NotificationProductOrderCreationNotificationServiceUnavailable {

	return &NotificationProductOrderCreationNotificationServiceUnavailable{}
}

// WithPayload adds the payload to the notification product order creation notification service unavailable response
func (o *NotificationProductOrderCreationNotificationServiceUnavailable) WithPayload(payload *models.ErrorRepresentation) *NotificationProductOrderCreationNotificationServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the notification product order creation notification service unavailable response
func (o *NotificationProductOrderCreationNotificationServiceUnavailable) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NotificationProductOrderCreationNotificationServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}