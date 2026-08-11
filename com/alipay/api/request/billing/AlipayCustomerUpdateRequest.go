package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayCustomerUpdateRequest struct {
	CustomerId            string   `json:"customerId,omitempty"`
	ReferenceCustomerId   string   `json:"referenceCustomerId,omitempty"`
	AlipayUserId          string   `json:"alipayUserId,omitempty"`
	Email                 string   `json:"email,omitempty"`
	FirstName             string   `json:"firstName,omitempty"`
	LastName              string   `json:"lastName,omitempty"`
	Country               string   `json:"country,omitempty"`
	State                 string   `json:"state,omitempty"`
	City                  string   `json:"city,omitempty"`
	Address               string   `json:"address,omitempty"`
	AddressDetail         string   `json:"addressDetail,omitempty"`
	Zipcode               string   `json:"zipcode,omitempty"`
	ShippingPhone         string   `json:"shippingPhone,omitempty"`
	ShippingCountry       string   `json:"shippingCountry,omitempty"`
	ShippingState         string   `json:"shippingState,omitempty"`
	ShippingCity          string   `json:"shippingCity,omitempty"`
	ShippingAddress       string   `json:"shippingAddress,omitempty"`
	ShippingAddressDetail string   `json:"shippingAddressDetail,omitempty"`
	ShippingZipcode       string   `json:"shippingZipcode,omitempty"`
	Description           string   `json:"description,omitempty"`
	Currency              string   `json:"currency,omitempty"`
	PreferredLocales      []string `json:"preferredLocales,omitempty"`
	DefaultPaymentMethod  string   `json:"defaultPaymentMethod,omitempty"`
	Metadata              string   `json:"metadata,omitempty"`
	PhoneNo               string   `json:"phoneNo,omitempty"`
	CountryCode           string   `json:"countryCode,omitempty"`
	BillingEmail          string   `json:"billingEmail,omitempty"`
	ShippingFirstName     string   `json:"shippingFirstName,omitempty"`
	ShippingLastName      string   `json:"shippingLastName,omitempty"`
	ShippingCountryCode   string   `json:"shippingCountryCode,omitempty"`
}

func NewAlipayCustomerUpdateRequest() (*request.AlipayRequest, *AlipayCustomerUpdateRequest) {
	alipayCustomerUpdateRequest := &AlipayCustomerUpdateRequest{}
	alipayRequest := request.NewAlipayRequest(alipayCustomerUpdateRequest, "/ams/api/v1/billing/customer/update", &responseBilling.AlipayCustomerUpdateResponse{})
	return alipayRequest, alipayCustomerUpdateRequest
}

func (alipayCustomerUpdateRequest *AlipayCustomerUpdateRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayCustomerUpdateRequest, "/ams/api/v1/billing/customer/update", &responseBilling.AlipayCustomerUpdateResponse{})
}
