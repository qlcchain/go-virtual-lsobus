/*
 * QLC DOD Service
 *
 * REST Api for QLC DOD Service
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DltOrderInfoOrderDetailTrackModel struct {
	ContractState string  `json:"contractState,omitempty"`
	OrderState    string  `json:"orderState,omitempty"`
	Reason        string  `json:"reason,omitempty"`
	Time          float64 `json:"time,omitempty"`
	Hash          string  `json:"hash,omitempty"`
}
