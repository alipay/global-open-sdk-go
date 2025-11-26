package model

type BrowserInfo struct {
	AcceptHeader      string `json:"acceptHeader,omitempty"`
	JavaEnabled       bool   `json:"javaEnabled,omitempty"`
	JavaScriptEnabled bool   `json:"javaScriptEnabled,omitempty"`
	Language          string `json:"language,omitempty"`
	UserAgent         string `json:"userAgent,omitempty"`
}
