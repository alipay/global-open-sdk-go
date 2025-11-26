package model

type PaymentMethodDetailType string

const (
	PaymentMethodDetailType_CARD            PaymentMethodDetailType = "CARD"
	PaymentMethodDetailType_EXTERNALACCOUNT PaymentMethodDetailType = "EXTERNALACCOUNT"
	PaymentMethodDetailType_COUPON          PaymentMethodDetailType = "COUPON"
	PaymentMethodDetailType_DISCOUNT        PaymentMethodDetailType = "DISCOUNT"
)
