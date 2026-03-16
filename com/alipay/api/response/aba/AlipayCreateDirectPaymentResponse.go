package responseAba

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayCreateDirectPaymentResponse struct {
	response.AlipayResponse
	PaymentId        string               `json:"paymentId,omitempty"`
	PaymentRequestId string               `json:"paymentRequestId,omitempty"`
	PayToMethod      *model.PaymentMethod `json:"payToMethod,omitempty"`
	PayFromAmount    *model.Amount        `json:"payFromAmount,omitempty"`
	PayToAmount      *model.Amount        `json:"payToAmount,omitempty"`
	PaymentTime      string               `json:"paymentTime,omitempty"`
	Result           *model.Result        `json:"result,omitempty"`
}
