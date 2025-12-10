package responsePay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayVaultsFetchNonceResponse struct {
	response.AlipayResponse
	CardToken string        `json:"cardToken,omitempty"`
	Result    *model.Result `json:"result,omitempty"`
}
