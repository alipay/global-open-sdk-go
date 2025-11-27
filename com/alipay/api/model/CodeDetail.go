package model

type CodeDetail struct {
	CodeValueType CodeValueType `json:"codeValueType,omitempty"`
	CodeValue     string        `json:"codeValue,omitempty"`
	DisplayType   DisplayType   `json:"displayType,omitempty"`
}
