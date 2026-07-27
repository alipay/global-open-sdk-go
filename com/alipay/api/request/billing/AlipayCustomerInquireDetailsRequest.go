package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayCustomerInquireDetailsRequest struct {
	CustomerId string `json:"customerId,omitempty"`
}

func NewAlipayCustomerInquireDetailsRequest() (*request.AlipayRequest, *AlipayCustomerInquireDetailsRequest) {
	alipayCustomerInquireDetailsRequest := &AlipayCustomerInquireDetailsRequest{}
	alipayRequest := request.NewAlipayRequest(alipayCustomerInquireDetailsRequest, "/ams/api/v1/billing/customer/inquireDetails", &responseBilling.AlipayCustomerInquireDetailsResponse{})
	return alipayRequest, alipayCustomerInquireDetailsRequest
}

func (alipayCustomerInquireDetailsRequest *AlipayCustomerInquireDetailsRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayCustomerInquireDetailsRequest, "/ams/api/v1/billing/customer/inquireDetails", &responseBilling.AlipayCustomerInquireDetailsResponse{})
}
