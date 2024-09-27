package responsePay

import "github.com/alipay/global-open-sdk-go/com/alipay/api/response"

type AlipayPaymentSessionResponse struct {
	response.AlipayResponse
	PaymentSessionData       string `json:"paymentSessionData"`
	PaymentSessionExpiryTime string `json:"paymentSessionExpiryTime"`
	PaymentSessionId         string `json:"paymentSessionId"`
	NormalUrl                string `json:"normalUrl"`
}
