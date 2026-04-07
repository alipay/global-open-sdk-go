package pay

import (
responsePay  "github.com/alipay/global-open-sdk-go/com/alipay/api/response/pay" 
 "github.com/alipay/global-open-sdk-go/com/alipay/api/model"
 "github.com/alipay/global-open-sdk-go/com/alipay/api/request"

)




type AlipayVaultsFetchNonceRequest struct {
        Card *model.Card `json:"card,omitempty"` 
}

func NewAlipayVaultsFetchNonceRequest() (*request.AlipayRequest, *AlipayVaultsFetchNonceRequest) { 
alipayVaultsFetchNonceRequest := &AlipayVaultsFetchNonceRequest{} 
alipayRequest := request.NewAlipayRequest (alipayVaultsFetchNonceRequest,  "null", &responsePay.AlipayVaultsFetchNonceResponse{}) 
return alipayRequest, alipayVaultsFetchNonceRequest 
} 
 
func (alipayVaultsFetchNonceRequest *AlipayVaultsFetchNonceRequest) NewRequest() *request.AlipayRequest { 
return request.NewAlipayRequest(&alipayVaultsFetchNonceRequest,  "null", &responsePay.AlipayVaultsFetchNonceResponse{}) 
}







