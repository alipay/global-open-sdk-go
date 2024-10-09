package responseVaulting

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayVaultingQueryResponse struct {
	response.AlipayResponse
	VaultingRequestId   string                     `json:"vaultingRequestId,omitempty"`
	NormalUrl           string                     `json:"normalUrl,omitempty"`
	SchemeUrl           string                     `json:"schemeUrl,omitempty"`
	ApplinkUrl          string                     `json:"applinkUrl,omitempty"`
	VaultingStatus      string                     `json:"vaultingStatus,omitempty"`
	PaymentMethodDetail *model.PaymentMethodDetail `json:"paymentMethodDetail,omitempty"`
}
