package aba

import (
responseAba  "github.com/alipay/global-open-sdk-go/com/alipay/api/response/aba" 
 "github.com/alipay/global-open-sdk-go/com/alipay/api/model"
 "github.com/alipay/global-open-sdk-go/com/alipay/api/request"

)




type AlipayInquireCardSensitiveInfoRequest struct {
        AssetId string `json:"assetId,omitempty"`
}

func NewAlipayInquireCardSensitiveInfoRequest() (*request.AlipayRequest, *AlipayInquireCardSensitiveInfoRequest) { 
alipayInquireCardSensitiveInfoRequest := &AlipayInquireCardSensitiveInfoRequest{} 
alipayRequest := request.NewAlipayRequest (alipayInquireCardSensitiveInfoRequest,  "null", &responseAba.AlipayInquireCardSensitiveInfoResponse{}) 
return alipayRequest, alipayInquireCardSensitiveInfoRequest 
} 
 
func (alipayInquireCardSensitiveInfoRequest *AlipayInquireCardSensitiveInfoRequest) NewRequest() *request.AlipayRequest { 
return request.NewAlipayRequest(&alipayInquireCardSensitiveInfoRequest,  "null", &responseAba.AlipayInquireCardSensitiveInfoResponse{}) 
}







