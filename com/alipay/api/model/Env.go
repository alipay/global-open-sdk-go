package model

type Env struct {
	TerminalType             TerminalType `json:"terminalType,omitempty"`
	OsType                   OsType       `json:"osType,omitempty"`
	UserAgent                string       `json:"userAgent,omitempty"`
	DeviceTokenId            string       `json:"deviceTokenId,omitempty"`
	ClientIp                 string       `json:"clientIp,omitempty"`
	CookieId                 string       `json:"cookieId,omitempty"`
	ExtendInfo               string       `json:"extendInfo,omitempty"`
	StoreTerminalId          string       `json:"storeTerminalId,omitempty"`
	StoreTerminalRequestTime string       `json:"storeTerminalRequestTime,omitempty"`
	BrowserInfo              *BrowserInfo `json:"browserInfo,omitempty"`
	ColorDepth               string       `json:"colorDepth,omitempty"`
	ScreenHeight             string       `json:"screenHeight,omitempty"`
	ScreenWidth              string       `json:"screenWidth,omitempty"`
	TimeZoneOffset           int32        `json:"timeZoneOffset,omitempty"`
	DeviceBrand              string       `json:"deviceBrand,omitempty"`
	DeviceModel              string       `json:"deviceModel,omitempty"`
	DeviceLanguage           string       `json:"deviceLanguage,omitempty"`
	DeviceId                 string       `json:"deviceId,omitempty"`
	EnvType                  string       `json:"envType,omitempty"`
}
