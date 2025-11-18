package model

type ResultProperties struct {
	ResultCode    *ResultPropertiesResultCode   `json:"resultCode,omitempty"`
	ResultStatus  *ResultPropertiesResultStatus `json:"resultStatus,omitempty"`
	ResultMessage *ResultPropertiesResultCode   `json:"resultMessage,omitempty"`
}
