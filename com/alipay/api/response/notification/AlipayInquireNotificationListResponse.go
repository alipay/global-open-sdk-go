package responseNotification

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayInquireNotificationListResponse struct {
	response.AlipayResponse
	Result            *model.Result              `json:"result,omitempty"`
	Notifications     []*model.NotificationEvent `json:"notifications,omitempty"`
	HasMore           bool                       `json:"hasMore,omitempty"`
	NextStartingAfter string                     `json:"nextStartingAfter,omitempty"`
}
