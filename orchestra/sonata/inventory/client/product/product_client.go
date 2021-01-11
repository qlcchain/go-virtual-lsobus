// Code generated by go-swagger; DO NOT EDIT.

package product

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new product API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for product API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ProductFind(params *ProductFindParams) (*ProductFindOK, error)

	ProductGet(params *ProductGetParams) (*ProductGetOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ProductFind products find list retrieve product list with summary view

  The Buyer requests a list of Products from the Seller based on filter criteria.
*/
func (a *Client) ProductFind(params *ProductFindParams) (*ProductFindOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProductFindParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "productFind",
		Method:             "GET",
		PathPattern:        "/product",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProductFindReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProductFindOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for productFind: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ProductGet products get by id retrieve one product with all information

  The Buyer requests the details associated with a single Product based on a Seller Product Identifier.
*/
func (a *Client) ProductGet(params *ProductGetParams) (*ProductGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProductGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "productGet",
		Method:             "GET",
		PathPattern:        "/product/{ProductId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProductGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProductGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for productGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}