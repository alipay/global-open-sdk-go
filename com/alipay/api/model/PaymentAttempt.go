package model

type PaymentAttempt struct {
	AttemptAt       string `json:"attemptAt,omitempty"`
	AttemptResponse string `json:"attemptResponse,omitempty"`
}
