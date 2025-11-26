package pay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responsePay "github.com/alipay/global-open-sdk-go/com/alipay/api/response/pay"
)

type AlipayVaultingQueryRequest struct {
	VaultingRequestId string `json:"vaultingRequestId,omitempty"`
}

func NewAlipayVaultingQueryRequest() (*request.AlipayRequest, *AlipayVaultingQueryRequest) {
	alipayVaultingQueryRequest := &AlipayVaultingQueryRequest{}
	alipayRequest := request.NewAlipayRequest(alipayVaultingQueryRequest, "/ams/api/v1/vaults/inquireVaulting", &responsePay.AlipayVaultingQueryResponse{})
	return alipayRequest, alipayVaultingQueryRequest
}

func (alipayVaultingQueryRequest *AlipayVaultingQueryRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayVaultingQueryRequest, "/ams/api/v1/vaults/inquireVaulting", &responsePay.AlipayVaultingQueryResponse{})
}
