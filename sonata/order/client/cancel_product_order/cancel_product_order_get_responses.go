// Code generated by go-swagger; DO NOT EDIT.

package cancel_product_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/iixlabs/virtual-lsobus/sonata/order/models"
)

// CancelProductOrderGetReader is a Reader for the CancelProductOrderGet structure.
type CancelProductOrderGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelProductOrderGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCancelProductOrderGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCancelProductOrderGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCancelProductOrderGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCancelProductOrderGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCancelProductOrderGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewCancelProductOrderGetMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewCancelProductOrderGetRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCancelProductOrderGetUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCancelProductOrderGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewCancelProductOrderGetServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCancelProductOrderGetOK creates a CancelProductOrderGetOK with default headers values
func NewCancelProductOrderGetOK() *CancelProductOrderGetOK {
	return &CancelProductOrderGetOK{}
}

/*CancelProductOrderGetOK handles this case with default header values.

Ok
*/
type CancelProductOrderGetOK struct {
	Payload *models.CancelProductOrder
}

func (o *CancelProductOrderGetOK) Error() string {
	return fmt.Sprintf("[GET /cancelProductOrder/{CancelProductOrderId}][%d] cancelProductOrderGetOK  %+v", 200, o.Payload)
}

func (o *CancelProductOrderGetOK) GetPayload() *models.CancelProductOrder {
	return o.Payload
}

func (o *CancelProductOrderGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CancelProductOrder)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelProductOrderGetBadRequest creates a CancelProductOrderGetBadRequest with default headers values
func NewCancelProductOrderGetBadRequest() *CancelProductOrderGetBadRequest {
	return &CancelProductOrderGetBadRequest{}
}

/*CancelProductOrderGetBadRequest handles this case with default header values.

Bad Request

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
*/
type CancelProductOrderGetBadRequest struct {
	Payload *models.ErrorRepresentation
}

func (o *CancelProductOrderGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /cancelProductOrder/{CancelProductOrderId}][%d] cancelProductOrderGetBadRequest  %+v", 400, o.Payload)
}

func (o *CancelProductOrderGetBadRequest) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *CancelProductOrderGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelProductOrderGetUnauthorized creates a CancelProductOrderGetUnauthorized with default headers values
func NewCancelProductOrderGetUnauthorized() *CancelProductOrderGetUnauthorized {
	return &CancelProductOrderGetUnauthorized{}
}

/*CancelProductOrderGetUnauthorized handles this case with default header values.

Unauthorized

List of supported error codes:
- 40: Missing credentials
- 41: Invalid credentials
- 42: Expired credentials
*/
type CancelProductOrderGetUnauthorized struct {
	Payload *models.ErrorRepresentation
}

func (o *CancelProductOrderGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cancelProductOrder/{CancelProductOrderId}][%d] cancelProductOrderGetUnauthorized  %+v", 401, o.Payload)
}

func (o *CancelProductOrderGetUnauthorized) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *CancelProductOrderGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelProductOrderGetForbidden creates a CancelProductOrderGetForbidden with default headers values
func NewCancelProductOrderGetForbidden() *CancelProductOrderGetForbidden {
	return &CancelProductOrderGetForbidden{}
}

/*CancelProductOrderGetForbidden handles this case with default header values.

Forbidden

List of supported error codes:
- 50: Access denied
- 51: Forbidden requester
- 52: Forbidden user
- 53: Too many requests
*/
type CancelProductOrderGetForbidden struct {
	Payload *models.ErrorRepresentation
}

func (o *CancelProductOrderGetForbidden) Error() string {
	return fmt.Sprintf("[GET /cancelProductOrder/{CancelProductOrderId}][%d] cancelProductOrderGetForbidden  %+v", 403, o.Payload)
}

func (o *CancelProductOrderGetForbidden) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *CancelProductOrderGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelProductOrderGetNotFound creates a CancelProductOrderGetNotFound with default headers values
func NewCancelProductOrderGetNotFound() *CancelProductOrderGetNotFound {
	return &CancelProductOrderGetNotFound{}
}

/*CancelProductOrderGetNotFound handles this case with default header values.

Not Found

List of supported error codes:
- 60: Resource not found
*/
type CancelProductOrderGetNotFound struct {
	Payload *models.ErrorRepresentation
}

func (o *CancelProductOrderGetNotFound) Error() string {
	return fmt.Sprintf("[GET /cancelProductOrder/{CancelProductOrderId}][%d] cancelProductOrderGetNotFound  %+v", 404, o.Payload)
}

func (o *CancelProductOrderGetNotFound) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *CancelProductOrderGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelProductOrderGetMethodNotAllowed creates a CancelProductOrderGetMethodNotAllowed with default headers values
func NewCancelProductOrderGetMethodNotAllowed() *CancelProductOrderGetMethodNotAllowed {
	return &CancelProductOrderGetMethodNotAllowed{}
}

/*CancelProductOrderGetMethodNotAllowed handles this case with default header values.

Method Not Allowed

List of supported error codes:
- 61: Method not allowed
*/
type CancelProductOrderGetMethodNotAllowed struct {
	Payload *models.ErrorRepresentation
}

func (o *CancelProductOrderGetMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /cancelProductOrder/{CancelProductOrderId}][%d] cancelProductOrderGetMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *CancelProductOrderGetMethodNotAllowed) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *CancelProductOrderGetMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelProductOrderGetRequestTimeout creates a CancelProductOrderGetRequestTimeout with default headers values
func NewCancelProductOrderGetRequestTimeout() *CancelProductOrderGetRequestTimeout {
	return &CancelProductOrderGetRequestTimeout{}
}

/*CancelProductOrderGetRequestTimeout handles this case with default header values.

Request Time-out

List of supported error codes:
- 63: Request time-out
*/
type CancelProductOrderGetRequestTimeout struct {
	Payload *models.ErrorRepresentation
}

func (o *CancelProductOrderGetRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /cancelProductOrder/{CancelProductOrderId}][%d] cancelProductOrderGetRequestTimeout  %+v", 408, o.Payload)
}

func (o *CancelProductOrderGetRequestTimeout) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *CancelProductOrderGetRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelProductOrderGetUnprocessableEntity creates a CancelProductOrderGetUnprocessableEntity with default headers values
func NewCancelProductOrderGetUnprocessableEntity() *CancelProductOrderGetUnprocessableEntity {
	return &CancelProductOrderGetUnprocessableEntity{}
}

/*CancelProductOrderGetUnprocessableEntity handles this case with default header values.

Unprocessable entity

Functional error
*/
type CancelProductOrderGetUnprocessableEntity struct {
	Payload *models.ErrorRepresentation
}

func (o *CancelProductOrderGetUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /cancelProductOrder/{CancelProductOrderId}][%d] cancelProductOrderGetUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CancelProductOrderGetUnprocessableEntity) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *CancelProductOrderGetUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelProductOrderGetInternalServerError creates a CancelProductOrderGetInternalServerError with default headers values
func NewCancelProductOrderGetInternalServerError() *CancelProductOrderGetInternalServerError {
	return &CancelProductOrderGetInternalServerError{}
}

/*CancelProductOrderGetInternalServerError handles this case with default header values.

Internal Server Error

List of supported error codes:
- 1: Internal error
*/
type CancelProductOrderGetInternalServerError struct {
	Payload *models.ErrorRepresentation
}

func (o *CancelProductOrderGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /cancelProductOrder/{CancelProductOrderId}][%d] cancelProductOrderGetInternalServerError  %+v", 500, o.Payload)
}

func (o *CancelProductOrderGetInternalServerError) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *CancelProductOrderGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelProductOrderGetServiceUnavailable creates a CancelProductOrderGetServiceUnavailable with default headers values
func NewCancelProductOrderGetServiceUnavailable() *CancelProductOrderGetServiceUnavailable {
	return &CancelProductOrderGetServiceUnavailable{}
}

/*CancelProductOrderGetServiceUnavailable handles this case with default header values.

Service Unavailable


*/
type CancelProductOrderGetServiceUnavailable struct {
	Payload *models.ErrorRepresentation
}

func (o *CancelProductOrderGetServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /cancelProductOrder/{CancelProductOrderId}][%d] cancelProductOrderGetServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CancelProductOrderGetServiceUnavailable) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *CancelProductOrderGetServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
