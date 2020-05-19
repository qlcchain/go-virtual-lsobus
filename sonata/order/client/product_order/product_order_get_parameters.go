// Code generated by go-swagger; DO NOT EDIT.

package product_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewProductOrderGetParams creates a new ProductOrderGetParams object
// with the default values initialized.
func NewProductOrderGetParams() *ProductOrderGetParams {
	var ()
	return &ProductOrderGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProductOrderGetParamsWithTimeout creates a new ProductOrderGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProductOrderGetParamsWithTimeout(timeout time.Duration) *ProductOrderGetParams {
	var ()
	return &ProductOrderGetParams{

		timeout: timeout,
	}
}

// NewProductOrderGetParamsWithContext creates a new ProductOrderGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewProductOrderGetParamsWithContext(ctx context.Context) *ProductOrderGetParams {
	var ()
	return &ProductOrderGetParams{

		Context: ctx,
	}
}

// NewProductOrderGetParamsWithHTTPClient creates a new ProductOrderGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProductOrderGetParamsWithHTTPClient(client *http.Client) *ProductOrderGetParams {
	var ()
	return &ProductOrderGetParams{
		HTTPClient: client,
	}
}

/*ProductOrderGetParams contains all the parameters to send to the API endpoint
for the product order get operation typically these are written to a http.Request
*/
type ProductOrderGetParams struct {

	/*ProductOrderID*/
	ProductOrderID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the product order get params
func (o *ProductOrderGetParams) WithTimeout(timeout time.Duration) *ProductOrderGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the product order get params
func (o *ProductOrderGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the product order get params
func (o *ProductOrderGetParams) WithContext(ctx context.Context) *ProductOrderGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the product order get params
func (o *ProductOrderGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the product order get params
func (o *ProductOrderGetParams) WithHTTPClient(client *http.Client) *ProductOrderGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the product order get params
func (o *ProductOrderGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProductOrderID adds the productOrderID to the product order get params
func (o *ProductOrderGetParams) WithProductOrderID(productOrderID string) *ProductOrderGetParams {
	o.SetProductOrderID(productOrderID)
	return o
}

// SetProductOrderID adds the productOrderId to the product order get params
func (o *ProductOrderGetParams) SetProductOrderID(productOrderID string) {
	o.ProductOrderID = productOrderID
}

// WriteToRequest writes these params to a swagger request
func (o *ProductOrderGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ProductOrderId
	if err := r.SetPathParam("ProductOrderId", o.ProductOrderID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}