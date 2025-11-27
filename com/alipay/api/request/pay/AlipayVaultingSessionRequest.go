package pay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responsePay "github.com/alipay/global-open-sdk-go/com/alipay/api/response/pay"
)

type AlipayVaultingSessionRequest struct {
	PaymentMethodType       string `json:"paymentMethodType,omitempty"`
	VaultingRequestId       string `json:"vaultingRequestId,omitempty"`
	VaultingNotificationUrl string `json:"vaultingNotificationUrl,omitempty"`
	RedirectUrl             string `json:"redirectUrl,omitempty"`
	MerchantRegion          string `json:"merchantRegion,omitempty"`
	Is3DSAuthentication     bool   `json:"is3DSAuthentication,omitempty"`
}

func NewAlipayVaultingSessionRequest() (*request.AlipayRequest, *AlipayVaultingSessionRequest) {
	alipayVaultingSessionRequest := &AlipayVaultingSessionRequest{}
	alipayRequest := request.NewAlipayRequest(alipayVaultingSessionRequest, "/ams/api/v1/vaults/createVaultingSession", &responsePay.AlipayVaultingSessionResponse{})
	return alipayRequest, alipayVaultingSessionRequest
}

func (alipayVaultingSessionRequest *AlipayVaultingSessionRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayVaultingSessionRequest, "/ams/api/v1/vaults/createVaultingSession", &responsePay.AlipayVaultingSessionResponse{})
}
