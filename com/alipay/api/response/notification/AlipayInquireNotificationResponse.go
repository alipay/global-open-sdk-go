package responseNotification

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayInquireNotificationResponse struct {
	response.AlipayResponse
	Result       *model.Result            `json:"result,omitempty"`
	Notification *model.NotificationEvent `json:"notification,omitempty"`
}
