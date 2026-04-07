package aba

import (
responseAba  "github.com/alipay/global-open-sdk-go/com/alipay/api/response/aba" 
 "github.com/alipay/global-open-sdk-go/com/alipay/api/model"
 "github.com/alipay/global-open-sdk-go/com/alipay/api/request"

)




type AlipayUpdateCardStatusRequest struct {
        AssetId string `json:"assetId,omitempty"`
        RequestId string `json:"requestId,omitempty"`
        OperateType string `json:"operateType,omitempty"`
        NotifyUrl string `json:"notifyUrl,omitempty"`
}

func NewAlipayUpdateCardStatusRequest() (*request.AlipayRequest, *AlipayUpdateCardStatusRequest) { 
alipayUpdateCardStatusRequest := &AlipayUpdateCardStatusRequest{} 
alipayRequest := request.NewAlipayRequest (alipayUpdateCardStatusRequest,  "null", &responseAba.AlipayUpdateCardStatusResponse{}) 
return alipayRequest, alipayUpdateCardStatusRequest 
} 
 
func (alipayUpdateCardStatusRequest *AlipayUpdateCardStatusRequest) NewRequest() *request.AlipayRequest { 
return request.NewAlipayRequest(&alipayUpdateCardStatusRequest,  "null", &responseAba.AlipayUpdateCardStatusResponse{}) 
}







