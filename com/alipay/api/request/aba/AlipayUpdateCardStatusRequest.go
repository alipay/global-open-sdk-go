package aba

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseAba "github.com/alipay/global-open-sdk-go/com/alipay/api/response/aba"
)

type AlipayUpdateCardStatusRequest struct {
	AssetId     string `json:"assetId,omitempty"`
	RequestId   string `json:"requestId,omitempty"`
	OperateType string `json:"operateType,omitempty"`
	NotifyUrl   string `json:"notifyUrl,omitempty"`
}

func NewAlipayUpdateCardStatusRequest() (*request.AlipayRequest, *AlipayUpdateCardStatusRequest) {
	alipayUpdateCardStatusRequest := &AlipayUpdateCardStatusRequest{}
	alipayRequest := request.NewAlipayRequest(alipayUpdateCardStatusRequest, "/ams/api/v1/aba/cards/updateCardStatus", &responseAba.AlipayUpdateCardStatusResponse{})
	return alipayRequest, alipayUpdateCardStatusRequest
}

func (alipayUpdateCardStatusRequest *AlipayUpdateCardStatusRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayUpdateCardStatusRequest, "/ams/api/v1/aba/cards/updateCardStatus", &responseAba.AlipayUpdateCardStatusResponse{})
}
