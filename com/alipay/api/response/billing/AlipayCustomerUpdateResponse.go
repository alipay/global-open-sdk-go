package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayCustomerUpdateResponse struct {
	response.AlipayResponse
	Result              *model.Result `json:"result,omitempty"`
	CustomerId          string        `json:"customerId,omitempty"`
	Status              string        `json:"status,omitempty"`
	PhoneNo             string        `json:"phoneNo,omitempty"`
	CountryCode         string        `json:"countryCode,omitempty"`
	BillingEmail        string        `json:"billingEmail,omitempty"`
	ShippingFirstName   string        `json:"shippingFirstName,omitempty"`
	ShippingLastName    string        `json:"shippingLastName,omitempty"`
	ShippingCountryCode string        `json:"shippingCountryCode,omitempty"`
}
