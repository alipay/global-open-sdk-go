package responseMarketplace

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type PaymentMethodResponse struct {
	response.AlipayResponse
	PaymentMethodType string `json:"paymentMethodType,omitempty"`
	CustomerId        string `json:"customerId,omitempty"`
}
