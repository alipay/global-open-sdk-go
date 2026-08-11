package model

type CouponInquireDetailsAppliesTo struct {
	Products []*CouponApplicableProduct `json:"products,omitempty"`
}
