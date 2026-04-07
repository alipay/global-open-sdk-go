package subscription

import (
responseSubscription  "github.com/alipay/global-open-sdk-go/com/alipay/api/response/subscription" 
 "github.com/alipay/global-open-sdk-go/com/alipay/api/model"
 "github.com/alipay/global-open-sdk-go/com/alipay/api/request"

)




type AlipayInquireSubscriptionRequest struct {
        MerchantAccountId string `json:"merchantAccountId,omitempty"`
        SubscriptionId string `json:"subscriptionId,omitempty"`
}

func NewAlipayInquireSubscriptionRequest() (*request.AlipayRequest, *AlipayInquireSubscriptionRequest) { 
alipayInquireSubscriptionRequest := &AlipayInquireSubscriptionRequest{} 
alipayRequest := request.NewAlipayRequest (alipayInquireSubscriptionRequest,  "null", &responseSubscription.AlipayInquireSubscriptionResponse{}) 
return alipayRequest, alipayInquireSubscriptionRequest 
} 
 
func (alipayInquireSubscriptionRequest *AlipayInquireSubscriptionRequest) NewRequest() *request.AlipayRequest { 
return request.NewAlipayRequest(&alipayInquireSubscriptionRequest,  "null", &responseSubscription.AlipayInquireSubscriptionResponse{}) 
}







