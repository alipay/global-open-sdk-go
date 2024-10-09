package responseVaulting

import "github.com/alipay/global-open-sdk-go/com/alipay/api/response"

type AlipayVaultingSessionResponse struct {
	response.AlipayResponse
	VaultingSessionData       string `json:"vaultingSessionData,omitempty"`
	VaultingSessionId         string `json:"vaultingSessionId,omitempty"`
	VaultingSessionExpiryTime string `json:"vaultingSessionExpiryTime,omitempty"`
}
