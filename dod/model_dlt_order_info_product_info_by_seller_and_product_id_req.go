/*
 * QLC DOD Service
 *
 * REST Api for QLC DOD Service
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DltOrderInfoProductInfoBySellerAndProductIdReq struct {
	QlcAddressSeller string `json:"qlcAddressSeller"`
	ProductId        string `json:"productId"`
}
