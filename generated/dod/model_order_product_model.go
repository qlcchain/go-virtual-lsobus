/*
 * QLC DOD Service
 *
 * REST Api for QLC DOD Service
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type OrderProductModel struct {
	Id                   string                                 `json:"id,omitempty"`
	ProductSpecification *OrderProductModelProductSpecification `json:"productSpecification,omitempty"`
	ProductRelationship  []OrderProductRelationshipModel        `json:"productRelationship,omitempty"`
	Place                []OrderPlaceModel                      `json:"place,omitempty"`
}