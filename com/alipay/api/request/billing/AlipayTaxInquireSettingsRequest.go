package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayTaxInquireSettingsRequest struct {
}

func NewAlipayTaxInquireSettingsRequest() (*request.AlipayRequest, *AlipayTaxInquireSettingsRequest) {
	alipayTaxInquireSettingsRequest := &AlipayTaxInquireSettingsRequest{}
	alipayRequest := request.NewAlipayRequest(alipayTaxInquireSettingsRequest, "/ams/api/v1/tax/inquireSettings", &responseBilling.AlipayTaxInquireSettingsResponse{})
	return alipayRequest, alipayTaxInquireSettingsRequest
}

func (alipayTaxInquireSettingsRequest *AlipayTaxInquireSettingsRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayTaxInquireSettingsRequest, "/ams/api/v1/tax/inquireSettings", &responseBilling.AlipayTaxInquireSettingsResponse{})
}
