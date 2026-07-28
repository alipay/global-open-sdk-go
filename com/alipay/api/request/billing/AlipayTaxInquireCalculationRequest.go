package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayTaxInquireCalculationRequest struct {
	TaxCalculationId        string `json:"taxCalculationId,omitempty"`
	TaxCalculationRequestId string `json:"taxCalculationRequestId,omitempty"`
	PaymentRequestId        string `json:"paymentRequestId,omitempty"`
}

func NewAlipayTaxInquireCalculationRequest() (*request.AlipayRequest, *AlipayTaxInquireCalculationRequest) {
	alipayTaxInquireCalculationRequest := &AlipayTaxInquireCalculationRequest{}
	alipayRequest := request.NewAlipayRequest(alipayTaxInquireCalculationRequest, "/ams/api/v1/tax/inquireCalculation", &responseBilling.AlipayTaxInquireCalculationResponse{})
	return alipayRequest, alipayTaxInquireCalculationRequest
}

func (alipayTaxInquireCalculationRequest *AlipayTaxInquireCalculationRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayTaxInquireCalculationRequest, "/ams/api/v1/tax/inquireCalculation", &responseBilling.AlipayTaxInquireCalculationResponse{})
}
