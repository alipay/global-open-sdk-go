package responsePay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayUpdateAmountResponse struct {
	response.AlipayResponse
	Result          *model.Result `json:"result,omitempty"`
	UpdateRequestId string        `json:"updateRequestId,omitempty"`
	PaymentId       string        `json:"paymentId,omitempty"`
	Amount          *model.Amount `json:"amount,omitempty"`
}
