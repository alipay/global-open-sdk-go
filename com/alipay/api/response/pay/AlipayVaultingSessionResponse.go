package responsePay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayVaultingSessionResponse struct {
	response.AlipayResponse
	Result                    *model.Result `json:"result,omitempty"`
	VaultingSessionData       string        `json:"vaultingSessionData,omitempty"`
	VaultingSessionId         string        `json:"vaultingSessionId,omitempty"`
	VaultingSessionExpiryTime string        `json:"vaultingSessionExpiryTime,omitempty"`
	NormalUrl                 string        `json:"normalUrl,omitempty"`
}
