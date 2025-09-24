package model

type PaymentMethod struct {
	PaymentMethodType           string         `json:"paymentMethodType,omitempty"`
	PaymentMethodId             string         `json:"paymentMethodId,omitempty"`
	PaymentMethodMetaData       map[string]any `json:"paymentMethodMetaData,omitempty"`
	CustomerId                  string         `json:"customerId,omitempty"`
	ExtendInfo                  string         `json:"extendInfo,omitempty"`
	RequireIssuerAuthentication bool           `json:"requireIssuerAuthentication,omitempty"`
	Funding                     FundingType    `json:"funding,omitempty"`
}

type FundingType string

const (
	FundingType_CREDIT         FundingType = "CREDIT"
	FundingType_DEBIT          FundingType = "DEBIT"
	FundingType_PREPAID        FundingType = "PREPAID"
	FundingType_CHARGE         FundingType = "CHARGE"
	FundingType_DEFERRED_DEBIT FundingType = "DEFERRED_DEBIT"
)
