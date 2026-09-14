package pay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responsePay "github.com/alipay/global-open-sdk-go/com/alipay/api/response/pay"
)

type AlipayInquireRuleListRequest struct {
	FundsType     model.FundsType   `json:"fundsType,omitempty"`
	TakeType      model.TakeType    `json:"takeType,omitempty"`
	ReleaseType   model.ReleaseType `json:"releaseType,omitempty"`
	Limit         int32             `json:"limit,omitempty"`
	StartingAfter string            `json:"startingAfter,omitempty"`
	EndingBefore  string            `json:"endingBefore,omitempty"`
}

func NewAlipayInquireRuleListRequest() (*request.AlipayRequest, *AlipayInquireRuleListRequest) {
	alipayInquireRuleListRequest := &AlipayInquireRuleListRequest{}
	alipayRequest := request.NewAlipayRequest(alipayInquireRuleListRequest, "/ams/api/v1/payments/reserve/inquireRuleList", &responsePay.AlipayInquireRuleListResponse{})
	return alipayRequest, alipayInquireRuleListRequest
}

func (alipayInquireRuleListRequest *AlipayInquireRuleListRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayInquireRuleListRequest, "/ams/api/v1/payments/reserve/inquireRuleList", &responsePay.AlipayInquireRuleListResponse{})
}
