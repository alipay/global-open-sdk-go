package model

type PaymentMethodDetail struct {
	PaymentMethodDetailType PaymentMethodDetailType      `json:"paymentMethodDetailType,omitempty"`
	Card                    *CardPaymentMethodDetail     `json:"card,omitempty"`
	ExternalAccount         *ExternalPaymentMethodDetail `json:"externalAccount,omitempty"`
	Discount                *DiscountPaymentMethodDetail `json:"discount,omitempty"`
	Coupon                  *CouponPaymentMethodDetail   `json:"coupon,omitempty"`
	PaymentMethodType       string                       `json:"paymentMethodType,omitempty"`
	ExtendInfo              string                       `json:"extendInfo,omitempty"`
	Wallet                  *Wallet                      `json:"wallet,omitempty"`
	InteractionType         string                       `json:"interactionType,omitempty"`
}
