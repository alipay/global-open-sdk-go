package notify

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayNotify struct {
	NotifyType string          `json:"notifyType,omitempty"`
	Result     response.Result `json:"result,omitempty"`
}
