package aba

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseAba "github.com/alipay/global-open-sdk-go/com/alipay/api/response/aba"
)

type AlipayInquiryStatementListRequest struct {
	CustomerId          string   `json:"customerId,customerId"`
	AccessToken         string   `json:"accessToken,omitempty"`
	StartTime           string   `json:"startTime,omitempty"`
	EndTime             string   `json:"endTime,omitempty"`
	TransactionTypeList []string `json:"transactionTypeList,omitempty"`
	CurrencyList        []string `json:"currencyList,omitempty"`
	PageSize            string   `json:"pageSize,omitempty"`
	PageNumber          string   `json:"pageNumber,omitempty"`
}

func (alipayInquiryStatementListRequest *AlipayInquiryStatementListRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayInquiryStatementListRequest, model.ABA_INQUERY_STATEMENT_LIST_PATH, &responseAba.AlipayInquiryStatementListResponse{})
}

func NewAlipayPayRequest() (*request.AlipayRequest, *AlipayInquiryStatementListRequest) {
	alipayInquiryStatementListRequest := &AlipayInquiryStatementListRequest{}
	alipayRequest := request.NewAlipayRequest(alipayInquiryStatementListRequest, model.ABA_INQUERY_STATEMENT_LIST_PATH, &responseAba.AlipayInquiryStatementListResponse{})
	return alipayRequest, alipayInquiryStatementListRequest
}
