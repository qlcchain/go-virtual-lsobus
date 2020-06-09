// Code generated by go-swagger; DO NOT EDIT.

package orchestra_api

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

	"github.com/qlcchain/go-lsobus/cmd/agent/models"
)

// NewExecCreateParams creates a new ExecCreateParams object
// with the default values initialized.
func NewExecCreateParams() *ExecCreateParams {
	var ()
	return &ExecCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExecCreateParamsWithTimeout creates a new ExecCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExecCreateParamsWithTimeout(timeout time.Duration) *ExecCreateParams {
	var ()
	return &ExecCreateParams{

		timeout: timeout,
	}
}

// NewExecCreateParamsWithContext creates a new ExecCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewExecCreateParamsWithContext(ctx context.Context) *ExecCreateParams {
	var ()
	return &ExecCreateParams{

		Context: ctx,
	}
}

// NewExecCreateParamsWithHTTPClient creates a new ExecCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExecCreateParamsWithHTTPClient(client *http.Client) *ExecCreateParams {
	var ()
	return &ExecCreateParams{
		HTTPClient: client,
	}
}

/*ExecCreateParams contains all the parameters to send to the API endpoint
for the exec create operation typically these are written to a http.Request
*/
type ExecCreateParams struct {

	/*Body*/
	Body *models.ProtoOrchestraCommonRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the exec create params
func (o *ExecCreateParams) WithTimeout(timeout time.Duration) *ExecCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the exec create params
func (o *ExecCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the exec create params
func (o *ExecCreateParams) WithContext(ctx context.Context) *ExecCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the exec create params
func (o *ExecCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the exec create params
func (o *ExecCreateParams) WithHTTPClient(client *http.Client) *ExecCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the exec create params
func (o *ExecCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the exec create params
func (o *ExecCreateParams) WithBody(body *models.ProtoOrchestraCommonRequest) *ExecCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the exec create params
func (o *ExecCreateParams) SetBody(body *models.ProtoOrchestraCommonRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ExecCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}