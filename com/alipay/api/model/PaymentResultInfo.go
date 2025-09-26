package model

type PaymentResultInfo struct {
	RefusalCodeRaw       string         `json:"refusalCodeRaw,omitempty"`
	RefusalReasonRaw     string         `json:"refusalReasonRaw,omitempty"`
	MerchantAdviceCode   string         `json:"merchantAdviceCode,omitempty"`
	AcquirerInfo         *AcquirerInfo  `json:"acquirerInfo,omitempty"`
	CardNo               string         `json:"cardNo,omitempty"`
	CardBrand            string         `json:"cardBrand,omitempty"`
	CardToken            string         `json:"cardToken,omitempty"`
	IssuingCountry       string         `json:"issuingCountry,omitempty"`
	Funding              string         `json:"funding,omitempty"`
	PaymentMethodRegion  string         `json:"paymentMethodRegion,omitempty"`
	ThreeDSResult        *ThreeDSResult `json:"threeDSResult,omitempty"`
	AvsResultRaw         string         `json:"avsResultRaw,omitempty"`
	CvvResultRaw         string         `json:"cvvResultRaw,omitempty"`
	NetworkTransactionId string         `json:"networkTransactionId,omitempty"`
	CreditPayPlan        *CreditPayPlan `json:"creditPayPlan,omitempty"`
	CardholderName       string         `json:"cardholderName,omitempty"`
	CardBin              string         `json:"cardBin,omitempty"`
	LastFour             string         `json:"lastFour,omitempty"`
	ExpiryMonth          string         `json:"expiryMonth,omitempty"`
	ExpiryYear           string         `json:"expiryYear,omitempty"`
}
