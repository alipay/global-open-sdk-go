package responsePay

import "github.com/alipay/global-open-sdk-go/com/alipay/api/response"

type AlipayPayCancelResponse struct {
	response.AlipayResponse
	PaymentId        string `json:"paymentId"`
	PaymentRequestId string `json:"paymentRequestId"`
	CancelTime       string `json:"cancelTime"`
}
