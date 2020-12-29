/*
 * QLC DOD Service
 *
 * REST Api for QLC DOD Service
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type QuoteQuoteItemSonataModel struct {
	Id                     string                             `json:"id,omitempty"`
	State                  string                             `json:"state,omitempty"`
	Action                 string                             `json:"action,omitempty"`
	Product                []QuoteProductModel                `json:"product,omitempty"`
	QuoteItemPrice         []QuoteQuoteItemPriceModel         `json:"quoteItemPrice,omitempty"`
	RequestedQuoteItemTerm []QuoteRequestedQuoteItemTermModel `json:"requestedQuoteItemTerm,omitempty"`
	QuoteItemTerm          []QuoteRequestedQuoteItemTermModel `json:"quoteItemTerm,omitempty"`
	QuoteItemRelationship  []QuoteItemRelationshipModel       `json:"quoteItemRelationship,omitempty"`
	Note                   []QuoteNoteModel                   `json:"note,omitempty"`
	Qualification          []QuoteQualificationModel          `json:"qualification,omitempty"`
	RelatedParty           []QuoteRelatedPartyModel           `json:"relatedParty,omitempty"`
}
