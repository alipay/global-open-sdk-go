package model

type CouponPaymentMethodDetail struct {
	CouponId                    string  `json:"couponId,omitempty"`
	AvailableAmount             *Amount `json:"availableAmount,omitempty"`
	CouponName                  string  `json:"couponName,omitempty"`
	CouponDescription           string  `json:"couponDescription,omitempty"`
	CouponExpireTime            string  `json:"couponExpireTime,omitempty"`
	PaymentMethodDetailMetadata string  `json:"paymentMethodDetailMetadata,omitempty"`
}
