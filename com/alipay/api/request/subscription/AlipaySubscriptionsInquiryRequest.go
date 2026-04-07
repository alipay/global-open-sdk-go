package subscription

import (
responseSubscription  "github.com/alipay/global-open-sdk-go/com/alipay/api/response/subscription" 
 "github.com/alipay/global-open-sdk-go/com/alipay/api/model"
 "github.com/alipay/global-open-sdk-go/com/alipay/api/request"

)




type AlipaySubscriptionsInquiryRequest struct {
        SubscriptionId string `json:"subscriptionId,omitempty"`
        SubscriptionRequestId string `json:"subscriptionRequestId,omitempty"`
}

func NewAlipaySubscriptionsInquiryRequest() (*request.AlipayRequest, *AlipaySubscriptionsInquiryRequest) { 
alipaySubscriptionsInquiryRequest := &AlipaySubscriptionsInquiryRequest{} 
alipayRequest := request.NewAlipayRequest (alipaySubscriptionsInquiryRequest,  "null", &responseSubscription.AlipaySubscriptionsInquiryResponse{}) 
return alipayRequest, alipaySubscriptionsInquiryRequest 
} 
 
func (alipaySubscriptionsInquiryRequest *AlipaySubscriptionsInquiryRequest) NewRequest() *request.AlipayRequest { 
return request.NewAlipayRequest(&alipaySubscriptionsInquiryRequest,  "null", &responseSubscription.AlipaySubscriptionsInquiryResponse{}) 
}







