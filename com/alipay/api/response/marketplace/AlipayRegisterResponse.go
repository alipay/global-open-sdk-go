package responseMarketplace

import "github.com/alipay/global-open-sdk-go/com/alipay/api/response"

type AlipayRegisterResponse struct {
	response.AlipayResponse
	RegistrationStatus string `json:"registrationStatus,omitempty"`
}
