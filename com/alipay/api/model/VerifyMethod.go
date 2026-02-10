package model

type VerifyMethod struct {
	VerifyMethodType string `json:"verifyMethodType,omitempty"`
	OtpValue         string `json:"otpValue,omitempty"`
}
