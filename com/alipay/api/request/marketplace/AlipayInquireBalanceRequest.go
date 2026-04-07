package marketplace

import (
responseMarketplace  "github.com/alipay/global-open-sdk-go/com/alipay/api/response/marketplace" 
 "github.com/alipay/global-open-sdk-go/com/alipay/api/model"
 "github.com/alipay/global-open-sdk-go/com/alipay/api/request"

)




type AlipayInquireBalanceRequest struct {
        ReferenceMerchantId string `json:"referenceMerchantId,omitempty"`
}

func NewAlipayInquireBalanceRequest() (*request.AlipayRequest, *AlipayInquireBalanceRequest) { 
alipayInquireBalanceRequest := &AlipayInquireBalanceRequest{} 
alipayRequest := request.NewAlipayRequest (alipayInquireBalanceRequest,  "null", &responseMarketplace.AlipayInquireBalanceResponse{}) 
return alipayRequest, alipayInquireBalanceRequest 
} 
 
func (alipayInquireBalanceRequest *AlipayInquireBalanceRequest) NewRequest() *request.AlipayRequest { 
return request.NewAlipayRequest(&alipayInquireBalanceRequest,  "null", &responseMarketplace.AlipayInquireBalanceResponse{}) 
}







