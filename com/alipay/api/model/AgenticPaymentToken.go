package model

type AgenticPaymentToken struct {
	Token                string `json:"token,omitempty"`
	TokenExpirationYear  string `json:"tokenExpirationYear,omitempty"`
	TokenExpirationMonth string `json:"tokenExpirationMonth,omitempty"`
	TokenExpiration      string `json:"tokenExpiration,omitempty"`
	Status               string `json:"status,omitempty"`
	DynamicDataType      string `json:"dynamicDataType,omitempty"`
	DynamicDataValue     string `json:"dynamicDataValue,omitempty"`
}
