package model

type Event struct {
	IdempotencyKey string        `json:"idempotencyKey,omitempty"`
	EventTimestamp int32         `json:"eventTimestamp,omitempty"`
	Payload        *EventPayload `json:"payload,omitempty"`
}
