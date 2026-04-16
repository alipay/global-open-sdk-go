package aba

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseAba "github.com/alipay/global-open-sdk-go/com/alipay/api/response/aba"
)

type AlipayInquiryStatementDetailRequest struct {
	StatementId string `json:"statementId,omitempty"`
}

func NewAlipayInquiryStatementDetailRequest() (*request.AlipayRequest, *AlipayInquiryStatementDetailRequest) {
	alipayInquiryStatementDetailRequest := &AlipayInquiryStatementDetailRequest{}
	alipayRequest := request.NewAlipayRequest(alipayInquiryStatementDetailRequest, "/ams/api/v1/aba/accounts/inquiryStatementDetail", &responseAba.AlipayInquiryStatementDetailResponse{})
	return alipayRequest, alipayInquiryStatementDetailRequest
}

func (alipayInquiryStatementDetailRequest *AlipayInquiryStatementDetailRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayInquiryStatementDetailRequest, "/ams/api/v1/aba/accounts/inquiryStatementDetail", &responseAba.AlipayInquiryStatementDetailResponse{})
}
