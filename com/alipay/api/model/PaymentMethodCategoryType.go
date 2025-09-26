package model

type PaymentMethodCategoryType string

const (
	PaymentMethodCategoryType_ALIPAY_PLUS        PaymentMethodCategoryType = "ALIPAY_PLUS"
	PaymentMethodCategoryType_WALLET             PaymentMethodCategoryType = "WALLET"
	PaymentMethodCategoryType_MOBILE_BANKING_APP PaymentMethodCategoryType = "MOBILE_BANKING_APP"
	PaymentMethodCategoryType_BANK_TRANSFER      PaymentMethodCategoryType = "BANK_TRANSFER"
	PaymentMethodCategoryType_ONLINE_BANKING     PaymentMethodCategoryType = "ONLINE_BANKING"
	PaymentMethodCategoryType_CARD               PaymentMethodCategoryType = "CARD"
	PaymentMethodCategoryType_OTC                PaymentMethodCategoryType = "OTC"
)
