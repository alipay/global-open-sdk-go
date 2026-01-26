package aba

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseAba "github.com/alipay/global-open-sdk-go/com/alipay/api/response/aba"
)

type AlipayInquireAvailableQuotaRequest struct {
	Currency string `json:"currency,omitempty"`
}

func NewAlipayInquireAvailableQuotaRequest() (*request.AlipayRequest, *AlipayInquireAvailableQuotaRequest) {
	alipayInquireAvailableQuotaRequest := &AlipayInquireAvailableQuotaRequest{}
	alipayRequest := request.NewAlipayRequest(alipayInquireAvailableQuotaRequest, "/ams/v1/aba/account/inquireAvailableQuota", &responseAba.AlipayInquireAvailableQuotaResponse{})
	return alipayRequest, alipayInquireAvailableQuotaRequest
}

func (alipayInquireAvailableQuotaRequest *AlipayInquireAvailableQuotaRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayInquireAvailableQuotaRequest, "/ams/v1/aba/account/inquireAvailableQuota", &responseAba.AlipayInquireAvailableQuotaResponse{})
}
