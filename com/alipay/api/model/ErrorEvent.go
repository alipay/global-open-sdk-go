package model

type ErrorEvent struct {
	ErrorCode      string        `json:"errorCode,omitempty"`
	IdempotencyKey string        `json:"idempotencyKey,omitempty"`
	EventTimestamp int64         `json:"eventTimestamp,omitempty"`
	Payload        *EventPayload `json:"payload,omitempty"`
}
