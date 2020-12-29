/*
 * QLC DOD Service
 *
 * REST Api for QLC DOD Service
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DltInvoiceGenerateInvoiceByProductIdModel struct {
	OrderId          string  `json:"orderId,omitempty"`
	InternalId       string  `json:"internalId,omitempty"`
	ConnectionName   string  `json:"connectionName,omitempty"`
	PaymentType      string  `json:"paymentType,omitempty"`
	BillingType      string  `json:"billingType,omitempty"`
	Currency         string  `json:"currency,omitempty"`
	ServiceClass     string  `json:"serviceClass,omitempty"`
	Bandwidth        string  `json:"bandwidth,omitempty"`
	Price            float64 `json:"price,omitempty"`
	StartTime        float64 `json:"startTime,omitempty"`
	EndTime          float64 `json:"endTime,omitempty"`
	InvoiceStartTime float64 `json:"invoiceStartTime,omitempty"`
	InvoiceEndTime   float64 `json:"invoiceEndTime,omitempty"`
	InvoiceUnitCount float64 `json:"invoiceUnitCount,omitempty"`
	Amount           float64 `json:"amount,omitempty"`
}
