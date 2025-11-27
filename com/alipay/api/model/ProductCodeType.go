package model

type ProductCodeType string

const (
	ProductCodeType_CASHIER_PAYMENT   ProductCodeType = "CASHIER_PAYMENT"
	ProductCodeType_AGREEMENT_PAYMENT ProductCodeType = "AGREEMENT_PAYMENT"
	ProductCodeType_IN_STORE_PAYMENT  ProductCodeType = "IN_STORE_PAYMENT"
)
