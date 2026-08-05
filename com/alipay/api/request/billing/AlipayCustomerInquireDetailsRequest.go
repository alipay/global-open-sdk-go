package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayCustomerInquireDetailsRequest struct {
	CustomerId          string `json:"customerId,omitempty"`
	PhoneNo             string `json:"phoneNo,omitempty"`
	CountryCode         string `json:"countryCode,omitempty"`
	BillingEmail        string `json:"billingEmail,omitempty"`
	ShippingFirstName   string `json:"shippingFirstName,omitempty"`
	ShippingLastName    string `json:"shippingLastName,omitempty"`
	ShippingCountryCode string `json:"shippingCountryCode,omitempty"`
}

func NewAlipayCustomerInquireDetailsRequest() (*request.AlipayRequest, *AlipayCustomerInquireDetailsRequest) {
	alipayCustomerInquireDetailsRequest := &AlipayCustomerInquireDetailsRequest{}
	alipayRequest := request.NewAlipayRequest(alipayCustomerInquireDetailsRequest, "/ams/api/v1/billing/customer/inquireDetails", &responseBilling.AlipayCustomerInquireDetailsResponse{})
	return alipayRequest, alipayCustomerInquireDetailsRequest
}

func (alipayCustomerInquireDetailsRequest *AlipayCustomerInquireDetailsRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayCustomerInquireDetailsRequest, "/ams/api/v1/billing/customer/inquireDetails", &responseBilling.AlipayCustomerInquireDetailsResponse{})
}
