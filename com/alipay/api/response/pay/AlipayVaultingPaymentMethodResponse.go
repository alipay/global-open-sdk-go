package responsePay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayVaultingPaymentMethodResponse struct {
	response.AlipayResponse
	Result              *model.Result              `json:"result,omitempty"`
	VaultingRequestId   string                     `json:"vaultingRequestId,omitempty"`
	PaymentMethodDetail *model.PaymentMethodDetail `json:"paymentMethodDetail,omitempty"`
	NormalUrl           string                     `json:"normalUrl,omitempty"`
	SchemeUrl           string                     `json:"schemeUrl,omitempty"`
	ApplinkUrl          string                     `json:"applinkUrl,omitempty"`
}
