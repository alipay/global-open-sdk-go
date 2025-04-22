package model

type PaymentMethodType string

const (
	PaymentMethodType_DISCOUNT                 PaymentMethodType = "DISCOUNT"
	PaymentMethodType_INTEREST_FREE            PaymentMethodType = "INTEREST_FREE"
	PaymentMethodType_BALANCE_ACCOUNT          PaymentMethodType = "BALANCE_ACCOUNT"
	PaymentMethodType_SETTLEMENT_CARD          PaymentMethodType = "SETTLEMENT_CARD"
	PaymentMethodType_APPLEPAY                 PaymentMethodType = "APPLEPAY"
	PaymentMethodType_UPI                      PaymentMethodType = "UPI"
	PaymentMethodType_ONLINEBANKING_NETBANKING PaymentMethodType = "ONLINEBANKING_NETBANKING"
)
