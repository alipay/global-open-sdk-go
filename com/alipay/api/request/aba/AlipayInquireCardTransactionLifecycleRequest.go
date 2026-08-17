package aba

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseAba "github.com/alipay/global-open-sdk-go/com/alipay/api/response/aba"
)

type AlipayInquireCardTransactionLifecycleRequest struct {
	StartTime               string   `json:"startTime,omitempty"`
	EndTime                 string   `json:"endTime,omitempty"`
	AssetIdList             []string `json:"assetIdList,omitempty"`
	TransactionCurrencyList []string `json:"transactionCurrencyList,omitempty"`
	LifecycleIdList         []string `json:"lifecycleIdList,omitempty"`
	PageSize                int32    `json:"pageSize,omitempty"`
	PageNumber              int32    `json:"pageNumber,omitempty"`
}

func NewAlipayInquireCardTransactionLifecycleRequest() (*request.AlipayRequest, *AlipayInquireCardTransactionLifecycleRequest) {
	alipayInquireCardTransactionLifecycleRequest := &AlipayInquireCardTransactionLifecycleRequest{}
	alipayRequest := request.NewAlipayRequest(alipayInquireCardTransactionLifecycleRequest, "/ams/api/v1/aba/cards/inquireCardTransactionLifecycle", &responseAba.AlipayInquireCardTransactionLifecycleResponse{})
	return alipayRequest, alipayInquireCardTransactionLifecycleRequest
}

func (alipayInquireCardTransactionLifecycleRequest *AlipayInquireCardTransactionLifecycleRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayInquireCardTransactionLifecycleRequest, "/ams/api/v1/aba/cards/inquireCardTransactionLifecycle", &responseAba.AlipayInquireCardTransactionLifecycleResponse{})
}
