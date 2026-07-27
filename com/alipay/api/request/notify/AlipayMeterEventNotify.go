package notify

import "github.com/alipay/global-open-sdk-go/com/alipay/api/model"

type AlipayMeterEventNotify struct {
	AlipayNotify
	EventName   string            `json:"eventName,omitempty"`
	ErrorEvents *model.ErrorEvent `json:"errorEvents,omitempty"`
}
