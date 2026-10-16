package agentic

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseAgentic "github.com/alipay/global-open-sdk-go/com/alipay/api/response/agentic"
)

type AlipayInquireAccountInfoRequest struct {
	RequestId   string                         `json:"requestId,omitempty"`
	UserLoginId string                         `json:"userLoginId,omitempty"`
	Agent       *model.AgenticAccountInfoAgent `json:"agent,omitempty"`
	Env         *model.AgenticAccountInfoEnv   `json:"env,omitempty"`
}

func NewAlipayInquireAccountInfoRequest() (*request.AlipayRequest, *AlipayInquireAccountInfoRequest) {
	alipayInquireAccountInfoRequest := &AlipayInquireAccountInfoRequest{}
	alipayRequest := request.NewAlipayRequest(alipayInquireAccountInfoRequest, "/ams/api/v1/agenticPayment/inquireAccountInfo", &responseAgentic.AlipayInquireAccountInfoResponse{})
	return alipayRequest, alipayInquireAccountInfoRequest
}

func (alipayInquireAccountInfoRequest *AlipayInquireAccountInfoRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayInquireAccountInfoRequest, "/ams/api/v1/agenticPayment/inquireAccountInfo", &responseAgentic.AlipayInquireAccountInfoResponse{})
}
