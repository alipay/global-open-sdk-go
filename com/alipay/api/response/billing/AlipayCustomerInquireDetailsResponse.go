package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayCustomerInquireDetailsResponse struct {
	response.AlipayResponse
	Result                *model.Result `json:"result,omitempty"`
	CustomerId            string        `json:"customerId,omitempty"`
	CustomerRequestId     string        `json:"customerRequestId,omitempty"`
	AlipayUserId          string        `json:"alipayUserId,omitempty"`
	Email                 string        `json:"email,omitempty"`
	FirstName             string        `json:"firstName,omitempty"`
	LastName              string        `json:"lastName,omitempty"`
	Country               string        `json:"country,omitempty"`
	State                 string        `json:"state,omitempty"`
	City                  string        `json:"city,omitempty"`
	Address               string        `json:"address,omitempty"`
	AddressDetail         string        `json:"addressDetail,omitempty"`
	Zipcode               string        `json:"zipcode,omitempty"`
	ShippingPhone         string        `json:"shippingPhone,omitempty"`
	ShippingCountry       string        `json:"shippingCountry,omitempty"`
	ShippingState         string        `json:"shippingState,omitempty"`
	ShippingCity          string        `json:"shippingCity,omitempty"`
	ShippingAddress       string        `json:"shippingAddress,omitempty"`
	ShippingAddressDetail string        `json:"shippingAddressDetail,omitempty"`
	Description           string        `json:"description,omitempty"`
	Currency              string        `json:"currency,omitempty"`
	PreferredLocales      []string      `json:"preferredLocales,omitempty"`
	DefaultPaymentMethod  string        `json:"defaultPaymentMethod,omitempty"`
	Status                string        `json:"status,omitempty"`
	ReferenceCustomerId   string        `json:"referenceCustomerId,omitempty"`
	Metadata              string        `json:"metadata,omitempty"`
	PhoneNo               string        `json:"phoneNo,omitempty"`
	CountryCode           string        `json:"countryCode,omitempty"`
	BillingEmail          string        `json:"billingEmail,omitempty"`
	ShippingFirstName     string        `json:"shippingFirstName,omitempty"`
	ShippingLastName      string        `json:"shippingLastName,omitempty"`
	ShippingCountryCode   string        `json:"shippingCountryCode,omitempty"`
	ShippingZipcode       string        `json:"shippingZipcode,omitempty"`
	GmtCreate             string        `json:"gmtCreate,omitempty"`
}
