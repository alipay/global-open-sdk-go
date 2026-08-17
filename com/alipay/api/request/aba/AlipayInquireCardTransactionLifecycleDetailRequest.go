package aba

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseAba "github.com/alipay/global-open-sdk-go/com/alipay/api/response/aba"
)

type AlipayInquireCardTransactionLifecycleDetailRequest struct {
	LifecycleId string `json:"lifecycleId,omitempty"`
}

func NewAlipayInquireCardTransactionLifecycleDetailRequest() (*request.AlipayRequest, *AlipayInquireCardTransactionLifecycleDetailRequest) {
	alipayInquireCardTransactionLifecycleDetailRequest := &AlipayInquireCardTransactionLifecycleDetailRequest{}
	alipayRequest := request.NewAlipayRequest(alipayInquireCardTransactionLifecycleDetailRequest, "/ams/api/v1/aba/cards/inquireCardTransactionLifecycleDetail", &responseAba.AlipayInquireCardTransactionLifecycleDetailResponse{})
	return alipayRequest, alipayInquireCardTransactionLifecycleDetailRequest
}

func (alipayInquireCardTransactionLifecycleDetailRequest *AlipayInquireCardTransactionLifecycleDetailRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayInquireCardTransactionLifecycleDetailRequest, "/ams/api/v1/aba/cards/inquireCardTransactionLifecycleDetail", &responseAba.AlipayInquireCardTransactionLifecycleDetailResponse{})
}
