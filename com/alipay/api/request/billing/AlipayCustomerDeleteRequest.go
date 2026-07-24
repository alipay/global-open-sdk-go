package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayCustomerDeleteRequest struct {
	CustomerId string `json:"customerId,omitempty"`
}

func NewAlipayCustomerDeleteRequest() (*request.AlipayRequest, *AlipayCustomerDeleteRequest) {
	alipayCustomerDeleteRequest := &AlipayCustomerDeleteRequest{}
	alipayRequest := request.NewAlipayRequest(alipayCustomerDeleteRequest, "/ams/api/v1/billing/customer/delete", &responseBilling.AlipayCustomerDeleteResponse{})
	return alipayRequest, alipayCustomerDeleteRequest
}

func (alipayCustomerDeleteRequest *AlipayCustomerDeleteRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayCustomerDeleteRequest, "/ams/api/v1/billing/customer/delete", &responseBilling.AlipayCustomerDeleteResponse{})
}
