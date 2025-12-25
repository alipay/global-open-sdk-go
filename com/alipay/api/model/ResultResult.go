package model

type ResultResult struct {
	ResultCode    string           `json:"resultCode,omitempty"`
	ResultStatus  ResultStatusType `json:"resultStatus,omitempty"`
	ResultMessage string           `json:"resultMessage,omitempty"`
}
