package subscription

import (
responseSubscription  "github.com/alipay/global-open-sdk-go/com/alipay/api/response/subscription" 
 "github.com/alipay/global-open-sdk-go/com/alipay/api/model"
 "github.com/alipay/global-open-sdk-go/com/alipay/api/request"

)




type AlipayInquireSubscriptionListRequest struct {
        MerchantAccountId string `json:"merchantAccountId,omitempty"`
        StartTimeFrom *model.time.Time `json:"startTimeFrom,omitempty"` 
        StartTimeTo *model.time.Time `json:"startTimeTo,omitempty"` 
        Statuses[] *model.SubscriptionStatus `json:"statuses,omitempty"` 
        PaymentMethodTypes []string `json:"paymentMethodTypes,omitempty"`
        Currencies []string `json:"currencies,omitempty"`
        PeriodTypes [] model.PeriodType `json:"periodTypes,omitempty"` 
        CurrentPage int32 `json:"currentPage,omitempty"`
        PageSize int32 `json:"pageSize,omitempty"`
}

func NewAlipayInquireSubscriptionListRequest() (*request.AlipayRequest, *AlipayInquireSubscriptionListRequest) { 
alipayInquireSubscriptionListRequest := &AlipayInquireSubscriptionListRequest{} 
alipayRequest := request.NewAlipayRequest (alipayInquireSubscriptionListRequest,  "null", &responseSubscription.AlipayInquireSubscriptionListResponse{}) 
return alipayRequest, alipayInquireSubscriptionListRequest 
} 
 
func (alipayInquireSubscriptionListRequest *AlipayInquireSubscriptionListRequest) NewRequest() *request.AlipayRequest { 
return request.NewAlipayRequest(&alipayInquireSubscriptionListRequest,  "null", &responseSubscription.AlipayInquireSubscriptionListResponse{}) 
}







