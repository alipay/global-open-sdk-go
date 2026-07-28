package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayTaxCalculateResponse struct {
	response.AlipayResponse
	Result             *model.Result                    `json:"result,omitempty"`
	TaxCalculationId   string                           `json:"taxCalculationId,omitempty"`
	Currency           string                           `json:"currency,omitempty"`
	TotalAmount        string                           `json:"totalAmount,omitempty"`
	ExclusiveTaxAmount string                           `json:"exclusiveTaxAmount,omitempty"`
	InclusiveTaxAmount string                           `json:"inclusiveTaxAmount,omitempty"`
	LineItems          []*model.TaxCalculatedLineItem   `json:"lineItems,omitempty"`
	TaxBreakdown       []*model.TaxBreakdown            `json:"taxBreakdown,omitempty"`
	ExpireAt           string                           `json:"expireAt,omitempty"`
	TaxDate            string                           `json:"taxDate,omitempty"`
	ShippingCost       *model.TaxCalculatedShippingCost `json:"shippingCost,omitempty"`
}
