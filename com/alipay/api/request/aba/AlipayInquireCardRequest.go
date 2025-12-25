package aba

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseAba "github.com/alipay/global-open-sdk-go/com/alipay/api/response/aba"
)

type AlipayInquireCardRequest struct {
	PageNumber     int32  `json:"pageNumber,omitempty"`
	PageSize       int32  `json:"pageSize,omitempty"`
	LastFourDigits string `json:"lastFourDigits,omitempty"`
	CardStatus     string `json:"cardStatus,omitempty"`
	CardNickName   string `json:"cardNickName,omitempty"`
}

func NewAlipayInquireCardRequest() (*request.AlipayRequest, *AlipayInquireCardRequest) {
	alipayInquireCardRequest := &AlipayInquireCardRequest{}
	alipayRequest := request.NewAlipayRequest(alipayInquireCardRequest, "/ams/api/v1/aba/cards/inquireCard", &responseAba.AlipayInquireCardResponse{})
	return alipayRequest, alipayInquireCardRequest
}

func (alipayInquireCardRequest *AlipayInquireCardRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayInquireCardRequest, "/ams/api/v1/aba/cards/inquireCard", &responseAba.AlipayInquireCardResponse{})
}
