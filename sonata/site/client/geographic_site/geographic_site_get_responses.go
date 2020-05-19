// Code generated by go-swagger; DO NOT EDIT.

package geographic_site

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/qlcchain/go-lsobus/sonata/site/models"
)

// GeographicSiteGetReader is a Reader for the GeographicSiteGet structure.
type GeographicSiteGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GeographicSiteGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGeographicSiteGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGeographicSiteGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGeographicSiteGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGeographicSiteGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGeographicSiteGetRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGeographicSiteGetUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGeographicSiteGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGeographicSiteGetServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGeographicSiteGetOK creates a GeographicSiteGetOK with default headers values
func NewGeographicSiteGetOK() *GeographicSiteGetOK {
	return &GeographicSiteGetOK{}
}

/*GeographicSiteGetOK handles this case with default header values.

Ok
*/
type GeographicSiteGetOK struct {
	Payload *models.GeographicSite
}

func (o *GeographicSiteGetOK) Error() string {
	return fmt.Sprintf("[GET /geographicSite/{SiteId}][%d] geographicSiteGetOK  %+v", 200, o.Payload)
}

func (o *GeographicSiteGetOK) GetPayload() *models.GeographicSite {
	return o.Payload
}

func (o *GeographicSiteGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeographicSite)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGeographicSiteGetBadRequest creates a GeographicSiteGetBadRequest with default headers values
func NewGeographicSiteGetBadRequest() *GeographicSiteGetBadRequest {
	return &GeographicSiteGetBadRequest{}
}

/*GeographicSiteGetBadRequest handles this case with default header values.

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
type GeographicSiteGetBadRequest struct {
	Payload *models.ErrorRepresentation
}

func (o *GeographicSiteGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /geographicSite/{SiteId}][%d] geographicSiteGetBadRequest  %+v", 400, o.Payload)
}

func (o *GeographicSiteGetBadRequest) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *GeographicSiteGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGeographicSiteGetUnauthorized creates a GeographicSiteGetUnauthorized with default headers values
func NewGeographicSiteGetUnauthorized() *GeographicSiteGetUnauthorized {
	return &GeographicSiteGetUnauthorized{}
}

/*GeographicSiteGetUnauthorized handles this case with default header values.

Unauthorized

List of supported error codes:
- 40: Missing credentials
- 41: Invalid credentials
- 42: Expired credentials
*/
type GeographicSiteGetUnauthorized struct {
	Payload *models.ErrorRepresentation
}

func (o *GeographicSiteGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /geographicSite/{SiteId}][%d] geographicSiteGetUnauthorized  %+v", 401, o.Payload)
}

func (o *GeographicSiteGetUnauthorized) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *GeographicSiteGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGeographicSiteGetNotFound creates a GeographicSiteGetNotFound with default headers values
func NewGeographicSiteGetNotFound() *GeographicSiteGetNotFound {
	return &GeographicSiteGetNotFound{}
}

/*GeographicSiteGetNotFound handles this case with default header values.

Not Found

List of supported error codes:
- 60: Resource not found
*/
type GeographicSiteGetNotFound struct {
	Payload *models.ErrorRepresentation
}

func (o *GeographicSiteGetNotFound) Error() string {
	return fmt.Sprintf("[GET /geographicSite/{SiteId}][%d] geographicSiteGetNotFound  %+v", 404, o.Payload)
}

func (o *GeographicSiteGetNotFound) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *GeographicSiteGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGeographicSiteGetRequestTimeout creates a GeographicSiteGetRequestTimeout with default headers values
func NewGeographicSiteGetRequestTimeout() *GeographicSiteGetRequestTimeout {
	return &GeographicSiteGetRequestTimeout{}
}

/*GeographicSiteGetRequestTimeout handles this case with default header values.

Request Time-out

List of supported error codes:
- 63: Request time-out
*/
type GeographicSiteGetRequestTimeout struct {
	Payload *models.ErrorRepresentation
}

func (o *GeographicSiteGetRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /geographicSite/{SiteId}][%d] geographicSiteGetRequestTimeout  %+v", 408, o.Payload)
}

func (o *GeographicSiteGetRequestTimeout) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *GeographicSiteGetRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGeographicSiteGetUnprocessableEntity creates a GeographicSiteGetUnprocessableEntity with default headers values
func NewGeographicSiteGetUnprocessableEntity() *GeographicSiteGetUnprocessableEntity {
	return &GeographicSiteGetUnprocessableEntity{}
}

/*GeographicSiteGetUnprocessableEntity handles this case with default header values.

Unprocessable entity

Functional error
*/
type GeographicSiteGetUnprocessableEntity struct {
	Payload *models.ErrorRepresentation
}

func (o *GeographicSiteGetUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /geographicSite/{SiteId}][%d] geographicSiteGetUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GeographicSiteGetUnprocessableEntity) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *GeographicSiteGetUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGeographicSiteGetInternalServerError creates a GeographicSiteGetInternalServerError with default headers values
func NewGeographicSiteGetInternalServerError() *GeographicSiteGetInternalServerError {
	return &GeographicSiteGetInternalServerError{}
}

/*GeographicSiteGetInternalServerError handles this case with default header values.

Internal Server Error

List of supported error codes:
- 1: Internal error
*/
type GeographicSiteGetInternalServerError struct {
	Payload *models.ErrorRepresentation
}

func (o *GeographicSiteGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /geographicSite/{SiteId}][%d] geographicSiteGetInternalServerError  %+v", 500, o.Payload)
}

func (o *GeographicSiteGetInternalServerError) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *GeographicSiteGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGeographicSiteGetServiceUnavailable creates a GeographicSiteGetServiceUnavailable with default headers values
func NewGeographicSiteGetServiceUnavailable() *GeographicSiteGetServiceUnavailable {
	return &GeographicSiteGetServiceUnavailable{}
}

/*GeographicSiteGetServiceUnavailable handles this case with default header values.

Service Unavailable


*/
type GeographicSiteGetServiceUnavailable struct {
	Payload *models.ErrorRepresentation
}

func (o *GeographicSiteGetServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /geographicSite/{SiteId}][%d] geographicSiteGetServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GeographicSiteGetServiceUnavailable) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *GeographicSiteGetServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
