package aba

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseAba "github.com/alipay/global-open-sdk-go/com/alipay/api/response/aba"
)

type AlipayInquiryStatementListRequest struct {
	FuzzyName           string   `json:"fuzzyName,omitempty"`
	CustomerId          string   `json:"customerId,omitempty"`
	AccessToken         string   `json:"accessToken,omitempty"`
	StartTime           string   `json:"startTime,omitempty"`
	EndTime             string   `json:"endTime,omitempty"`
	TransactionTypeList []string `json:"transactionTypeList,omitempty"`
	CurrencyList        []string `json:"currencyList,omitempty"`
	PageSize            string   `json:"pageSize,omitempty"`
	PageNumber          string   `json:"pageNumber,omitempty"`
}

func NewAlipayInquiryStatementListRequest() (*request.AlipayRequest, *AlipayInquiryStatementListRequest) {
	alipayInquiryStatementListRequest := &AlipayInquiryStatementListRequest{}
	alipayRequest := request.NewAlipayRequest(alipayInquiryStatementListRequest, "/ams/api/v1/aba/accounts/inquiryStatementList", &responseAba.AlipayInquiryStatementListResponse{})
	return alipayRequest, alipayInquiryStatementListRequest
}

func (alipayInquiryStatementListRequest *AlipayInquiryStatementListRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayInquiryStatementListRequest, "/ams/api/v1/aba/accounts/inquiryStatementList", &responseAba.AlipayInquiryStatementListResponse{})
}
