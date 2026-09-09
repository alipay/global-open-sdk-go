package pay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responsePay "github.com/alipay/global-open-sdk-go/com/alipay/api/response/pay"
)

type AlipayInquireHoldListRequest struct {
	HoldRequestId string `json:"holdRequestId,omitempty"`
	HoldId        string `json:"holdId,omitempty"`
	Limit         int32  `json:"limit,omitempty"`
	StartingAfter string `json:"startingAfter,omitempty"`
	EndingBefore  string `json:"endingBefore,omitempty"`
}

func NewAlipayInquireHoldListRequest() (*request.AlipayRequest, *AlipayInquireHoldListRequest) {
	alipayInquireHoldListRequest := &AlipayInquireHoldListRequest{}
	alipayRequest := request.NewAlipayRequest(alipayInquireHoldListRequest, "/ams/api/v1/payments/reserve/inquireHoldList", &responsePay.AlipayInquireHoldListResponse{})
	return alipayRequest, alipayInquireHoldListRequest
}

func (alipayInquireHoldListRequest *AlipayInquireHoldListRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayInquireHoldListRequest, "/ams/api/v1/payments/reserve/inquireHoldList", &responsePay.AlipayInquireHoldListResponse{})
}
