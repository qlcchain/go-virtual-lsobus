/*
 * QLC DOD Service
 *
 * REST Api for QLC DOD Service
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type QualificationValidateAddressModelFieldedAddress struct {
	Id                   string                                                               `json:"id,omitempty"`
	StreetName           string                                                               `json:"streetName,omitempty"`
	Postcode             string                                                               `json:"postcode,omitempty"`
	Locality             string                                                               `json:"locality,omitempty"`
	City                 string                                                               `json:"city,omitempty"`
	StateOrProvince      string                                                               `json:"stateOrProvince,omitempty"`
	Country              string                                                               `json:"country,omitempty"`
	GeographicSubAddress *QualificationValidateAddressModelFieldedAddressGeographicSubAddress `json:"geographicSubAddress,omitempty"`
}