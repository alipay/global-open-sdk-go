package aba

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseAba "github.com/alipay/global-open-sdk-go/com/alipay/api/response/aba"
)

type AlipayInquireCardDetailRequest struct {
	AssetId string `json:"assetId,omitempty"`
}

func NewAlipayInquireCardDetailRequest() (*request.AlipayRequest, *AlipayInquireCardDetailRequest) {
	alipayInquireCardDetailRequest := &AlipayInquireCardDetailRequest{}
	alipayRequest := request.NewAlipayRequest(alipayInquireCardDetailRequest, "/ams/api/v1/aba/cards/inquireCardDetail", &responseAba.AlipayInquireCardDetailResponse{})
	return alipayRequest, alipayInquireCardDetailRequest
}

func (alipayInquireCardDetailRequest *AlipayInquireCardDetailRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayInquireCardDetailRequest, "/ams/api/v1/aba/cards/inquireCardDetail", &responseAba.AlipayInquireCardDetailResponse{})
}
