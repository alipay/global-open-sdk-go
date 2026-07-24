package model

type Meter struct {
	EventName string   `json:"eventName,omitempty"`
	Events    []*Event `json:"events,omitempty"`
}
