package pay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responsePay "github.com/alipay/global-open-sdk-go/com/alipay/api/response/pay"
)

type AlipayVaultingPaymentMethodRequest struct {
	MerchantAccountId       string                     `json:"merchantAccountId,omitempty"`
	Metadata                string                     `json:"metadata,omitempty"`
	VaultingRequestId       string                     `json:"vaultingRequestId,omitempty"`
	VaultingNotificationUrl string                     `json:"vaultingNotificationUrl,omitempty"`
	RedirectUrl             string                     `json:"redirectUrl,omitempty"`
	MerchantRegion          string                     `json:"merchantRegion,omitempty"`
	PaymentMethodDetail     *model.PaymentMethodDetail `json:"paymentMethodDetail,omitempty"`
	Env                     *model.Env                 `json:"env,omitempty"`
	VaultingCurrency        string                     `json:"vaultingCurrency,omitempty"`
	CustomizedInfo          *model.CustomizedInfo      `json:"customizedInfo,omitempty"`
}

func NewAlipayVaultingPaymentMethodRequest() (*request.AlipayRequest, *AlipayVaultingPaymentMethodRequest) {
	alipayVaultingPaymentMethodRequest := &AlipayVaultingPaymentMethodRequest{}
	alipayRequest := request.NewAlipayRequest(alipayVaultingPaymentMethodRequest, "/ams/api/v1/vaults/vaultPaymentMethod", &responsePay.AlipayVaultingPaymentMethodResponse{})
	return alipayRequest, alipayVaultingPaymentMethodRequest
}

func (alipayVaultingPaymentMethodRequest *AlipayVaultingPaymentMethodRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayVaultingPaymentMethodRequest, "/ams/api/v1/vaults/vaultPaymentMethod", &responsePay.AlipayVaultingPaymentMethodResponse{})
}
