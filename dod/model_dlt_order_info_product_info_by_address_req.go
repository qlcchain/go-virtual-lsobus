/*
 * QLC DOD Service
 *
 * REST Api for QLC DOD Service
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DltOrderInfoProductInfoByAddressReq struct {
	QlcAddressBuyer string  `json:"qlcAddressBuyer"`
	Count           float64 `json:"count"`
	Offset          float64 `json:"offset"`
}
