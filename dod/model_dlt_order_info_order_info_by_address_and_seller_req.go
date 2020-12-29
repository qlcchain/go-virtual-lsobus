/*
 * QLC DOD Service
 *
 * REST Api for QLC DOD Service
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DltOrderInfoOrderInfoByAddressAndSellerReq struct {
	QlcAddressBuyer  string  `json:"qlcAddressBuyer"`
	QlcAddressSeller string  `json:"qlcAddressSeller"`
	Count            float64 `json:"count"`
	Offset           float64 `json:"offset"`
}
