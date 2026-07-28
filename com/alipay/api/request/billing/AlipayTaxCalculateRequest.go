package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayTaxCalculateRequest struct {
	TaxCalculationRequestId string                          `json:"taxCalculationRequestId,omitempty"`
	Currency                string                          `json:"currency,omitempty"`
	LineItems               []*model.TaxCalculationLineItem `json:"lineItems,omitempty"`
	CustomerId              string                          `json:"customerId,omitempty"`
	CustomerDetails         *model.TaxCustomerDetails       `json:"customerDetails,omitempty"`
	ShipFromDetails         *model.TaxShipFromDetails       `json:"shipFromDetails,omitempty"`
	ShippingCost            *model.TaxShippingCost          `json:"shippingCost,omitempty"`
	TaxDate                 string                          `json:"taxDate,omitempty"`
}

func NewAlipayTaxCalculateRequest() (*request.AlipayRequest, *AlipayTaxCalculateRequest) {
	alipayTaxCalculateRequest := &AlipayTaxCalculateRequest{}
	alipayRequest := request.NewAlipayRequest(alipayTaxCalculateRequest, "/ams/api/v1/tax/calculate", &responseBilling.AlipayTaxCalculateResponse{})
	return alipayRequest, alipayTaxCalculateRequest
}

func (alipayTaxCalculateRequest *AlipayTaxCalculateRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayTaxCalculateRequest, "/ams/api/v1/tax/calculate", &responseBilling.AlipayTaxCalculateResponse{})
}
