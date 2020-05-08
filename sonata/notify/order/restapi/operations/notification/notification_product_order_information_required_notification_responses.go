// Code generated by go-swagger; DO NOT EDIT.

package notification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/iixlabs/virtual-lsobus/sonata/notify/order/models"
)

// NotificationProductOrderInformationRequiredNotificationNoContentCode is the HTTP code returned for type NotificationProductOrderInformationRequiredNotificationNoContent
const NotificationProductOrderInformationRequiredNotificationNoContentCode int = 204

/*NotificationProductOrderInformationRequiredNotificationNoContent No Content

swagger:response notificationProductOrderInformationRequiredNotificationNoContent
*/
type NotificationProductOrderInformationRequiredNotificationNoContent struct {
}

// NewNotificationProductOrderInformationRequiredNotificationNoContent creates NotificationProductOrderInformationRequiredNotificationNoContent with default headers values
func NewNotificationProductOrderInformationRequiredNotificationNoContent() *NotificationProductOrderInformationRequiredNotificationNoContent {

	return &NotificationProductOrderInformationRequiredNotificationNoContent{}
}

// WriteResponse to the client
func (o *NotificationProductOrderInformationRequiredNotificationNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// NotificationProductOrderInformationRequiredNotificationBadRequestCode is the HTTP code returned for type NotificationProductOrderInformationRequiredNotificationBadRequest
const NotificationProductOrderInformationRequiredNotificationBadRequestCode int = 400

/*NotificationProductOrderInformationRequiredNotificationBadRequest Bad Request

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

swagger:response notificationProductOrderInformationRequiredNotificationBadRequest
*/
type NotificationProductOrderInformationRequiredNotificationBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewNotificationProductOrderInformationRequiredNotificationBadRequest creates NotificationProductOrderInformationRequiredNotificationBadRequest with default headers values
func NewNotificationProductOrderInformationRequiredNotificationBadRequest() *NotificationProductOrderInformationRequiredNotificationBadRequest {

	return &NotificationProductOrderInformationRequiredNotificationBadRequest{}
}

// WithPayload adds the payload to the notification product order information required notification bad request response
func (o *NotificationProductOrderInformationRequiredNotificationBadRequest) WithPayload(payload *models.ErrorRepresentation) *NotificationProductOrderInformationRequiredNotificationBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the notification product order information required notification bad request response
func (o *NotificationProductOrderInformationRequiredNotificationBadRequest) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NotificationProductOrderInformationRequiredNotificationBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NotificationProductOrderInformationRequiredNotificationUnauthorizedCode is the HTTP code returned for type NotificationProductOrderInformationRequiredNotificationUnauthorized
const NotificationProductOrderInformationRequiredNotificationUnauthorizedCode int = 401

/*NotificationProductOrderInformationRequiredNotificationUnauthorized Unauthorized

List of supported error codes:
- 40: Missing credentials
- 41: Invalid credentials
- 42: Expired credentials

swagger:response notificationProductOrderInformationRequiredNotificationUnauthorized
*/
type NotificationProductOrderInformationRequiredNotificationUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewNotificationProductOrderInformationRequiredNotificationUnauthorized creates NotificationProductOrderInformationRequiredNotificationUnauthorized with default headers values
func NewNotificationProductOrderInformationRequiredNotificationUnauthorized() *NotificationProductOrderInformationRequiredNotificationUnauthorized {

	return &NotificationProductOrderInformationRequiredNotificationUnauthorized{}
}

// WithPayload adds the payload to the notification product order information required notification unauthorized response
func (o *NotificationProductOrderInformationRequiredNotificationUnauthorized) WithPayload(payload *models.ErrorRepresentation) *NotificationProductOrderInformationRequiredNotificationUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the notification product order information required notification unauthorized response
func (o *NotificationProductOrderInformationRequiredNotificationUnauthorized) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NotificationProductOrderInformationRequiredNotificationUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NotificationProductOrderInformationRequiredNotificationForbiddenCode is the HTTP code returned for type NotificationProductOrderInformationRequiredNotificationForbidden
const NotificationProductOrderInformationRequiredNotificationForbiddenCode int = 403

/*NotificationProductOrderInformationRequiredNotificationForbidden Forbidden

List of supported error codes:
- 50: Access denied
- 51: Forbidden requester
- 52: Forbidden user
- 53: Too many requests

swagger:response notificationProductOrderInformationRequiredNotificationForbidden
*/
type NotificationProductOrderInformationRequiredNotificationForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewNotificationProductOrderInformationRequiredNotificationForbidden creates NotificationProductOrderInformationRequiredNotificationForbidden with default headers values
func NewNotificationProductOrderInformationRequiredNotificationForbidden() *NotificationProductOrderInformationRequiredNotificationForbidden {

	return &NotificationProductOrderInformationRequiredNotificationForbidden{}
}

// WithPayload adds the payload to the notification product order information required notification forbidden response
func (o *NotificationProductOrderInformationRequiredNotificationForbidden) WithPayload(payload *models.ErrorRepresentation) *NotificationProductOrderInformationRequiredNotificationForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the notification product order information required notification forbidden response
func (o *NotificationProductOrderInformationRequiredNotificationForbidden) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NotificationProductOrderInformationRequiredNotificationForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NotificationProductOrderInformationRequiredNotificationNotFoundCode is the HTTP code returned for type NotificationProductOrderInformationRequiredNotificationNotFound
const NotificationProductOrderInformationRequiredNotificationNotFoundCode int = 404

/*NotificationProductOrderInformationRequiredNotificationNotFound Not Found

List of supported error codes:
- 60: Resource not found

swagger:response notificationProductOrderInformationRequiredNotificationNotFound
*/
type NotificationProductOrderInformationRequiredNotificationNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewNotificationProductOrderInformationRequiredNotificationNotFound creates NotificationProductOrderInformationRequiredNotificationNotFound with default headers values
func NewNotificationProductOrderInformationRequiredNotificationNotFound() *NotificationProductOrderInformationRequiredNotificationNotFound {

	return &NotificationProductOrderInformationRequiredNotificationNotFound{}
}

// WithPayload adds the payload to the notification product order information required notification not found response
func (o *NotificationProductOrderInformationRequiredNotificationNotFound) WithPayload(payload *models.ErrorRepresentation) *NotificationProductOrderInformationRequiredNotificationNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the notification product order information required notification not found response
func (o *NotificationProductOrderInformationRequiredNotificationNotFound) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NotificationProductOrderInformationRequiredNotificationNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NotificationProductOrderInformationRequiredNotificationMethodNotAllowedCode is the HTTP code returned for type NotificationProductOrderInformationRequiredNotificationMethodNotAllowed
const NotificationProductOrderInformationRequiredNotificationMethodNotAllowedCode int = 405

/*NotificationProductOrderInformationRequiredNotificationMethodNotAllowed Method Not Allowed

List of supported error codes:
- 61: Method not allowed

swagger:response notificationProductOrderInformationRequiredNotificationMethodNotAllowed
*/
type NotificationProductOrderInformationRequiredNotificationMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewNotificationProductOrderInformationRequiredNotificationMethodNotAllowed creates NotificationProductOrderInformationRequiredNotificationMethodNotAllowed with default headers values
func NewNotificationProductOrderInformationRequiredNotificationMethodNotAllowed() *NotificationProductOrderInformationRequiredNotificationMethodNotAllowed {

	return &NotificationProductOrderInformationRequiredNotificationMethodNotAllowed{}
}

// WithPayload adds the payload to the notification product order information required notification method not allowed response
func (o *NotificationProductOrderInformationRequiredNotificationMethodNotAllowed) WithPayload(payload *models.ErrorRepresentation) *NotificationProductOrderInformationRequiredNotificationMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the notification product order information required notification method not allowed response
func (o *NotificationProductOrderInformationRequiredNotificationMethodNotAllowed) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NotificationProductOrderInformationRequiredNotificationMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NotificationProductOrderInformationRequiredNotificationRequestTimeoutCode is the HTTP code returned for type NotificationProductOrderInformationRequiredNotificationRequestTimeout
const NotificationProductOrderInformationRequiredNotificationRequestTimeoutCode int = 408

/*NotificationProductOrderInformationRequiredNotificationRequestTimeout Request Time-out

List of supported error codes:
- 63: Request time-out

swagger:response notificationProductOrderInformationRequiredNotificationRequestTimeout
*/
type NotificationProductOrderInformationRequiredNotificationRequestTimeout struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewNotificationProductOrderInformationRequiredNotificationRequestTimeout creates NotificationProductOrderInformationRequiredNotificationRequestTimeout with default headers values
func NewNotificationProductOrderInformationRequiredNotificationRequestTimeout() *NotificationProductOrderInformationRequiredNotificationRequestTimeout {

	return &NotificationProductOrderInformationRequiredNotificationRequestTimeout{}
}

// WithPayload adds the payload to the notification product order information required notification request timeout response
func (o *NotificationProductOrderInformationRequiredNotificationRequestTimeout) WithPayload(payload *models.ErrorRepresentation) *NotificationProductOrderInformationRequiredNotificationRequestTimeout {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the notification product order information required notification request timeout response
func (o *NotificationProductOrderInformationRequiredNotificationRequestTimeout) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NotificationProductOrderInformationRequiredNotificationRequestTimeout) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(408)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NotificationProductOrderInformationRequiredNotificationUnprocessableEntityCode is the HTTP code returned for type NotificationProductOrderInformationRequiredNotificationUnprocessableEntity
const NotificationProductOrderInformationRequiredNotificationUnprocessableEntityCode int = 422

/*NotificationProductOrderInformationRequiredNotificationUnprocessableEntity Unprocessable entity

Functional error

swagger:response notificationProductOrderInformationRequiredNotificationUnprocessableEntity
*/
type NotificationProductOrderInformationRequiredNotificationUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewNotificationProductOrderInformationRequiredNotificationUnprocessableEntity creates NotificationProductOrderInformationRequiredNotificationUnprocessableEntity with default headers values
func NewNotificationProductOrderInformationRequiredNotificationUnprocessableEntity() *NotificationProductOrderInformationRequiredNotificationUnprocessableEntity {

	return &NotificationProductOrderInformationRequiredNotificationUnprocessableEntity{}
}

// WithPayload adds the payload to the notification product order information required notification unprocessable entity response
func (o *NotificationProductOrderInformationRequiredNotificationUnprocessableEntity) WithPayload(payload *models.ErrorRepresentation) *NotificationProductOrderInformationRequiredNotificationUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the notification product order information required notification unprocessable entity response
func (o *NotificationProductOrderInformationRequiredNotificationUnprocessableEntity) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NotificationProductOrderInformationRequiredNotificationUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NotificationProductOrderInformationRequiredNotificationInternalServerErrorCode is the HTTP code returned for type NotificationProductOrderInformationRequiredNotificationInternalServerError
const NotificationProductOrderInformationRequiredNotificationInternalServerErrorCode int = 500

/*NotificationProductOrderInformationRequiredNotificationInternalServerError Internal Server Error

List of supported error codes:
- 1: Internal error

swagger:response notificationProductOrderInformationRequiredNotificationInternalServerError
*/
type NotificationProductOrderInformationRequiredNotificationInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewNotificationProductOrderInformationRequiredNotificationInternalServerError creates NotificationProductOrderInformationRequiredNotificationInternalServerError with default headers values
func NewNotificationProductOrderInformationRequiredNotificationInternalServerError() *NotificationProductOrderInformationRequiredNotificationInternalServerError {

	return &NotificationProductOrderInformationRequiredNotificationInternalServerError{}
}

// WithPayload adds the payload to the notification product order information required notification internal server error response
func (o *NotificationProductOrderInformationRequiredNotificationInternalServerError) WithPayload(payload *models.ErrorRepresentation) *NotificationProductOrderInformationRequiredNotificationInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the notification product order information required notification internal server error response
func (o *NotificationProductOrderInformationRequiredNotificationInternalServerError) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NotificationProductOrderInformationRequiredNotificationInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NotificationProductOrderInformationRequiredNotificationServiceUnavailableCode is the HTTP code returned for type NotificationProductOrderInformationRequiredNotificationServiceUnavailable
const NotificationProductOrderInformationRequiredNotificationServiceUnavailableCode int = 503

/*NotificationProductOrderInformationRequiredNotificationServiceUnavailable Service Unavailable



swagger:response notificationProductOrderInformationRequiredNotificationServiceUnavailable
*/
type NotificationProductOrderInformationRequiredNotificationServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewNotificationProductOrderInformationRequiredNotificationServiceUnavailable creates NotificationProductOrderInformationRequiredNotificationServiceUnavailable with default headers values
func NewNotificationProductOrderInformationRequiredNotificationServiceUnavailable() *NotificationProductOrderInformationRequiredNotificationServiceUnavailable {

	return &NotificationProductOrderInformationRequiredNotificationServiceUnavailable{}
}

// WithPayload adds the payload to the notification product order information required notification service unavailable response
func (o *NotificationProductOrderInformationRequiredNotificationServiceUnavailable) WithPayload(payload *models.ErrorRepresentation) *NotificationProductOrderInformationRequiredNotificationServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the notification product order information required notification service unavailable response
func (o *NotificationProductOrderInformationRequiredNotificationServiceUnavailable) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NotificationProductOrderInformationRequiredNotificationServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
