package notification

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseNotification "github.com/alipay/global-open-sdk-go/com/alipay/api/response/notification"
)

type AlipayInquireNotificationRequest struct {
	EventId string `json:"eventId,omitempty"`
}

func NewAlipayInquireNotificationRequest() (*request.AlipayRequest, *AlipayInquireNotificationRequest) {
	alipayInquireNotificationRequest := &AlipayInquireNotificationRequest{}
	alipayRequest := request.NewAlipayRequest(alipayInquireNotificationRequest, "/ams/api/v1/notification/inquireNotification", &responseNotification.AlipayInquireNotificationResponse{})
	return alipayRequest, alipayInquireNotificationRequest
}

func (alipayInquireNotificationRequest *AlipayInquireNotificationRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayInquireNotificationRequest, "/ams/api/v1/notification/inquireNotification", &responseNotification.AlipayInquireNotificationResponse{})
}
