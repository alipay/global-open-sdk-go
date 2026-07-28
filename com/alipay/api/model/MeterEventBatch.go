package model

type MeterEventBatch struct {
	EventName string   `json:"eventName,omitempty"`
	Events    []*Event `json:"events,omitempty"`
}
