package responsePay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayPayCancelResponse struct {
	response.AlipayResponse
	Result           *model.Result `json:"result,omitempty"`
	PaymentId        string        `json:"paymentId,omitempty"`
	PaymentRequestId string        `json:"paymentRequestId,omitempty"`
	CancelTime       string        `json:"cancelTime,omitempty"`
}
