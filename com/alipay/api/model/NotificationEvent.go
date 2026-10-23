package model

type NotificationEvent struct {
	EventId             string                 `json:"eventId,omitempty"`
	NotifyType          string                 `json:"notifyType,omitempty"`
	EventCreateTime     string                 `json:"eventCreateTime,omitempty"`
	DeliveryStatus      string                 `json:"deliveryStatus,omitempty"`
	ReferenceMerchantId string                 `json:"referenceMerchantId,omitempty"`
	NotificationData    map[string]interface{} `json:"notificationData,omitempty"`
}
