package model

type EventPayload struct {
	Value      string `json:"value,omitempty"`
	CustomerId string `json:"customerId,omitempty"`
}
