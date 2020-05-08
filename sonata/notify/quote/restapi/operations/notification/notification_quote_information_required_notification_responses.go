// Code generated by go-swagger; DO NOT EDIT.

package notification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/iixlabs/virtual-lsobus/sonata/notify/quote/models"
)

// NotificationQuoteInformationRequiredNotificationNoContentCode is the HTTP code returned for type NotificationQuoteInformationRequiredNotificationNoContent
const NotificationQuoteInformationRequiredNotificationNoContentCode int = 204

/*NotificationQuoteInformationRequiredNotificationNoContent Success

swagger:response notificationQuoteInformationRequiredNotificationNoContent
*/
type NotificationQuoteInformationRequiredNotificationNoContent struct {
}

// NewNotificationQuoteInformationRequiredNotificationNoContent creates NotificationQuoteInformationRequiredNotificationNoContent with default headers values
func NewNotificationQuoteInformationRequiredNotificationNoContent() *NotificationQuoteInformationRequiredNotificationNoContent {

	return &NotificationQuoteInformationRequiredNotificationNoContent{}
}

// WriteResponse to the client
func (o *NotificationQuoteInformationRequiredNotificationNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// NotificationQuoteInformationRequiredNotificationBadRequestCode is the HTTP code returned for type NotificationQuoteInformationRequiredNotificationBadRequest
const NotificationQuoteInformationRequiredNotificationBadRequestCode int = 400

/*NotificationQuoteInformationRequiredNotificationBadRequest Bad Request

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

swagger:response notificationQuoteInformationRequiredNotificationBadRequest
*/
type NotificationQuoteInformationRequiredNotificationBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewNotificationQuoteInformationRequiredNotificationBadRequest creates NotificationQuoteInformationRequiredNotificationBadRequest with default headers values
func NewNotificationQuoteInformationRequiredNotificationBadRequest() *NotificationQuoteInformationRequiredNotificationBadRequest {

	return &NotificationQuoteInformationRequiredNotificationBadRequest{}
}

// WithPayload adds the payload to the notification quote information required notification bad request response
func (o *NotificationQuoteInformationRequiredNotificationBadRequest) WithPayload(payload *models.ErrorRepresentation) *NotificationQuoteInformationRequiredNotificationBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the notification quote information required notification bad request response
func (o *NotificationQuoteInformationRequiredNotificationBadRequest) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NotificationQuoteInformationRequiredNotificationBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NotificationQuoteInformationRequiredNotificationUnauthorizedCode is the HTTP code returned for type NotificationQuoteInformationRequiredNotificationUnauthorized
const NotificationQuoteInformationRequiredNotificationUnauthorizedCode int = 401

/*NotificationQuoteInformationRequiredNotificationUnauthorized Unauthorized

List of supported error codes:
- 40: Missing credentials
- 41: Invalid credentials
- 42: Expired credentials

swagger:response notificationQuoteInformationRequiredNotificationUnauthorized
*/
type NotificationQuoteInformationRequiredNotificationUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewNotificationQuoteInformationRequiredNotificationUnauthorized creates NotificationQuoteInformationRequiredNotificationUnauthorized with default headers values
func NewNotificationQuoteInformationRequiredNotificationUnauthorized() *NotificationQuoteInformationRequiredNotificationUnauthorized {

	return &NotificationQuoteInformationRequiredNotificationUnauthorized{}
}

// WithPayload adds the payload to the notification quote information required notification unauthorized response
func (o *NotificationQuoteInformationRequiredNotificationUnauthorized) WithPayload(payload *models.ErrorRepresentation) *NotificationQuoteInformationRequiredNotificationUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the notification quote information required notification unauthorized response
func (o *NotificationQuoteInformationRequiredNotificationUnauthorized) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NotificationQuoteInformationRequiredNotificationUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NotificationQuoteInformationRequiredNotificationForbiddenCode is the HTTP code returned for type NotificationQuoteInformationRequiredNotificationForbidden
const NotificationQuoteInformationRequiredNotificationForbiddenCode int = 403

/*NotificationQuoteInformationRequiredNotificationForbidden Forbidden

List of supported error codes:
- 50: Access denied
- 51: Forbidden requester
- 52: Forbidden user
- 53: Too many requests

swagger:response notificationQuoteInformationRequiredNotificationForbidden
*/
type NotificationQuoteInformationRequiredNotificationForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewNotificationQuoteInformationRequiredNotificationForbidden creates NotificationQuoteInformationRequiredNotificationForbidden with default headers values
func NewNotificationQuoteInformationRequiredNotificationForbidden() *NotificationQuoteInformationRequiredNotificationForbidden {

	return &NotificationQuoteInformationRequiredNotificationForbidden{}
}

// WithPayload adds the payload to the notification quote information required notification forbidden response
func (o *NotificationQuoteInformationRequiredNotificationForbidden) WithPayload(payload *models.ErrorRepresentation) *NotificationQuoteInformationRequiredNotificationForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the notification quote information required notification forbidden response
func (o *NotificationQuoteInformationRequiredNotificationForbidden) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NotificationQuoteInformationRequiredNotificationForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NotificationQuoteInformationRequiredNotificationNotFoundCode is the HTTP code returned for type NotificationQuoteInformationRequiredNotificationNotFound
const NotificationQuoteInformationRequiredNotificationNotFoundCode int = 404

/*NotificationQuoteInformationRequiredNotificationNotFound Not Found

List of supported error codes:
- 60: Resource not found

swagger:response notificationQuoteInformationRequiredNotificationNotFound
*/
type NotificationQuoteInformationRequiredNotificationNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewNotificationQuoteInformationRequiredNotificationNotFound creates NotificationQuoteInformationRequiredNotificationNotFound with default headers values
func NewNotificationQuoteInformationRequiredNotificationNotFound() *NotificationQuoteInformationRequiredNotificationNotFound {

	return &NotificationQuoteInformationRequiredNotificationNotFound{}
}

// WithPayload adds the payload to the notification quote information required notification not found response
func (o *NotificationQuoteInformationRequiredNotificationNotFound) WithPayload(payload *models.ErrorRepresentation) *NotificationQuoteInformationRequiredNotificationNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the notification quote information required notification not found response
func (o *NotificationQuoteInformationRequiredNotificationNotFound) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NotificationQuoteInformationRequiredNotificationNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NotificationQuoteInformationRequiredNotificationMethodNotAllowedCode is the HTTP code returned for type NotificationQuoteInformationRequiredNotificationMethodNotAllowed
const NotificationQuoteInformationRequiredNotificationMethodNotAllowedCode int = 405

/*NotificationQuoteInformationRequiredNotificationMethodNotAllowed Method Not Allowed

List of supported error codes:
- 61: Method not allowed

swagger:response notificationQuoteInformationRequiredNotificationMethodNotAllowed
*/
type NotificationQuoteInformationRequiredNotificationMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewNotificationQuoteInformationRequiredNotificationMethodNotAllowed creates NotificationQuoteInformationRequiredNotificationMethodNotAllowed with default headers values
func NewNotificationQuoteInformationRequiredNotificationMethodNotAllowed() *NotificationQuoteInformationRequiredNotificationMethodNotAllowed {

	return &NotificationQuoteInformationRequiredNotificationMethodNotAllowed{}
}

// WithPayload adds the payload to the notification quote information required notification method not allowed response
func (o *NotificationQuoteInformationRequiredNotificationMethodNotAllowed) WithPayload(payload *models.ErrorRepresentation) *NotificationQuoteInformationRequiredNotificationMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the notification quote information required notification method not allowed response
func (o *NotificationQuoteInformationRequiredNotificationMethodNotAllowed) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NotificationQuoteInformationRequiredNotificationMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NotificationQuoteInformationRequiredNotificationUnprocessableEntityCode is the HTTP code returned for type NotificationQuoteInformationRequiredNotificationUnprocessableEntity
const NotificationQuoteInformationRequiredNotificationUnprocessableEntityCode int = 422

/*NotificationQuoteInformationRequiredNotificationUnprocessableEntity Unprocessable entity

Functional error

swagger:response notificationQuoteInformationRequiredNotificationUnprocessableEntity
*/
type NotificationQuoteInformationRequiredNotificationUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewNotificationQuoteInformationRequiredNotificationUnprocessableEntity creates NotificationQuoteInformationRequiredNotificationUnprocessableEntity with default headers values
func NewNotificationQuoteInformationRequiredNotificationUnprocessableEntity() *NotificationQuoteInformationRequiredNotificationUnprocessableEntity {

	return &NotificationQuoteInformationRequiredNotificationUnprocessableEntity{}
}

// WithPayload adds the payload to the notification quote information required notification unprocessable entity response
func (o *NotificationQuoteInformationRequiredNotificationUnprocessableEntity) WithPayload(payload *models.ErrorRepresentation) *NotificationQuoteInformationRequiredNotificationUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the notification quote information required notification unprocessable entity response
func (o *NotificationQuoteInformationRequiredNotificationUnprocessableEntity) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NotificationQuoteInformationRequiredNotificationUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NotificationQuoteInformationRequiredNotificationInternalServerErrorCode is the HTTP code returned for type NotificationQuoteInformationRequiredNotificationInternalServerError
const NotificationQuoteInformationRequiredNotificationInternalServerErrorCode int = 500

/*NotificationQuoteInformationRequiredNotificationInternalServerError Internal Server Error

List of supported error codes:
- 1: Internal error

swagger:response notificationQuoteInformationRequiredNotificationInternalServerError
*/
type NotificationQuoteInformationRequiredNotificationInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewNotificationQuoteInformationRequiredNotificationInternalServerError creates NotificationQuoteInformationRequiredNotificationInternalServerError with default headers values
func NewNotificationQuoteInformationRequiredNotificationInternalServerError() *NotificationQuoteInformationRequiredNotificationInternalServerError {

	return &NotificationQuoteInformationRequiredNotificationInternalServerError{}
}

// WithPayload adds the payload to the notification quote information required notification internal server error response
func (o *NotificationQuoteInformationRequiredNotificationInternalServerError) WithPayload(payload *models.ErrorRepresentation) *NotificationQuoteInformationRequiredNotificationInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the notification quote information required notification internal server error response
func (o *NotificationQuoteInformationRequiredNotificationInternalServerError) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NotificationQuoteInformationRequiredNotificationInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NotificationQuoteInformationRequiredNotificationServiceUnavailableCode is the HTTP code returned for type NotificationQuoteInformationRequiredNotificationServiceUnavailable
const NotificationQuoteInformationRequiredNotificationServiceUnavailableCode int = 503

/*NotificationQuoteInformationRequiredNotificationServiceUnavailable Service Unavailable



swagger:response notificationQuoteInformationRequiredNotificationServiceUnavailable
*/
type NotificationQuoteInformationRequiredNotificationServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewNotificationQuoteInformationRequiredNotificationServiceUnavailable creates NotificationQuoteInformationRequiredNotificationServiceUnavailable with default headers values
func NewNotificationQuoteInformationRequiredNotificationServiceUnavailable() *NotificationQuoteInformationRequiredNotificationServiceUnavailable {

	return &NotificationQuoteInformationRequiredNotificationServiceUnavailable{}
}

// WithPayload adds the payload to the notification quote information required notification service unavailable response
func (o *NotificationQuoteInformationRequiredNotificationServiceUnavailable) WithPayload(payload *models.ErrorRepresentation) *NotificationQuoteInformationRequiredNotificationServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the notification quote information required notification service unavailable response
func (o *NotificationQuoteInformationRequiredNotificationServiceUnavailable) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NotificationQuoteInformationRequiredNotificationServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
