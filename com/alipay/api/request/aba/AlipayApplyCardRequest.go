package aba

import (
responseAba  "github.com/alipay/global-open-sdk-go/com/alipay/api/response/aba" 
 "github.com/alipay/global-open-sdk-go/com/alipay/api/model"
 "github.com/alipay/global-open-sdk-go/com/alipay/api/request"

)




type AlipayApplyCardRequest struct {
        RequestId string `json:"requestId,omitempty"`
        CardNickName string `json:"cardNickName,omitempty"`
        Note string `json:"note,omitempty"`
        CardBinRule string `json:"cardBinRule,omitempty"`
        Purpose string `json:"purpose,omitempty"`
        Metadata map[string]string `json:"metadata,omitempty"`
        AuthorizationControl *model.AuthorizationControl `json:"authorizationControl,omitempty"` 
}

func NewAlipayApplyCardRequest() (*request.AlipayRequest, *AlipayApplyCardRequest) { 
alipayApplyCardRequest := &AlipayApplyCardRequest{} 
alipayRequest := request.NewAlipayRequest (alipayApplyCardRequest,  "null", &responseAba.AlipayApplyCardResponse{}) 
return alipayRequest, alipayApplyCardRequest 
} 
 
func (alipayApplyCardRequest *AlipayApplyCardRequest) NewRequest() *request.AlipayRequest { 
return request.NewAlipayRequest(&alipayApplyCardRequest,  "null", &responseAba.AlipayApplyCardResponse{}) 
}







