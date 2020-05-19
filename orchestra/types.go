package orchestra

import (
	"time"

	invmod "github.com/qlcchain/go-lsobus/sonata/inventory/models"
	ordmod "github.com/qlcchain/go-lsobus/sonata/order/models"
	poqmod "github.com/qlcchain/go-lsobus/sonata/poq/models"
	quomod "github.com/qlcchain/go-lsobus/sonata/quote/models"
	sitmod "github.com/qlcchain/go-lsobus/sonata/site/models"
)

const (
	BillingTypePAYG  = "PAYG"
	BillingTypeDOD   = "DOD"
	BillingTypeUsage = "USAGE"
)

type Partner struct {
	ID   string
	Name string
}

type BillingParams struct {
	BillingType   string
	BillingUnit   string    // used for PAYG, etc day/month/year
	MeasureUnit   string    // used for USAGE, etc minute/hour/Mbps/MByte
	StartTime     time.Time // used for DOD Duration
	EndTime       time.Time // used for DOD Duration
	CurrencyUnit  string    // etc USA/HKD/CNY
	CurrencyPrice float32
}

type OrderParams struct {
	OrderActivity string
	ItemAction    string
	ProductID     string

	ContractID string
	Buyer      *Partner
	Seller     *Partner

	ExternalID  string
	Description string
	ProjectID   string

	SrcSiteID    string
	SrcPortSpeed uint
	DstSiteID    string
	DstPortSpeed uint

	SrcPortID string
	SrcVlanID []uint
	DstPortID string
	DstVlanID []uint

	Bandwidth uint
	SVlanID   uint
	CosName   string

	BillingParams *BillingParams

	RspPoq   *poqmod.ProductOfferingQualification
	RspQuote *quomod.Quote
	RspOrder *ordmod.ProductOrder
}

type FindParams struct {
	Offset     string
	Limit      string
	ProjectID  string
	ExternalID string
	BuyerID    string
	SiteID     string
	State      string

	ProductSpecificationID string

	RspSiteList  []*sitmod.GeographicSiteFindResp
	RspPoqList   []*poqmod.ProductOfferingQualificationFind
	RspQuoteList []*quomod.QuoteFind
	RspOrderList []*ordmod.ProductOrderSummary
	RspInvList   []*invmod.ProductSummary
}

type GetParams struct {
	ID string

	RspSite  *sitmod.GeographicSite
	RspPoq   *poqmod.ProductOfferingQualification
	RspQuote *quomod.Quote
	RspOrder *ordmod.ProductOrder
	RspInv   *invmod.Product
}
