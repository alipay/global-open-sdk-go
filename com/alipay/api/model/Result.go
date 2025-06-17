package model

type Result struct {
	ResultCode    string           `json:"resultCode,omitempty"`
	ResultStatus  ResultStatusType `json:"resultStatus,omitempty"`
	ResultMessage string           `json:"resultMessage,omitempty"`
}
