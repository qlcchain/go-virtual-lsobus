/*
 * QLC DOD Service
 *
 * REST Api for QLC DOD Service
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type OrderCreateReqBuyer struct {
	Address string `json:"address"`
	Seed    string `json:"seed"`
	Name    string `json:"name"`
}
