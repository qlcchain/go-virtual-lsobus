/*
 * QLC DOD Service
 *
 * REST Api for QLC DOD Service
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DltOrderInfoProductInfoBySellerAndProductIdResData struct {
	ItemId            string                                                  `json:"itemId,omitempty"`
	BuyerProductId    string                                                  `json:"buyerProductId,omitempty"`
	ProductOfferingId string                                                  `json:"productOfferingId,omitempty"`
	ProductId         string                                                  `json:"productId,omitempty"`
	SrcCompanyName    string                                                  `json:"srcCompanyName,omitempty"`
	SrcRegion         string                                                  `json:"srcRegion,omitempty"`
	SrcCity           string                                                  `json:"srcCity,omitempty"`
	SrcDataCenter     string                                                  `json:"srcDataCenter,omitempty"`
	SrcPort           string                                                  `json:"srcPort,omitempty"`
	DstCompanyName    string                                                  `json:"dstCompanyName,omitempty"`
	DstRegion         string                                                  `json:"dstRegion,omitempty"`
	DstCity           string                                                  `json:"dstCity,omitempty"`
	DstDataCenter     string                                                  `json:"dstDataCenter,omitempty"`
	DstPort           string                                                  `json:"dstPort,omitempty"`
	Active            []DltOrderInfoProductInfoBySellerAndProductIdDoneModel  `json:"active,omitempty"`
	Done              []DltOrderInfoProductInfoBySellerAndProductIdDoneModel  `json:"done,omitempty"`
	Track             []DltOrderInfoProductInfoBySellerAndProductIdTrackModel `json:"track,omitempty"`
}
