package model

type InquireCardDetailAuthorizationControl struct {
	CardActiveTime              string                              `json:"cardActiveTime,omitempty"`
	CardCancelTime              string                              `json:"cardCancelTime,omitempty"`
	AllowedMerchantCategoryList []string                            `json:"allowedMerchantCategoryList,omitempty"`
	AllowedAuthTimes            int32                               `json:"allowedAuthTimes,omitempty"`
	AllowedCurrencies           []string                            `json:"allowedCurrencies,omitempty"`
	PaymentPreference           *InquireCardDetailPaymentPreference `json:"paymentPreference,omitempty"`
	ThreeDSMode                 string                              `json:"threeDSMode,omitempty"`
	PhoneNo                     string                              `json:"phoneNo,omitempty"`
	Email                       string                              `json:"email,omitempty"`
	CardLimitDetail             *CardLimitDetail                    `json:"cardLimitDetail,omitempty"`
	CardLimitInfo               *CardLimitInfo                      `json:"cardLimitInfo,omitempty"`
	RefundPreference            *RefundPreference                   `json:"refundPreference,omitempty"`
}
