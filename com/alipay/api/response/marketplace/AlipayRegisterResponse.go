package responseMarketplace

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayRegisterResponse struct {
	response.AlipayResponse
	Result             *model.Result `json:"result,omitempty"`
	RegistrationStatus string        `json:"registrationStatus,omitempty"`
}
