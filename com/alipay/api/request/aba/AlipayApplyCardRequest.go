package aba

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseAba "github.com/alipay/global-open-sdk-go/com/alipay/api/response/aba"
)

type AlipayApplyCardRequest struct {
	RequestId            string                      `json:"requestId,omitempty"`
	CardNickName         string                      `json:"cardNickName,omitempty"`
	Note                 string                      `json:"note,omitempty"`
	CardBinRule          string                      `json:"cardBinRule,omitempty"`
	Purpose              string                      `json:"purpose,omitempty"`
	Metadata             map[string]string           `json:"metadata,omitempty"`
	AuthorizationControl *model.AuthorizationControl `json:"authorizationControl,omitempty"`
}

func NewAlipayApplyCardRequest() (*request.AlipayRequest, *AlipayApplyCardRequest) {
	alipayApplyCardRequest := &AlipayApplyCardRequest{}
	alipayRequest := request.NewAlipayRequest(alipayApplyCardRequest, "/ams/api/v1/aba/cards/applyCard", &responseAba.AlipayApplyCardResponse{})
	return alipayRequest, alipayApplyCardRequest
}

func (alipayApplyCardRequest *AlipayApplyCardRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayApplyCardRequest, "/ams/api/v1/aba/cards/applyCard", &responseAba.AlipayApplyCardResponse{})
}
