package dod

import (
	"encoding/json"
	"testing"

	sw "github.com/qlcchain/go-lsobus/generated/dod"
)

func Test_quoteItem2Quote(t *testing.T) {
	data := `{"apimXUserId":"587055","httpStatus":"OK","id":"7a1407e8f50b4f7c86636fc354400f48","href":"https://test-api.hgc.com.hk/quote/v1/quotes/7a1407e8f50b4f7c86636fc354400f48","projectId":"MyProject-03","description":"","state":"READY","quoteDate":"2021-01-16T12:43:02.006Z","instantSyncQuoting":true,"quoteLevel":"FIRM","requestedQuoteCompletionDate":"2021-01-19T08:19:10.528Z","expectedQuoteCompletionDate":"2021-01-17T12:43:02.006Z","expectedFulfillmentStartDate":"","effectiveQuoteCompletionDate":"2021-01-16T12:43:02.006Z","validFor":{"startDate":"2021-01-16T12:43:02.006Z","endDate":"2021-02-16T12:43:02.006Z"},"note":[{"date":"2021-01-16T12:43:02.120Z","author":"HGC","text":"Interconnection Type = PTP, Interface Type = FE"}],"agreement":[{"id":"1","href":"","name":"Buyer","path":"/"}],"relatedParty":[{"id":"1","role":["buyer"],"name":"Buyer","emailAddress":"buyer@me.com","number":"12123","numberExtension":"1"}],"quoteItem":[{"id":"1","state":"READY","action":"INSTALL","product":{"id":"PRD0017HCVF44ED","productSpecification":{"id":"UNISpec","describing":{"physicalLayer":["1000BASE-LX"],"interconnectionType":"GatewayInterconnect"}},"productRelationship":[{"type":"UNISpec","product":{"id":"QLINK01"}}],"place":[{"id":"32443101","role":"UNI_LOCATION","@type":"FieldedAddress"}]},"quoteItemPrice":[{"priceType":"RECURRING","recurringChargePeriod":"MONTH","name":"Quote Item Price","price":{"preTaxAmount":{"value":0,"unit":"USD"}}},{"priceType":"NON_RECURRING","price":{"preTaxAmount":{"value":0,"unit":"USD"}}}],"requestedQuoteItemTerm":{"name":"","description":"","duration":{"value":12,"unit":"MONTH"}},"quoteItemTerm":{"name":"","description":"","duration":{"value":12,"unit":"MONTH"}},"quoteItemRelationship":[{"type":"RELIES_ON","id":""}],"note":[{"date":"2021-01-12T12:08:22.936Z","author":"Sample Note Author 1","text":""}],"qualification":[{"id":"8e32872f156b4f25814720d17f2fc0c8","href":"","qualificationItem":"1"}],"relatedParty":[{"id":"1","role":["buyer"],"name":"Buyer","emailAddress":"buyer@buyer.com","number":"213123","numberExtension":"1","@referredType":""}]},{"id":"2","state":"READY","action":"INSTALL","product":{"id":"PRD00FXUZY1GCVE","href":"","productSpecification":{"id":"ELineSpec","describing":{"sVlanId":"123","ENNIIngressBWProfile":{"cir":{"amount":100,"unit":"Mbps"}},"UNIIngressBWProfile":{"cir":{"amount":100,"unit":"Mbps"}}}},"productRelationship":[{"type":"RELIES_ON","product":{"id":"QLINK01"}}],"place":[]},"quoteItemPrice":[{"priceType":"RECURRING","recurringChargePeriod":"MONTH","name":"Quote Item Price for METRONET_METROPORT. Interconnection Type = PTP, Interface Type = FE","price":{"preTaxAmount":{"value":396,"unit":"USD"}}},{"priceType":"NON_RECURRING","price":{"preTaxAmount":{"value":0,"unit":"USD"}}}],"requestedQuoteItemTerm":{"name":"","description":"","duration":{"value":12,"unit":"MONTH"}},"quoteItemTerm":{"name":"","description":"","duration":{"value":12,"unit":"MONTH"}},"quoteItemRelationship":[{"type":"RELIES_ON","id":"1"}],"qualification":[{"id":"8e32872f156b4f25814720d17f2fc0c8","href":"","qualificationItem":"2"}],"relatedParty":[{"id":"","role":["buyer"],"name":"","emailAddress":"","number":"","numberExtension":""}]}]}`
	var quote sw.QuoteRes
	if err := json.Unmarshal([]byte(data), &quote); err != nil {
		t.Fatal(err)
	} else {
		t.Log(quote.ID)
	}
}
