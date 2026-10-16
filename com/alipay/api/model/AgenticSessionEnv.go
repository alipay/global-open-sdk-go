package model

type AgenticSessionEnv struct {
	TerminalType string `json:"terminalType,omitempty"`
	DeviceBrand  string `json:"deviceBrand,omitempty"`
	DeviceId     string `json:"deviceId,omitempty"`
	OsType       string `json:"osType,omitempty"`
	ClientIp     string `json:"clientIp,omitempty"`
}
