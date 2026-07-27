package model

type Error struct {
	Index          int32  `json:"index,omitempty"`
	GroupIndex     int32  `json:"groupIndex,omitempty"`
	EventName      string `json:"eventName,omitempty"`
	IdempotencyKey string `json:"idempotencyKey,omitempty"`
	CustomerId     string `json:"customerId,omitempty"`
	ErrorCode      string `json:"errorCode,omitempty"`
}
