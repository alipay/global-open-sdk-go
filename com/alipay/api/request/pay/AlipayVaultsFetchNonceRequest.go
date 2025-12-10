package pay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responsePay "github.com/alipay/global-open-sdk-go/com/alipay/api/response/pay"
)

type AlipayVaultsFetchNonceRequest struct {
	Card *model.Card `json:"card,omitempty"`
}

func NewAlipayVaultsFetchNonceRequest() (*request.AlipayRequest, *AlipayVaultsFetchNonceRequest) {
	alipayVaultsFetchNonceRequest := &AlipayVaultsFetchNonceRequest{}
	alipayRequest := request.NewAlipayRequest(alipayVaultsFetchNonceRequest, "/ams/api/v1/vaults/fetchNonce", &responsePay.AlipayVaultsFetchNonceResponse{})
	return alipayRequest, alipayVaultsFetchNonceRequest
}

func (alipayVaultsFetchNonceRequest *AlipayVaultsFetchNonceRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayVaultsFetchNonceRequest, "/ams/api/v1/vaults/fetchNonce", &responsePay.AlipayVaultsFetchNonceResponse{})
}
