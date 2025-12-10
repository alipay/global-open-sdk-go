package model

type Result struct {
	ResultCode    string                 `json:"resultCode,omitempty"`
	ResultStatus  map[string]interface{} `json:"resultStatus,omitempty"`
	ResultMessage string                 `json:"resultMessage,omitempty"`
}
