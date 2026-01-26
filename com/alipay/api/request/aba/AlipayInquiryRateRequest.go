package aba

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseAba "github.com/alipay/global-open-sdk-go/com/alipay/api/response/aba"
)

type AlipayInquiryRateRequest struct {
	RateConditionList []*model.InquiryRateCondition `json:"rateConditionList,omitempty"`
}

func NewAlipayInquiryRateRequest() (*request.AlipayRequest, *AlipayInquiryRateRequest) {
	alipayInquiryRateRequest := &AlipayInquiryRateRequest{}
	alipayRequest := request.NewAlipayRequest(alipayInquiryRateRequest, "/ams/v1/aba/funds/inquireRate", &responseAba.AlipayInquiryRateResponse{})
	return alipayRequest, alipayInquiryRateRequest
}

func (alipayInquiryRateRequest *AlipayInquiryRateRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayInquiryRateRequest, "/ams/v1/aba/funds/inquireRate", &responseAba.AlipayInquiryRateResponse{})
}
