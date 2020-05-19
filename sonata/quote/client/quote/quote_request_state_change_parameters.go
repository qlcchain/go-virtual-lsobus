// Code generated by go-swagger; DO NOT EDIT.

package quote

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

	"github.com/qlcchain/go-lsobus/sonata/quote/models"
)

// NewQuoteRequestStateChangeParams creates a new QuoteRequestStateChangeParams object
// with the default values initialized.
func NewQuoteRequestStateChangeParams() *QuoteRequestStateChangeParams {
	var ()
	return &QuoteRequestStateChangeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewQuoteRequestStateChangeParamsWithTimeout creates a new QuoteRequestStateChangeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewQuoteRequestStateChangeParamsWithTimeout(timeout time.Duration) *QuoteRequestStateChangeParams {
	var ()
	return &QuoteRequestStateChangeParams{

		timeout: timeout,
	}
}

// NewQuoteRequestStateChangeParamsWithContext creates a new QuoteRequestStateChangeParams object
// with the default values initialized, and the ability to set a context for a request
func NewQuoteRequestStateChangeParamsWithContext(ctx context.Context) *QuoteRequestStateChangeParams {
	var ()
	return &QuoteRequestStateChangeParams{

		Context: ctx,
	}
}

// NewQuoteRequestStateChangeParamsWithHTTPClient creates a new QuoteRequestStateChangeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQuoteRequestStateChangeParamsWithHTTPClient(client *http.Client) *QuoteRequestStateChangeParams {
	var ()
	return &QuoteRequestStateChangeParams{
		HTTPClient: client,
	}
}

/*QuoteRequestStateChangeParams contains all the parameters to send to the API endpoint
for the quote request state change operation typically these are written to a http.Request
*/
type QuoteRequestStateChangeParams struct {

	/*ChangeQuoteStateRequest*/
	ChangeQuoteStateRequest *models.ChangelQuoteStateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the quote request state change params
func (o *QuoteRequestStateChangeParams) WithTimeout(timeout time.Duration) *QuoteRequestStateChangeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the quote request state change params
func (o *QuoteRequestStateChangeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the quote request state change params
func (o *QuoteRequestStateChangeParams) WithContext(ctx context.Context) *QuoteRequestStateChangeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the quote request state change params
func (o *QuoteRequestStateChangeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the quote request state change params
func (o *QuoteRequestStateChangeParams) WithHTTPClient(client *http.Client) *QuoteRequestStateChangeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the quote request state change params
func (o *QuoteRequestStateChangeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChangeQuoteStateRequest adds the changeQuoteStateRequest to the quote request state change params
func (o *QuoteRequestStateChangeParams) WithChangeQuoteStateRequest(changeQuoteStateRequest *models.ChangelQuoteStateRequest) *QuoteRequestStateChangeParams {
	o.SetChangeQuoteStateRequest(changeQuoteStateRequest)
	return o
}

// SetChangeQuoteStateRequest adds the changeQuoteStateRequest to the quote request state change params
func (o *QuoteRequestStateChangeParams) SetChangeQuoteStateRequest(changeQuoteStateRequest *models.ChangelQuoteStateRequest) {
	o.ChangeQuoteStateRequest = changeQuoteStateRequest
}

// WriteToRequest writes these params to a swagger request
func (o *QuoteRequestStateChangeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ChangeQuoteStateRequest != nil {
		if err := r.SetBodyParam(o.ChangeQuoteStateRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
