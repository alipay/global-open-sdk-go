package notification

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseNotification "github.com/alipay/global-open-sdk-go/com/alipay/api/response/notification"
)

type AlipayInquireNotificationListRequest struct {
	NotifyType          []string `json:"notifyType,omitempty"`
	EventCreateTimeFrom string   `json:"eventCreateTimeFrom,omitempty"`
	EventCreateTimeTo   string   `json:"eventCreateTimeTo,omitempty"`
	Limit               int32    `json:"limit,omitempty"`
	StartingAfter       string   `json:"startingAfter,omitempty"`
}

func NewAlipayInquireNotificationListRequest() (*request.AlipayRequest, *AlipayInquireNotificationListRequest) {
	alipayInquireNotificationListRequest := &AlipayInquireNotificationListRequest{}
	alipayRequest := request.NewAlipayRequest(alipayInquireNotificationListRequest, "/ams/api/v1/notification/inquireNotificationList", &responseNotification.AlipayInquireNotificationListResponse{})
	return alipayRequest, alipayInquireNotificationListRequest
}

func (alipayInquireNotificationListRequest *AlipayInquireNotificationListRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayInquireNotificationListRequest, "/ams/api/v1/notification/inquireNotificationList", &responseNotification.AlipayInquireNotificationListResponse{})
}
