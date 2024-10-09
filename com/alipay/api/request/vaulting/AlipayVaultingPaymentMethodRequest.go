package vaulting

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseVaulting "github.com/alipay/global-open-sdk-go/com/alipay/api/response/vaulting"
)

type AlipayVaultingPaymentMethodRequest struct {
	VaultingRequestId       string                     `json:"vaultingRequestId,omitempty"`
	VaultingNotificationUrl string                     `json:"vaultingNotificationUrl,omitempty"`
	RedirectUrl             string                     `json:"redirectUrl,omitempty"`
	MerchantRegion          string                     `json:"merchantRegion,omitempty"`
	PaymentMethodDetail     *model.PaymentMethodDetail `json:"paymentMethodDetail,omitempty"`
	Env                     *model.Env                 `json:"env,omitempty"`
}

func NewAlipayVaultingPaymentMethodRequest() (*request.AlipayRequest, *AlipayVaultingPaymentMethodRequest) {
	alipayVaultingPaymentMethodRequest := &AlipayVaultingPaymentMethodRequest{}
	alipayRequest := request.NewAlipayRequest(alipayVaultingPaymentMethodRequest, model.VAULT_PAYMENT_METHOD_PATH, &responseVaulting.AlipayVaultingPaymentMethodResponse{})
	return alipayRequest, alipayVaultingPaymentMethodRequest
}
