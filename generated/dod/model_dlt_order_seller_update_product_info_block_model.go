/*
 * QLC DOD Service
 *
 * REST Api for QLC DOD Service
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DltOrderSellerUpdateProductInfoBlockModel struct {
	OrderItemId string `json:"orderItemId,omitempty"`
	ProductId   string `json:"productId,omitempty"`
	Active      string `json:"active,omitempty"`
}