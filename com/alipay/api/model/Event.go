package model

type Event struct {
	IdempotencyKey string        `json:"idempotencyKey,omitempty"`
	EventTimestamp int64         `json:"eventTimestamp,omitempty"`
	Payload        *EventPayload `json:"payload,omitempty"`
}
