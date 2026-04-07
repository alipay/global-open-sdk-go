package auth

import (
responseAuth  "github.com/alipay/global-open-sdk-go/com/alipay/api/response/auth" 
 "github.com/alipay/global-open-sdk-go/com/alipay/api/model"
 "github.com/alipay/global-open-sdk-go/com/alipay/api/request"

)




type AlipayAuthRevokeTokenRequest struct {
        AccessToken string `json:"accessToken,omitempty"`
        ExtendInfo string `json:"extendInfo,omitempty"`
        MerchantAccountId string `json:"merchantAccountId,omitempty"`
}

func NewAlipayAuthRevokeTokenRequest() (*request.AlipayRequest, *AlipayAuthRevokeTokenRequest) { 
alipayAuthRevokeTokenRequest := &AlipayAuthRevokeTokenRequest{} 
alipayRequest := request.NewAlipayRequest (alipayAuthRevokeTokenRequest,  "null", &responseAuth.AlipayAuthRevokeTokenResponse{}) 
return alipayRequest, alipayAuthRevokeTokenRequest 
} 
 
func (alipayAuthRevokeTokenRequest *AlipayAuthRevokeTokenRequest) NewRequest() *request.AlipayRequest { 
return request.NewAlipayRequest(&alipayAuthRevokeTokenRequest,  "null", &responseAuth.AlipayAuthRevokeTokenResponse{}) 
}







