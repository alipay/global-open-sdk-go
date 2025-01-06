package responseAuth

import "github.com/alipay/global-open-sdk-go/com/alipay/api/response"

type AlipayAuthCreateSessionResponse struct {
	response.AlipayResponse
	PaymentSessionId         string `json:"paymentSessionId,omitempty"`
	PaymentSessionData       string `json:"paymentSessionData,omitempty"`
	PaymentSessionExpiryTime string `json:"paymentSessionExpiryTime,omitempty"`
}
