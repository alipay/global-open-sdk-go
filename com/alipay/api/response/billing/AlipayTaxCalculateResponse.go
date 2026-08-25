package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayTaxCalculateResponse struct {
	response.AlipayResponse
	Result             *model.Result                       `json:"result,omitempty"`
	TaxCalculationId   string                              `json:"taxCalculationId,omitempty"`
	TotalAmount        *model.Amount                       `json:"totalAmount,omitempty"`
	ExclusiveTaxAmount *model.Amount                       `json:"exclusiveTaxAmount,omitempty"`
	InclusiveTaxAmount *model.Amount                       `json:"inclusiveTaxAmount,omitempty"`
	LineItems          []*model.TaxCalculatedLineItem      `json:"lineItems,omitempty"`
	TaxBreakdown       []*model.TaxBreakdown               `json:"taxBreakdown,omitempty"`
	ExpireAt           string                              `json:"expireAt,omitempty"`
	TaxDate            string                              `json:"taxDate,omitempty"`
	ShippingCost       *model.TaxCalculatedShippingCost    `json:"shippingCost,omitempty"`
	CustomerDetails    *model.TaxCalculatedCustomerDetails `json:"customerDetails,omitempty"`
}
