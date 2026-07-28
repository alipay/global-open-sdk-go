package model

type CouponUpdateResult struct {
	ResultStatus  string `json:"resultStatus,omitempty"`
	ResultCode    string `json:"resultCode,omitempty"`
	ResultMessage string `json:"resultMessage,omitempty"`
}
