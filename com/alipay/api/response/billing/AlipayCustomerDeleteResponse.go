package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayCustomerDeleteResponse struct {
	response.AlipayResponse
	Result     *model.Result `json:"result,omitempty"`
	CustomerId string        `json:"customerId,omitempty"`
	Status     string        `json:"status,omitempty"`
}
