package aba

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseAba "github.com/alipay/global-open-sdk-go/com/alipay/api/response/aba"
)

type AlipayInquireCardSensitiveInfoRequest struct {
	AssetId string `json:"assetId,omitempty"`
}

func NewAlipayInquireCardSensitiveInfoRequest() (*request.AlipayRequest, *AlipayInquireCardSensitiveInfoRequest) {
	alipayInquireCardSensitiveInfoRequest := &AlipayInquireCardSensitiveInfoRequest{}
	alipayRequest := request.NewAlipayRequest(alipayInquireCardSensitiveInfoRequest, "", &responseAba.AlipayInquireCardSensitiveInfoResponse{})
	return alipayRequest, alipayInquireCardSensitiveInfoRequest
}

func (alipayInquireCardSensitiveInfoRequest *AlipayInquireCardSensitiveInfoRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayInquireCardSensitiveInfoRequest, "", &responseAba.AlipayInquireCardSensitiveInfoResponse{})
}
