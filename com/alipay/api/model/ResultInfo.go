package model

type ResultInfo struct {
	ResultCode    string `json:"resultCode,omitempty"`
	ResultStatus  string `json:"resultStatus,omitempty"`
	ResultMessage string `json:"resultMessage,omitempty"`
}
