/*
 * QLC DOD Service
 *
 * REST Api for QLC DOD Service
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DltOrderInfoOrderInfoByAddressResData struct {
	TotalOrders float64                         `json:"totalOrders,omitempty"`
	OrderInfo   []DltOrderInfoOrderDetailResRpc `json:"orderInfo,omitempty"`
}