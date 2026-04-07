package subscription

import (
responseSubscription  "github.com/alipay/global-open-sdk-go/com/alipay/api/response/subscription" 
 "github.com/alipay/global-open-sdk-go/com/alipay/api/model"
 "github.com/alipay/global-open-sdk-go/com/alipay/api/request"

)




type AlipaySubscriptionCancelRequest struct {
        SubscriptionId string `json:"subscriptionId,omitempty"`
        SubscriptionRequestId string `json:"subscriptionRequestId,omitempty"`
        CancellationType model.CancellationType `json:"cancellationType,omitempty"`
}

func NewAlipaySubscriptionCancelRequest() (*request.AlipayRequest, *AlipaySubscriptionCancelRequest) { 
alipaySubscriptionCancelRequest := &AlipaySubscriptionCancelRequest{} 
alipayRequest := request.NewAlipayRequest (alipaySubscriptionCancelRequest,  "null", &responseSubscription.AlipaySubscriptionCancelResponse{}) 
return alipayRequest, alipaySubscriptionCancelRequest 
} 
 
func (alipaySubscriptionCancelRequest *AlipaySubscriptionCancelRequest) NewRequest() *request.AlipayRequest { 
return request.NewAlipayRequest(&alipaySubscriptionCancelRequest,  "null", &responseSubscription.AlipaySubscriptionCancelResponse{}) 
}







