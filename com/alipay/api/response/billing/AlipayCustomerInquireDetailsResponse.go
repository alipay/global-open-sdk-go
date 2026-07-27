package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayCustomerInquireDetailsResponse struct {
	response.AlipayResponse
	Result                *model.Result     `json:"result,omitempty"`
	CustomerId            string            `json:"customerId,omitempty"`
	CustomerRequestId     string            `json:"customerRequestId,omitempty"`
	AlipayUserId          string            `json:"alipayUserId,omitempty"`
	Email                 string            `json:"email,omitempty"`
	FirstName             string            `json:"firstName,omitempty"`
	LastName              string            `json:"lastName,omitempty"`
	MobileNo              string            `json:"mobileNo,omitempty"`
	Country               string            `json:"country,omitempty"`
	State                 string            `json:"state,omitempty"`
	City                  string            `json:"city,omitempty"`
	Address               string            `json:"address,omitempty"`
	AddressDetail         string            `json:"addressDetail,omitempty"`
	Zipcode               string            `json:"zipcode,omitempty"`
	ShippingName          string            `json:"shippingName,omitempty"`
	ShippingPhone         string            `json:"shippingPhone,omitempty"`
	ShippingCountry       string            `json:"shippingCountry,omitempty"`
	ShippingState         string            `json:"shippingState,omitempty"`
	ShippingCity          string            `json:"shippingCity,omitempty"`
	ShippingAddress       string            `json:"shippingAddress,omitempty"`
	ShippingAddressDetail string            `json:"shippingAddressDetail,omitempty"`
	ShippingZipCode       string            `json:"shippingZipCode,omitempty"`
	Description           string            `json:"description,omitempty"`
	Currency              string            `json:"currency,omitempty"`
	PreferredLocales      string            `json:"preferredLocales,omitempty"`
	DefaultPaymentMethod  string            `json:"defaultPaymentMethod,omitempty"`
	Status                string            `json:"status,omitempty"`
	ReferenceCustomerId   string            `json:"referenceCustomerId,omitempty"`
	Metadata              map[string]string `json:"metadata,omitempty"`
}
