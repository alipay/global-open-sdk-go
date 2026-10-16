package model

type AgenticAccountInfoEnv struct {
	TerminalType   string `json:"terminalType,omitempty"`
	OsType         string `json:"osType,omitempty"`
	DeviceId       string `json:"deviceId,omitempty"`
	DeviceTokenId  string `json:"deviceTokenId,omitempty"`
	ClientIp       string `json:"clientIp,omitempty"`
	ExtendInfo     string `json:"extendInfo,omitempty"`
	DeviceLanguage string `json:"deviceLanguage,omitempty"`
	UserAgent      string `json:"userAgent,omitempty"`
}
