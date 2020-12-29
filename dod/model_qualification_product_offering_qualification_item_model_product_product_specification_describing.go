/*
 * QLC DOD Service
 *
 * REST Api for QLC DOD Service
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type QualificationProductOfferingQualificationItemModelProductProductSpecificationDescribing struct {
	PhysicalLayer        []string                                                                                                     `json:"physicalLayer,omitempty"`
	ENNIIngressBWProfile *QualificationProductOfferingQualificationItemModelProductProductSpecificationDescribingEnniIngressBwProfile `json:"ENNIIngressBWProfile,omitempty"`
	InterconnectionType  string                                                                                                       `json:"interconnectionType,omitempty"`
	SVlanId              string                                                                                                       `json:"sVlanId,omitempty"`
}
