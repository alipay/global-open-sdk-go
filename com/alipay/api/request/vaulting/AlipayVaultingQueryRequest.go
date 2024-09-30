package vaulting

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseVaulting "github.com/alipay/global-open-sdk-go/com/alipay/api/response/vaulting"
)

type AlipayVaultingQueryRequest struct {
	VaultingRequestId string `json:"vaultingRequestId,omitempty"`
}

func NewAlipayVaultingQueryRequest() (*request.AlipayRequest, *AlipayVaultingQueryRequest) {
	alipayVaultingPaymentMethodRequest := &AlipayVaultingQueryRequest{}
	alipayRequest := request.NewAlipayRequest(alipayVaultingPaymentMethodRequest, model.INQUIRE_VAULTING_PATH, &responseVaulting.AlipayVaultingQueryResponse{})
	return alipayRequest, alipayVaultingPaymentMethodRequest
}
