package responseAba

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayUpdateCardStatusResponse struct {
	response.AlipayResponse
	Status    string        `json:"status,omitempty"`
	RequestId string        `json:"requestId,omitempty"`
	Result    *model.Result `json:"result,omitempty"`
}
