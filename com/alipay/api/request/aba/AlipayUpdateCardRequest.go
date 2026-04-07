package aba

import (
responseAba  "github.com/alipay/global-open-sdk-go/com/alipay/api/response/aba" 
 "github.com/alipay/global-open-sdk-go/com/alipay/api/model"
 "github.com/alipay/global-open-sdk-go/com/alipay/api/request"

)




type AlipayUpdateCardRequest struct {
        AssetId string `json:"assetId,omitempty"`
        RequestId string `json:"requestId,omitempty"`
        CardNickName string `json:"cardNickName,omitempty"`
        Note string `json:"note,omitempty"`
        Purpose string `json:"purpose,omitempty"`
        Metadata map[string]string `json:"metadata,omitempty"`
        AuthorizationControl *model.AuthorizationControl `json:"authorizationControl,omitempty"` 
}

func NewAlipayUpdateCardRequest() (*request.AlipayRequest, *AlipayUpdateCardRequest) { 
alipayUpdateCardRequest := &AlipayUpdateCardRequest{} 
alipayRequest := request.NewAlipayRequest (alipayUpdateCardRequest,  "null", &responseAba.AlipayUpdateCardResponse{}) 
return alipayRequest, alipayUpdateCardRequest 
} 
 
func (alipayUpdateCardRequest *AlipayUpdateCardRequest) NewRequest() *request.AlipayRequest { 
return request.NewAlipayRequest(&alipayUpdateCardRequest,  "null", &responseAba.AlipayUpdateCardResponse{}) 
}







