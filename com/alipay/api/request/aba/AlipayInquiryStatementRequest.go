package aba

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseAba "github.com/alipay/global-open-sdk-go/com/alipay/api/response/aba"
)

type AlipayInquiryStatementRequest struct {
	CustomerId          string      `json:"customerId,omitempty"`
	AccessToken         string      `json:"accessToken,omitempty"`
	StartTime           interface{} `json:"startTime,omitempty"`
	EndTime             string      `json:"endTime,omitempty"`
	TransactionTypeList []string    `json:"transactionTypeList,omitempty"`
	CurrencyList        []string    `json:"currencyList,omitempty"`
	PageSize            int32       `json:"pageSize,omitempty"`
	PageNumber          int32       `json:"pageNumber,omitempty"`
}

func NewAlipayInquiryStatementRequest() (*request.AlipayRequest, *AlipayInquiryStatementRequest) {
	alipayInquiryStatementRequest := &AlipayInquiryStatementRequest{}
	alipayRequest := request.NewAlipayRequest(alipayInquiryStatementRequest, "/ams/api/v1/aba/accounts/inquiryStatement", &responseAba.AlipayInquiryStatementResponse{})
	return alipayRequest, alipayInquiryStatementRequest
}

func (alipayInquiryStatementRequest *AlipayInquiryStatementRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayInquiryStatementRequest, "/ams/api/v1/aba/accounts/inquiryStatement", &responseAba.AlipayInquiryStatementResponse{})
}
