package model

type AuthorizationControl struct {
	CardActiveTime              string           `json:"cardActiveTime,omitempty"`
	CardCancelTime              string           `json:"cardCancelTime,omitempty"`
	AllowedMerchantCategoryList []string         `json:"allowedMerchantCategoryList,omitempty"`
	AllowedAuthTimes            int32            `json:"allowedAuthTimes,omitempty"`
	AllowedCurrencies           []string         `json:"allowedCurrencies,omitempty"`
	CardLimitDetail             *CardLimitDetail `json:"cardLimitDetail,omitempty"`
	CardLimitInfo               *CardLimitInfo   `json:"cardLimitInfo,omitempty"`
}
