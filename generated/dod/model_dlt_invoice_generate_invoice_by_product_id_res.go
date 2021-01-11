/*
 * QLC DOD Service
 *
 * REST Api for QLC DOD Service
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DltInvoiceGenerateInvoiceByProductIdRes struct {
	TotalAmount float64                                            `json:"totalAmount,omitempty"`
	Currency    string                                             `json:"currency,omitempty"`
	StartTime   float64                                            `json:"startTime,omitempty"`
	EndTime     float64                                            `json:"endTime,omitempty"`
	Buyer       *DltInvoiceGenerateInvoiceByOrderIdResBuyer        `json:"buyer,omitempty"`
	Connection  *DltInvoiceGenerateInvoiceByProductIdResConnection `json:"connection,omitempty"`
}