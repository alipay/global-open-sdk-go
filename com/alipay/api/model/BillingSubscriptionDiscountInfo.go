package model

type BillingSubscriptionDiscountInfo struct {
	DiscountId string  `json:"discountId,omitempty"`
	CouponId   string  `json:"couponId,omitempty"`
	Type       string  `json:"type,omitempty"`
	PercentOff int32   `json:"percentOff,omitempty"`
	AmountOff  *Amount `json:"amountOff,omitempty"`
	Duration   string  `json:"duration,omitempty"`
	Times      int32   `json:"times,omitempty"`
	Status     string  `json:"status,omitempty"`
}
