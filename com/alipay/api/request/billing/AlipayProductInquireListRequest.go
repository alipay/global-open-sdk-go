package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayProductInquireListRequest struct {
	StartingAfter string `json:"startingAfter,omitempty"`
	EndingBefore  string `json:"endingBefore,omitempty"`
	Limit         int32  `json:"limit,omitempty"`
	Active        bool   `json:"active,omitempty"`
	Type          string `json:"type,omitempty"`
	Keyword       string `json:"keyword,omitempty"`
	IncludeTotal  bool   `json:"includeTotal,omitempty"`
}

func NewAlipayProductInquireListRequest() (*request.AlipayRequest, *AlipayProductInquireListRequest) {
	alipayProductInquireListRequest := &AlipayProductInquireListRequest{}
	alipayRequest := request.NewAlipayRequest(alipayProductInquireListRequest, "/ams/api/v1/billing/product/inquireList", &responseBilling.AlipayProductInquireListResponse{})
	return alipayRequest, alipayProductInquireListRequest
}

func (alipayProductInquireListRequest *AlipayProductInquireListRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayProductInquireListRequest, "/ams/api/v1/billing/product/inquireList", &responseBilling.AlipayProductInquireListResponse{})
}
