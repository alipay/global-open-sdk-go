package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayMeterCreateSessionResponse struct {
	response.AlipayResponse
	Result            *model.Result `json:"result,omitempty"`
	SessionId         string        `json:"sessionId,omitempty"`
	SessionExpiryTime string        `json:"sessionExpiryTime,omitempty"`
}
