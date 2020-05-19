// Code generated by go-swagger; DO NOT EDIT.

package product

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/qlcchain/go-lsobus/sonata/inventory/models"
)

// ProductGetReader is a Reader for the ProductGet structure.
type ProductGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProductGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProductGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProductGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProductGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProductGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProductGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewProductGetUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewProductGetServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProductGetOK creates a ProductGetOK with default headers values
func NewProductGetOK() *ProductGetOK {
	return &ProductGetOK{}
}

/*ProductGetOK handles this case with default header values.

Ok
*/
type ProductGetOK struct {
	Payload *models.Product
}

func (o *ProductGetOK) Error() string {
	return fmt.Sprintf("[GET /product/{ProductId}][%d] productGetOK  %+v", 200, o.Payload)
}

func (o *ProductGetOK) GetPayload() *models.Product {
	return o.Payload
}

func (o *ProductGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Product)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProductGetBadRequest creates a ProductGetBadRequest with default headers values
func NewProductGetBadRequest() *ProductGetBadRequest {
	return &ProductGetBadRequest{}
}

/*ProductGetBadRequest handles this case with default header values.

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
type ProductGetBadRequest struct {
	Payload *models.ErrorRepresentation
}

func (o *ProductGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /product/{ProductId}][%d] productGetBadRequest  %+v", 400, o.Payload)
}

func (o *ProductGetBadRequest) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *ProductGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProductGetUnauthorized creates a ProductGetUnauthorized with default headers values
func NewProductGetUnauthorized() *ProductGetUnauthorized {
	return &ProductGetUnauthorized{}
}

/*ProductGetUnauthorized handles this case with default header values.

Unauthorized

List of supported error codes:
- 40: Missing credentials
- 41: Invalid credentials
- 42: Expired credentials
*/
type ProductGetUnauthorized struct {
	Payload *models.ErrorRepresentation
}

func (o *ProductGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /product/{ProductId}][%d] productGetUnauthorized  %+v", 401, o.Payload)
}

func (o *ProductGetUnauthorized) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *ProductGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProductGetForbidden creates a ProductGetForbidden with default headers values
func NewProductGetForbidden() *ProductGetForbidden {
	return &ProductGetForbidden{}
}

/*ProductGetForbidden handles this case with default header values.

Forbidden

List of supported error codes:
- 50: Access denied
- 51: Forbidden requester
- 52: Forbidden user
- 53: Too many requests
*/
type ProductGetForbidden struct {
	Payload *models.ErrorRepresentation
}

func (o *ProductGetForbidden) Error() string {
	return fmt.Sprintf("[GET /product/{ProductId}][%d] productGetForbidden  %+v", 403, o.Payload)
}

func (o *ProductGetForbidden) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *ProductGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProductGetNotFound creates a ProductGetNotFound with default headers values
func NewProductGetNotFound() *ProductGetNotFound {
	return &ProductGetNotFound{}
}

/*ProductGetNotFound handles this case with default header values.

Not Found

List of supported error codes:
- 60: Resource not found
*/
type ProductGetNotFound struct {
	Payload *models.ErrorRepresentation
}

func (o *ProductGetNotFound) Error() string {
	return fmt.Sprintf("[GET /product/{ProductId}][%d] productGetNotFound  %+v", 404, o.Payload)
}

func (o *ProductGetNotFound) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *ProductGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProductGetUnprocessableEntity creates a ProductGetUnprocessableEntity with default headers values
func NewProductGetUnprocessableEntity() *ProductGetUnprocessableEntity {
	return &ProductGetUnprocessableEntity{}
}

/*ProductGetUnprocessableEntity handles this case with default header values.

Unprocessable entity

Functional error
*/
type ProductGetUnprocessableEntity struct {
	Payload *models.ErrorRepresentation
}

func (o *ProductGetUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /product/{ProductId}][%d] productGetUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ProductGetUnprocessableEntity) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *ProductGetUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProductGetServiceUnavailable creates a ProductGetServiceUnavailable with default headers values
func NewProductGetServiceUnavailable() *ProductGetServiceUnavailable {
	return &ProductGetServiceUnavailable{}
}

/*ProductGetServiceUnavailable handles this case with default header values.

Service Unavailable


*/
type ProductGetServiceUnavailable struct {
	Payload *models.ErrorRepresentation
}

func (o *ProductGetServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /product/{ProductId}][%d] productGetServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ProductGetServiceUnavailable) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *ProductGetServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
