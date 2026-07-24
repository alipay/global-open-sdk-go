package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayCustomerCreateResponse struct {
	response.AlipayResponse
	Result            *model.Result `json:"result,omitempty"`
	CustomerId        string        `json:"customerId,omitempty"`
	CustomerRequestId string        `json:"customerRequestId,omitempty"`
	Email             string        `json:"email,omitempty"`
	Status            string        `json:"status,omitempty"`
}
