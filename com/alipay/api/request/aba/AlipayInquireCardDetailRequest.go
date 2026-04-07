package aba

import (
responseAba  "github.com/alipay/global-open-sdk-go/com/alipay/api/response/aba" 
 "github.com/alipay/global-open-sdk-go/com/alipay/api/model"
 "github.com/alipay/global-open-sdk-go/com/alipay/api/request"

)




type AlipayInquireCardDetailRequest struct {
        AssetId string `json:"assetId,omitempty"`
}

func NewAlipayInquireCardDetailRequest() (*request.AlipayRequest, *AlipayInquireCardDetailRequest) { 
alipayInquireCardDetailRequest := &AlipayInquireCardDetailRequest{} 
alipayRequest := request.NewAlipayRequest (alipayInquireCardDetailRequest,  "null", &responseAba.AlipayInquireCardDetailResponse{}) 
return alipayRequest, alipayInquireCardDetailRequest 
} 
 
func (alipayInquireCardDetailRequest *AlipayInquireCardDetailRequest) NewRequest() *request.AlipayRequest { 
return request.NewAlipayRequest(&alipayInquireCardDetailRequest,  "null", &responseAba.AlipayInquireCardDetailResponse{}) 
}







