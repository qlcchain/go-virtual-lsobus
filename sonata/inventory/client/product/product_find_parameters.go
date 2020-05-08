// Code generated by go-swagger; DO NOT EDIT.

package product

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

// NewProductFindParams creates a new ProductFindParams object
// with the default values initialized.
func NewProductFindParams() *ProductFindParams {
	var ()
	return &ProductFindParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProductFindParamsWithTimeout creates a new ProductFindParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProductFindParamsWithTimeout(timeout time.Duration) *ProductFindParams {
	var ()
	return &ProductFindParams{

		timeout: timeout,
	}
}

// NewProductFindParamsWithContext creates a new ProductFindParams object
// with the default values initialized, and the ability to set a context for a request
func NewProductFindParamsWithContext(ctx context.Context) *ProductFindParams {
	var ()
	return &ProductFindParams{

		Context: ctx,
	}
}

// NewProductFindParamsWithHTTPClient creates a new ProductFindParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProductFindParamsWithHTTPClient(client *http.Client) *ProductFindParams {
	var ()
	return &ProductFindParams{
		HTTPClient: client,
	}
}

/*ProductFindParams contains all the parameters to send to the API endpoint
for the product find operation typically these are written to a http.Request
*/
type ProductFindParams struct {

	/*BillingAccountID
	  The Billing Account associated with the Product.

	*/
	BillingAccountID *string
	/*BuyerID
	  Identifier of the party who role is buyer.

	*/
	BuyerID *string
	/*BuyerProductID
	  A reference to the buyerProductId provided in the order

	*/
	BuyerProductID *string
	/*GeographicalSiteID
	  A site identifier which is associated to the product

	*/
	GeographicalSiteID *string
	/*LastUpdateDateGt
	  Greater than the date that the last change affecting this Product was complet-ed

	*/
	LastUpdateDateGt *strfmt.DateTime
	/*LastUpdateDateLt
	  Less than date that the last change affecting this Product was completed

	*/
	LastUpdateDateLt *strfmt.DateTime
	/*Limit
	  Requested number of resources to be provided in response requested by client

	*/
	Limit *string
	/*Offset
	  Requested index for start of resources to be provided in response requested by client

	*/
	Offset *string
	/*ProductOfferingID
	  A reference to a product offering by id

	*/
	ProductOfferingID *string
	/*ProductOrderID
	  Identifies Product Order(s) associated with the Product

	*/
	ProductOrderID *string
	/*ProductSpecificationID
	  A reference to a product spec by id

	*/
	ProductSpecificationID *string
	/*RelatedProductID
	  This criteria allows to retrieve all Product records with a Product Relationship to a specified Product.
	E.g. All Products related to Product with ID 5

	*/
	RelatedProductID *string
	/*StartDateGt
	  Greater than the date that is the initial install date for the Product

	*/
	StartDateGt *strfmt.DateTime
	/*StartDateLt
	  Less than the date that is the initial install date for the Product

	*/
	StartDateLt *strfmt.DateTime
	/*Status
	  The status of the product

	*/
	Status *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the product find params
func (o *ProductFindParams) WithTimeout(timeout time.Duration) *ProductFindParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the product find params
func (o *ProductFindParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the product find params
func (o *ProductFindParams) WithContext(ctx context.Context) *ProductFindParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the product find params
func (o *ProductFindParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the product find params
func (o *ProductFindParams) WithHTTPClient(client *http.Client) *ProductFindParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the product find params
func (o *ProductFindParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBillingAccountID adds the billingAccountID to the product find params
func (o *ProductFindParams) WithBillingAccountID(billingAccountID *string) *ProductFindParams {
	o.SetBillingAccountID(billingAccountID)
	return o
}

// SetBillingAccountID adds the billingAccountId to the product find params
func (o *ProductFindParams) SetBillingAccountID(billingAccountID *string) {
	o.BillingAccountID = billingAccountID
}

// WithBuyerID adds the buyerID to the product find params
func (o *ProductFindParams) WithBuyerID(buyerID *string) *ProductFindParams {
	o.SetBuyerID(buyerID)
	return o
}

// SetBuyerID adds the buyerId to the product find params
func (o *ProductFindParams) SetBuyerID(buyerID *string) {
	o.BuyerID = buyerID
}

// WithBuyerProductID adds the buyerProductID to the product find params
func (o *ProductFindParams) WithBuyerProductID(buyerProductID *string) *ProductFindParams {
	o.SetBuyerProductID(buyerProductID)
	return o
}

// SetBuyerProductID adds the buyerProductId to the product find params
func (o *ProductFindParams) SetBuyerProductID(buyerProductID *string) {
	o.BuyerProductID = buyerProductID
}

// WithGeographicalSiteID adds the geographicalSiteID to the product find params
func (o *ProductFindParams) WithGeographicalSiteID(geographicalSiteID *string) *ProductFindParams {
	o.SetGeographicalSiteID(geographicalSiteID)
	return o
}

// SetGeographicalSiteID adds the geographicalSiteId to the product find params
func (o *ProductFindParams) SetGeographicalSiteID(geographicalSiteID *string) {
	o.GeographicalSiteID = geographicalSiteID
}

// WithLastUpdateDateGt adds the lastUpdateDateGt to the product find params
func (o *ProductFindParams) WithLastUpdateDateGt(lastUpdateDateGt *strfmt.DateTime) *ProductFindParams {
	o.SetLastUpdateDateGt(lastUpdateDateGt)
	return o
}

// SetLastUpdateDateGt adds the lastUpdateDateGt to the product find params
func (o *ProductFindParams) SetLastUpdateDateGt(lastUpdateDateGt *strfmt.DateTime) {
	o.LastUpdateDateGt = lastUpdateDateGt
}

// WithLastUpdateDateLt adds the lastUpdateDateLt to the product find params
func (o *ProductFindParams) WithLastUpdateDateLt(lastUpdateDateLt *strfmt.DateTime) *ProductFindParams {
	o.SetLastUpdateDateLt(lastUpdateDateLt)
	return o
}

// SetLastUpdateDateLt adds the lastUpdateDateLt to the product find params
func (o *ProductFindParams) SetLastUpdateDateLt(lastUpdateDateLt *strfmt.DateTime) {
	o.LastUpdateDateLt = lastUpdateDateLt
}

// WithLimit adds the limit to the product find params
func (o *ProductFindParams) WithLimit(limit *string) *ProductFindParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the product find params
func (o *ProductFindParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithOffset adds the offset to the product find params
func (o *ProductFindParams) WithOffset(offset *string) *ProductFindParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the product find params
func (o *ProductFindParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WithProductOfferingID adds the productOfferingID to the product find params
func (o *ProductFindParams) WithProductOfferingID(productOfferingID *string) *ProductFindParams {
	o.SetProductOfferingID(productOfferingID)
	return o
}

// SetProductOfferingID adds the productOfferingId to the product find params
func (o *ProductFindParams) SetProductOfferingID(productOfferingID *string) {
	o.ProductOfferingID = productOfferingID
}

// WithProductOrderID adds the productOrderID to the product find params
func (o *ProductFindParams) WithProductOrderID(productOrderID *string) *ProductFindParams {
	o.SetProductOrderID(productOrderID)
	return o
}

// SetProductOrderID adds the productOrderId to the product find params
func (o *ProductFindParams) SetProductOrderID(productOrderID *string) {
	o.ProductOrderID = productOrderID
}

// WithProductSpecificationID adds the productSpecificationID to the product find params
func (o *ProductFindParams) WithProductSpecificationID(productSpecificationID *string) *ProductFindParams {
	o.SetProductSpecificationID(productSpecificationID)
	return o
}

// SetProductSpecificationID adds the productSpecificationId to the product find params
func (o *ProductFindParams) SetProductSpecificationID(productSpecificationID *string) {
	o.ProductSpecificationID = productSpecificationID
}

// WithRelatedProductID adds the relatedProductID to the product find params
func (o *ProductFindParams) WithRelatedProductID(relatedProductID *string) *ProductFindParams {
	o.SetRelatedProductID(relatedProductID)
	return o
}

// SetRelatedProductID adds the relatedProductId to the product find params
func (o *ProductFindParams) SetRelatedProductID(relatedProductID *string) {
	o.RelatedProductID = relatedProductID
}

// WithStartDateGt adds the startDateGt to the product find params
func (o *ProductFindParams) WithStartDateGt(startDateGt *strfmt.DateTime) *ProductFindParams {
	o.SetStartDateGt(startDateGt)
	return o
}

// SetStartDateGt adds the startDateGt to the product find params
func (o *ProductFindParams) SetStartDateGt(startDateGt *strfmt.DateTime) {
	o.StartDateGt = startDateGt
}

// WithStartDateLt adds the startDateLt to the product find params
func (o *ProductFindParams) WithStartDateLt(startDateLt *strfmt.DateTime) *ProductFindParams {
	o.SetStartDateLt(startDateLt)
	return o
}

// SetStartDateLt adds the startDateLt to the product find params
func (o *ProductFindParams) SetStartDateLt(startDateLt *strfmt.DateTime) {
	o.StartDateLt = startDateLt
}

// WithStatus adds the status to the product find params
func (o *ProductFindParams) WithStatus(status *string) *ProductFindParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the product find params
func (o *ProductFindParams) SetStatus(status *string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *ProductFindParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BillingAccountID != nil {

		// query param billingAccountId
		var qrBillingAccountID string
		if o.BillingAccountID != nil {
			qrBillingAccountID = *o.BillingAccountID
		}
		qBillingAccountID := qrBillingAccountID
		if qBillingAccountID != "" {
			if err := r.SetQueryParam("billingAccountId", qBillingAccountID); err != nil {
				return err
			}
		}

	}

	if o.BuyerID != nil {

		// query param buyerId
		var qrBuyerID string
		if o.BuyerID != nil {
			qrBuyerID = *o.BuyerID
		}
		qBuyerID := qrBuyerID
		if qBuyerID != "" {
			if err := r.SetQueryParam("buyerId", qBuyerID); err != nil {
				return err
			}
		}

	}

	if o.BuyerProductID != nil {

		// query param buyerProductId
		var qrBuyerProductID string
		if o.BuyerProductID != nil {
			qrBuyerProductID = *o.BuyerProductID
		}
		qBuyerProductID := qrBuyerProductID
		if qBuyerProductID != "" {
			if err := r.SetQueryParam("buyerProductId", qBuyerProductID); err != nil {
				return err
			}
		}

	}

	if o.GeographicalSiteID != nil {

		// query param geographicalSiteId
		var qrGeographicalSiteID string
		if o.GeographicalSiteID != nil {
			qrGeographicalSiteID = *o.GeographicalSiteID
		}
		qGeographicalSiteID := qrGeographicalSiteID
		if qGeographicalSiteID != "" {
			if err := r.SetQueryParam("geographicalSiteId", qGeographicalSiteID); err != nil {
				return err
			}
		}

	}

	if o.LastUpdateDateGt != nil {

		// query param lastUpdateDate.gt
		var qrLastUpdateDateGt strfmt.DateTime
		if o.LastUpdateDateGt != nil {
			qrLastUpdateDateGt = *o.LastUpdateDateGt
		}
		qLastUpdateDateGt := qrLastUpdateDateGt.String()
		if qLastUpdateDateGt != "" {
			if err := r.SetQueryParam("lastUpdateDate.gt", qLastUpdateDateGt); err != nil {
				return err
			}
		}

	}

	if o.LastUpdateDateLt != nil {

		// query param lastUpdateDate.lt
		var qrLastUpdateDateLt strfmt.DateTime
		if o.LastUpdateDateLt != nil {
			qrLastUpdateDateLt = *o.LastUpdateDateLt
		}
		qLastUpdateDateLt := qrLastUpdateDateLt.String()
		if qLastUpdateDateLt != "" {
			if err := r.SetQueryParam("lastUpdateDate.lt", qLastUpdateDateLt); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit string
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := qrLimit
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset string
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := qrOffset
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.ProductOfferingID != nil {

		// query param productOfferingId
		var qrProductOfferingID string
		if o.ProductOfferingID != nil {
			qrProductOfferingID = *o.ProductOfferingID
		}
		qProductOfferingID := qrProductOfferingID
		if qProductOfferingID != "" {
			if err := r.SetQueryParam("productOfferingId", qProductOfferingID); err != nil {
				return err
			}
		}

	}

	if o.ProductOrderID != nil {

		// query param productOrderId
		var qrProductOrderID string
		if o.ProductOrderID != nil {
			qrProductOrderID = *o.ProductOrderID
		}
		qProductOrderID := qrProductOrderID
		if qProductOrderID != "" {
			if err := r.SetQueryParam("productOrderId", qProductOrderID); err != nil {
				return err
			}
		}

	}

	if o.ProductSpecificationID != nil {

		// query param productSpecificationId
		var qrProductSpecificationID string
		if o.ProductSpecificationID != nil {
			qrProductSpecificationID = *o.ProductSpecificationID
		}
		qProductSpecificationID := qrProductSpecificationID
		if qProductSpecificationID != "" {
			if err := r.SetQueryParam("productSpecificationId", qProductSpecificationID); err != nil {
				return err
			}
		}

	}

	if o.RelatedProductID != nil {

		// query param relatedProductId
		var qrRelatedProductID string
		if o.RelatedProductID != nil {
			qrRelatedProductID = *o.RelatedProductID
		}
		qRelatedProductID := qrRelatedProductID
		if qRelatedProductID != "" {
			if err := r.SetQueryParam("relatedProductId", qRelatedProductID); err != nil {
				return err
			}
		}

	}

	if o.StartDateGt != nil {

		// query param startDate.gt
		var qrStartDateGt strfmt.DateTime
		if o.StartDateGt != nil {
			qrStartDateGt = *o.StartDateGt
		}
		qStartDateGt := qrStartDateGt.String()
		if qStartDateGt != "" {
			if err := r.SetQueryParam("startDate.gt", qStartDateGt); err != nil {
				return err
			}
		}

	}

	if o.StartDateLt != nil {

		// query param startDate.lt
		var qrStartDateLt strfmt.DateTime
		if o.StartDateLt != nil {
			qrStartDateLt = *o.StartDateLt
		}
		qStartDateLt := qrStartDateLt.String()
		if qStartDateLt != "" {
			if err := r.SetQueryParam("startDate.lt", qStartDateLt); err != nil {
				return err
			}
		}

	}

	if o.Status != nil {

		// query param status
		var qrStatus string
		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {
			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
