package responsePay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayPaymentSessionResponse struct {
	response.AlipayResponse
	Result                   *model.Result `json:"result,omitempty"`
	PaymentSessionData       string        `json:"paymentSessionData,omitempty"`
	PaymentSessionExpiryTime string        `json:"paymentSessionExpiryTime,omitempty"`
	PaymentSessionId         string        `json:"paymentSessionId,omitempty"`
	NormalUrl                string        `json:"normalUrl,omitempty"`
	Url                      string        `json:"url,omitempty"`
}
