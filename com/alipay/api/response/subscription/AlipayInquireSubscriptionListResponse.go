package responseSubscription

import (
"github.com/alipay/global-open-sdk-go/com/alipay/api/model" 
"github.com/alipay/global-open-sdk-go/com/alipay/api/response" 

)




type AlipayInquireSubscriptionListResponse struct {
response.AlipayResponse 
        Result *model.Result `json:"result,omitempty"` 
        Subscriptions[] *model.SubscriptionInfo `json:"subscriptions,omitempty"` 
        Paginator *model.Paginator `json:"paginator,omitempty"` 
}








