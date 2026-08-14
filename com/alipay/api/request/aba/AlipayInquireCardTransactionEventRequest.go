package aba

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseAba "github.com/alipay/global-open-sdk-go/com/alipay/api/response/aba"
)

type AlipayInquireCardTransactionEventRequest struct {
	StartTime               string                                  `json:"startTime,omitempty"`
	EndTime                 string                                  `json:"endTime,omitempty"`
	AssetIdList             []string                                `json:"assetIdList,omitempty"`
	EventIdList             []string                                `json:"eventIdList,omitempty"`
	LifecycleIdList         []string                                `json:"lifecycleIdList,omitempty"`
	TransactionCurrencyList []string                                `json:"transactionCurrencyList,omitempty"`
	EventTypeList           []model.CardTransactionEventFilterType  `json:"eventTypeList,omitempty"`
	StatusList              []model.CardTransactionStatusFilterType `json:"statusList,omitempty"`
	PageSize                int32                                   `json:"pageSize,omitempty"`
	PageNumber              int32                                   `json:"pageNumber,omitempty"`
}

func NewAlipayInquireCardTransactionEventRequest() (*request.AlipayRequest, *AlipayInquireCardTransactionEventRequest) {
	alipayInquireCardTransactionEventRequest := &AlipayInquireCardTransactionEventRequest{}
	alipayRequest := request.NewAlipayRequest(alipayInquireCardTransactionEventRequest, "/ams/api/v1/aba/cards/inquireCardTransactionEvent", &responseAba.AlipayInquireCardTransactionEventResponse{})
	return alipayRequest, alipayInquireCardTransactionEventRequest
}

func (alipayInquireCardTransactionEventRequest *AlipayInquireCardTransactionEventRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayInquireCardTransactionEventRequest, "/ams/api/v1/aba/cards/inquireCardTransactionEvent", &responseAba.AlipayInquireCardTransactionEventResponse{})
}
