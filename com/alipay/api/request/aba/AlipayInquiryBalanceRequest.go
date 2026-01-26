package aba

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseAba "github.com/alipay/global-open-sdk-go/com/alipay/api/response/aba"
)

type AlipayInquiryBalanceRequest struct {
	CurrencyList []string `json:"currencyList,omitempty"`
	Accesstoken  string   `json:"accesstoken,omitempty"`
	CustomerId   string   `json:"customerId,omitempty"`
}

func NewAlipayInquiryBalanceRequest() (*request.AlipayRequest, *AlipayInquiryBalanceRequest) {
	alipayInquiryBalanceRequest := &AlipayInquiryBalanceRequest{}
	alipayRequest := request.NewAlipayRequest(alipayInquiryBalanceRequest, "/ams/v1/aba/accounts/inquiryBalance", &responseAba.AlipayInquiryBalanceResponse{})
	return alipayRequest, alipayInquiryBalanceRequest
}

func (alipayInquiryBalanceRequest *AlipayInquiryBalanceRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayInquiryBalanceRequest, "/ams/v1/aba/accounts/inquiryBalance", &responseAba.AlipayInquiryBalanceResponse{})
}
